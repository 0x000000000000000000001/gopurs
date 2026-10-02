module Gopurs.ClosedDictionaries.Admission
  ( SharedDictionaries
  , sharedDictionaries
  , reusable
  , liftable
  ) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty (toArray)
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.ClosedDictionaries.Scope (closedConstruction)
import Gopurs.CodegenState (CodegenMetadata)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType, Ident(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Substitute (unify)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))
import PureScript.Backend.Optimizer.TypeSubstitution (substitute)

-- Reuse and lifting have different proofs. Reuse recovers a missing annotation
-- from an existing imported application; lifting trusts the site's class
-- annotation and proves that the whole construction is closed and synchronous.
newtype SharedDictionaries = SharedDictionaries
  (Map.Map ApplicationKey { ident :: Ident, type :: ExprType })

-- Keep qualifications, curried/uncurried conventions and spine grouping exact.
-- Locals, type applications and annotated children cannot form keys: erasing
-- them could conflate different instantiations or coercions.
data ApplicationKey
  = GlobalKey (Qualified Ident)
  | CurriedKey ApplicationKey (Array ApplicationKey)
  | UncurriedKey ApplicationKey (Array ApplicationKey)

derive instance eqApplicationKey :: Eq ApplicationKey
derive instance ordApplicationKey :: Ord ApplicationKey

applicationKey :: NeutralExpr -> Maybe ApplicationKey
applicationKey (NeutralExpr syn) = case syn of
  Var name -> Just (GlobalKey name)
  App fn args -> CurriedKey <$> applicationKey fn <*> traverse applicationKey (toArray args)
  UncurriedApp fn args -> UncurriedKey <$> applicationKey fn <*> traverse applicationKey args
  _ -> Nothing

-- Collect once from the original module, preserving the first admissible
-- binding in source order. Newly lifted sites never join this reuse index.
sharedDictionaries :: CodegenMetadata -> BackendModule -> SharedDictionaries
sharedDictionaries metadata mod = SharedDictionaries (foldl addGroup Map.empty mod.bindings)
  where
  addGroup entries group
    | group.recursive = entries
    | otherwise = foldl addBinding entries group.bindings

  addBinding entries (Tuple ident expr) = case expr of
    NeutralExpr (Typed ty body) -> case dictionaryClass ty, applicationKey body of
      Just name, Just key | Map.member name metadata.classDeclsFields && isApplication body
        && monomorphic ty && importedApplicationType metadata mod.name key == Just ty ->
        if Map.member key entries then entries
        else Map.insert key { ident, type: ty } entries
      _, _ -> entries
    _ -> entries

reusable :: SharedDictionaries -> NeutralExpr -> Maybe { ident :: Ident, type :: ExprType }
reusable (SharedDictionaries entries) expr = applicationKey expr >>= flip Map.lookup entries

-- This is intentionally not general inference. Missing types, local-module or
-- unqualified globals, higher-rank arguments, unresolved variables and dynamic
-- Any decline reuse. Imported globals avoid new local initialization dependencies.
importedApplicationType :: CodegenMetadata -> ModuleName -> ApplicationKey -> Maybe ExprType
importedApplicationType metadata current = case _ of
  GlobalKey (Qualified (Just moduleName@(CoreFn.ModuleName name)) (Ident ident)) -> do
    guard (moduleName /= current)
    Map.lookup (name <> "." <> ident) metadata.globalTypes
  CurriedKey fn args -> infer fn args
  UncurriedKey fn args -> infer fn args
  _ -> Nothing
  where
  infer fn args = do
    fnType <- importedApplicationType metadata current fn
    argTypes <- traverse (importedApplicationType metadata current) args
    guard (Array.all monomorphic argTypes)
    applicationResult fnType argTypes

applicationResult :: ExprType -> Array ExprType -> Maybe ExprType
applicationResult ty args = case Array.uncons args of
  Nothing -> Just ty
  Just { head: arg, tail: rest } -> case ty of
    CoreFn.ForAll _ body -> applicationResult body args
    CoreFn.ConstrainedType constraints body -> applicationResult
      (CoreFn.Func (map (\(Tuple path types) -> CoreFn.ADT (String.joinWith "." path) path types) constraints) body) args
    CoreFn.Func params result -> do
      { head: param, tail: remaining } <- Array.uncons params
      let substitution = unify param arg Map.empty
      guard (substitute substitution param == arg)
      let next = if Array.null remaining then result else CoreFn.Func remaining result
      applicationResult (substitute substitution next) rest
    _ -> Nothing

monomorphic :: ExprType -> Boolean
monomorphic = case _ of
  CoreFn.Any -> false
  CoreFn.TypeVar _ -> false
  CoreFn.ForAll _ _ -> false
  CoreFn.ConstrainedType _ _ -> false
  CoreFn.ADT _ _ args -> Array.all monomorphic args
  CoreFn.Array item -> monomorphic item
  CoreFn.Func args result -> Array.all monomorphic args && monomorphic result
  CoreFn.TypeApp fn args -> monomorphic fn && Array.all monomorphic args
  CoreFn.Record row -> monomorphic row
  CoreFn.Row fields tail -> Array.all (\(Tuple _ field) -> monomorphic field) fields
    && case tail of
      Nothing -> true
      Just row -> monomorphic row
  _ -> true

-- A new getter requires an outer class annotation and a closed application.
-- Unlike annotation-lost reuse, it need not infer a monomorphic imported type.
liftable :: CodegenMetadata -> NeutralExpr -> Maybe ExprType
liftable metadata expr = do
  guard (isApplication expr)
  guard (closedConstruction expr)
  ty <- annotatedType expr
  className <- dictionaryClass ty
  guard (Map.member className metadata.classDeclsFields)
  pure ty

isApplication :: NeutralExpr -> Boolean
isApplication expr = case strip expr of
  App _ _ -> true
  UncurriedApp _ _ -> true
  _ -> false

annotatedType :: NeutralExpr -> Maybe ExprType
annotatedType (NeutralExpr syn) = case syn of
  Typed ty _ -> Just ty
  _ -> Nothing

dictionaryClass :: ExprType -> Maybe String
dictionaryClass = case _ of
  CoreFn.ADT name _ _ -> Just name
  CoreFn.ForAll _ ty -> dictionaryClass ty
  CoreFn.ConstrainedType _ ty -> dictionaryClass ty
  CoreFn.TypeApp ty _ -> dictionaryClass ty
  _ -> Nothing

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syn) = case syn of
  Typed _ inner -> strip inner
  TypeApp inner _ -> strip inner
  _ -> syn
