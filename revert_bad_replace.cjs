const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  'exprTypeToGenericGoType ptrPaths typeVars modNameStr (Record fields) = \n  let _ = unsafePerformEffect (Ref.write Set.empty globalReusedVars)\n  in TypeRecord',
  'exprTypeToGenericGoType ptrPaths typeVars modNameStr (Record fields) = TypeRecord'
);

code = code.replace(
  'translateDecl helpers modNameStr bound decl = case decl of',
  'translateDecl helpers modNameStr bound decl =\n  let _ = unsafePerformEffect (Ref.write Set.empty globalReusedVars)\n  in case decl of'
);

code = code.replace(
  'deadVarOptRaw = Array.head (Array.mapMaybe (\\(Tuple _ v) -> \n              if v.goType == expectedGoType \n                 && not (Set.member v.name liveOut) \n                 && not (Set.member v.name (freeVars tcoExpr)) \n              then Just v.name \n              else Nothing\n            ) (Map.toUnfoldable bound :: Array (Tuple String { name :: String, goType :: GoType })))\n            \n          result =',
  'deadVarOptRaw = Array.head (Array.mapMaybe (\\(Tuple _ v) -> \n              let reused = unsafePerformEffect (Ref.read globalReusedVars)\n              in\n              if v.goType == expectedGoType \n                 && not (Set.member v.name reused)\n                 && not (Set.member v.name liveOut) \n                 && not (Set.member v.name (freeVars tcoExpr)) \n              then Just v.name \n              else Nothing\n            ) (Array.fromFoldable (map (\\v -> Tuple v.name v) (Map.values bound))))\n            \n          deadVarOpt = deadVarOptRaw\n          _ = unsafePerformEffect (case deadVarOpt of\n                Just n -> Ref.modify_ (Set.insert n) globalReusedVars\n                Nothing -> pure unit)\n          result ='
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
