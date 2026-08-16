const fs = require('fs');
const data = JSON.parse(fs.readFileSync(process.argv[2], 'utf8'));

function findArrayMap(node) {
  if (!node) return;
  if (Array.isArray(node)) {
    for (const n of node) findArrayMap(n);
    return;
  }
  if (typeof node === 'object') {
    if (node.type === 'App' && node.abstraction && node.abstraction.type === 'Var' && node.abstraction.value.identifier === 'arrayMap') {
      console.log("FOUND APP arrayMap!");
      console.log("Type annotation on App:", JSON.stringify(node.ann.type, null, 2));
    }
    for (const key in node) findArrayMap(node[key]);
  }
}

for (const decl of data.decls) {
  findArrayMap(decl);
}
