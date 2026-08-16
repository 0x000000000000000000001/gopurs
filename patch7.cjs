const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(`        Let ident lvl val body ->`, `        Let ident lvl val body ->\n          let _dump = if ident == Ident "__local_var_2_0" then Debug.trace ("!!! DUMPING LOCAL VAR !!!\\n" <> printTcoExprShape val) \\_ -> unit else unit\n          in`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
