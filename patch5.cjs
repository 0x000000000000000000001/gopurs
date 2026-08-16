const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(`_trace_msg = unsafePerformEffect (Console.log ("=== APP TRACE ===\\nresFn.exprType: " <> goTypeToStr resFn.exprType <> "\\nflatArgs len: " <> show (Array.length flatArgs)))`, `_trace_msg = unsafePerformEffect (Console.log ("=== APP TRACE ===\\nresFn.exprType: " <> goTypeToStr resFn.exprType <> "\\nflatArgs len: " <> show (Array.length flatArgs)))`);
// wait, let's use Debug.trace
code = code.replace(`import Effect.Console as Console`, `import Debug as Debug`);
code = code.replace(`_trace_msg = unsafePerformEffect (Console.log ("=== APP TRACE ===\\nresFn.exprType: " <> goTypeToStr resFn.exprType <> "\\nflatArgs len: " <> show (Array.length flatArgs)))\n                            finalExprType = case resFn.exprType of`, `finalExprType = Debug.trace ("=== APP TRACE ===\\nresFn.exprType: " <> goTypeToStr resFn.exprType <> "\\nflatArgs len: " <> show (Array.length flatArgs)) \\_ -> case resFn.exprType of`);
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
