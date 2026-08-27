const fs = require('fs');
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
content = content.replace(/          stripTyped :: TcoExpr -> BackendSyntax TcoExpr\n          stripTyped e@\(TcoExpr _ syn\) = case syn of\n            Typed _ inner -> stripTyped inner\n            Let _ _ _ inner -> stripTyped inner\n            LetRec _ _ inner -> stripTyped inner\n            _ -> syn\n\n        in case stripTyped a, expectedGoType of/, '        in case unwrapTcoExpr a, expectedGoType of');
fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
