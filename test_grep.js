const fs = require('fs');
const content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
const lines = content.split('\n');
lines.forEach((line, i) => {
  if (line.includes('exprTypeToGoType') && line.includes('Func')) {
    console.log((i+1) + ': ' + line);
  }
});
