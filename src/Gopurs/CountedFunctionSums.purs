module Gopurs.CountedFunctionSums (optimize) where

import Prelude hiding (add, apply, compare)

import Control.Alternative (guard)
import Control.Monad.State (State, StateT, evalState, get, put, runStateT)
import Control.Monad.Trans.Class (lift)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set (Set)
import Data.Set as Set
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.ImmediateApplications.Scope (maximumLevel)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level(..), Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type Ref = Tuple (Maybe Ident) Level
type Env = Map.Map Ref Value
type Context =
  { owner :: ModuleName
  , globals :: Map.Map Ident NeutralExpr
  , producers :: Set Ident
  , active :: Set Ident
  , outer :: Boolean
  , depth :: Int
  }

data Scalar
  = Constant Int
  | Input Ref
  | Parameter Int
  | Sum Scalar Scalar
  | Product Scalar Scalar

derive instance eqScalar :: Eq Scalar

data Value
  = Integer Scalar
  | Closure (Array Ref) NeutralExpr Env (Set Ident) Boolean
  | Producer
  | Iterator Scalar
  | Repeated Scalar Value

type Proof =
  { fuel :: Int
  , next :: Int
  , inputs :: Set Ref
  , counts :: Array Scalar
  , used :: Boolean
  }
type Eval = StateT Proof Maybe

-- Only fully consumed, structurally known function iterators are reduced.
-- Their arbitrary pure wrappers are interpreted with bounded fuel; foreign
-- calls, effects, opaque callbacks and recursive initialization are rejected.
-- A runtime guard retains the original expression for negative counts and
-- values outside signed Int32. No benchmark names or chosen inputs enter here.
optimize :: Set Ident -> BackendModule -> BackendModule
optimize producers mod
  | Set.isEmpty producers = mod
  | otherwise =
      let
        globals = Map.fromFoldable $ Array.filter (\(Tuple name _) -> not (Map.member name mod.foreign))
          (Array.concatMap _.bindings (Array.filter (not <<< _.recursive) mod.bindings))
        context = { owner: mod.name, globals, producers, active: Set.empty, outer: true, depth: 0 }
      in mod { bindings = map (\group -> group { bindings = map (\(Tuple name body) ->
        Tuple name (rewrite context body)) group.bindings }) mod.bindings }

rewrite :: Context -> NeutralExpr -> NeutralExpr
rewrite context original@(NeutralExpr syntax) = case candidate context original of
  Just rewritten -> rewritten
  Nothing -> NeutralExpr (map (rewrite context) syntax)

candidate :: Context -> NeutralExpr -> Maybe NeutralExpr
candidate context original@(NeutralExpr syntax) = do
  case syntax of
    Typed Int _ -> pure unit
    _ -> Nothing
  Tuple result proof <- runStateT (evaluate context Map.empty original)
    { fuel: 768, next: 0, inputs: Set.empty, counts: [], used: false }
  scalar <- case result of
    Integer value | proof.used && closed value -> Just value
    _ -> Nothing
  let
    counts = Array.nubEq proof.counts
    nonNegativeInputs = Set.fromFoldable (Array.mapMaybe (case _ of
      Input ref -> Just ref
      _ -> Nothing) counts)
    inputGuards = Array.concatMap (\ref ->
      (if Set.member ref nonNegativeInputs then [] else [ compare Syn.OpGte (reference ref) (integer (-2147483647 - 1)) ])
      <> [ compare Syn.OpLte (reference ref) (integer 2147483647) ]
      ) (Set.toUnfoldable proof.inputs :: Array Ref)
    emit = do
      countGuards <- traverse (\count -> do
        value <- emitScalar count
        pure (compare Syn.OpGte value (integer 0))) counts
      resultExpr <- emitScalar scalar
      let conditions = inputGuards <> countGuards
      pure case NEA.fromArray conditions of
        Nothing -> resultExpr
        Just nonempty -> typed Int (NeutralExpr (Branch
          (NEA.singleton (Pair (foldl (binary Boolean Syn.OpBooleanAnd) (NEA.head nonempty) (NEA.tail nonempty)) resultExpr)) original))
  pure (evalState emit (maximumLevel original + 1))

