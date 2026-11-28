module Gopurs.IndexedAddLanes
  ( eligible
  , extractContribution
  , emit
  , emitContribution
  , emitContributionReverse
  , emitContribution32
  , emitContributionReverse32
  , inlineInt32Expr
  , countLocalUses
  ) where

import Prelude
import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, sum)
import Data.Int.Bits as Bits
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), isJust)
import Data.Tuple (Tuple(..), snd)
import Gopurs.ArrayIndexedFold (readOnlyScalar)
import Gopurs.CallAnalysis (collectCurriedAbs)
import Gopurs.ClosedIndexedFolds (closedPrimitiveCallback)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.GoAst (GoExpr(..), rawGo)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident, Literal(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendSyntax(..), Level)
import PureScript.Backend.Optimizer.Syntax as Syn

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

countLocalUses :: Level -> TcoExpr -> Int
countLocalUses target (TcoExpr _ syntax) = case syntax of
  Local _ level -> if level == target then 1 else 0
  _ -> sum (map (countLocalUses target) syntax)

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

-- Translate a pure, total closed scalar Int expression directly into an inline
-- Go `int32` expression, without closure calls (avoiding Go's InlMark NOOPs
-- and redundant int64 sign-extensions). Every OpMultiply requires a static
-- IEEE-754 exactness proof (leftBits + rightBits <= 53) using only proven
-- input magnitudes in `mags`. Constant-only subexpressions are folded in 32-bit
-- two's-complement Int arithmetic so Go never hits compile-time overflow.
inlineInt32Expr :: Map Level Int -> Map Level GoExpr -> TcoExpr -> Maybe GoExpr
inlineInt32Expr = go
  where
  lit32 n = GoCall (GoVar "int32") [ GoInt n ]
  unlit32 = case _ of
    GoCall (GoVar "int32") [ GoInt n ] -> Just n
    _ -> Nothing
  intLiteral (TcoExpr _ syntax) = case syntax of
    Typed Int inner -> intLiteral inner
    Typed Any inner -> intLiteral inner
    Syn.TypeApp inner _ -> intLiteral inner
    Lit (LitInt value) -> Just value
    _ -> Nothing
  litMag value
    | value < bottom || value > top = Nothing
    | value == bottom = Just 31
    | otherwise = Just (bits (if value < 0 then negate value else value) 0)
    where
    bits remaining count
      | remaining <= 1 = count
      | otherwise = bits (div remaining 2 + mod remaining 2) (count + 1)
  exprMag mags (TcoExpr _ syntax) = case syntax of
    Typed Int inner -> exprMag mags inner
    Typed Any inner -> exprMag mags inner
    Syn.TypeApp inner _ -> exprMag mags inner
    Lit (LitInt value) -> litMag value
    Local _ level -> Map.lookup level mags
    Let _ level rhs body ->
      let mags' = case exprMag mags rhs of
            Just m -> Map.insert level m mags
            Nothing -> mags
      in exprMag mags' body
    PrimOp (Op1 OpIntNegate _) -> Just 31
    PrimOp (Op1 OpIntBitNot _) -> Just 31
    PrimOp (Op2 operation left right) -> case operation of
      OpIntBitAnd -> case intLiteral left, intLiteral right of
        Just mask, _ | mask >= 0 -> litMag mask
        _, Just mask | mask >= 0 -> litMag mask
        _, _ -> Just 31
      OpIntBitShiftRight -> case intLiteral right of
        Just shift | shift >= 0 && shift <= 31 -> Just (31 - shift)
        _ -> Just 31
      OpIntBitOr -> Just 31
      OpIntBitXor -> Just 31
      OpIntBitShiftLeft -> Just 31
      OpIntNum OpAdd -> case exprMag mags left, exprMag mags right of
        Just a, Just b -> Just (min 31 (max a b + 1))
        _, _ -> Just 31
      OpIntNum OpSubtract -> case exprMag mags left, exprMag mags right of
        Just a, Just b -> Just (min 31 (max a b + 1))
        _, _ -> Just 31
      OpIntNum OpMultiply -> Just 31
      OpIntNum OpMod -> case intLiteral right of
        Just modulus | modulus > 0 && modulus <= 2147483647 -> litMag (modulus - 1)
        _ -> Nothing
      _ -> Nothing
    _ -> Nothing
  binFold op foldFn a b = case unlit32 a, unlit32 b of
    Just x, Just y -> lit32 (foldFn x y)
    _, _ -> GoBinOp op a b
  go mags env (TcoExpr _ syntax) = case syntax of
    Typed Int inner -> go mags env inner
    Typed Any inner -> go mags env inner
    Syn.TypeApp inner _ -> go mags env inner
    Lit (LitInt v) | v >= bottom && v <= top -> Just (lit32 v)
    Local _ level -> Map.lookup level env
    Let _ level rhs body | countLocalUses level body <= 1 -> do
      rhsGo <- go mags env rhs
      let mags' = case exprMag mags rhs of
            Just m -> Map.insert level m mags
            Nothing -> mags
      go mags' (Map.insert level rhsGo env) body
    PrimOp (Op1 OpIntNegate a) -> do
      x <- go mags env a
      pure case unlit32 x of
        Just n -> lit32 (negate n)
        Nothing -> GoPrefixOp "-" x
    PrimOp (Op1 OpIntBitNot a) -> do
      x <- go mags env a
      pure case unlit32 x of
        Just n -> lit32 (Bits.complement n)
        Nothing -> GoPrefixOp "^" x
    PrimOp (Op2 op a b) -> case op of
      OpIntNum OpAdd -> binFold "+" (+) <$> go mags env a <*> go mags env b
      OpIntNum OpSubtract -> binFold "-" (-) <$> go mags env a <*> go mags env b
      OpIntNum OpMultiply -> do
        mA <- exprMag mags a
        mB <- exprMag mags b
        guard (mA + mB <= 53)
        binFold "*" (*) <$> go mags env a <*> go mags env b
      OpIntNum OpMod -> do
        d <- intLiteral b
        guard (d > 0 && d <= 2147483647)
        x <- go mags env a
        if Bits.and d (d - 1) == 0 then
          pure (binFold "&" Bits.and x (lit32 (d - 1)))
        else do
          mA <- exprMag mags a
          guard (mA <= 31)
          pure (GoCall (GoSelector (GoVar "gopurs_runtime") "IntModPos32") [ x, lit32 d ])
      OpIntBitAnd -> binFold "&" Bits.and <$> go mags env a <*> go mags env b
      OpIntBitOr -> binFold "|" Bits.or <$> go mags env a <*> go mags env b
      OpIntBitXor -> binFold "^" Bits.xor <$> go mags env a <*> go mags env b
      OpIntBitShiftLeft -> do
        x <- go mags env a
        y <- go mags env b
        pure case unlit32 x, unlit32 y of
          Just nx, Just ny -> lit32 (Bits.shl nx ny)
          _, _ -> GoBinOp "<<" x (GoBinOp "&" (GoCall (GoVar "uint32") [ y ]) (GoInt 31))
      OpIntBitShiftRight -> do
        x <- go mags env a
        y <- go mags env b
        pure case unlit32 x, unlit32 y of
          Just nx, Just ny -> lit32 (Bits.shr nx ny)
          _, _ -> GoBinOp ">>" x (GoBinOp "&" (GoCall (GoVar "uint32") [ y ]) (GoInt 31))
      _ -> Nothing
    _ -> Nothing

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

