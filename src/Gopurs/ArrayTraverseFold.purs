module Gopurs.ArrayTraverseFold (emit, justMapper, annotated) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (evalState)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (any, foldM, foldr)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..), snd)
import Gopurs.IndexedAddLanes as IndexedAddLanes
import Gopurs.ArrayProducerFold (onlyCurried)
import Gopurs.ArrayTraverse (recognizeMaybe)
import Gopurs.ArrayTraverseTotal (totalIntegerMaybe)
import Gopurs.CallAnalysis (collectCurriedAbs, collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.ClosedIndexedFolds (closedPrimitiveCallback, validDescriptor)
import Gopurs.CodegenState (ArrayFoldDescriptor)
import Gopurs.ExprAnalysis (extractFuncType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, bindParameters, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.ImmediateApplications.Scope (freshen, maximumLevel)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName(..), ProperName(..), Qualified(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendSyntax(..), Level, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syntax

maybeInt :: ExprType
maybeInt = ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ Int ]

-- The callback must construct Just directly. Its payload is then checked by
-- the existing closed scalar proof, including every nested Typed annotation.
-- Nothing/branches/captures/foreign calls retain the ordinary traversal.
justMapper :: Int -> TcoExpr -> Maybe NeutralExpr
justMapper arity callback = do
  signature <- extractFuncType callback
  guard (signature.fArgs == Array.replicate arity Int && signature.fRet == maybeInt)
  case unwrapTcoExpr callback of
    Abs parameters body -> do
      let grouped = collectCurriedAbs parameters body
      guard (NEA.length grouped.args == arity && totalIntegerMaybe (map (\(Tuple _ level) -> level) (NEA.toArray grouped.args)) grouped.body)
    _ -> Nothing
  mapped <- convert arity callback
  guard (closedPrimitiveCallback arity mapped)
  pure mapped
  where
  convert remaining value = do
    guard (annotated (Func (Array.replicate remaining Int) maybeInt) value)
    case unwrapTcoExpr value of
      Abs parameters body -> do
        let count = NEA.length parameters
        guard (count <= remaining)
        scalar <- if count < remaining then convert (remaining - count) body else do
          guard (annotated maybeInt body)
          case unwrapTcoExpr body of
            CtorSaturated (Qualified (Just (ModuleName "Data.Maybe")) (Ident "Just")) _ (ProperName "Maybe") (Ident "Just") [ Tuple _ payload ] ->
              pure (neutral payload)
            _ -> Nothing
        pure (NeutralExpr (Typed (Func (Array.replicate remaining Int) Int) (NeutralExpr (Abs parameters scalar))))
      _ -> Nothing

-- TAST can retain generic m (t b) around a canonical traversal. Refine those
-- transparent applications against the proved result, but reject any concrete
-- contradiction, including nested Maybe String under an outer Maybe Int.
annotated :: ExprType -> TcoExpr -> Boolean
annotated expected (TcoExpr _ syntax) = case syntax of
  Typed actual value -> compatible expected actual && annotated expected value
  Syntax.TypeApp value _ -> annotated expected value
  _ -> true

compatible :: ExprType -> ExprType -> Boolean
compatible expected actual = isJust (refine expected actual Map.empty)

refine :: ExprType -> ExprType -> Map.Map String ExprType -> Maybe (Map.Map String ExprType)
refine expected actual variables = case expected, actual of
  _, Any -> Just variables
  _, TypeVar name -> bindVariable name expected variables
  Array wanted, Array supplied -> refine wanted supplied variables
  Array wanted, ADT "Prim.Array" [ "Prim", "Array" ] [ supplied ] -> refine wanted supplied variables
  Array wanted, CoreFn.TypeApp (TypeVar name) [ supplied ] -> do
    known <- bindVariable name (ADT "Prim.Array" [ "Prim", "Array" ] []) variables
    refine wanted supplied known
  ADT name path arguments, ADT suppliedName suppliedPath suppliedArguments -> do
    guard (name == suppliedName && path == suppliedPath)
    sameArguments arguments suppliedArguments variables
  ADT name path [ wanted ], CoreFn.TypeApp (TypeVar variable) [ supplied ] -> do
    known <- bindVariable variable (ADT name path []) variables
    refine wanted supplied known
  Func arguments result, Func suppliedArguments suppliedResult -> do
    known <- sameArguments arguments suppliedArguments variables
    refine result suppliedResult known
  _, _ | expected == actual -> Just variables
  _, _ -> Nothing
  where
  sameArguments wanted supplied known = do
    guard (Array.length wanted == Array.length supplied)
    foldM (\current (Tuple left right) -> refine left right current) known (Array.zip wanted supplied)
  bindVariable name value known = case Map.lookup name known of
    Just previous | previous /= value -> Nothing
    _ -> Just (Map.insert name value known)

emit :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> Maybe ExprResult
emit translate context nextId mbIdent level producer continuation = do
  guard (Array.null context.recVars && onlyCurried producer)
  let Tuple function spine = collectGoSpine producer
  -- Start with direct class-method forms. FFI methods may contain their own
  -- application stages and are not flattened by this fusion.
  guard (directTraversal function)
  traversal <- recognizeMaybe function (getGoSpineArgs spine)
  guard (Array.length traversal.args == traversal.callbackOffset + 2)
  callback <- Array.index traversal.args traversal.callbackOffset
  input <- Array.index traversal.args (traversal.callbackOffset + 1)
  guard (annotated (Array Int) input)
  guard (annotated (ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ Array Int ]) producer)
  branch <- case unwrapTcoExpr continuation of
    EffectPure value -> do
      -- Do not move traversal work into or out of a deferred effect value.
      guard context.options.inEffectBlock
      guard (annotated (ADT "Effect.Effect" [ "Effect", "Effect" ] [ Int ]) continuation)
      pure value
    _ -> pure continuation
  consumed <- consumedArray context level branch
  let whole = NeutralExpr (Let mbIdent level (neutral producer) (neutral continuation))
      arity = if traversal.indexed then 2 else 1
  case justMapper arity callback of
    Just mapper ->
      pure (emitLoop translate context nextId traversal mapper consumed.descriptor (maximumLevel whole))
    Nothing -> do
      totalCb <- fallibleTotalMapper arity callback
      guard (not (usesLevel level consumed.failure))
      pure (emitFallibleLoop translate context nextId traversal totalCb consumed.descriptor consumed.failure (maximumLevel whole))

fallibleTotalMapper :: Int -> TcoExpr -> Maybe { parameters :: Array (Tuple (Maybe Ident) Level), body :: TcoExpr }
fallibleTotalMapper arity callback = do
  signature <- extractFuncType callback
  guard (signature.fArgs == Array.replicate arity Int && signature.fRet == maybeInt)
  case unwrapTcoExpr callback of
    Abs parameters body -> do
      let grouped = collectCurriedAbs parameters body
          params = NEA.toArray grouped.args
      guard (Array.length params == arity && totalIntegerMaybe (map snd params) grouped.body)
      pure { parameters: params, body: grouped.body }
    _ -> Nothing

usesLevel :: Level -> TcoExpr -> Boolean
usesLevel target (TcoExpr _ syntax) = case syntax of
  Local _ level -> level == target
  _ -> any (usesLevel target) syntax

hasNothing :: TcoExpr -> Boolean
hasNothing expression = case unwrapTcoExpr expression of
  CtorSaturated _ _ _ (Ident "Nothing") [] -> true
  Branch branches otherwise -> any (\(Pair _ result) -> hasNothing result) branches || hasNothing otherwise
  _ -> false

directTraversal :: TcoExpr -> Boolean
directTraversal function = case unwrapTcoExpr function of
  Accessor _ (GetProp "traverse") -> true
  _ -> case qualifiedTarget function of
    Just { mbMod: Just (ModuleName "Data.Traversable"), name: "traverse" } -> true
    Just { mbMod: Just (ModuleName "Data.TraversableWithIndex"), name } -> name == "traverseWithIndex" || name == "traverseWithIndexDefault"
    _ -> false

consumedArray :: ExprContext -> Level -> TcoExpr -> Maybe { descriptor :: ArrayFoldDescriptor, failure :: TcoExpr }
consumedArray context level expression = do
  guard (annotated Int expression)
  { success, failure } <- case unwrapTcoExpr expression of
    Branch branches otherwise -> do
      guard (case unwrapTcoExpr otherwise of
        Fail _ -> true
        _ -> false)
      case NEA.toArray branches of
        [ Pair first firstResult, Pair second secondResult ]
          | tagged "Nothing" first && tagged "Just" second -> Just { success: secondResult, failure: firstResult }
          | tagged "Just" first && tagged "Nothing" second -> Just { success: firstResult, failure: secondResult }
        _ -> Nothing
    _ -> Nothing
  guard (annotated Int success && onlyCurried success)
  let Tuple fn spine = collectGoSpine success
  guard (annotated (Func [ Array Int ] Int) fn)
  case getGoSpineArgs spine of
    [ value ] -> do
      guard (annotated (Array Int) value)
      guard (case unwrapTcoExpr value of
        Accessor object (GetCtorField (Qualified (Just (ModuleName "Data.Maybe")) (Ident "Just")) _ (ProperName "Maybe") (Ident "Just") _ 0) -> sameLocal object
        _ -> false)
    _ -> Nothing
  target <- qualifiedTarget fn
  let local = target.mbMod == Nothing || map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) target.mbMod == Just context.modNameStr
  info <- if local then Map.lookup target.name context.moduleFunctions else do
    owner <- target.mbMod
    Map.lookup (unwrap owner <> "." <> target.name) context.metadata.globalFunctions
  guard (info.arity == 1 && info.fArgs == [ TypeNativeArray TypeInt64 ] && info.fRet == TypeInt64)
  descriptor <- info.arrayFold
  guard (validDescriptor descriptor)
  pure { descriptor, failure }
  where
  sameLocal object = annotated (ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ Array Int ]) object
    && case unwrapTcoExpr object of
      Local _ actual -> actual == level
      _ -> false
  tagged tag condition = annotated Boolean condition && case unwrapTcoExpr condition of
    PrimOp (Op1 (OpIsTag (Qualified (Just (ModuleName "Data.Maybe")) (Ident name))) object) -> name == tag && sameLocal object
    _ -> false

