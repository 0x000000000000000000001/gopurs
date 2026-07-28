const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  'deadVarOpt = deadVarOptRaw',
  'deadVarOpt = Nothing'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
