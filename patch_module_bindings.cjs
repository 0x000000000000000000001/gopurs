const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const replace = `
                                  expectedRetType = case extractExprFuncType (getExprType fn.val) of
                                    Just { fArgs, fRet } -> 
                                      let 
                                        arityStr = Array.length fn.args
                                        arityFArgs = Array.length fArgs
                                        h = unsafePerformEffect (Ref.read helpersRef)
                                      in if arityStr < arityFArgs then
                                        TypeFunc (map (exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr) (Array.drop arityStr fArgs)) (exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr fRet)
                                      else
                                        exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr fRet
                                    Nothing -> TypeValue
`;

code = code.replace(/expectedRetType = case extractExprFuncType \(getExprType fn\.val\) of\n\s*Just \{ fRet \} -> exprTypeToGoType [^\n]+\n\s*Nothing -> TypeValue/, replace);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
