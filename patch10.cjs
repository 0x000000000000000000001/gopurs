const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(`        Typed t inner ->`, `        Typed t inner ->\\n          if isStandardPursFunc (getExprType inner) then\\n             translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound expectedGoType loopCtx isTail inEffectBlock nextId inner\\n          else`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
