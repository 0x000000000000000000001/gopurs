const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /GoCall \(GoRaw \("func\(\) gopurs_runtime\.Value"\)\) \[ GoRaw \("{\\n" <> printGoExpr checkBlock <> "\\nreturn " <> printGoExpr allocBlock <> "\\n}\(\)"\) \]/,
  'GoRaw ("func() gopurs_runtime.Value {\\n" <> printGoExpr checkBlock <> "\\nreturn " <> printGoExpr allocBlock <> "\\n}()")'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
