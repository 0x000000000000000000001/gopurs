const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
`          stripEffectDefer e@(TcoExpr a syn) = 
             let _ = Debug.trace ("STRIPPING: " <> printTcoExpr e) (\\_ -> unit)
             in case syn of`,
`          stripEffectDefer (TcoExpr a syn) = case syn of`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
