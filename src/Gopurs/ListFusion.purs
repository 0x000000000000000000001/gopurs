module Gopurs.ListFusion (optimize) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (State, evalState, get, put)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), snd)
import Gopurs.ImmediateApplications.Scope (maximumLevel)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendSyntax(..), Level(..), Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type Ref = Tuple (Maybe Ident) Level
type Shape = { ty :: ExprType, empty :: Qualified Ident, cons :: Qualified Ident }
type Loop = { level :: Level, name :: Ident, args :: Array Ref, body :: NeutralExpr, initial :: Array NeutralExpr, signature :: { args :: Array ExprType, result :: ExprType } }
type Folder = { shape :: Shape, name :: Ident }
type Filter = { shape :: Shape, name :: Ident, item :: Ref, predicate :: NeutralExpr }
type Scalar = { ty :: ExprType, expr :: NeutralExpr }

-- A descending constructor loop reverses its visitation order. The admitted
-- filter prepends again, so its integer fold visits elements in the original
-- loop order. Fuse only this closed, immediately consumed pipeline. Its lists
-- never escape, and the moved predicate contains only total scalar primitives.
-- Keep the original bindings for other callers; TCO analyses the new scalar
-- loop afterwards. Source usage annotations are not an ownership proof here.
optimize :: BackendModule -> BackendModule
optimize mod =
  let bindings = Array.concatMap _.bindings mod.bindings
      folders = Map.fromFoldable $ Array.mapMaybe (\(Tuple name value) ->
        map (Tuple name) (recognizeFolder mod name value)) bindings
      filters = Map.fromFoldable $ Array.mapMaybe (\(Tuple name value) ->
        map (Tuple name) (recognizeFilter mod name value)) bindings
  in mod { bindings = map (\group -> group { bindings = map (\(Tuple name value) ->
    Tuple name (evalState (rewrite mod.name folders filters value) (maximumLevel value + 1))) group.bindings }) mod.bindings }

rewrite :: ModuleName -> Map.Map Ident Folder -> Map.Map Ident Filter -> NeutralExpr -> State Int NeutralExpr
rewrite current folders filters original@(NeutralExpr syntax) = do
  base <- get
  case fuse current folders filters base original of
    Just result -> put (base + 3) *> pure result
    Nothing -> NeutralExpr <$> traverse (rewrite current folders filters) syntax

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

arrow :: ExprType -> { args :: Array ExprType, result :: ExprType }
arrow (Func args result) = let rest = arrow result in { args: args <> rest.args, result: rest.result }
arrow ty = { args: [], result: ty }

abstractions :: NeutralExpr -> { args :: Array Ref, body :: NeutralExpr }
abstractions value = case strip value of
  Abs args body -> let rest = abstractions body in { args: NEA.toArray args <> rest.args, body: rest.body }
  _ -> { args: [], body: value }

application :: NeutralExpr -> { head :: NeutralExpr, args :: Array NeutralExpr }
application value = case strip value of
  App fn args -> let prior = application fn in { head: prior.head, args: prior.args <> NEA.toArray args }
  _ -> { head: value, args: [] }

typed :: ExprType -> NeutralExpr -> NeutralExpr
typed ty = NeutralExpr <<< Typed ty

local :: ExprType -> Ref -> NeutralExpr
local ty (Tuple name level) = typed ty (NeutralExpr (Local name level))

isLocal :: Ref -> NeutralExpr -> Boolean
isLocal ref value = case strip value of
  Local name level -> ref == Tuple name level
  _ -> false

global :: ModuleName -> NeutralExpr -> Maybe Ident
global current value = case strip value of
  Var (Qualified owner name) | owner == Just current || owner == Nothing -> Just name
  _ -> Nothing

integer :: Int -> NeutralExpr
integer value = typed Int (NeutralExpr (Lit (LitInt value)))

literal :: Int -> NeutralExpr -> Boolean
literal expected value = case strip value of
  Lit (LitInt found) -> expected == found
  _ -> false

