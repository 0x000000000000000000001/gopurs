import fs from 'fs';

let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// Restore original just in case it was half patched
content = content.replace(
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }',
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }'
);
content = content.replace(
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) =\n' +
  '  let\n' +
  '    isEff = isEffectNode expr\n' +
  '  in\n' +
  '    if isEff && not inEffectBlock then\n' +
  '      let\n' +
  '        res = translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx false true nextId tcoExpr\n' +
  '        funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\\n" <> printGoExpr (GoBlock (flattenStmts res.stmts <> [ GoReturn res.expr ])) <> "\\n})")\n' +
  '      in\n' +
  '        { stmts: StmtEmpty, expr: funcExpr, nextId: res.nextId }\n' +
  '    else\n',
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =\n'
);

// Now do the proper rename!
// 1. Rename ALL `translateExprImpl` to `translateExprImpl_`
content = content.replace(/translateExprImpl /g, 'translateExprImpl_ ');
content = content.replace(/translateExprImpl\n/g, 'translateExprImpl_\n');
content = content.replace(/translateExprImpl\(/g, 'translateExprImpl_(');
content = content.replace(/translateExprImpl::/g, 'translateExprImpl_::');
content = content.replace(/translateExprImpl ::/g, 'translateExprImpl_ ::');

// 2. Define `translateExprImpl` wrapper that calls `translateExprImpl_` with inEffectBlock = false
let wrapper = `
translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr =
  translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail false nextId tcoExpr

translateExprImpl_ :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) =
  let
    isEff = isEffectNode expr
  in
    if isEff && not inEffectBlock then
      let
        res = translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx false true nextId tcoExpr
        funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\\n" <> printGoExpr (GoBlock (flattenStmts res.stmts <> [ GoReturn res.expr ])) <> "\\n})")
      in
        { stmts: StmtEmpty, expr: funcExpr, nextId: res.nextId }
    else
`;

content = content.replace(
  'translateExprImpl_ :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }\n' +
  'translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr tcoAnalysis expr) =',
  wrapper
);

// 3. Find EffectBind and modify it to pass `true` to `inEffectBlock`
// Inside EffectBind:
// resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false (nextId + 1) binding
// resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail resBinding.nextId body
// We need to change the recursive calls inside EffectBind to pass `true` instead of `false` for `inEffectBlock`!
// Actually, since they all call `translateExprImpl_`, we can just replace them.
content = content.replace(
  /EffectBind mbIdent lvl binding body ->\s+let\s+originalName = localId mbIdent lvl\s+name = originalName <> "_" <> show nextId\s+newBound = Map.insert originalName name bound\s+resBinding = translateExprImpl_ helpersRef \(depth \+ 1\) modNameStr recVars namedBound bound Nothing \[\] false \(nextId \+ 1\) binding\s+resBody = translateExprImpl_ helpersRef \(depth \+ 1\) modNameStr recVars namedBound newBound Nothing loopCtx isTail resBinding.nextId body/g,
  `EffectBind mbIdent lvl binding body ->
        let
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          newBound = Map.insert originalName name bound
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false true (nextId + 1) binding
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail true resBinding.nextId body`
);

// Note: `isTail` is passed as `isTail` to resBody. We need to insert `true` before it!
// For EffectBind, the original was:
// resBody = translateExprImpl_ ... loopCtx isTail resBinding.nextId body
// We changed it to pass `true` before `resBinding.nextId`!

// Wait! We ALSO need to update the definition of liftIfNeeded!
// Inside translateExprImpl_, `liftIfNeeded` calls `translateExprImpl_`!
content = content.replace(
  'let res = translateExprImpl_ helpersRef 0 modNameStr recVars namedBound bound Nothing [] false newNextId tcoExpr',
  'let res = translateExprImpl_ helpersRef 0 modNameStr recVars namedBound bound Nothing [] false inEffectBlock newNextId tcoExpr'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
