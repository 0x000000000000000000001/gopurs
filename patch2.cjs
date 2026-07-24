const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/Runtime.purs', 'utf8');

code = code.replace(/StrVal: tag, /g, '');

fs.writeFileSync('src/Gopurs/Runtime.purs', code);
