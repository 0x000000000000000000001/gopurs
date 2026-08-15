const fs = require('fs');
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

content = content.replace(
  /exprTypeToGenericGoType ptrPaths enumAdts typeVars modNameStr/g,
  'exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modNameStr'
);

content = content.replace(
  /exprTypeToGenericGoType _ _ _ _ \(Record _\)/g,
  'exprTypeToGenericGoType _ _ _ _ _ (Record _)'
);

content = content.replace(
  /exprTypeToGenericGoType _ _ typeVars _ \(TypeVar/g,
  'exprTypeToGenericGoType _ _ _ typeVars _ (TypeVar'
);

content = content.replace(
  /exprTypeToGenericGoType ptrPaths enumAdts _ modNameStr ty = exprTypeToGoType ptrPaths enumAdts elidedCtors modNameStr ty/g,
  'exprTypeToGenericGoType ptrPaths enumAdts elidedCtors _ modNameStr ty = exprTypeToGoType ptrPaths enumAdts elidedCtors modNameStr ty'
);

content = content.replace(
  /structFieldGoType ptrPaths enumAdts typeVars modStr ty =/g,
  'structFieldGoType ptrPaths enumAdts elidedCtors typeVars modStr ty ='
);

content = content.replace(
  /structFieldGoType :: Map\.Map String \{ ctorName :: String, arity :: Int \} -> Set\.Set String -> Array String -> String -> ExprType -> GoType/g,
  'structFieldGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> Array String -> String -> ExprType -> GoType'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
