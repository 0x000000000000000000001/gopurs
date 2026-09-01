const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(/coerceGoExpr :: String -> GoExpr -> GoType -> GoType -> GoExpr/, 'import Debug (traceM)\ncoerceGoExpr :: String -> GoExpr -> GoType -> GoType -> GoExpr');

code = code.replace(/coerceGoExpr modNameStr expr srcT@\(TypeStructPointer b1 f1 s1 a1\) destT@\(TypeStructPointer b2 f2 s2 a2\) \| b1 == b2 =/g, `coerceGoExpr modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) | b1 == b2 =
  let
    _trace = unsafePerformEffect (traceM ("coerceGoExpr Rebox match: " <> goTypeToStr srcT <> " -> " <> goTypeToStr destT))
`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
