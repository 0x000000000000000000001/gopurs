sed -i '' 's/goFieldTypes = map (exprTypeToGoType modNameStr) monomorphizedFields/goFieldTypes = map (structFieldGoType modNameStr) monomorphizedFields/g' src/Gopurs/CodeGen.purs
sed -i '' 's/goFieldTypes = map (exprTypeToGoType modNameStr) monomorphizedFieldTypes/goFieldTypes = map (structFieldGoType modNameStr) monomorphizedFieldTypes/g' src/Gopurs/CodeGen.purs
