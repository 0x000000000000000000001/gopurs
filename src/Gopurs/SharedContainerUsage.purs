module Gopurs.SharedContainerUsage (collectSharedContainerPairs) where

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
collectSharedContainerPairs = foldl collectModule Set.empty
  where
  collectModule pairs (Module mod) =
    foldl (\acc ty -> case ty of
      Just value -> collectType false value acc
      Nothing -> acc)
      (foldl collectBind pairs mod.decls)
      (Map.values mod.foreign)

collectBind :: Pairs -> Bind Ann -> Pairs
collectBind pairs = case _ of
  NonRec binding -> collectBinding pairs binding
  Rec bindings -> foldl collectBinding pairs bindings

collectBinding :: Pairs -> Binding Ann -> Pairs
collectBinding pairs (Binding ann _ body) =
  let value = not (leadingLambda body)
  in collectExpr value body (collectAnn value ann pairs)

leadingLambda :: Expr Ann -> Boolean
leadingLambda (ExprAbs _ _ _) = true
leadingLambda (ExprTypeApp _ body _) = leadingLambda body
leadingLambda _ = false

collectAnn :: Boolean -> Ann -> Pairs -> Pairs
collectAnn value (Ann ann) pairs = case ann.type of
  Just ty -> collectType value ty pairs
  Nothing -> pairs

collectExpr :: Boolean -> Expr Ann -> Pairs -> Pairs
collectExpr value expression pairs =
  let withType = collectAnn value (exprAnn expression) pairs
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

collectAlternative :: Boolean -> Pairs -> CaseAlternative Ann -> Pairs
collectAlternative value pairs (CaseAlternative _ result) = case result of
  Unconditional body -> collectExpr value body pairs
  Guarded guards -> foldl (\acc (Guard condition body) -> collectExpr value body (collectExpr true condition acc)) pairs guards

-- The arrows of an ordinary curried signature are one direct function, not
-- returned callbacks. A function in one of its argument/container slots is a
-- value. Actual returned functions are also observed in their expression use.
collectType :: Boolean -> ExprType -> Pairs -> Pairs
collectType value ty pairs = case ty of
  ForAll _ body -> collectType value body pairs
  ConstrainedType _ body -> collectType value body pairs
  Func _ _ | value -> collectPairs ty pairs
  Func args result -> collectDirectResult result (foldl (flip (collectType true)) pairs args)
  Array element -> collectType true element pairs
  Record row -> collectType true row pairs
  Row fields _ -> foldl (\acc (Tuple _ field) -> collectType true field acc) pairs fields
  _ -> case applied ty [] of
    Just adt -> foldl (flip (collectType true)) pairs adt.args
    Nothing -> pairs
  where
  collectDirectResult (Func args result) acc = collectDirectResult result (foldl (flip (collectType true)) acc args)
  collectDirectResult (ForAll _ body) acc = collectDirectResult body acc
  collectDirectResult (ConstrainedType _ body) acc = collectDirectResult body acc
  collectDirectResult result acc = collectType true result acc

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
