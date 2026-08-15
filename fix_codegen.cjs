const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
`isEffectNode :: TcoExpr -> Boolean
isEffectNode tcoExpr = case unwrapTcoExpr expr of`,
`isEffectNode :: TcoExpr -> Boolean
isEffectNode expr = case unwrapTcoExpr expr of`
);

code = code.replace(
`executeIfOpaque :: forall a. BackendSyntax a -> GoExpr -> GoExpr

executeIfOpaque expr goExpr =
  if isEffectNode tcoExpr then goExpr`,
`executeIfOpaque :: TcoExpr -> GoExpr -> GoExpr

executeIfOpaque expr goExpr =
  if isEffectNode expr then goExpr`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
