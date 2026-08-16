const fs = require('fs');
const path = require('path');
const corefn = JSON.parse(fs.readFileSync('output/Data.String.CodeUnits/tcorefn.json', 'utf8'));
const contains = corefn.decls.find(d => d.identifier === 'contains');
console.log(JSON.stringify(contains, null, 2));
