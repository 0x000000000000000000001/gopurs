module Gopurs.FunctionFusion (optimizeFunctionProducers) where

import Prelude hiding (one)

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (Tuple(..), fst, snd)
import Gopurs.GoAst (sanitizeName)
import PureScript.Backend.Optimizer.Convert (BackendBindingGroup, BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Abs, App, Branch, Let, Lit, Local, PrimOp, Typed, Var), Level(..), Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type LocalRef = Tuple (Maybe Ident) Level
type Group = BackendBindingGroup Ident NeutralExpr
type Producer =
  { counter :: LocalRef
  , callback :: LocalRef
  , seed :: LocalRef
  , body :: NeutralExpr
  }

-- Run after PBO and before TCO. A nonnegative recursive function producer
-- can capture its count instead of constructing a chain of closures. Its
-- callback still runs only when the returned function is applied. Keeping a
-- branch before that function preserves the producer's public arity and its
-- original negative path, including evaluation before a partial application.
optimizeFunctionProducers :: BackendModule -> BackendModule
optimizeFunctionProducers mod =
  let
    names = Set.fromFoldable (Array.concatMap (map fst <<< _.bindings) mod.bindings)
      <> Map.keys mod.foreign
    identities = Set.fromFoldable $ Array.concatMap
      (Array.mapMaybe (\(Tuple name expr) -> if isIdentity expr then Just name else Nothing) <<< _.bindings)
      mod.bindings
    rewritten = foldl (rewriteGroup mod.name identities) { names, bindings: [] } mod.bindings
  in mod { bindings = rewritten.bindings }

rewriteGroup
  :: ModuleName
  -> Set Ident
  -> { names :: Set Ident, bindings :: Array Group }
  -> Group
  -> { names :: Set Ident, bindings :: Array Group }
rewriteGroup moduleName identities acc group = case group.bindings of
  [ Tuple name expr ] | group.recursive -> case recognize moduleName identities name expr of
    Just producer ->
      let
        worker = freshWorker acc.names name 0
        qualifiedWorker = Qualified (Just moduleName) worker
        workerGroup = { recursive: true, bindings: [ Tuple worker (iteratorWorker qualifiedWorker) ] }
        originalGroup = group { bindings = [ Tuple name (compactProducer qualifiedWorker producer) ] }
      in
        { names: Set.insert worker acc.names
        , bindings: acc.bindings <> [ workerGroup, originalGroup ]
        }
    Nothing -> acc { bindings = Array.snoc acc.bindings group }
  _ -> acc { bindings = Array.snoc acc.bindings group }

freshWorker :: Set Ident -> Ident -> Int -> Ident
freshWorker names (Ident original) index =
  let
    name = original <> "__gopurs_counted_function_" <> show index
    candidate = Ident name
    emittedNames = Set.map (\(Ident ident) -> sanitizeName ident) names
  in if Set.member candidate names || Set.member (sanitizeName name) emittedNames
     then freshWorker names (Ident original) (index + 1)
     else candidate

recognize :: ModuleName -> Set Ident -> Ident -> NeutralExpr -> Maybe Producer
recognize current identities name expr = do
  guard (annotation expr == Just producerType)
  lambda <- abstractions producerType expr
  counter <- one lambda.args
  branch <- case at functionType lambda.body of
    Just (Branch branches step) -> case NEA.toArray branches of
      [ Pair condition seed ] -> Just { condition, seed, step }
      _ -> Nothing
    _ -> Nothing
  Tuple left right <- case at Boolean branch.condition of
    Just (PrimOp (Syn.Op2 (Syn.OpIntOrd Syn.OpEq) left right)) -> Just (Tuple left right)
    _ -> Nothing
  guard ((local Int counter left && literal 0 right) || (literal 0 left && local Int counter right))
  guard (identitySeed branch.seed)
  step <- case at functionType branch.step of
    Just (Let _ previous built continuation) -> Just { previous, built, continuation }
    _ -> Nothing
  built <- application functionType step.built
  qualified <- case at producerType built.head of
    Just (Var qualifiedName) -> Just qualifiedName
    _ -> Nothing
  guard (here name qualified)
  decrement <- one built.args
  Tuple n amount <- case at Int decrement of
    Just (PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpSubtract) n amount)) -> Just (Tuple n amount)
    _ -> Nothing
  guard (local Int counter n && literal 1 amount)
  returned <- abstractions functionType step.continuation
  Tuple callback seed <- two returned.args
  applied <- application Int returned.body
  guard (local callbackType callback applied.head)
  inner <- one applied.args
  nested <- application Int inner
  guard (case at functionType nested.head of
    Just (Local _ level) -> level == step.previous
    _ -> false)
  Tuple passedCallback passedSeed <- two nested.args
  guard (local callbackType callback passedCallback && local Int seed passedSeed)
  pure { counter, callback, seed, body: lambda.body }
  where
  here ident (Qualified moduleName found) = ident == found && (moduleName == Nothing || moduleName == Just current)
  identitySeed seed = case at functionType seed of
    Just (Var qualified@(Qualified _ ident)) -> here ident qualified && Set.member ident identities
    _ -> isIdentity seed

callbackType :: ExprType
callbackType = Func [ Int ] Int

functionType :: ExprType
functionType = Func [ callbackType, Int ] Int

producerType :: ExprType
producerType = Func [ Int, callbackType, Int ] Int

