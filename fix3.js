import fs from 'fs';

let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const newSig = `
translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr _ expr) =
  let
    isEffectNode = case expr of
      EffectBind _ _ _ _ -> true
      EffectPure _ -> true
      EffectDefer _ -> true
      PrimEffect _ -> true
      _ -> false
  in if isEffectNode then
    let res = translateExprImplWorker true helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr
    in { stmts: StmtEmpty, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(_ gopurs_runtime.Value) gopurs_runtime.Value {\\n" + "return " + printGoExpr (GoBlock (res.stmts `flattenStmts` res.stmts /* wait, flattenStmts takes StmtTree */)) + "\\n}") ], nextId: res.nextId }
  else
    translateExprImplWorker false helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr

translateExprImplWorker :: Boolean -> Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImplWorker inEffect helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =`;

// The above is pseudocode. Let's write the actual PureScript code for the wrapper!
