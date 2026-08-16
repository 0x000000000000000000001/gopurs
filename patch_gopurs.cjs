const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
const search = `
unwrapTcoExpr (TcoExpr _ syn) = case syn of
  Typed _ e -> unwrapTcoExpr e
  _ -> syn`;
const replace = `
unwrapTcoExpr (TcoExpr _ syn) = case syn of
  Typed _ e -> unwrapTcoExpr e
  _ -> syn`;
// wait, I don't need to change unwrapTcoExpr if I can just change translateExprImpl_
