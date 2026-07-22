import fs from 'fs';
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

content = content.replace(
  'executeIfOpaque :: forall a. BackendSyntax a -> GoExpr -> GoExpr\n' +
  'unwrapTcoExpr :: TcoExpr -> BackendSyntax TcoExpr\n' +
  'unwrapTcoExpr (TcoExpr _ expr) = expr\n' +
  'executeIfOpaque expr goExpr =',
  'unwrapTcoExpr :: TcoExpr -> BackendSyntax TcoExpr\n' +
  'unwrapTcoExpr (TcoExpr _ expr) = expr\n' +
  'executeIfOpaque :: forall a. BackendSyntax a -> GoExpr -> GoExpr\n' +
  'executeIfOpaque expr goExpr ='
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
