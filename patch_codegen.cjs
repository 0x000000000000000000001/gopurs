const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// Fix unboxGoExpr for TypeString
code = code.replace(/TypeString -> GoSelector expr "StrVal"/, 'TypeString -> GoCall (GoSelector expr "StrVal") []');

// Fix isTagHelper
code = code.replace(/isTagHelper = "func Is_" \+ structName \+ "\\(v gopurs_runtime.Value\\) bool \\{\\n\\t_, ok := v.PtrVal.\\(\\*" \+ structName \+ "\\)\\n\\treturn ok\\n\\}"/, 'isTagHelper = "func Is_" <> structName <> "(v gopurs_runtime.Value) bool {\\n\\treturn v.Type == 9 && v.IntVal == " <> hashStr <> "\\n}"');
// Wait, my replace regex might be wrong above, let's use a simpler one.
code = code.replace(/_, ok := v.PtrVal\.\(\\\*"\s*<>\s*structName\s*<>\s*"\)/g, 'return v.Type == 9 /* Replaced */');
// Actually, let's just do:
code = code.replace(/isTagHelper = "func Is_"[^}]*return ok\\n\}"/g, 'isTagHelper = "func Is_" <> structName <> "(v gopurs_runtime.Value) bool {\\n\\treturn v.Type == 9 && v.IntVal == " <> hashStr <> "\\n}"');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
