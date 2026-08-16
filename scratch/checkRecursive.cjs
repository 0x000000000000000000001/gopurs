const fs = require('fs');
const json = JSON.parse(fs.readFileSync('tests/runner/output/Data.String.CodeUnits/corefn.json', 'utf8'));
const decl = json.decls.find(d => d.identifier === 'contains' || d.binds?.some(b => b.identifier === 'contains'));
console.log(decl.binds ? 'recursive' : 'non-recursive');
