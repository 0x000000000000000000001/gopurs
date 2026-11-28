module Gopurs.ArrayProducerFold (emit, onlyCurried) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (evalState)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Int.Bits ((.&.))
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..))
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.ArrayIndexedFold (readOnlyScalar, scalarType)
import Gopurs.ArraySort as ArraySort
import Gopurs.CongruenceSum as CongruenceSum
import Gopurs.IndexedAddLanes as IndexedAddLanes
import Gopurs.CallAnalysis (collectCurriedAbs, collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.ClosedIndexedFolds (closedPrimitiveCallback, validDescriptor)
import Gopurs.CodegenState (ArrayFoldDescriptor)
import Gopurs.ExprAnalysis (extractFuncType, getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, bindParameters, childContext, flattenStmts)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.GoSymbol (GoReference(..), generated)
import Gopurs.ImmediateApplications.Scope (freshen, maximumLevel)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.CoreFn (ExprType(Array, Boolean, Int), Literal(..), ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendOperator2(..), BackendOperatorOrd(..), BackendSyntax(..))

data ProducerPlan
  = ConcatPlan (Array TcoExpr)
  | MapPlan (Array TcoExpr)
  | FilterPlan (Array TcoExpr)
  | SortPlan (Array TcoExpr)

data SourceKind
  = SrcInt32
  | SrcInt64
  | SrcValue

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
      SortPlan args -> args
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
    affineParity = case plan of
      MapPlan args -> IndexedAddLanes.extractAffineParity (Array.head args) foldedCallback
      ConcatPlan _ -> IndexedAddLanes.extractAffineParity Nothing foldedCallback
      FilterPlan _ -> IndexedAddLanes.extractAffineParity Nothing foldedCallback
      SortPlan _ -> IndexedAddLanes.extractAffineParity Nothing foldedCallback
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
    acc32Name = variable "acc32"
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
    elemInt64 expr = case _ of
      SrcInt32 -> GoCall (GoVar "int64") [ expr ]
      SrcInt64 -> expr
      SrcValue -> integer expr TypeValue
    withIntSource index render =
      let
        inputName = argName index
        baseName = variable ("source" <> show index)
      in case argType index of
        TypeNativeArray TypeInt64 ->
          [ GoAssign baseName (GoVar inputName) ] <> render { target: baseName, kind: SrcInt64 }
        TypeNativeArray _ ->
          [ GoAssign baseName (GoVar inputName) ] <> render { target: baseName, kind: SrcValue }
        _ ->
          [ GoIfElse (rawGo (inputName <> ".Type == gopurs_runtime.TypeInt32Array"))
              ([ GoAssign (baseName <> "_i32") (rawGo ("*(*[]int32)(" <> inputName <> ".UnsafePtr)")) ]
                <> render { target: baseName <> "_i32", kind: SrcInt32 })
              [ GoIfElse (rawGo (inputName <> ".Type == gopurs_runtime.TypeIntArray"))
                  ([ GoAssign (baseName <> "_i") (rawGo ("*(*[]int64)(" <> inputName <> ".UnsafePtr)")) ]
                    <> render { target: baseName <> "_i", kind: SrcInt64 })
                  ([ GoAssign (baseName <> "_v") (rawGo ("*(*[]gopurs_runtime.Value)(" <> inputName <> ".UnsafePtr)")) ]
                    <> render { target: baseName <> "_v", kind: SrcValue })
              ]
          ]
    isMapped = case plan of
      MapPlan _ -> true
      _ -> false
    loop sourceInfo offset = GoForRange (indexName <> ", " <> valueName <> " := range " <> sourceInfo.target)
      [ GoMutate resultName (GoCall (GoVar callbackName)
          [ GoCall (GoVar "int64") [ GoBinOp "+" offset (GoVar indexName) ]
          , GoVar resultName
          , if isMapped then GoCall (arg 0) [ elemInt64 (GoVar valueName) sourceInfo.kind ]
            else elemInt64 (GoVar valueName) sourceInfo.kind
          ]) ]
    mappedLoop sourceInfo = case sourceInfo.kind, affineParity of
      SrcInt32, Just aff ->
        IndexedAddLanes.emitAffineSimdCall sourceInfo.target resultName aff false
      _, _ -> case plan, Array.head planArgs, contribLambda of
        MapPlan _, Just mapper, Just _ | IndexedAddLanes.eligible mapper foldedCallback ->
          [ IndexedAddLanes.emitContribution { source: sourceInfo.target, prefix: variable "lane", result: resultName, offset: Nothing }
              (\i value -> GoCall (GoVar contribName)
                [ GoCall (GoVar "int64") [ i ], GoCall (arg 0) [ elemInt64 value sourceInfo.kind ] ])
              (loop sourceInfo (GoInt 0)) ]
        _, _, _ -> [ loop sourceInfo (GoInt 0) ]
    concatLoop sourceInfo label offsetOpt fallbackOffset = case sourceInfo.kind, affineParity, offsetOpt of
      SrcInt32, Just aff, Nothing ->
        IndexedAddLanes.emitAffineSimdCall sourceInfo.target resultName aff false
      SrcInt32, Just aff, Just off ->
        [ GoIfElse (GoBinOp "==" (GoBinOp "&" off (GoInt 1)) (GoInt 0))
            (IndexedAddLanes.emitAffineSimdCall sourceInfo.target resultName aff false)
            (IndexedAddLanes.emitAffineSimdCall sourceInfo.target resultName aff true)
        ]
      _, _, _ -> case contribLambda of
        Just _ ->
          [ IndexedAddLanes.emitContribution { source: sourceInfo.target, prefix: variable label, result: resultName, offset: offsetOpt }
              (\i value -> GoCall (GoVar contribName)
                [ GoCall (GoVar "int64") [ i ], elemInt64 value sourceInfo.kind ])
              (loop sourceInfo fallbackOffset) ]
        Nothing -> [ loop sourceInfo fallbackOffset ]
    filterMaskPred = case plan of
      FilterPlan [ predicate, _ ] -> extractMaskPredicate predicate
      _ -> Nothing
    filterBody sourceInfo = case sourceInfo.kind, affineParity, filterMaskPred of
      SrcInt32, Just aff, Just { mask, target } ->
        [ rawGo (acc32Name <> ", " <> outIndexName <> " := gopurs_runtime.FilterMaskAffineParitySumInt32("
            <> sourceInfo.target <> ", int32(" <> show mask <> "), int32(" <> show target
            <> "), int32(" <> show aff.aEven <> "), int32(" <> show aff.bEven
            <> "), int32(" <> show aff.aOdd <> "), int32(" <> show aff.bOdd <> "))")
        , GoIfElse (GoBinOp ">" (GoVar outIndexName) (rawGo "int64(0)"))
            [ GoMutate resultName (GoCall (GoVar "int64") [ GoBinOp "+" (GoCall (GoVar "int32") [ GoVar resultName ]) (GoVar acc32Name) ]) ] []
        ]
      SrcInt32, Just aff, Nothing ->
        [ GoAssign outIndexName (rawGo "int64(0)")
        , GoAssign acc32Name (rawGo "int32(0)")
        , GoForRange ("_, " <> valueName <> " := range " <> sourceInfo.target)
            [ GoIfElse (GoCall (arg 0) [ GoCall (GoVar "int64") [ GoVar valueName ] ])
                [ IndexedAddLanes.emitAffineStep32 acc32Name (GoVar outIndexName) (GoVar valueName) aff
                , GoMutate outIndexName (GoBinOp "+" (GoVar outIndexName) (rawGo "int64(1)"))
                ] []
            ]
        , GoIfElse (GoBinOp ">" (GoVar outIndexName) (rawGo "int64(0)"))
            [ GoMutate resultName (GoCall (GoVar "int64") [ GoBinOp "+" (GoCall (GoVar "int32") [ GoVar resultName ]) (GoVar acc32Name) ]) ] []
        ]
      _, _, _ -> case contribLambda of
        Just _ ->
          [ GoAssign outIndexName (rawGo "int64(0)")
          , GoForRange ("_, " <> valueName <> " := range " <> sourceInfo.target)
              [ GoAssign itemName (elemInt64 (GoVar valueName) sourceInfo.kind)
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
              [ GoAssign itemName (elemInt64 (GoVar valueName) sourceInfo.kind)
              , GoIfElse (GoCall (arg 0) [ GoVar itemName ])
                  [ GoMutate resultName (GoCall (GoVar callbackName) [ GoVar outIndexName, GoVar resultName, GoVar itemName ])
                  , GoMutate outIndexName (GoBinOp "+" (GoVar outIndexName) (rawGo "int64(1)"))
                  ] []
              ]
          ]
    slicesSort = GoRaw { text: "slices.Sort", imports: [ "slices" ] }
    foldSortedSlice target kind = case kind, affineParity of
      SrcInt32, Just aff ->
        IndexedAddLanes.emitAffineSimdCall target resultName aff false
      _, _ -> case contribLambda of
        Just _ ->
          [ IndexedAddLanes.emitContribution { source: target, prefix: variable "sortLane", result: resultName, offset: Nothing }
              (\i value -> GoCall (GoVar contribName)
                [ GoCall (GoVar "int64") [ i ], elemInt64 value kind ])
              (loop { target, kind } (GoInt 0)) ]
        Nothing -> [ loop { target, kind } (GoInt 0) ]
    sortLoop sourceInfo =
      let
        nName = variable "sortN"
        prevName = variable "sortPrev"
        curName = variable "sortCur"
        monoName = variable "sortMono"
        outName = variable "sortOut"
        iName = variable "sortI"
        vName = variable "sortV"
        lastName = variable "sortLast"
        elemGoType = case sourceInfo.kind of
          SrcInt32 -> "[]int32"
          _ -> "[]int64"
        outKind = case sourceInfo.kind of
          SrcInt32 -> SrcInt32
          _ -> SrcInt64
        copyToOut = case sourceInfo.kind of
          SrcValue ->
            [ GoForRange (iName <> ", " <> vName <> " := range " <> sourceInfo.target)
                [ GoMutate (outName <> "[" <> iName <> "]") (elemInt64 (GoVar vName) SrcValue) ]
            ]
          _ -> [ GoCall (GoVar "copy") [ GoVar outName, rawGo sourceInfo.target ] ]
        allocSortAndFold =
          [ GoAssign outName (GoCall (GoVar "make") [ rawGo elemGoType, GoVar nName ]) ]
            <> copyToOut
            <> [ GoCall slicesSort [ GoVar outName ] ]
            <> foldSortedSlice outName outKind
        foldReversedSlice = case sourceInfo.kind, affineParity of
          SrcInt32, Just aff ->
            if aff.aEven == aff.aOdd && aff.bEven == aff.bOdd then
              IndexedAddLanes.emitAffineSimdCall sourceInfo.target resultName aff false
            else
              [ GoIfElse (GoBinOp "==" (GoBinOp "&" (GoVar nName) (GoInt 1)) (GoInt 0))
                  (IndexedAddLanes.emitAffineSimdCall sourceInfo.target resultName aff true)
                  (IndexedAddLanes.emitAffineSimdCall sourceInfo.target resultName aff false)
              ]
          _, _ ->
            [ GoAssign lastName (GoBinOp "-" (GoVar nName) (GoInt 1))
            , GoForRange (iName <> " := 0; " <> iName <> " < " <> nName <> "; " <> iName <> "++")
                [ GoMutate resultName (GoCall (GoVar callbackName)
                    [ GoCall (GoVar "int64") [ GoVar iName ]
                    , GoVar resultName
                    , elemInt64 (GoIndex (rawGo sourceInfo.target) (GoBinOp "-" (GoVar lastName) (GoVar iName))) sourceInfo.kind
                    ]) ]
            ]
        readAt idx = case sourceInfo.kind of
          SrcValue -> GoSelector (GoIndex (rawGo sourceInfo.target) idx) "IntVal"
          _ -> GoIndex (rawGo sourceInfo.target) idx
        curExpr = case sourceInfo.kind of
          SrcValue -> GoSelector (GoVar curName) "IntVal"
          _ -> GoVar curName
        descBranch = case sourceInfo.kind of
          SrcInt32 ->
            [ GoIfElse (GoCall (GoSelector (GoVar "gopurs_runtime") "IsNonIncreasingInt32") [ rawGo sourceInfo.target ])
                foldReversedSlice
                allocSortAndFold
            ]
          _ ->
            [ GoAssign monoName (rawGo "true")
            , GoForRange ("_, " <> curName <> " := range " <> sourceInfo.target <> "[2:]")
                [ GoIfElse (GoBinOp ">" curExpr (GoVar prevName))
                    [ GoMutate monoName (rawGo "false"), rawGo "break" ] []
                , GoMutate prevName curExpr
                ]
            , GoIfElse (GoVar monoName)
                foldReversedSlice
                allocSortAndFold
            ]
        ascBranch = case sourceInfo.kind of
          SrcInt32 ->
            [ GoIfElse (GoCall (GoSelector (GoVar "gopurs_runtime") "IsNonDecreasingInt32") [ rawGo sourceInfo.target ])
                (foldSortedSlice sourceInfo.target sourceInfo.kind)
                allocSortAndFold
            ]
          _ ->
            [ GoAssign monoName (rawGo "true")
            , GoForRange ("_, " <> curName <> " := range " <> sourceInfo.target <> "[2:]")
                [ GoIfElse (GoBinOp "<" curExpr (GoVar prevName))
                    [ GoMutate monoName (rawGo "false"), rawGo "break" ] []
                , GoMutate prevName curExpr
                ]
            , GoIfElse (GoVar monoName)
                (foldSortedSlice sourceInfo.target sourceInfo.kind)
                allocSortAndFold
            ]
      in
        [ GoAssign nName (GoCall (GoVar "len") [ rawGo sourceInfo.target ])
        , GoIfElse (GoBinOp "<" (GoVar nName) (GoInt 2))
            (foldSortedSlice sourceInfo.target sourceInfo.kind)
            [ GoAssign prevName (readAt (GoInt 1))
            , GoIfElse (GoBinOp "<" (GoVar prevName) (readAt (GoInt 0)))
                descBranch
                ascBranch
            ]
        ]
    usesContrib = case contribLambda, plan of
      Just _, ConcatPlan _ -> true
      Just _, FilterPlan _ -> true
      Just _, SortPlan _ -> true
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
      SortPlan _ -> withIntSource 0 sortLoop
      ConcatPlan _ ->
        [ GoAssign offsetName leftLengthExpr ]
          <> withIntSource 0 (\src -> concatLoop src "laneL" Nothing (GoInt 0))
          <> withIntSource 1 (\src -> concatLoop src "laneR" (Just (GoVar offsetName)) (GoVar offsetName))
    body = foldDecl <> contribDecl <> [ GoAssign resultName (GoCall (GoVar "int64") [ GoInt descriptor.seed ]) ] <> loopStmts
    isNativeArray = case _ of
      TypeNativeArray _ -> true
      _ -> false
    outlineSlow params callArgs = unsafePerformEffect do
      state <- Ref.read context.codegenStateRef
      let
        coldName = "Call_" <> context.modNameStr <> "___gopurs_arrayProducer_fold_" <> show state.globalId
        decl = GoFunctionDecl
          { name: generated context.modNameStr coldName
          , params
          , result: TypeInt64
          , body: GoBlock (body <> [ GoReturn (GoVar resultName) ])
          }
      Ref.modify_ (\current -> current
        { globalId = current.globalId + 1
        , declarations = current.declarations <> [ decl ]
        }) context.codegenStateRef
      pure (GoCall (GoRef (SymbolReference (generated context.modNameStr coldName))) callArgs)
    finalBody = case plan, affineParity of
      ConcatPlan _, Just aff | not (isNativeArray (argType 0)) && not (isNativeArray (argType 1)) ->
        let
          s0 = variable "fast0"
          s1 = variable "fast1"
          coldCall = outlineSlow
            [ Tuple (argName 0) TypeValue, Tuple (argName 1) TypeValue ]
            [ GoVar (argName 0), GoVar (argName 1) ]
        in
          [ GoAssign resultName (GoCall (GoVar "int64") [ GoInt descriptor.seed ])
          , GoIfElse (rawGo (argName 0 <> ".Type == gopurs_runtime.TypeInt32Array && " <> argName 1 <> ".Type == gopurs_runtime.TypeInt32Array"))
              ([ GoAssign s0 (rawGo ("*(*[]int32)(" <> argName 0 <> ".UnsafePtr)"))
               , GoAssign s1 (rawGo ("*(*[]int32)(" <> argName 1 <> ".UnsafePtr)"))
               ]
               <> IndexedAddLanes.emitAffineSimdCall s0 resultName aff false
               <> [ GoIfElse (GoBinOp "==" (GoBinOp "&" (GoCall (GoVar "len") [ GoVar s0 ]) (GoInt 1)) (GoInt 0))
                      (IndexedAddLanes.emitAffineSimdCall s1 resultName aff false)
                      (IndexedAddLanes.emitAffineSimdCall s1 resultName aff true)
                  ])
              [ GoMutate resultName coldCall ]
          ]
      MapPlan _, Just aff | not (isNativeArray (argType 1)) ->
        let
          s1 = variable "fast1"
          coldCall = outlineSlow
            [ Tuple (argName 0) (argType 0), Tuple (argName 1) TypeValue ]
            [ GoVar (argName 0), GoVar (argName 1) ]
        in
          [ GoAssign resultName (GoCall (GoVar "int64") [ GoInt descriptor.seed ])
          , GoIfElse (rawGo (argName 1 <> ".Type == gopurs_runtime.TypeInt32Array"))
              ([ GoAssign s1 (rawGo ("*(*[]int32)(" <> argName 1 <> ".UnsafePtr)")) ]
               <> IndexedAddLanes.emitAffineSimdCall s1 resultName aff false)
              [ GoMutate resultName coldCall ]
          ]
      FilterPlan _, Just _ | Just _ <- filterMaskPred, not (isNativeArray (argType 1)) ->
        let
          s1 = variable "fast1"
          coldCall = outlineSlow
            [ Tuple (argName 0) (argType 0), Tuple (argName 1) TypeValue ]
            [ GoVar (argName 0), GoVar (argName 1) ]
        in
          [ GoAssign resultName (GoCall (GoVar "int64") [ GoInt descriptor.seed ])
          , GoIfElse (rawGo (argName 1 <> ".Type == gopurs_runtime.TypeInt32Array"))
              ([ GoAssign s1 (rawGo ("*(*[]int32)(" <> argName 1 <> ".UnsafePtr)")) ]
               <> filterBody { target: s1, kind: SrcInt32 })
              [ GoMutate resultName coldCall ]
          ]
      SortPlan _, Just _ | not (isNativeArray (argType 0)) ->
        let
          s0 = variable "fast0"
          coldCall = outlineSlow
            [ Tuple (argName 0) TypeValue ]
            [ GoVar (argName 0) ]
        in
          [ GoAssign resultName (GoCall (GoVar "int64") [ GoInt descriptor.seed ])
          , GoIfElse (rawGo (argName 0 <> ".Type == gopurs_runtime.TypeInt32Array"))
              ([ GoAssign s0 (rawGo ("*(*[]int32)(" <> argName 0 <> ".UnsafePtr)")) ]
               <> sortLoop { target: s0, kind: SrcInt32 })
              [ GoMutate resultName coldCall ]
          ]
      _, _ -> body
  pure { stmts: captured.stmts, expr: GoCall (GoFuncLit [] finalBody (GoVar resultName) TypeInt64) []
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
  UncurriedApp fn [ comparator, conversion, input ] -> do
    target <- qualifiedTarget fn
    guard (target.mbMod == Just (ModuleName "Data.Array") && target.name == "sortByImpl")
    guard (ArraySort.intComparator comparator)
    { lt, eq, gt } <- ArraySort.staticOrderingConversion conversion
    guard (lt < 0 && eq == 0 && gt > 0)
    pure (SortPlan [ input ])
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
    else if target.mbMod == Just (ModuleName "Data.Array") && (target.name == "sortBy" || String.indexOf (Pattern "sortBy__") target.name == Just 0) then do
      guard (Array.length arguments == 2)
      comparator <- Array.head arguments
      guard (ArraySort.intComparator comparator)
      input <- Array.index arguments 1
      pure (SortPlan [ input ])
    else if target.mbMod == Just (ModuleName "Data.Array") && target.name == "sort" then do
      guard (Array.length arguments == 2)
      dictionary <- Array.head arguments
      guard (case qualifiedTarget dictionary of
        Just d -> d.mbMod == Just (ModuleName "Data.Ord") && d.name == "ordInt"
        Nothing -> false)
      input <- Array.index arguments 1
      pure (SortPlan [ input ])
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

-- Recognize pure bitmask/power-of-two congruence predicates `(x & mask) == target`
-- or `mod x (2^k) == residue` on Int32 elements. For Euclidean `mod` by a power
-- of two 2^k, `mod x (2^k) == r` is bit-identical to `(x & (2^k - 1)) == r` for
-- all signed 32-bit integers (positive, zero, and negative).
extractMaskPredicate :: TcoExpr -> Maybe { mask :: Int, target :: Int }
extractMaskPredicate predicate = case unwrapTcoExpr predicate of
  Abs parameters body -> do
    let grouped = collectCurriedAbs parameters body
    guard (NEA.length grouped.args == 1)
    Tuple name level <- Just (NEA.head grouped.args)
    let paramKey = localId name level
    case CongruenceSum.recognize paramKey grouped.body of
      Just { modulus, residue } | modulus > 0 && (modulus .&. (modulus - 1)) == 0 ->
        Just { mask: modulus - 1, target: residue }
      _ -> case unwrapTcoExpr grouped.body of
        PrimOp (Op2 (OpIntOrd OpEq) left right) ->
          case extractBitAnd paramKey left, extractLitInt right of
            Just mask, Just target -> Just { mask, target }
            _, _ -> case extractBitAnd paramKey right, extractLitInt left of
              Just mask, Just target -> Just { mask, target }
              _, _ -> Nothing
        _ -> Nothing
  _ -> Nothing
  where
  extractLitInt expr = case unwrapTcoExpr expr of
    Lit (LitInt n) -> Just n
    _ -> Nothing
  extractBitAnd paramKey expr = case unwrapTcoExpr expr of
    PrimOp (Op2 OpIntBitAnd a b) ->
      case unwrapTcoExpr a, extractLitInt b of
        Local n l, Just m | localId n l == paramKey -> Just m
        _, _ -> case unwrapTcoExpr b, extractLitInt a of
          Local n l, Just m | localId n l == paramKey -> Just m
          _, _ -> Nothing
    _ -> Nothing

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
