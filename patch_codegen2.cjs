const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(/isTagHelper = "func Is_" <> structName <> "\([^)]+\) bool \{\\n\\treturn v\.Type == 9 && v\.IntVal == " <> hashStr <> "\\n\}"/g, 
  'hashStr = hashString structName\n              isTagHelper = "func Is_" <> structName <> "(v gopurs_runtime.Value) bool {\\n\\treturn v.Type == 9 && v.IntVal == " <> hashStr <> "\\n}"');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
