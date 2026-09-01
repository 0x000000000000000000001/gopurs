const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

// Replace the buggy func signature and instantiation
code = code.replace(
  /"func " <> funcName <> "\(in \*" <> b1 <> "\[" <> String\.joinWith ", " \(map goTypeToStr a1\) <> "\]\) \*" <> b2 <> "\[" <> String\.joinWith ", " \(map goTypeToStr a2\) <> "\] \{\\n\\tif in == nil \{ return nil \}\\n\\tout := \&" <> b2 <> "\[" <> String\.joinWith ", " \(map goTypeToStr a2\) <> "\]\{\}\\n" <> assignments <> "\\n\\treturn out\\n\}"/g,
  '"func " <> funcName <> "(in *" <> s1 <> ") *" <> s2 <> " {\\n\\tif in == nil { return nil }\\n\\tout := &" <> s2 <> "{}\\n" <> assignments <> "\\n\\treturn out\\n}"'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