emitLoop :: TranslateExpr -> ExprContext -> Int -> { indexed :: Boolean, callbackOffset :: Int, args :: Array TcoExpr } -> NeutralExpr -> ArrayFoldDescriptor -> Int -> ExprResult
emitLoop translate context nextId traversal mapper descriptor lastLevel =
  let
    start = max lastLevel (maximumLevel descriptor.callback) + nextId + 1
    folded = Tco.analyze [] (evalState (freshen descriptor.callback) start)
    mapped = Tco.analyze [] mapper
    native currentId callback = case unwrapTcoExpr callback of
      Abs parameters callbackBody ->
        let
          grouped = collectCurriedAbs parameters callbackBody
          params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) (NEA.toArray grouped.args)
          callbackContext = context { depth = context.depth + 1, bound = bindParameters params context.bound
            , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
            , mbExpectedExprType = Just Int, mbExpectedGoType = Just TypeInt64 }
          result = translate callbackContext currentId grouped.body
        in { expr: GoFuncLit params (flattenStmts result.stmts) (integer result.expr result.exprType) TypeInt64, nextId: result.nextId }
      _ -> { expr: rawGo "nil", nextId: currentId }
    -- Canonical dictionaries and the adjacent closed mapper are pure and unused
    -- at runtime. Capture only the array operand.
    inputArg = Array.slice (traversal.callbackOffset + 1) (traversal.callbackOffset + 2) traversal.args
    captured = CallArguments.capture { prefix: "traverseFold_arg_", markUsed: true }
      (\_ currentId argument -> translate (childContext context Nothing) currentId argument) nextId inputArg
    mapResult = native captured.nextId mapped
    foldResult = native mapResult.nextId folded
    contribResult = if IndexedAddLanes.eligible mapped folded then
      case IndexedAddLanes.extractContribution folded of
        Just info ->
          let
            params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) [ info.index, info.value ]
            cbCtx = context { depth = context.depth + 1, bound = bindParameters params context.bound
              , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
              , mbExpectedExprType = Just Int, mbExpectedGoType = Just TypeInt64 }
            res = translate cbCtx foldResult.nextId info.body
          in Just { expr: GoFuncLit params (flattenStmts res.stmts) (integer res.expr res.exprType) TypeInt64, nextId: res.nextId }
        Nothing -> Nothing
      else Nothing
    affineParity = IndexedAddLanes.extractAffineParity (Just mapped) folded
    nextAfterContrib = case contribResult of
      Just c -> c.nextId
      Nothing -> foldResult.nextId
    suffix = show nextAfterContrib
    variable label = "traverseFold_" <> label <> "_" <> suffix
    arg position = fromMaybe (rawGo "nil") (Array.index captured.exprs position)
    inputName = case arg 0 of
      GoVar name -> name
      _ -> ""
    inputType = fromMaybe TypeValue (Array.index captured.exprTypes 0)
    index = GoCall (GoVar "int64") [ GoVar (variable "index") ]
    emitForSource src =
      let
        valueExpr raw = if src.isInt32 then GoCall (GoVar "int64") [ raw ] else integer raw src.elementType
        value = valueExpr (GoVar (variable "value"))
        mappedValue = GoCall (GoVar (variable "map")) (if traversal.indexed then [ index, value ] else [ value ])
        ordinary = GoForRange (variable "index" <> ", " <> variable "value" <> " := range " <> src.target)
          [ GoMutate (variable "result") (GoCall (GoVar (variable "fold")) [ index, GoVar (variable "result"), mappedValue ]) ]
      in case src.isInt32, affineParity of
        true, Just aff ->
          IndexedAddLanes.emitAffineSimdCall src.target (variable "result") aff false
        _, _ -> case contribResult of
          Just _ ->
            [ IndexedAddLanes.emitContribution { source: src.target, prefix: variable "lane", result: variable "result", offset: Nothing }
                (\i item ->
                  let current = GoCall (GoVar "int64") [ i ]
                      produced = GoCall (GoVar (variable "map")) (if traversal.indexed then [ current, valueExpr item ] else [ valueExpr item ])
                  in GoCall (GoVar (variable "contrib")) [ current, produced ]) ordinary ]
          Nothing -> [ ordinary ]
    loopStmts = case inputType of
      TypeNativeArray inner ->
        [ GoAssign (variable "source") (GoVar inputName) ] <> emitForSource { target: variable "source", elementType: inner, isInt32: false }
      _ ->
        [ GoIfElse (rawGo (inputName <> ".Type == gopurs_runtime.TypeInt32Array"))
            ([ GoAssign (variable "source_i32") (rawGo ("*(*[]int32)(" <> inputName <> ".UnsafePtr)")) ]
              <> emitForSource { target: variable "source_i32", elementType: TypeInt64, isInt32: true })
            [ GoIfElse (rawGo (inputName <> ".Type == gopurs_runtime.TypeIntArray"))
                ([ GoAssign (variable "source_i") (rawGo ("*(*[]int64)(" <> inputName <> ".UnsafePtr)")) ]
                  <> emitForSource { target: variable "source_i", elementType: TypeInt64, isInt32: false })
                ([ GoAssign (variable "source_v") (rawGo ("*(*[]gopurs_runtime.Value)(" <> inputName <> ".UnsafePtr)")) ]
                  <> emitForSource { target: variable "source_v", elementType: TypeValue, isInt32: false })
            ]
        ]
    contribDecl = case contribResult of
      Just c -> [ GoAssign (variable "contrib") c.expr ]
      Nothing -> []
    body =
      [ GoAssign (variable "map") mapResult.expr, GoAssign (variable "fold") foldResult.expr ] <> contribDecl <>
      [ GoAssign (variable "result") (GoCall (GoVar "int64") [ GoInt descriptor.seed ]) ] <> loopStmts
  in { stmts: captured.stmts, expr: GoCall (GoFuncLit [] body (GoVar (variable "result")) TypeInt64) []
     , exprType: TypeInt64, nextId: nextAfterContrib + 1 }
  where
  integer value from = coerceGoExpr context.codegenStateRef context.modNameStr value from TypeInt64

