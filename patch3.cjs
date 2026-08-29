const fs = require('fs');
const content = fs.readFileSync('src/Main.purs', 'utf8');

let newContent = content.replace(
  'buildModules\n  :: BuildOptions\n  -> List (Module Ann)\n  -> Aff Unit\nbuildModules options modules = do',
  `import Effect.Ref as Ref\n\nbuildModules\n  :: BuildOptions\n  -> Ref.Ref (List (Module Ann))\n  -> Aff Unit\nbuildModules options modulesRef = do`
);

newContent = newContent.replace(
  'go env remainingModules = case remainingModules of',
  `go env = do
      remainingModules <- liftEffect $ Ref.read modulesRef
      case remainingModules of`
);

newContent = newContent.replace(
  '    go\n      { directives: newDirectives\n      , implementations: Map.empty\n      , moduleIndex: moduleIndex + 1\n      , exports: newExports\n      }\n      remaining',
  `    liftEffect $ Ref.write remaining modulesRef\n    go\n      { directives: newDirectives\n      , implementations: Map.empty\n      , moduleIndex: moduleIndex + 1\n      , exports: newExports\n      }`
);

newContent = newContent.replace(
  '  go { directives: options.directives, implementations: Map.empty, moduleIndex: 1, exports: Map.empty } modules',
  '  go { directives: options.directives, implementations: Map.empty, moduleIndex: 1, exports: Map.empty }'
);

newContent = newContent.replace(
  '    (List.fromFoldable prepared.monomorphizedModules)',
  '    (unsafePerformEffect (Ref.new (List.fromFoldable prepared.monomorphizedModules)))'
);

fs.writeFileSync('src/Main.purs', newContent);
