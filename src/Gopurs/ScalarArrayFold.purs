module Gopurs.ScalarArrayFold (emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Foldable (foldl)
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
    intBody = case argType 2 of
      TypeNativeArray _ ->
        common <> [ GoAssign "scalarFold_result" initial ] <>
          integerSumBody source.target element
            (\value -> coerceGoExpr codegenStateRef modNameStr value source.elementType TypeInt64)
      _ ->
        [ GoAssign "scalarFold_input" (arg 2)
        , GoAssign "scalarFold_result" initial
        , GoIfElse (rawGo "scalarFold_input.Type == gopurs_runtime.TypeInt32Array")
            [ GoAssign "scalarFold_ints32" (rawGo "*(*[]int32)(scalarFold_input.UnsafePtr)")
            , GoIfElse (GoBinOp ">" (GoCall (GoVar "len") [ GoVar "scalarFold_ints32" ]) (GoInt 0))
                [ GoMutate "scalarFold_result"
                    (GoCall (GoVar "int64")
                      [ GoBinOp "+"
                          (GoCall (GoVar "int32") [ GoVar "scalarFold_result" ])
                          (GoCall (GoSelector (GoVar "gopurs_runtime") "SumInt32") [ GoVar "scalarFold_ints32" ])
                      ])
                ] []
            ]
            [ GoMutate "scalarFold_result"
                (GoCall (GoSelector (GoVar "gopurs_runtime") "SumIntArraySlow")
                  [ GoVar "scalarFold_input", GoVar "scalarFold_result" ])
            ]
        ]
    expr = if scalar == TypeInt64 then
      GoCall (GoFuncLit [] intBody (GoVar "scalarFold_result") scalar) []
      else
        GoCall (GoFuncLit [] common
          (GoCall (GoSelector (GoVar "gopurs_runtime") "ConcatStrings")
            [ initial, rawGo source.target
            , GoFuncLit [ Tuple "scalarFold_value" source.elementType ] [] element TypeString
            ]) TypeString) []
  pure { stmts: captured.stmts, expr, exprType: scalar, nextId: captured.nextId }

-- IntAdd is addition modulo 2^32. Independent accumulators may therefore use
-- wrapping int64 additions and narrow once after a nonempty fold. Small inputs
-- keep the original loop, including returning a wide seed unchanged when empty.
-- This applies only to the canonical callback selected above, not user folds.
integerSumBody :: String -> GoExpr -> (GoExpr -> GoExpr) -> Array GoExpr
integerSumBody source element unbox =
  [ GoAssign values (rawGo source)
  , GoIfElse (GoBinOp "<" (GoCall (GoVar "len") [ GoVar values ]) (GoInt 8))
      [ GoForRange ("_, scalarFold_value := range " <> values)
          [ GoMutate "scalarFold_result" (GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd")
              [ GoVar "scalarFold_result", element ]) ]
      , GoReturn (GoVar "scalarFold_result")
      ] []
  ] <>
  map (\lane -> GoAssign (accumulator lane) (GoCall (GoVar "int64") [ GoInt 0 ])) lanes <>
  [ GoAssign index (GoInt 0)
  , GoForRange ("; " <> index <> " <= len(" <> values <> ")-8; " <> index <> " += 8")
      ([ GoAssign chunk (rawGo (values <> "[" <> index <> ":" <> index <> "+8]")) ] <>
        map (\lane -> GoMutate (accumulator lane)
          (GoBinOp "+" (GoVar (accumulator lane)) (unbox (GoIndex (GoVar chunk) (GoInt lane))))) lanes)
  , GoForRange ("; " <> index <> " < len(" <> values <> "); " <> index <> "++")
      [ GoMutate (accumulator 0)
          (GoBinOp "+" (GoVar (accumulator 0)) (unbox (GoIndex (GoVar values) (GoVar index)))) ]
  , GoMutate "scalarFold_result" (GoCall (GoVar "int64")
      [ GoCall (GoVar "int32") [ foldl (GoBinOp "+") (GoVar "scalarFold_result") (map (GoVar <<< accumulator) lanes) ] ])
  ]
  where
  lanes = Array.range 0 7
  values = "scalarFold_items"
  index = "scalarFold_index"
  chunk = "scalarFold_chunk"
  accumulator lane = "scalarFold_acc" <> show lane

named :: String -> String -> TcoExpr -> Boolean
named moduleName name expression = case qualifiedTarget expression of
  Just target -> target.mbMod == Just (ModuleName moduleName) && target.name == name
  Nothing -> false
