const fs = require('fs');
const file = 'src/Gopurs/CodeGen.purs';
let code = fs.readFileSync(file, 'utf8');
code = code.replace(
  `                      let funcName = "Rebox_" <> modNameStr <> "_" <> hashString srcFullPath <> "_" <> hashString destFullPath`,
  `                      let funcName = "Rebox_" <> modNameStr <> "_" <> hashString srcFullPath <> "_" <> hashString destFullPath\n                      _ = Debug.trace ("REBOX_DUP_CHECK: func=" <> funcName <> " srcT=" <> (goTypeToStr srcT) <> " destT=" <> (goTypeToStr destT) <> " srcFull=" <> srcFullPath <> " destFull=" <> destFullPath) \\_ -> unit`
);
fs.writeFileSync(file, code);
