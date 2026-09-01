const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const target1 = `                                  isSelfRecursiveLoop = group.recursive && Array.length group.bindings == 1
                                  currentLoopCtx = if isSelfRecursiveLoop then [ { ident: fn.ident, params: map fst paramsWithTypes, loopParams: map (\\p -> fst p <> "_loop") paramsWithTypes, goTypes: map snd paramsWithTypes } ] else []
                                  resBodyMut = translateExprImpl_ helpersRef 0 modNameStr recVars moduleArities newBound (Just fn.ident) currentLoopCtx isSelfRecursiveLoop false 0 fn.body`;

const replacement1 = `                                  isSelfRecursiveLoop = group.recursive && Array.length group.bindings == 1
                                  fRet = case extractExprFuncType (getExprType fn.val) of
                                    Just { fRet: rt } -> exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr rt
                                    Nothing -> TypeValue
                                  currentLoopCtx = if isSelfRecursiveLoop then [ { ident: fn.ident, params: map fst paramsWithTypes, loopParams: map (\\p -> fst p <> "_loop") paramsWithTypes, goTypes: map snd paramsWithTypes, fRet } ] else []
                                  resBodyMut = translateExprImpl_ helpersRef 0 modNameStr recVars moduleArities newBound (Just fn.ident) currentLoopCtx isSelfRecursiveLoop false 0 fn.body`;

if (!code.includes(target1)) {
  console.error("Target string not found!");
  process.exit(1);
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.split(target1).join(replacement1));
console.log("Patched CodeGen.purs successfully.");
