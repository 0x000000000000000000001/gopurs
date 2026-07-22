const fs = require('fs');

let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }',
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }'
);

code = code.replace(
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =',
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail inEffect nextId tcoExpr@(TcoExpr tcoAnalysis expr) ='
);

// We need to replace all calls. Most calls look like:
// translateExprImpl helpersRef ... nextId expr
// Let's find all occurrences of `translateExprImpl` and see if we can insert `false` (or `inEffect`).

const lines = code.split('\n');
for (let i = 0; i < lines.length; i++) {
  if (lines[i].includes('translateExprImpl') && !lines[i].includes('::') && !lines[i].includes('tcoExpr@')) {
    // If it's a call, we replace the `nextId` argument with `false nextId`.
    // We know nextId is always named `nextId` or `accNext.nextId` or `resCond.nextId` etc.
    // It is always the 10th argument, and is followed by the expression.
    // Since naming conventions are consistent, let's just replace `nextId` with `inEffect nextId` where it is passed as an argument.
    // Actually, `inEffect` should default to `false`, EXCEPT inside `Effect` computations where it is `true`.
  }
}
