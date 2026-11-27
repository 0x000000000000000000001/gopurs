module Gopurs.IndexedAddLanes (eligible, emit) where

import Prelude
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIndexedFold (readOnlyScalar)
import Gopurs.CallAnalysis (collectCurriedAbs)
import Gopurs.ClosedIndexedFolds (closedPrimitiveCallback)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.GoAst (GoExpr(..), rawGo)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Literal(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendOperator(..), BackendOperator2(..), BackendOperatorNum(..), Level)

-- Existing fusion admits closed primitive callbacks. Recheck that contract,
-- then restrict lanes to a normalizing mapper and an additive accumulator.
-- Neither an Int annotation nor a local-use annotation is a range proof.
eligible :: TcoExpr -> TcoExpr -> Boolean
eligible mapper folder = case unwrapTcoExpr mapper, unwrapTcoExpr folder of
  Abs mapArgs mapBody, Abs foldArgs foldBody ->
    let mapped = collectCurriedAbs mapArgs mapBody
        folded = collectCurriedAbs foldArgs foldBody
        normalized = case unwrapTcoExpr mapped.body of
          PrimOp (Op2 (OpIntNum operation) _ _) -> operation == OpAdd || operation == OpSubtract || operation == OpMultiply
          _ -> false
    in normalized && total mapped.body && (NEA.length mapped.args == 1 || NEA.length mapped.args == 2)
      && closedPrimitiveCallback (NEA.length mapped.args) (neutral mapper)
      && closedPrimitiveCallback 3 (neutral folder)
      && case NEA.toArray folded.args, unwrapTcoExpr folded.body of
        [ Tuple _ index, Tuple _ accumulator, Tuple _ value ], PrimOp (Op2 (OpIntNum OpAdd) lhs contribution) ->
          (case unwrapTcoExpr lhs of
            Local _ level -> level == accumulator
            _ -> false)
            && absent accumulator contribution && total contribution && readOnlyScalar [ index, value ] contribution
        _, _ -> false
  _, _ -> false

absent :: Level -> TcoExpr -> Boolean
absent excluded (TcoExpr _ syntax) = case syntax of
  Local _ level -> level /= excluded
  _ -> all (absent excluded) syntax

-- Keep the proof deliberately small: division and dynamic modulo stay on the
-- original loop. Positive literal modulo is total and includes indexed masks.
total :: TcoExpr -> Boolean
total (TcoExpr _ syntax) = all total syntax && case syntax of
  PrimOp (Op2 (OpIntNum OpDivide) _ _) -> false
  PrimOp (Op2 (OpIntNum OpMod) _ divisor) -> case unwrapTcoExpr divisor of
    Lit (LitInt value) -> value > 0
    _ -> false
  _ -> true

neutral :: TcoExpr -> NeutralExpr
neutral (TcoExpr _ syntax) = NeutralExpr (map neutral syntax)

-- Reads and callback invocations retain their original order. Only independent
-- IntAdd accumulators differ; their final combination is modulo 2^32. The exact
-- original loop handles lengths below eight, including an unnormalized seed.
emit :: { source :: String, prefix :: String, result :: String } -> (GoExpr -> GoExpr -> GoExpr -> GoExpr) -> GoExpr -> GoExpr
emit names step fallback =
  let
    index = names.prefix <> "_index"
    lane 0 = names.result
    lane n = names.prefix <> "_" <> show n
    element i = GoIndex (rawGo ("(" <> names.source <> ")")) i
    update n =
      let i = GoBinOp "+" (GoVar index) (GoInt n)
      in GoMutate (lane n) (step i (GoVar (lane n)) (element i))
    add left right = GoCall (rawGo "gopurs_runtime.IntAdd") [ left, right ]
    body = map (\n -> GoAssign (lane n) (rawGo "int64(0)")) (Array.range 1 7) <>
      [ GoAssign index (GoInt 0)
      , GoForRange ("; " <> index <> " <= len(" <> names.source <> ")-8; " <> index <> " += 8") (map update (Array.range 0 7))
      ] <> map (\n -> GoMutate names.result (add (GoVar names.result) (GoVar (lane n)))) (Array.range 1 7) <>
      [ GoForRange ("; " <> index <> " < len(" <> names.source <> "); " <> index <> "++")
          [ GoMutate names.result (step (GoVar index) (GoVar names.result) (element (GoVar index))) ] ]
  in GoIfElse (rawGo ("len(" <> names.source <> ") < 8")) [ fallback ] body
