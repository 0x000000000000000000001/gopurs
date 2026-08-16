const fs = require('fs');
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
if (content.includes('unboxGoExpr (GoCall (GoSelector (GoVar "gopurs_runtime") "Int")')) {
  console.log("Already has optimization");
} else {
  console.log("No optimization");
}
