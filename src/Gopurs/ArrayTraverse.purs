module Gopurs.ArrayTraverse
  ( emit
  ) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Foldable (foldl)
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.AdtExprs as AdtExprs
import Gopurs.CallAnalysis (CallTarget, qualifiedTarget)
import Gopurs.CallArguments (applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName(..))

-- Only the standard Array / Either dictionaries justify replacing mapWithIndex
-- followed by generic sequence. An arbitrary Applicative must keep its traversal.
-- Three supplied arguments create the same reusable unary closure as the source;
-- four supplied arguments execute it immediately.
emit :: TranslateExpr -> ExprContext -> Int -> Maybe CallTarget -> Array TcoExpr -> Maybe ExprResult
emit translate context nextId target args = do
  { mbMod, name } <- target
  guard (mbMod == Just (ModuleName "Data.TraversableWithIndex"))
  guard (name == "traverseWithIndexDefault" || name == "traverseWithIndex")
  guard (Array.length args == 3 || Array.length args == 4)
  traversal <- Array.index args 0 >>= qualifiedTarget
  applicative <- Array.index args 1 >>= qualifiedTarget
  guard (traversal.mbMod == Just (ModuleName "Data.TraversableWithIndex") && traversal.name == "traversableWithIndexArray")
  guard (applicative.mbMod == Just (ModuleName "Data.Either") && applicative.name == "applicativeEither")
  pure (emitKnownEither translate context nextId args)

emitKnownEither :: TranslateExpr -> ExprContext -> Int -> Array TcoExpr -> ExprResult
emitKnownEither translate context@{ metadata, codegenStateRef, modNameStr } nextId args =
  let
    -- Capture each expression before translating the next argument's statements.
    -- Dictionary getter evaluation and callback construction remain at the same
    -- application stage even when this call returns a closure.
    captured = foldl capture { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId } args
    argument index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argumentType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    callback = boxGoExpr codegenStateRef modNameStr (argument 2) (argumentType 2)
    suffix = show captured.nextId
    inputName = "traverseEither_input_" <> suffix
    valuesName = "traverseEither_values_" <> suffix
    outputName = "traverseEither_output_" <> suffix
    indexName = "traverseEither_index_" <> suffix
    valueName = "traverseEither_value_" <> suffix
    resultName = "traverseEither_result_" <> suffix
    errorName = "traverseEither_firstError_" <> suffix
    failedName = "traverseEither_failed_" <> suffix
    isPartial = Array.length args == 3
    input = if isPartial then GoVar inputName else argument 3
    inputType = if isPartial then TypeValue else argumentType 3
    source = case inputType of
      TypeNativeArray elementType -> { expr: input, elementType }
      _ ->
        { expr: rawArray input
        , elementType: TypeValue
        }
    eitherModule = Just (ModuleName "Data.Either")
    eitherType = ADT "Data.Either.Either" [ "Data", "Either", "Either" ] [ Any, Any ]
    result = { expr: GoVar resultName, exprType: TypeValue, stmts: StmtEmpty, nextId: captured.nextId }
    tag name = (AdtExprs.isTag metadata codegenStateRef modNameStr eitherModule name result).expr
    payload = AdtExprs.getField metadata codegenStateRef modNameStr
      { moduleName: eitherModule, ctorName: "Right", index: 0 }
      { expr: GoVar resultName, exprType: TypeValue, sourceType: eitherType }
    prepared = AdtExprs.prepareSaturated metadata modNameStr eitherModule "Right" eitherType
    arrayValue = boxGoExpr codegenStateRef modNameStr (GoVar outputName) (TypeNativeArray TypeValue)
    success = AdtExprs.saturated codegenStateRef prepared { exprs: [ arrayValue ], exprTypes: [ TypeValue ] }
    boxedSuccess = boxGoExpr codegenStateRef modNameStr success.expr success.exprType
    loop =
      [ GoAssign resultName (applyBoxed callback
          [ GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoVar indexName ] ]
          , boxGoExpr codegenStateRef modNameStr (GoVar valueName) source.elementType
          ])
      , GoIfElse (tag "Right")
          [ GoMutate (outputName <> "[" <> indexName <> "]") (boxGoExpr codegenStateRef modNameStr payload.expr payload.exprType) ]
          [ GoIfElse (tag "Left")
              [ GoIfElse (GoPrefixOp "!" (GoVar failedName))
                  [ GoMutate errorName (GoVar resultName), GoMutate failedName (rawGo "true") ] []
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
    expression = if isPartial then
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
      result = translate (CallArguments.childContext context Nothing) acc.nextId arg
      name = "traverseEither_arg_" <> show result.nextId
    in
      { stmts: acc.stmts <> result.stmts <> StmtLeaf (GoAssign name result.expr) <> StmtLeaf (GoMutate "_" (GoVar name))
      , exprs: Array.snoc acc.exprs (GoVar name)
      , exprTypes: Array.snoc acc.exprTypes result.exprType
      , nextId: result.nextId + 1
      }

  rawArray input = GoPrefixOp "*" (GoCall (rawGo "(*[]gopurs_runtime.Value)") [ GoSelector input "UnsafePtr" ])
