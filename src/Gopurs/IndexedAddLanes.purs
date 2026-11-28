module Gopurs.IndexedAddLanes
  ( eligible
  , extractContribution
  , emit
  , emitContribution
  , emitContributionReverse
  ) where

import Prelude
import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all)
import Data.Maybe (Maybe(..), isJust)
import Data.Tuple (Tuple(..), snd)
import Gopurs.ArrayIndexedFold (readOnlyScalar)
import Gopurs.CallAnalysis (collectCurriedAbs)
import Gopurs.ClosedIndexedFolds (closedPrimitiveCallback)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.GoAst (GoExpr(..), rawGo)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident, Literal(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendOperator2(..), BackendOperatorNum(..), BackendSyntax(..), Level)

-- Existing fusion admits closed primitive callbacks. Recheck that contract,
-- then restrict lanes to a normalizing mapper and an additive accumulator.
-- Neither an Int annotation nor a local-use annotation is a range proof.
eligible :: TcoExpr -> TcoExpr -> Boolean
eligible mapper folder = case unwrapTcoExpr mapper of
  Abs mapArgs mapBody ->
    let mapped = collectCurriedAbs mapArgs mapBody
        normalized = case unwrapTcoExpr mapped.body of
          PrimOp (Op2 (OpIntNum operation) _ _) -> operation == OpAdd || operation == OpSubtract || operation == OpMultiply
          _ -> false
    in normalized && total mapped.body && (NEA.length mapped.args == 1 || NEA.length mapped.args == 2)
      && closedPrimitiveCallback (NEA.length mapped.args) (neutral mapper)
      && isJust (extractContribution folder)
  _ -> false

-- Extract the independent additive contribution of a closed 3-argument folder
-- `\index acc value -> acc + contribution(index, value)` whose contribution
-- does not reference `acc` and is total.
extractContribution :: TcoExpr -> Maybe { index :: Tuple (Maybe Ident) Level, value :: Tuple (Maybe Ident) Level, body :: TcoExpr }
extractContribution folder = case unwrapTcoExpr folder of
  Abs foldArgs foldBody -> do
    let folded = collectCurriedAbs foldArgs foldBody
    guard (closedPrimitiveCallback 3 (neutral folder))
    case NEA.toArray folded.args, unwrapTcoExpr folded.body of
      [ indexArg, Tuple _ accumulator, valueArg ], PrimOp (Op2 (OpIntNum OpAdd) lhs contribution) -> do
        guard case unwrapTcoExpr lhs of
          Local _ level -> level == accumulator
          _ -> false
        guard (absent accumulator contribution && total contribution && readOnlyScalar [ snd indexArg, snd valueArg ] contribution)
        pure { index: indexArg, value: valueArg, body: contribution }
      _, _ -> Nothing
  _ -> Nothing

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
-- Slicing an 8-element chunk eliminates per-element bounds checks, and masking
-- `index &^ 7 | k` exposes the constant low bits of each lane's index to Go's
-- SSA optimizer (folding power-of-two modulos such as `i & 1` per lane).
emit :: { source :: String, prefix :: String, result :: String } -> (GoExpr -> GoExpr -> GoExpr -> GoExpr) -> GoExpr -> GoExpr
emit names step fallback =
  let
    index = names.prefix <> "_index"
    chunk = names.prefix <> "_chunk"
    base = names.prefix <> "_base"
    lane 0 = names.result
    lane n = names.prefix <> "_" <> show n
    chunkElem n = GoIndex (GoVar chunk) (GoInt n)
    tailElem = GoIndex (rawGo ("(" <> names.source <> ")")) (GoVar index)
    update n =
      let i = GoBinOp "|" (GoVar base) (GoInt n)
      in GoMutate (lane n) (step i (GoVar (lane n)) (chunkElem n))
    add left right = GoCall (rawGo "gopurs_runtime.IntAdd") [ left, right ]
    body = map (\n -> GoAssign (lane n) (rawGo "int64(0)")) (Array.range 1 7) <>
      [ GoAssign index (GoInt 0)
      , GoForRange ("; " <> index <> " <= len(" <> names.source <> ")-8; " <> index <> " += 8")
          ([ GoAssign chunk (rawGo ("(" <> names.source <> ")[" <> index <> " : " <> index <> "+8]"))
           , GoAssign base (rawGo (index <> " &^ 7"))
           ] <> map update (Array.range 0 7))
      ] <> map (\n -> GoMutate names.result (add (GoVar names.result) (GoVar (lane n)))) (Array.range 1 7) <>
      [ GoForRange ("; " <> index <> " < len(" <> names.source <> "); " <> index <> "++")
          [ GoMutate names.result (step (GoVar index) (GoVar names.result) tailElem) ] ]
  in GoIfElse (rawGo ("len(" <> names.source <> ") < 8")) [ fallback ] body

