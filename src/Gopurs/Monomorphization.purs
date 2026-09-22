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
import PureScript.Backend.Optimizer.CoreFn (Ann, Bind(..), Binding(..), ExprType(..), Ident(..), Module(..))
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
    globalAstMap = Map.filterKeys (not <<< flip Set.member intrinsicGlobals) (buildGlobalAstMap modules)
    -- Row-only readers can share one native worker across record shapes. The
    -- emitter rechecks their uses after PBO; other polymorphism stays eligible.
    sharedRecordWorkers = Map.keys (Map.filter candidateToShare globalAstMap)
    rawInstantiations = foldl (collectInstantiations globalAstMap) Map.empty modules
    foreignGlobals = Set.union intrinsicGlobals (collectForeignGlobals modules)
  transitiveInstantiations <- collectTransitive globalAstMap rawInstantiations
  let instantiations = Map.filterKeys
        (\name -> not (Set.member name sharedRecordWorkers) && shouldMonomorphize globalTypes foreignGlobals name)
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
