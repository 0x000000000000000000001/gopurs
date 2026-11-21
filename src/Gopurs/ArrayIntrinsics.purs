module Gopurs.ArrayIntrinsics
  ( ArrayIntrinsic
  , Convention(..)
  , recognize
  , emitCurried
  , emitUncurried
  , module Index
  ) where

import Prelude
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIntrinsics.Index (safeIndex, unsafeIndex) as Index
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.ArrayIntrinsics.Callback as Callback
import Gopurs.CallAnalysis (CallTarget)
import Gopurs.CallArguments (Arguments, applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (extractFuncType, getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), childContext)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..), goTypeToStr)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

data ArrayIntrinsic = MapArray | FoldlArray | FilterArray

data Convention = Curried | Uncurried

-- Preserve each calling convention's name and qualification guards.
recognize :: Convention -> String -> Maybe CallTarget -> Int -> Maybe ArrayIntrinsic
recognize convention modNameStr target count = case target of
  Just { name: "arrayMap" } | count >= 2 -> Just MapArray
  Just { name: "foldlArray" } | count >= 3 -> Just FoldlArray
  Just { mbMod, name } | count >= 2 && (mbMod == Just (ModuleName "Data.Array") || (mbMod == Nothing && modNameStr == "Data.Array")) ->
    case convention, name of
      Curried, "filter" -> Just FilterArray
      Uncurried, "filterImpl" -> Just FilterArray
      _, _ -> Nothing
  _ -> Nothing

loopNames :: ArrayIntrinsic -> Int -> { arrValName :: String, arrGoName :: String, resGoName :: String, iName :: String, vName :: String }
loopNames intrinsic depth =
  let
    name = case intrinsic of
      MapArray -> "arrayMap"
      FoldlArray -> "foldlArray"
      FilterArray -> "filterImpl"
    iifeName = name <> show depth
  in
    { arrValName: "arr_val_" <> iifeName, arrGoName: "arr_go_" <> iifeName
    , resGoName: "res_go_" <> iifeName, iName: "i_" <> iifeName, vName: "v_" <> iifeName
    }

-- Capture the intrinsic's strict prefix in source order. The callback must
-- be constructed once, including when the input array is empty. Surplus
-- curried arguments belong to the application of the intrinsic's result.
emitCurried :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> ArrayIntrinsic -> ExprResult
emitCurried translate context@{ codegenStateRef, modNameStr } nextId fn args intrinsic = case intrinsic of
  FoldlArray -> emitCurriedFold translate context nextId fn args
  _ ->
    let
      result = emitUncurried translate context nextId (Array.take 2 args) intrinsic
      remaining = Array.drop 2 args
    in
      if Array.null remaining then result
      else
        let
          name = "arrayIntrinsic_result_" <> show result.nextId
          extras = CallArguments.capture { prefix: "arrayIntrinsic_extra_", markUsed: false }
            (\_ currentId arg ->
              let value = translate (childContext context Nothing) currentId arg
              in value { expr = boxGoExpr codegenStateRef modNameStr value.expr value.exprType, exprType = TypeValue })
            (result.nextId + 1) remaining
        in
          { stmts: result.stmts <> StmtLeaf (GoAssign name (boxGoExpr codegenStateRef modNameStr result.expr result.exprType)) <> extras.stmts
          , expr: applyBoxed (GoVar name) extras.exprs, exprType: TypeValue, nextId: extras.nextId }

