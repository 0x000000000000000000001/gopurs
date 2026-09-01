const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const targetStr = `                    combinedLoopCtx = loopCtxs <> loopCtx

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
                            nativeAssignment = GoMutate ("Call_local_" <> modNameStr <> "_" <> newName) (GoRaw ("func(" <> goParamsNative <> ") " <> goTypeToStr fRet <> " {\\n" <> printGoExpr funcBody <> "\\n}"))

                            nativeCallExpr = GoCall (GoVar ("Call_local_" <> modNameStr <> "_" <> newName)) (map (\\(Tuple p goT) -> unboxGoExpr (GoVar (p <> "_loop_val")) TypeValue goT) paramsWithTypes)
                            funcExpr = Array.foldr (\\(Tuple p goT) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_loop_val gopurs_runtime.Value) gopurs_runtime.Value {\\nreturn " <> printGoExpr acc <> "\\n}") ]) nativeCallExpr paramsWithTypes
                          in
                            Tuple (accStmts <> [ nativeAssignment, GoMutate newName funcExpr ]) resBodyMut.nextId
                      )
                      (Tuple [] allocRes.nextId)
                      fns

                    resBodyOuter = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars localModuleArities allocRes.newBound Nothing loopCtx isTail inEffectBlock nextId' body`;

const replacementStr = `                    combinedLoopCtx = loopCtxs <> loopCtx

                    resData = foldl
                      ( \\acc fn ->
                          let
                            oldName = localId (Just (Ident fn.ident)) lvl
                            boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                            newName = boundInfo.name
                            goType = boundInfo.goType
                            pTypes = paramTypes fn.body
                            paramsWithTypes = map
                              ( \\idStr ->
                                  let
                                    t = fromMaybe Any (Map.lookup idStr pTypes)
                                  in
                                    Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr t)
                              )
                              fn.args
                            currentLoopCtx = [ { ident: newName, params: fn.args, loopParams: map (\\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes, fRet: TypeValue } ]
                            loopBound = foldl (\\acc2 (Tuple idStr goT) -> Map.insert idStr { name: idStr, goType: goT } acc2) allocRes.newBound paramsWithTypes
                            resBodyMut = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars acc.modArities loopBound (Just newName) currentLoopCtx true false acc.nextId fn.body
                            trueFRet = resBodyMut.exprType

                            loopParams = map (\\(Tuple idStr _) -> idStr <> "_loop") paramsWithTypes
                            initVars = Array.concatMap (\\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes

                            funcBody = GoFor newName (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn resBodyMut.expr ])

                            goParamsNative = String.joinWith ", " (map (\\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)
                            nativeAssignment = GoMutate ("Call_local_" <> modNameStr <> "_" <> newName) (GoRaw ("func(" <> goParamsNative <> ") " <> goTypeToStr trueFRet <> " {\\n" <> printGoExpr funcBody <> "\\n}"))

                            nativeCallExpr = GoCall (GoVar ("Call_local_" <> modNameStr <> "_" <> newName)) (map (\\(Tuple p goT) -> unboxGoExpr (GoVar (p <> "_loop_val")) TypeValue goT) paramsWithTypes)
                            funcExpr = Array.foldr (\\(Tuple p goT) accExpr -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_loop_val gopurs_runtime.Value) gopurs_runtime.Value {\\nreturn " <> printGoExpr accExpr <> "\\n}") ]) (boxGoExpr nativeCallExpr trueFRet) paramsWithTypes
                            
                            newArities = Map.insert newName { fullName: "Call_local_" <> modNameStr <> "_" <> newName, fArgs: map snd paramsWithTypes, fRet: trueFRet, arity: Array.length fn.args } acc.modArities
                            newBound2 = Map.insert oldName { name: newName, goType: TypeFunc (map snd paramsWithTypes) trueFRet } acc.newBound
                            declStmtsLocal = [ GoRaw ("var Call_local_" <> modNameStr <> "_" <> newName <> " func(" <> String.joinWith ", " (map goTypeToStr (map snd paramsWithTypes)) <> ") " <> goTypeToStr trueFRet), GoRaw ("_ = Call_local_" <> modNameStr <> "_" <> newName), GoRaw ("var " <> newName <> " gopurs_runtime.Value"), GoRaw ("_ = " <> newName) ]
                          in
                            { stmts: acc.stmts <> declStmtsLocal <> [ nativeAssignment, GoMutate newName funcExpr ], nextId: resBodyMut.nextId, modArities: newArities, newBound: newBound2 }
                      )
                      { stmts: [], nextId: allocRes.nextId, modArities: moduleArities, newBound: allocRes.newBound }
                      fns

                    resBodyOuter = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars resData.modArities resData.newBound Nothing loopCtx isTail inEffectBlock resData.nextId body`;

if (!code.includes(targetStr.substring(0, 50))) {
  console.error("Target string not found!");
  process.exit(1);
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.replace(targetStr, replacementStr));
console.log("Patched CodeGen.purs successfully.");
