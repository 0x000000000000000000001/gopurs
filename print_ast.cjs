const fs = require('fs');
const json = JSON.parse(fs.readFileSync('tests/runner/output/Data.String.CodeUnits/corefn.json', 'utf8'));
const contains = json.decls.find(d => (d.identifier || d.binds?.[0]?.identifier) === 'contains');
console.log(JSON.stringify(contains, null, 2));
