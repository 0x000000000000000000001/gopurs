const fs = require('fs');
const lines = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8').split('\n');
const start = lines.findIndex(l => l.includes('expectedRetType = case extractExprFuncType'));
console.log(lines.slice(start - 2, start + 10).join('\n'));
