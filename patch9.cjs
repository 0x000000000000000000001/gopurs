const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(/synName = \(\\\(TcoExpr _ syn\) -> case syn of \{ Typed _ _ -> "Typed"; PrimOp _ -> "PrimOp"; App _ _ -> "App"; Var _ -> "Var"; Let _ _ _ _ -> "Let"; LetRec _ _ _ -> "LetRec"; Abs _ _ -> "Abs"; Lit _ -> "Lit"; CtorDef _ _ _ _ -> "CtorDef"; _ -> "Other" \}\) binding/,
`synName = getSynName binding`);

code += `\n
getSynName :: TcoExpr -> String
getSynName (TcoExpr _ syn) = case syn of
  Typed _ _ -> "Typed"
  PrimOp _ -> "PrimOp"
  App _ _ -> "App"
  Var _ -> "Var"
  Let _ _ _ _ -> "Let"
  LetRec _ _ _ -> "LetRec"
  Abs _ _ -> "Abs"
  Lit _ -> "Lit"
  CtorDef _ _ _ _ -> "CtorDef"
  _ -> "Other"
`;

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
