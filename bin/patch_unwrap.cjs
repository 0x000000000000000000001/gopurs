const fs = require('fs');
let content = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');

content = content.replace(
`                  fRetGo = case typeSig of
                    Just { fRet } -> exprTypeToGoType fRet
                    Nothing -> TypeValue
              in Just (Tuple name { fullName: name, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args })`,
`                  fRetGo = case typeSig of
                    Just { fRet } -> exprTypeToGoType fRet
                    Nothing -> TypeValue
              in if Array.all (_ /= TypeValue) goFArgs && goFRet /= TypeValue then
                Just (Tuple name { fullName: name, fArgs: goFArgs, fRet: goFRet, arity: Array.length fArgs })
              else Nothing`);

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', content);