emitCurriedFold :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> ExprResult
emitCurriedFold translate context@{ codegenStateRef, modNameStr } nextId fn args =
  let
    translateArgument index argumentId arg =
      let
        result = translate (childContext context Nothing) argumentId arg
        boxed = boxGoExpr codegenStateRef modNameStr result.expr result.exprType
        -- Inspect the fresh-array marker before capture hides it behind a
        -- variable. Its existing ownership proof and normalization stay intact.
        expr = if index == 2 && isIntArrayFold fn args then
          normalizeFreshIntArrayRoundtrip (show result.nextId) boxed
        else boxed
      in result { expr = expr, exprType = TypeValue }
    captured = CallArguments.capture { prefix: "arrayIntrinsic_arg_", markUsed: false }
      translateArgument nextId (Array.take 3 args)
    loop = curriedFold context captured
    remaining = Array.drop 3 args
  in
    if Array.null remaining then
      { stmts: captured.stmts, expr: loop, exprType: TypeValue, nextId: captured.nextId }
    else
      let
        resultName = "arrayIntrinsic_result_" <> show captured.nextId
        capturedRemaining = CallArguments.capture { prefix: "arrayIntrinsic_extra_", markUsed: false }
          (\_ argumentId arg ->
            let result = translate (childContext context Nothing) argumentId arg
            in result { expr = boxGoExpr codegenStateRef modNameStr result.expr result.exprType, exprType = TypeValue })
          (captured.nextId + 1) remaining
      in
        { stmts: captured.stmts <> StmtLeaf (GoAssign resultName loop) <> capturedRemaining.stmts
        , expr: applyBoxed (GoVar resultName) capturedRemaining.exprs
        , exprType: TypeValue
        , nextId: capturedRemaining.nextId
        }

