module Gopurs.PrimitiveExprs
  ( PrimitiveExpr
  , literal
  , unary
  , binary
  , modularIntProduct
  , boundedIntProduct
  , guardedIntProduct
  , powerOfTwoModulo
  , constantPositiveDivMod
  , moduloZeroComparison
  , stringLengthComparison
  ) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Int.Bits as Bits
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.String.CodeUnits as SCU
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Data.Foldable (foldl)
import Effect.Ref (Ref)
import Gopurs.Cache.LookupTable as LookupTable
import Gopurs.CallAnalysis (qualifiedTarget)
import Gopurs.CodegenState (CodegenState)
import Gopurs.ExprAnalysis (getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext, flattenStmts, wrapInStmts)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr, unboxGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Literal(..), ModuleName(..))
import PureScript.Backend.Optimizer.Syntax (BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendSyntax(..), Level, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type PrimitiveExpr =
  { expr :: GoExpr
  , exprType :: GoType
  }

-- Exact modulo-2^32 product encoded as 16-bit split JS-safe arithmetic:
--   ((x & 65535) * (y | 0)) + (((x >> 16) * (y | 0)) << 16)
-- In Go two's-complement int64 arithmetic, int64(int32(x * y)) is exact for
-- every 64-bit x, y. Restrict to pure Local operands so evaluation order and
-- side effects remain unchanged and no IIFE/NOOP overhead is introduced.
modularIntProduct :: TranslateExpr -> ExprContext -> Int -> BackendOperator2 -> TcoExpr -> TcoExpr -> Maybe ExprResult
modularIntProduct translate context nextId operator lhs rhs = do
  guard (operator == OpIntNum OpAdd)
  Tuple lowFactor lowOther <- matchLow lhs
  Tuple highFactor highOther <- matchHigh rhs
  guard (sameLocal lowFactor highFactor && sameLocal lowOther highOther)
  let
    child = childContext context (Just Int)
    left = translate child nextId lowFactor
    right = translate child left.nextId lowOther
    leftVal = unboxGoExpr context.codegenStateRef context.modNameStr left.expr left.exprType TypeInt64
    rightVal = unboxGoExpr context.codegenStateRef context.modNameStr right.expr right.exprType TypeInt64
    cast name argument = GoCall (GoVar name) [ argument ]
  pure
    { stmts: left.stmts <> right.stmts
    , expr: cast "int64" (cast "int32" (GoBinOp "*" leftVal rightVal))
    , exprType: TypeInt64
    , nextId: right.nextId
    }
  where
  matchLow expr = case unwrapTcoExpr expr of
    PrimOp (Syn.Op2 (OpIntNum OpMultiply) a b) -> do
      x <- matchMask16 a
      y <- matchBitOr0 b
      pure (Tuple x y)
    _ -> Nothing

  matchHigh expr = case unwrapTcoExpr expr of
    PrimOp (Syn.Op2 OpIntBitShiftLeft inner shift) | integerLiteral shift == Just 16 ->
      case unwrapTcoExpr inner of
        PrimOp (Syn.Op2 (OpIntNum OpMultiply) a b) -> do
          x <- matchShiftRight16 a
          y <- matchBitOr0 b
          pure (Tuple x y)
        _ -> Nothing
    _ -> Nothing

  matchMask16 expr = case unwrapTcoExpr expr of
    PrimOp (Syn.Op2 OpIntBitAnd x mask) | integerLiteral mask == Just 65535 -> localOperand x
    _ -> Nothing

  matchShiftRight16 expr = case unwrapTcoExpr expr of
    PrimOp (Syn.Op2 OpIntBitShiftRight x shift) | integerLiteral shift == Just 16 -> localOperand x
    _ -> Nothing

  matchBitOr0 expr = case unwrapTcoExpr expr of
    PrimOp (Syn.Op2 OpIntBitOr y zero) | integerLiteral zero == Just 0 -> localOperand y
    _ -> Nothing

  localOperand expr = case unwrapTcoExpr expr of
    Local _ _ -> Just expr
    _ -> Nothing

  sameLocal a b = case unwrapTcoExpr a, unwrapTcoExpr b of
    Local id1 lvl1, Local id2 lvl2 -> id1 == id2 && lvl1 == lvl2
    _, _ -> false

-- IntMul follows JS Number multiplication before narrowing. A native integer
-- product is equivalent only when the syntax bounds its magnitude by 2^53.
-- TAST Int alone is insufficient: public Int values can lie outside Int32.
boundedIntProduct :: TranslateExpr -> ExprContext -> Int -> BackendOperator2 -> TcoExpr -> TcoExpr -> Maybe ExprResult
boundedIntProduct translate context nextId operator lhs rhs = do
  guard (operator == OpIntNum OpMultiply)
  leftBits <- integerMagnitude lhs
  rightBits <- integerMagnitude rhs
  guard (leftBits + rightBits <= 53)
  let
    child = childContext context (Just Int)
    left = translate child nextId lhs
    right = translate child left.nextId rhs
    leftVal = unboxGoExpr context.codegenStateRef context.modNameStr left.expr left.exprType TypeInt64
    rightVal = unboxGoExpr context.codegenStateRef context.modNameStr right.expr right.exprType TypeInt64
    operand result = wrapInStmts context.recVars result.stmts TypeInt64
      (unboxGoExpr context.codegenStateRef context.modNameStr result.expr result.exprType TypeInt64)
    cast name argument = GoCall (GoVar name) [ argument ]
    product = GoBinOp "*" (GoVar "boundedLeft") (GoVar "boundedRight")
    -- When operands have no statements and at least one is a runtime value,
    -- emit direct int64(int32(a * b)) without an IIFE (avoiding Go's InlMark
    -- NOOP). Constant-only products retain the IIFE to prevent compile-time
    -- constant conversion overflow.
    direct = Array.null (flattenStmts left.stmts) && Array.null (flattenStmts right.stmts)
      && (isRuntimeGoExpr leftVal || isRuntimeGoExpr rightVal)
    function = GoFuncBlock [ Tuple "boundedLeft" TypeInt64, Tuple "boundedRight" TypeInt64 ]
      [ GoReturn (cast "int64" (cast "int32" product)) ] TypeInt64
    expr = if direct then cast "int64" (cast "int32" (GoBinOp "*" leftVal rightVal))
      else GoCall function [ operand left, operand right ]
  pure { stmts: StmtEmpty, expr
    , exprType: TypeInt64, nextId: right.nextId }
  where
  isRuntimeGoExpr = case _ of
    GoVar _ -> true
    GoIndex _ _ -> true
    GoSelector _ _ -> true
    GoCall (GoVar name) [ arg ] | name == "int64" || name == "int32" || name == "int" -> isRuntimeGoExpr arg
    GoCall _ _ -> true
    GoBinOp _ l r -> isRuntimeGoExpr l || isRuntimeGoExpr r
    GoPrefixOp _ v -> isRuntimeGoExpr v
    _ -> false

-- A small proved operand makes a cheap exact-integer path useful even when
-- the other operand is an opaque Int64. With b <= 21, the checked range still
-- includes all signed/unsigned 32-bit values. Neither a type annotation nor a
-- local binding is treated as an Int32 range proof. Outside the guard, retain
-- JS Number multiplication and its rounding through the existing runtime.
guardedIntProduct :: TranslateExpr -> ExprContext -> Int -> BackendOperator2 -> TcoExpr -> TcoExpr -> Maybe ExprResult
guardedIntProduct translate context nextId operator lhs rhs = do
  guard (operator == OpIntNum OpMultiply)
  guard (integerAnnotation lhs && integerAnnotation rhs)
  checked <- case integerMagnitude lhs, integerMagnitude rhs of
    Just bits, _ | bits <= 21 -> Just { name: "guardedRight", bits }
    _, Just bits | bits <= 21 -> Just { name: "guardedLeft", bits }
    _, _ -> Nothing
  let
    child = childContext context (Just Int)
    left = translate child nextId lhs
    right = translate child left.nextId rhs
    operand result = wrapInStmts context.recVars result.stmts TypeInt64
      (unboxGoExpr context.codegenStateRef context.modNameStr result.expr result.exprType TypeInt64)
    cast name argument = GoCall (GoVar name) [ argument ]
    limit = GoBinOp "<<" (cast "int64" (GoInt 1)) (GoInt (53 - checked.bits))
    doubleLimit = GoBinOp "<<" (cast "uint64" (GoInt 2)) (GoInt (53 - checked.bits))
    value = GoVar checked.name
    within = GoBinOp "<=" (cast "uint64" (GoBinOp "+" value limit)) doubleLimit
    a = GoVar "guardedLeft"
    b = GoVar "guardedRight"
    function = GoFuncBlock [ Tuple "guardedLeft" TypeInt64, Tuple "guardedRight" TypeInt64 ]
      [ GoIfElse within [ GoReturn (cast "int64" (cast "int32" (GoBinOp "*" a b))) ] []
      , GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "IntMul") [ a, b ])
      ] TypeInt64
  pure { stmts: StmtEmpty, expr: GoCall function [ operand left, operand right ]
    , exprType: TypeInt64, nextId: right.nextId }

