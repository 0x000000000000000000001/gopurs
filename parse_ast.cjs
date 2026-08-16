const fs = require('fs');
const code = fs.readFileSync('tests/runner/output/purescript/Data_String_CodeUnits.go', 'utf8');
const lines = code.split('\n');
const start = lines.findIndex(l => l.includes('func Call_Data_String_CodeUnits_contains'));
const end = lines.findIndex((l, i) => i > start && l.startsWith('}'));
console.log(lines.slice(start, end+1).join('\n'));
