const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(/GoMutate paramName \(unboxGoExpr argExpr argType expectedType\)/g, 'GoMutate paramName (coerceGoExpr modNameStr argExpr argType expectedType)');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
