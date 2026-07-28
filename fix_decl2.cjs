const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx liveOut isTail inEffectBlock nextId tcoExpr/,
  'translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId tcoExpr'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
