const fs = require('fs');
const ast = JSON.parse(fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs-test/output/Data.Map.Internal/corefn.json', 'utf8'));
const insert = ast.decls.find(d => d.binds && d.binds.some(b => b.identifier === 'insert'));
if (insert) {
  const b = insert.binds.find(b => b.identifier === 'insert');
  console.log(JSON.stringify(b.expression.annotation.type, null, 2));
} else {
  console.log("insert not found in decls");
}
