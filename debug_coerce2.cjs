const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(/coerceGoExpr modNameStr expr from to = unboxGoExpr \(boxGoExpr expr from\) TypeValue to/, `coerceGoExpr modNameStr expr from to =
  let
    _trace = unsafePerformEffect (traceM ("coerceGoExpr FALLBACK: " <> goTypeToStr from <> " -> " <> goTypeToStr to))
  in unboxGoExpr (boxGoExpr expr from) TypeValue to`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
