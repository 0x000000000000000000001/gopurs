const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(`        Typed t inner ->\\n          let\\n            expectedGoType = exprTypeToGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors modNameStr t\\n            res = translateExprImpl_ helpers modNameStr nextId isTail newBound inner\\n          in`, `        Typed t inner ->\\n          let\\n            expectedGoType = exprTypeToGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors modNameStr t\\n            res = translateExprImpl_ helpers modNameStr nextId isTail newBound inner\\n            _trace = if goTypeToStr expectedGoType == "[]gopurs_runtime.Value" && goTypeToStr res.exprType == "gopurs_runtime.Value" then Debug.trace ("\\n==== BUGGY TYPED ===\\nexpectedType=" <> goTypeToStr expectedGoType <> "\\nactualType=" <> goTypeToStr res.exprType <> "\\nAST=" <> printTcoExprShape inner <> "\\nTAST_TYPE=" <> printExprType t <> "\\n==================\\n") (\\_ -> unit) else unit\\n          in`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
