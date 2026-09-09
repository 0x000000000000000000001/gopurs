module Gopurs.PrimitiveExprs
  ( PrimitiveExpr
  , literal
  , unary
  , binary
  ) where

import Prelude

import Data.Maybe (Maybe(..))
import Data.String.CodeUnits as SCU
import Effect.Ref (Ref)
import Gopurs.CodegenState (CodegenState)
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr, unboxGoExpr)
import PureScript.Backend.Optimizer.CoreFn (Literal(..))
import PureScript.Backend.Optimizer.Syntax (BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..))

type PrimitiveExpr =
  { expr :: GoExpr
  , exprType :: GoType
  }

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
        else GoRaw nStr
    in
      Just { expr, exprType: TypeFloat64 }
  LitBoolean b -> Just { expr: GoRaw (if b then "true" else "false"), exprType: TypeBool }
  LitChar c -> Just { expr: GoString (SCU.singleton c), exprType: TypeString }
  _ -> Nothing

-- Select the scalar emitter before translating operands. Operators that need
-- statements or control flow remain in CodeGen; emitted expressions use only
-- already translated operands.
unary :: Ref CodegenState -> String -> BackendOperator1 -> Maybe (PrimitiveExpr -> PrimitiveExpr)
unary codegenStateRef modNameStr = case _ of
  OpBooleanNot -> Just \resE -> { expr: GoBinOp "!=" (unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeBool) (GoRaw "true"), exprType: TypeBool }
  OpIntNegate -> Just \resE -> { expr: GoPrefixOp "-" (unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeInt64), exprType: TypeInt64 }
  OpIntBitNot -> Just \resE -> { expr: GoBinOp "^" (GoRaw "^0") (unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeInt64), exprType: TypeInt64 }
  OpNumberNegate -> Just \resE -> { expr: GoPrefixOp "-" (unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeFloat64), exprType: TypeFloat64 }
  _ -> Nothing

binary :: Ref CodegenState -> String -> BackendOperator2 -> Maybe (PrimitiveExpr -> PrimitiveExpr -> PrimitiveExpr)
binary codegenStateRef modNameStr = case _ of
  OpIntNum OpAdd -> Just \res1 res2 -> { expr: GoBinOp "+" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
  OpIntNum OpSubtract -> Just \res1 res2 -> { expr: GoBinOp "-" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
  OpIntNum OpMultiply -> Just \res1 res2 -> { expr: GoBinOp "*" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
  OpIntNum OpDivide -> Just \res1 res2 -> { expr: GoBinOp "/" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
  OpIntNum OpMod -> Just \res1 res2 -> { expr: GoBinOp "%" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
  OpIntBitAnd -> Just \res1 res2 -> { expr: GoBinOp "&" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
  OpIntBitOr -> Just \res1 res2 -> { expr: GoBinOp "|" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
  OpIntBitXor -> Just \res1 res2 -> { expr: GoBinOp "^" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
  OpIntBitShiftLeft -> Just \res1 res2 -> { expr: GoBinOp "<<" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
  OpIntBitShiftRight -> Just \res1 res2 -> { expr: GoBinOp ">>" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeInt64) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
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
  OpStringAppend -> Just \res1 res2 -> { expr: GoBinOp "+" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeString) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeString), exprType: TypeString }
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
  OpBooleanOrd OpLt -> Just \res1 res2 -> { expr: GoBinOp "<" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool }
  OpBooleanOrd OpLte -> Just \res1 res2 -> { expr: GoBinOp "<=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool }
  OpBooleanOrd OpGt -> Just \res1 res2 -> { expr: GoBinOp ">" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool }
  OpBooleanOrd OpGte -> Just \res1 res2 -> { expr: GoBinOp ">=" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool }
  _ -> Nothing
