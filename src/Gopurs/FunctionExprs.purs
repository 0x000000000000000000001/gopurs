module Gopurs.FunctionExprs
  ( abstraction
  , uncurriedAbstraction
  , effectAbstraction
  ) where

import Prelude
import Gopurs.GoSymbol (generated)
import Data.Array as Array
import Data.Array.NonEmpty (NonEmptyArray, toArray)
import Data.Array.NonEmpty as NonEmptyArray
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..), fst)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CallAnalysis (collectCurriedAbs)
import Gopurs.ExprAnalysis (FunctionType, extractExprFuncType, extractFuncType, functionResultAfter, getExprType)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), bindParameters, childContext, flattenStmts)
import Gopurs.FunctionFusion as FunctionFusion
import Gopurs.GoAst (rawGo, GoExpr(..), GoDecl(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoFunctions (Parameters, curriedFunction, hasNativeRecordArgument)
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (Level)

abstraction :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> NonEmptyArray (Tuple (Maybe Ident) Level) -> TcoExpr -> ExprResult
abstraction translate context@{ codegenStateRef, metadata, modNameStr, mbExpectedExprType } nextId tcoExpr args body =
  let
    grouped = collectCurriedAbs args body
    consumed = NonEmptyArray.length grouped.args
    mbFuncTy = functionTypeOrExpected (extractFuncType tcoExpr) mbExpectedExprType

    -- A computation may still separate us from further lambdas.
    -- Consume only the parameters actually collected, not the full
    -- flattened function type's argument list.
    mbBodyType = case mbFuncTy of
      Just signature | consumed <= Array.length signature.fArgs -> Just (functionResultAfter consumed signature)
      _ -> Nothing

    paramsWithTypes = boxedParameters (toArray grouped.args)
    params = map fst paramsWithTypes
    resBody = translate (closureContext context paramsWithTypes mbBodyType) nextId grouped.body

    -- Anonymous curried closures group by five. Do not substitute the module
    -- wrapper's ten-argument policy or cross a computation between lambdas.
    buildFunc :: Array String -> GoExpr -> GoExpr
    buildFunc ps innerExpr =
      let
        len = Array.length ps
        bodyStmts = case innerExpr of
          GoBlock stmts -> stmts
          _ -> [ GoReturn innerExpr ]
      in
        if len == 1 then
          GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoFuncBlock [ Tuple (fromMaybe "" (Array.index ps 0)) TypeValue ] bodyStmts TypeValue ]
        else if len >= 2 && len <= 5 then
          let
            goParams = map (\p -> Tuple p TypeValue) ps
          in
            GoCall (GoSelector (GoVar "gopurs_runtime") ("Func" <> show len)) [ GoFuncBlock goParams bodyStmts TypeValue ]
        else
          let
            chunk = Array.take 5 ps
            rest = Array.drop 5 ps
          in
            buildFunc chunk (buildFunc rest innerExpr)

    funcExpr = buildFunc params (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ]))

    -- A second entry point uses the exact TAST argument representation. Its
    -- result stays Value so consumers polymorphic in that result can use it.
    -- Generate the body with native bindings: projections need no whole-record
    -- box, and escaping uses receive ordinary immutable values as usual.
    nativeTypes = case mbFuncTy of
      Just signature | consumed <= Array.length signature.fArgs ->
        map (exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr)
          (Array.take consumed signature.fArgs)
      _ -> []
  in
    if consumed <= 5 && hasNativeRecordArgument nativeTypes then
      let
        nativeParams = Array.zipWith Tuple params nativeTypes
        nativeBody = translate (closureContext context nativeParams mbBodyType) resBody.nextId grouped.body
        nativeFunction = GoFuncLit nativeParams (flattenStmts nativeBody.stmts)
          (boxGoExpr codegenStateRef modNameStr nativeBody.expr nativeBody.exprType) TypeValue
        expr = GoCall (GoSelector (GoVar "gopurs_runtime") "WithFunctionData") [ funcExpr, nativeFunction ]
      in
        { stmts: StmtEmpty, expr, exprType: TypeValue, nextId: nativeBody.nextId }
    else if consumed == 1 && nativeTypes == [ TypeInt64 ] && mbBodyType == Just Int then
      let
        nativeParams = Array.zipWith Tuple params nativeTypes
        nativeBody = translate ((closureContext context nativeParams mbBodyType) { mbExpectedGoType = Just TypeInt64 }) resBody.nextId grouped.body
        nativeFunction = GoFuncLit nativeParams (flattenStmts nativeBody.stmts)
          (coerceGoExpr codegenStateRef modNameStr nativeBody.expr nativeBody.exprType TypeInt64) TypeInt64
        expr = GoCall (GoSelector (GoVar "gopurs_runtime") "WithFunctionData") [ funcExpr, nativeFunction ]
      in
        { stmts: StmtEmpty, expr, exprType: TypeValue, nextId: nativeBody.nextId }
    else if consumed == 2 then
      case FunctionFusion.countedWorkerCallCount context params grouped.body of
        Just countExpr ->
          { stmts: StmtEmpty
          , expr: GoCall (GoSelector (GoVar "gopurs_runtime") "CountedIntIterator") [ countExpr, funcExpr ]
          , exprType: TypeValue
          , nextId: resBody.nextId
          }
        Nothing -> case FunctionFusion.countedCompositionOperands context params grouped.body of
          Just (Tuple mExpr nExpr) ->
            { stmts: StmtEmpty
            , expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ComposeCountedIntIterators") [ mExpr, nExpr, funcExpr ]
            , exprType: TypeValue
            , nextId: resBody.nextId
            }
          Nothing ->
            { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
    else
      { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }

uncurriedAbstraction :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array (Tuple (Maybe Ident) Level) -> TcoExpr -> ExprResult
uncurriedAbstraction translate context@{ codegenStateRef, modNameStr, tcoIdent, mbExpectedExprType } nextId tcoExpr args body =
  let
    mbFuncTy = functionTypeOrExpected (extractExprFuncType (getExprType tcoExpr)) mbExpectedExprType
    paramsWithTypes = boxedParameters args
    resBody = translate (closureContext context paramsWithTypes (map _.fRet mbFuncTy)) nextId body
    arity = Array.length args
  in
    if arity >= 2 && arity <= 10 then
      case tcoIdent of
        Just topName ->
          let
            callFuncDecl = GoFunctionDecl { name: generated modNameStr ("Call_" <> modNameStr <> "_" <> topName), params: paramsWithTypes, result: TypeValue, body: GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ]) }
            funcExpr = unsafePerformEffect do
              Ref.modify_ (\r -> r { declarations = Array.snoc r.declarations callFuncDecl }) codegenStateRef
              pure $ GoCall (GoSelector (GoVar "gopurs_runtime") ("Func" <> show arity)) [ GoVar ("Call_" <> modNameStr <> "_" <> topName) ]
          in
            { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
        Nothing ->
          let
            funcExpr = GoCall (GoSelector (GoVar "gopurs_runtime") ("Func" <> show arity))
              [ GoFuncLit paramsWithTypes (flattenStmts resBody.stmts)
                  (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType)
                  TypeValue
              ]
          in
            { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
    else
      let
        funcExpr = curriedFunction paramsWithTypes TypeValue
          (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ]))
      in
        { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }

