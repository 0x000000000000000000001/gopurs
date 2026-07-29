const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('../altbak.pub/output/TestDump/corefn.json', 'utf8'));
const main = corefn.decls.find(d => d.identifier === 'main');
console.log(JSON.stringify(main, null, 2));
