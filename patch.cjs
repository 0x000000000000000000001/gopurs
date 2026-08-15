const fs = require('fs');
const code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
const synNameCode = `
synName :: TcoExpr -> String
synName (TcoExpr _ syn) = case syn of
  Typed _ _ -> "Typed"
  PrimOp _ -> "PrimOp"
  App _ _ -> "App"
  Var _ -> "Var"
  Let _ _ _ _ -> "Let"
  LetRec _ _ _ -> "LetRec"
  Abs _ _ -> "Abs"
  Lit _ -> "Lit"
  CtorDef _ _ _ _ _ -> "CtorDef"
  _ -> "Other"
`;
const newCode = code.replace(/getExprType :: TcoExpr -> ExprType/, synNameCode + '\ngetExprType :: TcoExpr -> ExprType');
fs.writeFileSync('src/Gopurs/CodeGen.purs', newCode);
