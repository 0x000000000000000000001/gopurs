const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(`        Let ident lvl val body ->\\n          let\\n            idStr = localId ident lvl\\n            expectedType =`, `        Let ident lvl val body ->\\n          let\\n            idStr = localId ident lvl\\n            _dump = if idStr == "__local_var_2_0" || idStr == "__local_var_2_1" then Debug.trace ("=== DUMP " <> idStr <> " ===\\n" <> printTcoExprShape val) (\\_ -> unit) else unit\\n            expectedType =`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