neutral :: TcoExpr -> NeutralExpr
neutral (TcoExpr _ syntax) = NeutralExpr (map neutral syntax)

-- Fuse a fallible total Maybe traversal directly into its immediate array fold
-- consumer. Because the callback is proved total and side-effect-free,
-- short-circuiting on Nothing without allocating the intermediate array is
-- observationally identical to traversing and discarding on Nothing.
emitFallibleLoop
  :: TranslateExpr
  -> ExprContext
  -> Int
  -> { indexed :: Boolean, callbackOffset :: Int, args :: Array TcoExpr }
  -> { parameters :: Array (Tuple (Maybe Ident) Level), body :: TcoExpr }
  -> ArrayFoldDescriptor
  -> TcoExpr
  -> Int
  -> ExprResult
emitFallibleLoop translate context nextId traversal totalCb descriptor failureExpr lastLevel =
  let
    start = max lastLevel (maximumLevel descriptor.callback) + nextId + 1
    folded = Tco.analyze [] (evalState (freshen descriptor.callback) start)
    inputArg = Array.slice (traversal.callbackOffset + 1) (traversal.callbackOffset + 2) traversal.args
    captured = CallArguments.capture { prefix: "traverseFold_arg_", markUsed: true }
      (\_ currentId argument -> translate (childContext context Nothing) currentId argument) nextId inputArg
    contribInfo = IndexedAddLanes.extractContribution folded
    foldLambda = case contribInfo of
      Just info ->
        let
          params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) [ info.index, info.value ]
          cbCtx = context { depth = context.depth + 1, bound = bindParameters params context.bound
            , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
            , mbExpectedExprType = Just Int, mbExpectedGoType = Just TypeInt64 }
          res = translate cbCtx captured.nextId info.body
        in { expr: GoFuncLit params (flattenStmts res.stmts) (integer res.expr res.exprType) TypeInt64, nextId: res.nextId, isContrib: true }
      Nothing -> case unwrapTcoExpr folded of
        Abs parameters callbackBody ->
          let
            grouped = collectCurriedAbs parameters callbackBody
            params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) (NEA.toArray grouped.args)
            cbCtx = context { depth = context.depth + 1, bound = bindParameters params context.bound
              , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
              , mbExpectedExprType = Just Int, mbExpectedGoType = Just TypeInt64 }
            res = translate cbCtx captured.nextId grouped.body
          in { expr: GoFuncLit params (flattenStmts res.stmts) (integer res.expr res.exprType) TypeInt64, nextId: res.nextId, isContrib: false }
        _ -> { expr: rawGo "nil", nextId: captured.nextId, isContrib: false }
    suffix = show foldLambda.nextId
    variable label = "traverseFold_" <> label <> "_" <> suffix
    arg0 = fromMaybe (rawGo "nil") (Array.index captured.exprs 0)
    inputName = case arg0 of
      GoVar name -> name
      _ -> ""
    inputType = fromMaybe TypeValue (Array.index captured.exprTypes 0)
    indexName = variable "index"
    valueName = variable "value"
    resultName = variable "result"
    acc32Name = variable "acc32"
    foldName = variable "fold"
    failLabel = variable "fail"
    parameterNames = map (\(Tuple name level) -> localId name level) totalCb.parameters
    parameterTypes = map (const TypeInt64) totalCb.parameters
    bindingsFor isInt32 elemType = Array.zipWith GoAssign parameterNames
      (let valExpr = if isInt32 then GoCall (GoVar "int64") [ GoVar valueName ] else integer (GoVar valueName) elemType
       in if Array.length totalCb.parameters == 2 then
         [ GoCall (GoVar "int64") [ GoVar indexName ], valExpr ]
         else [ valExpr ])
    boundContext = (childContext context (Just Int))
      { bound = bindParameters (Array.zip parameterNames parameterTypes) context.bound }
    canFail = hasNothing totalCb.body
    valParamLevel = map snd (Array.last totalCb.parameters)
    eqRejectAffine = extractEqRejectAffine totalCb folded
    fallibleAffine = case IndexedAddLanes.extractAffineParity Nothing folded of
      Just aff | allPayloadsInt32 valParamLevel totalCb.body -> Just aff
      _ -> Nothing
    stepUpdate mbAff valExpr = case mbAff of
      Just (Tuple aff (Just isEven)) ->
        IndexedAddLanes.emitAffineStep32Fixed acc32Name isEven (GoCall (GoVar "int32") [ valExpr ]) aff
      Just (Tuple aff Nothing) ->
        IndexedAddLanes.emitAffineStep32 acc32Name (GoVar indexName) (GoCall (GoVar "int32") [ valExpr ]) aff
      Nothing ->
        if foldLambda.isContrib then
          GoMutate resultName (GoBinOp "+" (GoVar resultName)
            (GoCall (GoVar foldName) [ GoCall (GoVar "int64") [ GoVar indexName ], valExpr ]))
        else
          GoMutate resultName (GoCall (GoVar foldName)
            [ GoCall (GoVar "int64") [ GoVar indexName ], GoVar resultName, valExpr ])
    emitStep mbAff ctx id expr = case unwrapTcoExpr expr of
      CtorSaturated _ _ _ (Ident "Nothing") [] -> { stmts: [ rawGo ("goto " <> failLabel) ], nextId: id }
      CtorSaturated _ _ _ (Ident "Just") [ Tuple _ payload ] ->
        let translated = translate ctx id payload
        in { stmts: flattenStmts translated.stmts <> [ stepUpdate mbAff (integer translated.expr translated.exprType) ]
           , nextId: translated.nextId }
      Branch branches otherwise ->
        foldr (\(Pair condition branchResult) remaining ->
          let
            test = translate (childContext ctx (Just Boolean)) remaining.nextId condition
            yes = emitStep mbAff ctx test.nextId branchResult
          in { stmts: flattenStmts test.stmts <> [ GoIfElse test.expr yes.stmts remaining.stmts ], nextId: yes.nextId })
          (emitStep mbAff ctx id otherwise) branches
      _ -> { stmts: [], nextId: id }
    loopStep = emitStep Nothing boundContext (foldLambda.nextId + 1) totalCb.body
    eqRejectTranslated = case eqRejectAffine of
      Just info ->
        let res = translate ((childContext context (Just Int)) { mbExpectedGoType = Just TypeInt64 }) loopStep.nextId info.rejected
        in Just { rejected: res, affine: info.affine }
      Nothing -> Nothing
    nextAfterEqReject = case eqRejectTranslated of
      Just e -> e.rejected.nextId
      Nothing -> loopStep.nextId
    loopStepAffine = case eqRejectTranslated, fallibleAffine of
      Nothing, Just aff ->
        let
          evenStep = emitStep (Just (Tuple aff (Just true))) boundContext nextAfterEqReject totalCb.body
          oddStep = emitStep (Just (Tuple aff (Just false))) boundContext evenStep.nextId totalCb.body
          singleStep = emitStep (Just (Tuple aff (Just true))) boundContext oddStep.nextId totalCb.body
        in Just { aff, evenStep, oddStep, singleStep, nextId: singleStep.nextId }
      _, _ -> Nothing
    nextAfterSteps = case loopStepAffine of
      Just s -> s.nextId
      Nothing -> nextAfterEqReject
    failureTranslated = if canFail then
      Just (translate ((childContext context (Just Int)) { mbExpectedGoType = Just TypeInt64 }) nextAfterSteps failureExpr)
      else Nothing
    finalNextId = case failureTranslated of
      Just f -> f.nextId
      Nothing -> nextAfterSteps
    normalizeFor target = if foldLambda.isContrib then
      [ GoIfElse (GoBinOp ">" (GoCall (GoVar "len") [ rawGo target ]) (GoInt 0))
          [ GoMutate resultName (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoVar resultName ] ]) ] [] ]
      else []
    bindingsMutate isInt32 elemType = Array.zipWith GoMutate parameterNames
      (let valExpr = if isInt32 then GoCall (GoVar "int64") [ GoVar valueName ] else integer (GoVar valueName) elemType
       in if Array.length totalCb.parameters == 2 then
         [ GoCall (GoVar "int64") [ GoVar indexName ], valExpr ]
         else [ valExpr ])
    emitFallibleSource target isInt32 elemType = case isInt32, eqRejectTranslated, loopStepAffine of
      true, Just eqRej, _ ->
        let
          rejName = variable "rej"
          rej32Name = variable "rej32"
        in
          [ GoIfElse (GoBinOp ">" (GoCall (GoVar "len") [ rawGo target ]) (GoInt 0))
              (flattenStmts eqRej.rejected.stmts
                <> [ GoAssign rejName (integer eqRej.rejected.expr eqRej.rejected.exprType)
                   , GoAssign rej32Name (GoCall (GoVar "int32") [ GoVar rejName ])
                   , GoIfElse (GoBinOp "&&"
                       (GoBinOp "==" (GoCall (GoVar "int64") [ GoVar rej32Name ]) (GoVar rejName))
                       (GoCall (GoSelector (GoVar "gopurs_runtime") "ContainsInt32") [ rawGo target, GoVar rej32Name ]))
                       [ rawGo ("goto " <> failLabel) ] []
                   ]
                <> IndexedAddLanes.emitAffineSimdCall target resultName eqRej.affine false)
              []
          ]
      true, Nothing, Just affineSteps ->
        let
          nName = variable "len"
          pairIdxName = variable "pairIdx"
        in
          [ GoAssign acc32Name (rawGo "int32(0)")
          , GoAssign nName (GoCall (GoVar "len") [ rawGo target ])
          , GoAssign pairIdxName (GoInt 0)
          , GoForRange ("; " <> pairIdxName <> " <= " <> nName <> "-2; " <> pairIdxName <> " += 2")
              ([ GoAssign indexName (GoVar pairIdxName)
               , GoMutate "_" (GoVar indexName)
               , GoAssign valueName (GoIndex (rawGo target) (GoVar pairIdxName))
               ] <> bindingsFor true TypeInt64 <> affineSteps.evenStep.stmts
                 <> [ GoMutate indexName (GoBinOp "+" (GoVar pairIdxName) (GoInt 1))
                    , GoMutate "_" (GoVar indexName)
                    , GoMutate valueName (GoIndex (rawGo target) (GoVar indexName))
                    ] <> bindingsMutate true TypeInt64 <> affineSteps.oddStep.stmts)
          , GoIfElse (GoBinOp "<" (GoVar pairIdxName) (GoVar nName))
              ([ GoAssign indexName (GoVar pairIdxName)
               , GoMutate "_" (GoVar indexName)
               , GoAssign valueName (GoIndex (rawGo target) (GoVar pairIdxName))
               ] <> bindingsFor true TypeInt64 <> affineSteps.singleStep.stmts)
              []
          , GoIfElse (GoBinOp ">" (GoVar nName) (GoInt 0))
              [ GoMutate resultName (GoCall (GoVar "int64") [ GoBinOp "+" (GoCall (GoVar "int32") [ GoVar resultName ]) (GoVar acc32Name) ]) ] []
          ]
      _, _, _ ->
        [ GoForRange (indexName <> ", " <> valueName <> " := range " <> target)
            ([ GoMutate "_" (GoVar indexName) ] <> bindingsFor isInt32 elemType <> loopStep.stmts)
        ] <> normalizeFor target
    sourceStmts = case inputType of
      TypeNativeArray inner ->
        [ GoAssign (variable "source") (GoVar inputName) ] <> emitFallibleSource (variable "source") false inner
      _ ->
        [ GoIfElse (rawGo (inputName <> ".Type == gopurs_runtime.TypeInt32Array"))
            ([ GoAssign (variable "source_i32") (rawGo ("*(*[]int32)(" <> inputName <> ".UnsafePtr)")) ]
              <> emitFallibleSource (variable "source_i32") true TypeInt64)
            [ GoIfElse (rawGo (inputName <> ".Type == gopurs_runtime.TypeIntArray"))
                ([ GoAssign (variable "source_i") (rawGo ("*(*[]int64)(" <> inputName <> ".UnsafePtr)")) ]
                  <> emitFallibleSource (variable "source_i") false TypeInt64)
                ([ GoAssign (variable "source_v") (rawGo ("*(*[]gopurs_runtime.Value)(" <> inputName <> ".UnsafePtr)")) ]
                  <> emitFallibleSource (variable "source_v") false TypeValue)
            ]
        ]
    failBlock = case failureTranslated of
      Just f ->
        [ GoReturn (GoVar resultName)
        , rawGo (failLabel <> ":")
        ] <> flattenStmts f.stmts <> [ GoReturn (integer f.expr f.exprType) ]
      Nothing -> [ GoReturn (GoVar resultName) ]
    body =
      [ GoAssign foldName foldLambda.expr
      , GoAssign resultName (GoCall (GoVar "int64") [ GoInt descriptor.seed ])
      ] <> sourceStmts <> failBlock
  in { stmts: captured.stmts
     , expr: GoCall (GoFuncBlock [] body TypeInt64) []
     , exprType: TypeInt64
     , nextId: finalNextId + 1 }
  where
  integer value from = coerceGoExpr context.codegenStateRef context.modNameStr value from TypeInt64

