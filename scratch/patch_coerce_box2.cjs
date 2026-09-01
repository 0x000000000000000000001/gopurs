const fs = require('fs');

const path = 'src/Gopurs/CodeGen.purs';
let content = fs.readFileSync(path, 'utf8');

const target = `coerceGoExpr modNameStr expr srcT@(TypeStructPointer b f s a) TypeValue | Array.any (_ /= TypeValue) a =
  let
    destT = TypeStructPointer b f (b <> if Array.length a > 0 then "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") a) <> "]" else "") (map (const TypeValue) a)
  in
    boxGoExpr (coerceGoExpr modNameStr expr srcT destT) destT

coerceGoExpr modNameStr expr TypeValue destT@(TypeStructPointer b f s a) | Array.any (_ /= TypeValue) a =
  let
    srcT = TypeStructPointer b f (b <> if Array.length a > 0 then "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") a) <> "]" else "") (map (const TypeValue) a)
  in
    coerceGoExpr modNameStr (unboxGoExpr expr TypeValue srcT) srcT destT`;

const replacement = `coerceGoExpr modNameStr expr srcT@(TypeStructPointer b f s a) TypeValue | Array.any (_ /= TypeValue) a =
  let
    basePath = case String.indexOf (Pattern "[") s of
      Just i -> String.take i s
      Nothing -> s
    destT = TypeStructPointer b f (basePath <> if Array.length a > 0 then "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") a) <> "]" else "") (map (const TypeValue) a)
  in
    boxGoExpr (coerceGoExpr modNameStr expr srcT destT) destT

coerceGoExpr modNameStr expr TypeValue destT@(TypeStructPointer b f s a) | Array.any (_ /= TypeValue) a =
  let
    basePath = case String.indexOf (Pattern "[") s of
      Just i -> String.take i s
      Nothing -> s
    srcT = TypeStructPointer b f (basePath <> if Array.length a > 0 then "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") a) <> "]" else "") (map (const TypeValue) a)
  in
    coerceGoExpr modNameStr (unboxGoExpr expr TypeValue srcT) srcT destT`;

if (content.includes(target)) {
  content = content.replace(target, replacement);
  fs.writeFileSync(path, content, 'utf8');
  console.log("Patched CodeGen.purs successfully");
} else {
  console.log("Target not found!");
}
