module Gopurs.ClosedIndexedFolds (optimize, describe, validDescriptor, closedPrimitiveCallback) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, foldM)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Tuple (Tuple(..), snd)
import Gopurs.ArrayIndexedFold (integerCallback, readOnlyScalar, scalarType)
import Gopurs.CodegenState (ArrayFoldDescriptor)
import Gopurs.CallAnalysis (collectCurriedAbs)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ImmediateApplications.Scope (maximumLevel)
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(Any, Array, Func, Int, ADT, TypeVar), Ident(..), Literal(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level(..))

-- A closed primitive callback and a literal seed have no observable capture
-- work. Expose the remaining array parameter before native worker signatures
-- are computed, using the existing cross-module worker/cache protocol.
optimize :: BackendModule -> BackendModule
optimize mod = mod { bindings = map (\group -> if group.recursive then group else
  group { bindings = map (\(Tuple name value) -> Tuple name (fromMaybe value (expand value))) group.bindings }) mod.bindings }

expand :: NeutralExpr -> Maybe NeutralExpr
expand original = do
  guard (compatible (Func [ Array Int ] Int) original)
  call <- case strip original of
    App function arguments -> case NEA.toArray arguments of
      [ callback, seed ] -> Just { function, callback, seed }
      _ -> Nothing
    _ -> Nothing
  guard (case strip call.function of
    Var (Qualified (Just (ModuleName "Data.FoldableWithIndex")) (Ident "foldlWithIndexImpl")) -> true
    _ -> false)
  guard (compatible Int call.seed)
  guard (case strip call.seed of
    Lit (LitInt _) -> true
    _ -> false)
  guard (closed [] call.callback)
  let callback = Tco.analyze [] call.callback
  guard (integerCallback callback)
  case unwrapTcoExpr callback of
    Abs parameters body -> do
      let grouped = collectCurriedAbs parameters body
      guard (NEA.length grouped.args == 3)
      guard (readOnlyScalar (map (\(Tuple _ level) -> level) (NEA.toArray grouped.args)) grouped.body)
      let
        level = Level (maximumLevel original + 1)
        name = Just (Ident "__gopurs_indexed_input")
        input = NeutralExpr (Typed (Array Int) (NeutralExpr (Local name level)))
        applied = NeutralExpr (Typed Int (NeutralExpr (App call.function (NEA.cons' call.callback [ call.seed, input ]))))
      pure (NeutralExpr (Typed (Func [ Array Int ] Int) (NeutralExpr (Abs (NEA.singleton (Tuple name level)) applied))))
    _ -> Nothing

-- PBO retains both the instantiated annotation and inner generic annotations.
-- Check that each has a consistent specialization, including Prim.Array's
-- equivalent TAST spelling; never erase a concrete contradiction.
compatible :: ExprType -> NeutralExpr -> Boolean
compatible expected (NeutralExpr syntax) = case syntax of
  Typed actual value -> isJust (refine expected actual Map.empty) && compatible expected value
  TypeApp value _ -> compatible expected value
  _ -> true

refine :: ExprType -> ExprType -> Map.Map String ExprType -> Maybe (Map.Map String ExprType)
refine expected actual variables = case actual, expected of
  Any, _ -> Just variables
  TypeVar name, _ -> case Map.lookup name variables of
    Just previous | previous /= expected -> Nothing
    _ -> Just (Map.insert name expected variables)
  Array element, Array wanted -> refine wanted element variables
  ADT "Prim.Array" [ "Prim", "Array" ] [ element ], Array wanted -> refine wanted element variables
  Func arguments result, Func wantedArguments wantedResult -> do
    guard (Array.length arguments == Array.length wantedArguments)
    foldM (\known (Tuple wanted supplied) -> refine wanted supplied known) variables
      (Array.snoc (Array.zip wantedArguments arguments) (Tuple wantedResult result))
  _, _ | actual == expected -> Just variables
  _, _ -> Nothing

closed :: Array Level -> NeutralExpr -> Boolean
closed locals expression = case strip expression of
  Local _ level -> Array.elem level locals
  Var _ -> false
  Abs parameters body -> closed (locals <> map (\(Tuple _ level) -> level) (NEA.toArray parameters)) body
  Let _ level value body -> closed locals value && closed (Array.snoc locals level) body
  LetRec _ _ _ -> false
  syntax -> all (closed locals) syntax

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syntax) = case syntax of
  Typed _ expression -> strip expression
  TypeApp expression _ -> strip expression
  _ -> syntax

-- Export only the same closed primitive fold accepted by the native emitter.
-- Local names remain inside the callback lambda; the consumer captures no locals.
describe :: NeutralExpr -> Maybe ArrayFoldDescriptor
describe original = do
  guard (compatible (Func [ Array Int ] Int) original)
  Tuple input body <- case strip original of
    Abs parameters body | NEA.length parameters == 1 -> Just (Tuple (snd (NEA.head parameters)) body)
    _ -> Nothing
  guard (compatible Int body)
  Tuple callback seed <- case strip body of
    App function arguments -> do
      guard (case strip function of
        Var (Qualified (Just (ModuleName "Data.FoldableWithIndex")) (Ident "foldlWithIndexImpl")) -> true
        _ -> false)
      case NEA.toArray arguments of
        [ callback, seed, values ] | isInput input values -> do
          guard (compatible Int seed && compatible (Array Int) values)
          case strip seed of
            Lit (LitInt initial) -> Just (Tuple callback initial)
            _ -> Nothing
        _ -> Nothing
    _ -> Nothing
  let descriptor = { callback, seed }
  guard (validDescriptor descriptor)
  pure descriptor

validDescriptor :: ArrayFoldDescriptor -> Boolean
validDescriptor descriptor = integerCallback (Tco.analyze [] descriptor.callback) && closedPrimitiveCallback 3 descriptor.callback

-- Do not let an outer instantiated function type hide a contradictory suffix
-- annotation. Likewise inspect every nested Typed layer in primitive bodies.
closedPrimitiveCallback :: Int -> NeutralExpr -> Boolean
closedPrimitiveCallback arity callback = closed [] callback && lambdaTypes arity callback && case strip callback of
  Abs parameters body ->
    let
      Tuple allParameters scalarBody = collectParameters (NEA.toArray parameters) body
      levels = map (\(Tuple _ level) -> level) allParameters
    in Array.length allParameters == arity && scalarConsistent (Map.fromFoldable (map (\level -> Tuple level Int) levels)) scalarBody
      && readOnlyScalar levels (Tco.analyze [] scalarBody)
  _ -> false

lambdaTypes :: Int -> NeutralExpr -> Boolean
lambdaTypes remaining value = compatible (Func (Array.replicate remaining Int) Int) value && case strip value of
  Abs parameters body ->
    let count = NEA.length parameters in
    if count == remaining then compatible Int body
    else count < remaining && lambdaTypes (remaining - count) body
  _ -> false

scalarConsistent :: Map.Map Level ExprType -> NeutralExpr -> Boolean
scalarConsistent locals value@(NeutralExpr syntax) = isJust (scalarType locals (Tco.analyze [] value)) && case syntax of
  Let _ level initial body -> case scalarType locals (Tco.analyze [] initial) of
    Just ty -> scalarConsistent locals initial && scalarConsistent (Map.insert level ty locals) body
    Nothing -> false
  _ -> all (scalarConsistent locals) syntax

isInput :: Level -> NeutralExpr -> Boolean
isInput expected value = case strip value of
  Local _ level -> level == expected
  _ -> false

collectParameters :: Array (Tuple (Maybe Ident) Level) -> NeutralExpr -> Tuple (Array (Tuple (Maybe Ident) Level)) NeutralExpr
collectParameters parameters body = case strip body of
  Abs more result -> collectParameters (parameters <> NEA.toArray more) result
  _ -> Tuple parameters body
