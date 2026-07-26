sed -i '' 's/goFieldTypes = map (exprTypeToGoType modNameStr) monomorphizedFields/mangledFields = map (mangleType modNameStr) monomorphizedFields/g' src/Gopurs/CodeGen.purs

sed -i '' 's/goFieldTypes = map (exprTypeToGoType modNameStr) monomorphizedFieldTypes/mangledFields = map (mangleType modNameStr) monomorphizedFieldTypes/g' src/Gopurs/CodeGen.purs

sed -i '' 's/if Array.length goFieldTypes == 0 then "" else "_" <> String.replaceAll (Pattern "*") (Replacement "ptr") (String.replaceAll (Pattern "\[\]") (Replacement "arr") (String.replaceAll (Pattern " ") (Replacement "_") (String.joinWith "_" (map goTypeToStr goFieldTypes))))/if Array.length mangledFields == 0 then "" else "_" <> String.joinWith "_" mangledFields/g' src/Gopurs/CodeGen.purs
