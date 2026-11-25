-- Evaluate only completely known, pure integer expressions. Dictionary methods
-- are ordinary immutable records and closures here; no class or benchmark name
-- participates in the proof. Unknown syntax leaves the source expression intact.
module Gopurs.ClosedScalar (optimize) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.Rec.Class (Step(..), tailRecM)
import Control.Monad.State (State, StateT, evalState, get, put, runStateT)
import Control.Monad.Trans.Class (lift)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldr)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.ImmediateApplications.Scope (maximumLevel)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(Int), Ident, Literal(..), ModuleName, Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendSyntax(..), Level(..), Pair(..))

type Ref = Tuple (Maybe Ident) Level
type Env = Map.Map Ref Value
type Context =
  { moduleName :: ModuleName
  , globals :: Map.Map Ident NeutralExpr
  , active :: Set.Set Ident
  , depth :: Int
  , stackDepth :: Int
  }

data Value
  = Integer Int
  | BooleanValue Boolean
  | GlobalValue (Qualified Ident)
  | Object (Map.Map String Value)
  | Closure (Array Ref) NeutralExpr Env (Set.Set Ident)
  | Constructed (Array Value)

type EvalState = { fuel :: Int, reads :: Array (Qualified Ident) }
type Eval = StateT EvalState Maybe
type Evaluation = { context :: Context, env :: Env, expression :: NeutralExpr }

optimize :: BackendModule -> BackendModule
optimize mod =
  let globals = Map.fromFoldable $ Array.filter
        (\(Tuple name _) -> not (Map.member name mod.foreign))
        (Array.concatMap _.bindings mod.bindings)
      context = { moduleName: mod.name, globals, active: Set.empty, depth: 0, stackDepth: 0 }
  in mod { bindings = map (\group -> group { bindings = map
    (\(Tuple name value) -> Tuple name (evalState (rewrite context value) (maximumLevel value + 1))) group.bindings }) mod.bindings }

rewrite :: Context -> NeutralExpr -> State Int NeutralExpr
rewrite context original = rewriteCandidate context original original

rewriteCandidate :: Context -> NeutralExpr -> NeutralExpr -> State Int NeutralExpr
rewriteCandidate context candidate original@(NeutralExpr syntax)
  | context.depth >= 64 = pure original
  | otherwise = case evaluateCandidate context candidate of
      Just result -> do
        level <- get
        put (level + Array.length result.reads)
        pure $ foldr (\(Tuple index name) body -> NeutralExpr (Typed Int (NeutralExpr
          (Let Nothing (Level (level + index)) (NeutralExpr (Var name)) body))))
          (NeutralExpr (Typed Int (NeutralExpr (Lit (LitInt result.value)))))
          (Array.mapWithIndex Tuple result.reads)
      Nothing ->
        let child = context { depth = context.depth + 1 }
        in case syntax of
          -- PBO can leave operand annotations implicit. Integer primitives
          -- establish their operand type without guessing from a binding name.
          PrimOp (Op2 op@(OpIntNum _) left right) -> do
            lhs <- rewriteInteger child left
            rhs <- rewriteInteger child right
            pure (NeutralExpr (PrimOp (Op2 op lhs rhs)))
          _ -> NeutralExpr <$> traverse (rewrite child) syntax

rewriteInteger :: Context -> NeutralExpr -> State Int NeutralExpr
rewriteInteger context original@(NeutralExpr syntax) = case syntax of
  Typed _ _ -> rewrite context original
  Lit _ -> rewrite context original
  Local _ _ -> rewrite context original
  _ -> rewriteCandidate context (NeutralExpr (Typed Int original)) original

-- Keep the output domain explicit in the TAST. A closure/record/constructor is
-- only an intermediate interpreter value, never a newly emitted shared object.
evaluateCandidate :: Context -> NeutralExpr -> Maybe { value :: Int, reads :: Array (Qualified Ident) }
evaluateCandidate context original@(NeutralExpr syntax) = do
  case syntax of
    Typed Int (NeutralExpr (Lit _)) -> Nothing
    Typed Int (NeutralExpr (Local _ _)) -> Nothing
    Typed Int _ -> pure unit
    _ -> Nothing
  Tuple value state <- runStateT (evaluate context Map.empty original) { fuel: 8192, reads: [] }
  case value of
    Integer result -> Just { value: result, reads: state.reads }
    _ -> Nothing

-- Fuel bounds total work; tail calls and selected branches use the trampoline.
-- Unknown global values may be carried in records, but every strict read is
-- retained in the replacement. Calls through those values remain unsupported.
evaluate :: Context -> Env -> NeutralExpr -> Eval Value
evaluate context env expression = tailRecM evaluateStep { context, env, expression }

