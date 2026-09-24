module Gopurs.BorrowedObjects (borrowReadOnlyObjects) where

import Prelude

import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.String as String
import Data.String.CodeUnits as CU
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (CodegenMetadata)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendSyntax(..), Level)

-- Elide an identity-object copy only when every use of the Either and its
-- Right payload is accounted for. The object may feed saturated public field
-- readers; closures, effects, recursion and arbitrary consumers are
-- rejected. Custom element decoders never match the producer. The native
-- helper also checks the decoder's identity tag before borrowing.
borrowReadOnlyObjects :: CodegenMetadata -> BackendModule -> BackendModule
borrowReadOnlyObjects metadata mod = case Map.lookup "Data.Argonaut.Decode.Internal.Record.borrowObject" metadata.globalTypes of
  Nothing -> mod
  Just helperType -> mod { bindings = map rewriteGroup mod.bindings }
    where
    definitions = Map.fromFoldable (Array.concatMap (\group -> if group.recursive then [] else group.bindings) mod.bindings)
    rewriteGroup group
      | group.recursive = group
      | otherwise = group { bindings = map (\(Tuple name expr) -> Tuple name (rewrite expr)) group.bindings }
    rewrite original@(NeutralExpr syn) = case syn of
      LetRec _ _ _ -> original
      Let name level value body ->
        let value' = case strip value of
              App fn args | identityMethod fn && safe level Set.empty body ->
                case NEA.toArray args of
                  [ json ] -> replace value (NeutralExpr (App
                    (NeutralExpr (Typed helperType (NeutralExpr (Var helper))))
                    (NEA.cons' fn [ json ])))
                  _ -> rewrite value
              _ -> rewrite value
        in NeutralExpr (Let name level value' (rewrite body))
      _ -> NeutralExpr (map rewrite syn)
    identityMethod fn = case strip fn of
      Accessor dict (GetProp "decodeJson") -> identityDictionary Set.empty dict
      _ -> false
    identityDictionary visited expr = case strip expr of
      Var (Qualified (Just owner) name) | owner == mod.name && not (Set.member name visited) ->
        case Map.lookup name definitions of
          Just value -> identityDictionary (Set.insert name visited) value
          Nothing -> false
      App fn args -> isGlobal "Data.Argonaut.Decode.Class" "decodeForeignObject" fn
        && case NEA.toArray args of
          [ inner ] -> isGlobal "Data.Argonaut.Decode.Class" "decodeJsonJson" inner
          _ -> false
      _ -> false

helper :: Qualified Ident
helper = Qualified (Just (ModuleName "Data.Argonaut.Decode.Internal.Record")) (Ident "borrowObject")

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syn) = case syn of
  Typed _ inner -> strip inner
  TypeApp inner _ -> strip inner
  _ -> syn

-- Preserve call-site annotations and type applications exactly.
replace :: NeutralExpr -> NeutralExpr -> NeutralExpr
replace (NeutralExpr syn) replacement = case syn of
  Typed ty inner -> NeutralExpr (Typed ty (replace inner replacement))
  TypeApp inner ty -> NeutralExpr (TypeApp (replace inner replacement) ty)
  _ -> replacement

isGlobal :: String -> String -> NeutralExpr -> Boolean
isGlobal owner name expr = case strip expr of
  Var (Qualified (Just (ModuleName m)) (Ident n)) -> m == owner && n == name
  _ -> false

isLocal :: Level -> NeutralExpr -> Boolean
isLocal level expr = case strip expr of
  Local _ found -> level == found
  _ -> false

isEither :: String -> Qualified Ident -> Boolean
isEither name = (_ == Qualified (Just (ModuleName "Data.Either")) (Ident name))

objectSource :: Level -> Set.Set Level -> NeutralExpr -> Boolean
objectSource result objects expr = case strip expr of
  Local _ level -> Set.member level objects
  Accessor source (GetCtorField ctor _ _ _ "value0" 0) -> isEither "Right" ctor && isLocal result source
  _ -> false

fieldReader :: NeutralExpr -> Boolean
fieldReader expr = case strip expr of
  Var (Qualified (Just (ModuleName "Data.Argonaut.Decode.Decoders")) (Ident name)) ->
    Array.any (\base -> name == base || case String.stripPrefix (Pattern (base <> "__")) name of
      Just suffix -> suffix /= "" && Array.all (\c -> c >= '0' && c <= '9') (CU.toCharArray suffix)
      Nothing -> false) [ "getField", "getFieldOptional", "getFieldOptional'" ]
  _ -> false

safe :: Level -> Set.Set Level -> NeutralExpr -> Boolean
safe result objects original@(NeutralExpr syn) = case syn of
  Typed _ inner -> safe result objects inner
  TypeApp inner _ -> safe result objects inner
  Local _ level -> level /= result && not (Set.member level objects)
  PrimOp (Op1 (OpIsTag ctor) source) | isEither "Left" ctor || isEither "Right" ctor ->
    isLocal result source || safe result objects source
  Accessor source (GetCtorField ctor _ _ _ "value0" 0) | isEither "Left" ctor && isLocal result source -> true
  Let _ level value body | objectSource result objects value ->
    level /= result && safe result (Set.insert level objects) body
  Let _ level value body -> safe result objects value
    && (if level == result then independent result objects body else safe result (Set.delete level objects) body)
  App fn args | fieldReader fn -> case NEA.toArray args of
    [ decoder, object, key ] | objectSource result objects object -> safe result objects decoder && safe result objects key
    _ -> all (safe result objects) syn
  Abs _ _ -> independent result objects original
  UncurriedAbs _ _ -> independent result objects original
  UncurriedEffectAbs _ _ -> independent result objects original
  LetRec _ _ _ -> independent result objects original
  EffectBind _ _ _ _ -> independent result objects original
  EffectDefer _ -> independent result objects original
  EffectPure _ -> independent result objects original
  PrimEffect _ -> independent result objects original
  UncurriedEffectApp _ _ -> independent result objects original
  _ -> all (safe result objects) syn

-- Conservative: even a shadowed occurrence in a deferred/recursive scope
-- rejects borrowing. False negatives cannot introduce a captured borrowed map.
independent :: Level -> Set.Set Level -> NeutralExpr -> Boolean
independent result objects (NeutralExpr syn) = case syn of
  Local _ level -> level /= result && not (Set.member level objects)
  _ -> foldl (\ok child -> ok && independent result objects child) true syn
