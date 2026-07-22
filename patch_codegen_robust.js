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

const wrapperSrc = `
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

// 1. Rename the definition signature and first line.
content = content.replace(
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }\n' +
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =\n',
  isEffectNodeSrc + '\n' + wrapperSrc
);

// 2. Rename all calls `translateExprImpl helpersRef` to `translateExprImpl_ helpersRef ... false ...`
content = content.replace(/translateExprImpl helpersRef (.*?) (false|true|isTail) (nextId[a-zA-Z']*|res[a-zA-Z]*\.nextId|\(nextId \+ 1\)|currNextId|0|acc[a-zA-Z]*\.nextId) /g,
  "translateExprImpl_ helpersRef $1 $2 false $3 ");

// 3. Fix EffectBind recursive calls to pass `true`
content = content.replace(
  /EffectBind mbIdent lvl binding body ->\n(\s+)let\n(\s+)originalName = localId mbIdent lvl\n(\s+)name = originalName <> "_" <> show nextId\n(\s+)newBound = Map.insert originalName name bound\n(\s+)resBinding = translateExprImpl_ helpersRef \(depth \+ 1\) modNameStr recVars namedBound bound Nothing \[\] false false \(nextId \+ 1\) binding\n(\s+)resBody = translateExprImpl_ helpersRef \(depth \+ 1\) modNameStr recVars namedBound newBound Nothing loopCtx isTail false resBinding\.nextId body/g,
  `EffectBind mbIdent lvl binding body ->\n$1let\n$2originalName = localId mbIdent lvl\n$3name = originalName <> "_" <> show nextId\n$4newBound = Map.insert originalName name bound\n$5resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false true (nextId + 1) binding\n$6resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail true resBinding.nextId body`
);

// 4. Update the `liftIfNeeded` helper call which was not matched by the regex because it's `newNextId tcoExpr` (no space after tcoExpr?)
// Wait! `newNextId` doesn't end with `nextId ` if it's `newNextId tcoExpr` at the end of the line.
// Let's replace it manually:
content = content.replace(
  'let res = translateExprImpl helpersRef 0 modNameStr recVars namedBound bound Nothing [] false newNextId tcoExpr',
  'let res = translateExprImpl_ helpersRef 0 modNameStr recVars namedBound bound Nothing [] false inEffectBlock newNextId tcoExpr'
);

// Also we need to check if there are any other `translateExprImpl` left!
fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
