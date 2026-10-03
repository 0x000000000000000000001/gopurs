module Gopurs.Monomorphization
  ( monomorphizeModules
  , monomorphizeModulesWith
  , collectForeignForwarders
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
import Gopurs.Monomorphization.ForeignForwarders as ForeignForwarders
import Gopurs.NativeRecordArgs (candidateToShare)
import PureScript.Backend.Optimizer.CoreFn (Ann, Bind(..), Binding(..), ExprType(..), Ident(..), Module(..))
import PureScript.Backend.Optimizer.CoreFn.Usage (invalidateSourceUsageModule)
import PureScript.Backend.Optimizer.Monomorphize (InstantiationMap, collectInstantiations, monomorphize, transitiveCollect)
import PureScript.Backend.Optimizer.NativeMaps (insertStringImpl, stringCompare)

type GlobalAstMap = Map String (Binding Ann)

-- These are late barriers: available bodies still participate in dictionary
-- evaluation and transitive discovery. Only intrinsicGlobals is removed from
-- the AST index itself, before collection can unfold it.
type SpecializationBarriers =
  { foreignGlobals :: Set String
  , foreignForwarders :: Set String
  , sharedRecordWorkers :: Set String
  }

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
    barriers =
      { foreignGlobals: Set.union intrinsicGlobals (collectForeignGlobals modules)
      , foreignForwarders: collectForeignForwarders modules
      , sharedRecordWorkers: Map.keys (Map.filter candidateToShare globalAstMap)
      }
    rawInstantiations = foldl (collectInstantiations globalAstMap) Map.empty modules
  -- PBO owns the immutable round snapshots, ordered merge and fixed point. Do
  -- not filter its inputs: excluded definitions can expose eligible callees.
  transitiveInstantiations <- collectTransitive globalAstMap rawInstantiations
  let instantiations = Map.filterKeys (eligible globalTypes barriers) transitiveInstantiations
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
  insertStringImpl stringCompare (moduleName <> "." <> unwrap ident) binding bindings

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

-- Admission uses the original source type, never a use-site instantiation or
-- an annotation from a specialized copy. Missing types and explicit Any reject.
eligible :: Map String ExprType -> SpecializationBarriers -> String -> Boolean
eligible globalTypes barriers name =
  not (Set.member name barriers.sharedRecordWorkers)
    && not (Set.member name barriers.foreignForwarders)
    && not (Set.member name barriers.foreignGlobals)
    && case Map.lookup name globalTypes of
      Just ty -> hasSpecializableVariables ty
      Nothing -> false

-- Retain the public recognition entry point for callers and contract tests.
collectForeignForwarders :: List (Module Ann) -> Set String
collectForeignForwarders = ForeignForwarders.collectForeignForwarders

-- Keep the source naming convention: uppercase opaque names and the runtime
-- Value placeholder are not specialization variables. Quantifiers alone do not
-- suffice; a variable must occur in the body, a row tail or a constraint argument.
hasSpecializableVariables :: ExprType -> Boolean
hasSpecializableVariables (TypeVar v) = String.take 1 v == String.toLower (String.take 1 v) && v /= "gopurs_runtime.Value"

hasSpecializableVariables (Func args ret) = Array.any hasSpecializableVariables args || hasSpecializableVariables ret
hasSpecializableVariables (Array t) = hasSpecializableVariables t
hasSpecializableVariables (Record row) = hasSpecializableVariables row
hasSpecializableVariables (Row props tail) =
  let tailHas = case tail of
        Nothing -> false
        Just t -> hasSpecializableVariables t
  in Array.any (\(Tuple _ v) -> hasSpecializableVariables v) props || tailHas
hasSpecializableVariables (TypeApp c args) = hasSpecializableVariables c || Array.any hasSpecializableVariables args
hasSpecializableVariables (ForAll _ body) = hasSpecializableVariables body
hasSpecializableVariables (ConstrainedType constraints body) = Array.any (\(Tuple _ a) -> Array.any hasSpecializableVariables a) constraints || hasSpecializableVariables body
hasSpecializableVariables Int = false
hasSpecializableVariables String = false
hasSpecializableVariables Char = false
hasSpecializableVariables Number = false
hasSpecializableVariables Boolean = false
hasSpecializableVariables Unit = false
hasSpecializableVariables (TypeLevelString _) = false
hasSpecializableVariables (ADT _ _ args) = Array.any hasSpecializableVariables args
hasSpecializableVariables Any = false
