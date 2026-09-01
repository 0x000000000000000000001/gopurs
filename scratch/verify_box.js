const fs = require('fs');

const path = 'src/Gopurs/CodeGen.purs';
let content = fs.readFileSync(path, 'utf8');
const lines = content.split('\n');

const coerceStart = lines.findIndex(l => l.startsWith('coerceGoExpr '));
console.log(lines.slice(coerceStart, coerceStart + 25).join('\n'));
