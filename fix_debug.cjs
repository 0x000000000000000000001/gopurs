const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  'deadVarOptRaw = Array.head (Array.mapMaybe (\\(Tuple _ v) -> if v.goType == expectedGoType && not (Set.member v.name liveOut) && not (Set.member v.name (freeVars tcoExpr)) then Just v.name else Nothing) (Map.toUnfoldable bound :: Array (Tuple String { name :: String, goType :: GoType })))\n          _ = unsafePerformEffect (Debug.trace ("CtorSaturated " <> name <> " liveOut=" <> show (Set.toUnfoldable liveOut :: Array String) <> " freeVars=" <> show (Set.toUnfoldable (freeVars tcoExpr) :: Array String) <> " picked=" <> show deadVarOptRaw) \\_ -> pure unit)\n          deadVarOpt = deadVarOptRaw',
  'deadVarOpt = Array.head'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
