const fs = require('fs');
const file = 'src/Gopurs/CodeGen.purs';
let content = fs.readFileSync(file, 'utf8');
content = content.replace(
    'pTypes = paramTypes fn.body\n',
    'pTypes = Debug.trace ("pTypes for " <> fn.ident <> ": " <> String.joinWith ", " (map (\\(Tuple k v) -> k <> " -> " <> goTypeToStr (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts modNameStr v)) (Map.toUnfoldable (paramTypes fn.body) :: Array (Tuple String GoType)))) \\_ -> paramTypes fn.body\n'
);
fs.writeFileSync(file, content);
