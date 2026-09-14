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
import Gopurs.CallAnalysis (CallTarget, collectGoSpine, getGoSpineArgs, curriedTarget, qualifiedTarget)
import Gopurs.CallArguments (Arguments, applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.ExprAnalysis (getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, LoopTarget, TranslateExpr, StmtTree(..))
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, sanitizeName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
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
      Nothing ->
        let target = curriedTarget context.bound fn
        in case ArrayIntrinsics.recognize ArrayIntrinsics.Curried context.modNameStr target (Array.length args) of
          Just intrinsic ->
            let translated = CallArguments.translateBoxed translate context { stmts: StmtEmpty, nextId } args
            in ArrayIntrinsics.emitCurried context fn args intrinsic translated
          Nothing -> case directFunction context target (Array.length args) of
            Just info ->
              let translated = CallArguments.translate translate context Nothing { stmts: StmtEmpty, nextId } args
              in nativeCall context (GoVar info.fullName) info.fArgs info.fRet info.arity translated
            Nothing -> curriedCall translate context nextId fn args

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
curriedCall translate context@{ codegenStateRef, modNameStr } nextId fn args =
  let
    resFn = translate (CallArguments.childContext context Nothing) nextId fn
    -- Dynamic calls consume boxed records. A native function instead supplies
    -- its expected argument types when the translated values are coerced.
    expected = case resFn.exprType of
      TypeFunc _ _ -> Nothing
      _ -> Just Any
    translated = CallArguments.translate translate context expected { stmts: resFn.stmts, nextId: resFn.nextId } args
  in
    case resFn.exprType of
      TypeFunc fArgs fRet | Array.length args >= Array.length fArgs ->
        nativeCall context resFn.expr fArgs fRet (Array.length fArgs) translated
      _ ->
        let
          boxedArgs = CallArguments.boxRemaining context 0 translated
          expr = applyBoxed (boxGoExpr codegenStateRef modNameStr resFn.expr resFn.exprType) boxedArgs
        in
          { stmts: translated.stmts, expr, exprType: TypeValue, nextId: translated.nextId }

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
  case ArrayIntrinsics.recognize ArrayIntrinsics.Uncurried context.modNameStr (qualifiedTarget fn) (Array.length args) of
    Just intrinsic ->
      let translated = CallArguments.translate translate context Nothing { stmts: StmtEmpty, nextId } args
      in ArrayIntrinsics.emitUncurried context args intrinsic translated
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
    resFn = translate (CallArguments.childContext context Nothing) nextId fn
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
    translated = CallArguments.translate translate context Nothing { stmts: StmtEmpty, nextId } args
    assigns = Array.mapWithIndex
      (\index param -> GoMutate param (coerceGoExpr codegenStateRef modNameStr
        (fromMaybe (GoRaw "nil") (Array.index translated.exprs index))
        (fromMaybe TypeValue (Array.index translated.exprTypes index))
        (fromMaybe TypeValue (Array.index target.goTypes index))))
      target.loopParams
    expectedGoType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr
      (case getExprType expression of
        Any -> fromMaybe Any mbExpectedExprType
        ty -> ty)
  in
    { stmts: translated.stmts <> foldMap StmtLeaf assigns <> StmtLeaf (GoContinue target.ident)
    , expr: GoRaw ("func() " <> goTypeToStr expectedGoType <> " { panic(\"unreachable\") }()")
    , exprType: expectedGoType
    , nextId: translated.nextId
    }
