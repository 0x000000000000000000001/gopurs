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
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Gopurs.ArrayIntrinsics.Index (safeIndex, unsafeIndex) as Index
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.CallAnalysis (CallTarget, qualifiedTarget)
import Gopurs.CallArguments (Arguments, applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.ExprAnalysis (extractFuncType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), childContext)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..), goTypeToStr)
import Gopurs.GoConversions (boxGoExpr, unboxGoExpr)
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

loopNames :: Convention -> ArrayIntrinsic -> Int -> { iifeName :: String, arrValName :: String, arrGoName :: String, resGoName :: String, iName :: String, vName :: String }
loopNames convention intrinsic depth =
  let
    name = case intrinsic, convention of
      MapArray, _ -> "arrayMap"
      FoldlArray, _ -> "foldlArray"
      FilterArray, Curried -> "filter"
      FilterArray, Uncurried -> "filterImpl"
    iifeName = name <> show depth
  in
    { iifeName, arrValName: "arr_val_" <> iifeName, arrGoName: "arr_go_" <> iifeName
    , resGoName: "res_go_" <> iifeName, iName: "i_" <> iifeName, vName: "v_" <> iifeName
    }

-- Capture the intrinsic's strict prefix in source order. The callback must
-- be constructed once, including when the input array is empty. Surplus
-- curried arguments belong to the application of the intrinsic's result.
emitCurried :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> ArrayIntrinsic -> ExprResult
emitCurried translate context@{ codegenStateRef, modNameStr } nextId fn args intrinsic =
  let
    arity = case intrinsic of
      FoldlArray -> 3
      _ -> 2
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
      translateArgument nextId (Array.take arity args)
    loop = curriedLoop context intrinsic captured
    remaining = Array.drop arity args
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

curriedLoop :: ExprContext -> ArrayIntrinsic -> Arguments -> GoExpr
curriedLoop { depth } intrinsic accArgs =
  let
    { arrValName, arrGoName, resGoName, iName, vName } = loopNames Curried intrinsic depth
  in
    case intrinsic of
      MapArray ->
        let
          fExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 0)
          arrExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 1)
          loopBody = GoMutate (resGoName <> "[" <> iName <> "]") (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, GoVar vName ])
          iifeBody = GoBlock
            [ GoAssign arrGoName (GoCall (rawGo "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
            , GoAssign resGoName (GoCall (GoVar "make") [ rawGo "[]gopurs_runtime.Value", GoCall (GoVar "len") [ rawGo ("*" <> arrGoName) ] ])
            , GoForRange (iName <> ", " <> vName <> " := range *" <> arrGoName) [ loopBody ]
            , GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoVar resGoName ])
            ]
        in
          GoIIFE arrValName arrExpr iifeBody

      FoldlArray ->
        let
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
        in
          GoIIFE arrValName arrExpr iifeBody

      FilterArray ->
        let
          fExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 0)
          arrExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 1)
          condExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, GoVar vName ]
          isTrueExpr = GoCall (GoSelector condExpr "BoolVal") []
          loopBody = GoIfElse isTrueExpr [ GoMutate resGoName (GoCall (GoVar "append") [ GoVar resGoName, GoVar vName ]) ] []
          iifeBody = GoBlock
            [ GoAssign arrGoName (GoCall (rawGo "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
            , GoAssign resGoName (GoCall (GoVar "make") [ rawGo "[]gopurs_runtime.Value", rawGo "0" ])
            , GoForRange ("_, " <> vName <> " := range *" <> arrGoName) [ loopBody ]
            , GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoVar resGoName ])
            ]
        in
          GoIIFE arrValName arrExpr iifeBody

