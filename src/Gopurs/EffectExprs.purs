module Gopurs.EffectExprs
  ( bindEffect
  , pureValue
  , defer
  , primitive
  , wrap
  , wrapCachedInt
  ) where

import Prelude
import Control.Alternative (guard)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Gopurs.ExprAnalysis (executeIfOpaque, getExprType, isEffectNode, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr)
import Gopurs.ImmediateFfiEffects as ImmediateFfiEffects
import Gopurs.ImmediateEffectBranches as ImmediateEffectBranches
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendEffect(..), BackendSyntax(Abs, EffectDefer, Let, LetRec), Level)

bindEffect :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> ExprResult
bindEffect translate context nextId mbIdent lvl binding body =
  case ImmediateFfiEffects.direct translate context nextId mbIdent lvl binding body of
    Just result -> result
    Nothing -> case ImmediateFfiEffects.invokeBoundThunk translate context nextId mbIdent lvl binding body of
      Just result -> result
      Nothing -> ordinaryBind translate context nextId mbIdent lvl binding body

ordinaryBind :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> ExprResult
ordinaryBind translate context@{ codegenStateRef, depth, modNameStr, bound, options: { isTail } } nextId mbIdent lvl binding body =
  let
    stripEffectDefer (TcoExpr a syn) = case unwrapTcoExpr (TcoExpr a syn) of
      EffectDefer inner -> let Tuple _ i = stripEffectDefer inner in Tuple true i
      Abs _ inner -> let Tuple _ i = stripEffectDefer inner in Tuple true i
      Let ident lvl_ val body_ -> let Tuple stripped i = stripEffectDefer body_ in Tuple stripped (TcoExpr a (Let ident lvl_ val i))
      LetRec lvl_ bindings_ body_ -> let Tuple stripped i = stripEffectDefer body_ in Tuple stripped (TcoExpr a (LetRec lvl_ bindings_ i))
      _ -> Tuple false (TcoExpr a syn)
    Tuple wasStripped realBinding = stripEffectDefer binding
    originalName = localId mbIdent lvl
    name = originalName <> "_" <> show nextId
    newBound = Map.insert originalName { name, goType: TypeValue } bound
    resBinding = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: true }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) (nextId + 1) realBinding
    bodyContext = context { depth = (depth + 1), bound = newBound, tcoIdent = Nothing, options = { isTail, inEffectBlock: true } }
    executedBody = ImmediateEffectBranches.emit translate bodyContext resBinding.nextId body
    resBody = case executedBody of
      Just result -> result
      Nothing -> translate bodyContext resBinding.nextId body
    bindingExpr = if wasStripped then boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType else executeIfOpaque realBinding (boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType)
    bodyExpr = case executedBody of
      Just _ -> resBody.expr
      Nothing -> executeIfOpaque body resBody.expr
  in
    { stmts: resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: bodyExpr, exprType: resBody.exprType, nextId: resBody.nextId }

pureValue :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
pureValue translate context@{ depth, mbExpectedExprType } nextId binding =
  let
    expectedInnerInt = case mbExpectedExprType of
      Just (ADT "Effect.Effect" _ [ ADT "Effect.Effect" _ [ Int ] ]) -> true
      _ -> false
    expectedScalarInt = case mbExpectedExprType of
      Just (ADT "Effect.Effect" _ [ Int ]) -> true
      _ -> false
    innerContext = context
      { depth = (depth + 1), tcoIdent = Nothing, loopCtx = []
      , options = { isTail: false, inEffectBlock: false }
      , mbExpectedExprType = if expectedScalarInt then Just Int else Nothing
      , mbExpectedGoType = if expectedScalarInt then Just TypeInt64 else Nothing
      }
    wrapped = if expectedInnerInt && isEffectNode binding then
      Just (wrapIntAction translate (innerContext { mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) nextId binding)
      else wrapCachedInt translate (innerContext { mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) nextId binding
  in
    case wrapped of
      Just result -> result
      Nothing -> translate innerContext nextId binding

defer :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
defer translate context@{ codegenStateRef, depth, modNameStr } nextId binding =
  let
    resBinding = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: true }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) nextId binding
    funcExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
      [ GoFuncBlock [ Tuple "_" TypeValue ]
          (flattenStmts resBinding.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType) ])
          TypeValue
      ]
  in
    { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBinding.nextId }

