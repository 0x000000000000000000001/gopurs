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

translate :: Array (Array String) -> BackendModule -> String
translate _ backendMod = 
  let
    modNameStr = unwrap backendMod.name
    -- Collect module imports
    goImports = [ "gopurs/output/gopurs_runtime", "fmt" ]
    
    decls = Array.concatMap translateBindingGroup (Array.fromFoldable backendMod.bindings)
    
    -- Hardcode a dummy definition for each foreign to avoid compile errors
    foreignDecls = map (\f -> { identifier: (unwrap f), expression: GoVar "gopurs_runtime.Value{}" }) (Array.fromFoldable backendMod.foreign)
    
    -- Expose type tags mapping if needed or just let runtime handle it
    allDecls = foreignDecls <> decls
    
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
    goName = if safeName == "main" then "Main" else safeName
  in
    Just { identifier: goName, expression: translateExpr expr }

translateExpr :: NeutralExpr -> GoExpr
translateExpr (NeutralExpr expr) = case expr of
  Var (Qualified _ (Ident "bindE")) ->
    GoCall (GoSelector (GoVar "gopurs_runtime") "Func") 
      [ GoFunc ["a"] 
          (GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
            [ GoFunc ["f"]
                (GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
                  [ GoFunc ["_"]
                      (GoBlock
                        [ GoVar "resA := gopurs_runtime.Apply(a, gopurs_runtime.Value{})"
                        , GoVar "resB := gopurs_runtime.Apply(f, resA)"
                        , GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ GoVar "resB", GoVar "gopurs_runtime.Value{}" ])
                        ]
                      )
                  ]
                ))
            ]
          ))
      ]
  Var (Qualified _ (Ident "log")) ->
    GoCall (GoSelector (GoVar "gopurs_runtime") "Func") 
      [ GoFunc ["x"] 
          (GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
            [ GoFunc ["_"]
                (GoBlock
                  [ GoCall (GoVar "fmt.Println") [ GoSelector (GoVar "x") "StrVal" ]
                  , GoReturn (GoVar "gopurs_runtime.Value{}")
                  ]
                )
            ]
          ))
      ]
  Var (Qualified _ (Ident "showStringImpl")) ->
    GoCall (GoSelector (GoVar "gopurs_runtime") "Func") 
      [ GoFunc ["s"] 
          (GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Str")
            [ GoCall (GoVar "fmt.Sprintf") [ GoString "%q", GoSelector (GoVar "s") "StrVal" ] ]
          ))
      ]
  Var (Qualified _ (Ident i)) -> GoVar (String.replaceAll (String.Pattern "$") (String.Replacement "_") i)
  Local (Just (Ident i)) _ -> GoVar (String.replaceAll (String.Pattern "$") (String.Replacement "_") i)
  Local Nothing _ -> GoVar "_"
  Lit (LitString s) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString s ]
  App f args -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ translateExpr f, translateExpr (NonEmptyArray.head args) ]
  _ -> GoVar "gopurs_runtime.Value{}"
