const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(`unboxGoExpr expr currentType desiredType = case desiredType of`, `unboxGoExpr expr currentType desiredType = case desiredType of\\n    (TypeFunc args ret) -> case currentType of\\n      TypeFunc currentArgs currentRet -> expr\\n      _ -> expr -- Wait! I need to see what unboxGoExpr does for TypeFunc!`);

// Actually, let's just dump unboxGoExpr for TypeFunc
