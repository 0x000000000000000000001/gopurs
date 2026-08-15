const fs = require('fs');
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

content = content.replace(
  /structFieldGoType ptrPaths enumAdts elidedCtors typeVars modStr ty = \n  case exprTypeToGenericGoType ptrPaths enumAdts typeVars modStr ty of/g,
  'structFieldGoType ptrPaths enumAdts elidedCtors typeVars modStr ty = \n  case exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modStr ty of'
);

content = content.replace(
  /structFieldGoType pointerAdtPaths enumAdts decl\.vars modNameStr/g,
  'structFieldGoType pointerAdtPaths enumAdts elidedCtors decl.vars modNameStr'
);

content = content.replace(
  /structFieldGoType \(unsafePerformEffect \(Ref\.read helpersRef\)\)\.pointerAdtPaths \(unsafePerformEffect \(Ref\.read helpersRef\)\)\.enumAdts ctorInfo\.vars modNameStr/g,
  'structFieldGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors ctorInfo.vars modNameStr'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