allPayloadsInt32 :: Maybe Level -> TcoExpr -> Boolean
allPayloadsInt32 valLevel expr = case unwrapTcoExpr expr of
  CtorSaturated _ _ _ (Ident "Nothing") [] -> true
  CtorSaturated _ _ _ (Ident "Just") [ Tuple _ payload ] -> isNormalizedInt32 valLevel payload
  Branch branches otherwise ->
    Array.all (\(Pair _ branchResult) -> allPayloadsInt32 valLevel branchResult) (NEA.toArray branches)
      && allPayloadsInt32 valLevel otherwise
  _ -> false

isNormalizedInt32 :: Maybe Level -> TcoExpr -> Boolean
isNormalizedInt32 valLevel expr = case unwrapTcoExpr expr of
  Lit (LitInt _) -> true
  Local _ level -> Just level == valLevel
  PrimOp (Op1 OpIntNegate _) -> true
  PrimOp (Op1 OpIntBitNot _) -> true
  PrimOp (Op2 (OpIntNum OpAdd) _ _) -> true
  PrimOp (Op2 (OpIntNum OpSubtract) _ _) -> true
  PrimOp (Op2 (OpIntNum OpMultiply) _ _) -> true
  PrimOp (Op2 OpIntBitAnd _ _) -> true
  PrimOp (Op2 OpIntBitOr _ _) -> true
  PrimOp (Op2 OpIntBitXor _ _) -> true
  PrimOp (Op2 OpIntBitShiftLeft _ _) -> true
  PrimOp (Op2 OpIntBitShiftRight _ _) -> true
  _ -> false

