import fs from 'fs';
const data = JSON.parse(fs.readFileSync('ntco3.json'));

let g_bind;
function walk(obj) {
  if (!obj) return;
  if (obj.identifier === 'g' && obj.expression) {
    g_bind = obj;
  }
  if (typeof obj === 'object') {
    for (const k in obj) walk(obj[k]);
  }
}
walk(data);
console.log("g_bind.expression.annotation.type =", JSON.stringify(g_bind.expression.annotation.type, null, 2));
