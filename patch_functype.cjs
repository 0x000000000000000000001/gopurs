const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(/exprTypeToGoType ptrPaths enumAdts elided modNameStr \(Func args ret\) = TypeFunc \(map \(exprTypeToGoType ptrPaths enumAdts elided modNameStr\) args\) \(exprTypeToGoType ptrPaths enumAdts elided modNameStr ret\)/g, 
  'exprTypeToGoType ptrPaths enumAdts elided modNameStr (Func args ret) = TypeValue');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
