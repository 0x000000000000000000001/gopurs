module Gopurs.TailRecEffectStep (Plan, StepPlan, recognize, emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.ExprAnalysis (executeIfOpaque, extractFuncType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, TranslateExpr, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Abs, Branch, CtorSaturated, EffectBind, EffectDefer, EffectPure, Fail, Let), Level, Pair(..))

-- A syntactic lambda is an inert value. All of its terminal action paths must
-- expose the canonical Step constructor before any translation is attempted.
-- Opaque callbacks and opaque action results retain the ordinary runtime path.
type Plan = { parameter :: String, stateType :: GoType, body :: StepPlan }

data StepPlan
  = Loop TcoExpr
  | Done TcoExpr
  | Branches (Array (Tuple TcoExpr StepPlan)) StepPlan
  | BindValue (Maybe Ident) Level TcoExpr StepPlan
  | BindAction (Maybe Ident) Level TcoExpr StepPlan
  | Failure String

recognize :: ExprContext -> TcoExpr -> Maybe Plan
recognize { metadata, modNameStr } callback = case unwrapTcoExpr callback of
  Abs args body -> do
    Tuple ident level <- case NEA.toArray args of
      [ parameter ] -> Just parameter
      _ -> Nothing
    plan <- analyze true body
    functionType <- extractFuncType callback
    state <- Array.head functionType.fArgs
    let stateType = exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr state
    pure { parameter: localId ident level, stateType, body: plan }
  _ -> Nothing

analyze :: Boolean -> TcoExpr -> Maybe StepPlan
analyze action expression = case unwrapTcoExpr expression of
  Branch branches fallback -> Branches
    <$> traverse (\(Pair condition body) -> Tuple condition <$> analyze action body) (NEA.toArray branches)
    <*> analyze action fallback
  Let ident level binding body -> BindValue ident level binding <$> analyze action body
  EffectDefer body | action -> analyze true body
  EffectPure result | action -> analyze false result
  EffectBind ident level binding body | action -> BindAction ident level binding <$> analyze true body
  CtorSaturated (Qualified owner _) _ _ (Ident name) fields | not action -> do
    guard (owner == Just (ModuleName "Control.Monad.Rec.Class"))
    value <- case fields of
      [ Tuple _ value ] -> Just value
      _ -> Nothing
    case name of
      "Loop" -> Just (Loop value)
      "Done" -> Just (Done value)
      _ -> Nothing
  Fail message -> Just (Failure message)
  _ -> Nothing

type Emission = { statements :: Array GoExpr, nextId :: Int }

emit :: TranslateExpr -> ExprContext -> Int -> String -> String -> Plan -> Emission
emit translate context nextId stateName loopName plan =
  let
    -- Escaping closures must capture this iteration's immutable parameter,
    -- never the mutable state used by the following iteration.
    parameterName = "tailRecEffect_parameter_" <> show nextId
    bound = Map.insert plan.parameter { name: parameterName, goType: plan.stateType } context.bound
    result = emitPlan translate (context { bound = bound }) (nextId + 1) stateName loopName plan.stateType plan.body
  in result { statements = Array.cons (GoAssign parameterName (GoVar stateName)) result.statements }

emitPlan :: TranslateExpr -> ExprContext -> Int -> String -> String -> GoType -> StepPlan -> Emission
emitPlan translate context@{ codegenStateRef, modNameStr } nextId stateName loopName stateType = case _ of
  Loop value ->
    let result = translate (childContext context Nothing) nextId value
    in { statements: flattenStmts result.stmts <>
          [ GoMutate stateName (coerceGoExpr codegenStateRef modNameStr result.expr result.exprType stateType)
          , GoContinue loopName ], nextId: result.nextId }
  Done value ->
    let result = translate (childContext context Nothing) nextId value
    in { statements: flattenStmts result.stmts <>
          [ GoReturn (boxGoExpr codegenStateRef modNameStr result.expr result.exprType) ], nextId: result.nextId }
  Failure message -> { statements: [ GoCall (GoVar "panic") [ GoString message ] ], nextId }
  BindValue ident level value body -> bind false ident level value body
  BindAction ident level action body -> bind true ident level action body
  Branches branches fallback ->
    let
      parts = foldl
        (\acc (Tuple condition body) ->
          let
            test = translate (childContext context Nothing) acc.nextId condition
            branch = emitPlan translate context test.nextId stateName loopName stateType body
          in
            { statements: acc.statements <> [ GoBlock (flattenStmts test.stmts <>
                [ GoIfElse (coerceGoExpr codegenStateRef modNameStr test.expr test.exprType TypeBool) branch.statements [] ]) ]
            , nextId: branch.nextId
            }) { statements: [], nextId } branches
      final = emitPlan translate context parts.nextId stateName loopName stateType fallback
    in { statements: parts.statements <> [ GoBlock final.statements ], nextId: final.nextId }
  where
  bind action ident level value body =
    let
      -- EffectDefer explicitly wraps an action; other recognized effect nodes
      -- execute in effect context, and opaque actions keep their ordinary Apply.
      runAction expr = case unwrapTcoExpr expr of
        EffectDefer inner -> runAction inner
        _ ->
          let actionResult = translate ((childContext context Nothing) { options = { isTail: false, inEffectBlock: true } }) (nextId + 1) expr
          in actionResult { expr = executeIfOpaque expr (boxGoExpr codegenStateRef modNameStr actionResult.expr actionResult.exprType), exprType = TypeValue }
      result = if action then runAction value else translate (childContext context Nothing) (nextId + 1) value
      name = localId ident level <> "_" <> show nextId
      bound = Map.insert (localId ident level) { name, goType: result.exprType } context.bound
      rest = emitPlan translate (context { bound = bound }) result.nextId stateName loopName stateType body
    in { statements: flattenStmts result.stmts <> [ GoAssign name result.expr ] <> rest.statements, nextId: rest.nextId }
