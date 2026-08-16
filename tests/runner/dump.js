const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('output-es/Main/corefn.json', 'utf8'));
const decls = corefn.decls;
const alpha = decls.find(d => d.identifier === 'alpha');
console.log(JSON.stringify(alpha, null, 2));
