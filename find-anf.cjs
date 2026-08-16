const fs = require('fs');

function search(dir) {
  for (const file of fs.readdirSync(dir)) {
    const full = dir + '/' + file;
    if (fs.statSync(full).isDirectory()) {
      search(full);
    } else if (full.endsWith('corefn.json')) {
      const data = JSON.parse(fs.readFileSync(full, 'utf8'));
      findArrayMap(data, full);
    }
  }
}

function findArrayMap(node, file) {
  if (!node) return;
  if (Array.isArray(node)) {
    for (const n of node) findArrayMap(n, file);
    return;
  }
  if (typeof node === 'object') {
    if (node.type === 'App' && node.abstraction && node.abstraction.type === 'Var' && node.abstraction.value.identifier === 'arrayMap') {
      console.log("FOUND in", file);
      console.log("App type:", JSON.stringify(node.ann.type));
    }
    for (const key in node) findArrayMap(node[key], file);
  }
}

search('tests/runner/output');
