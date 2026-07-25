const fs = require('fs');
let content = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');

if (!content.includes('import Effect.Console as Console')) {
  content = content.replace('import Prelude', 'import Prelude\nimport Effect.Console as Console\nimport Effect.Unsafe (unsafePerformEffect)');
}

content = content.replace(
  '    expandBind (Tuple id@(Ident name) val) =\n      let qual = modNameStr <> "." <> name\n      in case Map.lookup qual instantiations of',
  '    expandBind (Tuple id@(Ident name) val) =\n      let\n        qual = modNameStr <> "." <> name\n        _ = unsafePerformEffect (Console.log ("EXPAND: " <> qual))\n      in case Map.lookup qual instantiations of'
);

content = content.replace(
  '           Just concretes ->',
  '           Just concretes ->\n             let _ = unsafePerformEffect (Console.log ("FOUND CONCRETES: " <> qual <> " -> " <> show (Set.size concretes))) in'
);

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', content);
