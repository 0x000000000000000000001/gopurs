import fs from 'fs';
const file = '/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/Convert.purs';
let content = fs.readFileSync(file, 'utf8');

const oldCode = `getReturnType :: ExprType -> Maybe ExprType
getReturnType (ForAll _ t) = getReturnType t
getReturnType (ConstrainedType _ t) = getReturnType t
getReturnType (Func _ ret) = Just ret
getReturnType _ = Nothing`;

const newCode = `getReturnType :: ExprType -> Maybe ExprType
getReturnType (ForAll _ t) = getReturnType t
getReturnType (ConstrainedType _ t) = getReturnType t
getReturnType (Func args ret) = case Array.uncons args of
  Just { tail } | Array.length tail > 0 -> Just (Func tail ret)
  _ -> Just ret
getReturnType _ = Nothing`;

if (content.includes(oldCode)) {
  content = content.replace(oldCode, newCode);
  fs.writeFileSync(file, content);
  console.log("Patched getReturnType successfully!");
} else {
  console.error("Could not find old code block.");
}
