module Gopurs.RefLoopRegions (optimize) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Gopurs.ImmediateApplications.Scope (occurrences, substitute)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(ADT, Int, Boolean), Ident(..), Literal(..), ModuleName(..), Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendEffect(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendOperator(..), BackendOperator2(..), BackendSyntax(..), Level, Pair(..))

data Kind = Increment NeutralExpr | LastWrite Level
type Plan = { parameters :: Array Level, reference :: Level, seed :: Level, kind :: Kind, st :: Boolean }
type Context = { owner :: ModuleName, plans :: Map.Map Ident Plan }

-- Only a fresh, nonescaping reference with a completely known lifetime may
-- lose intermediate writes. Public helpers and shared references stay intact.
-- Usage annotations are not ownership proofs: admission accounts for every
-- use through the exact allocation -> helper -> read continuation instead.
optimize :: BackendModule -> BackendModule
optimize mod =
  let
    plans = Map.fromFoldable (Array.concatMap (\group -> if group.recursive then [] else
      Array.mapMaybe (\(Tuple name value) -> Tuple name <$> helper value) group.bindings) mod.bindings)
    context = { owner: mod.name, plans }
    rewriteGroup group = if group.recursive then group else group
      { bindings = map (\(Tuple name value) -> Tuple name (rewrite context value)) group.bindings }
  in if Map.isEmpty plans then mod else mod { bindings = map rewriteGroup mod.bindings }

rewrite :: Context -> NeutralExpr -> NeutralExpr
rewrite context original@(NeutralExpr syntax) = case reduce context original of
  Just result -> result
  Nothing -> case syntax of
    LetRec _ _ _ -> original
    _ -> NeutralExpr (map (rewrite context) syntax)

reduce :: Context -> NeutralExpr -> Maybe NeutralExpr
reduce context original = case call "Control.Monad.ST.Internal" "run" original of
  Just [ action ] -> region context true original action
  _ -> region context false original original

region :: Context -> Boolean -> NeutralExpr -> NeutralExpr -> Maybe NeutralExpr
region context runST original action = do
  allocated <- allocation action
  guard (allocated.st == runST)
  guard (scalar allocated.initial)
  case strip allocated.body of
    EffectBind _ discarded loop final -> do
      guard (occurrences discarded final == 0)
      guard (reads allocated.st allocated.reference final)
      let Tuple fn args = spine loop
      name <- case strip fn of
        Var (Qualified (Just owner) name) | owner == context.owner -> Just name
        _ -> Nothing
      plan <- Map.lookup name context.plans
      guard (plan.st == allocated.st && Array.length args == Array.length plan.parameters)
      guard (all identity (Array.zipWith (\parameter arg ->
        if parameter == plan.reference then local allocated.reference arg
        else scalar arg && occurrences allocated.reference arg == 0) plan.parameters args))
      let
        actual = replaceParameters (Map.fromFoldable (Array.zip plan.parameters args))
        parameter level = actual (NeutralExpr (Local Nothing level))
        seed = parameter plan.seed
        pureResult value = if runST then value else NeutralExpr (EffectPure value)
        initial = allocated.initial
      pure case plan.kind of
        Increment amount ->
          let delta = actual amount
              stepSum = case strip delta of
                Lit (LitInt 0) -> num OpAdd initial (integer 0)
                Lit (LitInt 1) -> num OpAdd initial seed
                Lit (LitInt (-1)) -> num OpSubtract initial seed
                _ -> num OpAdd initial (num OpMultiply seed delta)
          in pureResult (choose (ord OpGt seed (integer 0)) stepSum initial)
        LastWrite limit ->
          let bound = parameter limit
              result = pureResult (choose (ord OpGt seed bound) initial bound)
              regionType = if runST then Int else ADT "Effect.Effect" [ "Effect", "Effect" ] [ Int ]
          -- Inclusive upward Int32 loops do not terminate at maxBound: keep
          -- the original allocation/action and all its looping behavior there.
          in NeutralExpr (Typed regionType (choose (ord OpLt bound (integer 2147483647)) result original))
    _ -> Nothing

type Allocation = { st :: Boolean, reference :: Level, initial :: NeutralExpr, body :: NeutralExpr }

allocation :: NeutralExpr -> Maybe Allocation
allocation original = case strip original of
  EffectBind _ reference value body -> do
    Tuple st initial <- newRef value
    pure { st, reference, initial, body }
  Let _ action value body -> case strip body of
    EffectBind _ reference invoked rest | local action invoked && occurrences action rest == 0 -> do
      Tuple st initial <- newRef value
      pure { st, reference, initial, body: rest }
    _ -> Nothing
  _ -> Nothing

newRef :: NeutralExpr -> Maybe (Tuple Boolean NeutralExpr)
newRef expression = case strip expression of
  PrimEffect (EffectRefNew initial) -> Just (Tuple true initial)
  _ -> case call "Effect.Ref" "_new" expression of
    Just [ initial ] -> Just (Tuple false initial)
    _ -> case call "Control.Monad.ST.Internal" "newImpl" expression of
      Just [ initial ] -> Just (Tuple true initial)
      _ -> Nothing

