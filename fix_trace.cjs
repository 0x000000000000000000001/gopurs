const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
`  App fn arg -> "App(" <> printTcoExprShape fn <> " " <> printTcoExprShape arg <> ")"`,
`  App fn arg -> "App(" <> printTcoExprShape fn <> ")"`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
