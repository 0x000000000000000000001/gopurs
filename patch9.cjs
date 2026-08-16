const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(`        Let ident lvl val body ->\\n          let\\n            idStr = localId ident lvl`, `        Let ident lvl val body ->\\n          let\\n            idStr = localId ident lvl\\n            _trace = if idStr == "__local_var_2_0" then Debug.trace ("=== LOCAL VAR 2_0 ===\\n" <> printTcoExprShape val) (\\_ -> unit) else unit`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
