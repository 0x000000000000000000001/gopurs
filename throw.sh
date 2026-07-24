#!/bin/bash
sed -i '' 's/isTailCallTo =/isTailCallTo = let _ = if isTail then Effect.Exception.Unsafe.unsafeThrow ("TCO Trace: flatFn = " <> show (getVar (unwrapExpr flatFn)) <> " getFuncType = " <> show (getFuncType (unwrapExpr flatFn))) else unit in/g' src/Gopurs/CodeGen.purs
npm run build && node bin/gopurs.js --main Test.TCO > compile.log 2>&1
cat compile.log | grep -A 5 "TCO Trace"
git checkout src/Gopurs/CodeGen.purs
