module Gopurs.GlobalTypes
  ( buildGlobalTypes
  ) where

import Prelude

import Data.Array as Array
import Data.Foldable (foldl)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binding(..), Expr(..), ExprType(..), Ident(..), Module(..))
import PureScript.Backend.Optimizer.Monomorphize (getExprAnn)

-- Index TAST types by qualified name, including foreign declarations.
buildGlobalTypes :: Array (Module Ann) -> Map String ExprType
buildGlobalTypes = Array.foldl addModuleTypes Map.empty

addModuleTypes :: Map String ExprType -> Module Ann -> Map String ExprType
addModuleTypes types (Module mod) =
  let
    moduleName = unwrap mod.name
    withDefinitions = Array.foldl (addBindTypes moduleName) types mod.decls
    foreignTypes = Map.toUnfoldable mod.foreign :: Array (Tuple Ident (Maybe ExprType))
  in
    foldl (addForeignType moduleName) withDefinitions foreignTypes

addBindTypes :: String -> Map String ExprType -> Bind Ann -> Map String ExprType
addBindTypes moduleName types = case _ of
  NonRec binding -> addBindingType moduleName types binding
  Rec bindings -> Array.foldl (addBindingType moduleName) types bindings

addBindingType :: String -> Map String ExprType -> Binding Ann -> Map String ExprType
addBindingType moduleName types binding@(Binding _ (Ident name) _) =
  case bindingType binding of
    Just ty -> Map.insert (moduleName <> "." <> name) ty types
    Nothing -> types

-- Prefer the binding annotation, then the expression annotation, then the
-- existing fallback for application results. An explicit Any is preserved.
bindingType :: Binding Ann -> Maybe ExprType
bindingType (Binding (Ann annotation) _ expr) = case annotation.type of
  Just ty -> Just ty
  Nothing -> case getExprAnn expr of
    Ann { type: Just ty } -> Just ty
    _ -> inferExprType expr

addForeignType :: String -> Map String ExprType -> Tuple Ident (Maybe ExprType) -> Map String ExprType
addForeignType moduleName types (Tuple (Ident name) mbType) = case mbType of
  Just ty -> Map.insert (moduleName <> "." <> name) ty types
  Nothing -> types

inferExprType :: Expr Ann -> Maybe ExprType
inferExprType (ExprApp _ fn _) = case getExprAnn fn of
  Ann { type: Just ty } -> getReturnType ty
  _ -> case inferExprType fn of
    Just ty -> getReturnType ty
    Nothing -> Nothing
inferExprType (ExprTypeApp _ fn _) = inferExprType fn
inferExprType _ = Nothing

getReturnType :: ExprType -> Maybe ExprType
getReturnType (ForAll _ ty) = getReturnType ty
getReturnType (ConstrainedType _ ty) = getReturnType ty
getReturnType (Func _ ret) = Just ret
getReturnType _ = Nothing