primitive :: TranslateExpr -> ExprContext -> Int -> BackendEffect TcoExpr -> ExprResult
primitive translate context@{ codegenStateRef, depth, modNameStr } nextId eff =
  case eff of
  EffectRefNew a ->
    let
      resA = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) nextId a
    in
      { stmts: resA.stmts
      , expr: GoCall (GoSelector (GoVar "gopurs_runtime") "NewSTRef")
          [ boxGoExpr codegenStateRef modNameStr resA.expr resA.exprType ]
      , exprType: TypeValue
      , nextId: resA.nextId + 1
      }
  EffectRefRead a ->
    let
      resA = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) nextId a
    in
      { stmts: resA.stmts
      , expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ReadSTRef")
          [ boxGoExpr codegenStateRef modNameStr resA.expr resA.exprType ]
      , exprType: TypeValue
      , nextId: resA.nextId
      }
  EffectRefWrite ref val ->
    let
      resRef = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) nextId ref
      resVal = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) resRef.nextId val
    in
      { stmts: resRef.stmts <> resVal.stmts
      , expr: GoCall (GoSelector (GoVar "gopurs_runtime") "WriteSTRef")
          [ boxGoExpr codegenStateRef modNameStr resRef.expr resRef.exprType
          , boxGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType
          ]
      , exprType: TypeValue
      , nextId: resVal.nextId
      }

wrap :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
wrap translate context nextId tcoExpr =
  let
    expectedEffect = case context.mbExpectedExprType of
      Just ty@(ADT "Effect.Effect" _ [ ADT "Effect.Effect" _ [ Int ] ]) -> Just ty
      _ -> Nothing
    res = translate (context { options = { isTail: false, inEffectBlock: true }, mbExpectedExprType = expectedEffect, mbExpectedGoType = Nothing }) nextId tcoExpr
  in
    { stmts: StmtEmpty, expr: boxedAction context res, exprType: TypeValue, nextId: res.nextId }

-- ModuleBindings calls this only for values initialized by a cached getter.
-- Match exactly the effect syntax wrap already defers: aliases, factories and
-- EffectDefer retain their existing construction/evaluation boundaries.
wrapCachedInt :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Maybe ExprResult
wrapCachedInt translate context nextId expression = do
  case getExprType expression of
    ADT "Effect.Effect" _ [ Int ] -> pure unit
    _ -> Nothing
  guard (isEffectNode expression)
  pure (wrapIntAction translate context nextId expression)

wrapIntAction :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
wrapIntAction translate context nextId expression =
  let
    effectIntType = ADT "Effect.Effect" [ "Effect", "Effect" ] [ Int ]
    result = translate (context { options = { isTail: false, inEffectBlock: true }, mbExpectedExprType = Just effectIntType, mbExpectedGoType = Nothing }) nextId expression
    -- Translate once, retaining the original boxed body as an independent
    -- entry point. Each invocation executes exactly one of these bodies.
    nativeResult = case result.exprType, result.expr of
      TypeInt64, value -> Just value
      TypeValue, GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ value ] -> Just value
      _, _ -> Nothing
    action = case nativeResult of
      Just value ->
        let native = GoFuncBlock [] (flattenStmts result.stmts <> [ GoReturn value ]) TypeInt64
        in GoCall (GoSelector (GoVar "gopurs_runtime") "WithFunctionData") [ boxedAction context result, native ]
      -- Boxed dynamic results retain the FFI's complete Unbox semantics.
      -- Reuse this translation even when it cannot provide a native entry.
      Nothing -> boxedAction context result
  in { stmts: StmtEmpty, expr: action, exprType: TypeValue, nextId: result.nextId }

boxedAction :: ExprContext -> ExprResult -> GoExpr
boxedAction { codegenStateRef, modNameStr } result =
  GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
    [ GoFuncBlock [ Tuple "_" TypeValue ]
        (flattenStmts result.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr result.expr result.exprType) ])
        TypeValue
    ]
