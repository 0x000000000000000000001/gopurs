const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('tests/runner/output/Data.String.CodeUnits/tcorefn.json', 'utf8'));
const contains = corefn.decls.find(d => d.identifier === 'contains');
console.log(JSON.stringify(contains, null, 2));
