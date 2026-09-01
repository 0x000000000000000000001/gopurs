const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

// 1. Fix the matchB1 split
code = code.replace(/String\.split \(Pattern "_"\) k/g, 'String.split (Pattern ".") k');

// 2. Remove the OLD rebox generator
// It looks like:
// ) <> (unsafePerformEffect do\n          pairsMap <- Ref.read globalReboxPairs\n          let reboxPairs = fromMaybe Set.empty (Map.lookup modNameStr pairsMap)\n          pure $ Array.fromFoldable $ map (\\(Tuple srcT destT) -> 
const oldReboxRegex = /\) <> \(unsafePerformEffect do\s*pairsMap <- Ref\.read globalReboxPairs\s*let reboxPairs = fromMaybe Set\.empty \(Map\.lookup modNameStr pairsMap\)\s*pure \$ Array\.fromFoldable \$ map \(\\\(Tuple srcT destT\) ->[\s\S]*?\) \(Array\.fromFoldable reboxPairs\)\s*\)/;

code = code.replace(oldReboxRegex, ')');

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
console.log("Patched loop and removed duplicate Rebox generator!");
