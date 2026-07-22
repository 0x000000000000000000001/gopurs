import fs from 'fs';
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
let matches = content.match(/translateExprImpl helpersRef (.*?) (false|true|isTail) (nextId[a-zA-Z']*|res[a-zA-Z]*\.nextId|\(nextId \+ 1\)|currNextId|0|acc[a-zA-Z]*\.nextId) /g);
console.log("Matched: " + (matches ? matches.length : 0));
