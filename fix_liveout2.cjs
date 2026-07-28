const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// For Branch (Line 1428 approx):
//           resDef = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing loopCtx Set.empty isTail false nextId def
// We need it to be `liveOut` instead of `Set.empty` for resDef.
code = code.replace(
  'resDef = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing loopCtx Set.empty isTail false nextId def',
  'resDef = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing loopCtx liveOut isTail false nextId def'
);

// For Branch conditions (resCond):
code = code.replace(
  'resCond = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] Set.empty false false acc.nextId condExpr',
  'resCond = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] (Set.union (freeVars bodyExpr) liveOut) false false acc.nextId condExpr'
);

// For Branch bodies (resBody):
code = code.replace(
  'resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing loopCtx Set.empty isTail false resCond.nextId bodyExpr',
  'resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing loopCtx liveOut isTail false resCond.nextId bodyExpr'
);

// For Let bindings (resBinding):
code = code.replace(
  'resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] Set.empty false false nextId val',
  'resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] (Set.union (Set.difference (freeVars body) (Set.singleton originalName)) liveOut) false false nextId val'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
console.log("Fixed liveOut threading");
