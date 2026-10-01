module Gopurs.ImmediateApplications (optimizeImmediateApplications) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (State, evalState)
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldl)
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.ImmediateApplications.Scope (freshen, hasRecursion, maximumLevel, occurrences, substitute)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (Ident, Literal(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level, Pair(..))

-- Admission records every place where the application will disappear. Emission
-- follows this plan without repeating guards or speculatively distributing calls.
data Reduction
  = SubstituteParameter Level Int NeutralExpr
  | PreserveLet (Maybe Ident) Level NeutralExpr Reduction
  | DistributeBranch (NEA.NonEmptyArray (Tuple NeutralExpr Reduction)) Reduction
  | PropagateFailure String

type ApplicationPlan = { reduction :: Reduction, copies :: Int }

-- Run after PBO, before ownership and TCO recompute local usage. This is
-- ordinary capture-avoiding beta reduction, including an immediately applied
-- branch result. It never substitutes an evaluated call/effect as an argument
-- or duplicates a callback value. Already-evaluated local/scalar arguments may
-- replace multiple parameter uses. Recursive initialization scopes are left
-- untouched. Bound duplicated syntax, not the size of a callback that moves
-- into just one live branch: a single insertion causes no callback code growth.
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
        [ arg ] | movable arg -> case admitApplication arg fn of
          Just plan | boundedGrowth plan.copies arg -> applyInto plan.reduction arg
          _ -> pure (NeutralExpr children)
        _ -> pure (NeutralExpr children)
      _ -> pure (NeutralExpr children)

-- Do not distribute applications of opaque functions merely because their
-- producer is a branch. Every possible result must eliminate the application.
admitApplication :: NeutralExpr -> NeutralExpr -> Maybe ApplicationPlan
admitApplication arg expr = case strip expr of
  Abs refs body -> case NEA.toArray refs of
    [ Tuple _ level ] -> do
      guard (not (hasRecursion body))
      let copies = occurrences level body
      guard (copies <= 1 || repeatable arg)
      pure { reduction: SubstituteParameter level copies body, copies }
    _ -> Nothing
  Let ident level value body -> do
    plan <- admitApplication arg body
    pure (plan { reduction = PreserveLet ident level value plan.reduction })
  Branch cases fallback -> do
    guard (NEA.length cases <= 4)
    plans <- traverse (\(Pair condition body) -> do
      plan <- admitApplication arg body
      pure { condition, plan }) cases
    fallbackPlan <- admitApplication arg fallback
    pure
      { reduction: DistributeBranch (map (\entry -> Tuple entry.condition entry.plan.reduction) plans) fallbackPlan.reduction
      , copies: foldl (\count entry -> count + entry.plan.copies) fallbackPlan.copies plans
      }
  Fail message -> pure { reduction: PropagateFailure message, copies: 0 }
  _ -> Nothing

-- Called only for a reducible function. Sum syntactic substitutions across
-- branches, rather than taking their maximum: mutually exclusive branches
-- still duplicate generated code. Zero/one substitution can admit a large body.
boundedGrowth :: Int -> NeutralExpr -> Boolean
boundedGrowth copies arg = copies <= 1 || size arg <= 128 / (copies - 1)

applyInto :: Reduction -> NeutralExpr -> State Int NeutralExpr
applyInto plan arg = case plan of
  SubstituteParameter level copies body ->
    -- An argument lambda may reuse levels from a sibling branch. Rename its
    -- bound locals before transplanting it below that branch's local binds.
    if copies == 0 then pure body
    else do
      renamed <- freshen arg
      rewrite (substitute level renamed body)
  PreserveLet ident level value body -> do
    applied <- applyInto body arg
    pure (NeutralExpr (Let ident level value applied))
  DistributeBranch cases fallback -> do
    cases' <- traverse (\(Tuple condition body) -> Pair condition <$> applyInto body arg) cases
    fallback' <- applyInto fallback arg
    pure (NeutralExpr (Branch cases' fallback'))
  PropagateFailure message -> pure (NeutralExpr (Fail message))

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syn) = case syn of
  Typed _ inner -> strip inner
  TypeApp inner _ -> strip inner
  _ -> syn

-- Reading a non-recursive local or scalar has no evaluation to move. A lambda
-- is already a value: its effectful/diverging body stays deferred until called.
-- Global getters, accessors, constructors, calls and arrays are not admitted.
movable :: NeutralExpr -> Boolean
movable expr = repeatable expr || case strip expr of
  Abs _ body -> not (hasRecursion body)
  _ -> false

-- A local reference reuses exactly the already-created value, including a
-- function or object. It never repeats evaluation or creates another closure.
-- A lambda is deliberately excluded even when its body is pure.
repeatable :: NeutralExpr -> Boolean
repeatable expr = case strip expr of
  Local _ _ -> true
  Lit (LitInt _) -> true
  Lit (LitNumber _) -> true
  Lit (LitString _) -> true
  Lit (LitChar _) -> true
  Lit (LitBoolean _) -> true
  _ -> false

size :: NeutralExpr -> Int
size (NeutralExpr syn) = 1 + foldl (\n child -> n + size child) 0 syn
