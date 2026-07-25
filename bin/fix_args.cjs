const fs = require('fs');
let content = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');

content = content.replace(
  'translate elidedCtors ctorTypes importsArray mod =',
  'translate elidedCtors ctorTypes instantiations importsArray mod ='
);

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', content);
