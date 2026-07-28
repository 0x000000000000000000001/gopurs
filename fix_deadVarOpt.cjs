const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  '             ) (Map.toUnfoldable bound :: Array (Tuple String { name :: String, goType :: GoType })))\n            \n          result =',
  '             ) (Map.toUnfoldable bound :: Array (Tuple String { name :: String, goType :: GoType })))\n          _ = unsafePerformEffect (case deadVarOptRaw of\n                Just n -> Debug.trace ("CtorSaturated " <> name <> " REUSING: " <> n <> " | liveOut=" <> show (Set.toUnfoldable liveOut :: Array String) <> " | bound=" <> String.joinWith "," (map _.name (Map.values bound)) <> " | fv=" <> show (Set.toUnfoldable (freeVars tcoExpr) :: Array String)) \\_ -> pure unit\n                Nothing -> pure unit)\n          deadVarOpt = deadVarOptRaw\n          result ='
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