typed :: ExprType -> NeutralExpr -> NeutralExpr
typed ty = NeutralExpr <<< Typed ty

reference :: ExprType -> LocalRef -> NeutralExpr
reference ty (Tuple ident level) = typed ty (NeutralExpr (Local ident level))

integer :: Int -> NeutralExpr
integer value = typed Int (NeutralExpr (Lit (LitInt value)))

workerCall :: Qualified Ident -> NeutralExpr -> NeutralExpr -> NeutralExpr -> NeutralExpr
workerCall worker count callback seed = typed Int
  (NeutralExpr (App (typed producerType (NeutralExpr (Var worker))) (NEA.cons' count [ callback, seed ])))

compactProducer :: Qualified Ident -> Producer -> NeutralExpr
compactProducer worker producer =
  let
    count = reference Int producer.counter
    condition = typed Boolean (NeutralExpr (PrimOp (Syn.Op2 (Syn.OpIntOrd Syn.OpGte) count (integer 0))))
    loop = workerCall worker count (reference callbackType producer.callback) (reference Int producer.seed)
    compact = typed functionType (NeutralExpr (Abs (NEA.cons' producer.callback [ producer.seed ]) loop))
    body = typed functionType (NeutralExpr (Branch (NEA.singleton (Pair condition compact)) producer.body))
  in typed producerType (NeutralExpr (Abs (NEA.singleton producer.counter) body))

iteratorWorker :: Qualified Ident -> NeutralExpr
iteratorWorker worker =
  let
    counter = Tuple (Just (Ident "remaining")) (Level 0)
    callback = Tuple (Just (Ident "callback")) (Level 1)
    seed = Tuple (Just (Ident "result")) (Level 2)
    count = reference Int counter
    fn = reference callbackType callback
    result = reference Int seed
    done = typed Boolean (NeutralExpr (PrimOp (Syn.Op2 (Syn.OpIntOrd Syn.OpEq) count (integer 0))))
    nextCount = typed Int (NeutralExpr (PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpSubtract) count (integer 1))))
    nextResult = typed Int (NeutralExpr (App fn (NEA.singleton result)))
    recurse = workerCall worker nextCount fn nextResult
    body = typed Int (NeutralExpr (Branch (NEA.singleton (Pair done result)) recurse))
  in typed producerType (NeutralExpr (Abs (NEA.cons' counter [ callback, seed ]) body))

-- Arrow grouping is irrelevant to this pattern. Representation-changing
-- annotations, dictionary constraints and opaque base expressions are not.
normalize :: ExprType -> ExprType
normalize (ForAll _ ty) = normalize ty
normalize (Func args ret) = case normalize ret of
  Func more result -> Func (map normalize args <> more) result
  result -> Func (map normalize args) result
normalize ty = ty

at :: ExprType -> NeutralExpr -> Maybe (Syn.BackendSyntax NeutralExpr)
at expected (NeutralExpr syn) = case syn of
  Typed ty inner -> do
    guard (normalize ty == normalize expected)
    at expected inner
  Syn.TypeApp inner _ -> at expected inner
  _ -> Just syn

annotation :: NeutralExpr -> Maybe ExprType
annotation (NeutralExpr syn) = case syn of
  Typed ty _ -> Just (normalize ty)
  Syn.TypeApp inner _ -> annotation inner
  _ -> Nothing

abstractions :: ExprType -> NeutralExpr -> Maybe { args :: Array LocalRef, body :: NeutralExpr }
abstractions ty expr = do
  syn <- at ty expr
  case syn of
    Abs parameters body -> do
      arrow <- case normalize ty of
        Func args result -> Just { args, result }
        _ -> Nothing
      let refs = NEA.toArray parameters
          count = Array.length refs
      guard (count <= Array.length arrow.args)
      let rest = Array.drop count arrow.args
          bodyType = if Array.null rest then arrow.result else Func rest arrow.result
      tail <- abstractions bodyType body
      pure { args: refs <> tail.args, body: tail.body }
    _ -> pure { args: [], body: expr }

isIdentity :: NeutralExpr -> Boolean
isIdentity expr = case check of
  Just _ -> true
  Nothing -> false
  where
  check = do
    ty <- annotation expr
    x <- case ty of
      Func [ Func [a] b, x ] result | a == x && b == x && result == x -> Just x
      _ -> Nothing
    lambda <- abstractions ty expr
    Tuple _ value <- two lambda.args
    guard (local x value lambda.body)

one :: forall a. Array a -> Maybe a
one = case _ of
  [ value ] -> Just value
  _ -> Nothing

two :: forall a. Array a -> Maybe (Tuple a a)
two = case _ of
  [ a, b ] -> Just (Tuple a b)
  _ -> Nothing

application :: ExprType -> NeutralExpr -> Maybe { head :: NeutralExpr, args :: Array NeutralExpr }
application ty expr = case at ty expr of
  Just (App head args) -> Just { head, args: NEA.toArray args }
  _ -> Nothing

local :: ExprType -> LocalRef -> NeutralExpr -> Boolean
local ty ref expr = case at ty expr of
  Just (Local _ found) -> found == snd ref
  _ -> false

literal :: Int -> NeutralExpr -> Boolean
literal value expr = case at Int expr of
  Just (Lit (LitInt found)) -> found == value
  _ -> false
