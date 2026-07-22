module Gopurs.CodeGen where

import Prelude
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendAccessor(..), Pair(..))
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
import Gopurs.Printer (printGoFile)

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
      ) importsArray
    
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
    safeName = String.replaceAll (String.Pattern "'") (String.Replacement "_prime") (String.replaceAll (String.Pattern "$") (String.Replacement "_") name)
    goName = capitalize safeName
  in
    Just { identifier: goName, expression: translateExpr modNameStr expr }

translateExpr :: String -> NeutralExpr -> GoExpr
translateExpr modNameStr (NeutralExpr expr) = case expr of
  Var (Qualified mbMn (Ident i)) -> 
    let safeName = String.replaceAll (String.Pattern "'") (String.Replacement "_prime") (String.replaceAll (String.Pattern "$") (String.Replacement "_") i)
        baseName = capitalize safeName
    in case mbMn of
      Just mn -> 
        let modStr = unwrap mn
            modPkg = String.replaceAll (String.Pattern ".") (String.Replacement "_") modStr
        in if modStr == modNameStr then GoVar baseName else GoSelector (GoVar modPkg) baseName
      Nothing -> GoVar baseName
  Local (Just (Ident i)) _ -> GoVar (String.replaceAll (String.Pattern "'") (String.Replacement "_prime") (String.replaceAll (String.Pattern "$") (String.Replacement "_") i))
  Local Nothing _ -> GoVar "_"
  Lit (LitString s) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString s ]
  Lit (LitInt i) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoInt i ]
  Lit (LitRecord props) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Record") [ GoMap (map (\(Prop k v) -> Tuple k (translateExpr modNameStr v)) props) ]
  App f args -> 
    Array.foldl (\acc arg -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ acc, translateExpr modNameStr arg ]) (translateExpr modNameStr f) (NonEmptyArray.toArray args)
  Abs idents body -> 
    Array.foldr (\(Tuple mbIdent _) inner ->
      let arg = case mbIdent of
            Just (Ident i) -> String.replaceAll (String.Pattern "'") (String.Replacement "_prime") (String.replaceAll (String.Pattern "$") (String.Replacement "_") i)
            Nothing -> "_"
      in GoFunc arg inner
    ) (translateExpr modNameStr body) (NonEmptyArray.toArray idents)
  Let mbIdent _ binding body -> 
    let name = case mbIdent of
          Just (Ident i) -> String.replaceAll (String.Pattern "'") (String.Replacement "_prime") (String.replaceAll (String.Pattern "$") (String.Replacement "_") i)
          Nothing -> "_"
    in GoIIFE name (translateExpr modNameStr binding) (translateExpr modNameStr body)
  Accessor obj (GetProp prop) -> GoRecordAccess (translateExpr modNameStr obj) prop
  Branch branches def ->
    GoBranch (map (\(Pair cond t) -> Tuple (translateExpr modNameStr cond) (translateExpr modNameStr t)) (NonEmptyArray.toArray branches)) (translateExpr modNameStr def)
  _ -> GoVar "gopurs_runtime.Value{}"
