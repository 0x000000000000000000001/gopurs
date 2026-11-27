module Gopurs.ArrayProducerFold (emit, onlyCurried) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (evalState)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIndexedFold (readOnlyScalar, scalarType)
import Gopurs.IndexedAddLanes as IndexedAddLanes
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
import PureScript.Backend.Optimizer.CoreFn (ExprType(Array, Boolean, Int), ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

data ProducerPlan
  = ConcatPlan (Array TcoExpr)
  | MapPlan (Array TcoExpr)
  | FilterPlan (Array TcoExpr)

-- Only an immediate producer is removed. Callbacks are primitive scalar
-- lambdas; unknown calls and shared produced arrays keep their original path.
-- Input buffers are read without mutation or ownership claims.
emit :: TranslateExpr -> ExprContext -> Int -> ArrayFoldDescriptor -> TcoExpr -> Maybe ExprResult
emit translate context nextId descriptor producer = do
  guard (Array.null context.recVars && validDescriptor descriptor)
  guard (getExprType producer == Array Int)
  plan <- recognizeProducer producer
  let
    planArgs = case plan of
      ConcatPlan args -> args
      MapPlan args -> args
      FilterPlan args -> args
    -- Imported levels are alpha-renamed independently for every insertion.
    -- The closed callback has no free references to reinterpret in this module.
    start = max (maximumLevel (neutral producer)) (maximumLevel descriptor.callback) + nextId + 1
    foldedCallback = Tco.analyze [] (evalState (freshen descriptor.callback) start)
    nativeLambda expectedExpr expectedGo currentId callback = case unwrapTcoExpr callback of
      Abs parameters callbackBody ->
        let
          grouped = collectCurriedAbs parameters callbackBody
          params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) (NEA.toArray grouped.args)
          callbackContext = context
            { depth = context.depth + 1, bound = bindParameters params context.bound
            , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
            , mbExpectedExprType = Just expectedExpr, mbExpectedGoType = Just expectedGo }
          result = translate callbackContext currentId grouped.body
        in { stmts: StmtEmpty, expr: GoFuncLit params (flattenStmts result.stmts) (coerceGoExpr context.codegenStateRef context.modNameStr result.expr result.exprType expectedGo) expectedGo
           , exprType: TypeFunc (map (const TypeInt64) params) expectedGo, nextId: result.nextId }
      _ -> { stmts: StmtEmpty, expr: rawGo "nil", exprType: TypeValue, nextId: currentId }
    -- An adjacent callback is captured before its input. Concat captures
    -- both input expressions before either source buffer is inspected.
    captured = CallArguments.capture { prefix: "arrayProducer_arg_", markUsed: false }
      (\index currentId argument -> case plan, index of
        MapPlan _, 0 -> nativeLambda Int TypeInt64 currentId argument
        FilterPlan _, 0 -> nativeLambda Boolean TypeBool currentId argument
        _, _ -> translate (childContext context Nothing) currentId argument) nextId planArgs
    folding = nativeLambda Int TypeInt64 captured.nextId foldedCallback
    contributionInfo = IndexedAddLanes.extractContribution foldedCallback
    contribLambda = case contributionInfo of
      Just info ->
        let
          params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) [ info.index, info.value ]
          cbCtx = context
            { depth = context.depth + 1, bound = bindParameters params context.bound
            , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
            , mbExpectedExprType = Just Int, mbExpectedGoType = Just TypeInt64 }
          res = translate cbCtx folding.nextId info.body
        in Just { expr: GoFuncLit params (flattenStmts res.stmts) (integer res.expr res.exprType) TypeInt64, nextId: res.nextId }
      Nothing -> Nothing
    nextAfterContrib = case contribLambda of
      Just c -> c.nextId
      Nothing -> folding.nextId
    suffix = show nextAfterContrib
    variable label = "arrayProducer_" <> label <> "_" <> suffix
    callbackName = variable "fold"
    contribName = variable "contrib"
    resultName = variable "result"
    indexName = variable "index"
    outIndexName = variable "outIndex"
    itemName = variable "item"
    valueName = variable "value"
    offsetName = variable "offset"
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    argName index = case arg index of
      GoVar name -> name
      _ -> ""
    withIntSource index render =
      let
        inputName = argName index
        baseName = variable ("source" <> show index)
      in case argType index of
        TypeNativeArray inner ->
          [ GoAssign baseName (GoVar inputName) ] <> render { target: baseName, elementType: inner }
        _ ->
          [ GoIfElse (rawGo (inputName <> ".Type == gopurs_runtime.TypeIntArray"))
              ([ GoAssign (baseName <> "_i") (rawGo ("*(*[]int64)(" <> inputName <> ".UnsafePtr)")) ]
                <> render { target: baseName <> "_i", elementType: TypeInt64 })
              ([ GoAssign (baseName <> "_v") (rawGo ("*(*[]gopurs_runtime.Value)(" <> inputName <> ".UnsafePtr)")) ]
                <> render { target: baseName <> "_v", elementType: TypeValue })
          ]
    isMapped = case plan of
      MapPlan _ -> true
      _ -> false
    loop sourceInfo offset = GoForRange (indexName <> ", " <> valueName <> " := range " <> sourceInfo.target)
      [ GoMutate resultName (GoCall (GoVar callbackName)
          [ GoCall (GoVar "int64") [ GoBinOp "+" offset (GoVar indexName) ]
          , GoVar resultName
          , if isMapped then GoCall (arg 0) [ integer (GoVar valueName) sourceInfo.elementType ]
            else integer (GoVar valueName) sourceInfo.elementType
          ]) ]
    mappedLoop sourceInfo = case plan, Array.head planArgs, contribLambda of
      MapPlan _, Just mapper, Just _ | IndexedAddLanes.eligible mapper foldedCallback ->
        [ IndexedAddLanes.emitContribution { source: sourceInfo.target, prefix: variable "lane", result: resultName, offset: Nothing }
            (\i value -> GoCall (GoVar contribName)
              [ GoCall (GoVar "int64") [ i ], GoCall (arg 0) [ integer value sourceInfo.elementType ] ])
            (loop sourceInfo (GoInt 0)) ]
      _, _, _ -> [ loop sourceInfo (GoInt 0) ]
    concatLoop sourceInfo label offsetOpt fallbackOffset = case contribLambda of
      Just _ ->
        [ IndexedAddLanes.emitContribution { source: sourceInfo.target, prefix: variable label, result: resultName, offset: offsetOpt }
            (\i value -> GoCall (GoVar contribName)
              [ GoCall (GoVar "int64") [ i ], integer value sourceInfo.elementType ])
            (loop sourceInfo fallbackOffset) ]
      Nothing -> [ loop sourceInfo fallbackOffset ]
    filterBody sourceInfo = case contribLambda of
      Just _ ->
        [ GoAssign outIndexName (rawGo "int64(0)")
        , GoForRange ("_, " <> valueName <> " := range " <> sourceInfo.target)
            [ GoAssign itemName (integer (GoVar valueName) sourceInfo.elementType)
            , GoIfElse (GoCall (arg 0) [ GoVar itemName ])
                [ GoMutate resultName (GoBinOp "+" (GoVar resultName) (GoCall (GoVar contribName) [ GoVar outIndexName, GoVar itemName ]))
                , GoMutate outIndexName (GoBinOp "+" (GoVar outIndexName) (rawGo "int64(1)"))
                ] []
            ]
        , GoIfElse (GoBinOp ">" (GoVar outIndexName) (rawGo "int64(0)"))
            [ GoMutate resultName (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoVar resultName ] ]) ] []
        ]
      Nothing ->
        [ GoAssign outIndexName (rawGo "int64(0)")
        , GoForRange ("_, " <> valueName <> " := range " <> sourceInfo.target)
            [ GoAssign itemName (integer (GoVar valueName) sourceInfo.elementType)
            , GoIfElse (GoCall (arg 0) [ GoVar itemName ])
                [ GoMutate resultName (GoCall (GoVar callbackName) [ GoVar outIndexName, GoVar resultName, GoVar itemName ])
                , GoMutate outIndexName (GoBinOp "+" (GoVar outIndexName) (rawGo "int64(1)"))
                ] []
            ]
        ]
    usesContrib = case contribLambda, plan of
      Just _, ConcatPlan _ -> true
      Just _, FilterPlan _ -> true
      Just _, MapPlan args -> case Array.head args of
        Just mapper -> IndexedAddLanes.eligible mapper foldedCallback
        Nothing -> false
      Nothing, _ -> false
    contribDecl = case contribLambda of
      Just c | usesContrib -> [ GoAssign contribName c.expr ]
      _ -> []
    foldDecl = case plan, contribLambda of
      FilterPlan _, Just _ -> []
      _, _ -> [ GoAssign callbackName folding.expr ]
    leftLengthExpr = case argType 0 of
      TypeNativeArray _ -> GoCall (GoVar "len") [ GoVar (argName 0) ]
      _ -> GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayLength") [ GoVar (argName 0) ]
    loopStmts = case plan of
      MapPlan _ -> withIntSource 1 mappedLoop
      FilterPlan _ -> withIntSource 1 filterBody
      ConcatPlan _ ->
        [ GoAssign offsetName leftLengthExpr ]
          <> withIntSource 0 (\src -> concatLoop src "laneL" Nothing (GoInt 0))
          <> withIntSource 1 (\src -> concatLoop src "laneR" (Just (GoVar offsetName)) (GoVar offsetName))
    body = foldDecl <> contribDecl <> [ GoAssign resultName (GoCall (GoVar "int64") [ GoInt descriptor.seed ]) ] <> loopStmts
  pure { stmts: captured.stmts, expr: GoCall (GoFuncLit [] body (GoVar resultName) TypeInt64) []
       , exprType: TypeInt64, nextId: nextAfterContrib + 1 }
  where
  integer expression from = coerceGoExpr context.codegenStateRef context.modNameStr expression from TypeInt64

