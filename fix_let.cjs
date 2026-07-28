const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  '                let tailProps = Array.drop 1 acc.remainingProps\n                    valLiveOut = Set.union liveOut (Set.unions (map (\\(Tuple _ p) -> freeVars p) tailProps))\n                let\n                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] Set.empty false false acc.nextId val',
  '                let\n                  tailProps = Array.drop 1 acc.remainingProps\n                  valLiveOut = Set.union liveOut (Set.unions (map (\\(Tuple _ p) -> freeVars p) tailProps))\n                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] valLiveOut false false acc.nextId val'
);

code = code.replace(
  '{ stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs coercedExpr, exprType: TypeValue, nextId: resVal.nextId, fieldIdx: acc.fieldIdx + 1 }',
  '{ stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs coercedExpr, exprType: TypeValue, nextId: resVal.nextId, fieldIdx: acc.fieldIdx + 1, remainingProps: tailProps }'
);

code = code.replace(
  '{ stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId, fieldIdx: 0 }',
  '{ stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId, fieldIdx: 0, remainingProps: props }'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
