const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const targetStr = `coerceGoExpr modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) | b1 == b2 =`;
const replacementStr = `coerceGoExpr modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) | b1 == b2 =
  let
    _register = unsafePerformEffect do
      pairsMap <- Ref.read globalReboxPairs
      let pairs = fromMaybe Set.empty (Map.lookup modNameStr pairsMap)
      if Set.member (Tuple srcT destT) pairs then pure unit
      else Ref.modify_ (\\m -> Map.insert modNameStr (Set.insert (Tuple srcT destT) pairs) m) globalReboxPairs
  in
    GoCall (GoVar ("Rebox_" <> modNameStr <> "_" <> hashString s1 <> "_" <> hashString s2)) [ expr ]

coerceGoExpr modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) = Debug.trace ("MISMATCH B1 B2: " <> b1 <> " vs " <> b2 <> " from " <> goTypeToStr srcT <> " to " <> goTypeToStr destT) \\_ ->
  unboxGoExpr (boxGoExpr expr srcT) TypeValue destT

coerceGoExpr modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) | b1 == b2 =`;

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.replace(targetStr, replacementStr));
console.log("Patched CodeGen.purs successfully.");
