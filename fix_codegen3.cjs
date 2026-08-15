const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
`          stripEffectDefer (TcoExpr a syn) = case syn of
            EffectDefer inner -> stripEffectDefer inner
            Let ident lvl val body -> TcoExpr a (Let ident lvl val (stripEffectDefer body))
            LetRec lvl bindings body -> TcoExpr a (LetRec lvl bindings (stripEffectDefer body))
            _ -> TcoExpr a syn`,
`          stripEffectDefer (TcoExpr a syn) = case syn of
            EffectDefer inner -> stripEffectDefer inner
            Abs _ inner -> stripEffectDefer inner
            Let ident lvl val body -> TcoExpr a (Let ident lvl val (stripEffectDefer body))
            LetRec lvl bindings body -> TcoExpr a (LetRec lvl bindings (stripEffectDefer body))
            _ -> TcoExpr a syn`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
