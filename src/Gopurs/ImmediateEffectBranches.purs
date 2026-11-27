module Gopurs.ImmediateEffectBranches (emit) where

import Prelude

import Control.Alternative (guard)
import Data.Foldable (any)
import Data.Maybe (Maybe(..))
import Gopurs.ControlExprs as ControlExprs
import Gopurs.ExprAnalysis (getExprType, isEffectNode, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (boxGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Branch), Pair(..))

-- Called only where ordinaryBind would immediately Apply the selected action.
-- It does not rewrite branch values saved for later execution. Known effect
-- nodes use exactly the context normally used inside EffectExprs.wrap; other
-- leaves retain their original construction and one runtime Apply. Conditions
-- keep their original order, short circuiting and ordinary translation.
emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Maybe ExprResult
emit translate context nextId expression = case unwrapTcoExpr expression of
  Branch branches otherwise -> do
    guard (benefits expression)
    let
      isEffectInt = case getExprType expression of
        ADT "Effect.Effect" _ [ Int ] -> true
        _ -> case context.mbExpectedExprType of
          Just (ADT "Effect.Effect" _ [ Int ]) -> true
          _ -> false
      targetGoType = if isEffectInt then TypeInt64 else TypeValue
      joined = context { mbExpectedExprType = Nothing, mbExpectedGoType = Just targetGoType }
    pure (ControlExprs.branchWith translate (execute isEffectInt translate) joined nextId branches otherwise)
  _ -> Nothing

benefits :: TcoExpr -> Boolean
benefits expression = case unwrapTcoExpr expression of
  Branch branches otherwise -> any (\(Pair _ body) -> benefits body) branches || benefits otherwise
  _ -> isEffectNode expression

execute :: Boolean -> TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
execute isEffectInt translate context nextId expression = case emit translate context nextId expression of
  Just result ->
    if isEffectInt && result.exprType == TypeValue then
      result { expr = unboxInt64 result.expr, exprType = TypeInt64 }
    else result
  Nothing ->
    let
      known = isEffectNode expression
      -- EffectExprs.wrap clears these expected types and disables tail calls
      -- inside its closure. Keep that same boundary when removing the closure.
      translated = translate (context
        { options = { isTail: false, inEffectBlock: known }
        , mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) nextId expression
      boxed = boxGoExpr context.codegenStateRef context.modNameStr translated.expr translated.exprType
      valueResult = if known then boxed else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply")
        [ boxed, rawGo "gopurs_runtime.Value{}" ]
    in
      if isEffectInt then
        case known, translated.exprType, translated.expr of
          true, TypeInt64, native -> translated { expr = native, exprType = TypeInt64 }
          true, TypeValue, GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ native ] -> translated { expr = native, exprType = TypeInt64 }
          _, _, _ -> translated { expr = unboxInt64 valueResult, exprType = TypeInt64 }
      else translated { expr = valueResult, exprType = TypeValue }
  where
  unboxInt64 value = GoCall (GoTypeApply (GoSelector (GoVar "gopurs_runtime") "Unbox") [ TypeInt64 ]) [ value ]
