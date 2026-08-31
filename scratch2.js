const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  `Console.log ("lookup fn.args: " <> show fn.args <> " fArgs length: " <> show (Array.length (fromMaybe [] (map _.fArgs (Map.lookup goName moduleArities)))) <> " fArgsGo length: " <> show (Array.length paramsWithTypes))`,
  `Console.log ("lookup fn.args: " <> show fn.args <> " expectedExprType: " <> printExprType (getExprType fn.val))`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