-- Accumulate additive contributions in raw int64 across 8 lanes and reduce
-- modulo 2^32 once after the loop. Because 2^32 divides 2^64, summing raw
-- int64 contributions and narrowing once is identical modulo 2^32 to calling
-- IntAdd on every element, while removing 8 sign-extensions per iteration.
emitContribution :: { source :: String, prefix :: String, result :: String, offset :: Maybe GoExpr } -> (GoExpr -> GoExpr -> GoExpr) -> GoExpr -> GoExpr
emitContribution names contribute fallback =
  let
    index = names.prefix <> "_index"
    chunk = names.prefix <> "_chunk"
    base = names.prefix <> "_base"
    lane 0 = names.result
    lane n = names.prefix <> "_" <> show n
    chunkElem n = GoIndex (GoVar chunk) (GoInt n)
    tailElem = GoIndex (rawGo ("(" <> names.source <> ")")) (GoVar index)
    tailIdx = case names.offset of
      Nothing -> GoVar index
      Just off -> GoBinOp "+" off (GoVar index)
    laneStep idxExpr n = GoMutate (lane n) (GoBinOp "+" (GoVar (lane n)) (contribute idxExpr (chunkElem n)))
    zeroLoop = GoForRange ("; " <> index <> " <= len(" <> names.source <> ")-8; " <> index <> " += 8")
      ([ GoAssign chunk (rawGo ("(" <> names.source <> ")[" <> index <> " : " <> index <> "+8]"))
       , GoAssign base (rawGo (index <> " &^ 7"))
       ] <> map (\n -> laneStep (GoBinOp "|" (GoVar base) (GoInt n)) n) (Array.range 0 7))
    offsetLoop off =
      let
        evenIdx n = if n `mod` 2 == 0 then
          GoBinOp "&^" (GoBinOp "+" (GoVar base) (GoInt n)) (GoInt 1)
          else GoBinOp "|" (GoBinOp "+" (GoVar base) (GoInt (n - 1))) (GoInt 1)
        oddIdx n = if n `mod` 2 == 0 then
          GoBinOp "|" (GoBinOp "+" (GoVar base) (GoInt n)) (GoInt 1)
          else GoBinOp "&^" (GoBinOp "+" (GoVar base) (GoInt n)) (GoInt 1)
        evenLoop = GoForRange ("; " <> index <> " <= len(" <> names.source <> ")-8; " <> index <> " += 8")
          ([ GoAssign chunk (rawGo ("(" <> names.source <> ")[" <> index <> " : " <> index <> "+8]"))
           , GoAssign base (GoBinOp "&^" (GoBinOp "+" off (GoVar index)) (GoInt 1))
           ] <> map (\n -> laneStep (evenIdx n) n) (Array.range 0 7))
        oddLoop = GoForRange ("; " <> index <> " <= len(" <> names.source <> ")-8; " <> index <> " += 8")
          ([ GoAssign chunk (rawGo ("(" <> names.source <> ")[" <> index <> " : " <> index <> "+8]"))
           , GoAssign base (GoBinOp "|" (GoBinOp "+" off (GoVar index)) (GoInt 1))
           ] <> map (\n -> laneStep (oddIdx n) n) (Array.range 0 7))
      in GoIfElse (GoBinOp "==" (GoBinOp "&" off (GoInt 1)) (GoInt 0)) [ evenLoop ] [ oddLoop ]
    mainLoop = case names.offset of
      Nothing -> zeroLoop
      Just off -> offsetLoop off
    sumPair a b = GoBinOp "+" a b
    combined = sumPair
      (sumPair (sumPair (GoVar (lane 0)) (GoVar (lane 1))) (sumPair (GoVar (lane 2)) (GoVar (lane 3))))
      (sumPair (sumPair (GoVar (lane 4)) (GoVar (lane 5))) (sumPair (GoVar (lane 6)) (GoVar (lane 7))))
    reduce = GoMutate names.result (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ combined ] ])
    body = map (\n -> GoAssign (lane n) (rawGo "int64(0)")) (Array.range 1 7) <>
      [ GoAssign index (GoInt 0)
      , mainLoop
      , reduce
      , GoForRange ("; " <> index <> " < len(" <> names.source <> "); " <> index <> "++")
          [ GoMutate names.result (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoBinOp "+" (GoVar names.result) (contribute tailIdx tailElem) ] ]) ]
      ]
  in GoIfElse (rawGo ("len(" <> names.source <> ") < 8")) [ fallback ] body

