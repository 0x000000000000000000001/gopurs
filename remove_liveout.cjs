const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  'Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Set String -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }',
  'Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }'
);

code = code.replace(
  'translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx liveOut isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) =',
  'translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) ='
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