extractEqRejectAffine
  :: { parameters :: Array (Tuple (Maybe Ident) Level), body :: TcoExpr }
  -> TcoExpr
  -> Maybe { rejected :: TcoExpr, affine :: IndexedAddLanes.AffineParity }
extractEqRejectAffine totalCb folded = case totalCb.parameters of
  [ param@(Tuple _ valLevel) ] -> do
    { rejected, payload } <- matchEqReject valLevel totalCb.body
    let syntheticMapper = Tco.analyze [] (NeutralExpr (Typed (Func [ Int ] Int) (NeutralExpr (Abs (NEA.singleton param) (neutral payload)))))
    affine <- IndexedAddLanes.extractAffineParity (Just syntheticMapper) folded
    pure { rejected, affine }
  _ -> Nothing

matchEqReject :: Level -> TcoExpr -> Maybe { rejected :: TcoExpr, payload :: TcoExpr }
matchEqReject valLevel expr = case unwrapTcoExpr expr of
  Branch branches otherwise -> case NEA.toArray branches of
    [ Pair cond branchRes ] -> case isNothingCtor branchRes, matchJustPayload otherwise of
      true, Just payload -> do
        rejected <- matchEqAgainst valLevel OpEq cond
        pure { rejected, payload }
      _, _ -> case matchJustPayload branchRes, isNothingCtor otherwise of
        Just payload, true -> do
          rejected <- matchEqAgainst valLevel OpNotEq cond
          pure { rejected, payload }
        _, _ -> Nothing
    _ -> Nothing
  _ -> Nothing

