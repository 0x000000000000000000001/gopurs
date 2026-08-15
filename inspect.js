const fs = require('fs');
const content = fs.readFileSync('output/Middle/Effect.json', 'utf8');
const data = JSON.parse(content);
const functorEffect = data.decls.find(d => d.ident === 'functorEffect');
console.log(JSON.stringify(functorEffect, null, 2));
