module Gopurs.LocalWorkers
  ( parameters
  , signature
  , emit
  ) where

import Prelude
import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.Tuple (Tuple(..), snd)
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.ExprAnalysis (extractExprFuncType, getExprType)
import Gopurs.ExprContext (ExprContext, ExprResult, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, rawGo)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoFunctions (Parameters, curriedFunction, iterationBindings, loopParameters, namedParameters)
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)

-- Local workers use annotation slots directly. Module workers have a separate
-- ABI policy, including read-only record projections and native sum results.
parameters :: ExprContext -> Array String -> TcoExpr -> Parameters
parameters { metadata, modNameStr } names value =
  let
    types = case extractExprFuncType (getExprType value) of
      Just { fArgs } -> map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) fArgs
      Nothing -> []
  in
    namedParameters names types

workerName :: String -> String -> String
workerName modName name = "Call_local_" <> modName <> "_" <> name

signature :: ExprContext -> String -> Parameters -> GoType -> FunctionInfo
signature { modNameStr } name params result =
  { fullName: workerName modNameStr name, fArgs: map snd params, fRet: result, arity: Array.length params }

-- The translated body determines the native result. Only the curried wrapper
-- crosses Value. Declarations stay separate so recursive callers can place the
-- whole group's declarations before any assignment captures a peer worker.
emit :: ExprContext -> String -> Parameters -> Boolean -> ExprResult -> { declarations :: Array GoExpr, assignments :: Array GoExpr }
emit { codegenStateRef, modNameStr } name params isLoop result =
  let
    nativeName = workerName modNameStr name
    resultType = result.exprType
    declarations =
      [ rawGo ("var " <> nativeName <> " func(" <> String.joinWith ", " (map (goTypeToStr <<< snd) params) <> ") " <> goTypeToStr resultType)
      , rawGo ("_ = " <> nativeName)
      , rawGo ("var " <> name <> " gopurs_runtime.Value")
      , rawGo ("_ = " <> name)
      ]
    body = iterationBindings params <> flattenStmts result.stmts <> [ GoReturn result.expr ]
    functionBody = if isLoop then GoFor name body else GoBlock body
    nativeAssignment = GoMutate nativeName (GoFuncBlock (loopParameters params) [ functionBody ] resultType)
    call = GoCall (GoVar nativeName) (map (\(Tuple param ty) ->
      coerceGoExpr codegenStateRef modNameStr (GoVar (param <> "_loop_val")) TypeValue ty) params)
    boxedResult = boxGoExpr codegenStateRef modNameStr call resultType
    wrapper = if Array.null params then
      curriedFunction [ Tuple "_" TypeValue ] TypeValue (GoBlock [ GoReturn boxedResult ])
      else Array.foldr (\(Tuple param _) inner ->
        GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
          [ GoFuncLit [ Tuple (param <> "_loop_val") TypeValue ] [] inner TypeValue ]) boxedResult params
  in
    { declarations, assignments: [ nativeAssignment, GoMutate name wrapper ] }
