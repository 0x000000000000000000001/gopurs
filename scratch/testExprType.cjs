const fs = require('fs');
const json = JSON.parse(fs.readFileSync('tests/runner/output/Data.String.CodeUnits/corefn.json', 'utf8'));
const contains = json.decls.find(d => d.identifier === 'contains');
console.log(JSON.stringify(contains.expression.annotation.type, null, 2));
