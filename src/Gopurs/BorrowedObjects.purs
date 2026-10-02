module Gopurs.BorrowedObjects (borrowReadOnlyObjects) where

import Prelude

import Control.Alternative (guard)
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
import PureScript.Backend.Optimizer.CoreFn (ExprType, Ident(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendSyntax(..), Level)

type BorrowContext =
  { moduleName :: ModuleName
  , definitions :: Map.Map Ident NeutralExpr
  , helperType :: ExprType
  }

type Borrowing = { method :: NeutralExpr, json :: NeutralExpr }

-- The result is the original Either envelope; objects are aliases of its Right
-- payload. Only lexical levels identify these references, not local spellings.
type BorrowScope = { result :: Level, objects :: Set.Set Level }

-- Elide an identity-object copy only when every use of the Either and its
-- Right payload is accounted for. The object may feed saturated public field
-- readers; borrowed references cannot escape through closures, effects,
-- recursion or arbitrary consumers. Custom element decoders never match the
-- producer. The native helper also checks its identity tag before borrowing.
borrowReadOnlyObjects :: CodegenMetadata -> BackendModule -> BackendModule
borrowReadOnlyObjects metadata mod = case Map.lookup "Data.Argonaut.Decode.Internal.Record.borrowObject" metadata.globalTypes of
  Nothing -> mod
  Just helperType ->
    let context =
          { moduleName: mod.name, helperType
          , definitions: Map.fromFoldable (Array.concatMap (\group -> if group.recursive then [] else group.bindings) mod.bindings)
          }
        rewriteGroup group
          | group.recursive = group
          | otherwise = group { bindings = map (\(Tuple name expr) -> Tuple name (rewrite context expr)) group.bindings }
    in mod { bindings = map rewriteGroup mod.bindings }

rewrite :: BorrowContext -> NeutralExpr -> NeutralExpr
rewrite context original@(NeutralExpr syn) = case syn of
  LetRec _ _ _ -> original
  Let name level value body ->
    let value' = case admitBorrowing context level value body of
          Just borrowing -> replaceAnnotated value (borrowedDecode context.helperType borrowing)
          Nothing -> rewrite context value
    in NeutralExpr (Let name level value' (rewrite context body))
  _ -> NeutralExpr (map (rewrite context) syn)

-- Prove the producer and every use of the result before constructing a plan.
-- Keep the exact unary curried shape; do not flatten an application spine or
-- infer identity from the dictionary's type alone.
admitBorrowing :: BorrowContext -> Level -> NeutralExpr -> NeutralExpr -> Maybe Borrowing
admitBorrowing context result value body = case strip value of
  App fn args -> do
    guard (identityMethod context fn)
    guard (readOnlyUses { result, objects: Set.empty } body)
    case NEA.toArray args of
      [ json ] -> Just { method: fn, json }
      _ -> Nothing
  _ -> Nothing

-- An admitted producer keeps its original method and JSON operands verbatim.
-- Only rejected producers recurse into their value; every continuation is still
-- rewritten. The helper already exists: this pass emits no declarations.
borrowedDecode :: ExprType -> Borrowing -> NeutralExpr
borrowedDecode helperType { method, json } = NeutralExpr (App
  (NeutralExpr (Typed helperType (NeutralExpr (Var helper))))
  (NEA.cons' method [ json ]))

identityMethod :: BorrowContext -> NeutralExpr -> Boolean
identityMethod context fn = case strip fn of
  Accessor dict (GetProp "decodeJson") -> identityDictionary context Set.empty dict
  _ -> false

-- Only qualified nonrecursive definitions in this module can be followed.
-- The visited set rejects cycles without introducing initialization dependencies.
identityDictionary :: BorrowContext -> Set.Set Ident -> NeutralExpr -> Boolean
identityDictionary context visited expr = case strip expr of
  Var (Qualified (Just owner) name) | owner == context.moduleName && not (Set.member name visited) ->
    case Map.lookup name context.definitions of
      Just value -> identityDictionary context (Set.insert name visited) value
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
replaceAnnotated :: NeutralExpr -> NeutralExpr -> NeutralExpr
replaceAnnotated (NeutralExpr syn) replacement = case syn of
  Typed ty inner -> NeutralExpr (Typed ty (replaceAnnotated inner replacement))
  TypeApp inner ty -> NeutralExpr (TypeApp (replaceAnnotated inner replacement) ty)
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

objectSource :: BorrowScope -> NeutralExpr -> Boolean
objectSource { result, objects } expr = case strip expr of
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

-- A synchronous reader may inspect the envelope, propagate Left, or pass the
-- Right payload in a saturated field reader's object slot. Its decoder and key
-- must independently satisfy the same proof; neither may carry the borrowed map.
readOnlyUses :: BorrowScope -> NeutralExpr -> Boolean
readOnlyUses scope@{ result, objects } original@(NeutralExpr syn) = case syn of
  Typed _ inner -> readOnlyUses scope inner
  TypeApp inner _ -> readOnlyUses scope inner
  Local _ level -> level /= result && not (Set.member level objects)
  PrimOp (Op1 (OpIsTag ctor) source) | isEither "Left" ctor || isEither "Right" ctor ->
    isLocal result source || readOnlyUses scope source
  Accessor source (GetCtorField ctor _ _ _ "value0" 0) | isEither "Left" ctor && isLocal result source -> true
  Let _ level value body | objectSource scope value ->
    level /= result && readOnlyUses (scope { objects = Set.insert level objects }) body
  Let _ level value body -> readOnlyUses scope value
    && (if level == result then independentScope scope body else readOnlyUses (scope { objects = Set.delete level objects }) body)
  App fn args | fieldReader fn -> case NEA.toArray args of
    [ decoder, object, key ] | objectSource scope object -> readOnlyUses scope decoder && readOnlyUses scope key
    _ -> all (readOnlyUses scope) syn
  Abs _ _ -> independentScope scope original
  UncurriedAbs _ _ -> independentScope scope original
  UncurriedEffectAbs _ _ -> independentScope scope original
  LetRec _ _ _ -> independentScope scope original
  EffectBind _ _ _ _ -> independentScope scope original
  EffectDefer _ -> independentScope scope original
  EffectPure _ -> independentScope scope original
  PrimEffect _ -> independentScope scope original
  UncurriedEffectApp _ _ -> independentScope scope original
  _ -> all (readOnlyUses scope) syn

-- Conservative: even a shadowed occurrence in a deferred/recursive scope
-- rejects borrowing. False negatives cannot introduce a captured borrowed map.
independentScope :: BorrowScope -> NeutralExpr -> Boolean
independentScope scope@{ result, objects } (NeutralExpr syn) = case syn of
  Local _ level -> level /= result && not (Set.member level objects)
  _ -> foldl (\ok child -> ok && independentScope scope child) true syn