-- Native int32 counterpart of emitContribution for inlined int32 expressions:
-- keeps all 8 lanes in 32-bit registers (eliminating per-element SXTW/SBFIZ
-- and closure InlMark NOOPs) and widens to int64 once after the loop.
-- Lengths below 8 keep `fallback`, preserving wide seeds when len == 0.
emitContribution32 :: { source :: String, prefix :: String, result :: String, offset :: Maybe GoExpr } -> (GoExpr -> GoExpr -> GoExpr) -> GoExpr -> GoExpr
emitContribution32 names contribute32 fallback =
  let
    index = names.prefix <> "_index"
    chunk = names.prefix <> "_chunk"
    base = names.prefix <> "_base"
    lane n = names.prefix <> "_" <> show n
    chunkElem n = GoIndex (GoVar chunk) (GoInt n)
    tailElem = GoIndex (rawGo ("(" <> names.source <> ")")) (GoVar index)
    tailIdx = case names.offset of
      Nothing -> GoVar index
      Just off -> GoBinOp "+" off (GoVar index)
    laneStep idxExpr n = GoMutate (lane n) (GoBinOp "+" (GoVar (lane n)) (contribute32 idxExpr (chunkElem n)))
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
    reduce = GoMutate (lane 0) combined
    tailLoop = GoForRange ("; " <> index <> " < len(" <> names.source <> "); " <> index <> "++")
      [ GoMutate (lane 0) (GoBinOp "+" (GoVar (lane 0)) (contribute32 tailIdx tailElem)) ]
    body = [ GoAssign (lane 0) (GoCall (GoVar "int32") [ GoVar names.result ]) ]
      <> map (\n -> GoAssign (lane n) (rawGo "int32(0)")) (Array.range 1 7)
      <> [ GoAssign index (GoInt 0)
         , mainLoop
         , reduce
         , tailLoop
         , GoMutate names.result (GoCall (GoVar "int64") [ GoVar (lane 0) ])
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

emitContributionReverse32 :: { source :: String, prefix :: String, result :: String } -> (GoExpr -> GoExpr -> GoExpr) -> GoExpr -> GoExpr
emitContributionReverse32 names contribute32 fallback =
  let
    index = names.prefix <> "_index"
    pos = names.prefix <> "_pos"
    chunk = names.prefix <> "_chunk"
    base = names.prefix <> "_base"
    lane n = names.prefix <> "_" <> show n
    chunkElem n = GoIndex (GoVar chunk) (GoInt (7 - n))
    tailElem = GoIndex (rawGo ("(" <> names.source <> ")")) (GoBinOp "-" (GoBinOp "-" (GoCall (GoVar "len") [ rawGo names.source ]) (GoInt 1)) (GoVar index))
    laneStep n = GoMutate (lane n) (GoBinOp "+" (GoVar (lane n)) (contribute32 (GoBinOp "|" (GoVar base) (GoInt n)) (chunkElem n)))
    mainLoop = GoForRange ("; " <> pos <> " >= 8; " <> pos <> " -= 8")
      ([ GoAssign chunk (rawGo ("(" <> names.source <> ")[" <> pos <> "-8 : " <> pos <> "]"))
       , GoAssign base (rawGo (index <> " &^ 7"))
       ] <> map laneStep (Array.range 0 7) <> [ GoMutate index (GoBinOp "+" (GoVar index) (GoInt 8)) ])
    sumPair a b = GoBinOp "+" a b
    combined = sumPair
      (sumPair (sumPair (GoVar (lane 0)) (GoVar (lane 1))) (sumPair (GoVar (lane 2)) (GoVar (lane 3))))
      (sumPair (sumPair (GoVar (lane 4)) (GoVar (lane 5))) (sumPair (GoVar (lane 6)) (GoVar (lane 7))))
    reduce = GoMutate (lane 0) combined
    tailLoop = GoForRange ("; " <> index <> " < len(" <> names.source <> "); " <> index <> "++")
      [ GoMutate (lane 0) (GoBinOp "+" (GoVar (lane 0)) (contribute32 (GoVar index) tailElem)) ]
    body = [ GoAssign (lane 0) (GoCall (GoVar "int32") [ GoVar names.result ]) ]
      <> map (\n -> GoAssign (lane n) (rawGo "int32(0)")) (Array.range 1 7)
      <> [ GoAssign index (GoInt 0)
         , GoAssign pos (GoCall (GoVar "len") [ rawGo names.source ])
         , mainLoop
         , reduce
         , tailLoop
         , GoMutate names.result (GoCall (GoVar "int64") [ GoVar (lane 0) ])
         ]
  in GoIfElse (rawGo ("len(" <> names.source <> ") < 8")) [ fallback ] body
