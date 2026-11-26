module Gopurs.CallExprs
  ( application
  , uncurriedApplication
  , effectApplication
  ) where

import Prelude
import Data.Array as Array
import Data.Foldable (foldMap)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIntrinsics as ArrayIntrinsics
import Gopurs.ArrayConcat as ArrayConcat
import Gopurs.ArrayIntrinsics.Callback as Callback
import Gopurs.ArrayTraverse as ArrayTraverse
import Gopurs.ArrayIndexedFold as ArrayIndexedFold
import Gopurs.ArrayEquality as ArrayEquality
import Gopurs.ScalarArrayFold as ScalarArrayFold
import Gopurs.ArrayRangeFold as ArrayRangeFold
import Gopurs.ArraySort as ArraySort
import Gopurs.ObjectTraverse as ObjectTraverse
import Gopurs.TailRecEffect as TailRecEffect
import Gopurs.PrivateSTRegions as PrivateSTRegions
import Gopurs.CallAnalysis (CallTarget, collectGoSpine, getGoSpineArgs, curriedTarget, qualifiedTarget)
import Gopurs.CallArguments (Arguments, applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.ExprAnalysis (getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, LoopTarget, TranslateExpr, StmtTree(..), childContext)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..), goTypeToStr, sanitizeName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoFunctions (hasNativeRecordArgument, nativeCallbackCall)
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

application :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
application translate context nextId expression =
  let
    Tuple fn spine = collectGoSpine expression
    args = getGoSpineArgs spine
  in
    -- Curried calls give a tail jump priority over intrinsic recognition.
    case tailTarget context fn of
      Just target -> tailCall translate context nextId expression args target
      Nothing -> case ObjectTraverse.emit translate context nextId (qualifiedTarget fn) args of
        Just result -> result
        Nothing -> case TailRecEffect.emit translate context nextId fn args of
          Just result -> result
          Nothing -> case PrivateSTRegions.emit translate context nextId fn args of
            Just result -> result
            Nothing -> regularApplication fn args
  where
  regularApplication fn args = case ArraySort.emit translate context nextId fn args of
    Just result -> result
    Nothing -> scalarApplication fn args

  scalarApplication fn args = case scalarFold translate context nextId fn args of
    Just result -> result
    Nothing -> arrayApplication fn args

  arrayApplication fn args = case ArrayEquality.emit translate context nextId fn args of
    Just result -> result
    Nothing -> case ArrayTraverse.emit translate context nextId fn args of
        Just result -> result
        Nothing ->
          let target = curriedTarget context.bound fn
          in case ArrayIntrinsics.recognize ArrayIntrinsics.Curried context.modNameStr target (Array.length args) of
            Just intrinsic ->
              ArrayIntrinsics.emitCurried translate context nextId fn args intrinsic
            Nothing -> case directFunction context target (Array.length args) of
              Just info ->
                let translated = directArguments translate context nextId info.fArgs args
                in nativeCall context (GoRef info.reference) info.fArgs info.fRet info.arity translated
              Nothing -> curriedCall translate context nextId fn args

scalarFold :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
scalarFold translate context nextId fn args = case ArrayRangeFold.emit translate context nextId fn args of
  Just fused -> Just fused
  Nothing -> case ScalarArrayFold.emit translate context nextId fn args of
    Just folded -> Just folded
    Nothing -> ArrayIndexedFold.emit translate context nextId fn args

-- Only a published native predicate slot enables lambda lowering here. Other
-- calls retain their current translation. Capture every argument in source
-- order before entering the worker, including boxed callback factories.
directArguments :: TranslateExpr -> ExprContext -> Int -> Array GoType -> Array TcoExpr -> Arguments
directArguments translate context@{ codegenStateRef, modNameStr } nextId expected args =
  let predicateType = TypeFunc [ TypeInt64 ] TypeBool
  in if Array.elem predicateType expected then
    CallArguments.capture { prefix: "nativePredicate_arg_", markUsed: false }
      (\index argumentId arg ->
        let
          native = if Array.index expected index == Just predicateType && getExprType arg == Func [ Int ] Boolean then
            Callback.nativeLambda translate context argumentId 1 arg
            else Nothing
          result = case native of
            Just callback@{ expr: GoFuncLit params statements returned resultType } ->
              callback
                { expr = GoFuncLit params statements (coerceGoExpr codegenStateRef modNameStr returned resultType TypeBool) TypeBool
                , exprType = predicateType }
            _ -> translate (childContext context Nothing) argumentId arg
        in if Array.index expected index == Just predicateType then
          result { expr = coerceGoExpr codegenStateRef modNameStr result.expr result.exprType predicateType, exprType = predicateType }
          else result)
      nextId args
    -- A unary native consumer has no later argument evaluation across which
    -- array materialization could move. In particular concat's empty-side
    -- alias must not be copied early relative to another argument factory.
    else if expected == [ TypeNativeArray TypeInt64 ] && Array.length args == 1 then CallArguments.capture { prefix: "nativeArray_arg_", markUsed: false }
      (\index argumentId arg ->
        let native = if Array.index expected index == Just (TypeNativeArray TypeInt64) then
              ArrayConcat.nativeArgument translate context argumentId arg
              else Nothing
        in case native of
          Just result -> result
          Nothing -> translate ((childContext context Nothing)
            { mbExpectedGoType = Just (TypeNativeArray TypeInt64) }) argumentId arg)
      nextId args
    else CallArguments.translate translate context Nothing { stmts: StmtEmpty, nextId } args

