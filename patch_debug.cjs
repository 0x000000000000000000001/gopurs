const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const targetStr = `assigns = Array.zipWith (\\(Tuple p goT) argExpr -> GoMutate p (coerceGoExpr modNameStr argExpr.expr argExpr.exprType goT)) paramsWithTypes argExprs`;
const replacementStr = `assigns = Array.zipWith (\\(Tuple p goT) argExpr -> GoMutate p (Debug.trace ("GoMutate coercion for " <> p <> ": from " <> goTypeToStr argExpr.exprType <> " to " <> goTypeToStr goT) \\_ -> coerceGoExpr modNameStr argExpr.expr argExpr.exprType goT)) paramsWithTypes argExprs`;

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.replace(targetStr, replacementStr));
console.log("Patched CodeGen.purs successfully.");
