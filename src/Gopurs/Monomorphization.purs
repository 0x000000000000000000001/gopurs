module Gopurs.Monomorphization
  ( monomorphizeModules
  ) where

import Prelude

import Data.Array as Array
import Data.Foldable (foldl)
import Data.List (List)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.CoreFn (Ann, Bind(..), Binding(..), ExprType(..), Ident(..), Module(..), Qualified(..))
import PureScript.Backend.Optimizer.Monomorphize (collectInstantiations, monomorphize, transitiveCollect)
import PureScript.Backend.Optimizer.Semantics.Foreign (coreForeignSemantics)

type GlobalAstMap = Map String (Binding Ann)

-- Receive the original global types and modules enriched with class declarations.
monomorphizeModules :: Map String ExprType -> List (Module Ann) -> List (Module Ann)
monomorphizeModules globalTypes modules =
  let
    -- Intrinsics must retain their names and dictionary arguments so the
    -- evaluator can recognize them before their PureScript bodies expand.
    intrinsicGlobals = Set.fromFoldable $ Array.mapMaybe
      (case _ of
        Tuple (Qualified (Just moduleName) ident) _ -> Just (unwrap moduleName <> "." <> unwrap ident)
        _ -> Nothing)
      (Map.toUnfoldable coreForeignSemantics)
    globalAstMap = Map.filterKeys (not <<< flip Set.member intrinsicGlobals) (buildGlobalAstMap modules)
    rawInstantiations = foldl (collectInstantiations globalAstMap) Map.empty modules
    transitiveInstantiations = transitiveCollect globalAstMap rawInstantiations
    foreignGlobals = Set.union intrinsicGlobals (collectForeignGlobals modules)
    instantiations = Map.filterKeys (shouldMonomorphize globalTypes foreignGlobals) transitiveInstantiations
  in
    if Map.isEmpty instantiations then
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
