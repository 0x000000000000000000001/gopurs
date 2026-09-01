const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(/_ = unsafePerformEffect do\n      h <- Ref\.read helpersRef\n      if Set\.member \(Tuple srcT destT\) h\.reboxPairs then pure unit\n      else Ref\.modify_ \(\\r -> r { reboxPairs = Set\.insert \(Tuple srcT destT\) r\.reboxPairs }\) helpersRef/, `_ = unsafePerformEffect do\n      pairs <- Ref.read globalReboxPairs\n      if Set.member (Tuple srcT destT) pairs then pure unit\n      else Ref.modify_ (\\s -> Set.insert (Tuple srcT destT) s) globalReboxPairs`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
