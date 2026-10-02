module Gopurs.Monomorphization.ForeignForwarders (collectForeignForwarders) where

import Prelude

import Data.Array as Array
import Data.Foldable (foldl)
import Data.List (List)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.CoreFn (Ann, Bind(..), Binding(..), Expr(..), ExprType, Ident(..), Module(..), Qualified(..))

-- These wrappers cannot unbox across their value-level FFI boundary. Recognize
-- the exact eta-forwarding body, including permitted unsafeCoerce wrappers;
-- the orchestration retains that body for collection but excludes its clones.
collectForeignForwarders :: List (Module Ann) -> Set String
collectForeignForwarders = foldl addModuleForwarders Set.empty

addModuleForwarders :: Set String -> Module Ann -> Set String
addModuleForwarders acc (Module mod) =
  let
    moduleName = unwrap mod.name
    foreignIdents = Set.fromFoldable
      (map (\(Tuple (Ident name) _) -> name) (Map.toUnfoldable mod.foreign :: Array (Tuple Ident (Maybe ExprType))))
  in
    Array.foldl (addForwardingBind moduleName foreignIdents) acc mod.decls

addForwardingBind :: String -> Set String -> Set String -> Bind Ann -> Set String
addForwardingBind moduleName foreignIdents acc = case _ of
  NonRec binding -> addForwarder moduleName foreignIdents acc binding
  Rec group -> Array.foldl (addForwarder moduleName foreignIdents) acc group

addForwarder :: String -> Set String -> Set String -> Binding Ann -> Set String
addForwarder moduleName foreignIdents acc (Binding _ (Ident ident) body) =
  if isForeignForwarder moduleName foreignIdents body then
    Set.insert (moduleName <> "." <> ident) acc
  else acc

isForeignForwarder :: String -> Set String -> Expr Ann -> Boolean
isForeignForwarder moduleName foreignIdents body = case collectLambdaParams body of
  { params, body: inner } ->
    not (Array.null params)
      && Set.size (Set.fromFoldable params) == Array.length params
      && case unapplyExpr inner of
      { head: ExprVar _ (Qualified mbModule (Ident ffi)), args } ->
        Set.member ffi foreignIdents
          && sameModule mbModule
          && (mbModule /= Nothing || not (Array.elem (Ident ffi) params))
          && Array.length args == Array.length params
          && foldl (&&) true (Array.zipWith isParameterReference args params)
      _ -> false
  where
  sameModule = case _ of
    Just mn -> unwrap mn == moduleName
    Nothing -> true

collectLambdaParams :: Expr Ann -> { params :: Array Ident, body :: Expr Ann }
collectLambdaParams (ExprAbs _ param body) =
  case collectLambdaParams body of
    rest -> rest { params = Array.cons param rest.params }
collectLambdaParams body = { params: [], body }

unapplyExpr :: Expr Ann -> { head :: Expr Ann, args :: Array (Expr Ann) }
unapplyExpr expr = case stripCoercions expr of
  ExprApp _ fn arg ->
    case unapplyExpr fn of
      inner -> inner { args = Array.snoc inner.args arg }
  other -> { head: other, args: [] }

stripCoercions :: Expr Ann -> Expr Ann
stripCoercions expr = case expr of
  ExprTypeApp _ inner _ -> stripCoercions inner
  ExprApp _ fn arg | isUnsafeCoerce fn -> stripCoercions arg
  _ -> expr

-- `unsafeCoerce` is polymorphic, so its use is wrapped in type applications.
isUnsafeCoerce :: Expr Ann -> Boolean
isUnsafeCoerce fn = case stripTypeApps fn of
  ExprVar _ (Qualified (Just moduleName) (Ident "unsafeCoerce")) -> unwrap moduleName == "Unsafe.Coerce"
  _ -> false

stripTypeApps :: Expr Ann -> Expr Ann
stripTypeApps (ExprTypeApp _ inner _) = stripTypeApps inner
stripTypeApps expr = expr

isParameterReference :: Expr Ann -> Ident -> Boolean
isParameterReference expr (Ident name) = case stripCoercions expr of
  ExprVar _ (Qualified Nothing (Ident used)) -> used == name
  _ -> false
