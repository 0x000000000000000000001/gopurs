module Gopurs.ClosedDictionaries.Scope (closedConstruction) where

import Prelude

import Data.Foldable (class Foldable, foldl)
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level(..))

-- This proof concerns the entire proposed getter, including deferred bodies.
-- Effects and recursive initialization cannot move into a shared construction.
closedConstruction :: NeutralExpr -> Boolean
closedConstruction = closedIn Set.empty

-- Carry the lexical environment down instead of allocating and merging every
-- subtree's free-variable set. Effect/recursive nodes reject the same proof
-- immediately, including when nested inside a deferred function body.
closedIn :: Set.Set String -> NeutralExpr -> Boolean
closedIn bound (NeutralExpr syn) = case syn of
  Local mbIdent lvl -> Set.member (localId mbIdent lvl) bound
  Abs args body -> closedIn (bindArgs bound args) body
  UncurriedAbs args body -> closedIn (bindArgs bound args) body
  Let mbIdent lvl val body -> closedIn bound val
    && closedIn (Set.insert (localId mbIdent lvl) bound) body
  LetRec _ _ _ -> false
  PrimEffect _ -> false
  EffectBind _ _ _ _ -> false
  EffectPure _ -> false
  EffectDefer _ -> false
  UncurriedEffectApp _ _ -> false
  UncurriedEffectAbs _ _ -> false
  _ -> foldl (\valid child -> valid && closedIn bound child) true syn

bindArgs :: forall f. Foldable f => Set.Set String -> f (Tuple (Maybe Ident) Level) -> Set.Set String
bindArgs = foldl (\bound (Tuple mbIdent lvl) -> Set.insert (localId mbIdent lvl) bound)

-- This proof uses the original source spelling and level. Go-sanitized names
-- can collide and must not identify two different source references here.
localId :: Maybe Ident -> Level -> String
localId (Just (Ident i)) (Level l) = i <> "_" <> show l
localId Nothing (Level l) = "__local_var_" <> show l
