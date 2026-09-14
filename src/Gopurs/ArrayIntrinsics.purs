module Gopurs.ArrayIntrinsics
  ( ArrayIntrinsic
  , Convention(..)
  , recognize
  , emitCurried
  , emitUncurried
  ) where

import Prelude
import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Gopurs.CallAnalysis (CallTarget, qualifiedTarget)
import Gopurs.CallArguments (Arguments, applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.ExprAnalysis (extractFuncType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult)
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

-- Arguments are already boxed in source order by the caller.
emitCurried :: ExprContext -> TcoExpr -> Array TcoExpr -> ArrayIntrinsic -> Arguments -> ExprResult
emitCurried context@{ depth } fn args intrinsic accArgs =
  let
    { iifeName, arrValName, arrGoName, resGoName, iName, vName } = loopNames Curried intrinsic depth

    iifeExpr = case intrinsic of
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
          boxedArrExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 2)
          arrExpr =
            if isIntArrayFold fn args then
              normalizeFreshIntArrayRoundtrip (iifeName <> "_" <> show accArgs.nextId) boxedArrExpr
            else boxedArrExpr
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

    arity = case intrinsic of
      FoldlArray -> 3
      _ -> 2
    accArgsRemainingBoxed = CallArguments.boxRemaining context arity accArgs

    finalExpr = applyBoxed iifeExpr accArgsRemainingBoxed
  in
    { stmts: accArgs.stmts, expr: finalExpr, exprType: TypeValue, nextId: accArgs.nextId }

-- Native arrays and callback signatures are preserved on this path.
emitUncurried :: ExprContext -> Array TcoExpr -> ArrayIntrinsic -> Arguments -> ExprResult
emitUncurried context@{ codegenStateRef, modNameStr, depth } args intrinsic accArgs =
  let
    { arrValName, arrGoName, resGoName, iName, vName } = loopNames Uncurried intrinsic depth
    { fullName: fnFullName, info: mbFnArityInfo } = callbackInfo context args

    iifeExpr = case intrinsic of
      MapArray ->
        let
          fExprRaw = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 0)
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

          arrGoAssignment = source.assignment
          arrGoRangeTarget = source.target

          loopBody = case mbFnArityInfo of
            Just info | info.arity == 1 ->
              let
                expectedArgType = fromMaybe TypeValue (Array.index info.fArgs 0)
              in
                GoMutate (resGoName <> "[" <> iName <> "]") (unboxGoExpr codegenStateRef modNameStr (GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArgType ]) info.fRet retType)
            _ ->
              GoMutate (resGoName <> "[" <> iName <> "]") (unboxGoExpr codegenStateRef modNameStr (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ boxGoExpr codegenStateRef modNameStr fExprRaw fExprType, boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]) TypeValue retType)

          iifeBodyStmts =
            [ arrGoAssignment
            , GoAssign resGoName (GoCall (GoVar "make") [ rawGo ("[]" <> goTypeToStr retType), GoCall (GoVar "len") [ rawGo arrGoRangeTarget ] ])
            , GoForRange (iName <> ", " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
            ]
        in
          { stmts: accArgs.stmts, expr: GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) finalRetType) [], exprType: finalRetType, nextId: accArgs.nextId }

      FoldlArray ->
        let
          fExpr = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 0)
          fExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 0)
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
              in
                GoMutate resGoName (unboxGoExpr codegenStateRef modNameStr (GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr codegenStateRef modNameStr (GoVar resGoName) initExprType expectedArg0, unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArg1 ]) info.fRet initExprType)
            _ ->
              GoMutate resGoName (unboxGoExpr codegenStateRef modNameStr (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply2") [ boxGoExpr codegenStateRef modNameStr fExpr fExprType, boxGoExpr codegenStateRef modNameStr (GoVar resGoName) initExprType, boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]) TypeValue initExprType)

          arrGoAssignment = source.assignment
          arrGoRangeTarget = source.target

          iifeBody = GoBlock
            [ GoAssign resGoName initExpr
            , arrGoAssignment
            , GoForRange ("_, " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
            , GoReturn (GoVar resGoName)
            ]
        in
          { stmts: accArgs.stmts, expr: GoIIFE arrValName arrExpr iifeBody, exprType: initExprType, nextId: accArgs.nextId }

      FilterArray ->
        let
          fExprRaw = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 0)
          fExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 0)
          arrExprRaw = fromMaybe (rawGo "nil") (Array.index accArgs.exprs 1)
          arrExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 1)

          source = arraySource arrValName arrGoName arrExprType
          elemType = source.elementType

          isTrueExpr = case mbFnArityInfo of
            Just info | info.arity == 1 ->
              let
                expectedArgType = fromMaybe TypeValue (Array.index info.fArgs 0)
              in
                GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArgType ]
            _ ->
              let
                condExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ boxGoExpr codegenStateRef modNameStr fExprRaw fExprType, boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]
              in
                GoCall (GoSelector condExpr "BoolVal") []

          loopBody = GoIfElse isTrueExpr [ GoMutate resGoName (GoCall (GoVar "append") [ GoVar resGoName, GoVar vName ]) ] []

          arrGoAssignment = source.assignment
          arrGoRangeTarget = source.target

          iifeBodyStmts =
            [ arrGoAssignment
            , GoAssign resGoName (GoCall (GoVar "make") [ rawGo ("[]" <> goTypeToStr elemType), rawGo "0" ])
            , GoForRange ("_, " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
            ]
          filterExpr = GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) (TypeNativeArray elemType)) []
          freshFilterExpr = if elemType == TypeValue && Array.length args == 2 then GoFreshFilterArray filterExpr else filterExpr
        in
          { stmts: accArgs.stmts, expr: freshFilterExpr, exprType: TypeNativeArray elemType, nextId: accArgs.nextId }

  in
    iifeExpr

callbackInfo :: ExprContext -> Array TcoExpr -> { fullName :: String, info :: Maybe FunctionInfo }
callbackInfo { modNameStr, moduleFunctions } args =
  let
    target = Array.index args 0 >>= qualifiedTarget
    fullName = case target of
      Just { mbMod: Just (ModuleName mn), name } -> String.replaceAll (Pattern ".") (Replacement "_") mn <> "." <> name
      Just { mbMod: Nothing, name } -> modNameStr <> "." <> name
      Nothing -> ""
  in
    { fullName, info: Map.lookup fullName moduleFunctions }

arraySource :: String -> String -> GoType -> { assignment :: GoExpr, target :: String, elementType :: GoType }
arraySource valueName arrayName = case _ of
  TypeNativeArray inner ->
    { assignment: GoAssign arrayName (GoVar valueName), target: arrayName, elementType: inner }
  _ ->
    { assignment: GoAssign arrayName (GoCall (rawGo "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar valueName) "UnsafePtr" ])
    , target: "*" <> arrayName, elementType: TypeValue
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

