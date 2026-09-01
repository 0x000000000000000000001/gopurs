const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const targetStr = `                    loopCtxs = map
                      ( \\fn ->
                          let
                            oldName = localId (Just (Ident fn.ident)) lvl
                            newName = (fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)).name
                            pTypes = paramTypes fn.body
                            paramsWithTypes = map
                              ( \\idStr ->
                                  let
                                    t = fromMaybe Any (Map.lookup idStr pTypes)
                                  in
                                    Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr t)
                              )
                              fn.args
                          in
                            { ident: newName, params: fn.args, loopParams: map (\\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes }
                      )
                      fns

                    combinedLoopCtx = loopCtxs <> loopCtx

                    localModuleArities = foldl (\\acc ctx -> Map.insert ctx.ident { fullName: "Call_local_" <> modNameStr <> "_" <> ctx.ident, fArgs: ctx.goTypes, fRet: TypeValue, arity: Array.length ctx.params } acc) moduleArities loopCtxs
                    declStmts = Array.concatMap (\\ctx -> [ GoRaw ("var Call_local_" <> modNameStr <> "_" <> ctx.ident <> " func(" <> String.joinWith ", " (map goTypeToStr ctx.goTypes) <> ") gopurs_runtime.Value"), GoRaw ("_ = Call_local_" <> modNameStr <> "_" <> ctx.ident), GoRaw ("var " <> ctx.ident <> " gopurs_runtime.Value"), GoRaw ("_ = " <> ctx.ident) ]) loopCtxs

                    Tuple fnWrapperStmts nextId' = foldl
                      ( \\(Tuple accStmts currNextId) fn ->
                          let
                            oldName = localId (Just (Ident fn.ident)) lvl
                            newName = (fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)).name
                            pTypes = paramTypes fn.body
                            paramsWithTypes = map
                              ( \\idStr ->
                                  let
                                    t = fromMaybe Any (Map.lookup idStr pTypes)
                                  in
                                    Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr t)
                              )
                              fn.args
                            currentLoopCtx = [ { ident: newName, params: fn.args, loopParams: map (\\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes } ]
                            loopBound = foldl (\\acc (Tuple idStr goT) -> Map.insert idStr { name: idStr, goType: goT } acc) allocRes.newBound paramsWithTypes
                            resBodyMut = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars localModuleArities loopBound (Just newName) currentLoopCtx true false currNextId fn.body

                            loopParams = map (\\(Tuple idStr _) -> idStr <> "_loop") paramsWithTypes
                            initVars = Array.concatMap (\\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes

                            funcBody = GoFor newName (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr resBodyMut.expr resBodyMut.exprType) ])

                            goParamsNative = String.joinWith ", " (map (\\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)
                            nativeAssignment = GoMutate ("Call_local_" <> modNameStr <> "_" <> newName) (GoRaw ("func(" <> goParamsNative <> ") gopurs_runtime.Value {\\n" <> printGoExpr funcBody <> "\\n}"))`;

const replacementStr = `                    loopCtxs = map
                      ( \\fn ->
                          let
                            oldName = localId (Just (Ident fn.ident)) lvl
                            boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                            newName = boundInfo.name
                            goType = boundInfo.goType
                            fRet = case goType of
                              TypeFunc _ r -> r
                              _ -> TypeValue
                            pTypes = paramTypes fn.body
                            paramsWithTypes = map
                              ( \\idStr ->
                                  let
                                    t = fromMaybe Any (Map.lookup idStr pTypes)
                                  in
                                    Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr t)
                              )
                              fn.args
                          in
                            { ident: newName, params: fn.args, loopParams: map (\\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes, fRet }
                      )
                      fns

                    combinedLoopCtx = loopCtxs <> loopCtx

                    localModuleArities = foldl (\\acc ctx -> Map.insert ctx.ident { fullName: "Call_local_" <> modNameStr <> "_" <> ctx.ident, fArgs: ctx.goTypes, fRet: ctx.fRet, arity: Array.length ctx.params } acc) moduleArities loopCtxs
                    declStmts = Array.concatMap (\\ctx -> [ GoRaw ("var Call_local_" <> modNameStr <> "_" <> ctx.ident <> " func(" <> String.joinWith ", " (map goTypeToStr ctx.goTypes) <> ") " <> goTypeToStr ctx.fRet), GoRaw ("_ = Call_local_" <> modNameStr <> "_" <> ctx.ident), GoRaw ("var " <> ctx.ident <> " gopurs_runtime.Value"), GoRaw ("_ = " <> ctx.ident) ]) loopCtxs

                    Tuple fnWrapperStmts nextId' = foldl
                      ( \\(Tuple accStmts currNextId) fn ->
                          let
                            oldName = localId (Just (Ident fn.ident)) lvl
                            boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                            newName = boundInfo.name
                            goType = boundInfo.goType
                            fRet = case goType of
                              TypeFunc _ r -> r
                              _ -> TypeValue
                            pTypes = paramTypes fn.body
                            paramsWithTypes = map
                              ( \\idStr ->
                                  let
                                    t = fromMaybe Any (Map.lookup idStr pTypes)
                                  in
                                    Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr t)
                              )
                              fn.args
                            currentLoopCtx = [ { ident: newName, params: fn.args, loopParams: map (\\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes, fRet } ]
                            loopBound = foldl (\\acc (Tuple idStr goT) -> Map.insert idStr { name: idStr, goType: goT } acc) allocRes.newBound paramsWithTypes
                            resBodyMut = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars localModuleArities loopBound (Just newName) currentLoopCtx true false currNextId fn.body

                            loopParams = map (\\(Tuple idStr _) -> idStr <> "_loop") paramsWithTypes
                            initVars = Array.concatMap (\\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes

                            funcBody = GoFor newName (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (coerceGoExpr modNameStr resBodyMut.expr resBodyMut.exprType fRet) ])

                            goParamsNative = String.joinWith ", " (map (\\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)
                            nativeAssignment = GoMutate ("Call_local_" <> modNameStr <> "_" <> newName) (GoRaw ("func(" <> goParamsNative <> ") " <> goTypeToStr fRet <> " {\\n" <> printGoExpr funcBody <> "\\n}"))`;

if (!code.includes(targetStr)) {
  console.error("Target string not found!");
  process.exit(1);
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.replace(targetStr, replacementStr));
console.log("Patched CodeGen.purs successfully.");
