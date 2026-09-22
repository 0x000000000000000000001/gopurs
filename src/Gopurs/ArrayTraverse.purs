module Gopurs.ArrayTraverse
  ( emit
  ) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NonEmptyArray
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.AdtExprs as AdtExprs
import Gopurs.CallAnalysis (collectCurriedAbs, collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments (applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (extractFuncType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, flattenStmts)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(GetProp), BackendSyntax(Abs, Accessor))

-- The known Array/Either instances permit a strict loop instead of building
-- intermediate arrays through generic applicative composition. PBO may expose
-- either the class method, its static dictionary accessor, or the FFI worker.
-- Every form must prove the same dictionaries; arbitrary Applicatives stay on
-- their existing traversal path.
emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
emit translate context nextId fn args = do
  traversal <- recognize fn args
  guard (Array.length traversal.args >= traversal.callbackOffset)
  guard (Array.length traversal.args <= traversal.callbackOffset + 2)
  pure (emitKnownEither translate context nextId traversal.indexed traversal.callbackOffset traversal.args)

type Traversal = { indexed :: Boolean, callbackOffset :: Int, args :: Array TcoExpr }

recognize :: TcoExpr -> Array TcoExpr -> Maybe Traversal
recognize fn args = case unwrapTcoExpr fn of
  Accessor dictionary (GetProp "traverse") -> do
    guard (isNamed "Data.Traversable" "traversableArray" dictionary)
    applicative <- Array.index args 0
    guard (isNamed "Data.Either" "applicativeEither" applicative)
    -- Retain the dictionary getter before the supplied arguments, just as for
    -- an ordinary method call. Reading this known field has no extra effects.
    pure { indexed: false, callbackOffset: 2, args: Array.cons dictionary args }
  _ -> do
    { mbMod, name } <- qualifiedTarget fn
    if mbMod == Just (ModuleName "Data.Traversable") && name == "traverseArrayImpl" then do
      applyFn <- Array.index args 0
      mapFn <- Array.index args 1
      pureFn <- Array.index args 2
      appendFn <- Array.index args 3
      guard (isMethod "Control.Apply" "apply" "Data.Either" "applyEither" applyFn)
      guard (isMethod "Data.Functor" "map" "Data.Either" "functorEither" mapFn)
      guard (isMethod "Control.Applicative" "pure" "Data.Either" "applicativeEither" pureFn)
      guard (isMethod "Data.Semigroup" "append" "Data.Semigroup" "semigroupArray" appendFn)
      pure { indexed: false, callbackOffset: 4, args }
    else do
      let indexed = mbMod == Just (ModuleName "Data.TraversableWithIndex")
      guard (if indexed then name == "traverseWithIndexDefault" || name == "traverseWithIndex"
        else mbMod == Just (ModuleName "Data.Traversable") && name == "traverse")
      dictionary <- Array.index args 0
      applicative <- Array.index args 1
      guard (if indexed then isNamed "Data.TraversableWithIndex" "traversableWithIndexArray" dictionary
        else isNamed "Data.Traversable" "traversableArray" dictionary)
      guard (isNamed "Data.Either" "applicativeEither" applicative)
      pure { indexed, callbackOffset: 2, args }

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

emitKnownEither :: TranslateExpr -> ExprContext -> Int -> Boolean -> Int -> Array TcoExpr -> ExprResult
emitKnownEither translate context@{ metadata, codegenStateRef, modNameStr } nextId indexed callbackOffset args =
  let
    -- Capture each expression before translating the next argument's statements.
    -- Dictionary getter evaluation and callback construction remain at the same
    -- application stage even when this call returns a closure.
    captured = foldl capture { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId } args
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
    callbackName = "traverseEither_callback_" <> suffix
    inputName = "traverseEither_input_" <> suffix
    valuesName = "traverseEither_values_" <> suffix
    outputName = "traverseEither_output_" <> suffix
    indexName = "traverseEither_index_" <> suffix
    valueName = "traverseEither_value_" <> suffix
    resultName = "traverseEither_result_" <> suffix
    errorName = "traverseEither_firstError_" <> suffix
    failedName = "traverseEither_failed_" <> suffix
    input = if hasInput then argument (callbackOffset + 1) else GoVar inputName
    inputType = if hasInput then argumentType (callbackOffset + 1) else TypeValue
    source = case inputType of
      TypeNativeArray elementType -> { expr: input, elementType }
      _ ->
        { expr: rawArray input
        , elementType: TypeValue
        }
    eitherModule = Just (ModuleName "Data.Either")
    eitherType = ADT "Data.Either.Either" [ "Data", "Either", "Either" ] [ Any, Any ]
    result = { expr: GoVar resultName, exprType: callbackResultType, stmts: StmtEmpty, nextId: captured.nextId }
    tag name = (AdtExprs.isTag metadata codegenStateRef modNameStr eitherModule name result).expr
    payload = AdtExprs.getField metadata codegenStateRef modNameStr
      { moduleName: eitherModule, ctorName: "Right", index: 0 }
      { expr: GoVar resultName, exprType: callbackResultType, sourceType: eitherType }
    prepared = AdtExprs.prepareSaturated metadata modNameStr eitherModule "Right" eitherType
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
      , GoIfElse (tag "Right")
          [ GoMutate (outputName <> "[" <> indexName <> "]") (boxGoExpr codegenStateRef modNameStr payload.expr payload.exprType) ]
          [ GoIfElse (tag "Left")
              [ GoIfElse (GoPrefixOp "!" (GoVar failedName))
                  [ GoMutate errorName (boxGoExpr codegenStateRef modNameStr (GoVar resultName) callbackResultType)
                  , GoMutate failedName (rawGo "true")
                  ] []
              ]
              [ GoCall (GoVar "panic") [ GoString "Failed pattern match" ] ]
          ]
      ]
    -- Keep invoking the callback after a Left. The original strict map evaluates
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
  capture acc arg =
    let
      result = case if Array.length acc.exprs == callbackOffset then nativeCallback translate context acc.nextId indexed arg else Nothing of
        Just native -> native
        Nothing -> translate (CallArguments.childContext context Nothing) acc.nextId arg
      name = "traverseEither_arg_" <> show result.nextId
    in
      { stmts: acc.stmts <> result.stmts <> StmtLeaf (GoAssign name result.expr) <> StmtLeaf (GoMutate "_" (GoVar name))
      , exprs: Array.snoc acc.exprs (GoVar name)
      , exprTypes: Array.snoc acc.exprTypes result.exprType
      , nextId: result.nextId + 1
      }

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
      bound = foldl (\acc (Tuple name goType) -> Map.insert name { name, goType } acc) context.bound params
      expectedResult = map _.fRet (extractFuncType callback)
      result = translate
        (context
          { depth = context.depth + 1
          , bound = bound
          , tcoIdent = Nothing
          , loopCtx = []
          , options = { isTail: true, inEffectBlock: false }
          , mbExpectedExprType = expectedResult
          }) nextId grouped.body
    pure
      { stmts: StmtEmpty
      , expr: GoFuncLit params (flattenStmts result.stmts) result.expr result.exprType
      , exprType: TypeFunc parameterTypes result.exprType
      , nextId: result.nextId
      }
  _ -> Nothing
