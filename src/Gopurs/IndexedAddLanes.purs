module Gopurs.IndexedAddLanes
  ( AffineParity
  , eligible
  , extractContribution
  , extractAffineParity
  , emit
  , emitContribution
  , emitAffineSimdCall
  , emitAffineStep32
  ) where

import Prelude
import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all)
import Data.Int.Bits as Bits
import Data.Map as Map
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
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), Level, Pair(..))

type AffineParity =
  { aEven :: Int
  , bEven :: Int
  , aOdd :: Int
  , bOdd :: Int
  }

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

data AffineVal
  = ConstInt Int
  | Affine Int Int
  | IndexVar

smallMul :: Int -> Boolean
smallMul c = c >= -2097152 && c <= 2097152

-- Symbolically evaluate a closed primitive expression over Z/2^32Z assuming
-- its value input lies in signed Int32 (guaranteed for TypeInt32Array elements
-- and for outputs of Int32-normalizing operations). Every OpMultiply requires
-- at least one factor in [-2^21, 2^21] so |a * b| <= 2^52 < 2^53 and JS
-- float64 rounding cannot differ from exact modulo-2^32 multiplication.
extractAffineParity :: Maybe TcoExpr -> TcoExpr -> Maybe AffineParity
extractAffineParity mbMapper folder = do
  contrib <- extractContribution folder
  let
    idxLevel = Just (snd contrib.index)
    valLevel = snd contrib.value
  Tuple evenVal oddVal <- case mbMapper of
    Nothing -> do
      c0 <- evalAffine (Map.singleton valLevel (Affine 1 0)) idxLevel 0 contrib.body
      c1 <- evalAffine (Map.singleton valLevel (Affine 1 0)) idxLevel 1 contrib.body
      pure (Tuple c0 c1)
    Just mapper -> case unwrapTcoExpr mapper of
      Abs mapArgs mapBody -> do
        let mapped = collectCurriedAbs mapArgs mapBody
            arity = NEA.length mapped.args
        guard ((arity == 1 || arity == 2) && total mapped.body && closedPrimitiveCallback arity (neutral mapper))
        Tuple mIdx mVal <- case NEA.toArray mapped.args of
          [ Tuple _ v ] -> Just (Tuple Nothing v)
          [ Tuple _ i, Tuple _ v ] -> Just (Tuple (Just i) v)
          _ -> Nothing
        m0 <- evalAffine (Map.singleton mVal (Affine 1 0)) mIdx 0 mapped.body
        m1 <- evalAffine (Map.singleton mVal (Affine 1 0)) mIdx 1 mapped.body
        guard (isScalarAffine m0 && isScalarAffine m1)
        c0 <- evalAffine (Map.singleton valLevel m0) idxLevel 0 contrib.body
        c1 <- evalAffine (Map.singleton valLevel m1) idxLevel 1 contrib.body
        pure (Tuple c0 c1)
      _ -> Nothing
  Tuple aEven bEven <- toPair evenVal
  Tuple aOdd bOdd <- toPair oddVal
  pure { aEven, bEven, aOdd, bOdd }
  where
  isScalarAffine = case _ of
    ConstInt _ -> true
    Affine _ _ -> true
    IndexVar -> false

  toPair = case _ of
    ConstInt b -> Just (Tuple 0 b)
    Affine a b -> Just (Tuple a b)
    IndexVar -> Nothing

  evalAffine :: Map.Map Level AffineVal -> Maybe Level -> Int -> TcoExpr -> Maybe AffineVal
  evalAffine env idxLevel parity expr = case unwrapTcoExpr expr of
    Lit (LitInt n) -> Just (ConstInt n)
    Local _ level ->
      if Just level == idxLevel then Just IndexVar
      else Map.lookup level env
    Let _ level rhs body -> do
      val <- evalAffine env idxLevel parity rhs
      evalAffine (Map.insert level val env) idxLevel parity body
    PrimOp (Op1 OpIntNegate inner) -> case evalAffine env idxLevel parity inner of
      Just (ConstInt c) -> Just (ConstInt (-c))
      Just (Affine a b) -> Just (Affine (-a) (-b))
      _ -> Nothing
    PrimOp (Op2 op left right) -> do
      lv <- evalAffine env idxLevel parity left
      rv <- evalAffine env idxLevel parity right
      case op of
        OpIntNum OpAdd -> case lv, rv of
          ConstInt c1, ConstInt c2 -> Just (ConstInt (c1 + c2))
          Affine a b, ConstInt c -> Just (Affine a (b + c))
          ConstInt c, Affine a b -> Just (Affine a (c + b))
          Affine a1 b1, Affine a2 b2 -> Just (Affine (a1 + a2) (b1 + b2))
          _, _ -> Nothing
        OpIntNum OpSubtract -> case lv, rv of
          ConstInt c1, ConstInt c2 -> Just (ConstInt (c1 - c2))
          Affine a b, ConstInt c -> Just (Affine a (b - c))
          ConstInt c, Affine a b -> Just (Affine (-a) (c - b))
          Affine a1 b1, Affine a2 b2 -> Just (Affine (a1 - a2) (b1 - b2))
          _, _ -> Nothing
        OpIntNum OpMultiply -> case lv, rv of
          ConstInt c1, ConstInt c2 | smallMul c1 || smallMul c2 -> Just (ConstInt (c1 * c2))
          ConstInt c, Affine a b | smallMul c -> Just (Affine (c * a) (c * b))
          Affine a b, ConstInt c | smallMul c -> Just (Affine (c * a) (c * b))
          _, _ -> Nothing
        OpIntNum OpMod -> case lv, rv of
          IndexVar, ConstInt 2 -> Just (ConstInt parity)
          ConstInt c1, ConstInt c2 | c2 > 0 -> Just (ConstInt (c1 `mod` c2))
          _, _ -> Nothing
        OpIntBitAnd -> case lv, rv of
          IndexVar, ConstInt 1 -> Just (ConstInt parity)
          ConstInt 1, IndexVar -> Just (ConstInt parity)
          ConstInt c1, ConstInt c2 -> Just (ConstInt (Bits.and c1 c2))
          _, _ -> Nothing
        OpIntBitOr -> case lv, rv of
          ConstInt c1, ConstInt c2 -> Just (ConstInt (Bits.or c1 c2))
          _, _ -> Nothing
        OpIntBitXor -> case lv, rv of
          ConstInt c1, ConstInt c2 -> Just (ConstInt (Bits.xor c1 c2))
          _, _ -> Nothing
        OpIntBitShiftLeft -> case lv, rv of
          ConstInt c, ConstInt k | k >= 0 && k <= 31 -> Just (ConstInt (Bits.shl c k))
          Affine a b, ConstInt k | k >= 0 && k <= 31 -> Just (Affine (Bits.shl a k) (Bits.shl b k))
          _, _ -> Nothing
        _ -> Nothing
    Branch branches otherwise ->
      evalBranches env idxLevel parity (NEA.toArray branches) otherwise
    _ -> Nothing

  evalBranches env idxLevel parity branches fallback = case Array.uncons branches of
    Nothing -> evalAffine env idxLevel parity fallback
    Just { head: Pair cond branchResult, tail } -> do
      taken <- evalBool env idxLevel parity cond
      if taken then evalAffine env idxLevel parity branchResult
      else evalBranches env idxLevel parity tail fallback

  evalBool env idxLevel parity expr = case unwrapTcoExpr expr of
    Lit (LitBoolean b) -> Just b
    PrimOp (Op1 OpBooleanNot inner) -> not <$> evalBool env idxLevel parity inner
    PrimOp (Op2 OpBooleanAnd left right) -> do
      l <- evalBool env idxLevel parity left
      if not l then Just false else evalBool env idxLevel parity right
    PrimOp (Op2 OpBooleanOr left right) -> do
      l <- evalBool env idxLevel parity left
      if l then Just true else evalBool env idxLevel parity right
    PrimOp (Op2 (OpIntOrd ordOp) left right) -> do
      lv <- evalAffine env idxLevel parity left
      rv <- evalAffine env idxLevel parity right
      case lv, rv of
        ConstInt a, ConstInt b -> Just case ordOp of
          OpEq -> a == b
          OpNotEq -> a /= b
          OpLt -> a < b
          OpLte -> a <= b
          OpGt -> a > b
          OpGte -> a >= b
        _, _ -> Nothing
    _ -> Nothing

