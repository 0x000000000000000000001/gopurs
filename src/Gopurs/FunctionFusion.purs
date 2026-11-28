module Gopurs.FunctionFusion
  ( optimizeFunctionProducers
  , iteratorWorkerGuard
  , countedWorkerCallCount
  , countedCompositionOperands
  ) where

import Prelude hiding (one)

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..), snd)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext)
import Gopurs.GoAst (GoExpr(..), GoType(..), sanitizeName)
import Gopurs.GoFunctions (Parameters)
import Gopurs.WorkerNames as WorkerNames
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.Convert (BackendBindingGroup, BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
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
    names = WorkerNames.fromModule mod
    identities = Set.fromFoldable $ Array.concatMap
      (Array.mapMaybe (\(Tuple name expr) -> if isIdentity expr then Just name else Nothing) <<< _.bindings)
      mod.bindings
    rewritten = foldl (rewriteGroup mod.name identities) { names, bindings: [] } mod.bindings
  in mod { bindings = rewritten.bindings }

rewriteGroup
  :: ModuleName
  -> Set Ident
  -> { names :: WorkerNames.Names, bindings :: Array Group }
  -> Group
  -> { names :: WorkerNames.Names, bindings :: Array Group }
rewriteGroup moduleName identities acc group = case group.bindings of
  [ Tuple name expr ] | group.recursive -> case recognize moduleName identities name expr of
    Just producer ->
      let
        worker = WorkerNames.fresh "__gopurs_counted_function_" name acc.names
        qualifiedWorker = Qualified (Just moduleName) worker
        workerGroup = { recursive: true, bindings: [ Tuple worker (iteratorWorker qualifiedWorker) ] }
        originalGroup = group { bindings = [ Tuple name (compactProducer qualifiedWorker producer) ] }
      in
        { names: WorkerNames.reserve worker acc.names
        , bindings: acc.bindings <> [ workerGroup, originalGroup ]
        }
    Nothing -> acc { bindings = Array.snoc acc.bindings group }
  _ -> acc { bindings = Array.snoc acc.bindings group }

-- Admission proves the complete recurrence before any worker is allocated or
-- emitted: identity at zero, decrement by one, then callback(previous f seed).
recognize :: ModuleName -> Set Ident -> Ident -> NeutralExpr -> Maybe Producer
recognize current identities name expr = do
  guard (annotation expr == Just producerType)
  lambda <- abstractions producerType expr
  counter <- one lambda.args
  branch <- zeroCase counter lambda.body
  guard (identitySeed current identities branch.seed)
  previous <- recursiveStep current name counter branch.step
  returned <- callbackChain previous.level previous.continuation
  pure { counter, callback: returned.callback, seed: returned.seed, body: lambda.body }

type ZeroCase = { seed :: NeutralExpr, step :: NeutralExpr }

zeroCase :: LocalRef -> NeutralExpr -> Maybe ZeroCase
zeroCase counter body = do
  branch <- case at functionType body of
    Just (Branch branches step) -> case NEA.toArray branches of
      [ Pair condition seed ] -> Just { condition, seed, step }
      _ -> Nothing
    _ -> Nothing
  Tuple left right <- case at Boolean branch.condition of
    Just (PrimOp (Syn.Op2 (Syn.OpIntOrd Syn.OpEq) left right)) -> Just (Tuple left right)
    _ -> Nothing
  guard ((local Int counter left && literal 0 right) || (literal 0 left && local Int counter right))
  pure { seed: branch.seed, step: branch.step }

recursiveStep :: ModuleName -> Ident -> LocalRef -> NeutralExpr -> Maybe { level :: Level, continuation :: NeutralExpr }
recursiveStep current name counter expr = do
  step <- case at functionType expr of
    Just (Let _ previous built continuation) -> Just { previous, built, continuation }
    _ -> Nothing
  built <- application functionType step.built
  qualified <- case at producerType built.head of
    Just (Var qualifiedName) -> Just qualifiedName
    _ -> Nothing
  guard (sameBinding current name qualified)
  decrement <- one built.args
  Tuple n amount <- case at Int decrement of
    Just (PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpSubtract) n amount)) -> Just (Tuple n amount)
    _ -> Nothing
  guard (local Int counter n && literal 1 amount)
  pure { level: step.previous, continuation: step.continuation }

callbackChain :: Level -> NeutralExpr -> Maybe { callback :: LocalRef, seed :: LocalRef }
callbackChain previous continuation = do
  returned <- abstractions functionType continuation
  Tuple callback seed <- two returned.args
  applied <- application Int returned.body
  guard (local callbackType callback applied.head)
  inner <- one applied.args
  nested <- application Int inner
  guard (case at functionType nested.head of
    Just (Local _ level) -> level == previous
    _ -> false)
  Tuple passedCallback passedSeed <- two nested.args
  guard (local callbackType callback passedCallback && local Int seed passedSeed)
  pure { callback, seed }

sameBinding :: ModuleName -> Ident -> Qualified Ident -> Boolean
sameBinding current ident (Qualified moduleName found) =
  ident == found && (moduleName == Nothing || moduleName == Just current)

identitySeed :: ModuleName -> Set Ident -> NeutralExpr -> Boolean
identitySeed current identities seed = case at functionType seed of
  Just (Var qualified@(Qualified _ ident)) -> sameBinding current ident qualified && Set.member ident identities
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
    -- Preserve the entire negative path before the returned lambda, including
    -- its divergence and the evaluation boundary of a partial application.
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

tcoLocalName :: TcoExpr -> Maybe String
tcoLocalName expr = case unwrapTcoExpr expr of
  Syn.Local ident level -> Just (localId ident level)
  _ -> Nothing

