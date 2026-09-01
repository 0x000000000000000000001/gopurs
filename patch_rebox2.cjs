const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(/"func " <> funcName <> "\(in \*" <> f1 <> "\["/g, '"func " <> funcName <> "(in *" <> b1 <> "["');
code = code.replace(/\]\) \*" <> f2 <> "\["/g, ']) *" <> b2 <> "["');
code = code.replace(/out := \&" <> f2 <> "\["/g, 'out := &" <> b2 <> "["');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
