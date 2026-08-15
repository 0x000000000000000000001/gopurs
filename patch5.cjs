const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(/CtorDef _ _ _ _ _/, "CtorDef _ _ _ _");
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
