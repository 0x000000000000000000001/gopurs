const fs = require('fs');
const contents = fs.readFileSync('tests/passing/4179.purs');
console.log(contents.toString());
