module Gopurs.SharedContainerUsage (collectSharedContainerPairs, foldFunctionValueTypes, foldFunctionTypes) where

import Prelude

import Data.Array as Array
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binding(..), CaseAlternative(..), CaseGuard(..), Expr(..), ExprType(..), Guard(..), Module(..), Prop(..), exprAnn)

type Pairs = Set.Set (Array ExprType)

-- A source higher-order boundary can retain the generic calling ABI after
-- specialization. Share only concrete product families observed there, rather
-- than changing every native list with the same structural shape. This is a
-- representation heuristic: it proves neither escape nor exclusive ownership.
-- Codegen still validates the enum, pointer declarations and complete layout.
collectSharedContainerPairs :: Array (Module Ann) -> Pairs
collectSharedContainerPairs = foldFunctionValueTypes collectPairs Set.empty

-- Visit each annotation observed in a function-value context. Consumers decide
-- which nested type families matter to their own representation policy.
foldFunctionValueTypes :: forall a. (ExprType -> a -> a) -> a -> Array (Module Ann) -> a
foldFunctionValueTypes collect = foldFunctionTypes { value: collect, recursive: \_ result -> result }

-- The same walk also exposes recursive binding signatures, including local
-- groups. Consumers can combine positive boundary evidence with loop guards
-- without walking expressions a second time or reinterpreting usage counts.
foldFunctionTypes
  :: forall a
   . { value :: ExprType -> a -> a, recursive :: ExprType -> a -> a }
  -> a
  -> Array (Module Ann)
  -> a
foldFunctionTypes visitors initial = foldl collectModule initial
  where
  collectModule result (Module mod) =
    foldl (\acc ty -> case ty of
      Just value -> collectType false value acc
      Nothing -> acc)
      (foldl collectBind result mod.decls)
      (Map.values mod.foreign)

  collectBind result = case _ of
    NonRec binding -> collectBinding result binding
    Rec bindings -> foldl collectRecursiveBinding result bindings

  collectRecursiveBinding result binding@(Binding (Ann ann) _ _) =
    let
      withSignature = case ann.type of
        Just ty -> visitors.recursive ty result
        Nothing -> result
    in collectBinding withSignature binding

  collectBinding result (Binding ann _ body) =
    let value = not (leadingLambda body)
    in collectExpr value body (collectAnn value ann result)

  leadingLambda (ExprAbs _ _ _) = true
  leadingLambda (ExprTypeApp _ body _) = leadingLambda body
  leadingLambda _ = false

  collectAnn value (Ann ann) result = case ann.type of
    Just ty -> collectType value ty result
    Nothing -> result

  collectExpr value expression result =
    let withType = collectAnn value (exprAnn expression) result
    in case expression of
      ExprLit _ literal -> foldl (flip (collectExpr true)) withType literal
      ExprApp _ fn argument -> collectExpr true argument (collectExpr false fn withType)
      ExprAbs _ _ body -> collectExpr (not (leadingLambda body)) body withType
      ExprLet _ bindings body -> collectExpr value body (foldl collectBind withType bindings)
      ExprTypeApp _ body _ -> collectExpr value body withType
      ExprCase _ scrutinees alternatives -> foldl (collectAlternative value)
        (foldl (flip (collectExpr true)) withType scrutinees) alternatives
      ExprAccessor _ record _ -> collectExpr true record withType
      ExprUpdate _ record props -> foldl (\acc (Prop _ field) -> collectExpr true field acc)
        (collectExpr true record withType) props
      ExprVar _ _ -> withType
      ExprConstructor _ _ _ _ -> withType

  collectAlternative value result (CaseAlternative _ alternative) = case alternative of
    Unconditional body -> collectExpr value body result
    Guarded guards -> foldl (\acc (Guard condition body) -> collectExpr value body (collectExpr true condition acc)) result guards

  -- Ordinary curried arrows describe one direct function. A function in an
  -- argument/container slot is a value; returned functions are also observed
  -- at their expression uses.
  collectType value ty result = case ty of
    ForAll _ body -> collectType value body result
    ConstrainedType _ body -> collectType value body result
    Func _ _ | value -> visitors.value ty result
    Func args output -> collectDirectResult output (foldl (flip (collectType true)) result args)
    Array element -> collectType true element result
    Record row -> collectType true row result
    Row fields _ -> foldl (\acc (Tuple _ field) -> collectType true field acc) result fields
    _ -> case applied ty [] of
      Just adt -> foldl (flip (collectType true)) result adt.args
      Nothing -> result

  collectDirectResult (Func args output) result = collectDirectResult output (foldl (flip (collectType true)) result args)
  collectDirectResult (ForAll _ body) result = collectDirectResult body result
  collectDirectResult (ConstrainedType _ body) result = collectDirectResult body result
  collectDirectResult output result = collectType true output result

-- Once a function value is observed, retain concrete product arguments nested
-- in its inputs/results. Unknown, open or polymorphic products are excluded.
collectPairs :: ExprType -> Pairs -> Pairs
collectPairs ty pairs = case ty of
  Func args result -> collectPairs result (foldl (flip collectPairs) pairs args)
  ForAll _ body -> collectPairs body pairs
  ConstrainedType _ body -> collectPairs body pairs
  Array element -> collectPairs element pairs
  Record row -> collectPairs row pairs
  Row fields _ -> foldl (\acc (Tuple _ field) -> collectPairs field acc) pairs fields
  _ -> case applied ty [] of
    Just adt ->
      let
        selected = case adt.args of
          [ record, tag ] | adt.name == "Data.Tuple.Tuple" && scalarRecord record -> case applied tag [] of
            Just enum | Array.null enum.args -> Set.insert [ orderedRecord record, ADT enum.name enum.path [] ] pairs
            _ -> pairs
          _ -> pairs
      in foldl (flip collectPairs) selected adt.args
    Nothing -> pairs

scalarRecord :: ExprType -> Boolean
scalarRecord (Record (Row fields Nothing)) = not (Array.null fields)
  && Array.all (\(Tuple _ field) -> scalar field) fields
  where
  scalar Int = true
  scalar Number = true
  scalar Boolean = true
  scalar Char = true
  scalar String = true
  scalar _ = false
scalarRecord _ = false

-- The native record layout also orders labels; row spelling must not produce
-- duplicate decisions for the same product family.
orderedRecord :: ExprType -> ExprType
orderedRecord (Record (Row fields Nothing)) = Record (Row (Array.sortBy (comparing \(Tuple name _) -> name) fields) Nothing)
orderedRecord other = other

applied :: ExprType -> Array ExprType -> Maybe { name :: String, path :: Array String, args :: Array ExprType }
applied (TypeApp fn extra) collected = applied fn (extra <> collected)
applied (ADT name path original) collected = Just { name, path, args: original <> collected }
applied _ _ = Nothing
