const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /checkBlockStmts = StmtBlock \[ GoIfElse condType \(flattenStmts \(mutateStmts <> \[assignMut\]\)\) \[assignAlloc\] \]/,
  'checkBlockStmts = StmtLeaf (GoIfElse condType (flattenStmts (mutateStmts <> [assignMut])) [assignAlloc])'
);
code = code.replace(
  /stmts: accProps\.stmts <> StmtList \[declTmp\] <> checkBlockStmts/,
  'stmts: accProps.stmts <> StmtLeaf declTmp <> checkBlockStmts'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
