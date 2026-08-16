const fs = require('fs');
const json = JSON.parse(fs.readFileSync('tests/runner/output/Data.String.CodeUnits/tcorefn.json', 'utf8'));
const contains = json.decls.find(d => d.identifier === 'contains');
console.log(JSON.stringify(contains, null, 2));
