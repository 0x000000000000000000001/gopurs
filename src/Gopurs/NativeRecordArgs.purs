module Gopurs.NativeRecordArgs
  ( projectedArgument
  , candidateToShare
  , workerArguments
  ) where

import Prelude

import Data.Array as Array
import Data.Foldable (all, any)
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..), fst)
import Gopurs.GoAst (GoType(..))
import Gopurs.GoTypes (visibleRecordFields)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binder(..), Binding(..), CaseAlternative(..), CaseGuard(..), Expr(..), ExprType(..), Guard(..), Ident, Qualified(..), propValue)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Monomorphize (getExprAnn)
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendSyntax(Abs, Accessor, EffectDefer, Local, Typed, UncurriedAbs, UncurriedEffectAbs))

-- This is a projection for a proven read-only parameter, not a representation
-- for arbitrary open records. Its tail must never be returned or reconstructed.
projectedArgument :: (ExprType -> GoType) -> ExprType -> Maybe GoType
projectedArgument toGoType ty = map
  (TypeRecord <<< map (\(Tuple label fieldType) -> Tuple label (toGoType fieldType)))
  (projectableFields ty)

projectableFields :: ExprType -> Maybe (Array (Tuple String ExprType))
projectableFields (Record (Row fields (Just (TypeVar _)))) =
  let visible = visibleRecordFields fields
  in if not (Array.null visible) && all (\(Tuple _ ty) -> isScalar ty) visible then
    Just (Array.sortBy (comparing fst) visible)
  else Nothing
projectableFields _ = Nothing

isScalar :: ExprType -> Boolean
isScalar = case _ of
  Int -> true
  Number -> true
  String -> true
  Char -> true
  Boolean -> true
  _ -> false

-- Excluding a binding from monomorphisation requires a source-level proof.
-- Code generation independently rechecks uses on the transformed TCO body.
-- Explicit Any annotations remain unknown; they never fall through to another
-- annotation that might happen to look more specific.
candidateToShare :: Binding Ann -> Boolean
candidateToShare (Binding (Ann ann) _ expr) =
  let
    Ann exprAnn = getExprAnn expr
    signature = case ann.type of
      Just ty -> Just ty
      Nothing -> exprAnn.type
  in case signature >>= sharedSignature [] of
    Just { args, result, quantified } ->
      let
        lambdas = collectArguments [] expr
        tails = Array.mapMaybe rowVariable args
        validType ty = case projectableFields ty of
          Just _ -> true
          Nothing -> isMonomorphic ty
        validParameter (Tuple ident ty) = case projectableFields ty of
          Just fields -> safeSourceUse true ident (map fst fields) lambdas.body
          Nothing -> true
      in
        isScalar result
          && not (Array.null tails)
          && all (flip Array.elem tails) quantified
          && all validType args
          && Array.length lambdas.args == Array.length args
          && Array.length (Array.nub lambdas.args) == Array.length lambdas.args
          && all validParameter (Array.zip lambdas.args args)
    Nothing -> false

sharedSignature :: Array String -> ExprType -> Maybe { args :: Array ExprType, result :: ExprType, quantified :: Array String }
sharedSignature quantified (ForAll vars body) = sharedSignature (quantified <> vars) body
sharedSignature quantified (Func args result) = Just { args, result, quantified }
sharedSignature _ _ = Nothing

rowVariable :: ExprType -> Maybe String
rowVariable (Record (Row _ (Just (TypeVar variable)))) = Just variable
rowVariable _ = Nothing

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

-- allowRead becomes false under closures: even a getter there would capture
-- the parameter. Lexical shadowing is respected for lambdas, lets, and cases.
safeSourceUse :: Boolean -> Ident -> Array String -> Expr Ann -> Boolean
safeSourceUse allowRead target fields = case _ of
  ExprVar _ (Qualified Nothing ident) -> ident /= target
  ExprVar _ _ -> true
  ExprLit _ literal -> all (safeSourceUse allowRead target fields) literal
  ExprConstructor _ _ _ _ -> true
  ExprAccessor _ (ExprVar _ (Qualified Nothing ident)) field | ident == target ->
    allowRead && Array.elem field fields
  ExprAccessor _ object _ -> safeSourceUse allowRead target fields object
  ExprUpdate _ object props ->
    safeSourceUse allowRead target fields object
      && all (safeSourceUse allowRead target fields <<< propValue) props
  ExprAbs _ ident body -> ident == target || safeSourceUse false target fields body
  ExprApp _ fn arg -> safeSourceUse allowRead target fields fn && safeSourceUse allowRead target fields arg
  ExprCase _ values branches ->
    all (safeSourceUse allowRead target fields) values
      && all safeAlternative branches
  ExprLet _ bindings body -> safeBindings bindings body
  ExprTypeApp _ inner _ -> safeSourceUse allowRead target fields inner
  where
  safeAlternative (CaseAlternative binders result) =
    any (binds target) binders || case result of
      Unconditional body -> safeSourceUse allowRead target fields body
      Guarded guards -> all (\(Guard condition body) ->
        safeSourceUse allowRead target fields condition && safeSourceUse allowRead target fields body) guards

  safeBindings bindings body = case Array.uncons bindings of
    Nothing -> safeSourceUse allowRead target fields body
    Just { head: NonRec (Binding _ ident value), tail } ->
      safeSourceUse allowRead target fields value
        && (ident == target || safeBindings tail body)
    Just { head: Rec group, tail } ->
      any (\(Binding _ ident _) -> ident == target) group
        || (all (\(Binding _ _ value) -> safeSourceUse allowRead target fields value) group
              && safeBindings tail body)

binds :: Ident -> Binder Ann -> Boolean
binds target = case _ of
  BinderNull _ -> false
  BinderVar _ ident -> ident == target
  BinderNamed _ ident binder -> ident == target || binds target binder
  BinderLit _ literal -> any (binds target) literal
  BinderConstructor _ _ _ binders -> any (binds target) binders

-- A failed or stale source proof cannot enable projection: every affected
-- parameter is rechecked on final IR, and the result must remain scalar.
-- Partial and dynamic calls continue to use the existing boxed wrapper.
workerArguments :: (ExprType -> GoType) -> Array String -> TcoExpr -> Array ExprType -> ExprType -> Array GoType
workerArguments toGoType names body types result =
  if isScalar result then Array.mapWithIndex choose types else map toGoType types
  where
  choose index ty = fromMaybe (toGoType ty) do
    name <- Array.index names index
    fields <- projectableFields ty
    projected <- projectedArgument toGoType ty
    if safeTcoUse name (map fst fields) body then Just projected else Nothing

isTarget :: String -> TcoExpr -> Boolean
isTarget target (TcoExpr _ syntax) = case syntax of
  Typed _ inner -> isTarget target inner
  Local ident level -> localId ident level == target
  _ -> false

containsTarget :: String -> TcoExpr -> Boolean
containsTarget target expr@(TcoExpr _ syntax) =
  isTarget target expr || any (containsTarget target) syntax

safeTcoUse :: String -> Array String -> TcoExpr -> Boolean
safeTcoUse target fields (TcoExpr _ syntax) = case syntax of
  Local ident level -> localId ident level /= target
  Accessor object (GetProp field) | isTarget target object -> Array.elem field fields
  Abs _ body -> not (containsTarget target body)
  UncurriedAbs _ body -> not (containsTarget target body)
  UncurriedEffectAbs _ body -> not (containsTarget target body)
  EffectDefer body -> not (containsTarget target body)
  _ -> all (safeTcoUse target fields) syntax
