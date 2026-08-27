const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  `                            finalExprType = case resFn.exprType of`,
  `                            finalExprTypeStr = case resFn.exprType of
                              TypeFunc fA fR -> "TypeFunc len:" <> show (Array.length fA) <> " ret:" <> goTypeToStr fR
                              _ -> "Other " <> goTypeToStr resFn.exprType
                            _ = unsafePerformEffect (Console.log ("App on " <> printGoExpr resFn.expr <> " with args len " <> show (Array.length flatArgs) <> " fnType: " <> finalExprTypeStr))
                            finalExprType = case resFn.exprType of`
);
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
