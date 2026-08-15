const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
`          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false true (nextId + 1) realBinding`,
`          resBinding = Debug.trace ("BINDING FOR " <> name <> ": " <> printTcoExpr binding) \\_ -> translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false true (nextId + 1) realBinding`
);

if (!code.includes('printTcoExpr')) {
    code += `
printTcoExpr :: TcoExpr -> String
printTcoExpr e = case unwrapTcoExpr e of
  EffectDefer inner -> "EffectDefer (" <> printTcoExpr inner <> ")"
  EffectBind _ _ binding body -> "EffectBind (" <> printTcoExpr binding <> ") (" <> printTcoExpr body <> ")"
  EffectPure inner -> "EffectPure (...)"
  Let ident lvl val body -> "Let " <> ident <> " (" <> printTcoExpr val <> ") (" <> printTcoExpr body <> ")"
  LetRec lvl bindings body -> "LetRec (...) (" <> printTcoExpr body <> ")"
  Abs mbIdent body -> "Abs " <> show mbIdent <> " (" <> printTcoExpr body <> ")"
  App f arg -> "App (" <> printTcoExpr f <> ") (" <> printTcoExpr arg <> ")"
  Var qual -> "Var " <> show qual
  Literal lit -> "Literal"
  PrimEffect _ -> "PrimEffect"
  Local ident lvl -> "Local " <> show ident
  _ -> "Other"
`;
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
