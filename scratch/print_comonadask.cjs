const fs = require('fs');
const json = JSON.parse(fs.readFileSync('../../altbak.pub/output/Control.Comonad.Env.Class/corefn.json', 'utf-8'));

const c = json.classDecls.find(x => x.name === 'ComonadAsk');
console.log(JSON.stringify(c, null, 2));