-- Euclidean modulo by a positive power of two is the low-bit mask, including
-- for negative/noncanonical Int64 inputs. Use Go's raw mask rather than the
-- PureScript bitwise helper, which also narrows its result to Int32.
powerOfTwoModulo :: TranslateExpr -> ExprContext -> Int -> BackendOperator2 -> TcoExpr -> TcoExpr -> Maybe ExprResult
powerOfTwoModulo translate context nextId operator lhs rhs = do
  guard (operator == OpIntNum OpMod && integerAnnotation lhs)
  modulus <- integerLiteral rhs
  guard (modulus > 0 && modulus <= 2147483647 && Bits.and modulus (modulus - 1) == 0)
  let
    value = translate (childContext context (Just Int)) nextId lhs
    operand = wrapInStmts context.recVars value.stmts TypeInt64
      (unboxGoExpr context.codegenStateRef context.modNameStr value.expr value.exprType TypeInt64)
  pure { stmts: StmtEmpty, expr: GoBinOp "&" operand (GoInt (modulus - 1))
    , exprType: TypeInt64, nextId: value.nextId }

-- Constant positive Euclidean div/mod (0 < d <= 2^31 - 1):
--   * Power of two: exact low-bit mask (& (d - 1)) for mod, exact arithmetic
--     right shift (>> log2(d)) for div, valid for all signed and wide int64.
--   * Other positive constants: IntModPos64 / IntDivPos64 (no zero-divisor
--     check, and 32-bit unsigned magic-multiply fast path for 0 <= a <= MaxInt32).
constantPositiveDivMod :: TranslateExpr -> ExprContext -> Int -> BackendOperator2 -> TcoExpr -> TcoExpr -> Maybe ExprResult
constantPositiveDivMod translate context nextId operator lhs rhs = do
  guard (integerAnnotation lhs)
  divisor <- integerLiteral rhs
  guard (divisor > 0 && divisor <= 2147483647)
  case operator of
    OpIntNum OpMod -> do
      let
        value = translate (childContext context (Just Int)) nextId lhs
        operand = wrapInStmts context.recVars value.stmts TypeInt64
          (unboxGoExpr context.codegenStateRef context.modNameStr value.expr value.exprType TypeInt64)
        expr =
          if Bits.and divisor (divisor - 1) == 0 then
            GoBinOp "&" operand (GoInt (divisor - 1))
          else
            GoCall (GoSelector (GoVar "gopurs_runtime") "IntModPos64")
              [ operand, GoCall (GoVar "int64") [ GoInt divisor ] ]
      pure { stmts: StmtEmpty, expr, exprType: TypeInt64, nextId: value.nextId }
    OpIntNum OpDivide -> do
      let
        value = translate (childContext context (Just Int)) nextId lhs
        operand = wrapInStmts context.recVars value.stmts TypeInt64
          (unboxGoExpr context.codegenStateRef context.modNameStr value.expr value.exprType TypeInt64)
        expr =
          if divisor == 1 then
            operand
          else if Bits.and divisor (divisor - 1) == 0 then
            GoBinOp ">>" operand (GoInt (powerOfTwoShift divisor 0))
          else
            GoCall (GoSelector (GoVar "gopurs_runtime") "IntDivPos64")
              [ operand, GoCall (GoVar "int64") [ GoInt divisor ] ]
      pure { stmts: StmtEmpty, expr, exprType: TypeInt64, nextId: value.nextId }
    _ -> Nothing
  where
  powerOfTwoShift value shift
    | value <= 1 = shift
    | otherwise = powerOfTwoShift (Bits.zshr value 1) (shift + 1)

