const fs = require('fs');
const json = JSON.parse(fs.readFileSync('ntco3.json', 'utf8'));
const finds = [];
function search(obj) {
  if (!obj) return;
  if (obj.identifier === 'g') finds.push(obj);
  if (typeof obj === 'object') {
    Object.values(obj).forEach(search);
  }
}
search(json);
console.log(JSON.stringify(finds[0], null, 2));
