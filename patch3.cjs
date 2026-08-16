const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(`_trace_msg = unsafePerformEffect (Console.log ("=== APP TRACE ===\\nflatFn: " <> show (unwrapTcoExpr flatFn) <> "\\nresFn.exprType: " <> show resFn.exprType <> "\\nflatArgs len: " <> show (Array.length flatArgs)))`, `_trace_msg = unsafePerformEffect (Console.log ("=== APP TRACE ===\\nresFn.exprType: " <> show resFn.exprType <> "\\nflatArgs len: " <> show (Array.length flatArgs)))`);
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