-- The general curried fold keeps its boxed accumulator and its existing
-- fresh-filter normalization. Map/filter use the shared native emitter above.
curriedFold :: ExprContext -> Arguments -> GoExpr
curriedFold { depth } accArgs =
  let
    { arrValName, arrGoName, resGoName, vName } = loopNames FoldlArray depth
    fExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 0)
    initExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 1)
    arrExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 2)
    loopBody = GoMutate resGoName (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply2") [ fExpr, GoVar resGoName, GoVar vName ])
    iifeBody = GoBlock
      [ GoAssign resGoName initExpr
      , GoAssign arrGoName (GoCall (rawGo "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
      , GoForRange ("_, " <> vName <> " := range *" <> arrGoName) [ loopBody ]
      , GoReturn (GoVar resGoName)
      ]
  in GoIIFE arrValName arrExpr iifeBody

-- Native arrays and callback signatures are preserved on this path.
emitUncurried :: TranslateExpr -> ExprContext -> Int -> Array TcoExpr -> ArrayIntrinsic -> ExprResult
emitUncurried translate context@{ codegenStateRef, modNameStr, depth } nextId args intrinsic =
  let
    { arrValName, arrGoName, resGoName, iName, vName } = loopNames intrinsic depth
    callbackArity = case intrinsic of
      FoldlArray -> 2
      _ -> 1
    mbFnArityInfo = Array.head args >>= Callback.worker context callbackArity
    sourceIndex = if callbackArity == 2 then 2 else 1
    requestedElement = Array.index args sourceIndex >>= scalarArrayElement context
    worker = case mbFnArityInfo of
      Just info -> GoRef info.reference
      Nothing -> rawGo "nil"
    -- Known native workers still evaluate their source callback argument once.
    -- Keep every captured native type, including the callback's result type.
    accArgs = CallArguments.capture { prefix: "arrayIntrinsic_arg_", markUsed: true }
      (\index argumentId arg -> case if index == 0 then Callback.nativeLambda translate context argumentId callbackArity arg else Nothing of
        Just native -> native
        Nothing ->
          let result = translate (childContext context Nothing) argumentId arg
          in case index == sourceIndex, requestedElement, result.expr of
            true, Just TypeInt64, GoUnboxIntArray source -> result { expr = source, exprType = TypeValue }
            _, _, _ -> result)
      nextId args
    capturedType = fromMaybe TypeValue (Array.index accArgs.exprTypes 0)
    nativeFunction = case capturedType of
      TypeFunc parameters _ -> Array.length parameters == callbackArity
      _ -> false
    directCallback = nativeFunction || case mbFnArityInfo of
      Just _ -> true
      _ -> false
    callCallback inputs inputTypes = case mbFnArityInfo of
      Just info ->
        { expr: GoCall worker (Array.zipWith (\value (Tuple actual target) -> coerceGoExpr codegenStateRef modNameStr value actual target) inputs (Array.zip inputTypes info.fArgs))
        , exprType: info.fRet }
      Nothing -> case capturedType of
        TypeFunc expected resultType | Array.length expected == callbackArity ->
          { expr: GoCall (fromMaybe (rawGo "nil") (Array.head accArgs.exprs))
              (Array.zipWith (\value (Tuple actual target) -> coerceGoExpr codegenStateRef modNameStr value actual target) inputs (Array.zip inputTypes expected))
          , exprType: resultType }
        _ ->
          { expr: applyBoxed (GoVar callbackName) (Array.zipWith (boxGoExpr codegenStateRef modNameStr) inputs inputTypes)
          , exprType: TypeValue }
    callbackName = "arrayIntrinsic_callback_" <> show accArgs.nextId
    capturedCallback = if directCallback then [] else
      [ GoAssign callbackName (boxGoExpr codegenStateRef modNameStr
          (fromMaybe (rawGo "nil") (Array.index accArgs.exprs 0))
          (fromMaybe TypeValue (Array.index accArgs.exprTypes 0)))
      ]
  in
    case intrinsic of
      MapArray ->
        let
          fExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 0)
          arrExprRaw = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 1)
          arrExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 1)

          source = arraySource arrValName arrGoName arrExprType
          elemType = source.elementType

          retType = case fExprType of
            TypeFunc _ ret -> ret
            _ -> case mbFnArityInfo of
              Just info -> info.fRet
              _ -> TypeValue

          finalRetType = TypeNativeArray retType

          inputType = fromMaybe elemType requestedElement
          input = coerceGoExpr codegenStateRef modNameStr (GoVar vName) elemType inputType
          call = callCallback [ input ] [ inputType ]
          loopBody = GoMutate (resGoName <> "[" <> iName <> "]")
            (coerceGoExpr codegenStateRef modNameStr call.expr call.exprType retType)

          iifeBodyStmts = capturedCallback <>
            [ source.assignment
            , GoAssign resGoName (GoCall (GoVar "make") [ rawGo ("[]" <> goTypeToStr retType), GoCall (GoVar "len") [ rawGo source.target ] ])
            , GoForRange (iName <> ", " <> vName <> " := range " <> source.target) [ loopBody ]
            ]
        in
          { stmts: accArgs.stmts, expr: GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) finalRetType) [], exprType: finalRetType, nextId: accArgs.nextId }

      FoldlArray ->
        let
          initExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 1)
          initExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 1)
          arrExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 2)
          arrExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 2)

          source = arraySource arrValName arrGoName arrExprType
          elemType = source.elementType

          inputType = fromMaybe elemType requestedElement
          input = coerceGoExpr codegenStateRef modNameStr (GoVar vName) elemType inputType
          call = callCallback [ GoVar resGoName, input ] [ initExprType, inputType ]
          loopBody = GoMutate resGoName (coerceGoExpr codegenStateRef modNameStr call.expr call.exprType initExprType)

          iifeBodyStmts = capturedCallback <>
            [ GoAssign resGoName initExpr
            , source.assignment
            , GoForRange ("_, " <> vName <> " := range " <> source.target) [ loopBody ]
            ]
          foldExpr = GoCall (GoFuncLit []
            (Array.cons (GoAssign arrValName arrExpr) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts))
            (GoVar resGoName) initExprType) []
        in
          { stmts: accArgs.stmts, expr: foldExpr, exprType: initExprType, nextId: accArgs.nextId }

      FilterArray ->
        let
          arrExprRaw = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 1)
          arrExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 1)

          source = arraySource arrValName arrGoName arrExprType
          elemType = source.elementType

          resultElement = fromMaybe elemType requestedElement
          value = coerceGoExpr codegenStateRef modNameStr (GoVar vName) elemType resultElement
          call = callCallback [ value ] [ resultElement ]
          isTrueExpr = coerceGoExpr codegenStateRef modNameStr call.expr call.exprType TypeBool
          -- Reserve a bounded buffer only on the first match. Reserving the
          -- entire input regresses sparse/all-false filters in time and memory.
          loopBody = GoIfElse isTrueExpr
            [ GoIfElse (GoBinOp "==" (GoCall (GoVar "len") [ GoVar resGoName ]) (GoInt 0))
                [ GoMutate resGoName (GoCall (GoVar "make")
                    [ rawGo ("[]" <> goTypeToStr resultElement), GoInt 0
                    , GoCall (GoVar "min") [ GoInt 256, GoCall (GoVar "len") [ rawGo source.target ] ] ]) ] []
            , GoMutate resGoName (GoCall (GoVar "append") [ GoVar resGoName, value ])
            ] []

          iifeBodyStmts = capturedCallback <>
            [ source.assignment
            , GoAssign resGoName (GoCall (GoVar "make") [ rawGo ("[]" <> goTypeToStr resultElement), rawGo "0" ])
            , GoForRange ("_, " <> vName <> " := range " <> source.target) [ loopBody ]
            ]
          filterExpr = GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) (TypeNativeArray resultElement)) []
          freshFilterExpr = if resultElement == TypeValue && Array.length args == 2 then GoFreshFilterArray filterExpr else filterExpr
        in
          { stmts: accArgs.stmts, expr: freshFilterExpr, exprType: TypeNativeArray resultElement, nextId: accArgs.nextId }

