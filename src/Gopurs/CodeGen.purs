module Gopurs.CodeGen where

import Prelude
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendAccessor(..), Pair(..), Level(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorOrd(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Convert (BackendModule, BackendBindingGroup)
import Data.String as String
import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.CoreFn (Literal(..), Ident(..), Qualified(..), Prop(..))
import Data.Tuple (Tuple(..))
import Data.Array.NonEmpty as NonEmptyArray
import Data.String.Pattern (Pattern(..))

import Gopurs.GoAst (GoFile, GoDecl, GoExpr(..))
import Gopurs.Printer (printGoFile, printGoExpr)

capitalize :: String -> String
capitalize s = 
  let firstChar = String.take 1 s
  in if String.toUpper firstChar == firstChar 
     then s <> "_" 
     else String.toUpper firstChar <> String.drop 1 s

translate :: Array (Array String) -> BackendModule -> String
translate importsArray backendMod = 
  let
    modNameStr = unwrap backendMod.name
    
    decls = Array.concatMap (translateBindingGroup modNameStr) (Array.fromFoldable backendMod.bindings)
    allDecls = decls
    
    dummyText = printGoFile { packageName: "", imports: [], decls: allDecls }
    
    goImports = Array.nub $ (if String.contains (String.Pattern "gopurs_runtime") dummyText then [ "gopurs/output/gopurs_runtime" ] else []) <> 
      Array.mapMaybe (\i -> 
        let modStr = String.joinWith "." i 
            modPkg = String.replaceAll (String.Pattern ".") (String.Replacement "_") modStr
        in if modStr /= modNameStr && modStr /= "Prim" && String.contains (String.Pattern (modPkg <> ".")) dummyText
           then Just ("gopurs/output/" <> modStr)
           else Nothing
      ) (importsArray <> [ ["Unsafe", "Coerce"], ["Partial", "Unsafe"], ["Data", "Function"], ["Data", "Function", "Uncurried"], ["Record", "Unsafe"], ["Type", "Proxy"], ["Data", "Unit"], ["Data", "Eq"], ["Data", "Semiring"], ["Data", "Ring"], ["Data", "EuclideanRing"] ])
    
    goFile = { packageName: String.replaceAll (String.Pattern ".") (String.Replacement "_") modNameStr
             , imports: goImports
             , decls: allDecls 
             }
  in
    printGoFile goFile

translateBindingGroup :: String -> BackendBindingGroup Ident NeutralExpr -> Array GoDecl
translateBindingGroup modNameStr bg =
  Array.mapMaybe (translateBinding modNameStr) bg.bindings

translateBinding :: String -> Tuple Ident NeutralExpr -> Maybe GoDecl
translateBinding modNameStr (Tuple (Ident name) expr) = 
  let 
    safeName = sanitizeName name
    goName = capitalize safeName
  in
    Just { identifier: goName, expression: translateExpr modNameStr expr }

sanitizeName :: String -> String
sanitizeName name = 
  let s1 = String.replaceAll (String.Pattern "'") (String.Replacement "_prime") (String.replaceAll (String.Pattern "$") (String.Replacement "_dollar") name)
  in if s1 == "break" || s1 == "default" || s1 == "func" || s1 == "interface" || s1 == "select" || s1 == "case" || s1 == "defer" || s1 == "go" || s1 == "map" || s1 == "struct" || s1 == "chan" || s1 == "else" || s1 == "goto" || s1 == "package" || s1 == "switch" || s1 == "const" || s1 == "fallthrough" || s1 == "if" || s1 == "range" || s1 == "type" || s1 == "continue" || s1 == "for" || s1 == "import" || s1 == "return" || s1 == "var" then s1 <> "_" else s1

translateExpr :: String -> NeutralExpr -> GoExpr
translateExpr modNameStr (NeutralExpr expr) = case expr of
  Var (Qualified mbMn (Ident i)) -> 
    let safeName = String.replaceAll (String.Pattern "'") (String.Replacement "_prime") (String.replaceAll (String.Pattern "$") (String.Replacement "_dollar") i)
        baseName = capitalize safeName
    in case mbMn of
      Just mn -> 
        let modStr = unwrap mn
            modPkg = String.replaceAll (String.Pattern ".") (String.Replacement "_") modStr
        in if modStr == modNameStr then GoVar baseName else GoSelector (GoVar modPkg) baseName
      Nothing -> GoVar baseName
  Local (Just (Ident i)) _ -> GoVar (String.replaceAll (String.Pattern "'") (String.Replacement "_prime") (String.replaceAll (String.Pattern "$") (String.Replacement "_dollar") i))
  Local Nothing (Level l) -> GoVar ("_unused_" <> show l)
  Lit (LitString s) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString s ]
  Lit (LitInt i) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoInt i ]
  Lit (LitArray xs) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoRaw ("[]gopurs_runtime.Value{" <> String.joinWith ", " (map (\x -> printGoExpr (translateExpr modNameStr x)) xs) <> "}") ]
  Lit (LitRecord props) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Record") [ GoMap (map (\(Prop k v) -> Tuple k (translateExpr modNameStr v)) props) ]
  App f args -> 
    Array.foldl (\acc arg -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ acc, translateExpr modNameStr arg ]) (translateExpr modNameStr f) (NonEmptyArray.toArray args)
  Abs idents body -> 
    Array.foldr (\(Tuple mbIdent (Level l)) inner ->
      let arg = case mbIdent of
            Just (Ident i) -> String.replaceAll (String.Pattern "'") (String.Replacement "_prime") (String.replaceAll (String.Pattern "$") (String.Replacement "_dollar") i)
            Nothing -> "_unused_" <> show l
      in GoFunc arg inner
    ) (translateExpr modNameStr body) (NonEmptyArray.toArray idents)
  Let mbIdent (Level l) binding body -> 
    let name = case mbIdent of
          Just (Ident i) -> sanitizeName i
          Nothing -> "_unused_" <> show l
    in GoIIFE name (translateExpr modNameStr binding) (translateExpr modNameStr body)
  LetRec (Level l) bindings body ->
    let mappedBindings = map (\(Tuple (Ident i) expr_) -> Tuple (sanitizeName i) (translateExpr modNameStr expr_)) (NonEmptyArray.toArray bindings)
    in GoLetRec mappedBindings (translateExpr modNameStr body)
  Accessor obj accessor -> case accessor of
    GetProp prop -> GoRecordAccess (translateExpr modNameStr obj) prop
    GetIndex idx -> GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ translateExpr modNameStr obj, GoInt idx ]
    GetCtorField _ _ _ _ fieldName _ -> GoRecordAccess (translateExpr modNameStr obj) fieldName
  CtorDef _ _ (Ident name) fields ->
    let recordMap = GoCall (GoSelector (GoVar "gopurs_runtime") "Record") [GoMap (Array.cons (Tuple "_tag" (GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [GoString name])) (map (\f -> Tuple f (GoVar (sanitizeName f))) fields))]
    in Array.foldr (\f inner -> GoFunc (sanitizeName f) inner) recordMap fields
  CtorSaturated _ _ _ (Ident name) props ->
    GoCall (GoSelector (GoVar "gopurs_runtime") "Record") [GoMap (Array.cons (Tuple "_tag" (GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [GoString name])) (map (\(Tuple k v) -> Tuple k (translateExpr modNameStr v)) props))]
  Fail msg ->
    GoRaw ("func() gopurs_runtime.Value { panic(" <> printGoExpr (GoString msg) <> ") }()")
  Branch branches def ->
    GoBranch (map (\(Pair cond t) -> Tuple (translateExpr modNameStr cond) (translateExpr modNameStr t)) (NonEmptyArray.toArray branches)) (translateExpr modNameStr def)
  PrimOp op -> case op of
    Op1 OpBooleanNot e -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "==" (GoSelector (translateExpr modNameStr e) "IntVal") (GoInt 0) ]
    Op1 OpIntNegate e -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "-" (GoInt 0) (GoSelector (translateExpr modNameStr e) "IntVal") ]
    Op1 (OpIsTag (Qualified _ (Ident tag))) e -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "==" (GoSelector (GoRecordAccess (translateExpr modNameStr e) "_tag") "StrVal") (GoString tag) ]
    Op1 OpArrayLength e -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "len") [ GoTypeAssertion (GoSelector (translateExpr modNameStr e) "PtrVal") "[]gopurs_runtime.Value" ] ]
    Op2 (OpIntOrd OpEq) e1 e2 -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "==" (GoSelector (translateExpr modNameStr e1) "IntVal") (GoSelector (translateExpr modNameStr e2) "IntVal") ]
    Op2 (OpIntOrd OpNotEq) e1 e2 -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "!=" (GoSelector (translateExpr modNameStr e1) "IntVal") (GoSelector (translateExpr modNameStr e2) "IntVal") ]
    Op2 (OpIntOrd OpLt) e1 e2 -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "<" (GoSelector (translateExpr modNameStr e1) "IntVal") (GoSelector (translateExpr modNameStr e2) "IntVal") ]
    Op2 (OpIntOrd OpLte) e1 e2 -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "<=" (GoSelector (translateExpr modNameStr e1) "IntVal") (GoSelector (translateExpr modNameStr e2) "IntVal") ]
    Op2 (OpIntOrd OpGt) e1 e2 -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp ">" (GoSelector (translateExpr modNameStr e1) "IntVal") (GoSelector (translateExpr modNameStr e2) "IntVal") ]
    Op2 (OpIntOrd OpGte) e1 e2 -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp ">=" (GoSelector (translateExpr modNameStr e1) "IntVal") (GoSelector (translateExpr modNameStr e2) "IntVal") ]
    Op2 OpBooleanAnd e1 e2 -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "&&" (GoBinOp "!=" (GoSelector (translateExpr modNameStr e1) "IntVal") (GoInt 0)) (GoBinOp "!=" (GoSelector (translateExpr modNameStr e2) "IntVal") (GoInt 0)) ]
    Op2 OpBooleanOr e1 e2 -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "||" (GoBinOp "!=" (GoSelector (translateExpr modNameStr e1) "IntVal") (GoInt 0)) (GoBinOp "!=" (GoSelector (translateExpr modNameStr e2) "IntVal") (GoInt 0)) ]
    _ -> GoVar "gopurs_runtime.Value{}"
  _ -> GoVar "gopurs_runtime.Value{}"
