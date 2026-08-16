const fs = require('fs');
const tast = JSON.parse(fs.readFileSync('tests/runner/output/Data.String.CodeUnits/corefn.json'));
const contains = tast.decls.find(d => d.identifier === 'contains');
console.dir(contains, { depth: null });
