module Gopurs.ArrayIndexedFold (emit, readOnlyScalar, integerCallback) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.CallAnalysis (collectCurriedAbs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (extractFuncType, getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, bindParameters, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Literal(..), ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendSyntax(..), Level, Pair(..))

-- The generic FFI adapter copies the source array. Reading it directly is safe
-- only if the callback cannot mutate it through an alias. This proof admits
-- primitive scalar work, not arbitrary calls merely carrying a pure type.
-- All primitive arithmetic keeps its existing semantics. The fold invokes each
-- callback in strict ascending order; no totality assumption is needed.
readOnlyScalar :: Array Level -> TcoExpr -> Boolean
readOnlyScalar locals expression = scalarType (Map.fromFoldable (map (\level -> Tuple level Int) locals)) expression == Just Int

scalarType :: Map.Map Level ExprType -> TcoExpr -> Maybe ExprType
scalarType locals expression = do
  inferred <- case unwrapTcoExpr expression of
    Local _ level -> case Map.lookup level locals of
      Just ty -> Just ty
      Nothing -> let ty = getExprType expression in if ty == Int || ty == Boolean then Just ty else Nothing
    Lit (LitInt _) -> Just Int
    Lit (LitBoolean _) -> Just Boolean
    PrimOp (Op1 operator value) -> do
      ty <- if operator == OpBooleanNot then Just Boolean
        else if operator == OpIntNegate || operator == OpIntBitNot then Just Int else Nothing
      guard (scalarType locals value == Just ty)
      pure ty
    PrimOp (Op2 operator left right) -> do
      signature <- scalarOperator operator
      guard (scalarType locals left == Just signature.operand && scalarType locals right == Just signature.operand)
      pure signature.result
    Branch branches otherwise -> do
      ty <- scalarType locals otherwise
      guard (all (\(Pair condition result) -> scalarType locals condition == Just Boolean && scalarType locals result == Just ty) branches)
      pure ty
    Let _ level value body -> do
      ty <- scalarType locals value
      scalarType (Map.insert level ty locals) body
    _ -> Nothing
  let annotated = getExprType expression
  guard (annotated == Any || annotated == inferred)
  pure inferred

scalarOperator :: BackendOperator2 -> Maybe { operand :: ExprType, result :: ExprType }
scalarOperator = case _ of
  OpIntNum _ -> integer
  OpIntOrd _ -> Just { operand: Int, result: Boolean }
  OpBooleanOrd _ -> boolean
  OpBooleanAnd -> boolean
  OpBooleanOr -> boolean
  OpIntBitAnd -> integer
  OpIntBitOr -> integer
  OpIntBitXor -> integer
  OpIntBitShiftLeft -> integer
  OpIntBitShiftRight -> integer
  OpIntBitZeroFillShiftRight -> integer
  _ -> Nothing
  where
  integer = Just { operand: Int, result: Int }
  boolean = Just { operand: Boolean, result: Boolean }

-- PBO may leave the first lambda unannotated while retaining the exact TAST
-- type of its suffix. The canonical indexed-fold contract proves that first
-- parameter is Int; the suffix must still prove both other parameters/result.
-- An explicit incompatible outer function type is never overridden.
integerCallback :: TcoExpr -> Boolean
integerCallback callback = case extractFuncType callback of
  Just signature -> signature.fArgs == [ Int, Int, Int ] && signature.fRet == Int
  Nothing -> case callback of
    TcoExpr _ (Typed _ _) -> false
    _ -> case unwrapTcoExpr callback of
      Abs parameters body | NEA.length parameters == 1 -> case extractFuncType body of
        Just signature -> signature.fArgs == [ Int, Int ] && signature.fRet == Int
        Nothing -> false
      _ -> false

emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
emit translate context nextId fn args = do
  target <- qualifiedTarget fn
  guard (target.mbMod == Just (ModuleName "Data.FoldableWithIndex") && target.name == "foldlWithIndexImpl")
  guard (Array.length args == 2 || Array.length args == 3)
  guard (Array.null context.recVars)
  callback <- Array.head args
  guard (integerCallback callback)
  case unwrapTcoExpr callback of
    Abs parameters result -> do
      let grouped = collectCurriedAbs parameters result
      guard (NEA.length grouped.args == 3)
      guard (readOnlyScalar (map (\(Tuple _ level) -> level) (NEA.toArray grouped.args)) grouped.body)
      let
        params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) (NEA.toArray grouped.args)
        callbackContext = context
          { depth = context.depth + 1
          , bound = bindParameters params context.bound
          , tcoIdent = Nothing
          , loopCtx = []
          , options = { isTail: true, inEffectBlock: false }
          , mbExpectedExprType = Just Int
          , mbExpectedGoType = Just TypeInt64
          }
        translateArgument index currentId value = if index == 0 then
          let translatedBody = translate callbackContext currentId grouped.body
          in { stmts: StmtEmpty
             , expr: GoFuncLit params (flattenStmts translatedBody.stmts) (integer translatedBody.expr translatedBody.exprType) TypeInt64
             , exprType: TypeFunc [ TypeInt64, TypeInt64, TypeInt64 ] TypeInt64
             , nextId: translatedBody.nextId
             }
          else translate ((childContext context Nothing) { mbExpectedGoType = if index == 2 then Just TypeValue else Nothing }) currentId value
        -- Captures preserve callback/seed/input order. With two arguments the
        -- first two are evaluated now; each later invocation gets a fresh acc.
        captured = CallArguments.capture { prefix: "indexedFold_arg_", markUsed: true } translateArgument nextId args
        argument index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
        argumentType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
        suffix = show captured.nextId
        inputName = "indexedFold_input_" <> suffix
        valuesName = "indexedFold_values_" <> suffix
        resultName = "indexedFold_result_" <> suffix
        indexName = "indexedFold_index_" <> suffix
        valueName = "indexedFold_value_" <> suffix
        complete = Array.length args == 3
        source = arraySource inputName valuesName (if complete then argumentType 2 else TypeValue)
        body =
          (if complete then [ GoAssign inputName (argument 2) ] else []) <>
          [ source.assignment
          , GoAssign resultName (integer (argument 1) (argumentType 1))
          , GoForRange (indexName <> ", " <> valueName <> " := range " <> source.target)
              [ GoMutate resultName (GoCall (argument 0)
                  [ GoCall (GoVar "int64") [ GoVar indexName ], GoVar resultName, integer (GoVar valueName) source.elementType ]) ]
          ]
        resultExpr = GoVar resultName
        expression = if complete then GoCall (GoFuncLit [] body resultExpr TypeInt64) []
          else GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
            [ GoFuncLit [ Tuple inputName TypeValue ] body
                (boxGoExpr context.codegenStateRef context.modNameStr resultExpr TypeInt64) TypeValue ]
      pure { stmts: captured.stmts, expr: expression, exprType: if complete then TypeInt64 else TypeValue, nextId: captured.nextId + 1 }
    _ -> Nothing
  where
  integer expression from = coerceGoExpr context.codegenStateRef context.modNameStr expression from TypeInt64