isNothingCtor :: TcoExpr -> Boolean
isNothingCtor expr = case unwrapTcoExpr expr of
  CtorSaturated (Qualified (Just (ModuleName "Data.Maybe")) (Ident "Nothing")) _ (ProperName "Maybe") (Ident "Nothing") [] -> true
  _ -> false

matchJustPayload :: TcoExpr -> Maybe TcoExpr
matchJustPayload expr = case unwrapTcoExpr expr of
  CtorSaturated (Qualified (Just (ModuleName "Data.Maybe")) (Ident "Just")) _ (ProperName "Maybe") (Ident "Just") [ Tuple _ payload ] -> Just payload
  _ -> Nothing

matchEqAgainst :: Level -> BackendOperatorOrd -> TcoExpr -> Maybe TcoExpr
matchEqAgainst valLevel expectedOp expr = case unwrapTcoExpr expr of
  PrimOp (Op2 (OpIntOrd op) left right) | op == expectedOp ->
    if isLocalLevel valLevel left && not (usesLevel valLevel right) then Just right
    else if isLocalLevel valLevel right && not (usesLevel valLevel left) then Just left
    else Nothing
  _ -> Nothing

isLocalLevel :: Level -> TcoExpr -> Boolean
isLocalLevel target expr = case unwrapTcoExpr expr of
  Local _ level -> level == target
  _ -> false
