const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const replace = `
                                  expectedRetType = case Map.lookup goName moduleArities of
                                    Just { fArgs, fRet } -> 
                                      if Array.length fn.args < Array.length fArgs then
                                        TypeFunc (Array.drop (Array.length fn.args) fArgs) fRet
                                      else
                                        fRet
                                    Nothing -> TypeValue
`;

code = code.replace(/expectedRetType = case Map\.lookup goName moduleArities of\n\s*Just \{ fRet \} -> fRet\n\s*Nothing -> TypeValue/, replace);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
