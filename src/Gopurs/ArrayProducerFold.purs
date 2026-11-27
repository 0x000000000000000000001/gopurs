module Gopurs.ArrayProducerFold (emit, onlyCurried) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (evalState)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIndexedFold (readOnlyScalar)
import Gopurs.IndexedAddLanes as IndexedAddLanes
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.CallAnalysis (collectCurriedAbs, collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.ClosedIndexedFolds (closedPrimitiveCallback, validDescriptor)
import Gopurs.CodegenState (ArrayFoldDescriptor)
import Gopurs.ExprAnalysis (extractFuncType, getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, bindParameters, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.ImmediateApplications.Scope (freshen, maximumLevel)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.CoreFn (ExprType(Array, Int), ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

-- Only an immediate producer is removed. Both callbacks are closed primitive
-- scalar lambdas; unknown calls, captures and shared produced arrays keep their
-- original path. The input buffers are read without mutation or ownership claims.
emit :: TranslateExpr -> ExprContext -> Int -> ArrayFoldDescriptor -> TcoExpr -> Maybe ExprResult
emit translate context nextId descriptor producer = do
  guard (Array.null context.recVars && validDescriptor descriptor && onlyCurried producer)
  guard (getExprType producer == Array Int)
  let Tuple function spine = collectGoSpine producer
  target <- qualifiedTarget function
  let arguments = getGoSpineArgs spine
  plan <- if target.mbMod == Just (ModuleName "Data.Semigroup") && target.name == "concatArray" then do
    guard (Array.length arguments == 2)
    pure { mapped: false, arguments }
    else if target.mbMod == Just (ModuleName "Data.Functor") && target.name == "arrayMap" then do
      guard (Array.length arguments == 2)
      callback <- Array.head arguments
      signature <- extractFuncType callback
      guard (signature.fArgs == [ Int ] && signature.fRet == Int && closedPrimitiveCallback 1 (neutral callback))
      case unwrapTcoExpr callback of
        Abs parameters body -> do
          let grouped = collectCurriedAbs parameters body
          guard (NEA.length grouped.args == 1 && readOnlyScalar (map (\(Tuple _ level) -> level) (NEA.toArray grouped.args)) grouped.body)
          pure { mapped: true, arguments }
        _ -> Nothing
      else Nothing
  let
    -- Imported levels are alpha-renamed independently for every insertion.
    -- The closed callback has no free references to reinterpret in this module.
    start = max (maximumLevel (neutral producer)) (maximumLevel descriptor.callback) + nextId + 1
    foldedCallback = Tco.analyze [] (evalState (freshen descriptor.callback) start)
    nativeLambda currentId callback = case unwrapTcoExpr callback of
      Abs parameters callbackBody ->
        let
          grouped = collectCurriedAbs parameters callbackBody
          params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) (NEA.toArray grouped.args)
          callbackContext = context
            { depth = context.depth + 1, bound = bindParameters params context.bound
            , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
            , mbExpectedExprType = Just Int, mbExpectedGoType = Just TypeInt64 }
          result = translate callbackContext currentId grouped.body
        in { stmts: StmtEmpty, expr: GoFuncLit params (flattenStmts result.stmts) (integer result.expr result.exprType) TypeInt64
           , exprType: TypeFunc (map (const TypeInt64) params) TypeInt64, nextId: result.nextId }
      _ -> { stmts: StmtEmpty, expr: rawGo "nil", exprType: TypeValue, nextId: currentId }
    -- An adjacent closed mapper is captured before its input. Concat captures
    -- both input expressions before either source buffer is inspected.
    captured = CallArguments.capture { prefix: "arrayProducer_arg_", markUsed: false }
      (\index currentId argument -> if plan.mapped && index == 0 then nativeLambda currentId argument
        else translate (childContext context Nothing) currentId argument) nextId plan.arguments
    folding = nativeLambda captured.nextId foldedCallback
    suffix = show folding.nextId
    variable label = "arrayProducer_" <> label <> "_" <> suffix
    callbackName = variable "fold"
    resultName = variable "result"
    indexName = variable "index"
    valueName = variable "value"
    offsetName = variable "offset"
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    source index =
      let inputName = case arg index of
            GoVar name -> name
            _ -> ""
      in arraySource inputName (variable ("source" <> show index)) (fromMaybe TypeValue (Array.index captured.exprTypes index))
    left = source (if plan.mapped then 1 else 0)
    right = source 1
    loop sourceInfo offset = GoForRange (indexName <> ", " <> valueName <> " := range " <> sourceInfo.target)
      [ GoMutate resultName (GoCall (GoVar callbackName)
          [ GoCall (GoVar "int64") [ GoBinOp "+" offset (GoVar indexName) ]
          , GoVar resultName
          , if plan.mapped then GoCall (arg 0) [ integer (GoVar valueName) sourceInfo.elementType ]
            else integer (GoVar valueName) sourceInfo.elementType
          ]) ]
    mappedLoop = case Array.head plan.arguments of
      Just mapper | plan.mapped && IndexedAddLanes.eligible mapper foldedCallback ->
        IndexedAddLanes.emit { source: left.target, prefix: variable "lane", result: resultName }
          (\i accumulator value -> GoCall (GoVar callbackName)
            [ GoCall (GoVar "int64") [ i ], accumulator, GoCall (arg 0) [ integer value left.elementType ] ])
          (loop left (GoInt 0))
      _ -> loop left (GoInt 0)
    body = [ GoAssign callbackName folding.expr, GoAssign resultName (GoCall (GoVar "int64") [ GoInt descriptor.seed ]), left.assignment ] <>
      (if plan.mapped then [ mappedLoop ] else
        [ right.assignment, GoAssign offsetName (GoCall (GoVar "len") [ rawGo left.target ])
        , loop left (GoInt 0), loop right (GoVar offsetName) ])
  pure { stmts: captured.stmts, expr: GoCall (GoFuncLit [] body (GoVar resultName) TypeInt64) []
       , exprType: TypeInt64, nextId: folding.nextId + 1 }
  where
  integer expression from = coerceGoExpr context.codegenStateRef context.modNameStr expression from TypeInt64

neutral :: TcoExpr -> NeutralExpr
neutral (TcoExpr _ syntax) = NeutralExpr (map neutral syntax)

-- An uncurried/staged invocation is a separate observable boundary. The broad
-- spine collector must not erase it when recognizing a saturated producer.
onlyCurried :: TcoExpr -> Boolean
onlyCurried (TcoExpr _ syntax) = case syntax of
  App function _ -> onlyCurried function
  Typed _ value -> onlyCurried value
  TypeApp value _ -> onlyCurried value
  UncurriedApp _ _ -> false
  UncurriedEffectApp _ _ -> false
  _ -> true
