module Gopurs.ObjectTraverse (emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.AdtExprs as AdtExprs
import Gopurs.CallAnalysis (CallTarget, qualifiedTarget)
import Gopurs.CallArguments (Arguments, applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName(..))

-- The standard Object traversal inserts into a copied immutable map for every
-- successful field. With the known Either applicative, only the final map can
-- escape, so a fresh map local to this invocation can collect those fields.
-- Unknown dictionaries retain the generic traversal and its applicative effects.
-- traverseWithIndexDefault first maps the object through its separate FFI;
-- its callback order is not the sorted fold order, so it must stay generic.
emit :: TranslateExpr -> ExprContext -> Int -> Maybe CallTarget -> Array TcoExpr -> Maybe ExprResult
emit translate context nextId target args = do
  recognize target args
  let captured = captureArguments translate context nextId args
  pure (emitLoop context captured)

recognize :: Maybe CallTarget -> Array TcoExpr -> Maybe Unit
recognize target args = do
  { mbMod, name } <- target
  guard (mbMod == Just (ModuleName "Data.TraversableWithIndex"))
  guard (name == "traverseWithIndex")
  guard (Array.length args >= 2 && Array.length args <= 4)
  traversal <- Array.index args 0 >>= qualifiedTarget
  applicative <- Array.index args 1 >>= qualifiedTarget
  guard (traversal.mbMod == Just (ModuleName "Foreign.Object") && traversal.name == "traversableWithIndexObject")
  guard (applicative.mbMod == Just (ModuleName "Data.Either") && applicative.name == "applicativeEither")

-- Box and capture each supplied argument before translating the next one.
-- A dictionary-only application neither evaluates an absent callback nor
-- allocates output storage: both remain inside the binary worker below.
captureArguments :: TranslateExpr -> ExprContext -> Int -> Array TcoExpr -> Arguments
captureArguments translate context@{ codegenStateRef, modNameStr } =
  CallArguments.capture { prefix: "traverseObjectEither_arg_", markUsed: true } boxArgument
  where
  boxArgument _ nextId arg =
    let result = translate (childContext context Nothing) nextId arg
    in result
      { expr = boxGoExpr codegenStateRef modNameStr result.expr result.exprType
      , exprType = TypeValue
      }

emitLoop :: ExprContext -> Arguments -> ExprResult
emitLoop { metadata, codegenStateRef, modNameStr } captured =
  let
    argument index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    suffix = show captured.nextId
    callbackName = "traverseObjectEither_callback_" <> suffix
    inputName = "traverseObjectEither_input_" <> suffix
    valuesName = "traverseObjectEither_values_" <> suffix
    keysName = "traverseObjectEither_keys_" <> suffix
    keyName = "traverseObjectEither_key_" <> suffix
    outputName = "traverseObjectEither_output_" <> suffix
    resultName = "traverseObjectEither_result_" <> suffix
    errorName = "traverseObjectEither_firstError_" <> suffix
    failedName = "traverseObjectEither_failed_" <> suffix
    supplied = Array.length captured.exprs
    callback = if supplied == 2 then GoVar callbackName else argument 2
    input = if supplied < 4 then GoVar inputName else argument 3
    runtime name expressions = GoCall (GoSelector (GoVar "gopurs_runtime") name) expressions
    eitherModule = Just (ModuleName "Data.Either")
    eitherType = ADT "Data.Either.Either" [ "Data", "Either", "Either" ] [ Any, Any ]
    result = { expr: GoVar resultName, exprType: TypeValue, stmts: StmtEmpty, nextId: captured.nextId }
    tag name = (AdtExprs.isTag metadata codegenStateRef modNameStr eitherModule name result).expr
    payload = AdtExprs.getField metadata codegenStateRef modNameStr
      { moduleName: eitherModule, ctorName: "Right", index: 0 }
      { expr: GoVar resultName, exprType: TypeValue, sourceType: eitherType }
    prepared = AdtExprs.prepareSaturated metadata modNameStr eitherModule "Right" eitherType
    success = AdtExprs.saturated codegenStateRef prepared
      { exprs: [ runtime "Any" [ GoVar outputName ] ], exprTypes: [ TypeValue ] }
    boxedSuccess = boxGoExpr codegenStateRef modNameStr success.expr success.exprType
    loop =
      [ GoAssign resultName (applyBoxed callback
          [ runtime "Str" [ GoVar keyName ]
          , runtime "Box" [ GoIndex (GoVar valuesName) (GoVar keyName) ]
          ])
      , GoIfElse (tag "Right")
          [ GoMutate (outputName <> "[" <> keyName <> "]") (boxGoExpr codegenStateRef modNameStr payload.expr payload.exprType) ]
          [ GoIfElse (tag "Left")
              [ GoIfElse (GoPrefixOp "!" (GoVar failedName))
                  [ GoMutate errorName (GoVar resultName)
                  , GoMutate failedName (rawGo "true")
                  ] []
              ]
              [ GoCall (GoVar "panic") [ GoString "Failed pattern match" ] ]
          ]
      ]
    body =
      [ GoAssign valuesName (runtime "UnboxObject" [ input ])
      , GoAssign keysName (GoCall (GoVar "make") [ rawGo "[]string", GoInt 0, GoCall (GoVar "len") [ GoVar valuesName ] ])
      , GoForRange (keyName <> " := range " <> valuesName)
          [ GoMutate keysName (GoCall (GoVar "append") [ GoVar keysName, GoVar keyName ]) ]
      -- Foreign.Object._foldM uses sort.Strings in this Go backend. Preserve
      -- that exact order, including numeric-looking keys and non-ASCII keys;
      -- changing it could change the first error or observable FFI callbacks.
      , GoCall (GoRaw { text: "sort.Strings", imports: [ "sort" ] }) [ GoVar keysName ]
      , GoAssign outputName (GoCall (GoVar "make") [ rawGo "map[string]any", GoCall (GoVar "len") [ GoVar valuesName ] ])
      , GoAssign errorName (rawGo "gopurs_runtime.Value{}")
      , GoAssign failedName (rawGo "false")
      -- The original strict fold evaluates every callback, even after Left.
      -- Keep its first error while continuing the remaining callbacks. The
      -- output never aliases the input or a previous invocation's result.
      , GoForRange ("_, " <> keyName <> " := range " <> keysName) loop
      , GoIfElse (GoVar failedName) [ GoReturn (GoVar errorName) ] []
      ]
    expression = case supplied of
      2 -> runtime "Func2"
        [ GoFuncLit [ Tuple callbackName TypeValue, Tuple inputName TypeValue ] body boxedSuccess TypeValue ]
      3 -> runtime "Func"
        [ GoFuncLit [ Tuple inputName TypeValue ] body boxedSuccess TypeValue ]
      _ -> GoCall (GoFuncLit [] body boxedSuccess TypeValue) []
  in
    { stmts: captured.stmts, expr: expression, exprType: TypeValue, nextId: captured.nextId + 1 }
