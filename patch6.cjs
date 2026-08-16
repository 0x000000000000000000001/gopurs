const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// We will add a Debug.trace in unboxGoExpr when desiredType is TypeNativeArray
code = code.replace(`    (TypeNativeArray inner) -> case currentType of`, `    (TypeNativeArray inner) -> Debug.trace ("!!! UNBOXING TO ARRAY !!! currentType=" <> goTypeToStr currentType <> "\\nexpr=" <> printGoExpr expr) \\_ -> case currentType of`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
