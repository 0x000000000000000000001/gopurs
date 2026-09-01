const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(
  'coerceGoExpr modNameStr expr (TypeStructPointer b1 f1 s1 a1) (TypeStructPointer b2 f2 s2 a2) | b1 == b2 && s1 == s2 && a1 == a2 = expr',
  `coerceGoExpr modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) | b1 == b2 && s1 == s2 && a1 == a2 = expr

coerceGoExpr modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) | b1 == b2 =
  let
    _register = unsafePerformEffect do
      pairsMap <- Ref.read globalReboxPairs
      let pairs = fromMaybe Set.empty (Map.lookup modNameStr pairsMap)
      if Set.member (Tuple srcT destT) pairs then pure unit
      else Ref.modify_ (\\m -> Map.insert modNameStr (Set.insert (Tuple srcT destT) pairs) m) globalReboxPairs
  in
    GoCall (GoVar ("Rebox_" <> modNameStr <> "_" <> hashString s1 <> "_" <> hashString s2)) [ expr ]`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
console.log("Patched coerceGoExpr TypeStructPointer logic!");
