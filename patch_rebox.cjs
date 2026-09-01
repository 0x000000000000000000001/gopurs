const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

// 1. Change globalReboxPairs definition
code = code.replace(
  'globalReboxPairs :: Ref.Ref (Set.Set (Tuple GoType GoType))\nglobalReboxPairs = unsafePerformEffect (Ref.new Set.empty)',
  'globalReboxPairs :: Ref.Ref (Map.Map String (Set.Set (Tuple GoType GoType)))\nglobalReboxPairs = unsafePerformEffect (Ref.new Map.empty)'
);

// 2. Change coerceGoExpr to use Map
code = code.replace(
  'pairs <- Ref.read globalReboxPairs\n      if Set.member (Tuple srcT destT) pairs then pure unit\n      else Ref.modify_ (\\s -> Set.insert (Tuple srcT destT) s) globalReboxPairs',
  `pairsMap <- Ref.read globalReboxPairs\n      let pairs = fromMaybe Set.empty (Map.lookup modNameStr pairsMap)\n      if Set.member (Tuple srcT destT) pairs then pure unit\n      else Ref.modify_ (\\m -> Map.insert modNameStr (Set.insert (Tuple srcT destT) pairs) m) globalReboxPairs`
);

// 3. Append to helpers.rawDecls
code = code.replace(
  ', rawDecls: helpers.rawDecls',
  `, rawDecls: helpers.rawDecls <> (unsafePerformEffect do
          pairsMap <- Ref.read globalReboxPairs
          let reboxPairs = fromMaybe Set.empty (Map.lookup modNameStr pairsMap)
          pure $ Array.fromFoldable $ map (\\(Tuple srcT destT) -> 
              case srcT, destT of
                TypeStructPointer b1 f1 s1 a1, TypeStructPointer b2 f2 s2 a2 ->
                  let
                    funcName = "Rebox_" <> modNameStr <> "_" <> hashString s1 <> "_" <> hashString s2
                    assignments = String.joinWith "\\n" (Array.mapWithIndex (\\i (Tuple t1 t2) -> 
                        "\\t\\tout.V" <> show i <> " = " <> printGoExpr (unboxGoExpr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)
                      ) (Array.zip a1 a2))
                  in
                    "func " <> funcName <> "(in *" <> f1 <> "[" <> String.joinWith ", " (map goTypeToStr a1) <> "]) *" <> f2 <> "[" <> String.joinWith ", " (map goTypeToStr a2) <> "] {\\n\\tif in == nil { return nil }\\n\\tout := &" <> f2 <> "[" <> String.joinWith ", " (map goTypeToStr a2) <> "]{}\\n" <> assignments <> "\\n\\treturn out\\n}"
                _, _ -> ""
            ) (Array.fromFoldable reboxPairs)
        )`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
