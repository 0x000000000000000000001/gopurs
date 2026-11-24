-- Evaluate only completely known, pure integer expressions. Dictionary methods
-- are ordinary immutable records and closures here; no class or benchmark name
-- participates in the proof. Unknown syntax leaves the source expression intact.
module Gopurs.ClosedScalar (optimize) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (StateT, evalStateT, get, put)
import Control.Monad.Trans.Class (lift)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(Int), Ident, Literal(..), ModuleName, Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator2(..), BackendOperatorNum(..), BackendSyntax(..), Level)

type Ref = Tuple (Maybe Ident) Level
type Env = Map.Map Ref Value
type Context =
  { moduleName :: ModuleName
  , globals :: Map.Map Ident NeutralExpr
  , active :: Set.Set Ident
  , depth :: Int
  }

data Value
  = Integer Int
  | Object (Map.Map String Value)
  | Closure (Array Ref) NeutralExpr Env (Set.Set Ident)
  | Constructed (Array Value)

type Eval = StateT Int Maybe

optimize :: BackendModule -> BackendModule
optimize mod =
  let globals = Map.fromFoldable $ Array.filter
        (\(Tuple name _) -> not (Map.member name mod.foreign))
        (Array.concatMap _.bindings (Array.filter (not <<< _.recursive) mod.bindings))
      context = { moduleName: mod.name, globals, active: Set.empty, depth: 0 }
  in mod { bindings = map (\group -> group { bindings = map
    (\(Tuple name value) -> Tuple name (rewrite context value)) group.bindings }) mod.bindings }

rewrite :: Context -> NeutralExpr -> NeutralExpr
rewrite context original@(NeutralExpr syntax) = case evaluateCandidate context original of
  Just value -> NeutralExpr (Typed Int (NeutralExpr (Lit (LitInt value))))
  Nothing -> NeutralExpr (map (rewrite context) syntax)

-- Keep the output domain explicit in the TAST. A closure/record/constructor is
-- only an intermediate interpreter value, never a newly emitted shared object.
evaluateCandidate :: Context -> NeutralExpr -> Maybe Int
evaluateCandidate context original@(NeutralExpr syntax) = do
  case syntax of
    Typed Int (NeutralExpr (Lit _)) -> Nothing
    Typed Int (NeutralExpr (Local _ _)) -> Nothing
    Typed Int _ -> pure unit
    _ -> Nothing
  value <- evalStateT (evaluate context Map.empty original) 512
  case value of
    Integer result -> Just result
    _ -> Nothing

-- Fuel and depth bound both work and stack use, even for malformed/cyclic IR.
-- Re-entering a global through a captured closure is rejected independently.
evaluate :: Context -> Env -> NeutralExpr -> Eval Value
evaluate context env (NeutralExpr syntax) = do
  remaining <- get
  lift $ guard (remaining > 0 && context.depth < 64)
  put (remaining - 1)
  let child = context { depth = context.depth + 1 }
      go = evaluate child env
  case syntax of
    Typed _ inner -> go inner
    TypeApp inner _ -> go inner
    Lit (LitInt value) -> pure (Integer value)
    Lit (LitRecord fields) -> do
      evaluated <- traverse (\(Prop name value) -> Tuple name <$> go value) fields
      let values = Map.fromFoldable evaluated
      lift $ guard (Map.size values == Array.length fields)
      pure (Object values)
    Local name level -> lift (Map.lookup (Tuple name level) env)
    Var (Qualified owner name) -> do
      lift $ guard ((owner == Nothing || owner == Just context.moduleName)
        && not (Set.member name context.active))
      value <- lift (Map.lookup name context.globals)
      -- A global initializer is lexically closed over module bindings, never
      -- over locals in whichever function happened to reference it.
      evaluate (child { active = Set.insert name context.active }) Map.empty value
    Abs args body -> pure (Closure (NEA.toArray args) body env context.active)
    Let name level value body -> do
      evaluated <- go value
      evaluate child (Map.insert (Tuple name level) evaluated env) body
    App fn args -> do
      callable <- go fn
      -- All arguments are evaluated, including unused parameters. Likewise
      -- every record/constructor field is checked before any projection.
      arguments <- traverse go (NEA.toArray args)
      applyKnown child callable arguments
    Accessor source (GetProp name) -> do
      value <- go source
      case value of
        Object fields -> lift (Map.lookup name fields)
        _ -> lift Nothing
    CtorSaturated _ _ _ _ fields ->
      Constructed <$> traverse (\(Tuple _ value) -> go value) fields
    PrimOp (Op2 (OpIntNum OpAdd) left right) -> do
      a <- go left
      b <- go right
      case a, b of
        Integer x, Integer y -> pure (Integer (x + y))
        _, _ -> lift Nothing
    -- In particular: imported/foreign calls, effects, Fail, recursive lets,
    -- branches, other primitives and Number conversions are never evaluated.
    _ -> lift Nothing

applyKnown :: Context -> Value -> Array Value -> Eval Value
applyKnown context value arguments = case value of
  Closure parameters body captured active -> do
    let count = min (Array.length parameters) (Array.length arguments)
        env = Map.union (Map.fromFoldable (Array.zip (Array.take count parameters)
          (Array.take count arguments))) captured
        pending = Array.drop count parameters
        supplied = Array.drop count arguments
    if not (Array.null pending) then pure (Closure pending body env active)
    else do
      result <- evaluate (context { active = active }) env body
      if Array.null supplied then pure result
      else applyKnown (context { depth = context.depth + 1 }) result supplied
  _ -> lift Nothing