-- Resolve the concrete layout from TAST rather than constructor spelling.
-- Only a local Nil/Cons-shaped ADT instantiated entirely at Int is admitted.
listShape :: BackendModule -> ExprType -> Maybe Shape
listShape mod ty = do
  info <- case ty of
    ADT full _ args -> Just { full, args }
    _ -> Nothing
  decl <- Array.find (\decl -> unwrap mod.name <> "." <> decl.name == info.full) mod.dataDecls
  guard (Array.length decl.vars == Array.length info.args)
  guard (all (_ == Int) info.args)
  let variables = Map.fromFoldable (Array.zip decl.vars info.args)
      substitute = case _ of
        TypeVar name -> fromMaybe (TypeVar name) (Map.lookup name variables)
        ADT name path args -> ADT name path (map substitute args)
        other -> other
  Tuple empty cons <- case decl.constructors of
    [ a, b ] | Array.null a.fields -> Just (Tuple a b)
    [ a, b ] | Array.null b.fields -> Just (Tuple b a)
    _ -> Nothing
  guard (map substitute cons.fields == [ Int, ty ])
  pure { ty, empty: Qualified (Just mod.name) (Ident empty.name), cons: Qualified (Just mod.name) (Ident cons.name) }

emptyList :: Shape -> NeutralExpr -> Boolean
emptyList shape value = case strip value of
  CtorSaturated name _ _ _ fields -> name == shape.empty && Array.null fields
  Var name -> name == shape.empty
  _ -> false

field :: Shape -> Ref -> Int -> NeutralExpr -> Boolean
field shape ref index value = case strip value of
  Accessor base (GetCtorField name _ _ _ _ found) -> name == shape.cons && index == found && isLocal ref base
  _ -> false

consFields :: Shape -> NeutralExpr -> Maybe (Tuple NeutralExpr NeutralExpr)
consFields shape value = case strip value of
  CtorSaturated name _ _ _ [ Tuple _ head, Tuple _ tail ] | name == shape.cons -> Just (Tuple head tail)
  _ -> Nothing

tag :: Qualified Ident -> Ref -> NeutralExpr -> Boolean
tag expected ref value = case strip value of
  PrimOp (Op1 (OpIsTag name) input) -> name == expected && isLocal ref input
  _ -> false

listCases :: Shape -> Ref -> NeutralExpr -> Maybe { empty :: NeutralExpr, cons :: NeutralExpr }
listCases shape ref value = case strip value of
  Branch cases failure -> case NEA.toArray cases, strip failure of
    [ Pair nilTest nilBody, Pair consTest consBody ], Fail _
      | tag shape.empty ref nilTest && tag shape.cons ref consTest -> Just { empty: nilBody, cons: consBody }
    _ , _ -> Nothing
  _ -> Nothing

two :: forall a. Array a -> Maybe (Tuple a a)
two [ a, b ] = Just (Tuple a b)
two _ = Nothing

single :: forall a. Array a -> Maybe a
single [ value ] = Just value
single _ = Nothing

recognizeFolder :: BackendModule -> Ident -> NeutralExpr -> Maybe Folder
recognizeFolder mod name value = do
  signature <- arrow <$> annotation value
  shape <- case signature.args, signature.result of
    [ Func [ Int, Int ] Int, Int, ty ], Int -> listShape mod ty
    _ , _ -> Nothing
  let lambda = abstractions value
  args <- case lambda.args of
    [ _, acc, input ] -> Just { acc, input }
    _ -> Nothing
  branches <- listCases shape args.input lambda.body
  guard (isLocal args.acc branches.empty)
  let recur = application branches.cons
  guard (global mod.name recur.head == Just name)
  step <- case recur.args of
    [ _, next, tail ] | field shape args.input 1 tail -> Just next
    _ -> Nothing
  guard case strip step of
    PrimOp (Op2 (OpIntNum OpAdd) acc head) -> isLocal args.acc acc && field shape args.input 0 head
    _ -> false
  -- The unused callback must remain a harmless value at every recursive call.
  -- A delayed/opaque computation here cannot be discarded during fusion.
  ignored <- Array.head recur.args
  guard (harmlessFunction ignored)
  pure { shape, name }

