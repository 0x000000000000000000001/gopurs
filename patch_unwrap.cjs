const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const replace = `
                  fRetGo = case typeSig of
                    Just { fArgs, fRet } -> 
                      if Array.length args < Array.length fArgs then
                        TypeFunc (map (exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr) (Array.drop (Array.length args) fArgs)) (exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr fRet)
                      else
                        exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr fRet
                    Nothing -> TypeValue
`;

code = code.replace(/fRetGo = case typeSig of\n\s*Just \{ fArgs, fRet \} -> \n\s*if Array\.length args < Array\.length fArgs then\n\s*TypeValue\n\s*else\n\s*exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr fRet\n\s*Nothing -> TypeValue/, replace);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
