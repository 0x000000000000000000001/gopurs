const fs = require('fs');
const path = 'src/Gopurs/CodeGen.purs';
let content = fs.readFileSync(path, 'utf8');

const target1 = `      , rawDecls: helpers.rawDecls <> (unsafePerformEffect do
          pairsMap <- Ref.read globalReboxPairs
          let reboxPairs = fromMaybe Set.empty (Map.lookup modNameStr pairsMap)`;

const replacement1 = `      , rawDecls: helpers.rawDecls <> (unsafePerformEffect do
          let
            loop generatedFuncs = do
              pairsMap <- Ref.read globalReboxPairs
              let reboxPairs = fromMaybe Set.empty (Map.lookup modNameStr pairsMap)`;

const target2 = `                    Just (Tuple _ info) ->
                      let
                        env1 = Map.fromFoldable (Array.zip info.vars a1)
                        env2 = Map.fromFoldable (Array.zip info.vars a2)
                        assignments = String.joinWith "\\n" (Array.mapWithIndex (\\i fieldExprType -> 
                            let
                              genericTy = structFieldGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors info.vars modNameStr fieldExprType
                              t1 = instantiateGenericGoType env1 genericTy
                              t2 = instantiateGenericGoType env2 genericTy
                            in
                              "\\t\\tout.V" <> show i <> " = " <> printGoExpr (coerceGoExpr modNameStr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)
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
                              "\\t\\tout.V" <> show i <> " = " <> printGoExpr (coerceGoExpr modNameStr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)
                          ) info.fields)
                        funcBody = "func " <> funcName <> "(in *" <> s1 <> ") *" <> s2 <> " {\\n\\tif in == nil { return nil }\\n\\tout := &" <> s2 <> "{}\\n" <> assignments <> "\\n\\treturn out\\n}"
                      in Just (Tuple funcName funcBody)
                    Nothing -> Nothing
                _, _ -> Nothing
            ) (Array.fromFoldable reboxPairs)
            
          pure $ Array.fromFoldable (Map.values generatedFuncs)
        )`;

const replacement2 = `                    Just (Tuple _ info) ->
                      let
                        env1 = Map.fromFoldable (Array.zip info.vars a1)
                        env2 = Map.fromFoldable (Array.zip info.vars a2)
                        assignments = String.joinWith "\\n" (Array.mapWithIndex (\\i fieldExprType -> 
                            let
                              genericTy = structFieldGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors info.vars modNameStr fieldExprType
                              t1 = instantiateGenericGoType env1 genericTy
                              t2 = instantiateGenericGoType env2 genericTy
                            in
                              "\\t\\tout.V" <> show i <> " = " <> printGoExpr (coerceGoExpr modNameStr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)
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
                              "\\t\\tout.V" <> show i <> " = " <> printGoExpr (coerceGoExpr modNameStr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)
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

content = content.replace(target1, replacement1);
content = content.replace(target2, replacement2);
fs.writeFileSync(path, content, 'utf8');
console.log("Patched successfully");
