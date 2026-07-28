const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  'Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Boolean -> Int -> TcoExpr',
  'Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Set String -> Boolean -> Boolean -> Int -> TcoExpr'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
