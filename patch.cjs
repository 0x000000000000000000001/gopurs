const fs = require('fs');
const file = 'src/Gopurs/CodeGen.purs';
let code = fs.readFileSync(file, 'utf8');

// Find the arrayMap case and insert logging for the RAW EXPRESSIONS of fExpr and arrExpr!
// Actually, fExpr is not a TcoExpr, it's just the translated GoExpr!
// Wait, the arguments to flattenApp are in flatArgs!
// flatArgs is an Array of AppArgs.
// I can just print the TcoExpr of flatArgs!
// But how?
// I can use `Debug.trace`!
