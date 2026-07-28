const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  'mutateBlock = GoBlock (mutateStmts <> [ GoRaw "fmt.Println(\\"FBIP MUTATION TRIGGERED!\\")", GoReturn (GoVar deadVar) ])',
  'mutateBlock = GoBlock (mutateStmts <> [ GoReturn (GoVar deadVar) ])'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
