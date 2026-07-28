const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /            \)\n                        \)\n            \{ stmts: StmtEmpty, exprs: \[\], exprType: TypeValue, nextId, fieldIdx: 0 \}/,
  '            )\n            { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId, fieldIdx: 0 }'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
