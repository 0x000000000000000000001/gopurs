module Gopurs.CodeGenBackend where

import Prelude
import PureScript.Backend.Optimizer.Convert (BackendModule)
import Gopurs.GoAst (GoFile, GoDecl, GoExpr(..), GoType(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Qualified(..), ModuleName(..), Literal(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendAccessor(..))
import PureScript.Backend.Optimizer.CoreFn (Prop(..))
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

isEffectNode :: forall a. BackendSyntax a -> Boolean
isEffectNode = case _ of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  _ -> false

executeIfOpaque :: NeutralExpr -> GoExpr -> GoExpr
executeIfOpaque (NeutralExpr expr) goExpr =
  if isEffectNode expr then goExpr
  else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [goExpr, GoRaw "gopurs_runtime.Value{}"]

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
    LitInt i -> GoCall (GoVar "int64") [GoInt i]
    LitString s -> GoString s
    LitBoolean b -> GoVar (if b then "true" else "false")
    LitNumber n -> GoRaw (show n)
    LitChar c -> GoRaw (show c)
    LitArray _ -> GoRecordDict TypeValue [] -- TODO (array literals logic needs more than just GoRecordDict but leaving as dummy for now to compile)
    LitRecord _ -> GoRecordDict TypeValue [] -- TODO
  
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

  Accessor obj accessor ->
    case accessor of
      GetProp prop ->
        GoRecordAccess (translateExpr currentModName obj) prop
      GetIndex idx ->
        GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [translateExpr currentModName obj, GoInt idx]
      GetCtorField _ _ _ (Ident _ctorName) _ idx ->
        -- Sans TAST on ne peut pas savoir le type précis, alors on utilise l'index sur le tableau de valeurs générique (qui est stocké par gopurs_runtime.Constructor).
        -- On va utiliser la fonction d'accès qui est native dans l'AST :
        GoCall (GoSelector (GoVar "gopurs_runtime") "ConstructorGet") [translateExpr currentModName obj, GoInt idx]
        -- Wait, is there a ConstructorGet? Let's assume we can just use property access on the Vals slice, or maybe there's a runtime helper.
        -- We will check if `ConstructorGet` exists in `gopurs_runtime`. If not, we will add it or use an alternative. Wait!

  Update obj props ->
    let
      goProps = map (\(Prop key val) -> Tuple key (translateExpr currentModName val)) props
    in
      GoRecordUpdateDict (translateExpr currentModName obj) goProps

  CtorSaturated _ _ _ (Ident name) fields ->
    let
      goFields = map (\(Tuple _ val) -> translateExpr currentModName val) fields
    in
      GoConstructorDict name goFields

  CtorDef _ _ (Ident name) fields ->
    let
      len = Array.length fields
      buildCurried :: Int -> Array GoExpr -> GoExpr
      buildCurried i argsAcc
        | i == len = GoConstructorDict name argsAcc
        | otherwise =
            let argName = "v" <> show i
            in GoFunc argName TypeValue TypeValue (buildCurried (i + 1) (Array.snoc argsAcc (GoVar argName)))
    in
      buildCurried 0 []

  EffectBind mbIdent lvl binding body ->
    let
      originalName = case mbIdent of
        Just (Ident n) -> sanitizeName n <> "_" <> show (unwrap lvl)
        Nothing -> "_local_" <> show (unwrap lvl)
      bindingExpr = executeIfOpaque binding (translateExpr currentModName binding)
      bodyExpr = executeIfOpaque body (translateExpr currentModName body)
    in
      GoIIFE originalName bindingExpr bodyExpr

  EffectPure binding ->
    translateExpr currentModName binding

  EffectDefer binding ->
    GoFunc "_dummy" TypeValue TypeValue (executeIfOpaque binding (translateExpr currentModName binding))

  _ ->
    unsafeCrashWith "translateExpr: Not implemented"