-- Euclidean IntMod(a, b) == 0 holds iff b == 0 || a % b == 0 (the Euclidean
-- sign fixup never turns a non-zero truncated remainder into 0).
moduloZeroComparison :: TranslateExpr -> ExprContext -> Int -> BackendOperator2 -> TcoExpr -> TcoExpr -> Maybe ExprResult
moduloZeroComparison translate context nextId operator lhs rhs = do
  negate <- case operator of
    OpIntOrd OpEq -> Just false
    OpIntOrd OpNotEq -> Just true
    _ -> Nothing
  Tuple dividend divisor <- case matchMod lhs, integerLiteral rhs of
    Just pair, Just 0 -> Just pair
    _, _ -> case integerLiteral lhs, matchMod rhs of
      Just 0, Just pair -> Just pair
      _, _ -> Nothing
  guard (integerAnnotation dividend && integerAnnotation divisor)
  let child = childContext context (Just Int)
  case integerLiteral divisor of
    Just d | d > 0 && d <= 2147483647 -> do
      resA <- Just (translate child nextId dividend)
      let
        opA = unboxGoExpr context.codegenStateRef context.modNameStr resA.expr resA.exprType TypeInt64
        remExpr =
          if Bits.and d (d - 1) == 0 then GoBinOp "&" opA (GoInt (d - 1))
          else GoBinOp "%" opA (GoCall (GoVar "int64") [ GoInt d ])
        cmpOp = if negate then "!=" else "=="
      pure { stmts: resA.stmts, expr: GoBinOp cmpOp remExpr (GoInt 0), exprType: TypeBool, nextId: resA.nextId }
    _ -> do
      let
        resA = translate child nextId dividend
        resB = translate child resA.nextId divisor
        opA = unboxGoExpr context.codegenStateRef context.modNameStr resA.expr resA.exprType TypeInt64
        opB = unboxGoExpr context.codegenStateRef context.modNameStr resB.expr resB.exprType TypeInt64
        zeroCheck = GoCall (GoSelector (GoVar "gopurs_runtime") "IntModZero") [ opA, opB ]
      pure
        { stmts: resA.stmts <> resB.stmts
        , expr: if negate then GoPrefixOp "!" zeroCheck else zeroCheck
        , exprType: TypeBool
        , nextId: resB.nextId
        }
  where
  matchMod expr = case unwrapTcoExpr expr of
    PrimOp (Syn.Op2 (OpIntNum OpMod) a b) -> Just (Tuple a b)
    _ -> Nothing

