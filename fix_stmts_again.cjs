const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /checkBlockStmts = StmtBlock \[ GoIfElse condType \(mutateStmts <> \[assignMut\]\) \[assignAlloc\] \]/,
  'checkBlockStmts = StmtLeaf (GoIfElse condType (mutateStmts <> [assignMut]) [assignAlloc])'
);
code = code.replace(
  /stmts: accProps\.stmts <> StmtLeaf \(GoRaw \("\/\/ deadVarOpt: MUTATION"\)\) <> StmtList \[declTmp\] <> checkBlockStmts/,
  'stmts: accProps.stmts <> StmtLeaf (GoRaw ("// deadVarOpt: MUTATION")) <> declTmp <> checkBlockStmts'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
