const fs = require('fs');
const path = 'src/Gopurs/CodeGen.purs';
let content = fs.readFileSync(path, 'utf8');

const targetRegex = /,\s*rawDecls:\s*helpers\.rawDecls\s*<>\s*\(unsafePerformEffect do[\s\S]*?\)\s*\(Array\.fromFoldable reboxPairs\)\n\s*\)/g;

const replacement = `, rawDecls: helpers.rawDecls <> (unsafePerformEffect do
          let
            loop generatedFuncs = do
              pairsMap <- Ref.read globalReboxPairs
              let reboxPairs = fromMaybe Set.empty (Map.lookup modNameStr pairsMap)
              
              let 
                newFuncs = Map.fromFoldable $ Array.mapMaybe (\\(Tuple srcT destT) -> 
                  case srcT, destT of
                    TypeStructPointer b1 f1 s1 a1, TypeStructPointer b2 f2 s2 a2 | b1 == b2 ->
                      let
                        funcName = "Rebox_" <> modNameStr <> "_" <> hashString s1 <> "_" <> hashString s2
                      in if Map.member funcName generatedFuncs then Nothing else
                      let
                        matchB1 k =
                          let
                            parts = String.split (Pattern "_") k
                            pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
                          in if Array.length parts >= 2 then
                              let
                                ctorName = fromMaybe "" (Array.last parts)
                                pkgName = String.joinWith "_" (Array.slice 0 (Array.length parts - 1) parts)
                                expectedB1 = "Constructor_" <> pkgName <> "_" <> sanitizeName ctorName
                                expectedB2 = "Data_" <> pkgName <> "_" <> sanitizeName ctorName
                              in expectedB1 == b1 || expectedB2 == b1
                            else false
                            
                        mbCtor = Array.find (\\(Tuple k _) -> matchB1 k) (Map.toUnfoldable helpers.ctorTypes :: Array (Tuple String { vars :: Array String, fields :: Array ExprType }))
                        mbClass = Array.find (\\(Tuple k _) -> matchB1 k) (Map.toUnfoldable helpers.classDeclsFields :: Array (Tuple String { vars :: Array String, fields :: Array { name :: String, "type" :: ExprType } }))
                      in case mbCtor of
                        Just (Tuple _ info) ->
                          let
                            env1 = Map.fromFoldable (Array.zip info.vars a1)
                            env2 = Map.fromFoldable (Array.zip info.vars a2)
                            assignments = String.joinWith "\\n" (Array.mapWithIndex (\\i fieldExprType -> 
                                let
                                  genericTy = structFieldGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors info.vars modNameStr fieldExprType
                                  t1 = instantiateGenericGoType env1 genericTy
                                  t2 = instantiateGenericGoType env2 genericTy
                                in
                                  "\\t\\tout.V" <> show i <> " = " <> printGoExpr (coerceGoExpr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)
                              ) info.fields)
                            funcBody = "func " <> funcName <> "(in *" <> s1 <> ") *" <> s2 <> " {\\n\\tif in == nil { return nil }\\n\\tout := &" <> s2 <> "{}\\n" <> assignments <> "\\n\\treturn out\\n}"
                          in Just (Tuple funcName funcBody)
                        Nothing ->
                          let
                            ctorVarsAndFields = case mbClass of
                              Just (Tuple _ classInfo) ->
                                Just { vars: classInfo.vars, fields: map (\\f -> f."type") classInfo.fields }
                              Nothing ->
                                let
                                  _trace = unsafePerformEffect (Console.log ("ERROR: Rebox missing! b1=" <> b1 <> " keysCtor: " <> String.joinWith ", " (map fst (Map.toUnfoldable helpers.ctorTypes :: Array (Tuple String _)))))
                                in Nothing
                          in case ctorVarsAndFields of
                            Just info ->
                              let
                                env1 = Map.fromFoldable (Array.zip info.vars a1)
                                env2 = Map.fromFoldable (Array.zip info.vars a2)
                                assignments = String.joinWith "\\n" (Array.mapWithIndex (\\i fieldExprType -> 
                                    let
                                      genericTy = structFieldGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors info.vars modNameStr fieldExprType
                                      t1 = instantiateGenericGoType env1 genericTy
                                      t2 = instantiateGenericGoType env2 genericTy
                                    in
                                      "\\t\\tout.V" <> show i <> " = " <> printGoExpr (coerceGoExpr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)
                                  ) info.fields)
                                funcBody = "func " <> funcName <> "(in *" <> s1 <> ") *" <> s2 <> " {\\n\\tif in == nil { return nil }\\n\\tout := &" <> s2 <> "{}\\n" <> assignments <> "\\n\\treturn out\\n}"
                              in Just (Tuple funcName funcBody)
                            Nothing -> Nothing
                    _, _ -> Nothing
                ) (Array.fromFoldable reboxPairs)
              
              let nextGeneratedFuncs = Map.union generatedFuncs newFuncs
              if Map.isEmpty newFuncs then
                pure $ Array.fromFoldable (Map.values nextGeneratedFuncs)
              else
                loop nextGeneratedFuncs

          loop Map.empty
        )`;

content = content.replace(targetRegex, replacement);
fs.writeFileSync(path, content, 'utf8');
console.log("Patched successfully!");