evaluate :: Context -> Env -> NeutralExpr -> Eval Value
evaluate context env (NeutralExpr syntax) = do
  proof <- get
  lift $ guard (proof.fuel > 0 && context.depth < 64)
  put (proof { fuel = proof.fuel - 1 })
  let child = context { depth = context.depth + 1 }
      go = evaluate child env
  case syntax of
    Typed Int inner -> case inner of
      NeutralExpr (Local ident level) -> case Map.lookup (Tuple ident level) env of
        Just value -> pure value
        Nothing -> do
          lift $ guard context.outer
          state <- get
          let ref = Tuple ident level
          put (state { inputs = Set.insert ref state.inputs })
          pure (Integer (Input ref))
      _ -> go inner
    Typed ty inner -> do
      lift $ guard (functionAnnotation ty)
      go inner
    Syn.TypeApp inner _ -> go inner
    Lit (LitInt value) -> do
      lift $ guard (value >= (-2147483647 - 1) && value <= 2147483647)
      pure (Integer (Constant value))
    Local ident level -> lift (Map.lookup (Tuple ident level) env)
    Var (Qualified owner name) -> do
      lift $ guard (owner == Nothing || owner == Just context.owner)
      if Set.member name context.producers then pure Producer
      else do
        lift $ guard (not (Set.member name context.active))
        body <- lift (Map.lookup name context.globals)
        evaluate (child { active = Set.insert name context.active, outer = false }) Map.empty body
    Abs refs body -> pure (Closure (NEA.toArray refs) body env context.active context.outer)
    Let ident level value body -> do
      evaluated <- go value
      evaluate child (Map.insert (Tuple ident level) evaluated env) body
    App fn args -> do
      callable <- go fn
      arguments <- traverse go (NEA.toArray args)
      apply child callable arguments
    PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpAdd) left right) -> do
      a <- go left >>= scalarValue
      b <- go right >>= scalarValue
      pure (Integer (add a b))
    PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpSubtract) left right) -> do
      a <- go left >>= scalarValue
      b <- go right >>= scalarValue
      case b of
        Constant value -> pure (Integer (add a (Constant (negate value))))
        _ -> lift Nothing
    _ -> lift Nothing

apply :: Context -> Value -> Array Value -> Eval Value
apply context value arguments = do
  lift $ guard (context.depth < 64)
  case value of
    Closure parameters body captured active outer -> do
      let count = min (Array.length parameters) (Array.length arguments)
          env = Map.union (Map.fromFoldable (Array.zip (Array.take count parameters) (Array.take count arguments))) captured
          pending = Array.drop count parameters
          supplied = Array.drop count arguments
      if not (Array.null pending) then pure (Closure pending body env active outer)
      else do
        result <- evaluate (context { active = active, outer = outer }) env body
        if Array.null supplied then pure result
        else apply (context { depth = context.depth + 1 }) result supplied
    _ -> case Array.uncons arguments of
      Nothing -> pure value
      Just { head, tail } -> do
        result <- applyOne context value head
        if Array.null tail then pure result
        else apply (context { depth = context.depth + 1 }) result tail

applyOne :: Context -> Value -> Value -> Eval Value
applyOne context value argument = case value of
  Producer -> do
    count <- scalarValue argument
    lift $ guard (closed count)
    proof <- get
    put (proof { counts = Array.snoc proof.counts count })
    pure (Iterator count)
  Iterator count -> pure (Repeated count argument)
  Repeated count callback -> do
    seed <- scalarValue argument
    proof <- get
    let parameter = proof.next
    put (proof { next = parameter + 1 })
    result <- apply (context { depth = context.depth + 1 }) callback [ Integer (Parameter parameter) ] >>= scalarValue
    offset <- lift (translation parameter result)
    current <- get
    put (current { used = true })
    pure (Integer (add seed (multiply count offset)))
  _ -> lift Nothing

scalarValue :: Value -> Eval Scalar
scalarValue (Integer scalar) = pure scalar
scalarValue _ = lift Nothing

-- A callback must be precisely a translation of its one argument. This rejects
-- multiplication, duplicated arguments, discarded arguments and opaque calls.
translation :: Int -> Scalar -> Maybe Scalar
translation parameter = case _ of
  Parameter found | found == parameter -> Just (Constant 0)
  Sum left right | closed right -> (\offset -> add offset right) <$> translation parameter left
  Sum left right | closed left -> (add left) <$> translation parameter right
  _ -> Nothing

closed :: Scalar -> Boolean
closed = case _ of
  Parameter _ -> false
  Sum a b -> closed a && closed b
  Product a b -> closed a && closed b
  _ -> true

add :: Scalar -> Scalar -> Scalar
add (Constant 0) right = right
add left (Constant 0) = left
add (Constant left) (Constant right) = Constant (left + right)
add left right = Sum left right

multiply :: Scalar -> Scalar -> Scalar
multiply (Constant 0) _ = Constant 0
multiply _ (Constant 0) = Constant 0
multiply (Constant 1) right = right
multiply left (Constant 1) = left
multiply left right = Product left right

