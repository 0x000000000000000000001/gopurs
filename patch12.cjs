const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(`            shouldBypassCoerce = isAppNode && res.exprType == TypeValue && expectedGoType /= TypeValue`, `            shouldBypassCoerce = isAppNode && res.exprType == TypeValue && expectedGoType /= TypeValue\\n            _traceBypass = if res.exprType == TypeValue && expectedGoType /= TypeValue then Debug.trace ("!!! BYPASS CHECK !!! shape=" <> printTcoExprShape inner <> " type=" <> goTypeToStr expectedGoType) (\\_ -> unit) else unit`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