integerAnnotation :: TcoExpr -> Boolean
integerAnnotation expression = case getExprType expression of
  Any -> true
  Int -> true
  _ -> false

-- Return b only when |value| <= 2^b. These operations normalize their results;
-- division, unsigned shifts, arbitrary calls and plain locals are not proofs.
integerMagnitude :: TcoExpr -> Maybe Int
integerMagnitude = integerMagnitudeWith Map.empty

integerMagnitudeWith :: Map Level Int -> TcoExpr -> Maybe Int
integerMagnitudeWith known (TcoExpr _ syntax) = case syntax of
  Typed Int inner -> integerMagnitudeWith known inner
  Typed Any inner -> integerMagnitudeWith known inner
  Syn.TypeApp inner _ -> integerMagnitudeWith known inner
  Lit (LitInt value) -> literalMagnitude value
  Local _ level -> Map.lookup level known
  Let _ level rhs body ->
    let known' = case integerMagnitudeWith known rhs of
          Just m -> Map.insert level m known
          Nothing -> known
    in integerMagnitudeWith known' body
  Branch branches otherwise -> do
    defMag <- integerMagnitudeWith known otherwise
    branchMags <- traverse (\(Pair _ res) -> integerMagnitudeWith known res) (NEA.toArray branches)
    pure (foldl max defMag branchMags)
  PrimOp (Syn.Op1 OpIntNegate _) -> Just 31
  PrimOp (Syn.Op1 OpIntBitNot _) -> Just 31
  PrimOp (Syn.Op2 operation left right) -> case operation of
    OpIntBitAnd -> case integerLiteral left, integerLiteral right of
      Just mask, _ | mask >= 0 -> literalMagnitude mask
      _, Just mask | mask >= 0 -> literalMagnitude mask
      _, _ -> Just 31
    OpIntBitShiftRight -> case integerLiteral right of
      Just shift | shift >= 0 && shift <= 31 -> Just (31 - shift)
      _ -> Just 31
    OpIntBitOr -> Just 31
    OpIntBitXor -> Just 31
    OpIntBitShiftLeft -> Just 31
    OpIntNum OpAdd -> sumMagnitude known left right
    OpIntNum OpSubtract -> sumMagnitude known left right
    OpIntNum OpMultiply -> Just 31
    OpIntNum OpMod -> case integerLiteral right of
      Just modulus | modulus > 0 && modulus <= 2147483647 -> literalMagnitude (modulus - 1)
      _ -> Nothing
    _ -> Nothing
  _ -> Nothing

