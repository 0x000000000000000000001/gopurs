import fs from 'fs';
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const isEffectNodeSrc = `
isEffectNode :: forall a. BackendSyntax a -> Boolean
isEffectNode = case _ of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> true
  PrimEffect _ -> true
  UncurriedEffectApp _ _ -> true
  _ -> false

executeIfOpaque :: forall a. BackendSyntax a -> GoExpr -> GoExpr
executeIfOpaque expr goExpr =
  if isEffectNode expr then goExpr
  else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ goExpr, GoRaw "gopurs_runtime.Value{}" ]
`;

// Insert right before translateExprImpl ::
content = content.replace(
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }',
  isEffectNodeSrc + '\ntranslateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }'
);

// 1. Rename ALL `translateExprImpl helpersRef` to `translateExprImpl_ helpersRef ... false ...`
content = content.replace(/translateExprImpl helpersRef (.*?) (false|true|isTail) (nextId[a-zA-Z']*|res[a-zA-Z]*\.nextId|\(nextId \+ 1\)|currNextId|0|acc[a-zA-Z]*\.nextId) /g,
  "translateExprImpl_ helpersRef $1 $2 false $3 ");

// 2. We need to replace the recursive calls inside EffectBind to pass true!
content = content.replace(
  /EffectBind mbIdent lvl binding body ->\n(\s+)let\n(\s+)originalName = localId mbIdent lvl\n(\s+)name = originalName <> "_" <> show nextId\n(\s+)newBound = Map.insert originalName name bound\n(\s+)resBinding = translateExprImpl_ helpersRef \(depth \+ 1\) modNameStr recVars namedBound bound Nothing \[\] false false \(nextId \+ 1\) binding\n(\s+)resBody = translateExprImpl_ helpersRef \(depth \+ 1\) modNameStr recVars namedBound newBound Nothing loopCtx isTail false resBinding\.nextId body/g,
  `EffectBind mbIdent lvl binding body ->\n$1let\n$2originalName = localId mbIdent lvl\n$3name = originalName <> "_" <> show nextId\n$4newBound = Map.insert originalName name bound\n$5resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false true (nextId + 1) binding\n$6resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail true resBinding.nextId body`
);

// 3. Define the wrapper
let wrapper = `
translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr =
  translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail false nextId tcoExpr

translateExprImpl_ :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) =
  let
    isEff = isEffectNode expr
  in
    if isEff && not inEffectBlock then
      let
        res = translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx false true nextId tcoExpr
        funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\\n" <> printGoExpr (GoBlock (flattenStmts res.stmts <> [ GoReturn res.expr ])) <> "\\n})")
      in
        { stmts: StmtEmpty, expr: funcExpr, nextId: res.nextId }
    else
`;

// Also, the original translateExprImpl definition line needs to be replaced.
// Since we didn't replace `translateExprImpl helpersRef depth modNameStr recVars` yet (it wasn't matched because the signature line has a different format, it has `=` at the end and no `expr`).
// Wait, the match was `/translateExprImpl helpersRef (.*?) (false|true|isTail) (nextId[a-zA-Z']*|res[a-zA-Z]*\.nextId|\(nextId \+ 1\)|currNextId|0|acc[a-zA-Z]*\.nextId) /g`
// The definition line is:
// `translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =`
// This line does NOT have `(false|true|isTail)` followed by `nextId `, it has `isTail nextId tcoExpr`.
// So it was NOT matched!
// So we can just replace the definition line!
content = content.replace(
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }\n' +
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =',
  wrapper
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
