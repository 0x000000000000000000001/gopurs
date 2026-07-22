import fs from 'fs';
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

content = content.replace(
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }\n' +
  'translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr =\n',
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }\n' +
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr =\n'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
