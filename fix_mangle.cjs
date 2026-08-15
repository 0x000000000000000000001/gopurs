const fs = require('fs');
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

content = content.replace(
  /mangleType :: Map\.Map String \{ ctorName :: String, arity :: Int \} -> Set\.Set String -> String -> ExprType -> String/g,
  'mangleType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> String -> ExprType -> String'
);

content = content.replace(
  /mangleType ptrPaths enumAdts modNameStr t =/g,
  'mangleType ptrPaths enumAdts elidedCtors modNameStr t ='
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