-- Native arrays and callback signatures are preserved on this path.
emitUncurried :: TranslateExpr -> ExprContext -> Int -> Array TcoExpr -> ArrayIntrinsic -> ExprResult
emitUncurried translate context@{ codegenStateRef, modNameStr, depth } nextId args intrinsic =
  let
    { arrValName, arrGoName, resGoName, iName, vName } = loopNames Uncurried intrinsic depth
    { worker, info: mbFnArityInfo } = callbackInfo context args
    callbackArity = case intrinsic of
      FoldlArray -> 2
      _ -> 1
    directCallback = case mbFnArityInfo of
      Just info -> info.arity == callbackArity
      _ -> false
    -- Known native workers still evaluate their source callback argument once.
    -- Keep every captured native type, including the callback's result type.
    accArgs = CallArguments.capture { prefix: "arrayIntrinsic_arg_", markUsed: true }
      (\_ argumentId arg -> translate (childContext context Nothing) argumentId arg)
      nextId args
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

          loopBody = case mbFnArityInfo of
            Just info | info.arity == 1 ->
              let
                expectedArgType = fromMaybe TypeValue (Array.index info.fArgs 0)
                value = unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArgType
                call = GoCall worker [ value ]
              in
                GoMutate (resGoName <> "[" <> iName <> "]") (unboxGoExpr codegenStateRef modNameStr call info.fRet retType)
            _ ->
              let call = applyBoxed (GoVar callbackName)
                    [ boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]
              in GoMutate (resGoName <> "[" <> iName <> "]") (unboxGoExpr codegenStateRef modNameStr call TypeValue retType)

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

          loopBody = case mbFnArityInfo of
            Just info | info.arity == 2 ->
              let
                expectedArg0 = fromMaybe TypeValue (Array.index info.fArgs 0)
                expectedArg1 = fromMaybe TypeValue (Array.index info.fArgs 1)
                accumulator = unboxGoExpr codegenStateRef modNameStr (GoVar resGoName) initExprType expectedArg0
                value = unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArg1
                call = GoCall worker [ accumulator, value ]
              in
                GoMutate resGoName (unboxGoExpr codegenStateRef modNameStr call info.fRet initExprType)
            _ ->
              let call = applyBoxed (GoVar callbackName)
                    [ boxGoExpr codegenStateRef modNameStr (GoVar resGoName) initExprType
                    , boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType
                    ]
              in GoMutate resGoName (unboxGoExpr codegenStateRef modNameStr call TypeValue initExprType)

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

          isTrueExpr = case mbFnArityInfo of
            Just info | info.arity == 1 ->
              let
                expectedArgType = fromMaybe TypeValue (Array.index info.fArgs 0)
              in
                GoCall worker [ unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArgType ]
            _ ->
              let
                condExpr = applyBoxed (GoVar callbackName)
                  [ boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]
              in
                GoCall (GoSelector condExpr "BoolVal") []

          loopBody = GoIfElse isTrueExpr [ GoMutate resGoName (GoCall (GoVar "append") [ GoVar resGoName, GoVar vName ]) ] []

          iifeBodyStmts = capturedCallback <>
            [ source.assignment
            , GoAssign resGoName (GoCall (GoVar "make") [ rawGo ("[]" <> goTypeToStr elemType), rawGo "0" ])
            , GoForRange ("_, " <> vName <> " := range " <> source.target) [ loopBody ]
            ]
          filterExpr = GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) (TypeNativeArray elemType)) []
          freshFilterExpr = if elemType == TypeValue && Array.length args == 2 then GoFreshFilterArray filterExpr else filterExpr
        in
          { stmts: accArgs.stmts, expr: freshFilterExpr, exprType: TypeNativeArray elemType, nextId: accArgs.nextId }

callbackInfo :: ExprContext -> Array TcoExpr -> { worker :: GoExpr, info :: Maybe FunctionInfo }
callbackInfo { modNameStr, moduleFunctions } args =
  let
    target = Array.index args 0 >>= qualifiedTarget
    fullName = case target of
      Just { mbMod: Just (ModuleName mn), name } -> String.replaceAll (Pattern ".") (Replacement "_") mn <> "." <> name
      Just { mbMod: Nothing, name } -> modNameStr <> "." <> name
      Nothing -> ""
  in
    { worker: GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fullName)
    , info: Map.lookup fullName moduleFunctions
    }

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
