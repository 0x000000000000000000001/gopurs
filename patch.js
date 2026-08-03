const fs = require('fs');
let code = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'utf-8');
code = code.replace(
  `    _ -> "gopurs_runtime.Any(" <> valName <> ")"`,
  `    _ -> if anyT == "gopurs_runtime.Value" then valName else "func() gopurs_runtime.Value { if v, ok := any(" <> valName <> ").(gopurs_runtime.Value); ok { return v }; return gopurs_runtime.Any(" <> valName <> ") }()"`
);
code = code.replace(
  `wrapReturn _ _ _ valName = "gopurs_runtime.Box(" <> valName <> ")"`,
  `wrapReturn _ _ _ valName = "gopurs_runtime.Box(" <> valName <> ")"` // keep the same
);
fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', code);
