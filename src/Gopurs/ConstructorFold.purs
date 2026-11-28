module Gopurs.ConstructorFold (optimize) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (StateT, evalStateT, get, modify_, put)
import Control.Monad.Trans.Class (lift)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, any, foldMap, foldl, foldr)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst, snd)
import Gopurs.ImmediateApplications.Scope (maximumLevel)
import Gopurs.WorkerNames as WorkerNames
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ConstructorType(..), ExprType(..), Ident(..), Literal(..), ModuleName, ProperName(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendSyntax(..), Level(..), Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type Ref = Tuple (Maybe Ident) Level
type Shape =
  { ty :: ExprType
  , name :: String
  , constructorType :: ConstructorType
  , constructors :: Map.Map (Qualified Ident) (Array ExprType)
  }

data Algebra
  = Field Int
  | Constant Int
  | Unary BackendOperator1 Algebra
  | Binary BackendOperator2 Algebra Algebra

type Consumer = { name :: Ident, shape :: Shape, cases :: Map.Map (Qualified Ident) Algebra }
type Plan = { consumer :: Ident, producer :: Ident, worker :: Ident, args :: Array ExprType, binding :: NeutralExpr }
type Scalar = { ty :: ExprType, expr :: NeutralExpr }
type BuildState =
  { next :: Int
  , calls :: Int
  , seen :: Array (Tuple (Array NeutralExpr) Ref)
  , lets :: Array (Tuple Ref NeutralExpr)
  }
type Build = StateT BuildState Maybe
type Producer = { moduleName :: ModuleName, name :: Ident, worker :: Ident, args :: Array ExprType, consumer :: Consumer }

-- Fuse only an immediate composition of a fresh constructor producer and a
-- structurally recursive integer fold. The algebra consumes every field once,
-- in constructor order. Both sides admit only total scalar primitives; unknown
-- calls, effects, borrowed trees, callbacks and recursive local lets stay intact.
-- Original functions remain available to callers that retain the tree.
optimize :: BackendModule -> BackendModule
optimize mod =
  let
    recursive = Array.mapMaybe (\group -> case group.recursive, group.bindings of
      true, [ binding ] | not (Map.member (fst binding) mod.foreign) -> Just binding
      _, _ -> Nothing) mod.bindings
    consumers = Array.mapMaybe (\(Tuple name value) -> recognizeConsumer mod name value) recursive
    found = foldl (collectPlans mod consumers)
      { names: WorkerNames.fromModule mod, plans: [] } recursive
    byPair = Map.fromFoldable (map (\plan -> Tuple (Tuple plan.consumer plan.producer) plan) found.plans)
    rewritten = map (\group -> group { bindings = map (\(Tuple name value) ->
      Tuple name (rewrite mod.name byPair value)) group.bindings }) mod.bindings
    used = foldMap (foldMap (_.used <<< snd) <<< _.bindings) rewritten
    workers = map (\plan -> { recursive: true, bindings: [ Tuple plan.worker plan.binding ] })
      (Array.filter (\plan -> Set.member plan.worker used) found.plans)
    original = map (\group -> group { bindings = map (\(Tuple name result) -> Tuple name result.expr) group.bindings }) rewritten
  in mod { bindings = workers <> original }

collectPlans
  :: BackendModule
  -> Array Consumer
  -> { names :: WorkerNames.Names, plans :: Array Plan }
  -> Tuple Ident NeutralExpr
  -> { names :: WorkerNames.Names, plans :: Array Plan }
collectPlans mod consumers initial (Tuple name value) = foldl collect initial consumers
  where
  collect acc consumer =
    let worker = WorkerNames.fresh "__gopurs_constructor_fold_" name acc.names
    in case recognizeProducer mod.name consumer name worker value of
      Nothing -> acc
      Just plan -> { names: WorkerNames.reserve worker acc.names, plans: Array.snoc acc.plans plan }

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syntax) = case syntax of
  Typed _ value -> strip value
  Syn.TypeApp value _ -> strip value
  _ -> syntax

