const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(/synName :: TcoExpr -> String[\s\S]*?_ -> "Other"/, '');
code = code.replace(/(import Gopurs\.OptimizeTAST[^\n]*\n)/, `$1\n
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
`);
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
