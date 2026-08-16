const fs = require('fs');
const file = 'src/Gopurs/CodeGen.purs';
let content = fs.readFileSync(file, 'utf8');
content = content.replace(
  /translateTopLevelDecl helpersRef modNameStr recVars moduleArities bound fn =/,
  `translateTopLevelDecl helpersRef modNameStr recVars moduleArities bound fn =\n  Debug.trace ("TRANSLATING: " <> sanitizeName (case fn of (NonRec (Ident name) _) -> name; (Rec arr) -> case Array.head arr of Just (Tuple (Ident name) _) -> name; _ -> ""; _ -> "")) \\_ ->`
);
fs.writeFileSync(file, content);