annotation :: NeutralExpr -> Maybe ExprType
annotation (NeutralExpr syntax) = case syntax of
  Typed ty _ -> Just ty
  Syn.TypeApp value _ -> annotation value
  _ -> Nothing

compatible :: ExprType -> NeutralExpr -> Boolean
compatible ty value = case annotation value of
  Nothing -> true
  Just actual -> actual == ty

arrow :: ExprType -> { args :: Array ExprType, result :: ExprType }
arrow (Func args result) = let rest = arrow result in { args: args <> rest.args, result: rest.result }
arrow result = { args: [], result }

abstractions :: NeutralExpr -> { args :: Array Ref, body :: NeutralExpr }
abstractions value = case strip value of
  Abs args body -> let rest = abstractions body in { args: NEA.toArray args <> rest.args, body: rest.body }
  _ -> { args: [], body: value }

application :: NeutralExpr -> { head :: NeutralExpr, args :: Array NeutralExpr }
application value = case strip value of
  App fn args -> let prior = application fn in { head: prior.head, args: prior.args <> NEA.toArray args }
  _ -> { head: value, args: [] }

global :: ModuleName -> NeutralExpr -> Maybe Ident
global owner value = case strip value of
  Var (Qualified actual name) | actual == Just owner || actual == Nothing -> Just name
  _ -> Nothing

typed :: ExprType -> NeutralExpr -> NeutralExpr
typed ty = NeutralExpr <<< Typed ty

local :: Ref -> NeutralExpr
local (Tuple name level) = typed Int (NeutralExpr (Local name level))

isLocal :: Ref -> NeutralExpr -> Boolean
isLocal ref value = case strip value of
  Local name level -> ref == Tuple name level
  _ -> false

integer :: Int -> NeutralExpr
integer value = typed Int (NeutralExpr (Lit (LitInt value)))

treeShape :: BackendModule -> ExprType -> Maybe Shape
treeShape mod ty = do
  info <- case ty of
    ADT full _ args -> Just { full, args }
    _ -> Nothing
  decl <- Array.find (\decl -> unwrap mod.name <> "." <> decl.name == info.full) mod.dataDecls
  guard (not (Array.null decl.constructors) && Array.length decl.vars == Array.length info.args)
  guard (all (_ == Int) info.args)
  let variables = Map.fromFoldable (Array.zip decl.vars info.args)
      substitute = case _ of
        TypeVar name -> fromMaybe (TypeVar name) (Map.lookup name variables)
        ADT full path args -> ADT full path (map substitute args)
        other -> other
      constructors = Map.fromFoldable (map (\ctor -> Tuple (Qualified (Just mod.name) (Ident ctor.name))
        (map substitute ctor.fields)) decl.constructors)
      fields = Array.concatMap (map substitute <<< _.fields) decl.constructors
  guard (Map.size constructors == Array.length decl.constructors)
  guard (all (\field -> field == Int || field == ty) fields && any (_ == ty) fields)
  pure { ty, name: decl.name, constructors, constructorType: if Array.length decl.constructors == 1 then ProductType else SumType }

recognizeConsumer :: BackendModule -> Ident -> NeutralExpr -> Maybe Consumer
recognizeConsumer mod name value = do
  signature <- arrow <$> annotation value
  ty <- case signature.args, signature.result of
    [ arg ], Int -> Just arg
    _, _ -> Nothing
  shape <- treeShape mod ty
  let lambda = abstractions value
  input <- case lambda.args of
    [ ref ] -> Just ref
    _ -> Nothing
  branches <- case strip lambda.body of
    Branch cases failure -> case strip failure of
      Fail _ -> Just (NEA.toArray cases)
      _ -> Nothing
    _ -> Nothing
  cases <- traverse (\(Pair condition body) -> do
    ctor <- case strip condition of
      PrimOp (Op1 (OpIsTag ctor) source) | isLocal input source -> Just ctor
      _ -> Nothing
    fields <- Map.lookup ctor shape.constructors
    algebra <- recognizeAlgebra mod.name name shape ctor fields input body
    -- Retain all strict recursive demand, and the producer's field order.
    guard (fieldOrder algebra == Array.mapWithIndex (\index _ -> index) fields)
    pure (Tuple ctor algebra)) branches
  let indexed = Map.fromFoldable cases
  guard (Map.size indexed == Array.length cases && Map.keys indexed == Map.keys shape.constructors)
  pure { name, shape, cases: indexed }

