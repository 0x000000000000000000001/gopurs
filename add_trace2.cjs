const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
`          realBinding = let _ = Debug.trace ("BINDING SHAPE BEFORE: " <> printTcoExprShape binding <> " AFTER: " <> printTcoExprShape (stripEffectDefer binding)) (\\_ -> unit) in stripEffectDefer binding`,
`          realBinding = stripEffectDefer binding`
);

code = code.replace(
`        in
          { stmts: resBinding.stmts <> foldMap StmtLeaf declStmts <> resBody.stmts, expr: bodyExpr, exprType: resBody.exprType, nextId: resBody.nextId }`,
`        in
          { stmts: StmtLeaf (GoRaw ("/* SHAPE BEFORE: " <> printTcoExprShape binding <> " AFTER: " <> printTcoExprShape realBinding <> " */\\n")) <> resBinding.stmts <> foldMap StmtLeaf declStmts <> resBody.stmts, expr: bodyExpr, exprType: resBody.exprType, nextId: resBody.nextId }`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
