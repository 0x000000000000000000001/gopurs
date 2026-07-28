const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// Add globalReusedVars
if (!code.includes('globalReusedVars ::')) {
  code = code.replace(
    'translate ::',
    'globalReusedVars :: Ref.Ref (Set.Set String)\nglobalReusedVars = unsafePerformEffect (Ref.new Set.empty)\n\ntranslate ::'
  );
}

// Reset in translateDecl
code = code.replace(
  'translateDecl helpers modNameStr bound decl = case decl of',
  'translateDecl helpers modNameStr bound decl =\n  let _ = unsafePerformEffect (Ref.write Set.empty globalReusedVars)\n  in case decl of'
);

// Update deadVarOptRaw
code = code.replace(
  'deadVarOptRaw = Array.head (Array.mapMaybe (\\(Tuple _ v) ->',
  'deadVarOptRaw = Array.head (Array.mapMaybe (\\(Tuple _ v) -> \n              let reused = unsafePerformEffect (Ref.read globalReusedVars)\n              in'
);

code = code.replace(
  'if v.goType == expectedGoType',
  'if v.goType == expectedGoType \\n                 && not (Set.member v.name reused)'
);

// Remove the old buggy fix_deadVarOpt2.cjs changes if they exist
code = code.replace(
  ', reusedVars: Set.insert deadVar accProps.reusedVars }',
  ' }'
);

// Add the Ref.modify_ side effect
code = code.replace(
  'deadVarOpt = deadVarOptRaw\n          result =',
  'deadVarOpt = deadVarOptRaw\n          _ = unsafePerformEffect (case deadVarOpt of\n                Just n -> Ref.modify_ (Set.insert n) globalReusedVars\n                Nothing -> pure unit)\n          result ='
);

code = code.replace(/\\n/g, '\n');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
