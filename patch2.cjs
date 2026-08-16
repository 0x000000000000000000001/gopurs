const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(/_ = unsafePerformEffect \\(Console.log \\("\\unwrapFunc: " <> show name <> " args: " <> show \\(Array.length args\\) <> " fArgs: " <> show \\(Array.length \\(fromMaybe \\[\\] \\(map \\_\\.fArgs typeSig\\)\\)\\)\\)\\)/g, '');
code = code.replace(/fRetGo = case typeSig of/g, `fRetGo = let _ = unsafePerformEffect (Console.log ("unwrapFunc: " <> show name <> " args: " <> show (Array.length args) <> " fArgs: " <> show (Array.length (fromMaybe [] (map _.fArgs typeSig))))) in case typeSig of`);
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
