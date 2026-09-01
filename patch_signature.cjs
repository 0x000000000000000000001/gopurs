const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const targetStr = `translateExprImpl :: Ref { decls :: Array GoDecl, rawDecls :: Array String, elidedCtors :: Set.Set String, ctorTypes :: Map String { vars :: Array String, fields :: Array ExprType }, pointerAdtPaths :: Map String { ctorName :: String, arity :: Int }, pointerAdtNodes :: Set String, pointerAdtLeaves :: Map String { nodeBaseStruct :: String, nodeCtor :: String }, enumAdts :: Set.Set String, enumCtors :: Set.Set String, globalTypes :: Map.Map String ExprType, classDeclsFields :: Map String { vars :: Array String, fields :: Array { name :: String, "type" :: ExprType } }, globalId :: Int } -> Int -> String -> Array String -> Map String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int } -> Map String { name :: String, goType :: GoType } -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }`;

const replacementStr = `translateExprImpl :: Ref { decls :: Array GoDecl, rawDecls :: Array String, elidedCtors :: Set.Set String, ctorTypes :: Map String { vars :: Array String, fields :: Array ExprType }, pointerAdtPaths :: Map String { ctorName :: String, arity :: Int }, pointerAdtNodes :: Set String, pointerAdtLeaves :: Map String { nodeBaseStruct :: String, nodeCtor :: String }, enumAdts :: Set.Set String, enumCtors :: Set.Set String, globalTypes :: Map.Map String ExprType, classDeclsFields :: Map String { vars :: Array String, fields :: Array { name :: String, "type" :: ExprType } }, globalId :: Int } -> Int -> String -> Array String -> Map String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int } -> Map String { name :: String, goType :: GoType } -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType, fRet :: GoType } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }`;

if (!code.includes(targetStr)) {
  console.error("Target string not found!");
  process.exit(1);
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.replace(targetStr, replacementStr));
console.log("Patched CodeGen.purs successfully.");
