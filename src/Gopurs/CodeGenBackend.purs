module Gopurs.CodeGenBackend where

import Prelude
import PureScript.Backend.Optimizer.Convert (BackendModule)
import Gopurs.GoAst (GoFile, GoDecl, GoExpr(..), GoType(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Qualified(..), ModuleName(..), Literal(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))
import PureScript.Backend.Optimizer.FreeVars (sanitizeName)
import Data.Maybe (Maybe(..))
import Partial.Unsafe (unsafeCrashWith)
import Data.String (replaceAll, Pattern(..), Replacement(..))
import Data.Newtype (unwrap)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Tuple (Tuple(..))
import Data.Foldable (foldl, foldr)

translateBackend :: BackendModule -> GoFile
translateBackend mod =
  let
    modNameStr = replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
    
    decls = Array.concatMap (\group ->
      map (\(Tuple (Ident name) val) ->
        { identifier: name
        , expression: translateExpr modNameStr val
        , goType: TypeValue
        }
      ) group.bindings
    ) mod.bindings

  in
    { packageName: "gopurs_module"
    , imports: [ "fmt", "gopurs/gopurs_runtime" ]
    , decls: decls
    , rawDecls: []
    , foreigns: []
    }

translateExpr :: String -> NeutralExpr -> GoExpr
translateExpr currentModName (NeutralExpr expr) = case expr of
  Var (Qualified mbMod (Ident name)) ->
    let
      pkgStr = case mbMod of
        Just (ModuleName mn) -> replaceAll (Pattern ".") (Replacement "_") mn
        Nothing -> currentModName
    in
      GoVar ("gopurs_module_" <> pkgStr <> "." <> sanitizeName name)
  
  Local mbIdent level ->
    case mbIdent of
      Just (Ident name) -> GoVar (sanitizeName name <> "_" <> show (unwrap level))
      Nothing -> GoVar ("_local_" <> show (unwrap level))

  Lit lit -> case lit of
    LitInt i -> GoInt i
    LitString s -> GoString s
    LitBoolean b -> GoVar (if b then "true" else "false")
    LitNumber n -> GoRaw (show n)
    LitChar c -> GoRaw (show c)
    LitArray _ -> GoRecordDict [] -- TODO (array literals logic needs more than just GoRecordDict but leaving as dummy for now to compile)
    LitRecord _ -> GoRecordDict [] -- TODO
  
  App f xs ->
    foldl (\acc x -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [acc, translateExpr currentModName x]) (translateExpr currentModName f) xs

  Abs args body ->
    foldr (\(Tuple mbIdent level) acc ->
      let
        argName = case mbIdent of
          Just (Ident name) -> sanitizeName name <> "_" <> show (unwrap level)
          Nothing -> "_local_" <> show (unwrap level)
      in
        GoFunc argName TypeValue TypeValue acc
    ) (translateExpr currentModName body) args

  Let mbIdent level val body ->
    let
      name = case mbIdent of
        Just (Ident n) -> sanitizeName n <> "_" <> show (unwrap level)
        Nothing -> "_local_" <> show (unwrap level)
    in
      GoIIFE name (translateExpr currentModName val) (translateExpr currentModName body)

  LetRec level binds body ->
    let
      goBinds = map (\(Tuple (Ident name) val) -> Tuple (sanitizeName name <> "_" <> show (unwrap level)) (translateExpr currentModName val)) (NEA.toArray binds)
    in
      GoLetRec goBinds (translateExpr currentModName body)

  _ ->
    unsafeCrashWith "translateExpr: Not implemented"