-- Scalar views are read-only: remove only an explicit Int-array copy and
-- normalize each selected element when producing the fresh result buffer.
scalarArrayElement :: ExprContext -> TcoExpr -> Maybe GoType
scalarArrayElement { metadata, modNameStr } expression = case getExprType expression of
  Array element ->
    let scalar = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr element
    in case scalar of
      TypeInt64 -> Just scalar
      TypeFloat64 -> Just scalar
      TypeString -> Just scalar
      TypeBool -> Just scalar
      _ -> Nothing
  _ -> Nothing

-- Require the resolved intrinsic and concrete scalar callback. Outer Typed
-- annotations on the seed can describe the enclosing application, so only a
-- literal establishes its Int type here; opaque/dynamic seeds keep the copies.
isIntArrayFold :: TcoExpr -> Array TcoExpr -> Boolean
isIntArrayFold fn args = case unwrapTcoExpr fn, args of
  Var (Qualified (Just (ModuleName "Data.Foldable")) (Ident "foldlArray")), [ callback, seed, _ ] ->
    case extractFuncType callback, unwrapTcoExpr seed of
      Just { fArgs: [ Int, Int ], fRet: Int }, Lit (LitInt _) -> true
      _, _ -> false
  _, _ -> false

normalizeFreshIntArrayRoundtrip :: String -> GoExpr -> GoExpr
normalizeFreshIntArrayRoundtrip suffix expr = case expr of
  GoBoxIntArray (GoUnboxIntArray (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoFreshFilterArray filtered ])) ->
    let
      sourceName = "source_int_array_" <> suffix
      itemsName = "items_int_array_" <> suffix
      indexName = "i_int_array_" <> suffix
      valueName = "v_int_array_" <> suffix
      source = GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ filtered ]
      -- The source runs once, before these IIFE-local names enter scope. The
      -- filter's make/append owns this buffer; do not search through variables,
      -- calls or storage for a marker. Normalize every Value before the fold
      -- to preserve IntVal/tag/pointer semantics of the two original copies.
      body = GoBlock
        [ GoAssign itemsName (GoCall (rawGo "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar sourceName) "UnsafePtr" ])
        , GoForRange (indexName <> ", " <> valueName <> " := range *" <> itemsName)
            [ GoMutate ("(*" <> itemsName <> ")[" <> indexName <> "]")
                (GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoSelector (GoVar valueName) "IntVal" ])
            ]
        , GoReturn (GoVar sourceName)
        ]
    in
      GoIIFE sourceName source body
  _ -> expr
