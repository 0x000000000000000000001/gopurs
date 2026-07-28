const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  '                 && not (Set.member v.name liveOut) \n                 && not (Set.member v.name (freeVars tcoExpr)) ',
  '                 && not (Set.member v.name (freeVars tcoExpr)) '
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
