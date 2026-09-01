const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const targetStr = `instantiateGenericGoType env (TypeStructPointer base key full typeArgs) = TypeStructPointer base key full (map (instantiateGenericGoType env) typeArgs)`;

const replacementStr = `instantiateGenericGoType env (TypeStructPointer base key full typeArgs) =
  let
    newTypeArgs = map (instantiateGenericGoType env) typeArgs
    typeArgsStr = if Array.length newTypeArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr newTypeArgs) <> "]" else ""
    monoStructName = case String.indexOf (Pattern "[") full of
      Just i -> String.take i full
      Nothing -> full
  in
    TypeStructPointer base key (monoStructName <> typeArgsStr) newTypeArgs`;

if (!code.includes(targetStr)) {
  console.error("Target string not found!");
  process.exit(1);
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.replace(targetStr, replacementStr));
console.log("Patched CodeGen.purs successfully.");
