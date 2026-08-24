const fs = require('fs');
const data = JSON.parse(fs.readFileSync('/Users/0x1/Documents/htdocs/altbak.pub/output/Data.Enum/corefn.json', 'utf8'));
const adt = data.decls.find(d => {
  if (d.bindType === 'NonRec' && d.identifier === 'Cardinality') return true;
  return false;
});
console.log("Cardinality exists?", data.dataDecls.some(d => d.typeName === 'Cardinality'));
