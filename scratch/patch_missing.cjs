const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(
  'globalRecordStructs :: Ref.Ref (Set.Set String)\nglobalRecordStructs = unsafePerformEffect (Ref.new Set.empty)',
  'globalReboxPairs :: Ref.Ref (Map.Map String (Set.Set (Tuple GoType GoType)))\nglobalReboxPairs = unsafePerformEffect (Ref.new Map.empty)\n\nglobalRecordStructs :: Ref.Ref (Set.Set String)\nglobalRecordStructs = unsafePerformEffect (Ref.new Set.empty)'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
console.log("Patched missing globalReboxPairs!");
