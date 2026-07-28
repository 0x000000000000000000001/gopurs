const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// We need to add `liveOut` to translateExprImpl_ signature
code = code.replace(
  'translateExprImpl_ :: Ref { decls :: Array GoDecl',
  'import Data.Set (Set)\nimport Data.Set as Set\nimport Gopurs.UsageAnalysis (freeVars)\n\ntranslateExprImpl_ :: Ref { decls :: Array GoDecl'
);

code = code.replace(
  'Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean',
  'Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Set String -> Boolean'
);

code = code.replace(
  'translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) =',
  'translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx liveOut isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) ='
);

// We need to replace all calls to translateExprImpl_
// There are calls like: translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail false nextId tcoExpr
// We will replace them with a function that computes the correct liveOut.
// Actually, it's easier to manually review the calls.