fieldIndex :: Shape -> Qualified Ident -> Array ExprType -> Ref -> NeutralExpr -> Maybe Int
fieldIndex shape ctor fields input value = case strip value of
  Accessor source (GetCtorField actual kind typeName ctorName label index) -> do
    guard (isLocal input source && actual == ctor && kind == shape.constructorType)
    guard (typeName == ProperName shape.name && Qualified (case ctor of Qualified owner _ -> owner) ctorName == ctor)
    guard (label == "value" <> show index)
    ty <- Array.index fields index
    guard (compatible ty value)
    pure index
  _ -> Nothing

recognizeAlgebra :: ModuleName -> Ident -> Shape -> Qualified Ident -> Array ExprType -> Ref -> NeutralExpr -> Maybe Algebra
recognizeAlgebra owner self shape ctor fields input value = do
  guard (compatible Int value)
  case strip value of
    Lit (LitInt n) -> Just (Constant n)
    Accessor _ _ -> do
      index <- fieldIndex shape ctor fields input value
      guard (Array.index fields index == Just Int)
      pure (Field index)
    App _ _ -> do
      let call = application value
      guard (global owner call.head == Just self)
      child <- case call.args of
        [ arg ] -> Just arg
        _ -> Nothing
      index <- fieldIndex shape ctor fields input child
      guard (Array.index fields index == Just shape.ty)
      pure (Field index)
    PrimOp (Op1 op child) | op == OpIntNegate || op == OpIntBitNot ->
      Unary op <$> go child
    PrimOp (Op2 op left right) | totalIntOperator op -> Binary op <$> go left <*> go right
    _ -> Nothing
  where
  go = recognizeAlgebra owner self shape ctor fields input

fieldOrder :: Algebra -> Array Int
fieldOrder = case _ of
  Field index -> [ index ]
  Constant _ -> []
  Unary _ child -> fieldOrder child
  Binary _ left right -> fieldOrder left <> fieldOrder right

totalIntOperator :: BackendOperator2 -> Boolean
totalIntOperator = case _ of
  OpIntNum OpAdd -> true
  OpIntNum OpSubtract -> true
  OpIntNum OpMultiply -> true
  OpIntBitAnd -> true
  OpIntBitOr -> true
  OpIntBitXor -> true
  OpIntBitShiftLeft -> true
  OpIntBitShiftRight -> true
  OpIntBitZeroFillShiftRight -> true
  _ -> false

scalar :: Map.Map Ref ExprType -> NeutralExpr -> Maybe Scalar
scalar env value = do
  result <- case strip value of
    Lit (LitInt n) -> Just { ty: Int, expr: integer n }
    Lit (LitBoolean b) -> Just { ty: Boolean, expr: typed Boolean (NeutralExpr (Lit (LitBoolean b))) }
    Local name level -> do
      ty <- Map.lookup (Tuple name level) env
      pure { ty, expr: typed ty (NeutralExpr (Local name level)) }
    PrimOp (Op1 op child) -> do
      next <- scalar env child
      guard ((next.ty == Int && (op == OpIntNegate || op == OpIntBitNot)) || (next.ty == Boolean && op == OpBooleanNot))
      pure { ty: next.ty, expr: typed next.ty (NeutralExpr (PrimOp (Op1 op next.expr))) }
    PrimOp (Op2 op left right) -> do
      a <- scalar env left
      b <- scalar env right
      ty <- case op of
        _ | totalIntOperator op && a.ty == Int && b.ty == Int -> Just Int
        OpIntOrd _ | a.ty == Int && b.ty == Int -> Just Boolean
        OpBooleanAnd | a.ty == Boolean && b.ty == Boolean -> Just Boolean
        OpBooleanOr | a.ty == Boolean && b.ty == Boolean -> Just Boolean
        _ -> Nothing
      pure { ty, expr: typed ty (NeutralExpr (PrimOp (Op2 op a.expr b.expr))) }
    _ -> Nothing
  guard (compatible result.ty value)
  pure result

