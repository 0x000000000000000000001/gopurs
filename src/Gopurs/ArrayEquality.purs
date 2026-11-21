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
    captured = CallArguments.capture { prefix: "arrayEq_arg_", markUsed: true }
      (\index currentId value ->
        let result = translate (childContext context Nothing) currentId value
        in case index > 0, scalarType, result.expr of
          true, TypeInt64, GoUnboxIntArray source -> result { expr = source, exprType = TypeValue }
          _, _, _ -> result)
      nextId (Array.cons dictionary inputs)
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    suffix = show captured.nextId
    leftName = "arrayEq_left_" <> suffix
    rightName = "arrayEq_right_" <> suffix
    indexName = "arrayEq_index_" <> suffix
    left = arraySource "arrayEq_leftValue" leftName (argType 1)
    right = arraySource "arrayEq_rightValue" rightName (argType 2)
    element source = coerceGoExpr codegenStateRef modNameStr
      (GoIndex (rawGo ("(" <> source.target <> ")")) (GoVar indexName)) source.elementType scalarType
    length source = GoCall (GoVar "len") [ rawGo source.target ]
    expr = GoCall (GoFuncLit []
      [ GoAssign "arrayEq_leftValue" (arg 1)
      , GoAssign "arrayEq_rightValue" (arg 2)
      , left.assignment
      , right.assignment
      , GoIfElse (GoBinOp "!=" (length left) (length right)) [ GoReturn (rawGo "false") ] []
      , GoForRange (indexName <> " := range " <> left.target)
          [ GoIfElse (GoBinOp "!=" (element left) (element right)) [ GoReturn (rawGo "false") ] [] ]
      ] (rawGo "true") TypeBool) []
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
