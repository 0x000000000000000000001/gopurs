const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  '{ stmts: StmtEmpty, exprs: [], fieldIdx: 0, nextId: nextId }',
  '{ stmts: StmtEmpty, exprs: [], fieldIdx: 0, nextId: nextId, remainingProps: props }'
);

code = code.replace(
  '( \\acc (Tuple _ val) ->',
  '( \\acc (Tuple _ val) ->\n                let tailProps = Array.drop 1 acc.remainingProps\n                    valLiveOut = Set.union liveOut (Set.unions (map (\\(Tuple _ p) -> freeVars p) tailProps))'
);

code = code.replace(
  'resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] Set.empty false false acc.nextId val',
  'resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] valLiveOut false false acc.nextId val'
);

code = code.replace(
  'fieldIdx: acc.fieldIdx + 1, nextId: resVal.nextId }',
  'fieldIdx: acc.fieldIdx + 1, nextId: resVal.nextId, remainingProps: tailProps }'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
