const fs = require('fs');

let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const newSig = `
translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr =
  translateExprImplWorker false helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr

translateExprImplWorker :: Boolean -> Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
`;

code = code.replace(
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }\n',
  newSig
);

code = code.replace(
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =',
  'translateExprImplWorker inEffect helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) ='
);

// We need to replace all `translateExprImpl` with `translateExprImplWorker inEffect` INSIDE translateExprImplWorker.
// Since translateExprImpl is only called inside its own body (except for one call in translateDeclImpl), we can just replace all occurrences of `translateExprImpl ` (with trailing space) to `translateExprImplWorker inEffect ` globally, and then fix translateDeclImpl if needed.
// Let's replace all `translateExprImpl ` -> `translateExprImplWorker inEffect `
code = code.replace(/translateExprImpl /g, 'translateExprImplWorker inEffect ');

// Then fix the ones that should NOT have `inEffect` (like in translateDeclImpl)
// translateDeclImpl uses `translateExprImpl ` ... wait, we can just let translateDeclImpl call translateExprImplWorker false.
code = code.replace(/translateDeclImpl(.*?)=([\s\S]*?)translateExprImplWorker inEffect/g, 'translateDeclImpl$1=$2translateExprImplWorker false');
code = code.replace(/translateDeclImpl(.*?)=([\s\S]*?)translateExprImplWorker inEffect/g, 'translateDeclImpl$1=$2translateExprImplWorker false');

// And in EffectBind / EffectPure / PrimEffect / EffectDefer, we should pass `true` instead of `inEffect`!
// Actually, let's just leave it as `inEffect`, and before EffectBind we wrap if !inEffect!
// YES! If !inEffect, we wrap the whole EffectBind in a thunk!
// Wait! If we wrap it in a thunk, do we need to pass `true` to its children?
// YES! Because its children are now INSIDE the thunk!
// So inside EffectBind / EffectPure / PrimEffect / EffectDefer, we MUST replace `translateExprImplWorker inEffect` with `translateExprImplWorker true`.

fs.writeFileSync('src/Gopurs/CodeGen.purs.tmp', code);