tcoIntLiteral :: Int -> TcoExpr -> Boolean
tcoIntLiteral expected expr = case unwrapTcoExpr expr of
  Syn.Lit (LitInt found) -> found == expected
  _ -> false

-- Structural recognition of the counted-function iterator worker so
-- ModuleWorkers can dispatch nonnegative counts through CountedIntApply while
-- keeping the original loop intact for negative counts.
iteratorWorkerGuard :: String -> Parameters -> GoType -> TcoExpr -> Maybe GoExpr
iteratorWorkerGuard workerIdent params resultType body = do
  guard (resultType == TypeInt64)
  { rem, cb, res } <- case params of
    [ Tuple rem TypeInt64, Tuple cb TypeValue, Tuple res TypeInt64 ]
      | rem /= cb && rem /= res && cb /= res -> Just { rem, cb, res }
    _ -> Nothing
  { cond, base, step } <- case unwrapTcoExpr body of
    Syn.Branch branches step -> case NEA.toArray branches of
      [ Pair cond base ] -> Just { cond, base, step }
      _ -> Nothing
    _ -> Nothing
  guard (tcoLocalName base == Just res)
  guard case unwrapTcoExpr cond of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntOrd Syn.OpEq) left right) ->
      (tcoLocalName left == Just rem && tcoIntLiteral 0 right)
        || (tcoIntLiteral 0 left && tcoLocalName right == Just rem)
    _ -> false
  let Tuple head spine = collectGoSpine step
      args = getGoSpineArgs spine
  guard case unwrapTcoExpr head of
    Syn.Var (Qualified _ (Ident target)) -> sanitizeName target == workerIdent
    _ -> false
  { dec, cbArg, nextRes } <- case args of
    [ dec, cbArg, nextRes ] -> Just { dec, cbArg, nextRes }
    _ -> Nothing
  guard (tcoLocalName cbArg == Just cb)
  guard case unwrapTcoExpr dec of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpSubtract) left right) ->
      tcoLocalName left == Just rem && tcoIntLiteral 1 right
    _ -> false
  let Tuple cbHead cbSpine = collectGoSpine nextRes
      cbArgs = getGoSpineArgs cbSpine
  guard (tcoLocalName cbHead == Just cb)
  guard (case cbArgs of
    [ arg ] -> tcoLocalName arg == Just res
    _ -> false)
  let remVar = GoVar (rem <> "_loop")
      cbVar = GoVar (cb <> "_loop")
      resVar = GoVar (res <> "_loop")
  pure $ GoIfElse
    (GoBinOp ">=" remVar (GoCall (GoVar "int64") [ GoInt 0 ]))
    [ GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "CountedIntApply") [ remVar, cbVar, resVar ]) ]
    []

-- Recognize the two-argument closure emitted by compactProducer around a
-- module-local counted iterator worker so CountedIntIterator can attach its
-- repetition metadata for partial applications and nested compositions.
countedWorkerCallCount :: ExprContext -> Array String -> TcoExpr -> Maybe GoExpr
countedWorkerCallCount context params body = do
  Tuple fParam xParam <- two params
  let Tuple head spine = collectGoSpine body
      args = getGoSpineArgs spine
  Qualified mbMod (Ident workerName) <- case unwrapTcoExpr head of
    Syn.Var qualified -> Just qualified
    _ -> Nothing
  let isLocal = mbMod == Nothing
        || map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) mbMod == Just context.modNameStr
  guard isLocal
  guard (String.contains (Pattern "__gopurs_counted_function_") workerName)
  info <- Map.lookup (sanitizeName workerName) context.moduleFunctions
  guard (info.fArgs == [ TypeInt64, TypeValue, TypeInt64 ] && info.fRet == TypeInt64)
  { countExpr, fArg, xArg } <- case args of
    [ countExpr, fArg, xArg ] -> Just { countExpr, fArg, xArg }
    _ -> Nothing
  guard (tcoLocalName fArg == Just fParam && tcoLocalName xArg == Just xParam)
  countKey <- tcoLocalName countExpr
  binding <- Map.lookup countKey context.bound
  guard (binding.goType == TypeInt64 && binding.name /= fParam && binding.name /= xParam)
  pure (GoVar binding.name)

-- Recognize higher-order iterator composition (\f x -> m (n f) x) where m and
-- n are already-evaluated local values. ComposeCountedIntIterators checks at
-- runtime whether both carry nonnegative counted-iterator metadata and falls
-- back to the original closure otherwise.
countedCompositionOperands :: ExprContext -> Array String -> TcoExpr -> Maybe (Tuple GoExpr GoExpr)
countedCompositionOperands context params body = do
  Tuple fParam xParam <- two params
  let Tuple outerHead outerSpine = collectGoSpine body
      outerArgs = getGoSpineArgs outerSpine
  mKey <- tcoLocalName outerHead
  mBinding <- Map.lookup mKey context.bound
  guard (mBinding.goType == TypeValue && mBinding.name /= fParam && mBinding.name /= xParam)
  Tuple innerApp xArg <- two outerArgs
  guard (tcoLocalName xArg == Just xParam)
  let Tuple innerHead innerSpine = collectGoSpine innerApp
      innerArgs = getGoSpineArgs innerSpine
  nKey <- tcoLocalName innerHead
  nBinding <- Map.lookup nKey context.bound
  guard (nBinding.goType == TypeValue && nBinding.name /= fParam && nBinding.name /= xParam)
  fArg <- one innerArgs
  guard (tcoLocalName fArg == Just fParam)
  pure (Tuple (GoVar mBinding.name) (GoVar nBinding.name))