evaluateStep :: Evaluation -> Eval (Step Evaluation Value)
evaluateStep { context, env, expression: NeutralExpr syntax } = do
  state <- get
  lift $ guard (state.fuel > 0 && context.depth < 1024 && context.stackDepth < 64)
  put (state { fuel = state.fuel - 1 })
  let child = context { depth = context.depth + 1 }
      go = evaluate (child { stackDepth = child.stackDepth + 1 }) env
      next expression = pure (Loop { context: child, env, expression })
      done = pure <<< Done
  case syntax of
    Typed _ inner -> next inner
    TypeApp inner _ -> next inner
    Lit (LitInt value) -> done (Integer value)
    Lit (LitBoolean value) -> done (BooleanValue value)
    Lit (LitRecord fields) -> do
      evaluated <- traverse (\(Prop name value) -> Tuple name <$> go value) fields
      let values = Map.fromFoldable evaluated
      lift $ guard (Map.size values == Array.length fields)
      done (Object values)
    Local name level -> Done <$> lift (Map.lookup (Tuple name level) env)
    Var qualified@(Qualified owner name) ->
      if (owner == Nothing || owner == Just context.moduleName) && Map.member name context.globals then do
        lift $ guard (not (Set.member name context.active))
        value <- lift (Map.lookup name context.globals)
        -- Initializers are closed over module bindings, never caller locals.
        pure (Loop { context: child { active = Set.insert name context.active }, env: Map.empty, expression: value })
      else do
        observed <- get
        -- Global bindings are immutable. Keep their first strict read in order,
        -- including potentially failing initialization, rather than assuming
        -- imported constants are harmless or knowing any library's names.
        if Array.elem qualified observed.reads then pure unit
        else put (observed { reads = Array.snoc observed.reads qualified })
        done (GlobalValue qualified)
    Abs args body -> done (Closure (NEA.toArray args) body env context.active)
    Let name level value body -> do
      evaluated <- go value
      pure (Loop { context: child, env: Map.insert (Tuple name level) evaluated env, expression: body })
    App fn args -> do
      callable <- go fn
      arguments <- traverse go (NEA.toArray args)
      applyKnown child callable arguments
    Branch branches otherwise -> chooseBranch child env (NEA.toArray branches) otherwise
    Accessor source (GetProp name) -> do
      value <- go source
      case value of
        Object fields -> Done <$> lift (Map.lookup name fields)
        _ -> lift Nothing
    CtorSaturated _ _ _ _ fields ->
      Done <<< Constructed <$> traverse (\(Tuple _ value) -> go value) fields
    PrimOp (Op2 (OpIntNum operator) left right) -> do
      a <- go left
      b <- go right
      case a, b, operator of
        Integer x, Integer y, OpAdd -> done (Integer (x + y))
        Integer x, Integer y, OpSubtract -> done (Integer (x - y))
        _, _, _ -> lift Nothing
    PrimOp (Op2 (OpIntOrd OpEq) left right) -> do
      a <- go left
      b <- go right
      case a, b of
        Integer x, Integer y -> done (BooleanValue (x == y))
        _, _ -> lift Nothing
    -- Effects, Fail, recursive lets, unknown calls and conversions stay intact.
    _ -> lift Nothing

chooseBranch :: Context -> Env -> Array (Pair NeutralExpr) -> NeutralExpr -> Eval (Step Evaluation Value)
chooseBranch context env branches otherwise = case Array.uncons branches of
  Nothing -> pure (Loop { context, env, expression: otherwise })
  Just { head: Pair condition value, tail } -> do
    result <- evaluate (context { stackDepth = context.stackDepth + 1 }) env condition
    case result of
      BooleanValue true -> pure (Loop { context, env, expression: value })
      BooleanValue false -> chooseBranch context env tail otherwise
      _ -> lift Nothing

applyKnown :: Context -> Value -> Array Value -> Eval (Step Evaluation Value)
applyKnown context value arguments = case value of
  Closure parameters body captured active -> do
    let count = min (Array.length parameters) (Array.length arguments)
        env = Map.union (Map.fromFoldable (Array.zip (Array.take count parameters)
          (Array.take count arguments))) captured
        pending = Array.drop count parameters
        supplied = Array.drop count arguments
    if not (Array.null pending) then pure (Done (Closure pending body env active))
    else if Array.null supplied then
      pure (Loop { context: context { active = Set.empty }, env, expression: body })
    else do
      result <- evaluate (context { active = Set.empty, stackDepth = context.stackDepth + 1 }) env body
      applyKnown (context { depth = context.depth + 1 }) result supplied
  _ -> lift Nothing
