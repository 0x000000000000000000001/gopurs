const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(`            isAppNode = case inner of\\n              App _ _ -> true\\n              _ -> false`, `            isAppNode = case unwrapExpr inner of\\n              App _ _ -> true\\n              _ -> false`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
