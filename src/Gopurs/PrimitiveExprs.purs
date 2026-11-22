module Gopurs.PrimitiveExprs
  ( PrimitiveExpr
  , literal
  , unary
  , binary
  , stringLengthComparison
  ) where

import Prelude

import Control.Alternative (guard)
import Data.Array.NonEmpty as NEA
import Data.Maybe (Maybe(..))
import Data.String.CodeUnits as SCU
import Effect.Ref (Ref)
import Gopurs.Cache.LookupTable as LookupTable
import Gopurs.CallAnalysis (qualifiedTarget)
import Gopurs.CodegenState (CodegenState)
import Gopurs.ExprAnalysis (getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext, wrapInStmts)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr, unboxGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Literal(..), ModuleName(..))
import PureScript.Backend.Optimizer.Syntax (BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type PrimitiveExpr =
  { expr :: GoExpr
  , exprType :: GoType
  }

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
