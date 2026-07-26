const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/Printer.purs', 'utf8');

code = code.replace(/  GoRecordAccessStatic obj size idx ->[\s\S]*?show idx/, `  GoRecordAccessNative expr prop ->
    printGoExpr expr <> "." <> prop`);

fs.writeFileSync('src/Gopurs/Printer.purs', code);
