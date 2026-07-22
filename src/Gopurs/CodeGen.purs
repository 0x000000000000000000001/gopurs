module Gopurs.CodeGen where

import Prelude
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Convert (BackendModule, BackendBindingGroup)
import Data.String as String
import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.CoreFn (Literal(..), Ident(..), Qualified(..))
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
    
    decls = Array.concatMap translateBindingGroup (Array.fromFoldable backendMod.bindings)
    allDecls = decls
    
    dummyText = printGoFile { packageName: "", imports: [], decls: allDecls }
    
    goImports = Array.nub $ [ "gopurs/output/gopurs_runtime" ] <> 
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

translateBindingGroup :: BackendBindingGroup Ident NeutralExpr -> Array GoDecl
translateBindingGroup bg =
  Array.mapMaybe translateBinding bg.bindings

translateBinding :: Tuple Ident NeutralExpr -> Maybe GoDecl
translateBinding (Tuple (Ident name) expr) = 
  let 
    safeName = String.replaceAll (String.Pattern "$") (String.Replacement "_") name
    goName = capitalize safeName
  in
    Just { identifier: goName, expression: translateExpr expr }

translateExpr :: NeutralExpr -> GoExpr
translateExpr (NeutralExpr expr) = case expr of
  Var (Qualified mbMn (Ident i)) -> 
    let safeName = String.replaceAll (String.Pattern "$") (String.Replacement "_") i
        baseName = capitalize safeName
    in case mbMn of
      Just mn -> 
        let modPkg = String.replaceAll (String.Pattern ".") (String.Replacement "_") (unwrap mn)
        in GoSelector (GoVar modPkg) baseName
      Nothing -> GoVar baseName
  Local (Just (Ident i)) _ -> GoVar (String.replaceAll (String.Pattern "$") (String.Replacement "_") i)
  Local Nothing _ -> GoVar "_"
  Lit (LitString s) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString s ]
  App f args -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ translateExpr f, translateExpr (NonEmptyArray.head args) ]
  _ -> GoVar "gopurs_runtime.Value{}"
