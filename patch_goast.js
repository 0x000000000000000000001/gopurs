const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/GoAst.purs', 'utf-8');

// 1. goTypeToStr (TypeGenericParam name) = "gopurs_runtime.Value"
code = code.replace(
  /goTypeToStr \(TypeGenericParam name\) = .*/,
  'goTypeToStr (TypeGenericParam _) = "gopurs_runtime.Value"'
);

// 2. erasedGoTypeToStr (TypeGenericParam _) = "gopurs_runtime.Value" (Already there, but just in case)
code = code.replace(
  /erasedGoTypeToStr \(TypeGenericParam _\) = .*/,
  'erasedGoTypeToStr (TypeGenericParam _) = "gopurs_runtime.Value"'
);

// 3. goTypeToStr TypeStructPointer
code = code.replace(
  /goTypeToStr \(TypeStructPointer _ _ monoStructName args\) = "\*" <> monoStructName <> \(if Array\.length args > 0 .*/,
  'goTypeToStr (TypeStructPointer _ _ monoStructName _) = "*" <> monoStructName'
);

// 4. erasedGoTypeToStr TypeStructPointer
code = code.replace(
  /erasedGoTypeToStr \(TypeStructPointer _ _ monoStructName args\) = "\*" <> monoStructName <> \(if Array\.length args > 0 .*/,
  'erasedGoTypeToStr (TypeStructPointer _ _ monoStructName _) = "*" <> monoStructName'
);

// 5. GoConstructor - let's check its definition
// data GoExpr = GoConstructor String String (Array GoType) (Array GoExpr)
// printGoExpr (GoConstructor hash monoStructName typeArgs fields) = ...
code = code.replace(
  /printGoExpr \(GoConstructor _ monoStructName typeArgs fields\) = .*/,
  'printGoExpr (GoConstructor _ monoStructName _ fields) = "(\\&" <> monoStructName <> "{" <> String.joinWith ", " (map printGoExpr fields) <> "})"'
);
// Wait, GoConstructor has a fixed format! Let me check how it is printed first!
fs.writeFileSync('patch_goast_code.txt', code);
