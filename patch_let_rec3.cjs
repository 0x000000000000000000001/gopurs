const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const targetStr = `                    resBodyOuter = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars resData.modArities resData.newBound Nothing loopCtx isTail inEffectBlock resData.nextId body
                  in
                    { stmts: foldMap StmtLeaf declStmts <> foldMap StmtLeaf fnWrapperStmts <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }`;

const replacementStr = `                    resBodyOuter = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars resData.modArities resData.newBound Nothing loopCtx isTail inEffectBlock resData.nextId body
                  in
                    { stmts: foldMap StmtLeaf resData.stmts <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }`;

if (!code.includes(targetStr)) {
  console.error("Target string not found!");
  process.exit(1);
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.replace(targetStr, replacementStr));
console.log("Patched CodeGen.purs successfully.");
