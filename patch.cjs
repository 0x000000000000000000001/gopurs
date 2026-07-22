const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// Update the type signature
code = code.replace(
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }',
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }'
);

// Update the function definition
code = code.replace(
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =',
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail inEffect nextId tcoExpr@(TcoExpr tcoAnalysis expr) ='
);

const parts = code.split(/(translateExprImpl\s+[\s\S]*?nextId\s+.*?tcoExpr)/g);

// Wait, doing this carefully: let's replace all `translateExprImpl ` correctly.
// There are exactly 42 occurrences of `translateExprImpl `.
// The first two are the signature and definition, which we already replaced.
// For the remaining 40 occurrences, we just need to insert `false ` before `nextId` or whatever the nextId argument is.
// Actually, it's safer to just split by `translateExprImpl ` and modify the rest manually or via script.
