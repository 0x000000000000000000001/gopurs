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

// Find all calls to translateExprImpl and insert `false` (for inEffect) before nextId
// A call looks like: translateExprImpl arg1 arg2 arg3 arg4 arg5 arg6 arg7 arg8 arg9 nextIdArg exprArg
// It might be split across lines, but in CodeGen.purs it's usually on one line.
// Let's use a regex to match the first 9 arguments (from helpersRef to isTail equivalent).
// Actually, it's easier to match `translateExprImpl ` and the 9 arguments.
const callRegex = /translateExprImpl\s+([a-zA-Z0-9_.]+)\s+(\([^)]+\)|[a-zA-Z0-9_.]+)\s+([a-zA-Z0-9_.]+)\s+([a-zA-Z0-9_.]+)\s+([a-zA-Z0-9_.]+)\s+([a-zA-Z0-9_.]+)\s+([a-zA-Z0-9_.]+)\s+([a-zA-Z0-9_\[\]]+)\s+([a-zA-Z0-9_.]+)/g;

code = code.replace(callRegex, (match) => {
  return match + ' false';
});

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
