const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(/globalRecordStructs :: Ref\.Ref \(Set\.Set String\)\nglobalReboxPairs :: Ref\.Ref \(Set\.Set \(Tuple GoType GoType\)\)\nglobalReboxPairs = unsafePerformEffect \(Ref\.new Set\.empty\)\nglobalRecordStructs = unsafePerformEffect \(Ref\.new Set\.empty\)/, 'globalReboxPairs :: Ref.Ref (Set.Set (Tuple GoType GoType))\nglobalReboxPairs = unsafePerformEffect (Ref.new Set.empty)\nglobalRecordStructs :: Ref.Ref (Set.Set String)\nglobalRecordStructs = unsafePerformEffect (Ref.new Set.empty)');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
