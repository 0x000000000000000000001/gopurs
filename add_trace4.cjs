const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
`          { stmts: StmtLeaf (GoRaw ("/* SHAPE BEFORE: " <> printTcoExprShape binding <> " AFTER: " <> printTcoExprShape realBinding <> " */\\n")) <> resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: bodyExpr, exprType: resBody.exprType, nextId: resBody.nextId }`,
`          { stmts: StmtLeaf (GoRaw ("_ = \\"SHAPE BEFORE: " <> printTcoExprShape binding <> " AFTER: " <> printTcoExprShape realBinding <> "\\"")) <> resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: bodyExpr, exprType: resBody.exprType, nextId: resBody.nextId }`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
