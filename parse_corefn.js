const fs = require('fs');
const json = JSON.parse(fs.readFileSync('tests/runner/output/Main/corefn.json'));
const main = json.decls.find(d => d.identifier === 'main');
fs.writeFileSync('main_corefn.json', JSON.stringify(main, null, 2));
