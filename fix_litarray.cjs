const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  '( \\acc val ->\n                let\n                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] valLiveOut false false acc.nextId val',
  '( \\acc val ->\n                let\n                  tailArgs = Array.drop 1 acc.remainingArgs\n                  valLiveOut = Set.union liveOut (Set.unions (map freeVars tailArgs))\n                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] valLiveOut false false acc.nextId val'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
