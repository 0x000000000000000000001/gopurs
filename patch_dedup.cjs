const fs = require('fs');
const file = 'src/Gopurs/CodeGen.purs';
let code = fs.readFileSync(file, 'utf8');

const targetStr = `              funcs = Array.fromFoldable $ Set.map (\\(Tuple srcT destT) -> 
                case Tuple srcT destT of
                  Tuple (TypeStructPointer baseName fullName srcFullPath srcArgs) (TypeStructPointer _ _ destFullPath destArgs) ->
                    let`;

const replaceStr = `              funcsWithNames = Array.fromFoldable $ Set.map (\\(Tuple srcT destT) -> 
                case Tuple srcT destT of
                  Tuple (TypeStructPointer baseName fullName srcFullPath srcArgs) (TypeStructPointer _ _ destFullPath destArgs) ->
                    let`;

code = code.replace(targetStr, replaceStr);

const targetStr2 = `                      funcName = "Rebox_" <> modNameStr <> "_" <> hashString srcFullPath <> "_" <> hashString destFullPath
                    in
                        "func " <> funcName <> "(src *" <> srcFullPath <> ") *" <> destFullPath <> " {\\n" <>
                        "\\tif src == nil { return nil }\\n" <>
                        "\\tdest := new(" <> destFullPath <> ")\\n" <>
                        String.joinWith "\\n" assignments <> "\\n" <>
                        "\\treturn dest\\n" <>
                        "}\\n"
              ) unprocessed`;

const replaceStr2 = `                      funcName = "Rebox_" <> modNameStr <> "_" <> hashString srcFullPath <> "_" <> hashString destFullPath
                    in
                        Tuple funcName $ "func " <> funcName <> "(src *" <> srcFullPath <> ") *" <> destFullPath <> " {\\n" <>
                        "\\tif src == nil { return nil }\\n" <>
                        "\\tdest := new(" <> destFullPath <> ")\\n" <>
                        String.joinWith "\\n" assignments <> "\\n" <>
                        "\\treturn dest\\n" <>
                        "}\\n"
              ) unprocessed
              
              funcs = Array.fromFoldable $ Map.values $ Map.fromFoldable funcsWithNames`;

code = code.replace(targetStr2, replaceStr2);

fs.writeFileSync(file, code);
