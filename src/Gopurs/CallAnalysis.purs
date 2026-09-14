module Gopurs.CallAnalysis
  ( CurriedAbs
  , collectCurriedAbs
  , extractUncurriedAbs
  , GoSpineArg
  , getGoSpineArgs
  , collectGoSpine
  , isClosureNode
  , CallTarget
  , qualifiedTarget
  , curriedTarget
  ) where

import Prelude
import Data.Array as Array
import Data.Array.NonEmpty (NonEmptyArray, toArray)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Tuple (Tuple(..))
import Gopurs.ExprContext (LocalEnv)
import Gopurs.GoAst (GoType(..))
import Gopurs.ExprAnalysis (getExprType, unwrapTcoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level)
import PureScript.Backend.Optimizer.Syntax as Syn

type CurriedAbs =
  { args :: NonEmptyArray (Tuple (Maybe Ident) Level)
  , body :: TcoExpr
  }

-- Collect adjacent curried lambdas without crossing a computation.
-- Typed is transparent only when looking for another Abs; the terminal body
-- keeps its original annotations for translation.
collectCurriedAbs :: NonEmptyArray (Tuple (Maybe Ident) Level) -> TcoExpr -> CurriedAbs
collectCurriedAbs args body =
  case lookAheadAbs body of
    Just next -> collectCurriedAbs (args <> next.args) next.body
    Nothing -> { args, body }
  where
  lookAheadAbs :: TcoExpr -> Maybe CurriedAbs
  lookAheadAbs (TcoExpr _ syntax) = case syntax of
    Typed _ inner -> lookAheadAbs inner
    Abs nextArgs nextBody -> Just { args: nextArgs, body: nextBody }
    _ -> Nothing

extractUncurriedAbs :: TcoExpr -> Maybe { args :: Array String, body :: TcoExpr }
extractUncurriedAbs (TcoExpr _ syntax) = case syntax of
  -- A zero-argument call is a separate evaluation step, even around another lambda.
  UncurriedAbs [] body -> Just { args: [], body }
  UncurriedAbs args body ->
    let
      thisArgs = map (\(Tuple mbI lvl) -> localId mbI lvl) args
    in
      case extractUncurriedAbs body of
        Just inner | not (Array.null inner.args) -> Just { args: thisArgs <> inner.args, body: inner.body }
        _ -> Just { args: thisArgs, body }
  Abs args body ->
    let
      thisArgs = map (\(Tuple mbI lvl) -> localId mbI lvl) (toArray args)
    in
      case extractUncurriedAbs body of
        Just inner | not (Array.null inner.args) -> Just { args: thisArgs <> inner.args, body: inner.body }
        _ -> Just { args: thisArgs, body }
  Typed _ inner -> extractUncurriedAbs inner
  _ -> Nothing

data GoSpineArg = GoSpineApp (Array TcoExpr) | GoSpineTypeApp ExprType

getGoSpineArgs :: Array GoSpineArg -> Array TcoExpr
getGoSpineArgs = Array.concatMap extractApp
  where
  extractApp (GoSpineApp a) = a
  extractApp _ = []

unwrapForSpine :: TcoExpr -> BackendSyntax TcoExpr
unwrapForSpine (TcoExpr _ syn) = case syn of
  Typed _ inner -> unwrapForSpine inner
  _ -> syn

collectGoSpine :: TcoExpr -> Tuple TcoExpr (Array GoSpineArg)
collectGoSpine e =
  case unwrapForSpine e of
    App f args ->
      let
        Tuple f' args' = collectGoSpine f
      in
        Tuple f' (args' <> [ GoSpineApp (toArray args) ])
    UncurriedApp f args ->
      let
        Tuple f' args' = collectGoSpine f
      in
        Tuple f' (args' <> [ GoSpineApp args ])
    Syn.TypeApp f ty ->
      let
        Tuple f' args' = collectGoSpine f
      in
        Tuple f' (args' <> [ GoSpineTypeApp ty ])
    _ -> Tuple e []

getArityFromType :: ExprType -> Int
getArityFromType = go 0
  where
  go acc (ForAll _ t) = go acc t
  go acc (ConstrainedType _ t) = go acc t
  go acc (Func args ret) = go (acc + Array.length args) ret
  go acc _ = acc

