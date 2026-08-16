const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('output/Main/corefn.json', 'utf8'));
const decls = corefn.decls;
const rec = decls.find(d => d.binds);
const bravo = rec.binds.find(b => b.identifier === 'bravo');
console.log(JSON.stringify(bravo, null, 2));
