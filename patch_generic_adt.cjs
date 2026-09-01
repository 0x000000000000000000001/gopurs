const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const targetStr = `        Just info ->
          let
            pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (String.joinWith "." (Array.slice 0 (Array.length path - 1) path))
            monoStructName = "Constructor_" <> pkgNameStr <> "_" <> sanitizeName info.ctorName
          in
          if Set.member monoStructName elided then TypeValue
          else if info.arity == 0 then TypeStructPointer monoStructName monoStructName monoStructName []
          else
            let
              finalArgs =
                if Array.length args == info.arity then
                  map (exprTypeToGenericGoType ptrPaths enumAdts elided typeVars modNameStr) args
                else if Array.length typeVars == info.arity then
                  map TypeGenericParam typeVars
                else
                  Array.replicate info.arity TypeValue
            in
              TypeStructPointer monoStructName monoStructName monoStructName finalArgs
        Nothing -> TypeValue`;

const replacementStr = `        Just info ->
          let
            pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (String.joinWith "." (Array.slice 0 (Array.length path - 1) path))
            monoStructName = "Constructor_" <> pkgNameStr <> "_" <> sanitizeName info.ctorName
            baseStructName = "Data_" <> pkgNameStr <> "_" <> sanitizeName info.ctorName
          in
          if Set.member monoStructName elided then TypeValue
          else if info.arity == 0 then TypeStructPointer baseStructName fullName monoStructName []
          else
            let
              finalArgs =
                if Array.length args == info.arity then
                  map (exprTypeToGenericGoType ptrPaths enumAdts elided typeVars modNameStr) args
                else if Array.length typeVars == info.arity then
                  map TypeGenericParam typeVars
                else
                  Array.replicate info.arity TypeValue
              typeArgsStr = if Array.length finalArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr finalArgs) <> "]" else ""
            in
              TypeStructPointer baseStructName fullName (monoStructName <> typeArgsStr) finalArgs
        Nothing -> TypeValue`;

if (!code.includes(targetStr)) {
  console.error("Target string not found!");
  process.exit(1);
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.replace(targetStr, replacementStr));
console.log("Patched CodeGen.purs successfully.");
