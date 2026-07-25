const fs = require('fs');
const ast = JSON.parse(fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs-test/output/TestMonomorphize/corefn.json', 'utf8'));

let insert = null;
for (const d of ast.decls) {
  if (d.binds) {
    for (const b of d.binds) {
      if (b.identifier === 'myIdentity') {
        insert = b;
        break;
      }
    }
  } else if (d.identifier === 'myIdentity') {
    insert = d;
    break;
  }
}
if (insert) {
  console.log(JSON.stringify(insert.expression.annotation.type, null, 2));
} else {
  console.log("myIdentity not found in decls");
}
