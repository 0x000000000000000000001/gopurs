import fs from 'fs';
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

content = content.replace(
  `      EffectBind mbIdent lvl binding body ->
        let
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          newBound = Map.insert originalName name bound
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false true (nextId + 1) binding
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail true resBinding.nextId body
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoAssign name resBinding.expr) <> resBody.stmts, expr: resBody.expr, nextId: resBody.nextId }`,
  `      EffectBind mbIdent lvl binding body ->
        let
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          newBound = Map.insert originalName name bound
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false true (nextId + 1) binding
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail true resBinding.nextId body
          bindingExpr = executeIfOpaque (unwrapTcoExpr binding) resBinding.expr
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: resBody.expr, nextId: resBody.nextId }`
);

// Note: I DID NOT wrap it in a thunk!
// Wait! If I wrap it in a thunk inside EffectBind, then EffectBind ALWAYS returns a thunk!
// BUT if we wrap it in a thunk inside EffectBind, does it mean we double-wrap?
// NO! The `translateExprImpl_` wrapper ALREADY checks `isEff && not inEffectBlock`!
// If we are `not inEffectBlock`, it WILL wrap it!
// If we ARE `inEffectBlock`, it should NOT be wrapped!
// But wait! If `EffectBind` is used as a statement inside another `EffectBind` (where `inEffectBlock` is true), we DO NOT want to wrap it!
// So `EffectBind` SHOULD NOT wrap its result in a thunk!!!
// It should just emit the statements!
// Wait! If it just emits the statements, then when `isEff && not inEffectBlock` triggers, the WRAPPER will wrap `EffectBind` in a thunk!
// YES! The wrapper will put ALL the statements inside `GoRaw("func ...")`!
// So my wrapper ALREADY handles the thunking perfectly!
// All I need to do is `executeIfOpaque`!

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
