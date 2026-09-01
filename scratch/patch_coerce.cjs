const fs = require('fs');
const path = 'src/Gopurs/CodeGen.purs';
let content = fs.readFileSync(path, 'utf8');

const target = `coerceGoExpr :: GoExpr -> GoType -> GoType -> GoExpr
coerceGoExpr expr from to | from == to = expr
coerceGoExpr expr (TypeStructPointer b1 f1 s1 a1) (TypeStructPointer b2 f2 s2 a2) | b1 == b2 && s1 == s2 && a1 == a2 = expr
coerceGoExpr expr from TypeValue = boxGoExpr expr from
coerceGoExpr expr TypeValue to = unboxGoExpr expr TypeValue to
coerceGoExpr expr from to = unboxGoExpr (boxGoExpr expr from) TypeValue to`;

const replacement = `coerceGoExpr :: GoExpr -> GoType -> GoType -> GoExpr
coerceGoExpr expr from to | from == to = expr
coerceGoExpr expr (TypeStructPointer b1 f1 s1 a1) (TypeStructPointer b2 f2 s2 a2) | b1 == b2 && s1 == s2 && a1 == a2 = expr

coerceGoExpr expr srcT@(TypeStructPointer b f s a) TypeValue | Array.any (_ /= TypeValue) a =
  let
    basePath = case String.indexOf (Pattern "[") s of
      Just i -> String.take i s
      Nothing -> s
    destT = TypeStructPointer b f (basePath <> if Array.length a > 0 then "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") a) <> "]" else "") (map (const TypeValue) a)
  in
    boxGoExpr (coerceGoExpr expr srcT destT) destT

coerceGoExpr expr TypeValue destT@(TypeStructPointer b f s a) | Array.any (_ /= TypeValue) a =
  let
    basePath = case String.indexOf (Pattern "[") s of
      Just i -> String.take i s
      Nothing -> s
    srcT = TypeStructPointer b f (basePath <> if Array.length a > 0 then "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") a) <> "]" else "") (map (const TypeValue) a)
  in
    coerceGoExpr (unboxGoExpr expr TypeValue srcT) srcT destT

coerceGoExpr expr from TypeValue = boxGoExpr expr from
coerceGoExpr expr TypeValue to = unboxGoExpr expr TypeValue to
coerceGoExpr expr from to = unboxGoExpr (boxGoExpr expr from) TypeValue to`;

if (content.includes(target)) {
  content = content.replace(target, replacement);
  fs.writeFileSync(path, content, 'utf8');
  console.log("Patched coerceGoExpr successfully!");
} else {
  console.log("Target not found!");
}
