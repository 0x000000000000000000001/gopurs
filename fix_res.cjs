const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  '          result =\n            let\n              monoStructName = "Constructor_" <> sanitizeName name',
  '          _ = unsafePerformEffect (case deadVarOptRaw of\n                Just n -> Debug.trace ("CtorSaturated " <> name <> " REUSING: " <> n <> " | liveOut=" <> show (Set.toUnfoldable liveOut :: Array String) <> " | bound=" <> String.joinWith "," (map _.name (Map.values bound)) <> " | fv=" <> show (Set.toUnfoldable (freeVars tcoExpr) :: Array String)) \\_ -> pure unit\n                Nothing -> pure unit)\n          deadVarOpt = deadVarOptRaw\n          result =\n            let\n              monoStructName = "Constructor_" <> sanitizeName name'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
