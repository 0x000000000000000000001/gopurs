module Gopurs.ArrayTraverse
  ( emit
  , recognizeMaybe
  ) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NonEmptyArray
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.AdtExprs as AdtExprs
import Gopurs.ArrayTraverseTotal as ArrayTraverseTotal
import Gopurs.CallAnalysis (collectCurriedAbs, collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments (Arguments, applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (extractFuncType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, bindParameters, childContext, flattenStmts)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(GetProp), BackendSyntax(Abs, Accessor))

-- The known Array/Either and Array/Maybe instances permit a strict loop instead of building
-- intermediate arrays through generic applicative composition. PBO may expose
-- either the class method, its static dictionary accessor, or the FFI worker.
-- Every form must prove the same dictionaries; arbitrary Applicatives stay on
-- their existing traversal path.
emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
emit translate context nextId fn args = do
  traversal <- recognize fn args
  pure case if traversal.plan.moduleName == "Data.Maybe" then
    ArrayTraverseTotal.emit translate context nextId traversal.indexed traversal.callbackOffset traversal.args
    else Nothing of
      Just result -> result
      Nothing -> emitLoop context traversal (captureArguments translate context nextId traversal)

type ApplicativePlan =
  { moduleName :: String
  , sourceType :: ExprType
  , success :: String
  , failure :: String
  , prefix :: String
  }

type Traversal = { indexed :: Boolean, callbackOffset :: Int, args :: Array TcoExpr, plan :: ApplicativePlan }

-- Consumers reuse the exact canonical dictionary proof, not the spelling of a
-- traversal method alone. They impose their own saturation/phase restrictions.
recognizeMaybe :: TcoExpr -> Array TcoExpr -> Maybe { indexed :: Boolean, callbackOffset :: Int, args :: Array TcoExpr }
recognizeMaybe fn args = do
  traversal <- recognize fn args
  guard (traversal.plan.moduleName == "Data.Maybe")
  pure { indexed: traversal.indexed, callbackOffset: traversal.callbackOffset, args: traversal.args }

-- These are the canonical library constructors, never inferred from a custom
-- dictionary's shape. The regular ADT emitter handles nullable Maybe layouts.
applicativePlan :: TcoExpr -> Maybe ApplicativePlan
applicativePlan dictionary
  | isNamed "Data.Either" "applicativeEither" dictionary = Just eitherPlan
  | isNamed "Data.Maybe" "applicativeMaybe" dictionary = Just maybePlan
  | otherwise = Nothing

eitherPlan :: ApplicativePlan
eitherPlan =
  { moduleName: "Data.Either"
  , sourceType: ADT "Data.Either.Either" [ "Data", "Either", "Either" ] [ Any, Any ]
  , success: "Right", failure: "Left", prefix: "traverseEither"
  }

maybePlan :: ApplicativePlan
maybePlan =
  { moduleName: "Data.Maybe"
  , sourceType: ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ Any ]
  , success: "Just", failure: "Nothing", prefix: "traverseMaybe"
  }

ffiPlan :: TcoExpr -> TcoExpr -> TcoExpr -> Maybe ApplicativePlan
ffiPlan applyFn mapFn pureFn =
  let
    matches mod suffix = isMethod "Control.Apply" "apply" mod ("apply" <> suffix) applyFn
      && isMethod "Data.Functor" "map" mod ("functor" <> suffix) mapFn
      && isMethod "Control.Applicative" "pure" mod ("applicative" <> suffix) pureFn
  in
    if matches "Data.Either" "Either" then Just eitherPlan
    else if matches "Data.Maybe" "Maybe" then Just maybePlan
    else Nothing

recognize :: TcoExpr -> Array TcoExpr -> Maybe Traversal
recognize fn args = do
  traversal <- recognizeForm fn args
  guard (Array.length traversal.args >= traversal.callbackOffset)
  guard (Array.length traversal.args <= traversal.callbackOffset + 2)
  pure traversal

