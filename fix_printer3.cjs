const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/Printer.purs', 'utf8');

code = code.replace(/  GoRecordAccessNative expr prop ->\n    printGoExpr expr <> "\." <> prop <> "]"\n    else\n      "\(\(\*gopurs_runtime\.RecordData" <> show size <> "\)\(" <> printGoExpr obj <> "\.UnsafePtr\)\)\.V" <> show idx/g, `  GoRecordAccessNative expr prop ->\n    printGoExpr expr <> "." <> prop`);

fs.writeFileSync('src/Gopurs/Printer.purs', code);
