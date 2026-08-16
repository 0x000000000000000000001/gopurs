const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(`            _traceBypass = if res.exprType == TypeValue && expectedGoType /= TypeValue then Debug.trace ("!!! BYPASS CHECK !!! shape=" <> printTcoExprShape inner <> " type=" <> goTypeToStr expectedGoType) (\\_ -> unit) else unit`, `            _traceBypass = if goTypeToStr expectedGoType == "[]gopurs_runtime.Value" && goTypeToStr res.exprType == "gopurs_runtime.Value" then Debug.trace ("!!! BYPASS ALL !!! shape=" <> printTcoExprShape inner <> " type=" <> goTypeToStr expectedGoType) (\\_ -> unit) else unit`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
