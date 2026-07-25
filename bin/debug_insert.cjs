const fs = require('fs');
const ast = JSON.parse(fs.readFileSync('/Users/0x1/Documents/htdocs/altbak.pub/output/Data.Map.Internal/corefn.json', 'utf8'));

let insert = null;
for (const d of ast.decls) {
  if (d.binds) {
    for (const b of d.binds) {
      if (b.identifier === 'insert') {
        insert = b;
        break;
      }
    }
  } else if (d.identifier === 'insert') {
    insert = d;
    break;
  }
}
if (insert) {
  console.log(JSON.stringify(insert.expression.annotation, null, 2));
} else {
  console.log("insert not found in decls");
}
