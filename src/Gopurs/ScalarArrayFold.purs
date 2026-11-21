module Gopurs.ScalarArrayFold (emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.CallAnalysis (qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, childContext)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ModuleName(..))

-- Only these exact builtins justify eliminating the callback and reading the
-- existing array representation. User dictionaries/functions keep their calls.
emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
emit translate context@{ codegenStateRef, modNameStr } nextId fn args = do
  guard (named "Data.Foldable" "foldlArray" fn && Array.length args == 3)
  callback <- Array.head args
  scalar <- if named "Data.Semiring" "intAdd" callback then Just TypeInt64
    else if named "Data.Semigroup" "concatString" callback then Just TypeString
    else Nothing
  let
    -- A read-only scalar fold needs no native array copy. Coerce each element
    -- at its use instead; array annotations still determine other call paths.
    captured = CallArguments.capture { prefix: "scalarFold_arg_", markUsed: true }
      (\index currentId value -> translate
        ((childContext context Nothing) { mbExpectedGoType = if index == 2 then Just TypeValue else Nothing })
        currentId value)
      nextId args
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    source = arraySource "scalarFold_input" "scalarFold_values" (argType 2)
    initial = coerceGoExpr codegenStateRef modNameStr (arg 1) (argType 1) scalar
    element = coerceGoExpr codegenStateRef modNameStr (GoVar "scalarFold_value") source.elementType scalar
    common = [ GoAssign "scalarFold_input" (arg 2), source.assignment ]
    expr = if scalar == TypeInt64 then
      GoCall (GoFuncLit []
        (common <>
          [ GoAssign "scalarFold_result" initial
          , GoForRange ("_, scalarFold_value := range " <> source.target)
              [ GoMutate "scalarFold_result" (GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd")
                  [ GoVar "scalarFold_result", element ]) ]
          ]) (GoVar "scalarFold_result") scalar) []
      else
        GoCall (GoFuncLit [] common
          (GoCall (GoSelector (GoVar "gopurs_runtime") "ConcatStrings")
            [ initial, rawGo source.target
            , GoFuncLit [ Tuple "scalarFold_value" source.elementType ] [] element TypeString
            ]) TypeString) []
  pure { stmts: captured.stmts, expr, exprType: scalar, nextId: captured.nextId }

named :: String -> String -> TcoExpr -> Boolean
named moduleName name expression = case qualifiedTarget expression of
  Just target -> target.mbMod == Just (ModuleName moduleName) && target.name == name
  Nothing -> false
