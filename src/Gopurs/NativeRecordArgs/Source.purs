module Gopurs.NativeRecordArgs.Source (candidateToShare) where

import Prelude

import Data.Array as Array
import Data.Foldable (all, any)
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..), fst)
import Gopurs.NativeRecordArgs.Projection (projectableFields, shareableResult)
import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binder(..), Binding(..), CaseAlternative(..), CaseGuard(..), Expr(..), ExprType(..), Guard(..), Ident, Qualified(..), propValue)
import PureScript.Backend.Optimizer.Monomorphize (getExprAnn)

type SharedSignature =
  { args :: Array ExprType
  , result :: ExprType
  , quantified :: Array String
  }

-- This source-level proof only suppresses monomorphisation. Code generation
-- independently rechecks every projected parameter on the transformed TCO body.
candidateToShare :: Binding Ann -> Boolean
candidateToShare binding@(Binding _ _ expr) = case bindingType binding >>= sharedSignature [] of
  Just { args, result, quantified } ->
    let
      lambdas = collectArguments [] expr
      tails = Array.mapMaybe rowVariable args
      validType ty = case projectableFields ty of
        Just _ -> true
        Nothing -> isMonomorphic ty
      validParameter (Tuple ident ty) = case projectableFields ty of
        Just fields -> readOnlyUses true ident (map fst fields) lambdas.body
        Nothing -> true
    in
      shareableResult result
        && not (Array.null tails)
        && all (flip Array.elem tails) quantified
        && all validType args
        && Array.length lambdas.args == Array.length args
        && Array.length (Array.nub lambdas.args) == Array.length lambdas.args
        && all validParameter (Array.zip lambdas.args args)
  Nothing -> false

-- Only an absent binding annotation permits expression-type fallback. Explicit
-- Any remains unknown even when the expression has a more specific annotation.
bindingType :: Binding Ann -> Maybe ExprType
bindingType (Binding (Ann ann) _ expr) = case ann.type of
  Just ty -> Just ty
  Nothing -> let Ann exprAnn = getExprAnn expr in exprAnn.type

sharedSignature :: Array String -> ExprType -> Maybe SharedSignature
sharedSignature quantified (ForAll vars body) = sharedSignature (quantified <> vars) body
sharedSignature quantified (Func args result) = Just { args, result, quantified }
sharedSignature _ _ = Nothing

rowVariable :: ExprType -> Maybe String
rowVariable (Record (Row _ (Just (TypeVar variable)))) = Just variable
rowVariable _ = Nothing

-- Ordinary companion parameters may contain closed rows, but no open tail,
-- dynamic Any, quantifier or constraint. This is stricter than result admission.
isMonomorphic :: ExprType -> Boolean
isMonomorphic = case _ of
  TypeVar _ -> false
  Any -> false
  ForAll _ _ -> false
  ConstrainedType _ _ -> false
  Array element -> isMonomorphic element
  ADT _ _ args -> all isMonomorphic args
  TypeApp fn args -> isMonomorphic fn && all isMonomorphic args
  Func args result -> all isMonomorphic args && isMonomorphic result
  Record row -> isMonomorphic row
  Row fields tail ->
    all (\(Tuple _ ty) -> isMonomorphic ty) fields
      && case tail of
        Nothing -> true
        Just _ -> false
  _ -> true

collectArguments :: Array Ident -> Expr Ann -> { args :: Array Ident, body :: Expr Ann }
collectArguments args (ExprAbs _ ident body) = collectArguments (Array.snoc args ident) body
collectArguments args body = { args, body }

-- allowRead becomes false under closures: even a getter would capture the row.
-- Source identifiers obey lexical shadowing in lambdas, lets and case patterns.
readOnlyUses :: Boolean -> Ident -> Array String -> Expr Ann -> Boolean
readOnlyUses allowRead target fields = case _ of
  ExprVar _ (Qualified Nothing ident) -> ident /= target
  ExprVar _ _ -> true
  ExprLit _ literal -> all (readOnlyUses allowRead target fields) literal
  ExprConstructor _ _ _ _ -> true
  ExprAccessor _ (ExprVar _ (Qualified Nothing ident)) field | ident == target ->
    allowRead && Array.elem field fields
  ExprAccessor _ object _ -> readOnlyUses allowRead target fields object
  ExprUpdate _ object props ->
    readOnlyUses allowRead target fields object
      && all (readOnlyUses allowRead target fields <<< propValue) props
  ExprAbs _ ident body -> ident == target || readOnlyUses false target fields body
  ExprApp _ fn arg -> readOnlyUses allowRead target fields fn && readOnlyUses allowRead target fields arg
  ExprCase _ values branches ->
    all (readOnlyUses allowRead target fields) values
      && all safeAlternative branches
  ExprLet _ bindings body -> safeBindings bindings body
  ExprTypeApp _ inner _ -> readOnlyUses allowRead target fields inner
  where
  safeAlternative (CaseAlternative binders result) =
    any (binds target) binders || case result of
      Unconditional body -> readOnlyUses allowRead target fields body
      Guarded guards -> all (\(Guard condition body) ->
        readOnlyUses allowRead target fields condition && readOnlyUses allowRead target fields body) guards

  safeBindings bindings body = case Array.uncons bindings of
    Nothing -> readOnlyUses allowRead target fields body
    Just { head: NonRec (Binding _ ident value), tail } ->
      readOnlyUses allowRead target fields value
        && (ident == target || safeBindings tail body)
    Just { head: Rec group, tail } ->
      any (\(Binding _ ident _) -> ident == target) group
        || (all (\(Binding _ _ value) -> readOnlyUses allowRead target fields value) group
              && safeBindings tail body)

binds :: Ident -> Binder Ann -> Boolean
binds target = case _ of
  BinderNull _ -> false
  BinderVar _ ident -> ident == target
  BinderNamed _ ident binder -> ident == target || binds target binder
  BinderLit _ literal -> any (binds target) literal
  BinderConstructor _ _ _ binders -> any (binds target) binders