recognizeForm :: TcoExpr -> Array TcoExpr -> Maybe Traversal
recognizeForm fn args = case unwrapTcoExpr fn of
  Accessor dictionary (GetProp "traverse") -> do
    guard (isNamed "Data.Traversable" "traversableArray" dictionary)
    applicative <- Array.index args 0
    plan <- applicativePlan applicative
    -- Retain the dictionary getter before the supplied arguments, just as for
    -- an ordinary method call. Reading this known field has no extra effects.
    pure { indexed: false, callbackOffset: 2, args: Array.cons dictionary args, plan }
  _ -> do
    { mbMod, name } <- qualifiedTarget fn
    if mbMod == Just (ModuleName "Data.Traversable") && name == "traverseArrayImpl" then do
      applyFn <- Array.index args 0
      mapFn <- Array.index args 1
      pureFn <- Array.index args 2
      appendFn <- Array.index args 3
      plan <- ffiPlan applyFn mapFn pureFn
      guard (isMethod "Data.Semigroup" "append" "Data.Semigroup" "semigroupArray" appendFn)
      pure { indexed: false, callbackOffset: 4, args, plan }
    else do
      let indexed = mbMod == Just (ModuleName "Data.TraversableWithIndex")
      guard (if indexed then name == "traverseWithIndexDefault" || name == "traverseWithIndex"
        else mbMod == Just (ModuleName "Data.Traversable") && name == "traverse")
      dictionary <- Array.index args 0
      applicative <- Array.index args 1
      guard (if indexed then isNamed "Data.TraversableWithIndex" "traversableWithIndexArray" dictionary
        else isNamed "Data.Traversable" "traversableArray" dictionary)
      plan <- applicativePlan applicative
      pure { indexed, callbackOffset: 2, args, plan }

isNamed :: String -> String -> TcoExpr -> Boolean
isNamed mod name expr = case qualifiedTarget expr of
  Just target -> target.mbMod == Just (ModuleName mod) && target.name == name
  Nothing -> false

isMethod :: String -> String -> String -> String -> TcoExpr -> Boolean
isMethod mod name dictMod dictName expression =
  let
    Tuple fn spine = collectGoSpine expression
    args = getGoSpineArgs spine
  in isNamed mod name fn && case args of
    [ dictionary ] -> isNamed dictMod dictName dictionary
    _ -> false

-- Dictionary getters and callback construction run at the current application
-- stage, including when the input is absent and the result is a reusable closure.
-- Only the callback position may bypass the ordinary boxed function ABI.
captureArguments :: TranslateExpr -> ExprContext -> Int -> Traversal -> Arguments
captureArguments translate context nextId traversal =
  CallArguments.capture { prefix: traversal.plan.prefix <> "_arg_", markUsed: true }
    translateArgument nextId traversal.args
  where
  translateArgument index currentId arg =
    case if index == traversal.callbackOffset then nativeCallback translate context currentId traversal.indexed arg else Nothing of
      Just native -> native
      Nothing -> translate (childContext context Nothing) currentId arg

