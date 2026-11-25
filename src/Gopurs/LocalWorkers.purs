module Gopurs.LocalWorkers
  ( parameters
  , signature
  , needsWrapper
  , emit
  ) where

import Prelude
import Data.Array as Array
import Data.Foldable (any)
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.Tuple (Tuple(..), snd)
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.ExprAnalysis (extractExprFuncType, getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, rawGo)
import Gopurs.GoSymbol (GoReference(..))
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoFunctions (Parameters, curriedFunction, iterationBindings, loopParameters, namedParameters)
import Gopurs.Int32Loops as Int32Loops
import Gopurs.AffineLoops as AffineLoops
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(App, Local))

-- Local workers use annotation slots directly. Module workers have a separate
-- ABI policy, including read-only record projections and native sum results.
parameters :: ExprContext -> Array String -> TcoExpr -> Parameters
parameters { metadata, modNameStr } names value =
  let
    types = case extractExprFuncType (getExprType value) of
      Just { fArgs } -> map (exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) fArgs
      Nothing -> []
  in
    namedParameters names types

workerName :: String -> String -> String
workerName modName name = "Call_local_" <> modName <> "_" <> name

signature :: ExprContext -> String -> Parameters -> GoType -> FunctionInfo
signature { modNameStr } name params result =
  { reference: LocalReference (workerName modNameStr name), fArgs: map snd params, fRet: result, arity: Array.length params }

-- Inspect the current TCO expression, not usage annotations from an earlier
-- IR. Saturated curried calls use the published native worker directly.
-- Every other occurrence keeps the wrapper, including partial applications,
-- returned functions, callback arguments and uses in peer closures.
needsWrapper :: String -> Int -> TcoExpr -> Boolean
needsWrapper target arity expression =
  let
    visit = needsWrapper target arity
    syntax = unwrapTcoExpr expression
  in case syntax of
    Local ident level -> localId ident level == target
    App _ _ ->
      let
        Tuple fn spine = collectGoSpine expression
        args = getGoSpineArgs spine
      in case unwrapTcoExpr fn of
        Local ident level | localId ident level == target && arity > 0 && Array.length args >= arity && onlyCurriedApplications expression -> any visit args
        _ -> any visit syntax
    _ -> any visit syntax

-- collectGoSpine also flattens uncurried applications. Keep their wrappers,
-- including when an uncurried application is nested under a curried one.
onlyCurriedApplications :: TcoExpr -> Boolean
onlyCurriedApplications expression = case unwrapTcoExpr expression of
  App fn _ -> onlyCurriedApplications fn
  Local _ _ -> true
  _ -> false

-- The translated body determines the native result. Only the curried wrapper
-- crosses Value. Declarations stay separate so recursive callers can place the
-- whole group's declarations before any assignment captures a peer worker.
emit :: ExprContext -> String -> Parameters -> Boolean -> Maybe TcoExpr -> ExprResult -> { declarations :: Array GoExpr, assignments :: Array GoExpr }
emit { codegenStateRef, modNameStr, bound } name params keepWrapper loopBody result =
  let
    nativeName = workerName modNameStr name
    resultType = result.exprType
    declarations =
      [ rawGo ("var " <> nativeName <> " func(" <> String.joinWith ", " (map (goTypeToStr <<< snd) params) <> ") " <> goTypeToStr resultType)
      , rawGo ("_ = " <> nativeName)
      ]
        <> if keepWrapper then
          [ rawGo ("var " <> name <> " gopurs_runtime.Value")
          , rawGo ("_ = " <> name)
          ]
        else []
    body = flattenStmts result.stmts <> [ GoReturn result.expr ]
    functionBody = case loopBody of
      Just source -> AffineLoops.wrap (AffineLoops.localTarget bound name) params resultType source
        (Int32Loops.loop name params resultType body)
      Nothing -> GoBlock (iterationBindings params <> body)
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
    { declarations, assignments: [ nativeAssignment ] <> if keepWrapper then [ GoMutate name wrapper ] else [] }
