module Gopurs.ImmediateEffectBranches (emit) where

import Prelude

import Control.Alternative (guard)
import Data.Foldable (any)
import Data.Maybe (Maybe(..))
import Gopurs.ControlExprs as ControlExprs
import Gopurs.ExprAnalysis (isEffectNode, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (boxGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
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
    let joined = context { mbExpectedExprType = Nothing, mbExpectedGoType = Just TypeValue }
    pure (ControlExprs.branchWith translate (execute translate) joined nextId branches otherwise)
  _ -> Nothing

benefits :: TcoExpr -> Boolean
benefits expression = case unwrapTcoExpr expression of
  Branch branches otherwise -> any (\(Pair _ body) -> benefits body) branches || benefits otherwise
  _ -> isEffectNode expression

execute :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
execute translate context nextId expression = case emit translate context nextId expression of
  Just result -> result
  Nothing ->
    let
      known = isEffectNode expression
      -- EffectExprs.wrap clears these expected types and disables tail calls
      -- inside its closure. Keep that same boundary when removing the closure.
      translated = translate (context
        { options = { isTail: false, inEffectBlock: known }
        , mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) nextId expression
      boxed = boxGoExpr context.codegenStateRef context.modNameStr translated.expr translated.exprType
      result = if known then boxed else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply")
        [ boxed, rawGo "gopurs_runtime.Value{}" ]
    in translated { expr = result, exprType = TypeValue }