-- A direct worker must be saturated; partial calls retain the boxed path.
directFunction :: ExprContext -> Maybe CallTarget -> Int -> Maybe FunctionInfo
directFunction { metadata, modNameStr, moduleFunctions } target count = do
  { mbMod, name } <- target
  let isLocal = map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) mbMod == Just modNameStr || mbMod == Nothing
  info <- if isLocal then Map.lookup name moduleFunctions
    else do
      mn <- mbMod
      Map.lookup (unwrap mn <> "." <> name) metadata.globalFunctions
  if count >= info.arity && info.arity >= 1 then Just info else Nothing

curriedCall :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> ExprResult
curriedCall translate context@{ codegenStateRef, metadata, modNameStr } nextId fn args =
  let
    resFn = translate (childContext context Nothing) nextId fn
    argumentTypes = map (exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr <<< getExprType) args
  in
    case resFn.exprType of
      TypeFunc fArgs fRet | Array.length args >= Array.length fArgs ->
        let translated = CallArguments.translate translate context Nothing { stmts: resFn.stmts, nextId: resFn.nextId } args
        in nativeCall context resFn.expr fArgs fRet (Array.length fArgs) translated
      _ | Array.length args <= 5 && hasNativeRecordArgument argumentTypes ->
        recordCallbackCall translate context resFn args
      _ ->
        let
          translated = CallArguments.translate translate context (Just Any) { stmts: resFn.stmts, nextId: resFn.nextId } args
          boxedArgs = CallArguments.boxRemaining context 0 translated
          expr = applyBoxed (boxGoExpr codegenStateRef modNameStr resFn.expr resFn.exprType) boxedArgs
        in
          { stmts: translated.stmts, expr, exprType: TypeValue, nextId: translated.nextId }

recordCallbackCall :: TranslateExpr -> ExprContext -> ExprResult -> Array TcoExpr -> ExprResult
recordCallbackCall translate context@{ codegenStateRef, modNameStr } fn args =
  let
    fnName = "__callback_fn_" <> show fn.nextId
    captured = CallArguments.capture { prefix: "__callback_arg_", markUsed: false }
      (\_ nextId arg ->
        let result = translate (childContext context Nothing) nextId arg
        -- A Go short declaration would default an untyped numeric literal to
        -- int/float64. Preserve its declared native ABI before capturing it.
        in result { expr = case result.exprType of
          TypeInt64 -> GoCall (GoVar "int64") [ result.expr ]
          TypeFloat64 -> GoCall (GoVar "float64") [ result.expr ]
          TypeUint32 -> GoCall (GoVar "uint32") [ result.expr ]
          _ -> result.expr
        })
      (fn.nextId + 1) args
    fnExpr = GoVar fnName
    fallback = applyBoxed fnExpr (CallArguments.boxRemaining context 0 captured)
    expr = if hasNativeRecordArgument captured.exprTypes then
      nativeCallbackCall fnExpr captured.exprTypes captured.exprs fallback
      else fallback
  in
    { stmts: fn.stmts
        <> StmtLeaf (GoAssign fnName (boxGoExpr codegenStateRef modNameStr fn.expr fn.exprType))
        <> captured.stmts
    , expr
    , exprType: TypeValue
    , nextId: captured.nextId
    }

-- Coerce the native prefix and apply any surplus arguments to its boxed
-- result. Boxing surplus arguments stays before prefix coercion, as before.
nativeCall :: ExprContext -> GoExpr -> Array GoType -> GoType -> Int -> Arguments -> ExprResult
nativeCall context@{ codegenStateRef, modNameStr } fn expected ret arity args =
  let
    remaining = CallArguments.boxRemaining context arity args
    callArgs = CallArguments.coercePrefix context arity expected args
    call = GoCall fn callArgs
    expr = if Array.null remaining then call else applyBoxed (boxGoExpr codegenStateRef modNameStr call ret) remaining
    exprType = if Array.null remaining then ret else TypeValue
  in
    { stmts: args.stmts, expr, exprType, nextId: args.nextId }

