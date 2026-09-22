module Gopurs.ImmediateApplications (optimizeImmediateApplications) where

import Prelude

import Control.Monad.State (State, evalState, get, put)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (class Foldable, foldl)
import Data.Map as Map
import Data.Maybe (Maybe, fromMaybe)
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), snd)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (Ident, Literal(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level(..), Pair(..))

-- Run after PBO, before ownership and TCO recompute local usage. This is
-- ordinary capture-avoiding beta reduction, including an immediately applied
-- branch result. It never substitutes an evaluated call/effect as an argument
-- or duplicates a parameter's uses. Recursive initialization scopes are left
-- untouched. Small syntax bounds limit branch-distribution code growth.
optimizeImmediateApplications :: BackendModule -> BackendModule
optimizeImmediateApplications mod = mod
  { bindings = map (\group -> group
      { bindings = map (\(Tuple name expr) ->
          Tuple name (evalState (rewrite expr) (maximumLevel expr + 1))) group.bindings
      }) mod.bindings
  }

rewrite :: NeutralExpr -> State Int NeutralExpr
rewrite original@(NeutralExpr syn) = case syn of
  LetRec _ _ _ -> pure original
  _ -> do
    children <- traverse rewrite syn
    case children of
      App fn args -> case NEA.toArray args of
        [ arg ] | movable arg && size arg <= 128 && reducible fn -> applyInto fn arg
        _ -> pure (NeutralExpr children)
      _ -> pure (NeutralExpr children)

-- Do not distribute applications of opaque functions merely because their
-- producer is a branch. Every possible result must eliminate the application.
reducible :: NeutralExpr -> Boolean
reducible expr = case strip expr of
  Abs refs body -> case NEA.toArray refs of
    [ Tuple _ level ] -> not (hasRecursion body) && occurrences level body <= 1
    _ -> false
  Let _ _ _ body -> reducible body
  Branch cases fallback -> NEA.length cases <= 4
    && foldl (\accepted (Pair _ body) -> accepted && reducible body) true cases
    && reducible fallback
  Fail _ -> true
  _ -> false

applyInto :: NeutralExpr -> NeutralExpr -> State Int NeutralExpr
applyInto original arg = case strip original of
  Abs refs body -> case NEA.toArray refs of
    [ Tuple _ level ] | not (hasRecursion body) && occurrences level body <= 1 -> do
      -- An argument lambda may reuse levels from a sibling branch. Rename its
      -- bound locals before transplanting it below that branch's local binds.
      renamed <- freshen Map.empty arg
      rewrite (substitute level renamed body)
    _ -> unchanged
  Let ident level value body -> do
    applied <- applyInto body arg
    pure (NeutralExpr (Let ident level value applied))
  Branch cases fallback | NEA.length cases <= 4 -> do
    cases' <- traverse (\(Pair condition body) -> Pair condition <$> applyInto body arg) cases
    fallback' <- applyInto fallback arg
    pure (NeutralExpr (Branch cases' fallback'))
  Fail message -> pure (NeutralExpr (Fail message))
  _ -> unchanged
  where
  unchanged = pure (NeutralExpr (App original (NEA.singleton arg)))

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syn) = case syn of
  Typed _ inner -> strip inner
  TypeApp inner _ -> strip inner
  _ -> syn

-- Reading a non-recursive local or scalar has no evaluation to move. A lambda
-- is already a value: its effectful/diverging body stays deferred until called.
-- Global getters, accessors, constructors, calls and arrays are not admitted.
movable :: NeutralExpr -> Boolean
movable expr = case strip expr of
  Local _ _ -> true
  Abs _ body -> not (hasRecursion body)
  Lit (LitInt _) -> true
  Lit (LitNumber _) -> true
  Lit (LitString _) -> true
  Lit (LitChar _) -> true
  Lit (LitBoolean _) -> true
  _ -> false

hasRecursion :: NeutralExpr -> Boolean
hasRecursion (NeutralExpr syn) = case syn of
  LetRec _ _ _ -> true
  _ -> foldl (\found child -> found || hasRecursion child) false syn

