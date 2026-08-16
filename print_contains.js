const fs = require('fs');
const contents = fs.readFileSync('tests/runner/output/purescript/Data.String.CodeUnits/tcorefn.json');
const tcorefn = JSON.parse(contents);
const decls = tcorefn.decls;
const containsDecl = decls.find(d => d.bindType === 'NonRec' && d.ident === 'contains');
console.dir(containsDecl, { depth: null });
