const fs = require('fs');
const contents = fs.readFileSync('output/Data.String.CodeUnits/corefn.json');
const tcorefn = JSON.parse(contents);
const decls = tcorefn.decls;
const containsDecl = decls.find(d => d.bindType === 'NonRec' && d.identifier === 'contains');
console.dir(containsDecl, { depth: null });
