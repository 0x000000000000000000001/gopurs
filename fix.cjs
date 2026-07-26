const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/GoAst.purs', 'utf8');
code = code.replace(/  \| GoRecordUpdateNative GoExpr Int \(Array \(Tuple Int GoExpr\)\)\n/g, '');
code = code.replace(/  \| GoRecordAccessNative GoExpr Int Int\n/g, '');
fs.writeFileSync('src/Gopurs/GoAst.purs', code);