reads :: Boolean -> Level -> NeutralExpr -> Boolean
reads st reference expression = case strip expression of
  PrimEffect (EffectRefRead ref) -> st && local reference ref
  _ -> case call (if st then "Control.Monad.ST.Internal" else "Effect.Ref") "read" expression of
    Just [ ref ] -> local reference ref
    _ -> false

helper :: NeutralExpr -> Maybe Plan
helper expression = do
  let Tuple parameters body = lambdas expression
      Tuple fn args = spine body
  guard (Array.length parameters >= 2)
  guard (Array.length (Array.nub parameters) == Array.length parameters)
  Tuple st supplied <- case strip fn of
    Accessor dictionary (GetProp "tailRecM") -> Tuple <$> instanceKind dictionary <*> pure args
    _ | named "Control.Monad.Rec.Class" "tailRecM" fn -> case Array.uncons args of
      Just { head: dictionary, tail } -> Tuple <$> instanceKind dictionary <*> pure tail
      _ -> Nothing
    _ -> Nothing
  case supplied of
    [ callback, initial ] -> do
      seed <- localLevel initial
      guard (Array.elem seed parameters)
      let Tuple stateParameters step = lambdas callback
      state <- case stateParameters of
        [ state ] -> Just state
        _ -> Nothing
      guard (not (Array.elem state parameters))
      case strip step of
        Branch branches repeated -> case NEA.toArray branches of
          [ Pair condition done ] -> do
            guard (doneUnit done)
            Tuple mutation next <- iteration repeated
            case strip condition of
              PrimOp (Op2 (OpIntOrd OpLte) left right) | local state left && literal 0 right -> do
                guard (advance OpSubtract state next)
                Tuple reference amount <- increment st mutation
                guard (Array.elem reference parameters && reference /= seed)
                guard (invariant (Array.filter (_ /= reference) parameters) amount)
                pure { parameters, reference, seed, kind: Increment amount, st }
              PrimOp (Op2 (OpIntOrd OpGt) left right) | local state left -> do
                limit <- localLevel right
                guard (Array.elem limit parameters && advance OpAdd state next)
                Tuple reference value <- writeRef st mutation
                guard (Array.elem reference parameters && reference /= seed && reference /= limit && local state value)
                pure { parameters, reference, seed, kind: LastWrite limit, st }
              _ -> Nothing
          _ -> Nothing
        _ -> Nothing
    _ -> Nothing

instanceKind :: NeutralExpr -> Maybe Boolean
instanceKind value
  | named "Control.Monad.ST.Internal" "monadRecST" value = Just true
  | named "Control.Monad.Rec.Class" "monadRecEffect" value = Just false
  | otherwise = Nothing

iteration :: NeutralExpr -> Maybe (Tuple NeutralExpr NeutralExpr)
iteration expression = case strip expression of
  EffectDefer inner -> iteration inner
  Let _ level value body -> case strip body of
    EffectBind _ discarded invoked rest | local level invoked && occurrences level rest == 0 -> finish discarded value rest
    _ -> Nothing
  EffectBind _ discarded value rest -> finish discarded value rest
  _ -> Nothing
  where
  finish discarded value rest = do
    next <- stepValue "Loop" rest
    guard (occurrences discarded next == 0)
    pure (Tuple value next)

stepValue :: String -> NeutralExpr -> Maybe NeutralExpr
stepValue name expression = case strip expression of
  EffectPure value -> case strip value of
    CtorSaturated (Qualified (Just (ModuleName owner)) _) _ _ (Ident ctor) [ Tuple _ result ]
      | owner == "Control.Monad.Rec.Class" && ctor == name -> Just result
    _ -> Nothing
  _ -> Nothing

doneUnit :: NeutralExpr -> Boolean
doneUnit expression = case stepValue "Done" expression of
  Just value -> named "Data.Unit" "unit" value
  Nothing -> false

advance :: BackendOperatorNum -> Level -> NeutralExpr -> Boolean
advance operator state expression = case strip expression of
  PrimOp (Op2 (OpIntNum found) left right) -> found == operator && local state left && literal 1 right
  _ -> false

writeRef :: Boolean -> NeutralExpr -> Maybe (Tuple Level NeutralExpr)
writeRef st expression = case strip expression of
  PrimEffect (EffectRefWrite ref value) | st -> flip Tuple value <$> localLevel ref
  _ -> case call (if st then "Control.Monad.ST.Internal" else "Effect.Ref") "write" expression of
    Just [ value, ref ] -> flip Tuple value <$> localLevel ref
    _ -> Nothing

increment :: Boolean -> NeutralExpr -> Maybe (Tuple Level NeutralExpr)
increment true expression = case strip expression of
  EffectBind _ previous source update -> do
    reference <- case strip source of
      PrimEffect (EffectRefRead ref) -> localLevel ref
      _ -> Nothing
    Tuple written value <- writeRef true update
    guard (written == reference)
    Tuple state amount <- addition value
    guard (local previous state && occurrences previous amount == 0)
    pure (Tuple reference amount)
  _ -> Nothing
