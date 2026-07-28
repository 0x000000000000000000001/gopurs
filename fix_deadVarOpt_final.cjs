const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  '            ) (Map.toUnfoldable bound :: Array (Tuple String { name :: String, goType :: GoType })))\n            \n          deadVarOpt = Nothing\n          _ = unsafePerformEffect (case deadVarOpt of\n                Just n -> Ref.modify_ (Set.insert n) globalReusedVars\n                Nothing -> pure unit)\n          result =',
  '            ) (Array.fromFoldable (map (\\v -> Tuple v.name v) (Map.values bound))))\n            \n          deadVarOpt = deadVarOptRaw\n          _ = unsafePerformEffect (case deadVarOpt of\n                Just n -> Ref.modify_ (Set.insert n) globalReusedVars\n                Nothing -> pure unit)\n          result ='
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
