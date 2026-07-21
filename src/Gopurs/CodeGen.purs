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

import Gopurs.GoAst (GoFile, GoDecl, GoExpr(..))
import Gopurs.Printer (printGoFile)

capitalize :: String -> String
capitalize s = String.toUpper (String.take 1 s) <> String.drop 1 s

translate :: Array (Array String) -> BackendModule -> String
translate importsArray backendMod = 
  let
    modNameStr = unwrap backendMod.name
    
    goImports = [ "gopurs/output/gopurs_runtime" ] <> map (\i -> "gopurs/output/" <> String.joinWith "." i) importsArray
    
    decls = Array.concatMap translateBindingGroup (Array.fromFoldable backendMod.bindings)
    
    allDecls = decls
    
    goFile = { packageName: String.replaceAll (String.Pattern ".") (String.Replacement "_") modNameStr
             , imports: Array.nub goImports
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
    goName = if safeName == "main" then "Main" else safeName
  in
    Just { identifier: goName, expression: translateExpr expr }

translateExpr :: NeutralExpr -> GoExpr
translateExpr (NeutralExpr expr) = case expr of
  Var (Qualified _ (Ident i)) -> GoVar (String.replaceAll (String.Pattern "$") (String.Replacement "_") i)
  Local (Just (Ident i)) _ -> GoVar (String.replaceAll (String.Pattern "$") (String.Replacement "_") i)
  Local Nothing _ -> GoVar "_"
  Lit (LitString s) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString s ]
  App f args -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ translateExpr f, translateExpr (NonEmptyArray.head args) ]
  _ -> GoVar "gopurs_runtime.Value{}"