size :: NeutralExpr -> Int
size (NeutralExpr syn) = 1 + foldl (\n child -> n + size child) 0 syn

maximumLevel :: NeutralExpr -> Int
maximumLevel (NeutralExpr syn) = foldl (\n child -> max n (maximumLevel child)) own syn
  where
  level (Level n) = n
  refs :: forall f. Foldable f => f (Tuple (Maybe Ident) Level) -> Int
  refs xs = foldl (\n ref -> max n (level (snd ref))) (-1) xs
  own = case syn of
    Local _ l -> level l
    Abs xs _ -> refs xs
    UncurriedAbs xs _ -> refs xs
    UncurriedEffectAbs xs _ -> refs xs
    Let _ l _ _ -> level l
    EffectBind _ l _ _ -> level l
    LetRec l xs _ -> level l + NEA.length xs - 1
    _ -> -1

occurrences :: Level -> NeutralExpr -> Int
occurrences target (NeutralExpr syn) = case syn of
  Local _ level -> if level == target then 1 else 0
  Abs refs body -> if bound refs then 0 else occurrences target body
  UncurriedAbs refs body -> if bound refs then 0 else occurrences target body
  UncurriedEffectAbs refs body -> if bound refs then 0 else occurrences target body
  Let _ level value body -> occurrences target value + if level == target then 0 else occurrences target body
  EffectBind _ level value body -> occurrences target value + if level == target then 0 else occurrences target body
  _ -> foldl (\n child -> n + occurrences target child) 0 syn
  where
  bound :: forall f. Foldable f => f (Tuple (Maybe Ident) Level) -> Boolean
  bound = foldl (\found ref -> found || snd ref == target) false

substitute :: Level -> NeutralExpr -> NeutralExpr -> NeutralExpr
substitute target replacement original@(NeutralExpr syn) = case syn of
  Local _ level | level == target -> replacement
  Abs refs _ | bound refs -> original
  UncurriedAbs refs _ | bound refs -> original
  UncurriedEffectAbs refs _ | bound refs -> original
  Let ident level value body -> NeutralExpr
    (Let ident level (go value) (if level == target then body else go body))
  EffectBind ident level value body -> NeutralExpr
    (EffectBind ident level (go value) (if level == target then body else go body))
  _ -> NeutralExpr (map go syn)
  where
  go = substitute target replacement
  bound :: forall f. Foldable f => f (Tuple (Maybe Ident) Level) -> Boolean
  bound = foldl (\found ref -> found || snd ref == target) false

fresh :: State Int Level
fresh = do
  next <- get
  put (next + 1)
  pure (Level next)

freshen :: Map.Map Level Level -> NeutralExpr -> State Int NeutralExpr
freshen env (NeutralExpr syn) = NeutralExpr <$> case syn of
  Local ident level -> pure (Local ident (fromMaybe level (Map.lookup level env)))
  Abs refs body -> do
    renamed <- traverse renameRef refs
    let env' = foldl addRef env (NEA.zip refs renamed)
    Abs renamed <$> freshen env' body
  UncurriedAbs refs body -> do
    renamed <- traverse renameRef refs
    let env' = foldl addRef env (Array.zip refs renamed)
    UncurriedAbs renamed <$> freshen env' body
  UncurriedEffectAbs refs body -> do
    renamed <- traverse renameRef refs
    let env' = foldl addRef env (Array.zip refs renamed)
    UncurriedEffectAbs renamed <$> freshen env' body
  Let ident level value body -> do
    level' <- fresh
    Let ident level' <$> freshen env value <*> freshen (Map.insert level level' env) body
  EffectBind ident level value body -> do
    level' <- fresh
    EffectBind ident level' <$> freshen env value <*> freshen (Map.insert level level' env) body
  _ -> traverse (freshen env) syn
  where
  renameRef :: Tuple (Maybe Ident) Level -> State Int (Tuple (Maybe Ident) Level)
  renameRef (Tuple ident _) = Tuple ident <$> fresh
  addRef acc (Tuple (Tuple _ old) (Tuple _ new)) = Map.insert old new acc