recognizeProducer :: TcoExpr -> Maybe ProducerPlan
recognizeProducer producer = case unwrapTcoExpr producer of
  UncurriedApp fn args@[ predicate, _ ] -> do
    target <- qualifiedTarget fn
    guard (target.mbMod == Just (ModuleName "Data.Array") && target.name == "filterImpl")
    guard (scalarPredicate predicate)
    pure (FilterPlan args)
  _ -> do
    guard (onlyCurried producer)
    let Tuple function spine = collectGoSpine producer
    target <- qualifiedTarget function
    let arguments = getGoSpineArgs spine
    if target.mbMod == Just (ModuleName "Data.Semigroup") && target.name == "concatArray" then do
      guard (Array.length arguments == 2)
      pure (ConcatPlan arguments)
    else if target.mbMod == Just (ModuleName "Data.Functor") && target.name == "arrayMap" then do
      guard (Array.length arguments == 2)
      callback <- Array.head arguments
      signature <- extractFuncType callback
      guard (signature.fArgs == [ Int ] && signature.fRet == Int && closedPrimitiveCallback 1 (neutral callback))
      case unwrapTcoExpr callback of
        Abs parameters body -> do
          let grouped = collectCurriedAbs parameters body
          guard (NEA.length grouped.args == 1 && readOnlyScalar (map (\(Tuple _ level) -> level) (NEA.toArray grouped.args)) grouped.body)
          pure (MapPlan arguments)
        _ -> Nothing
    else if target.mbMod == Just (ModuleName "Data.Array") && target.name == "filter" then do
      guard (Array.length arguments == 2)
      predicate <- Array.head arguments
      guard (scalarPredicate predicate)
      pure (FilterPlan arguments)
    else Nothing

scalarPredicate :: TcoExpr -> Boolean
scalarPredicate predicate = case extractFuncType predicate of
  Just signature | signature.fArgs == [ Int ] && signature.fRet == Boolean -> case unwrapTcoExpr predicate of
    Abs parameters body ->
      let grouped = collectCurriedAbs parameters body
      in NEA.length grouped.args == 1 && case NEA.head grouped.args of
        Tuple _ level -> scalarType (Map.singleton level Int) grouped.body == Just Boolean
    _ -> false
  _ -> false

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
