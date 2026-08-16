const fs = require('fs');
let code = fs.readFileSync('tests/runner/output/purescript/Data_String_CodeUnits.go', 'utf8');
const start = code.indexOf('func Call_Data_String_CodeUnits_contains');
const end = code.indexOf('}', code.indexOf('func Get_Data_String_CodeUnits__charAt'));
code = code.substring(start, end);
console.log(code);
