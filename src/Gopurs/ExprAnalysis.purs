module Gopurs.ExprAnalysis
  ( isEffectNode
  , unwrapTcoExpr
  , printTcoExprShape
  , extractExprFuncType
  , extractFuncType
  , getExprType
  , executeIfOpaque
  , bindFieldFunctionParameters
  , hasTypeVars
  ) where

import Prelude
import Data.Array as Array
import Data.Array.NonEmpty (toArray)
import Data.Foldable (foldl, any)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.Tuple (Tuple(..))
import Gopurs.ExprContext (LocalEnv)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType)
import Gopurs.GoTypes (printExprType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Literal(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendSyntax(UncurriedAbs, Abs, PrimOp, Typed, EffectPure, EffectBind, UncurriedEffectApp, UncurriedApp, UncurriedEffectAbs, Lit, LetRec, Var, Branch, App, Let, PrimEffect, EffectDefer), Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

isEffectNode :: TcoExpr -> Boolean
isEffectNode expr = case unwrapTcoExpr expr of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> false
  PrimEffect _ -> true
  UncurriedEffectApp _ _ -> true
  Let _ _ _ body -> isEffectNode body
  LetRec _ _ body -> isEffectNode body
  _ -> false

unwrapTcoExpr :: TcoExpr -> BackendSyntax TcoExpr
unwrapTcoExpr (TcoExpr _ syn) = case syn of
  Typed _ inner -> unwrapTcoExpr inner
  Syn.TypeApp inner _ -> unwrapTcoExpr inner
  _ -> syn

printTcoExprShape :: TcoExpr -> String
printTcoExprShape e = case unwrapTcoExpr e of
  Let _ _ _ body -> "Let(" <> printTcoExprShape body <> ")"
  Abs _ body -> "Abs(" <> printTcoExprShape body <> ")"
  App fn _ -> "App(" <> printTcoExprShape fn <> ")"
  Branch branches def -> "Branch(" <> String.joinWith ", " (map (\(Pair _ expr) -> printTcoExprShape expr) (toArray branches)) <> ", def=" <> printTcoExprShape def <> ")"
  Var _ -> "Var"
  LetRec _ _ body -> "LetRec(" <> printTcoExprShape body <> ")"
  Lit (LitInt _) -> "LitInt"
  Lit (LitNumber _) -> "LitNumber"
  Lit (LitString _) -> "LitString"
  Lit (LitChar _) -> "LitChar"
  Lit (LitBoolean _) -> "LitBoolean"
  Lit (LitArray _) -> "LitArray"
  Lit (LitRecord _) -> "LitRecord"
  UncurriedAbs _ body -> "UncurriedAbs(" <> printTcoExprShape body <> ")"
  UncurriedEffectAbs _ body -> "UncurriedEffectAbs(" <> printTcoExprShape body <> ")"
  UncurriedApp fn _ -> "UncurriedApp(" <> printTcoExprShape fn <> ")"
  UncurriedEffectApp fn _ -> "UncurriedEffectApp(" <> printTcoExprShape fn <> ")"
  EffectBind _ _ _ body -> "EffectBind(" <> printTcoExprShape body <> ")"
  EffectPure _ -> "EffectPure"
  Typed tp inner -> "Typed(" <> printExprType tp <> ", " <> printTcoExprShape inner <> ")"
  _ -> "Other"

extractExprFuncType :: ExprType -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }
extractExprFuncType ty =
  let
    flattenFuncType acc (Func args ret) = flattenFuncType (acc <> args) ret
    flattenFuncType acc ret = { fArgs: acc, fRet: ret }

    getFunc (Func a r) = Just (flattenFuncType a r)
    getFunc (ConstrainedType constraints innerTy) =
      case getFunc innerTy of
        Just i ->
          let
            constraintTypes = map
              ( \(Tuple qual args) ->
                  let
                    qualStr = String.joinWith "." qual
                  in
                    ADT qualStr qual args
              )
              constraints
          in
            Just (i { fArgs = constraintTypes <> i.fArgs })
        Nothing -> Nothing
    getFunc (ForAll _ innerTy) = getFunc innerTy
    getFunc _ = Nothing
  in
    getFunc ty

extractFuncType :: TcoExpr -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }
extractFuncType (TcoExpr _ (Typed ty inner)) =
  case extractExprFuncType ty of
    Just r -> Just r
    Nothing -> extractFuncType inner
extractFuncType _ = Nothing

getExprType :: TcoExpr -> ExprType
getExprType (TcoExpr _ syn) = case syn of
  Typed t _ -> t
  PrimOp op -> case op of
    Op1 OpIntNegate _ -> Int
    Op1 OpIntBitNot _ -> Int
    Op1 OpNumberNegate _ -> Number
    Op1 OpBooleanNot _ -> Boolean
    Op1 (OpIsTag _) _ -> Boolean
    Op1 OpArrayLength _ -> Int
    Op2 (OpIntNum _) _ _ -> Int
    Op2 (OpIntOrd _) _ _ -> Boolean
    Op2 OpIntBitZeroFillShiftRight _ _ -> Int
    Op2 (OpNumberNum _) _ _ -> Number
    Op2 (OpNumberOrd _) _ _ -> Boolean
    Op2 OpStringAppend _ _ -> String
    Op2 (OpStringOrd _) _ _ -> Boolean
    Op2 (OpCharOrd _) _ _ -> Boolean
    Op2 (OpBooleanOrd _) _ _ -> Boolean
    Op2 OpBooleanAnd _ _ -> Boolean
    Op2 OpBooleanOr _ _ -> Boolean
    _ -> Any
  _ -> Any

executeIfOpaque :: TcoExpr -> GoExpr -> GoExpr

executeIfOpaque expr goExpr =
  if isEffectNode expr then goExpr
  else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ goExpr, rawGo "gopurs_runtime.Value{}" ]

bindFieldFunctionParameters :: (ExprType -> GoType) -> LocalEnv -> ExprType -> TcoExpr -> LocalEnv
bindFieldFunctionParameters toGoType bound expectedExprType value =
  let
    mbArgs = case unwrapTcoExpr value of
      Abs args _ -> Just (toArray args)
      UncurriedAbs args _ -> Just args
      _ -> Nothing
  in
    case mbArgs, extractExprFuncType expectedExprType of
      Just args, Just { fArgs } ->
        let
          paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (toGoType fArgTy)) args (fArgs <> Array.replicate (Array.length args - Array.length fArgs) Any)
        in
          foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
      _, _ -> bound

hasTypeVars :: ExprType -> Boolean
hasTypeVars = case _ of
  TypeVar _ -> true
  Array t -> hasTypeVars t
  ADT _ _ ts -> any hasTypeVars ts
  TypeApp t ts -> hasTypeVars t || any hasTypeVars ts
  Func ts t -> any hasTypeVars ts || hasTypeVars t
  Record t -> hasTypeVars t
  Row ts tail ->
    any (\(Tuple _ t) -> hasTypeVars t) ts ||
      case tail of
        Just t -> hasTypeVars t
        Nothing -> false
  ForAll _ t -> hasTypeVars t
  ConstrainedType cs t -> any (\(Tuple _ ts) -> any hasTypeVars ts) cs || hasTypeVars t
  _ -> false
