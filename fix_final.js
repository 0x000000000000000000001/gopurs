import fs from 'fs';

let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// Replace the incorrect translateExprImplWorker definition
code = code.replace(
  'translateExprImplWorker inEffect modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr _ expr) = case expr of',
  'translateExprImplWorker inEffect helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) = case expr of'
);

// We also need to fix the `isEffectNode` wrapper.
code = code.replace(
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr _ expr) =',
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) ='
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
