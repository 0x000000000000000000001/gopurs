module Gopurs.ConstructorExprs
  ( definition
  , saturated
  ) where

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.AdtExprs as AdtExprs
import Gopurs.ExprAnalysis (bindFieldFunctionParameters, extractExprFuncType, getExprType, hasTypeVars, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, LocalEnv, StmtTree(..), TranslateExpr, childContext)
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Literal(..), ModuleName)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(CtorDef, CtorSaturated, Lit))

definition :: ExprContext -> Int -> TcoExpr -> String -> Array String -> ExprResult
definition { metadata, codegenStateRef, modNameStr, mbExpectedExprType } nextId node name fields =
  let
    annotatedType = fromMaybe (getExprType node) mbExpectedExprType
    ctorType = case extractExprFuncType annotatedType of
      Just { fRet } -> fRet
      Nothing -> annotatedType
    result = AdtExprs.definition metadata codegenStateRef modNameStr name fields ctorType
  in
    { stmts: StmtEmpty, expr: result.expr, exprType: result.exprType, nextId }

-- Prepare the result layout once. Translate each field in source order and
-- coerce it immediately, before the next field can inspect the shared state.
saturated :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Maybe ModuleName -> String -> Array (Tuple String TcoExpr) -> ExprResult
saturated translateExpr context@{ metadata, codegenStateRef, modNameStr, bound, mbExpectedExprType } nextId node moduleName name props =
  let
    ctorType = saturatedResultType (getExprType node) mbExpectedExprType
    prepared = AdtExprs.prepareSaturated metadata modNameStr moduleName name ctorType
    child = childContext context Nothing
    toGoType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr
    fields = Array.foldl
      (\acc (Tuple _ value) ->
        let
          expected = AdtExprs.saturatedFieldType prepared acc.fieldIdx
          fieldBound = bindFieldFunctionParameters toGoType bound expected.exprType value
          valueResult = translateExpr (child { bound = fieldBound }) acc.nextId value
          coerced = coerceGoExpr codegenStateRef modNameStr valueResult.expr valueResult.exprType expected.goType
          constant = isReusableConstant expected.goType valueResult.exprType value
        in
          { stmts: acc.stmts <> valueResult.stmts
          , exprs: Array.snoc acc.exprs coerced
          , exprTypes: Array.snoc acc.exprTypes expected.goType
          , nextId: valueResult.nextId
          , fieldIdx: acc.fieldIdx + 1
          , constants: Array.snoc acc.constants constant
          })
      { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId, fieldIdx: 0, constants: [] }
      props
    result = AdtExprs.saturated codegenStateRef prepared { exprs: fields.exprs, exprTypes: fields.exprTypes }
  in
    reuseConstructor bound fields.constants
      { stmts: fields.stmts, expr: result.expr, exprType: result.exprType, nextId: fields.nextId }

-- Prefer a concrete expected type only when the node has no type or still has
-- variables. An already concrete annotation keeps priority.
saturatedResultType :: ExprType -> Maybe ExprType -> ExprType
saturatedResultType actual expected = case actual of
  Any -> fromMaybe Any expected
  ty | hasTypeVars ty -> case expected of
    Just expectedType | not (hasTypeVars expectedType) -> expectedType
    _ -> ty
  ty -> ty

-- Reuse requires a representation-preserving scalar replacement. Number is
-- deliberately absent: numeric equality cannot distinguish signed zero.
isReusableConstant :: GoType -> GoType -> TcoExpr -> Boolean
isReusableConstant expected actual value = expected == actual && case expected, unwrapTcoExpr value of
  TypeInt64, Lit (LitInt _) -> true
  TypeBool, Lit (LitBoolean _) -> true
  TypeUint32, CtorSaturated _ _ _ _ fields -> Array.null fields
  TypeUint32, CtorDef _ _ _ fields -> Array.null fields
  _, _ -> false

-- Reuse is considered only after field translation and only without statements.
-- The conditional owns one new local name; all other paths keep the counter.
reuseConstructor :: LocalEnv -> Array Boolean -> ExprResult -> ExprResult
reuseConstructor bound constants result =
  let
    reuse = case result.stmts of
      StmtEmpty -> AdtExprs.constructorReuse bound result.exprType constants result.expr
      _ -> Nothing
  in
    case reuse of
      Just { source, condition } ->
        let
          resultName = "__reuse_" <> show result.nextId
          declare = rawGo ("var " <> resultName <> " " <> goTypeToStr result.exprType)
          choose = GoIfElse condition [ GoMutate resultName source ] [ GoMutate resultName result.expr ]
        in
          { stmts: StmtLeaf declare <> StmtLeaf choose, expr: GoVar resultName, exprType: result.exprType, nextId: result.nextId + 1 }
      Nothing -> result
