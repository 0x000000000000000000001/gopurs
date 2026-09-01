const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const targetStr1 = `GoMutate paramName (unboxGoExpr argExpr argType expectedType)`;
const replacementStr1 = `GoMutate paramName (coerceGoExpr modNameStr argExpr argType expectedType)`;

// We replace ALL occurrences.
code = code.split(targetStr1).join(replacementStr1);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
console.log("Patched CodeGen.purs successfully.");
