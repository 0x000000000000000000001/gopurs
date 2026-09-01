const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(/import Debug \(traceM\)\n/, '');
if (!code.includes('import Debug')) {
  code = code.replace(/import Prelude\n/, 'import Prelude\nimport Debug (traceM)\n');
}

code = code.replace(/coerceGoExpr modNameStr expr srcT@\(TypeStructPointer b1 f1 s1 a1\) destT@\(TypeStructPointer b2 f2 s2 a2\) \| b1 == b2 =\n  let\n    _trace = unsafePerformEffect \(traceM \("coerceGoExpr Rebox match: " <> goTypeToStr srcT <> " -> " <> goTypeToStr destT\)\)\n\n  let\n    srcPath = s1/, `coerceGoExpr modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) | b1 == b2 =
  let
    _trace = unsafePerformEffect (traceM ("coerceGoExpr Rebox match: " <> goTypeToStr srcT <> " -> " <> goTypeToStr destT))
    srcPath = s1`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