harmlessFunction :: NeutralExpr -> Boolean
harmlessFunction value = case strip value of
  Var (Qualified (Just owner) (Ident name)) -> unwrap owner == "Data.Semiring" && name == "intAdd"
  _ -> false

localLoop :: NeutralExpr -> Maybe Loop
localLoop value = case strip value of
  LetRec level bindings body -> do
    let entry = application body
    name <- case strip entry.head of
      Local (Just name) found | found == level -> Just name
      _ -> Nothing
    selected <- snd <$> Array.find (\(Tuple ident _) -> ident == name) (NEA.toArray bindings)
    signature <- arrow <$> annotation selected
    -- Unselected recursive bindings are only function values, never effects.
    guard (all (not <<< Array.null <<< _.args <<< abstractions <<< snd) bindings)
    let lambda = abstractions selected
    pure { level, name, args: lambda.args, body: lambda.body, initial: entry.args, signature }
  _ -> Nothing

selfCall :: Loop -> NeutralExpr -> Maybe (Array NeutralExpr)
selfCall loop value = do
  let call = application value
  guard (isLocal (Tuple (Just loop.name) loop.level) call.head)
  pure call.args

recognizeFilter :: BackendModule -> Ident -> NeutralExpr -> Maybe Filter
recognizeFilter mod name value = do
  signature <- arrow <$> annotation value
  ty <- single signature.args
  guard (ty == signature.result)
  shape <- listShape mod ty
  let lambda = abstractions value
  input <- single lambda.args
  loop <- localLoop lambda.body
  guard (loop.signature.args == [ ty, ty ] && loop.signature.result == ty)
  Tuple item acc <- two loop.args
  Tuple initial seed <- two loop.initial
  guard (isLocal input initial && emptyList shape seed)
  branches <- listCases shape item loop.body
  guard (isLocal acc branches.empty)
  branch <- case strip branches.cons of
    Branch cases no -> case NEA.toArray cases of
      [ Pair predicate yes ] -> Just { predicate, yes, no }
      _ -> Nothing
    _ -> Nothing
  Tuple yesTail yesAcc <- selfCall loop branch.yes >>= two
  Tuple noTail noAcc <- selfCall loop branch.no >>= two
  Tuple head tail <- consFields shape yesAcc
  guard (field shape item 1 yesTail && field shape item 1 noTail && isLocal acc noAcc)
  guard (field shape item 0 head && isLocal acc tail)
  predicate <- scalar (\expr -> if field shape item 0 expr then Just (integer 0) else Nothing) branch.predicate
  guard (predicate.ty == Boolean)
  pure { shape, name, item, predicate: branch.predicate }

-- No arbitrary calls, effect nodes, callbacks, partial matches or captured
-- values may move from the list traversal to the producer. Modulo requires a
-- nonzero constant divisor; division and floating-point operations are rejected.
scalar :: (NeutralExpr -> Maybe NeutralExpr) -> NeutralExpr -> Maybe Scalar
scalar replace value = case replace value of
  Just expr -> Just { ty: Int, expr }
  Nothing -> case strip value of
    Lit (LitInt _) -> Just { ty: Int, expr: typed Int (NeutralExpr (strip value)) }
    Lit (LitBoolean _) -> Just { ty: Boolean, expr: typed Boolean (NeutralExpr (strip value)) }
    PrimOp (Op1 OpBooleanNot child) -> do
      next <- scalar replace child
      guard (next.ty == Boolean)
      pure { ty: Boolean, expr: typed Boolean (NeutralExpr (PrimOp (Op1 OpBooleanNot next.expr))) }
    PrimOp (Op2 op left right) -> do
      a <- scalar replace left
      b <- scalar replace right
      result <- case op of
        OpIntNum OpAdd | a.ty == Int && b.ty == Int -> Just Int
        OpIntNum OpSubtract | a.ty == Int && b.ty == Int -> Just Int
        OpIntNum OpMultiply | a.ty == Int && b.ty == Int -> Just Int
        OpIntNum OpMod | a.ty == Int && b.ty == Int -> case strip right of
          Lit (LitInt divisor) | divisor /= 0 -> Just Int
          _ -> Nothing
        OpIntOrd _ | a.ty == Int && b.ty == Int -> Just Boolean
        OpBooleanAnd | a.ty == Boolean && b.ty == Boolean -> Just Boolean
        OpBooleanOr | a.ty == Boolean && b.ty == Boolean -> Just Boolean
        _ -> Nothing
      pure { ty: result, expr: typed result (NeutralExpr (PrimOp (Op2 op a.expr b.expr))) }
    _ -> Nothing

