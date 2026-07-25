const fs = require('fs');
let content = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'utf8');

if (!content.includes('import Node.FS.Sync as FS')) {
  content = content.replace('import Effect.Console as Console', 'import Effect.Console as Console\nimport Node.FS.Sync as FS');
}

content = content.replace(
  '  Console.log (show (Array.fromFoldable (Map.keys instantiations)))',
  '  liftEffect $ FS.writeTextFile UTF8 "instantiations.txt" (show (Array.fromFoldable (Map.keys instantiations)))'
);

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', content);