increment false expression = case call "Effect.Ref" "modifyImpl" expression of
  Just [ callback, ref ] -> do
    reference <- localLevel ref
    let Tuple params body = lambdas callback
    previous <- case params of
      [ parameter ] -> Just parameter
      _ -> Nothing
    record <- recordBody body
    value <- Array.findMap (\(Prop key value) -> if key == "state" then Just value else Nothing) record
    guard (Array.length record == 2 && all (\(Prop key term) -> (key == "state" || key == "value") && safeInteger term) record)
    Tuple state amount <- addition value
    guard (local previous state && occurrences previous amount == 0)
    pure (Tuple reference amount)
  _ -> Nothing

recordBody :: NeutralExpr -> Maybe (Array (Prop NeutralExpr))
recordBody expression = case strip expression of
  Lit (LitRecord fields) -> Just fields
  Let _ level value body | safeInteger value -> recordBody (substitute level value body)
  _ -> Nothing

addition :: NeutralExpr -> Maybe (Tuple NeutralExpr NeutralExpr)
addition expression = case strip expression of
  PrimOp (Op2 (OpIntNum OpAdd) left right) -> Just (Tuple left right)
  _ -> Nothing

-- Captured immutable integer operands cannot run effects or throw while being
-- moved into the formula. Reject function calls and non-integer representations.
scalar :: NeutralExpr -> Boolean
scalar (NeutralExpr syntax) = case syntax of
  Typed Int value -> case strip value of
    Local _ _ -> true
    Lit (LitInt _) -> true
    _ -> false
  Typed _ _ -> false
  TypeApp value _ -> scalar value
  Lit (LitInt _) -> true
  _ -> false

invariant :: Array Level -> NeutralExpr -> Boolean
invariant parameters expression = case strip expression of
  Local _ level -> Array.elem level parameters
  Lit (LitInt _) -> true
  PrimOp (Op2 (OpIntNum operator) left right) | operator == OpAdd || operator == OpSubtract || operator == OpMultiply ->
    invariant parameters left && invariant parameters right
  _ -> false

safeInteger :: NeutralExpr -> Boolean
safeInteger expression = case strip expression of
  Local _ _ -> true
  Lit (LitInt _) -> true
  PrimOp (Op2 (OpIntNum operator) left right) | operator == OpAdd || operator == OpSubtract || operator == OpMultiply -> safeInteger left && safeInteger right
  _ -> false

-- Simultaneous substitution: caller levels may numerically overlap callee
-- parameters. Never substitute into an argument inserted by another parameter.
replaceParameters :: Map.Map Level NeutralExpr -> NeutralExpr -> NeutralExpr
replaceParameters parameters original@(NeutralExpr syntax) = case syntax of
  Local _ level -> case Map.lookup level parameters of
    Just value -> value
    Nothing -> original
  _ -> NeutralExpr (map (replaceParameters parameters) syntax)

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syntax) = case syntax of
  Typed _ value -> strip value
  TypeApp value _ -> strip value
  _ -> syntax

lambdas :: NeutralExpr -> Tuple (Array Level) NeutralExpr
lambdas expression = case strip expression of
  Abs params body -> let Tuple rest result = lambdas body
    in Tuple (map (\(Tuple _ level) -> level) (NEA.toArray params) <> rest) result
  _ -> Tuple [] expression

spine :: NeutralExpr -> Tuple NeutralExpr (Array NeutralExpr)
spine expression = case strip expression of
  App fn args -> let Tuple target before = spine fn in Tuple target (before <> NEA.toArray args)
  _ -> Tuple expression []

call :: String -> String -> NeutralExpr -> Maybe (Array NeutralExpr)
call owner name expression = let Tuple fn args = spine expression
  in if named owner name fn then Just args else Nothing

named :: String -> String -> NeutralExpr -> Boolean
named owner name expression = case strip expression of
  Var (Qualified (Just (ModuleName foundOwner)) (Ident foundName)) -> foundOwner == owner && foundName == name
  _ -> false

localLevel :: NeutralExpr -> Maybe Level
localLevel expression = case strip expression of
  Local _ level -> Just level
  _ -> Nothing

local :: Level -> NeutralExpr -> Boolean
local level expression = localLevel expression == Just level

literal :: Int -> NeutralExpr -> Boolean
literal expected expression = case strip expression of
  Lit (LitInt value) -> value == expected
  _ -> false

integer :: Int -> NeutralExpr
integer value = NeutralExpr (Typed Int (NeutralExpr (Lit (LitInt value))))

num :: BackendOperatorNum -> NeutralExpr -> NeutralExpr -> NeutralExpr
num operator left right = NeutralExpr (Typed Int (NeutralExpr (PrimOp (Op2 (OpIntNum operator) left right))))

ord :: BackendOperatorOrd -> NeutralExpr -> NeutralExpr -> NeutralExpr
ord operator left right = NeutralExpr (Typed Boolean (NeutralExpr (PrimOp (Op2 (OpIntOrd operator) left right))))

choose :: NeutralExpr -> NeutralExpr -> NeutralExpr -> NeutralExpr
choose condition yes no = NeutralExpr (Branch (NEA.singleton (Pair condition yes)) no)
