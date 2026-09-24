module Gopurs.Monomorphization
  ( monomorphizeModules
  , monomorphizeModulesWith
  ) where

import Prelude

import Data.Array as Array
import Data.Foldable (foldl)
import Data.Identity (Identity(..))
import Data.List (List)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.Tuple (Tuple(..))
import Gopurs.NativeRecordArgs (candidateToShare)
import PureScript.Backend.Optimizer.CoreFn (Ann, Bind(..), Binding(..), Expr(..), ExprType(..), Ident(..), Module(..), Qualified(..))
import PureScript.Backend.Optimizer.CoreFn.Usage (invalidateSourceUsageModule)
import PureScript.Backend.Optimizer.Monomorphize (InstantiationMap, collectInstantiations, monomorphize, transitiveCollect)

type GlobalAstMap = Map String (Binding Ann)

-- Receive the original global types and modules enriched with class declarations.
monomorphizeModules :: Map String ExprType -> List (Module Ann) -> List (Module Ann)
monomorphizeModules globalTypes inputModules = runIdentity $
  monomorphizeModulesWith (\ast instantiations -> pure (transitiveCollect ast instantiations)) globalTypes inputModules
  where
  runIdentity (Identity result) = result

-- The sequential entry point and Aff preparation share the same barriers and
-- final module ordering. Only the transitive collector is supplied by callers.
monomorphizeModulesWith
  :: forall m
   . Monad m
  => (GlobalAstMap -> InstantiationMap -> m InstantiationMap)
  -> Map String ExprType
  -> List (Module Ann)
  -> m (List (Module Ann))
monomorphizeModulesWith collectTransitive globalTypes inputModules = do
  let
    -- Source identities and usage proofs belong to the exported CoreFn.
    -- Specialization copies bindings; analyze the final IR afresh instead.
    modules = map invalidateSourceUsageModule inputModules
    -- Keep negate's dictionary until its signed-zero intrinsic is recognized.
    -- Other known definitions must remain available for static evaluation.
    intrinsicGlobals = Set.singleton "Data.Ring.negate"
    -- Thin wrappers over value-level foreign imports gain nothing from
    -- specialization: the FFI boundary is boxed by construction, so a
    -- specialized copy only rewrites the wrapper's own boundary to native
    -- representations and adds conversions around a foreign call that cannot
    -- use them. This subsumes the named `caseJson*` accessors.
    foreignForwarders = collectForeignForwarders modules
    globalAstMap = Map.filterKeys (not <<< flip Set.member intrinsicGlobals) (buildGlobalAstMap modules)
    -- Row-only readers can share one native worker across record shapes. The
    -- emitter rechecks their uses after PBO; other polymorphism stays eligible.
    sharedRecordWorkers = Map.keys (Map.filter candidateToShare globalAstMap)
    rawInstantiations = foldl (collectInstantiations globalAstMap) Map.empty modules
    foreignGlobals = Set.union intrinsicGlobals (collectForeignGlobals modules)
  transitiveInstantiations <- collectTransitive globalAstMap rawInstantiations
  let instantiations = Map.filterKeys
        (\name ->
          not (Set.member name sharedRecordWorkers)
            && not (Set.member name foreignForwarders)
            && shouldMonomorphize globalTypes foreignGlobals name)
        transitiveInstantiations
  pure $ if Map.isEmpty instantiations then
      modules
    else
      map (monomorphize globalAstMap instantiations) modules

buildGlobalAstMap :: List (Module Ann) -> GlobalAstMap
buildGlobalAstMap = foldl addModuleBindings Map.empty

addModuleBindings :: GlobalAstMap -> Module Ann -> GlobalAstMap
addModuleBindings bindings (Module mod) =
  Array.foldl (addBind (unwrap mod.name)) bindings mod.decls

addBind :: String -> GlobalAstMap -> Bind Ann -> GlobalAstMap
addBind moduleName bindings = case _ of
  NonRec binding -> addBinding moduleName bindings binding
  Rec group -> Array.foldl (addBinding moduleName) bindings group

addBinding :: String -> GlobalAstMap -> Binding Ann -> GlobalAstMap
addBinding moduleName bindings binding@(Binding _ ident _) =
  Map.insert (moduleName <> "." <> unwrap ident) binding bindings

collectForeignGlobals :: List (Module Ann) -> Set String
collectForeignGlobals = foldl addModuleForeignGlobals Set.empty

addModuleForeignGlobals :: Set String -> Module Ann -> Set String
addModuleForeignGlobals globals (Module mod) =
  let
    foreigns = Map.toUnfoldable mod.foreign :: Array (Tuple Ident (Maybe ExprType))
  in
    Array.foldl (addForeignGlobal (unwrap mod.name)) globals foreigns

addForeignGlobal :: String -> Set String -> Tuple Ident (Maybe ExprType) -> Set String
addForeignGlobal moduleName globals (Tuple (Ident name) _) =
  Set.insert (moduleName <> "." <> name) globals

-- Filter only after transitive collection, using the original global type table.
shouldMonomorphize :: Map String ExprType -> Set String -> String -> Boolean
shouldMonomorphize globalTypes foreignGlobals name =
  not (Set.member name foreignGlobals) && case Map.lookup name globalTypes of
    Just ty -> hasTypeVariables ty
    Nothing -> false

