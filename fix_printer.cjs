const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/Printer.purs', 'utf8');

// Replace the entire GoRecordUpdateStatic block with GoRecordUpdateNative
const blockRegex = /  GoRecordUpdateStatic orig size updates ->[\s\S]*?\(val2, ok2 := orig \(\*\)\(gopurs_runtime\.RecordData\)\)\\n" <>\n      "          _ = ok2\\n"\n      \)/;

// Wait, the regex might be tricky. Let's just do a string replacement.
code = code.replace(/  GoRecordUpdateStatic orig size updates ->[\s\S]*?\)\n/, `  GoRecordUpdateNative expr typeName fields ->
    "func() *" <> typeName <> " {\\n" <>
    "  var newRec = *" <> printGoExpr expr <> "\\n" <>
    String.joinWith "\\n" (map (\\(Tuple prop val) -> "  newRec." <> prop <> " = " <> printGoExpr val) fields) <> "\\n" <>
    "  return &newRec\\n" <>
    "}()"\n`);

fs.writeFileSync('src/Gopurs/Printer.purs', code);
