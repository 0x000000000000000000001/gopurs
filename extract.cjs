const fs = require('fs');
const json = JSON.parse(fs.readFileSync('tests/runner/output/Main/corefn.json', 'utf8'));
const functor = json.decls.find(d => d.identifier === 'functorM' || (d.binds && d.binds.some(b => b.identifier === 'functorM')));
console.log(JSON.stringify(functor, null, 2));
