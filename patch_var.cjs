const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(/\{ stmts: StmtEmpty, expr: unboxGoExpr rawCall TypeValue vType, exprType: vType, nextId \}/g, 'let actualExprType = case vType of TypeFunc _ _ -> TypeValue; _ -> vType in { stmts: StmtEmpty, expr: unboxGoExpr rawCall TypeValue vType, exprType: actualExprType, nextId }');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
