const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(/fRetGo = case typeSig of/, `_ = unsafePerformEffect (Console.log ("unwrapFunc: " <> show name <> " args: " <> show (Array.length args) <> " fArgs: " <> show (Array.length (fromMaybe [] (map _.fArgs typeSig)))))\n                  fRetGo = case typeSig of`);
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
