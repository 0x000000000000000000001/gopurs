import fs from 'fs';

let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const newSig = `
translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr =
  translateExprImplWorker false helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr

translateExprImplWorker :: Boolean -> Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImplWorker inEffect helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =`;

code = code.replace(
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }\ntranslateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =',
  newSig
);

// Now for all occurrences of `translateExprImpl(` or `translateExprImpl helpersRef` inside translateExprImplWorker,
// we replace them with `translateExprImplWorker inEffect`.
// We have to be careful NOT to replace `translateExprImpl` in `translateDeclImpl`.
// Let's split by `translateDeclImpl` and only replace in the first part!

const parts = code.split('translateDeclImpl ::');
parts[0] = parts[0].replace(/translateExprImpl /g, 'translateExprImplWorker inEffect ');
// Restore the one we messed up in the signature
parts[0] = parts[0].replace('translateExprImplWorker inEffect ::', 'translateExprImpl ::');
parts[0] = parts[0].replace('translateExprImplWorker inEffect helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr =', 'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr =');
parts[0] = parts[0].replace('translateExprImplWorker inEffect false', 'translateExprImplWorker false');

// Now, for EffectBind, EffectPure, EffectDefer, PrimEffect, we want to replace `inEffect` with `true`.
// Let's do this by finding those nodes and wrapping them if !inEffect!

// Actually, wait! The EASIEST way is to do the wrapping AT THE TOP of `translateExprImplWorker`!
// If `isEffectNode expr` AND `!inEffect`, we return `gopurs_runtime.Func(...)` containing the evaluation with `inEffect=true`!
// Wait! If we just intercept it at the TOP of `translateExprImplWorker`, we don't need to change ANY recursive calls!
