const fs = require('fs');
const ast = JSON.parse(fs.readFileSync('../altbak.pub/output/Main/corefn.json', 'utf8'));
function findLog(node) {
  if (Array.isArray(node)) return node.map(findLog).find(x => x);
  if (typeof node !== 'object' || node === null) return null;
  if (node.type === 'Var' && node.value && node.value.identifier === 'log') {
    return node;
  }
  for (let key in node) {
    let res = findLog(node[key]);
    if (res) return res;
  }
  return null;
}
console.log(JSON.stringify(findLog(ast), null, 2));
