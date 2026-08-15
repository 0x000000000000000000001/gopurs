const fs = require('fs');
const json = JSON.parse(fs.readFileSync('ntco3.json', 'utf8'));
console.log(JSON.stringify(json.expression.expression.binds[0].expression.abstraction.expression.abstraction.expression.binds[0].annotation, null, 2));
