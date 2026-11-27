module Gopurs.ArrayEquality (emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, childContext)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ModuleName(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(GetProp), BackendSyntax(Accessor))

-- Only canonical scalar Eq dictionaries justify native equality. An arbitrary
-- Eq instance, even for a newtype over the same scalar, retains its callback.
emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
emit translate context@{ codegenStateRef, modNameStr } nextId fn args = do
  { dictionary, inputs } <- case unwrapTcoExpr fn, args of
    Accessor dictionary (GetProp "eq"), [ left, right ] ->
      Just { dictionary, inputs: [ left, right ] }
    _, [ dictionary, left, right ] | isEqName context "eq" fn ->
      Just { dictionary, inputs: [ left, right ] }
    _, _ -> Nothing
  let Tuple arrayEq spine = collectGoSpine dictionary
  guard (isEqName context "eqArray" arrayEq)
  scalar <- case getGoSpineArgs spine of
    [ scalar ] -> Just scalar
    _ -> Nothing
  scalarType <- scalarEquality context scalar
  let
    -- Canonical scalar Eq dictionaries are pure library instances whose runtime
    -- representation is unused by the native equality loop. Capture only the
    -- two array operands in source order.
    captured = CallArguments.capture { prefix: "arrayEq_arg_", markUsed: true }
      (\_ currentId value ->
        let result = translate (childContext context Nothing) currentId value
        in case scalarType, result.expr of
          TypeInt64, GoUnboxIntArray source -> result { expr = source, exprType = TypeValue }
          _, _ -> result)
      nextId inputs
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    suffix = show captured.nextId
    leftName = "arrayEq_left_" <> suffix
    rightName = "arrayEq_right_" <> suffix
    indexName = "arrayEq_index_" <> suffix
    left = arraySource "arrayEq_leftValue" leftName (argType 0)
    right = arraySource "arrayEq_rightValue" rightName (argType 1)
    element source = coerceGoExpr codegenStateRef modNameStr
      (GoIndex (rawGo ("(" <> source.target <> ")")) (GoVar indexName)) source.elementType scalarType
    length source = GoCall (GoVar "len") [ rawGo source.target ]
    bothNative = left.elementType == scalarType && right.elementType == scalarType
    slicesEqual = GoRaw { text: "slices.Equal", imports: [ "slices" ] }
    equalIntSlices = GoSelector (GoVar "gopurs_runtime") "EqualIntArray"
    nativeEqual = if scalarType == TypeInt64 then equalIntSlices else slicesEqual
    intArrayFastPath = if scalarType == TypeInt64 && not bothNative then
      case left.elementType, right.elementType of
        TypeInt64, TypeValue ->
          [ GoIfElse (rawGo "arrayEq_rightValue.Type == gopurs_runtime.TypeIntArray")
              [ GoReturn (GoCall equalIntSlices [ GoVar "arrayEq_leftValue", rawGo "*(*[]int64)(arrayEq_rightValue.UnsafePtr)" ]) ] [] ]
        TypeValue, TypeInt64 ->
          [ GoIfElse (rawGo "arrayEq_leftValue.Type == gopurs_runtime.TypeIntArray")
              [ GoReturn (GoCall equalIntSlices [ rawGo "*(*[]int64)(arrayEq_leftValue.UnsafePtr)", GoVar "arrayEq_rightValue" ]) ] [] ]
        TypeValue, TypeValue ->
          [ GoIfElse (rawGo "arrayEq_leftValue.Type == gopurs_runtime.TypeIntArray")
              [ GoIfElse (rawGo "arrayEq_rightValue.Type == gopurs_runtime.TypeIntArray")
                  [ GoReturn (GoCall equalIntSlices [ rawGo "*(*[]int64)(arrayEq_leftValue.UnsafePtr)", rawGo "*(*[]int64)(arrayEq_rightValue.UnsafePtr)" ]) ]
                  [ GoAssign "arrayEq_li" (rawGo "*(*[]int64)(arrayEq_leftValue.UnsafePtr)")
                  , GoAssign "arrayEq_rv" (rawGo "*(*[]gopurs_runtime.Value)(arrayEq_rightValue.UnsafePtr)")
                  , GoIfElse (rawGo "len(arrayEq_li) != len(arrayEq_rv)") [ GoReturn (rawGo "false") ] []
                  , GoForRange "i, v := range arrayEq_li"
                      [ GoIfElse (rawGo "v != arrayEq_rv[i].IntVal") [ GoReturn (rawGo "false") ] [] ]
                  , GoReturn (rawGo "true")
                  ]
              ]
              [ GoIfElse (rawGo "arrayEq_rightValue.Type == gopurs_runtime.TypeIntArray")
                  [ GoAssign "arrayEq_lv" (rawGo "*(*[]gopurs_runtime.Value)(arrayEq_leftValue.UnsafePtr)")
                  , GoAssign "arrayEq_ri" (rawGo "*(*[]int64)(arrayEq_rightValue.UnsafePtr)")
                  , GoIfElse (rawGo "len(arrayEq_lv) != len(arrayEq_ri)") [ GoReturn (rawGo "false") ] []
                  , GoForRange "i, v := range arrayEq_ri"
                      [ GoIfElse (rawGo "arrayEq_lv[i].IntVal != v") [ GoReturn (rawGo "false") ] [] ]
                  , GoReturn (rawGo "true")
                  ] []
              ]
          ]
        _, _ -> []
      else []
    compareBody = if bothNative then
      [ left.assignment
      , right.assignment
      , GoReturn (GoCall nativeEqual [ rawGo left.target, rawGo right.target ])
      ]
      else
      intArrayFastPath <>
      [ left.assignment
      , right.assignment
      , GoIfElse (GoBinOp "!=" (length left) (length right)) [ GoReturn (rawGo "false") ] []
      , GoForRange (indexName <> " := range " <> left.target)
          [ GoIfElse (GoBinOp "!=" (element left) (element right)) [ GoReturn (rawGo "false") ] [] ]
      ]
    expr = GoCall (GoFuncLit []
      ([ GoAssign "arrayEq_leftValue" (arg 0)
       , GoAssign "arrayEq_rightValue" (arg 1)
       ] <> compareBody) (rawGo "true") TypeBool) []
  pure { stmts: captured.stmts, expr, exprType: TypeBool, nextId: captured.nextId }

isEqName :: ExprContext -> String -> TcoExpr -> Boolean
isEqName { modNameStr } name expression = case qualifiedTarget expression of
  Just target -> target.name == name && case target.mbMod of
    Just (ModuleName "Data.Eq") -> true
    Nothing -> modNameStr == "Data_Eq" || modNameStr == "Data.Eq"
    _ -> false
  Nothing -> false

scalarEquality :: ExprContext -> TcoExpr -> Maybe GoType
scalarEquality context scalar
  | isEqName context "eqInt" scalar = Just TypeInt64
  | isEqName context "eqNumber" scalar = Just TypeFloat64
  | isEqName context "eqBoolean" scalar = Just TypeBool
  | isEqName context "eqChar" scalar || isEqName context "eqString" scalar = Just TypeString
  | otherwise = Nothing
