const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const traceCode = `
printTcoExprShape :: TcoExpr -> String
printTcoExprShape e = case unwrapTcoExpr e of
  EffectDefer inner -> "EffectDefer(" <> printTcoExprShape inner <> ")"
  Abs _ inner -> "Abs(" <> printTcoExprShape inner <> ")"
  Let _ _ _ body -> "Let(" <> printTcoExprShape body <> ")"
  LetRec _ _ body -> "LetRec(" <> printTcoExprShape body <> ")"
  EffectBind _ _ _ body -> "EffectBind(" <> printTcoExprShape body <> ")"
  EffectPure _ -> "EffectPure"
  App fn arg -> "App(" <> printTcoExprShape fn <> " " <> printTcoExprShape arg <> ")"
  _ -> "Other"
`;

code = code.replace('extractExprFuncType :: ExprType -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }', traceCode + '\nextractExprFuncType :: ExprType -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }');

code = code.replace(
`          realBinding = stripEffectDefer binding`,
`          realBinding = let _ = Debug.trace ("BINDING SHAPE BEFORE: " <> printTcoExprShape binding <> " AFTER: " <> printTcoExprShape (stripEffectDefer binding)) (\\_ -> unit) in stripEffectDefer binding`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
