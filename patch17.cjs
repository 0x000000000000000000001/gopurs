const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace("exprTypeToGoType ptrPaths enumAdts elided modNameStr (Array ty) = TypeNativeArray (exprTypeToGoType ptrPaths enumAdts elided modNameStr ty)\\n", "exprTypeToGoType ptrPaths enumAdts elided modNameStr (Array ty) = TypeNativeArray (exprTypeToGoType ptrPaths enumAdts elided modNameStr ty)\nexprTypeToGoType ptrPaths enumAdts elided modNameStr (Func args ret) = TypeFunc (map (exprTypeToGoType ptrPaths enumAdts elided modNameStr) args) (exprTypeToGoType ptrPaths enumAdts elided modNameStr ret)\n");

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