emitAffineSimdCall :: String -> String -> AffineParity -> Boolean -> Array GoExpr
emitAffineSimdCall target resultName aff swapped =
  let
    aE = if swapped then aff.aOdd else aff.aEven
    bE = if swapped then aff.bOdd else aff.bEven
    aO = if swapped then aff.aEven else aff.aOdd
    bO = if swapped then aff.bEven else aff.bOdd
    call =
      if aE == 1 && bE == 0 && aO == 1 && bO == 0 then
        GoCall (GoSelector (GoVar "gopurs_runtime") "SumInt32") [ rawGo target ]
      else
        GoCall (GoSelector (GoVar "gopurs_runtime") "AffineParitySumInt32")
          [ rawGo target
          , GoCall (GoVar "int32") [ GoInt aE ]
          , GoCall (GoVar "int32") [ GoInt bE ]
          , GoCall (GoVar "int32") [ GoInt aO ]
          , GoCall (GoVar "int32") [ GoInt bO ]
          ]
  in
    [ GoIfElse (GoBinOp ">" (GoCall (GoVar "len") [ rawGo target ]) (GoInt 0))
        [ GoMutate resultName
            (GoCall (GoVar "int64")
              [ GoBinOp "+" (GoCall (GoVar "int32") [ GoVar resultName ]) call ])
        ] []
    ]

emitAffineStep32 :: String -> GoExpr -> GoExpr -> AffineParity -> GoExpr
emitAffineStep32 acc32Name idxExpr val32Expr aff =
  let
    int32Lit n = GoCall (GoVar "int32") [ GoInt n ]
    term a b = case a, b of
      0, _ -> int32Lit b
      1, 0 -> val32Expr
      1, _ -> GoBinOp "+" val32Expr (int32Lit b)
      (-1), 0 -> GoPrefixOp "-" val32Expr
      _, 0 -> GoBinOp "*" (int32Lit a) val32Expr
      _, _ -> GoBinOp "+" (GoBinOp "*" (int32Lit a) val32Expr) (int32Lit b)
    addTo expr = GoMutate acc32Name (GoBinOp "+" (GoVar acc32Name) expr)
  in
    if aff.aEven == aff.aOdd && aff.bEven == aff.bOdd then
      addTo (term aff.aEven aff.bEven)
    else
      GoIfElse (GoBinOp "==" (GoBinOp "&" idxExpr (GoInt 1)) (GoInt 0))
        [ addTo (term aff.aEven aff.bEven) ]
        [ addTo (term aff.aOdd aff.bOdd) ]
