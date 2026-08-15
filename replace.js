const fs = require('fs');
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// Update signature
content = content.replace(
  'exprTypeToGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> String -> ExprType -> GoType',
  'exprTypeToGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> String -> ExprType -> GoType'
);

// Update base definition
content = content.replace(
  /exprTypeToGoType _ _ _ Int = TypeInt64\nexprTypeToGoType _ _ _ Number = TypeFloat64\nexprTypeToGoType _ _ _ String = TypeString\nexprTypeToGoType _ _ _ Char = TypeString\nexprTypeToGoType _ _ _ Boolean = TypeBool\nexprTypeToGoType ptrPaths enumAdts modNameStr \(Array ty\) = TypeNativeArray \(exprTypeToGoType ptrPaths enumAdts modNameStr ty\)\nexprTypeToGoType ptrPaths enumAdts modNameStr \(Record \(Row fields _\)\) = TypeRecord \(map \(\\\(Tuple k v\) -> Tuple k \(exprTypeToGoType ptrPaths enumAdts modNameStr v\)\) \(Array\.sortBy \(comparing \\\(Tuple k _\) -> k\) fields\)\)\nexprTypeToGoType ptrPaths enumAdts modNameStr \(Record _\) = TypeValue\nexprTypeToGoType ptrPaths enumAdts modNameStr \(ADT fullName path args\) =/g,
  `exprTypeToGoType _ _ _ _ Int = TypeInt64
exprTypeToGoType _ _ _ _ Number = TypeFloat64
exprTypeToGoType _ _ _ _ String = TypeString
exprTypeToGoType _ _ _ _ Char = TypeString
exprTypeToGoType _ _ _ _ Boolean = TypeBool
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Array ty) = TypeNativeArray (exprTypeToGoType ptrPaths enumAdts elided modNameStr ty)
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Record (Row fields _)) = TypeRecord (map (\\(Tuple k v) -> Tuple k (exprTypeToGoType ptrPaths enumAdts elided modNameStr v)) (Array.sortBy (comparing \\(Tuple k _) -> k) fields))
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Record _) = TypeValue
exprTypeToGoType ptrPaths enumAdts elided modNameStr (ADT fullName path args) =
  if Set.member (getStructName "" Nothing (snd path)) elided then TypeValue else`
);

// We need to replace all calls:
// 1. `exprTypeToGoType ptrPaths enumAdts modNameStr` -> `exprTypeToGoType ptrPaths enumAdts elidedCtors modNameStr`
// 2. `exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts` -> 
//    `exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors`

content = content.replace(
  /exprTypeToGoType ptrPaths enumAdts modNameStr/g,
  'exprTypeToGoType ptrPaths enumAdts elidedCtors modNameStr'
);

content = content.replace(
  /exprTypeToGoType \(unsafePerformEffect \(Ref\.read helpersRef\)\)\.pointerAdtPaths \(unsafePerformEffect \(Ref\.read helpersRef\)\)\.enumAdts/g,
  'exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors'
);

// Fix getRefType signature? Wait, where is ptrPaths passed to exprTypeToGoType?
// In `exprTypeToGenericGoType`:
content = content.replace(
  /exprTypeToGenericGoType :: Map\.Map String \{ ctorName :: String, arity :: Int \} -> Set\.Set String -> Array String -> String -> ExprType -> GoType/,
  'exprTypeToGenericGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> Array String -> String -> ExprType -> GoType'
);

content = content.replace(
  /exprTypeToGenericGoType ptrPaths enumAdts vars modNameStr/g,
  'exprTypeToGenericGoType ptrPaths enumAdts elidedCtors vars modNameStr'
);

content = content.replace(
  /exprTypeToGenericGoType \(unsafePerformEffect \(Ref\.read helpersRef\)\)\.pointerAdtPaths \(unsafePerformEffect \(Ref\.read helpersRef\)\)\.enumAdts/g,
  'exprTypeToGenericGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
