const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/Printer.purs', 'utf8');

const target = `  GoMutate name expr ->
    name <> " = " <> printGoExpr expr`;

const replacement = `  GoMutate name expr ->
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
    "}\\n}()"`;

if (code.includes(target)) {
    code = code.replace(target, replacement);
} else {
    console.log("Could not find GoMutate block. Trying regex.");
    code = code.replace(/  GoMutate name expr ->\s*name <> " = " <> printGoExpr expr/, replacement);
}

fs.writeFileSync('src/Gopurs/Printer.purs', code);
