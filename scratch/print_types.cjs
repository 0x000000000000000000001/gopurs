const fs = require('fs');
const json = JSON.parse(fs.readFileSync('../../altbak.pub/output/Control.Comonad.Env.Class/corefn.json', 'utf-8'));
console.log(Object.keys(json));
