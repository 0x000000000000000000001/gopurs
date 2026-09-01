const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

// 1. Update the signature
code = code.replace(
  'coerceGoExpr :: GoExpr -> GoType -> GoType -> GoExpr',
  'coerceGoExpr :: String -> GoExpr -> GoType -> GoType -> GoExpr'
);

// 2. Add modNameStr to all parameter lists
code = code.replace(/coerceGoExpr expr/g, 'coerceGoExpr modNameStr expr');

// 3. Update all CALLS to coerceGoExpr to pass modNameStr
code = code.replace(/coerceGoExpr \(/g, 'coerceGoExpr modNameStr (');
code = code.replace(/coerceGoExpr res/g, 'coerceGoExpr modNameStr res');
code = code.replace(/coerceGoExpr p/g, 'coerceGoExpr modNameStr p');
code = code.replace(/coerceGoExpr exprAccess/g, 'coerceGoExpr modNameStr exprAccess');
code = code.replace(/coerceGoExpr \$/g, 'coerceGoExpr modNameStr $');
code = code.replace(/coerceGoExpr GoVar/g, 'coerceGoExpr modNameStr GoVar');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
console.log("Patched coerceGoExpr calls!");