uncurriedApplication :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> TcoExpr -> Array TcoExpr -> ExprResult
uncurriedApplication translate context nextId expression fn args =
  -- Uncurried calls recognize intrinsics before considering a tail jump.
  case ArraySort.emit translate context nextId fn args of
    Just result -> result
    Nothing -> uncurriedScalarApplication translate context nextId expression fn args

uncurriedScalarApplication :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> TcoExpr -> Array TcoExpr -> ExprResult
uncurriedScalarApplication translate context nextId expression fn args =
  case scalarFold translate context nextId fn args of
    Just result -> result
    Nothing -> uncurriedArrayApplication translate context nextId expression fn args

uncurriedArrayApplication :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> TcoExpr -> Array TcoExpr -> ExprResult
uncurriedArrayApplication translate context nextId expression fn args =
  case ArrayEquality.emit translate context nextId fn args of
    Just result -> result
    Nothing -> case ArrayIntrinsics.safeIndex translate context nextId (qualifiedTarget fn) args of
      Just result -> result
      Nothing -> case ArrayIntrinsics.recognize ArrayIntrinsics.Uncurried context.modNameStr (qualifiedTarget fn) (Array.length args) of
        Just intrinsic ->
          ArrayIntrinsics.emitUncurried translate context nextId args intrinsic
        Nothing ->
          let
            Tuple flatFn spine = collectGoSpine expression
            flatArgs = getGoSpineArgs spine
          in
            case tailTarget context flatFn of
              Just target -> tailCall translate context nextId expression flatArgs target
              Nothing -> uncurriedCall translate context 10 nextId fn args

effectApplication :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> ExprResult
effectApplication translate context = uncurriedCall translate context 5

-- Effect applications and ordinary uncurried applications differ only in
-- the largest runtime wrapper they select. Both box their final result.
uncurriedCall :: TranslateExpr -> ExprContext -> Int -> Int -> TcoExpr -> Array TcoExpr -> ExprResult
uncurriedCall translate context@{ codegenStateRef, modNameStr } maxArity nextId fn args =
  let
    resFn = translate (childContext context Nothing) nextId fn
    translated = CallArguments.translate translate context Nothing { stmts: resFn.stmts, nextId: resFn.nextId } args
    len = Array.length args
    goFuncName = if len >= 2 && len <= maxArity then "UncurriedApp" <> show len else "UncurriedApp"
    expr = case resFn.exprType of
      TypeFunc fArgs fRet | Array.length fArgs == len ->
        let callArgs = CallArguments.coercePrefix context len fArgs translated
        in boxGoExpr codegenStateRef modNameStr (GoCall resFn.expr callArgs) fRet
      _ ->
        let boxedArgs = CallArguments.boxRemaining context 0 translated
        in GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName)
          (Array.cons (boxGoExpr codegenStateRef modNameStr resFn.expr resFn.exprType) boxedArgs)
  in
    { stmts: translated.stmts, expr, exprType: TypeValue, nextId: translated.nextId }

tailTarget :: ExprContext -> TcoExpr -> Maybe LoopTarget
tailTarget { bound, loopCtx, options: { isTail } } fn =
  if isTail then case unwrapTcoExpr fn of
    Local mbIdent lvl ->
      let binding = fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)
      in Array.find (\target -> target.ident == binding.name) loopCtx
    Var (Qualified _ (Ident name)) -> Array.find (\target -> target.ident == sanitizeName name) loopCtx
    _ -> Nothing
  else Nothing

tailCall :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> LoopTarget -> ExprResult
tailCall translate context@{ metadata, codegenStateRef, modNameStr, mbExpectedExprType } nextId expression args target =
  let
    -- Translate all operands outside tail position before updating the loop
    -- slots. Their references use the current iteration's immutable parameters
    -- (GoFunctions.iterationBindings), not slots changed by earlier assignments.
    translated = CallArguments.translate translate context Nothing { stmts: StmtEmpty, nextId } args
    assigns = Array.mapWithIndex
      (\index param -> GoMutate param (coerceGoExpr codegenStateRef modNameStr
        (fromMaybe (rawGo "nil") (Array.index translated.exprs index))
        (fromMaybe TypeValue (Array.index translated.exprTypes index))
        (fromMaybe TypeValue (Array.index target.goTypes index))))
      target.loopParams
    expectedGoType = exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr
      (case getExprType expression of
        Any -> fromMaybe Any mbExpectedExprType
        ty -> ty)
  in
    { stmts: translated.stmts <> foldMap StmtLeaf assigns <> StmtLeaf (GoContinue target.ident)
    , expr: rawGo ("func() " <> goTypeToStr expectedGoType <> " { panic(\"unreachable\") }()")
    , exprType: expectedGoType
    , nextId: translated.nextId
    }