sumMagnitude :: Map Level Int -> TcoExpr -> TcoExpr -> Maybe Int
sumMagnitude known left right = case integerMagnitudeWith known left, integerMagnitudeWith known right of
  Just a, Just b -> Just (min 31 (max a b + 1))
  _, _ -> Just 31

integerLiteral :: TcoExpr -> Maybe Int
integerLiteral (TcoExpr _ syntax) = case syntax of
  Typed Int inner -> integerLiteral inner
  Typed Any inner -> integerLiteral inner
  Syn.TypeApp inner _ -> integerLiteral inner
  Lit (LitInt value) -> Just value
  _ -> Nothing

literalMagnitude :: Int -> Maybe Int
literalMagnitude value
  | value < (-2147483647 - 1) || value > 2147483647 = Nothing
  | value == (-2147483647 - 1) = Just 31
  | otherwise = Just (bits (if value < 0 then negate value else value) 0)
  where
  bits remaining count
    | remaining <= 1 = count
    | otherwise = bits (div remaining 2 + mod remaining 2) (count + 1)

-- Length equality needs neither complete scan when both immutable strings are
-- the same view. Admit the canonical, typed operation before lowering its
-- calls; similarly named functions and other comparisons keep their fallback.
stringLengthComparison :: TranslateExpr -> ExprContext -> Int -> BackendOperator2 -> TcoExpr -> TcoExpr -> Maybe ExprResult
stringLengthComparison translate context nextId operator lhs rhs = do
  negate <- case operator of
    OpIntOrd OpEq -> Just false
    OpIntOrd OpNotEq -> Just true
    _ -> Nothing
  left <- lengthInput context lhs
  right <- lengthInput context rhs
  let
    child = childContext context (Just String)
    res1 = translate child nextId left
    res2 = translate child res1.nextId right
    operand result = wrapInStmts context.recVars result.stmts TypeValue
      (boxGoExpr context.codegenStateRef context.modNameStr result.expr result.exprType)
    equal = GoCall (GoSelector (GoVar "gopurs_runtime") "StringCodeUnitsLengthsEqual")
      [ operand res1, operand res2 ]
  -- Keep each operand's statements with its value. Hoisting the right-hand
  -- statements ahead of the left-hand call would change evaluation order.
  pure { stmts: StmtEmpty, expr: if negate then GoPrefixOp "!" equal else equal
    , exprType: TypeBool, nextId: res2.nextId }

lengthInput :: ExprContext -> TcoExpr -> Maybe TcoExpr
lengthInput context expression = case unwrapTcoExpr expression of
  Syn.App fn args -> case NEA.toArray args of
    [ argument ] -> admit fn argument
    _ -> Nothing
  Syn.UncurriedApp fn [ argument ] -> admit fn argument
  _ -> Nothing
  where
  admit fn argument = do
    target <- qualifiedTarget fn
    guard (target.mbMod == Just (ModuleName "Data.String.CodeUnits") && target.name == "length")
    let
      signature = case getExprType fn of
        Any -> LookupTable.lookup "Data.String.CodeUnits.length" context.metadata.globalTypes
        ty -> Just ty
    guard (signature == Just (Func [ String ] Int))
    guard (case getExprType expression of
      Any -> true
      Int -> true
      _ -> false)
    guard (case getExprType argument of
      Any -> true
      String -> true
      _ -> false)
    pure argument

