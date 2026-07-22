import fs from 'fs';

let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const helper = `
isEffectNode :: forall a. BackendSyntax a -> Boolean
isEffectNode = case _ of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> true
  PrimEffect _ -> true
  _ -> false

executeIfOpaque :: forall a. BackendSyntax a -> GoExpr -> GoExpr
executeIfOpaque expr goExpr =
  if isEffectNode expr then goExpr
  else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ goExpr, GoRaw "gopurs_runtime.Value{}" ]
`;

code = code.replace(
  'translateExprImpl ::',
  helper + '\ntranslateExprImpl ::'
);

code = code.replace(
  /EffectBind mbIdent lvl binding body \-\>\n        let\n          originalName = localId mbIdent lvl\n          name = originalName \<\> "_" \<\> show nextId\n          newBound = Map.insert originalName name bound\n          resBinding = translateExprImpl helpersRef \(depth \+ 1\) modNameStr recVars namedBound bound Nothing \[\] false \(nextId \+ 1\) binding\n          resBody = translateExprImpl helpersRef \(depth \+ 1\) modNameStr recVars namedBound newBound Nothing loopCtx isTail resBinding.nextId body\n        in\n          \{ stmts: resBinding.stmts \<\> StmtLeaf \(GoAssign name resBinding.expr\) \<\> resBody.stmts, expr: resBody.expr, nextId: resBody.nextId \}/,
  `EffectBind mbIdent lvl binding body ->
        let
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          newBound = Map.insert originalName name bound
          resBinding = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false (nextId + 1) binding
          resBody = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail resBinding.nextId body
          
          bindingExpr = executeIfOpaque (unwrap binding) resBinding.expr
          bodyExpr = executeIfOpaque (unwrap body) resBody.expr
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: bodyExpr, nextId: resBody.nextId }`
);

code = code.replace(
  /EffectDefer binding \-\>\n        let\n          resBinding = translateExprImpl helpersRef \(depth \+ 1\) modNameStr recVars namedBound bound Nothing \[\] false nextId binding\n          funcExpr = GoRaw \("gopurs_runtime.Func\\(func\\(_ gopurs_runtime.Value\\) gopurs_runtime.Value \{\\n" \<\> printGoExpr \(GoBlock \(flattenStmts resBinding.stmts \<\> \[ GoReturn resBinding.expr \]\)\) \<\> "\\n\}\\)"\)\n        in\n          \{ stmts: StmtEmpty, expr: funcExpr, nextId: resBinding.nextId \}/,
  `EffectDefer binding ->
        let
          resBinding = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false nextId binding
          bindingExpr = executeIfOpaque (unwrap binding) resBinding.expr
          funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\\n" <> printGoExpr (GoBlock (flattenStmts resBinding.stmts <> [ GoReturn bindingExpr ])) <> "\\n})")
        in
          { stmts: StmtEmpty, expr: funcExpr, nextId: resBinding.nextId }`
);

code = code.replace(
  /EffectPure binding \-\>\n        translateExprImpl helpersRef \(depth \+ 1\) modNameStr recVars namedBound bound Nothing \[\] false nextId binding/,
  `EffectPure binding ->
        translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false nextId binding`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
