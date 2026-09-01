const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const target1 = `Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }`;
const replacement1 = `Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType, fRet :: GoType } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }`;

const target2 = `Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Boolean -> Maybe ExprType -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }`;
const replacement2 = `Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType, fRet :: GoType } -> Boolean -> Boolean -> Maybe ExprType -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }`;


if (!code.includes(target1) || !code.includes(target2)) {
  console.error("Target string not found!");
  process.exit(1);
}

code = code.split(target1).join(replacement1);
code = code.split(target2).join(replacement2);
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
console.log("Patched CodeGen.purs successfully.");