-- Reverse-order counterpart of emitContribution: visits source from len-1 down
-- to 0 with logical indices 0..len-1 across 8 unrolled BCE lanes, reducing
-- modulo 2^32 once after the loop.
emitContributionReverse :: { source :: String, prefix :: String, result :: String } -> (GoExpr -> GoExpr -> GoExpr) -> GoExpr -> GoExpr
emitContributionReverse names contribute fallback =
  let
    index = names.prefix <> "_index"
    pos = names.prefix <> "_pos"
    chunk = names.prefix <> "_chunk"
    base = names.prefix <> "_base"
    lane 0 = names.result
    lane n = names.prefix <> "_" <> show n
    chunkElem n = GoIndex (GoVar chunk) (GoInt (7 - n))
    tailElem = GoIndex (rawGo ("(" <> names.source <> ")")) (GoBinOp "-" (GoBinOp "-" (GoCall (GoVar "len") [ rawGo names.source ]) (GoInt 1)) (GoVar index))
    laneStep n = GoMutate (lane n) (GoBinOp "+" (GoVar (lane n)) (contribute (GoBinOp "|" (GoVar base) (GoInt n)) (chunkElem n)))
    mainLoop = GoForRange ("; " <> pos <> " >= 8; " <> pos <> " -= 8")
      ([ GoAssign chunk (rawGo ("(" <> names.source <> ")[" <> pos <> "-8 : " <> pos <> "]"))
       , GoAssign base (rawGo (index <> " &^ 7"))
       ] <> map laneStep (Array.range 0 7) <> [ GoMutate index (GoBinOp "+" (GoVar index) (GoInt 8)) ])
    sumPair a b = GoBinOp "+" a b
    combined = sumPair
      (sumPair (sumPair (GoVar (lane 0)) (GoVar (lane 1))) (sumPair (GoVar (lane 2)) (GoVar (lane 3))))
      (sumPair (sumPair (GoVar (lane 4)) (GoVar (lane 5))) (sumPair (GoVar (lane 6)) (GoVar (lane 7))))
    reduce = GoMutate names.result (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ combined ] ])
    body = map (\n -> GoAssign (lane n) (rawGo "int64(0)")) (Array.range 1 7) <>
      [ GoAssign index (GoInt 0)
      , GoAssign pos (GoCall (GoVar "len") [ rawGo names.source ])
      , mainLoop
      , reduce
      , GoForRange ("; " <> index <> " < len(" <> names.source <> "); " <> index <> "++")
          [ GoMutate names.result (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoBinOp "+" (GoVar names.result) (contribute (GoVar index) tailElem) ] ]) ]
      ]
  in GoIfElse (rawGo ("len(" <> names.source <> ") < 8")) [ fallback ] body