recognizeProducer :: ModuleName -> Consumer -> Ident -> Ident -> NeutralExpr -> Maybe Plan
recognizeProducer moduleName consumer name worker value = do
  signature <- arrow <$> annotation value
  guard (signature.result == consumer.shape.ty && not (Array.null signature.args) && all (_ == Int) signature.args)
  let lambda = abstractions value
      env = Map.fromFoldable (Array.zip lambda.args signature.args)
      context = { moduleName, name, worker, args: signature.args, consumer }
  args <- NEA.fromArray lambda.args
  guard (Array.length lambda.args == Array.length signature.args && Map.size env == Array.length lambda.args)
  body <- evalStateT (do
    result <- withScope identity (produce context env lambda.body)
    state <- get
    lift (guard (state.calls > 0))
    pure result) { next: maximumLevel value + 1, calls: 0, seen: [], lets: [] }
  pure { consumer: consumer.name, producer: name, worker, args: signature.args
       , binding: typed (Func signature.args Int) (NeutralExpr (Abs args body)) }

fresh :: Build Ref
fresh = do
  state <- get
  put (state { next = state.next + 1 })
  pure (Tuple Nothing (Level state.next))

wrapLets :: Array (Tuple Ref NeutralExpr) -> NeutralExpr -> NeutralExpr
wrapLets bindings body =
  foldr (\(Tuple (Tuple name level) child) rest -> typed Int (NeutralExpr (Let name level child rest))) body bindings

withScope
  :: (Array (Tuple (Array NeutralExpr) Ref) -> Array (Tuple (Array NeutralExpr) Ref))
  -> Build NeutralExpr
  -> Build NeutralExpr
withScope adjustSeen action = do
  outer <- get
  put (outer { seen = adjustSeen outer.seen, lets = [] })
  expr <- action
  inner <- get
  put (inner { seen = outer.seen, lets = outer.lets })
  pure (wrapLets inner.lets expr)

sameExpr :: NeutralExpr -> NeutralExpr -> Boolean
sameExpr a b = case strip a, strip b of
  Lit (LitInt x), Lit (LitInt y) -> x == y
  Lit (LitBoolean x), Lit (LitBoolean y) -> x == y
  Local nameA levelA, Local nameB levelB -> nameA == nameB && levelA == levelB
  PrimOp (Op1 opA childA), PrimOp (Op1 opB childB) -> opA == opB && sameExpr childA childB
  PrimOp (Op2 opA lA rA), PrimOp (Op2 opB lB rB) -> opA == opB && sameExpr lA lB && sameExpr rA rB
  _, _ -> false

sameExprs :: Array NeutralExpr -> Array NeutralExpr -> Boolean
sameExprs xs ys = Array.length xs == Array.length ys && all (\(Tuple a b) -> sameExpr a b) (Array.zip xs ys)

mentionsLevel :: Level -> NeutralExpr -> Boolean
mentionsLevel target value = case strip value of
  Local _ level -> level == target
  PrimOp (Op1 _ child) -> mentionsLevel target child
  PrimOp (Op2 _ left right) -> mentionsLevel target left || mentionsLevel target right
  _ -> false

