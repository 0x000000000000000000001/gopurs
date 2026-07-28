const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  '{ stmts: accProps.stmts, expr: GoRaw ("func() gopurs_runtime.Value {\\n" <> printGoExpr ifStmt <> "\\n}()"), exprType: expectedGoType, nextId: accProps.nextId }',
  '{ stmts: accProps.stmts, expr: GoRaw ("func() gopurs_runtime.Value {\\n" <> printGoExpr ifStmt <> "\\n}()"), exprType: expectedGoType, nextId: accProps.nextId, reusedVars: Set.insert deadVar accProps.reusedVars }'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
