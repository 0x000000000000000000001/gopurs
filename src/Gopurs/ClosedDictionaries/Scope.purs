module Gopurs.ClosedDictionaries.Scope (closedConstruction) where

import Prelude

import Data.Array.NonEmpty (toArray)
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
closedConstruction expr = Set.isEmpty (freeVars expr)
  && not (containsEffect expr)
  && not (containsRecursion expr)

containsEffect :: NeutralExpr -> Boolean
containsEffect (NeutralExpr syn) = case syn of
  PrimEffect _ -> true
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> true
  UncurriedEffectApp _ _ -> true
  UncurriedEffectAbs _ _ -> true
  _ -> foldl (\found child -> found || containsEffect child) false syn

containsRecursion :: NeutralExpr -> Boolean
containsRecursion (NeutralExpr syn) = case syn of
  LetRec _ _ _ -> true
  _ -> foldl (\found child -> found || containsRecursion child) false syn

-- Only binding forms change scope; every other form unions its children's free
-- references. A let binds its body, never its own initializer.
freeVars :: NeutralExpr -> Set.Set String
freeVars (NeutralExpr syn) = case syn of
  Local mbIdent lvl -> Set.singleton (localId mbIdent lvl)
  Abs args body -> differenceBound args body
  UncurriedAbs args body -> differenceBound args body
  UncurriedEffectAbs args body -> differenceBound args body
  LetRec lvl binds body ->
    let
      bindsSet = foldl (\acc (Tuple ident _) -> Set.insert (localId (Just ident) lvl) acc) Set.empty (toArray binds)
      bodyVars = freeVars body
      bindsVars = foldl (\acc (Tuple _ e) -> Set.union acc (freeVars e)) Set.empty (toArray binds)
    in
      Set.difference (Set.union bodyVars bindsVars) bindsSet
  Let mbIdent lvl val body ->
    Set.union (freeVars val) (Set.difference (freeVars body) (Set.singleton (localId mbIdent lvl)))
  EffectBind mbIdent lvl val body ->
    Set.union (freeVars val) (Set.difference (freeVars body) (Set.singleton (localId mbIdent lvl)))
  _ -> foldl (\acc child -> Set.union acc (freeVars child)) Set.empty syn

differenceBound :: forall f. Foldable f => f (Tuple (Maybe Ident) Level) -> NeutralExpr -> Set.Set String
differenceBound args body =
  let
    argsSet = foldl (\acc (Tuple mbIdent lvl) -> Set.insert (localId mbIdent lvl) acc) Set.empty args
  in
    Set.difference (freeVars body) argsSet

-- This proof uses the original source spelling and level. Go-sanitized names
-- can collide and must not identify two different source references here.
localId :: Maybe Ident -> Level -> String
localId (Just (Ident i)) (Level l) = i <> "_" <> show l
localId Nothing (Level l) = "__local_var_" <> show l
