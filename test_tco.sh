#!/bin/bash
sed -i '' 's/resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing \[\] false false nextId flatFn/let _ = unsafePerformEffect $ Console.log ("mbDirectCall returned Nothing for flatFn: " <> show (getVar (unwrapExpr flatFn)) <> " type: " <> show (getFuncType (unwrapExpr flatFn)) <> " isTailCallTo: " <> show isTailCallTo) in\n                      resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId flatFn/g' src/Gopurs/CodeGen.purs
npm run build && node bin/gopurs.js --main AppX > compile.log 2>&1
cat compile.log | grep -A 2 -B 2 "mbDirectCall"
