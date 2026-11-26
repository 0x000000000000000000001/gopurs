module Gopurs.ClosedIndexedFolds (optimize) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, foldM)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIndexedFold (integerCallback, readOnlyScalar)
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
