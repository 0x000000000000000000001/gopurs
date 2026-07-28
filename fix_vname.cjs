const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /then Just { name: vname, goType: v\.goType }/,
  'then Just { name: v.name, goType: v.goType }'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
