const fs = require('fs');
const lines = fs.readFileSync('tests/runner/output/purescript/Data_String_CodeUnits.go', 'utf8').split('\n');
console.log(lines.slice(418, 435).join('\n'));
