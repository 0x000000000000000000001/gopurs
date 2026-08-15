const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
`          resBinding = Debug.trace ("BINDING FOR " <> name <> ": " <> printTcoExpr binding) \\_ -> translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false true (nextId + 1) realBinding`,
`          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false true (nextId + 1) realBinding`
);

code = code.replace(
`          stripEffectDefer (TcoExpr a syn) = case syn of
            EffectDefer inner -> stripEffectDefer inner
            Abs _ inner -> stripEffectDefer inner
            Let ident lvl val body -> TcoExpr a (Let ident lvl val (stripEffectDefer body))
            LetRec lvl bindings body -> TcoExpr a (LetRec lvl bindings (stripEffectDefer body))
            _ -> TcoExpr a syn`,
`          stripEffectDefer e@(TcoExpr a syn) = 
             let _ = Debug.trace ("STRIPPING: " <> printTcoExpr e) (\\_ -> unit)
             in case syn of
              EffectDefer inner -> stripEffectDefer inner
              Abs _ inner -> stripEffectDefer inner
              Let ident lvl val body -> TcoExpr a (Let ident lvl val (stripEffectDefer body))
              LetRec lvl bindings body -> TcoExpr a (LetRec lvl bindings (stripEffectDefer body))
              _ -> TcoExpr a syn`
);

if (!code.includes('printTcoExpr')) {
    code += `
printTcoExpr :: TcoExpr -> String
printTcoExpr e = case unwrapExpr e of
  EffectDefer inner -> "EffectDefer (" <> printTcoExpr inner <> ")"
  EffectBind _ _ binding body -> "EffectBind (" <> printTcoExpr binding <> ") (" <> printTcoExpr body <> ")"
  EffectPure inner -> "EffectPure (...)"
  Let ident lvl val body -> "Let (" <> printTcoExpr val <> ") (" <> printTcoExpr body <> ")"
  LetRec lvl bindings body -> "LetRec (...) (" <> printTcoExpr body <> ")"
  Abs mbIdent body -> "Abs (" <> printTcoExpr body <> ")"
  App f arg -> "App (" <> printTcoExpr f <> ") (" <> printTcoExpr arg <> ")"
  Var qual -> "Var"
  Literal lit -> "Literal"
  PrimEffect _ -> "PrimEffect"
  Local ident lvl -> "Local"
  _ -> "Other"
`;
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
