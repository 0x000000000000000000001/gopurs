const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(/_trace = Debug\.trace \(([^)]+)\) \\_ -> unit\n\s*funcStr =/g, 'funcStr = Debug.trace ($1) \\_ ->');
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
