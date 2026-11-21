module Gopurs.ArraySort (emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.AdtExprs as AdtExprs
import Gopurs.ArrayIntrinsics.Callback as Callback
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments (applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, childContext)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(GetProp), BackendSyntax(Accessor))

-- The standard Int comparator is pure and fully known. Preserve the separately
-- supplied Ordering conversion, including all of its calls; custom Ord or
-- comparison functions keep their ordinary FFI path. The sorting algorithm and
-- strict-decreasing fast path mirror the library's stable sort implementation.
emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
emit translate context@{ metadata, codegenStateRef, modNameStr } nextId fn args = do
  guard (named "Data.Array" "sortByImpl" fn && Array.length args == 3)
  comparator <- Array.head args
  guard (intComparator comparator)
  let
    captured = CallArguments.capture { prefix: "arraySort_arg_", markUsed: true }
      (\index currentId expression -> case if index == 1 then Callback.nativeLambda translate context currentId 1 expression else Nothing of
        Just native -> native
        Nothing ->
          let result = translate (childContext context Nothing) currentId expression
          in case index, result.expr of
            2, GoUnboxIntArray boxedSource -> result { expr = boxedSource, exprType = TypeValue }
            _, _ -> result)
      nextId args
    argument index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argumentType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    source = arraySource "arraySort_input" "arraySort_values" (argumentType 2)
    orderingType = ADT "Data.Ordering.Ordering" [ "Data", "Ordering", "Ordering" ] []
    converted name =
      let
        prepared = AdtExprs.prepareSaturated metadata modNameStr (Just (ModuleName "Data.Ordering")) name orderingType
        ordering = AdtExprs.saturated codegenStateRef prepared { exprs: [], exprTypes: [] }
      in case argumentType 1 of
        TypeFunc [ expected ] resultType -> coerceGoExpr codegenStateRef modNameStr
          (GoCall (argument 1) [ coerceGoExpr codegenStateRef modNameStr ordering.expr ordering.exprType expected ]) resultType TypeInt64
        _ -> coerceGoExpr codegenStateRef modNameStr
          (applyBoxed (boxGoExpr codegenStateRef modNameStr (argument 1) (argumentType 1))
            [ boxGoExpr codegenStateRef modNameStr ordering.expr ordering.exprType ]) TypeValue TypeInt64
    compare = GoFuncBlock [ Tuple "a" TypeInt64, Tuple "b" TypeInt64 ]
      [ GoIfElse (GoBinOp "<" (GoVar "a") (GoVar "b")) [ GoReturn (converted "LT") ] []
      , GoIfElse (GoBinOp "==" (GoVar "a") (GoVar "b")) [ GoReturn (converted "EQ") ] []
      , GoReturn (converted "GT")
      ] TypeInt64
    length = GoCall (GoVar "len") [ GoVar "arraySort_output" ]
    at index = GoIndex (GoVar "arraySort_output") index
    comparison a b = GoCall (GoVar "arraySort_compare") [ a, b ]
    body =
      [ GoAssign "arraySort_input" (argument 2), source.assignment
      , GoAssign "arraySort_output" (GoCall (GoVar "make") [ rawGo "[]int64", GoCall (GoVar "len") [ rawGo source.target ] ])
      , GoForRange ("index, value := range " <> source.target)
          [ GoMutate "arraySort_output[index]" (coerceGoExpr codegenStateRef modNameStr (GoVar "value") source.elementType TypeInt64) ]
      , GoIfElse (GoBinOp "<" length (GoInt 2)) [ GoReturn (GoVar "arraySort_output") ] []
      , GoAssign "arraySort_compare" compare
      , GoAssign "arraySort_descending" (GoBinOp "<" (comparison (at (GoInt 1)) (at (GoInt 0))) (GoInt 0))
      , GoAssign "arraySort_monotone" (rawGo "true")
      , GoForRange "index := 2; index < len(arraySort_output); index++"
          [ GoAssign "order" (comparison (at (GoVar "index")) (at (GoBinOp "-" (GoVar "index") (GoInt 1))))
          , GoIfElse (GoBinOp "||"
              (GoBinOp "&&" (GoVar "arraySort_descending") (GoBinOp ">=" (GoVar "order") (GoInt 0)))
              (GoBinOp "&&" (GoPrefixOp "!" (GoVar "arraySort_descending")) (GoBinOp "<" (GoVar "order") (GoInt 0))))
              [ GoMutate "arraySort_monotone" (rawGo "false"), rawGo "break" ] []
          ]
      , GoIfElse (GoVar "arraySort_monotone")
          [ GoIfElse (GoVar "arraySort_descending") [ GoCall (slices "Reverse") [ GoVar "arraySort_output" ] ] []
          , GoReturn (GoVar "arraySort_output") ] []
      , GoCall (slices "SortStableFunc")
          [ GoVar "arraySort_output"
          , GoFuncLit [ Tuple "a" TypeInt64, Tuple "b" TypeInt64 ] []
              (GoCall (GoVar "int") [ comparison (GoVar "a") (GoVar "b") ]) (TypeInterface "int")
          ]
      ]
  pure { stmts: captured.stmts
    , expr: GoCall (GoFuncLit [] body (GoVar "arraySort_output") (TypeNativeArray TypeInt64)) []
    , exprType: TypeNativeArray TypeInt64, nextId: captured.nextId }
  where
  slices name = GoRaw { text: "slices." <> name, imports: [ "slices" ] }

intComparator :: TcoExpr -> Boolean
intComparator expression = case unwrapTcoExpr expression of
  Accessor dictionary (GetProp "compare") -> named "Data.Ord" "ordInt" dictionary
  _ ->
    let Tuple fn spine = collectGoSpine expression
    in named "Data.Ord" "compare" fn && case getGoSpineArgs spine of
      [ dictionary ] -> named "Data.Ord" "ordInt" dictionary
      _ -> false

named :: String -> String -> TcoExpr -> Boolean
named moduleName name expression = case qualifiedTarget expression of
  Just target -> target.mbMod == Just (ModuleName moduleName) && target.name == name
  Nothing -> false
