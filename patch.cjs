const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// We need to trace finalExprType and resFn.exprType in App
const target = `                            finalExprType = case resFn.exprType of`;
const replacement = `                            _trace_msg = unsafePerformEffect (Console.log ("=== APP TRACE ===\\nflatFn: " <> show (unwrapTcoExpr flatFn) <> "\\nresFn.exprType: " <> show resFn.exprType <> "\\nflatArgs len: " <> show (Array.length flatArgs)))
                            finalExprType = case resFn.exprType of`;

if (!code.includes('_trace_msg')) {
  code = code.replace(target, replacement);
  // Add imports
  code = `import Effect.Console as Console\n` + code;
  fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
  console.log("Patched!");
} else {
  console.log("Already patched.");
}
