const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace("import Effect.Console as Console\n", "");
code = code.replace("import Data.Tuple (Tuple(..), fst, snd)", "import Effect.Console as Console\nimport Data.Tuple (Tuple(..), fst, snd)");
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
