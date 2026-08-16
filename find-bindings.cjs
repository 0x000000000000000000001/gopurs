const fs = require('fs');
const json = JSON.parse(fs.readFileSync('functorM.json', 'utf8'));

function walk(node, path) {
  if (!node) return;
  if (Array.isArray(node)) {
    node.forEach((n, i) => walk(n, path + '[' + i + ']'));
    return;
  }
  if (typeof node === 'object') {
    if (node.type === 'Let') {
      console.log('Let binding found at', path);
      node.binds.forEach(b => {
        console.log('  Identifier:', b.identifier, 'Type:', b.expression.type);
        if (b.expression.type === 'App') {
          console.log('    App Function:', b.expression.abstraction.type, b.expression.abstraction.value?.identifier);
        }
      });
    }
    Object.keys(node).forEach(k => walk(node[k], path + '.' + k));
  }
}

walk(json.expression || json, 'root');