-- | Bindings whose body eta-forwards to a foreign import of the same module,
-- | optionally wrapping arguments and result in `unsafeCoerce`. Specializing
-- | such a binding cannot unbox anything across the FFI, so it is skipped.
collectForeignForwarders :: List (Module Ann) -> Set String
collectForeignForwarders = foldl addModuleForwarders Set.empty

addModuleForwarders :: Set String -> Module Ann -> Set String
addModuleForwarders acc (Module mod) =
  let
    moduleName = unwrap mod.name
    foreignIdents = Set.fromFoldable
      (map (\(Tuple (Ident name) _) -> name) (Map.toUnfoldable mod.foreign :: Array (Tuple Ident (Maybe ExprType))))
  in
    Array.foldl (addForwardingBind moduleName foreignIdents) acc mod.decls

addForwardingBind :: String -> Set String -> Set String -> Bind Ann -> Set String
addForwardingBind moduleName foreignIdents acc = case _ of
  NonRec binding -> addForwarder moduleName foreignIdents acc binding
  Rec group -> Array.foldl (addForwarder moduleName foreignIdents) acc group

addForwarder :: String -> Set String -> Set String -> Binding Ann -> Set String
addForwarder moduleName foreignIdents acc (Binding _ (Ident ident) body) =
  if isForeignForwarder moduleName foreignIdents body then
    Set.insert (moduleName <> "." <> ident) acc
  else acc

isForeignForwarder :: String -> Set String -> Expr Ann -> Boolean
isForeignForwarder moduleName foreignIdents body = case collectLambdaParams body of
  { params, body: inner } ->
    not (Array.null params) && case unapplyExpr inner of
      { head: ExprVar _ (Qualified mbModule (Ident ffi)), args } ->
        Set.member ffi foreignIdents
          && sameModule mbModule
          && Array.length args == Array.length params
          && foldl (&&) true (Array.zipWith isParameterReference args params)
      _ -> false
  where
  sameModule = case _ of
    Just mn -> unwrap mn == moduleName
    Nothing -> true

collectLambdaParams :: Expr Ann -> { params :: Array Ident, body :: Expr Ann }
collectLambdaParams (ExprAbs _ param body) =
  case collectLambdaParams body of
    rest -> rest { params = Array.cons param rest.params }
collectLambdaParams body = { params: [], body }

unapplyExpr :: Expr Ann -> { head :: Expr Ann, args :: Array (Expr Ann) }
unapplyExpr expr = case stripCoercions expr of
  ExprApp _ fn arg ->
    case unapplyExpr fn of
      inner -> inner { args = Array.snoc inner.args arg }
  other -> { head: other, args: [] }

stripCoercions :: Expr Ann -> Expr Ann
stripCoercions expr = case expr of
  ExprTypeApp _ inner _ -> stripCoercions inner
  ExprApp _ fn arg | isUnsafeCoerce fn -> stripCoercions arg
  _ -> expr

-- `unsafeCoerce` is polymorphic, so its use is wrapped in type applications.
isUnsafeCoerce :: Expr Ann -> Boolean
isUnsafeCoerce fn = case stripTypeApps fn of
  ExprVar _ (Qualified _ (Ident "unsafeCoerce")) -> true
  _ -> false

stripTypeApps :: Expr Ann -> Expr Ann
stripTypeApps (ExprTypeApp _ inner _) = stripTypeApps inner
stripTypeApps expr = expr

isParameterReference :: Expr Ann -> Ident -> Boolean
isParameterReference expr (Ident name) = case stripCoercions expr of
  ExprVar _ (Qualified Nothing (Ident used)) -> used == name
  _ -> false

hasTypeVariables :: ExprType -> Boolean
hasTypeVariables (TypeVar v) = String.take 1 v == String.toLower (String.take 1 v) && v /= "gopurs_runtime.Value"

hasTypeVariables (Func args ret) = Array.any hasTypeVariables args || hasTypeVariables ret
hasTypeVariables (Array t) = hasTypeVariables t
hasTypeVariables (Record row) = hasTypeVariables row
hasTypeVariables (Row props tail) =
  let tailHas = case tail of
        Nothing -> false
        Just t -> hasTypeVariables t
  in Array.any (\(Tuple _ v) -> hasTypeVariables v) props || tailHas
hasTypeVariables (TypeApp c args) = hasTypeVariables c || Array.any hasTypeVariables args
hasTypeVariables (ForAll _ body) = hasTypeVariables body
hasTypeVariables (ConstrainedType constraints body) = Array.any (\(Tuple _ a) -> Array.any hasTypeVariables a) constraints || hasTypeVariables body
hasTypeVariables Int = false
hasTypeVariables String = false
hasTypeVariables Char = false
hasTypeVariables Number = false
hasTypeVariables Boolean = false
hasTypeVariables Unit = false
hasTypeVariables (TypeLevelString _) = false
hasTypeVariables (ADT _ _ args) = Array.any hasTypeVariables args
hasTypeVariables Any = false
