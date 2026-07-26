const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/Printer.purs', 'utf8');

const updateTarget = `  GoRecordUpdateStatic orig size updates ->
    let
      structName = if size >= 6 then "gopurs_runtime.RecordData" else "gopurs_runtime.RecordData" <> show size
      typeVal = if size >= 6 then "gopurs_runtime.TypeRecordData" else "gopurs_runtime.TypeRecord" <> show size
    in
      if size >= 6 then
        let
          assignments = String.joinWith "\\n" (map (\\(Tuple idx val) -> "newVals[" <> show idx <> "] = " <> printGoExpr val) updates)
        in
          "func() gopurs_runtime.Value {\\nr := (*" <> structName <> ")(" <> printGoExpr orig <> ".UnsafePtr)\\nnewVals := make([]gopurs_runtime.Value, len(r.Vals))\\ncopy(newVals, r.Vals)\\n" <> assignments <> "\\nnewR := gopurs_runtime.RecordData{Keys: r.Keys, Vals: newVals}\\nreturn gopurs_runtime.Value{Type: " <> typeVal <> ", UnsafePtr: unsafe.Pointer(&newR)}\\n}()"
      else
        let
          assignments = String.joinWith "\\n" (map (\\(Tuple idx val) -> "clone.V" <> show idx <> " = " <> printGoExpr val) updates)
        in
          "func() gopurs_runtime.Value {\\nclone := *((*" <> structName <> ")(" <> printGoExpr orig <> ".UnsafePtr))\\n" <> assignments <> "\\nreturn gopurs_runtime.Value{Type: " <> typeVal <> ", UnsafePtr: unsafe.Pointer(&clone)}\\n}()"`;

const updateReplacement = `  GoRecordUpdateNative expr typeName fields ->
    "func() *" <> typeName <> " {\\n" <>
    "  var newRec = *" <> printGoExpr expr <> "\\n" <>
    String.joinWith "\\n" (map (\\(Tuple prop val) -> "  newRec." <> prop <> " = " <> printGoExpr val) fields) <> "\\n" <>
    "  return &newRec\\n" <>
    "}()"`;

const accessTarget = `  GoRecordAccessStatic obj size idx ->
    if size >= 6 then
      "((*gopurs_runtime.RecordData)(" <> printGoExpr obj <> ".UnsafePtr)).Vals[" <> show idx <> "]"
    else
      "((*gopurs_runtime.RecordData" <> show size <> ")(" <> printGoExpr obj <> ".UnsafePtr)).V" <> show idx`;

const accessReplacement = `  GoRecordAccessNative expr prop ->
    printGoExpr expr <> "." <> prop`;

if (code.includes(updateTarget)) {
    code = code.replace(updateTarget, updateReplacement);
} else {
    console.log("Could not find GoRecordUpdateStatic block");
}

if (code.includes(accessTarget)) {
    code = code.replace(accessTarget, accessReplacement);
} else {
    console.log("Could not find GoRecordAccessStatic block");
}

fs.writeFileSync('src/Gopurs/Printer.purs', code);