-- Repeated IntAdd is exact modulo 2^32. Prelude IntMul instead rounds a JS
-- Number product before narrowing, so it is not a valid replacement at large
-- inputs. Splitting one signed operand into 16-bit pieces keeps both products
-- exact in Number arithmetic, including overflow and negative increments.
emitScalar :: Scalar -> State Int NeutralExpr
emitScalar = case _ of
  Constant value -> pure (integer value)
  Input ref -> pure (reference ref)
  Parameter _ -> pure (NeutralExpr (Fail "unbound counted-function parameter"))
  Sum left right -> binary Int (Syn.OpIntNum Syn.OpAdd) <$> emitScalar left <*> emitScalar right
  product@(Product _ _) ->
    let
      factors = flattenProduct product
    in case Array.uncons factors of
      Just { head, tail } | Array.length tail >= 2 && Array.all (_ == head) tail ->
        emitPower head (Array.length factors)
      _ -> case product of
        Product left right -> do
          a <- emitScalar left
          b <- emitScalar right
          emitModularProduct a b
        _ -> pure (integer 1)

flattenProduct :: Scalar -> Array Scalar
flattenProduct = case _ of
  Product left right -> flattenProduct left <> flattenProduct right
  other -> [ other ]

emitPower :: Scalar -> Int -> State Int NeutralExpr
emitPower base exponent
  | exponent <= 1 = emitScalar base
  | otherwise = do
      xExpr <- emitScalar base
      case base of
        Input _ -> buildPow xExpr exponent
        Constant _ -> buildPow xExpr exponent
        _ -> do
          lvl <- get
          put (lvl + 1)
          let ref = Tuple (Just (Ident "countedBase")) (Level lvl)
          body <- buildPow (reference ref) exponent
          pure (typed Int (NeutralExpr (Let (Just (Ident "countedBase")) (Level lvl) xExpr body)))
  where
  buildPow x n
    | n == 1 = pure x
    | n == 2 = emitModularProduct x x
    | mod n 2 == 0 = do
        half <- buildPow x (div n 2)
        lvl <- get
        put (lvl + 1)
        let ref = Tuple (Just (Ident "countedPow")) (Level lvl)
        sq <- emitModularProduct (reference ref) (reference ref)
        pure (typed Int (NeutralExpr (Let (Just (Ident "countedPow")) (Level lvl) half sq)))
    | otherwise = do
        evenPart <- buildPow x (n - 1)
        emitModularProduct evenPart x

emitModularProduct :: NeutralExpr -> NeutralExpr -> State Int NeutralExpr
emitModularProduct a b = do
  next <- get
  put (next + 2)
  let
    first = Tuple (Just (Ident "countedFactor")) (Level next)
    second = Tuple (Just (Ident "countedIncrement")) (Level (next + 1))
    x = reference first
    -- Inputs have signed Int32 guards, constants are checked during the
    -- proof, and sums/products narrow. Expose that fact to scalar emission
    -- without changing the arithmetic or relying on local binding names.
    y = binary Int Syn.OpIntBitOr (reference second) (integer 0)
    low = binary Int (Syn.OpIntNum Syn.OpMultiply) (binary Int Syn.OpIntBitAnd x (integer 65535)) y
    high = binary Int Syn.OpIntBitShiftLeft
      (binary Int (Syn.OpIntNum Syn.OpMultiply) (binary Int Syn.OpIntBitShiftRight x (integer 16)) y) (integer 16)
    result = binary Int (Syn.OpIntNum Syn.OpAdd) low high
  pure (typed Int (NeutralExpr (Let (Just (Ident "countedFactor")) (Level next) a
    (typed Int (NeutralExpr (Let (Just (Ident "countedIncrement")) (Level (next + 1)) b result))))))

-- Representation annotations must not turn a Number or another primitive
-- into a proof about Int32 arithmetic. Function/newtype-polymorphic wrappers
-- remain admissible; their bodies must still use the supported integer syntax.
functionAnnotation :: ExprType -> Boolean
functionAnnotation = case _ of
  Func _ _ -> true
  ForAll _ inner -> functionAnnotation inner
  TypeVar _ -> true
  Any -> true
  _ -> false

typed :: ExprType -> NeutralExpr -> NeutralExpr
typed ty = NeutralExpr <<< Typed ty

reference :: Ref -> NeutralExpr
reference (Tuple name level) = typed Int (NeutralExpr (Local name level))

integer :: Int -> NeutralExpr
integer value = typed Int (NeutralExpr (Lit (LitInt value)))

binary :: ExprType -> Syn.BackendOperator2 -> NeutralExpr -> NeutralExpr -> NeutralExpr
binary ty operator left right = typed ty (NeutralExpr (PrimOp (Syn.Op2 operator left right)))

compare :: Syn.BackendOperatorOrd -> NeutralExpr -> NeutralExpr -> NeutralExpr
compare operator = binary Boolean (Syn.OpIntOrd operator)
