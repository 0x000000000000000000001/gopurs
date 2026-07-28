const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

function fixFoldl(code, targetArrayName, startString, endString) {
  let idx = 0;
  while ((idx = code.indexOf(startString, idx)) !== -1) {
    let nextIdx = code.indexOf(endString, idx);
    if (nextIdx === -1) break;
    nextIdx += endString.length;
    
    let block = code.substring(idx, nextIdx);
    
    // We expect block to look like:
    // ... foldl
    // ( \acc arg ->
    //     let
    //       argRes = translateExprImpl_ ... bound Nothing [] Set.empty false false acc.nextId arg
    //     in
    //       { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
    // )
    // { stmts: ... nextId ... }
    // targetArrayName
    
    let newBlock = block.replace(
      '( \\acc arg ->',
      `( \\acc arg ->\n                let tailArgs = Array.drop 1 acc.remainingArgs\n                    argLiveOut = Set.union liveOut (Set.unions (map freeVars tailArgs))`
    );
    
    // Replace Set.empty with argLiveOut
    newBlock = newBlock.replace(
      'bound Nothing [] Set.empty',
      'bound Nothing [] argLiveOut'
    );
    newBlock = newBlock.replace(
      'loopCtx Set.empty false false',
      'loopCtx argLiveOut false false'
    );
    
    // Add remainingArgs: tailArgs to the return object of the lambda
    newBlock = newBlock.replace(
      ', nextId: argRes.nextId }',
      ', nextId: argRes.nextId, remainingArgs: tailArgs }'
    );
    newBlock = newBlock.replace(
      ', nextId: resVal.nextId }',
      ', nextId: resVal.nextId, remainingArgs: tailArgs }'
    );
    
    // Add remainingArgs: targetArrayName to the initial accumulator
    // E.g. { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
    newBlock = newBlock.replace(
      ', nextId }',
      `, nextId, remainingArgs: toArray ${targetArrayName} }` // Use toArray just in case it's NonEmptyArray, though for Array it might fail if toArray is only for NonEmptyArray. Actually we can just use `(Array.fromFoldable ${targetArrayName})`.
    );
    newBlock = newBlock.replace(
      ', nextId: resFn.nextId }',
      `, nextId: resFn.nextId, remainingArgs: (Array.fromFoldable ${targetArrayName}) }`
    );
    // for xs in LitArray
    newBlock = newBlock.replace(
      ', nextId }\n            xs',
      `, nextId, remainingArgs: (Array.fromFoldable xs) }\n            xs`
    );
    
    code = code.substring(0, idx) + newBlock + code.substring(nextIdx);
    idx += newBlock.length;
  }
  return code;
}

// LitArray
code = fixFoldl(code, 'xs', 'accXs = foldl', 'xs');

// App (isTailCallTo)
code = fixFoldl(code, 'flatArgs', 'accArgs = foldl\\n                              ( \\\\acc arg ->', 'flatArgs');

// App (Intrinsic)
code = fixFoldl(code, 'args', 'accArgs = foldl\\n                  ( \\\\acc arg ->', 'args');

// UncurriedApp
code = fixFoldl(code, 'args', 'accArgs = foldl\\n                  ( \\\\acc arg ->', 'args');
code = fixFoldl(code, 'args', 'accArgs = foldl\\n            ( \\\\acc arg ->', 'args'); // UncurriedEffectApp

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
