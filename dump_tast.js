const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(/fRetGo = let _ = unsafePerformEffect \\(Console.log \\("\\unwrapFunc/g, `
                  _ = if show name == "(Ident \\"contains\\")" then unsafePerformEffect (Console.log ("TAST of contains: " <> show val)) else unit
                  fRetGo = let _ = unsafePerformEffect (Console.log ("unwrapFunc`);
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