produce :: Producer -> Map.Map Ref ExprType -> NeutralExpr -> Build NeutralExpr
produce context env value = do
  lift (guard (compatible context.consumer.shape.ty value))
  case strip value of
    Branch cases fallback -> do
      cases' <- traverse (\(Pair condition body) -> do
        test <- lift (scalar env condition)
        lift (guard (test.ty == Boolean))
        Pair test.expr <$> withScope identity (produce context env body)) cases
      fallback' <- withScope identity (produce context env fallback)
      pure (typed Int (NeutralExpr (Branch cases' fallback')))
    Let name level initial body -> do
      initial' <- lift (scalar env initial)
      let keep = Array.filter (\(Tuple args _) -> not (any (mentionsLevel level) args))
      body' <- withScope keep (produce context (Map.insert (Tuple name level) initial'.ty env) body)
      pure (typed Int (NeutralExpr (Let name level initial'.expr body')))
    CtorSaturated ctor kind typeName ctorName fields -> do
      let shape = context.consumer.shape
      fieldTypes <- lift (Map.lookup ctor shape.constructors)
      algebra <- lift (Map.lookup ctor context.consumer.cases)
      lift (guard (kind == shape.constructorType && typeName == ProperName shape.name
        && Qualified (Just context.moduleName) ctorName == ctor && Array.length fields == Array.length fieldTypes))
      lift (guard (map fst fields == Array.mapWithIndex (\index _ -> "value" <> show index) fieldTypes))
      refs <- traverse (\(Tuple ty (Tuple _ field)) -> do
        child <- if ty == Int then do
          result <- lift (scalar env field)
          lift (guard (result.ty == Int))
          pure result.expr
        else produce context env field
        r <- fresh
        modify_ (\state -> state { lets = Array.snoc state.lets (Tuple r child) })
        pure r) (Array.zip fieldTypes fields)
      lift (instantiate refs algebra)
    App _ _ -> do
      let call = application value
      lift (guard (global context.moduleName call.head == Just context.name && Array.length call.args == Array.length context.args))
      args <- lift (traverse (scalar env) call.args)
      lift (guard (all (\arg -> arg.ty == Int) args))
      let argExprs = map _.expr args
      state <- get
      case Array.find (\(Tuple prev _) -> sameExprs prev argExprs) state.seen of
        Just (Tuple _ existing) -> do
          modify_ (\s -> s { calls = s.calls + 1 })
          pure (local existing)
        Nothing -> do
          callExpr <- lift (workerCall context.moduleName context.worker context.args argExprs)
          r <- fresh
          modify_ (\s -> s
            { calls = s.calls + 1
            , lets = Array.snoc s.lets (Tuple r callExpr)
            , seen = Array.snoc s.seen (Tuple argExprs r)
            })
          pure (local r)
    _ -> lift Nothing

instantiate :: Array Ref -> Algebra -> Maybe NeutralExpr
instantiate refs = case _ of
  Field index -> local <$> Array.index refs index
  Constant n -> Just (integer n)
  Unary op child -> typed Int <<< NeutralExpr <<< PrimOp <<< Op1 op <$> instantiate refs child
  Binary op left right -> do
    a <- instantiate refs left
    b <- instantiate refs right
    pure (typed Int (NeutralExpr (PrimOp (Op2 op a b))))

workerCall :: ModuleName -> Ident -> Array ExprType -> Array NeutralExpr -> Maybe NeutralExpr
workerCall owner worker types args = do
  supplied <- NEA.fromArray args
  let fn = typed (Func types Int) (NeutralExpr (Var (Qualified (Just owner) worker)))
  pure (typed Int (NeutralExpr (App fn supplied)))

rewrite :: ModuleName -> Map.Map (Tuple Ident Ident) Plan -> NeutralExpr -> { expr :: NeutralExpr, used :: Set.Set Ident }
rewrite owner plans value@(NeutralExpr syntax) = case fuse of
  Just result -> result
  Nothing ->
    let children = map (rewrite owner plans) syntax
    in { expr: NeutralExpr (map _.expr children), used: foldMap _.used children }
  where
  fuse = do
    let consumed = application value
    consumer <- global owner consumed.head
    input <- case consumed.args of
      [ arg ] -> Just arg
      _ -> Nothing
    let built = application input
    producer <- global owner built.head
    plan <- Map.lookup (Tuple consumer producer) plans
    guard (Array.length built.args == Array.length plan.args && compatible Int value)
    expr <- workerCall owner plan.worker plan.args built.args
    pure { expr, used: Set.singleton plan.worker }