-- Arrays and records retain their recursive translation in CodeGen.
literal :: forall a. Literal a -> Maybe PrimitiveExpr
literal = case _ of
  LitString s -> Just { expr: GoString s, exprType: TypeString }
  LitInt i -> Just { expr: GoCall (GoVar "int64") [GoInt i], exprType: TypeInt64 }
  LitNumber n ->
    let
      nStr = show n
      expr =
        if n == 0.0 && 1.0 / n < 0.0 then GoCall (GoSelector (GoVar "gopurs_runtime") "NegativeZero") []
        else if nStr == "Infinity" then GoCall (GoSelector (GoVar "math") "Inf") [ GoInt 1 ]
        else if nStr == "-Infinity" then GoCall (GoSelector (GoVar "math") "Inf") [ GoInt (-1) ]
        else if nStr == "NaN" then GoCall (GoSelector (GoVar "math") "NaN") []
        else rawGo nStr
    in
      Just { expr, exprType: TypeFloat64 }
  LitBoolean b -> Just { expr: rawGo (if b then "true" else "false"), exprType: TypeBool }
  LitChar c -> Just { expr: GoString (SCU.singleton c), exprType: TypeString }
  _ -> Nothing

-- Select the scalar emitter before translating operands. Operators that need
-- statements or control flow remain in CodeGen; emitted expressions use only
-- already translated operands.
unary :: Ref CodegenState -> String -> BackendOperator1 -> Maybe (PrimitiveExpr -> PrimitiveExpr)
unary codegenStateRef modNameStr = case _ of
  OpBooleanNot -> Just \resE -> { expr: GoBinOp "!=" (unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeBool) (rawGo "true"), exprType: TypeBool }
  OpIntNegate -> Just \resE -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntNegate") [unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeInt64], exprType: TypeInt64 }
  OpIntBitNot -> Just \resE -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntBitNot") [unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeInt64], exprType: TypeInt64 }
  OpNumberNegate -> Just \resE -> { expr: GoPrefixOp "-" (unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeFloat64), exprType: TypeFloat64 }
  _ -> Nothing