fuse :: ModuleName -> Map.Map Ident Folder -> Map.Map Ident Filter -> Int -> NeutralExpr -> Maybe NeutralExpr
fuse current folders filters base original = do
  let call = application original
  folder <- global current call.head >>= flip Map.lookup folders
  args <- case call.args of
    [ callback, seed, input ] | harmlessFunction callback -> Just { seed, input }
    _ -> Nothing
  -- The seed is evaluated before the producer in the original application.
  -- Restrict it to an already evaluated scalar, so moving it into the loop
  -- cannot change strictness or introduce captures from another local scope.
  guard case strip args.seed of
    Lit (LitInt _) -> true
    Local _ _ -> annotation args.seed == Just Int
    _ -> false
  let filtered = application args.input
  filter <- global current filtered.head >>= flip Map.lookup filters
  guard (filter.shape.ty == folder.shape.ty)
  produced <- single filtered.args
  loop <- localLoop produced
  guard (loop.signature.args == [ Int, folder.shape.ty ] && loop.signature.result == folder.shape.ty)
  Tuple counter acc <- two loop.args
  Tuple initial nil <- two loop.initial
  guard (emptyList folder.shape nil)
  guard case strip initial of
    Lit (LitInt _) -> true
    Local _ _ -> annotation initial == Just Int
    _ -> false
  branch <- case strip loop.body of
    Branch cases step -> case NEA.toArray cases of
      [ Pair condition result ] | isLocal acc result -> Just { condition, step }
      _ -> Nothing
    _ -> Nothing
  guard case strip branch.condition of
    PrimOp (Op2 (OpIntOrd OpLt) test limit) -> isLocal counter test && case strip limit of
      Lit (LitInt _) -> true
      _ -> false
    _ -> false
  Tuple next constructed <- selfCall loop branch.step >>= two
  guard case strip next of
    PrimOp (Op2 (OpIntNum OpSubtract) value amount) -> isLocal counter value && literal 1 amount
    _ -> false
  Tuple head tail <- consFields folder.shape constructed
  guard (isLocal counter head && isLocal acc tail)
  -- The binding-wide supply is above all original levels and advances for
  -- each fusion, including siblings in the same expression.
  let worker = Tuple (Just (Ident "listFold")) (Level base)
      count = Tuple (Just (Ident "remaining")) (Level (base + 1))
      total = Tuple (Just (Ident "total")) (Level (base + 2))
      countExpr = local Int count
      totalExpr = local Int total
      workerType = Func [ Int, Int ] Int
      replaceCounter value = if isLocal counter value then Just countExpr else Nothing
  condition <- scalar replaceCounter branch.condition
  decrement <- scalar replaceCounter next
  predicate <- scalar (\expr -> if field filter.shape filter.item 0 expr then Just countExpr else Nothing) filter.predicate
  let added = typed Int (NeutralExpr (PrimOp (Op2 (OpIntNum OpAdd) totalExpr countExpr)))
      selected = typed Int (NeutralExpr (Branch (NEA.singleton (Pair predicate.expr added)) totalExpr))
      recur = typed Int (NeutralExpr (App (local workerType worker) (NEA.cons' decrement.expr [ selected ])))
      body = typed Int (NeutralExpr (Branch (NEA.singleton (Pair condition.expr totalExpr)) recur))
      function = typed workerType (NeutralExpr (Abs (NEA.cons' count [ total ]) body))
      entry = typed Int (NeutralExpr (App (local workerType worker) (NEA.cons' initial [ args.seed ])))
  pure $ typed Int (NeutralExpr (LetRec (Level base) (NEA.singleton (Tuple (Ident "listFold") function)) entry))
