module Gopurs.ImmediateApplications.Scope
  ( hasRecursion
  , maximumLevel
  , occurrences
  , substitute
  , freshen
  ) where

import Prelude
import Control.Monad.State (State, get, put)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (class Foldable, foldl)
import Data.Map as Map
import Data.Maybe (Maybe, fromMaybe)
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), snd)
import PureScript.Backend.Optimizer.CoreFn (Ident)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level(..))

hasRecursion :: NeutralExpr -> Boolean
hasRecursion (NeutralExpr syn) = case syn of
  LetRec _ _ _ -> true
  _ -> foldl (\found child -> found || hasRecursion child) false syn

-- The supply starts above every level in the original binding, including scopes
-- that the rewriting pass will leave untouched. Each module binding owns a supply.
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

-- Counting, substitution and freshening operate on admitted, recursion-free
-- fragments. Parameter shadowing stops at bodies, not Let/EffectBind initializers.
occurrences :: Level -> NeutralExpr -> Int
occurrences target (NeutralExpr syn) = case syn of
  Local _ level -> if level == target then 1 else 0
  Abs refs body -> if binds target refs then 0 else occurrences target body
  UncurriedAbs refs body -> if binds target refs then 0 else occurrences target body
  UncurriedEffectAbs refs body -> if binds target refs then 0 else occurrences target body
  Let _ level value body -> occurrences target value + if level == target then 0 else occurrences target body
  EffectBind _ level value body -> occurrences target value + if level == target then 0 else occurrences target body
  _ -> foldl (\n child -> n + occurrences target child) 0 syn

substitute :: Level -> NeutralExpr -> NeutralExpr -> NeutralExpr
substitute target replacement original@(NeutralExpr syn) = case syn of
  Local _ level | level == target -> replacement
  Abs refs _ | binds target refs -> original
  UncurriedAbs refs _ | binds target refs -> original
  UncurriedEffectAbs refs _ | binds target refs -> original
  Let ident level value body -> NeutralExpr
    (Let ident level (go value) (if level == target then body else go body))
  EffectBind ident level value body -> NeutralExpr
    (EffectBind ident level (go value) (if level == target then body else go body))
  _ -> NeutralExpr (map go syn)
  where
  go = substitute target replacement

binds :: forall f. Foldable f => Level -> f (Tuple (Maybe Ident) Level) -> Boolean
binds target = foldl (\found ref -> found || snd ref == target) false

fresh :: State Int Level
fresh = do
  next <- get
  put (next + 1)
  pure (Level next)

-- Rename bound locals only; free references still capture the same outer values.
-- One traversal per insertion preserves the pass's deterministic level order.
freshen :: NeutralExpr -> State Int NeutralExpr
freshen = renameWith Map.empty

renameWith :: Map.Map Level Level -> NeutralExpr -> State Int NeutralExpr
renameWith env (NeutralExpr syn) = NeutralExpr <$> case syn of
  Local ident level -> pure (Local ident (fromMaybe level (Map.lookup level env)))
  Abs refs body -> do
    renamed <- traverse renameRef refs
    let env' = foldl addRef env (NEA.zip refs renamed)
    Abs renamed <$> renameWith env' body
  UncurriedAbs refs body -> do
    renamed <- traverse renameRef refs
    let env' = foldl addRef env (Array.zip refs renamed)
    UncurriedAbs renamed <$> renameWith env' body
  UncurriedEffectAbs refs body -> do
    renamed <- traverse renameRef refs
    let env' = foldl addRef env (Array.zip refs renamed)
    UncurriedEffectAbs renamed <$> renameWith env' body
  Let ident level value body -> do
    level' <- fresh
    Let ident level' <$> renameWith env value <*> renameWith (Map.insert level level' env) body
  EffectBind ident level value body -> do
    level' <- fresh
    EffectBind ident level' <$> renameWith env value <*> renameWith (Map.insert level level' env) body
  _ -> traverse (renameWith env) syn
  where
  renameRef :: Tuple (Maybe Ident) Level -> State Int (Tuple (Maybe Ident) Level)
  renameRef (Tuple ident _) = Tuple ident <$> fresh
  addRef acc (Tuple (Tuple _ old) (Tuple _ new)) = Map.insert old new acc