binary :: Ref CodegenState -> String -> BackendOperator2 -> Maybe (PrimitiveExpr -> PrimitiveExpr -> PrimitiveExpr)
binary codegenStateRef modNameStr = case _ of
  OpIntNum OpAdd -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64], exprType: TypeInt64 }
  OpIntNum OpSubtract -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntSub") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64], exprType: TypeInt64 }
  OpIntNum OpMultiply -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntMul") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64], exprType: TypeInt64 }
  OpIntNum OpDivide -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntDiv") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64], exprType: TypeInt64 }
  OpIntNum OpMod -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntMod") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64], exprType: TypeInt64 }
  OpIntBitAnd -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntBitAnd") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64], exprType: TypeInt64 }
  OpIntBitOr -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntBitOr") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64], exprType: TypeInt64 }
  OpIntBitXor -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntBitXor") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64], exprType: TypeInt64 }
  OpIntBitShiftLeft -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntShl") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64], exprType: TypeInt64 }
  OpIntBitShiftRight -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntShr") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64], exprType: TypeInt64 }
  OpIntBitZeroFillShiftRight -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Zshr") [ boxGoExpr codegenStateRef modNameStr res1.expr res1.exprType, boxGoExpr codegenStateRef modNameStr res2.expr res2.exprType ], exprType: TypeValue }
  OpIntOrd OpEq -> Just \res1 res2 -> { expr: GoBinOp "==" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
  OpIntOrd OpNotEq -> Just \res1 res2 -> { expr: GoBinOp "!=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
  OpIntOrd OpLt -> Just \res1 res2 -> { expr: GoBinOp "<" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
  OpIntOrd OpLte -> Just \res1 res2 -> { expr: GoBinOp "<=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
  OpIntOrd OpGt -> Just \res1 res2 -> { expr: GoBinOp ">" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
  OpIntOrd OpGte -> Just \res1 res2 -> { expr: GoBinOp ">=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
  OpNumberNum OpAdd -> Just \res1 res2 -> { expr: GoBinOp "+" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64), exprType: TypeFloat64 }
  OpNumberNum OpSubtract -> Just \res1 res2 -> { expr: GoBinOp "-" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64), exprType: TypeFloat64 }
  OpNumberNum OpMultiply -> Just \res1 res2 -> { expr: GoBinOp "*" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64), exprType: TypeFloat64 }
  OpNumberNum OpDivide -> Just \res1 res2 -> { expr: GoBinOp "/" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64), exprType: TypeFloat64 }
  OpNumberNum OpMod -> Just \res1 res2 -> { expr: GoCall (GoVar "math.Mod") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64], exprType: TypeFloat64 }
  OpNumberOrd OpEq -> Just \res1 res2 -> { expr: GoBinOp "==" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
  OpNumberOrd OpNotEq -> Just \res1 res2 -> { expr: GoBinOp "!=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
  OpNumberOrd OpLt -> Just \res1 res2 -> { expr: GoBinOp "<" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
  OpNumberOrd OpLte -> Just \res1 res2 -> { expr: GoBinOp "<=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
  OpNumberOrd OpGt -> Just \res1 res2 -> { expr: GoBinOp ">" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
  OpNumberOrd OpGte -> Just \res1 res2 -> { expr: GoBinOp ">=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeFloat64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
  OpStringAppend -> Just \res1 res2 -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ConcatString") [unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString, unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString], exprType: TypeString }
  OpStringOrd OpEq -> Just \res1 res2 -> { expr: GoBinOp "==" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpStringOrd OpNotEq -> Just \res1 res2 -> { expr: GoBinOp "!=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpStringOrd OpLt -> Just \res1 res2 -> { expr: GoBinOp "<" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpStringOrd OpLte -> Just \res1 res2 -> { expr: GoBinOp "<=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpStringOrd OpGt -> Just \res1 res2 -> { expr: GoBinOp ">" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpStringOrd OpGte -> Just \res1 res2 -> { expr: GoBinOp ">=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpCharOrd OpEq -> Just \res1 res2 -> { expr: GoBinOp "==" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpCharOrd OpNotEq -> Just \res1 res2 -> { expr: GoBinOp "!=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpCharOrd OpLt -> Just \res1 res2 -> { expr: GoBinOp "<" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpCharOrd OpLte -> Just \res1 res2 -> { expr: GoBinOp "<=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpCharOrd OpGt -> Just \res1 res2 -> { expr: GoBinOp ">" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpCharOrd OpGte -> Just \res1 res2 -> { expr: GoBinOp ">=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeBool }
  OpBooleanOrd OpEq -> Just \res1 res2 -> { expr: GoBinOp "==" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool }
  OpBooleanOrd OpNotEq -> Just \res1 res2 -> { expr: GoBinOp "!=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool }
  -- Go cannot order bools. Their canonical payloads are 0/1; comparing those
  -- preserves false < true and evaluates both operands exactly once.
  OpBooleanOrd OpLt -> Just \res1 res2 -> { expr: GoBinOp "<" (GoSelector (boxGoExpr codegenStateRef modNameStr res1.expr res1.exprType) "IntVal") (GoSelector (boxGoExpr codegenStateRef modNameStr res2.expr res2.exprType) "IntVal"), exprType: TypeBool }
  OpBooleanOrd OpLte -> Just \res1 res2 -> { expr: GoBinOp "<=" (GoSelector (boxGoExpr codegenStateRef modNameStr res1.expr res1.exprType) "IntVal") (GoSelector (boxGoExpr codegenStateRef modNameStr res2.expr res2.exprType) "IntVal"), exprType: TypeBool }
  OpBooleanOrd OpGt -> Just \res1 res2 -> { expr: GoBinOp ">" (GoSelector (boxGoExpr codegenStateRef modNameStr res1.expr res1.exprType) "IntVal") (GoSelector (boxGoExpr codegenStateRef modNameStr res2.expr res2.exprType) "IntVal"), exprType: TypeBool }
  OpBooleanOrd OpGte -> Just \res1 res2 -> { expr: GoBinOp ">=" (GoSelector (boxGoExpr codegenStateRef modNameStr res1.expr res1.exprType) "IntVal") (GoSelector (boxGoExpr codegenStateRef modNameStr res2.expr res2.exprType) "IntVal"), exprType: TypeBool }
  _ -> Nothing
