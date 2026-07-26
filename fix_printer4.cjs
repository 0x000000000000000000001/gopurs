const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/Printer.purs', 'utf8');

code = code.replace(/  GoMutate name expr ->\n      name <> " = " <> printGoExpr expr/, `  GoMutate name expr ->
      name <> " = " <> printGoExpr expr
  GoStructInit structName fields ->
    "&" <> structName <> "{" <> String.joinWith ", " (map (\\(Tuple prop val) -> prop <> ": " <> printGoExpr val) fields) <> "}"
  GoStructUpdate orig structName updates ->
    "func() *" <> structName <> " {\\n" <>
    "  var newRec = *" <> printGoExpr orig <> "\\n" <>
    String.joinWith "\\n" (map (\\(Tuple prop val) -> "  newRec." <> prop <> " = " <> printGoExpr val) updates) <> "\\n" <>
    "  return &newRec\\n" <>
    "}()"
  GoTypeSwitch expr cases ->
    "func() gopurs_runtime.Value {\\nswitch v := (" <> printGoExpr expr <> ").(type) {\\n" <>
    String.joinWith "\\n" (map (\\(Tuple typeName val) -> "  case " <> typeName <> ":\\n    return " <> printGoExpr val) cases) <> "\\n" <>
    "  default:\\n    panic(\\"unhandled type switch\\")\\n" <>
    "}\\n}()"`);

fs.writeFileSync('src/Gopurs/Printer.purs', code);
