const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(`        Let ident lvl val body ->\\n          let _dump =`, `        Let ident lvl val body ->\\n          let\\n            _dump =`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