isClosureNode :: forall r. { globalTypes :: Map.Map String ExprType | r } -> TcoExpr -> Boolean
isClosureNode metadata expr = case unwrapTcoExpr expr of
  Abs _ _ -> true
  UncurriedAbs _ _ -> true
  App _ _ ->
    let
      Tuple flatFn flatArgsSpine = collectGoSpine expr
      flatArgs = getGoSpineArgs flatArgsSpine
      expectedArity = getArityFromType (getExprType flatFn)
      actualArity = Array.length flatArgs
    in
      case unwrapTcoExpr flatFn of
        Var (Qualified mbMn (Ident i)) ->
          let
            vType = case mbMn of
              Just mn -> Map.lookup (unwrap mn <> "." <> i) metadata.globalTypes
              Nothing -> Nothing

            expectedArity2 = case vType of
              Just t -> getArityFromType t
              Nothing -> 0
          in
            actualArity < expectedArity || actualArity < expectedArity2 || i == "foldrArray" || i == "foldlArray" || i == "traverse_" || i == "for_" || i == "traverseArrayImpl"
        _ -> actualArity < expectedArity
  UncurriedApp _ _ ->
    let
      Tuple flatFn flatArgsSpine = collectGoSpine expr
      flatArgs = getGoSpineArgs flatArgsSpine
      expectedArity = getArityFromType (getExprType flatFn)
      actualArity = Array.length flatArgs
    in
      case unwrapTcoExpr flatFn of
        Var (Qualified mbMn (Ident i)) ->
          let
            vType = case mbMn of
              Just mn -> Map.lookup (unwrap mn <> "." <> i) metadata.globalTypes
              Nothing -> Nothing

            expectedArity2 = case vType of
              Just t -> getArityFromType t
              Nothing -> 0
          in
            actualArity < expectedArity || actualArity < expectedArity2 || i == "foldrArray" || i == "foldlArray" || i == "traverse_" || i == "for_" || i == "traverseArrayImpl"
        _ -> actualArity < expectedArity
  Let _ _ _ body -> isClosureNode metadata body
  LetRec _ _ body -> isClosureNode metadata body
  Typed _ inner -> isClosureNode metadata inner
  _ -> false

type CallTarget = { mbMod :: Maybe ModuleName, name :: String }

qualifiedTarget :: TcoExpr -> Maybe CallTarget
qualifiedTarget expr = case unwrapTcoExpr expr of
  Var (Qualified mbMod (Ident name)) -> Just { mbMod, name }
  _ -> Nothing

curriedTarget :: LocalEnv -> TcoExpr -> Maybe CallTarget
curriedTarget bound expr = getVar (unwrapTcoExpr expr)
  where
  getVar :: BackendSyntax TcoExpr -> Maybe CallTarget
  getVar (Typed _ inner) = getVar (unwrapTcoExpr inner)
  getVar (Var (Qualified mbMod (Ident name))) = Just { mbMod, name }
  getVar (Local mbIdent lvl) =
    let
      resolvedName = (fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)).name
    in
      Just { mbMod: Nothing, name: resolvedName }
  getVar (Lit _) = Just { mbMod: Nothing, name: "Lit" }
  getVar (App _ _) = Just { mbMod: Nothing, name: "App" }
  getVar (Abs _ _) = Just { mbMod: Nothing, name: "Abs" }
  getVar (UncurriedApp _ _) = Just { mbMod: Nothing, name: "UncurriedApp" }
  getVar (UncurriedAbs _ _) = Just { mbMod: Nothing, name: "UncurriedAbs" }
  getVar (UncurriedEffectApp _ _) = Just { mbMod: Nothing, name: "UncurriedEffectApp" }
  getVar (UncurriedEffectAbs _ _) = Just { mbMod: Nothing, name: "UncurriedEffectAbs" }
  getVar (Accessor _ _) = Just { mbMod: Nothing, name: "Accessor" }
  getVar (Update _ _) = Just { mbMod: Nothing, name: "Update" }
  getVar (CtorSaturated _ _ _ _ _) = Just { mbMod: Nothing, name: "CtorSaturated" }
  getVar (CtorDef _ _ _ _) = Just { mbMod: Nothing, name: "CtorDef" }
  getVar (LetRec _ _ _) = Just { mbMod: Nothing, name: "LetRec" }
  getVar (Let _ _ _ _) = Just { mbMod: Nothing, name: "Let" }
  getVar (EffectBind _ _ _ _) = Just { mbMod: Nothing, name: "EffectBind" }
  getVar (EffectPure _) = Just { mbMod: Nothing, name: "EffectPure" }
  getVar (EffectDefer _) = Just { mbMod: Nothing, name: "EffectDefer" }
  getVar _ = Just { mbMod: Nothing, name: "Unknown" }
