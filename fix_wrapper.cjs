const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  'translateExprImpl helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail nextId tcoExpr =',
  'translateExprImpl helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx liveOut isTail nextId tcoExpr ='
);

code = code.replace(
  'translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx Set.empty isTail false nextId tcoExpr',
  'translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx liveOut isTail false nextId tcoExpr'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
