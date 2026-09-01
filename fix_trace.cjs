const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(/import Debug \(traceM\)\n/, 'import Effect.Console as Console\n');
code = code.replace(/_trace = unsafePerformEffect \(traceM \("coerceGoExpr Rebox match: " <> goTypeToStr srcT <> " -> " <> goTypeToStr destT\)\)/, '_trace = unsafePerformEffect (Console.log ("coerceGoExpr Rebox match: " <> goTypeToStr srcT <> " -> " <> goTypeToStr destT))');
code = code.replace(/_trace = unsafePerformEffect \(traceM \("coerceGoExpr FALLBACK: " <> goTypeToStr from <> " -> " <> goTypeToStr to\)\)/, '_trace = unsafePerformEffect (Console.log ("coerceGoExpr FALLBACK: " <> goTypeToStr from <> " -> " <> goTypeToStr to))');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