emitLoop :: ExprContext -> Traversal -> Arguments -> ExprResult
emitLoop { metadata, codegenStateRef, modNameStr } { indexed, callbackOffset, args, plan } captured =
  let
    argument index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argumentType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    hasCallback = Array.length args > callbackOffset
    hasInput = Array.length args > callbackOffset + 1
    callbackType = if hasCallback then argumentType callbackOffset else TypeValue
    callback = if hasCallback then argument callbackOffset else GoVar callbackName
    callbackResultType = case callbackType of
      TypeFunc _ resultType -> resultType
      _ -> TypeValue
    suffix = show captured.nextId
    callbackName = plan.prefix <> "_callback_" <> suffix
    inputName = plan.prefix <> "_input_" <> suffix
    valuesName = plan.prefix <> "_values_" <> suffix
    outputName = plan.prefix <> "_output_" <> suffix
    indexName = plan.prefix <> "_index_" <> suffix
    valueName = plan.prefix <> "_value_" <> suffix
    resultName = plan.prefix <> "_result_" <> suffix
    errorName = plan.prefix <> "_firstError_" <> suffix
    failedName = plan.prefix <> "_failed_" <> suffix
    input = if hasInput then argument (callbackOffset + 1) else GoVar inputName
    inputType = if hasInput then argumentType (callbackOffset + 1) else TypeValue
    source = case inputType of
      TypeNativeArray elementType -> { expr: input, elementType }
      _ ->
        { expr: rawArray input
        , elementType: TypeValue
        }
    adtModule = Just (ModuleName plan.moduleName)
    result = { expr: GoVar resultName, exprType: callbackResultType, stmts: StmtEmpty, nextId: captured.nextId }
    tag name = (AdtExprs.isTag metadata codegenStateRef modNameStr adtModule name result).expr
    payload = AdtExprs.getField metadata codegenStateRef modNameStr
      { moduleName: adtModule, ctorName: plan.success, index: 0 }
      { expr: GoVar resultName, exprType: callbackResultType, sourceType: plan.sourceType }
    prepared = AdtExprs.prepareSaturated metadata modNameStr adtModule plan.success plan.sourceType
    arrayValue = boxGoExpr codegenStateRef modNameStr (GoVar outputName) (TypeNativeArray TypeValue)
    success = AdtExprs.saturated codegenStateRef prepared { exprs: [ arrayValue ], exprTypes: [ TypeValue ] }
    boxedSuccess = boxGoExpr codegenStateRef modNameStr success.expr success.exprType
    indexExpr = GoCall (GoVar "int64") [ GoVar indexName ]
    value = boxGoExpr codegenStateRef modNameStr (GoVar valueName) source.elementType
    callbackCall = case callbackType of
      TypeFunc _ _ -> GoCall callback (if indexed then [ indexExpr, value ] else [ value ])
      _ -> applyBoxed callback
        (if indexed then [ GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ indexExpr ], value ] else [ value ])
    loop =
      [ GoAssign resultName callbackCall
      , GoIfElse (tag plan.success)
          [ GoMutate (outputName <> "[" <> indexName <> "]") (boxGoExpr codegenStateRef modNameStr payload.expr payload.exprType) ]
          [ GoIfElse (tag plan.failure)
              [ GoIfElse (GoPrefixOp "!" (GoVar failedName))
                  [ GoMutate errorName (boxGoExpr codegenStateRef modNameStr (GoVar resultName) callbackResultType)
                  , GoMutate failedName (rawGo "true")
                  ] []
              ]
              [ GoCall (GoVar "panic") [ GoString "Failed pattern match" ] ]
          ]
      ]
    -- Keep invoking the callback after Left/Nothing. The original strict map evaluates
    -- every element before sequence selects the first error.
    body =
      [ GoAssign valuesName source.expr
      , GoAssign outputName (GoCall (GoVar "make") [ rawGo "[]gopurs_runtime.Value", GoCall (GoVar "len") [ GoVar valuesName ] ])
      , GoAssign errorName (rawGo "gopurs_runtime.Value{}")
      , GoAssign failedName (rawGo "false")
      , GoForRange (indexName <> ", " <> valueName <> " := range " <> valuesName) loop
      , GoIfElse (GoVar failedName) [ GoReturn (GoVar errorName) ] []
      ]
    expression = if not hasCallback then
      GoCall (GoSelector (GoVar "gopurs_runtime") "Func2")
        [ GoFuncLit [ Tuple callbackName TypeValue, Tuple inputName TypeValue ] body boxedSuccess TypeValue ]
      else if not hasInput then
      GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
        [ GoFuncLit [ Tuple inputName TypeValue ] body boxedSuccess TypeValue ]
      else GoCall (GoFuncLit [] body boxedSuccess TypeValue) []
  in
    { stmts: captured.stmts
    , expr: expression
    , exprType: TypeValue
    , nextId: captured.nextId + 1
    }
  where
  rawArray input = GoPrefixOp "*" (GoCall (rawGo "(*[]gopurs_runtime.Value)") [ GoSelector input "UnsafePtr" ])

-- This callback is consumed only by the known traversal, so it needs no boxed
-- function ABI. Adjacent lambdas are values: capturing them here preserves the
-- original evaluation stage, including when the traversal is partially applied.
-- Do not cross a let/call between lambdas: its evaluation belongs to each first
-- application. The ordinary callback path retains that behavior.
nativeCallback :: TranslateExpr -> ExprContext -> Int -> Boolean -> TcoExpr -> Maybe ExprResult
nativeCallback translate context nextId indexed callback = case unwrapTcoExpr callback of
  Abs args body -> do
    let
      grouped = collectCurriedAbs args body
      parameterTypes = if indexed then [ TypeInt64, TypeValue ] else [ TypeValue ]
    guard (NonEmptyArray.length grouped.args == Array.length parameterTypes)
    let
      names = map (\(Tuple ident level) -> localId ident level) (NonEmptyArray.toArray grouped.args)
      params = Array.zip names parameterTypes
      bound = bindParameters params context.bound
      expectedResult = map _.fRet (extractFuncType callback)
      result = translate
        (context
          { depth = context.depth + 1
          , bound = bound
          , tcoIdent = Nothing
          , loopCtx = []
          , options = { isTail: true, inEffectBlock: false }
          , mbExpectedExprType = expectedResult
          , mbExpectedGoType = Nothing
          }) nextId grouped.body
    pure
      { stmts: StmtEmpty
      , expr: GoFuncLit params (flattenStmts result.stmts) result.expr result.exprType
      , exprType: TypeFunc parameterTypes result.exprType
      , nextId: result.nextId
      }
  _ -> Nothing
