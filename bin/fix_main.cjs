const fs = require('fs');
let content = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'utf8');

if (!content.includes('import Gopurs.Monomorphize')) {
  content = content.replace('import Gopurs.FfiSupport (findFfiFile, appendFfiWrappers)', 'import Gopurs.FfiSupport (findFfiFile, appendFfiWrappers)\nimport Gopurs.Monomorphize (collectInstantiations)');
}

content = content.replace('  buildModules', '  let instantiations = foldl collectInstantiations Map.empty finalModules\n  FS.writeTextFile UTF8 "instantiations.txt" (show (Array.fromFoldable (Map.keys instantiations)))\n\n  buildModules');

content = content.replace('let goFile = translate elidedCtors ctorTypes importsArray backendMod', 'let goFile = translate elidedCtors ctorTypes instantiations importsArray backendMod');

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', content);