effectAbstraction :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array (Tuple (Maybe Ident) Level) -> TcoExpr -> ExprResult
effectAbstraction translate context@{ codegenStateRef, modNameStr, mbExpectedExprType } nextId tcoExpr args body =
  let
    mbFuncTy = functionTypeOrExpected (extractExprFuncType (getExprType tcoExpr)) mbExpectedExprType
    paramsWithTypes = boxedParameters args
    resBody = translate (closureContext context paramsWithTypes (map _.fRet mbFuncTy)) nextId body
    arity = Array.length args
  in
    if arity >= 2 && arity <= 5 then
      let
        funcExpr = GoCall (GoSelector (GoVar "gopurs_runtime") ("Func" <> show arity))
          [ GoFuncLit paramsWithTypes (flattenStmts resBody.stmts)
              (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType, rawGo "gopurs_runtime.Value{}" ])
              TypeValue
          ]
      in
        { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }

    else
      let
        funcExpr = curriedFunction paramsWithTypes TypeValue
          (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType, rawGo "gopurs_runtime.Value{}" ]) ]))
      in
        { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }

functionTypeOrExpected :: Maybe FunctionType -> Maybe ExprType -> Maybe FunctionType
functionTypeOrExpected actual expected = case actual of
  Just signature -> Just signature
  Nothing -> expected >>= extractExprFuncType

boxedParameters :: Array (Tuple (Maybe Ident) Level) -> Parameters
boxedParameters = map (\(Tuple ident level) -> Tuple (localId ident level) TypeValue)

-- Capture the outer lexical bindings, but never inherit its loop targets: a
-- deferred function body cannot jump into the loop that created the closure.
closureContext :: ExprContext -> Parameters -> Maybe ExprType -> ExprContext
closureContext context params expected = (childContext context expected)
  { bound = bindParameters params context.bound
  , mbExpectedGoType = Just TypeValue
  , options = { isTail: context.options.isTail, inEffectBlock: false }
  }
