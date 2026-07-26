const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/Printer.purs', 'utf8');

code = code.replace(/\\\\n/g, '\\n');
code = code.replace(/\\\\/g, '\\');

fs.writeFileSync('src/Gopurs/Printer.purs', code);
