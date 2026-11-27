-- Fuse a closed range / head-dependent exclusion / sum pipeline. The proof
-- describes its complete recurrences; no function or constructor spelling is
-- significant. The private bitmap represents the surviving ordered interval.
module Gopurs.ListProgressionFold (prepare) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (State, put, runState)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst, snd)
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), rawGo, sanitizeName)
import Gopurs.GoSymbol (GoReference(..), GoSymbol, generated)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendSyntax(..), Level, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type Ref = Tuple (Maybe Ident) Level
type Shape = { ty :: ExprType, empty :: Qualified Ident, cons :: Qualified Ident }
type Lambda = { args :: Array Ref, body :: NeutralExpr }
type Loop = { self :: Ref, args :: Array Ref, body :: NeutralExpr, initial :: Array NeutralExpr }
type Context =
  { module_ :: BackendModule
  , sieves :: Map.Map Ident Shape
  , locals :: Map.Map Ref NeutralExpr
  , helper :: Ident
  }

prepare :: BackendModule -> { module :: BackendModule, declarations :: Array GoDecl, functions :: Map.Map String FunctionInfo }
prepare mod =
  let bindings = Array.concatMap _.bindings mod.bindings
      filters = Map.fromFoldable $ Array.mapMaybe (\(Tuple name value) ->
        map (Tuple name) (recognizeFilter mod value)) bindings
      sieves = Map.fromFoldable $ Array.mapMaybe (\(Tuple name value) ->
        map (Tuple name) (recognizeSieve mod filters name value)) bindings
  in if Map.isEmpty sieves then { module: mod, declarations: [], functions: Map.empty }
     else prepareRecognized mod sieves

prepareRecognized :: BackendModule -> Map.Map Ident Shape -> { module :: BackendModule, declarations :: Array GoDecl, functions :: Map.Map String FunctionInfo }
prepareRecognized mod sieves =
  let helper = reserveName mod
      context = { module_: mod, sieves, locals: Map.empty, helper }
      Tuple groups changed = runState (traverse (\group -> do
        rewritten <- traverse (\(Tuple name value) -> Tuple name <$> rewrite context value) group.bindings
        pure (group { bindings = rewritten })) mod.bindings) false
      owner = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
      symbol = generated owner ("Call_" <> owner <> "_" <> sanitizeName (unwrap helper))
  in { module: mod { bindings = groups }
     , declarations: if changed then [ declaration symbol ] else []
     , functions: if changed then Map.singleton (sanitizeName (unwrap helper))
         { arrayFold: Nothing, codePointFold: Nothing, reference: SymbolReference symbol, fArgs: [ TypeInt64, TypeInt64, TypeInt64 ], fRet: TypeInt64, arity: 3 }
       else Map.empty
     }

reserveName :: BackendModule -> Ident
reserveName mod = choose 0
  where
  reserved = Set.fromFoldable $ map (sanitizeName <<< unwrap)
    (map fst (Array.concatMap _.bindings mod.bindings) <> Array.fromFoldable (Map.keys mod.foreign))
  choose index =
    let name = "__gopurs_list_progression_fold_" <> show index
    in if Set.member (sanitizeName name) reserved then choose (index + 1) else Ident name

-- Bounds are guarded in the source IR. Wider Go intermediates make 2*p and
-- the final increment safe even near MaxInt32. The sum retains Int32 wrapping.
-- A prefix starting at 2 is a prime sieve, so its odd bitmap can start marking
-- at p*p. Other lower bounds may select composite heads and retain 2*p below.
declaration :: GoSymbol -> GoDecl
declaration symbol = GoFunctionDecl
  { name: symbol
  , params: [ Tuple "lower" TypeInt64, Tuple "upper" TypeInt64, Tuple "seed" TypeInt64 ]
  , result: TypeInt64
  , body: GoBlock
      [ GoIfElse (GoBinOp "==" (GoVar "lower") (rawGo "int64(2)")) primePrefix []
      , GoAssign "rejected" (rawGo "make([]bool, upper-lower+1)")
      , GoAssign "total" (GoVar "seed")
      , GoForRange "value := lower; value <= upper; value++"
          [ GoIfElse (GoIndex (GoVar "rejected") (GoBinOp "-" (GoVar "value") (GoVar "lower")))
              [ GoContinue "" ] []
          , GoMutate "total" (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoBinOp "+" (GoVar "total") (GoVar "value") ] ])
          , GoForRange "multiple := value + value; multiple <= upper; multiple += value"
              [ GoMutate "rejected[multiple-lower]" (GoVar "true") ]
          ]
      , GoReturn (GoVar "total")
      ]
  }

-- The existing admission guard bounds this prefix to at most 1,048,577.
-- Its positive sum and Int32 seed therefore fit Int64 before the final wrap.
-- Even values are excluded by 2; entry i represents the odd value 2*i+1.
-- The local scratch buffer is per invocation and never escapes or aliases.
primePrefix :: Array GoExpr
primePrefix =
  [ rawGo "var local [512]bool"
  , rawGo "var rejected []bool"
  , GoAssign "size" (rawGo "(upper+1)/2")
  , GoIfElse (rawGo "size <= int64(len(local))")
      [ GoMutate "rejected" (rawGo "local[:size]") ]
      [ GoMutate "rejected" (rawGo "make([]bool, size)") ]
  , GoForRange "prime := int64(3); prime*prime <= upper; prime += 2"
      [ GoIfElse (rawGo "!rejected[prime/2]")
          [ GoForRange "multiple := prime*prime; multiple <= upper; multiple += prime+prime"
              [ GoMutate "rejected[multiple/2]" (GoVar "true") ]
          ] []
      ]
  , GoAssign "total" (rawGo "seed+2")
  , GoForRange "value := int64(3); value <= upper; value += 2"
      [ GoIfElse (rawGo "!rejected[value/2]")
          [ GoMutate "total" (GoBinOp "+" (GoVar "total") (GoVar "value")) ] []
      ]
  , GoReturn (rawGo "int64(int32(total))")
  ]

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

signature :: NeutralExpr -> Maybe { args :: Array ExprType, result :: ExprType }
signature value = flatten <$> annotation value
  where
  flatten (Func args result) = let rest = flatten result in { args: args <> rest.args, result: rest.result }
  flatten result = { args: [], result }

abstractions :: NeutralExpr -> Lambda
abstractions value = case strip value of
  Abs args body -> let rest = abstractions body in { args: NEA.toArray args <> rest.args, body: rest.body }
  _ -> { args: [], body: value }

application :: NeutralExpr -> { head :: NeutralExpr, args :: Array NeutralExpr }
application value = case strip value of
  App fn args -> let prior = application fn in { head: prior.head, args: prior.args <> NEA.toArray args }
  _ -> { head: value, args: [] }

localRef :: NeutralExpr -> Maybe Ref
localRef value = case strip value of
  Local name level -> Just (Tuple name level)
  _ -> Nothing

isLocal :: Ref -> NeutralExpr -> Boolean
isLocal ref value = localRef value == Just ref

global :: ModuleName -> NeutralExpr -> Maybe Ident
global current value = case strip value of
  Var (Qualified owner name) | owner == Just current || owner == Nothing -> Just name
  _ -> Nothing

single :: forall a. Array a -> Maybe a
single [ a ] = Just a
single _ = Nothing

two :: forall a. Array a -> Maybe (Tuple a a)
two [ a, b ] = Just (Tuple a b)
two _ = Nothing

literal :: Int -> NeutralExpr -> Boolean
literal expected value = case strip value of
  Lit (LitInt found) -> expected == found
  _ -> false

scalarValue :: NeutralExpr -> Boolean
scalarValue value = annotation value == Just Int && case strip value of
  Local _ _ -> true
  Lit (LitInt _) -> true
  _ -> false

listShape :: BackendModule -> ExprType -> Maybe Shape
listShape mod ty = do
  info <- case ty of
    ADT full _ args -> Just { full, args }
    _ -> Nothing
  decl <- Array.find (\decl -> unwrap mod.name <> "." <> decl.name == info.full) mod.dataDecls
  guard (Array.length decl.vars == Array.length info.args && all (_ == Int) info.args)
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

listCases :: Shape -> Ref -> NeutralExpr -> Maybe { empty :: NeutralExpr, cons :: NeutralExpr }
listCases shape ref value = case strip value of
  Branch cases failure -> case NEA.toArray cases, strip failure of
    [ Pair nilTest nilBody, Pair consTest consBody ], Fail _
      | tag shape.empty nilTest && tag shape.cons consTest -> Just { empty: nilBody, cons: consBody }
    _, _ -> Nothing
  _ -> Nothing
  where
  tag expected input = case strip input of
    PrimOp (Op1 (OpIsTag found) argument) -> expected == found && isLocal ref argument
    _ -> false

localLoop :: Array ExprType -> ExprType -> NeutralExpr -> Maybe Loop
localLoop args result value = case strip value of
  LetRec level bindings body -> do
    let entry = application body
    self <- localRef entry.head
    selected <- snd <$> Array.find (\(Tuple name _) -> self == Tuple (Just name) level) (NEA.toArray bindings)
    guard (all (not <<< Array.null <<< _.args <<< abstractions <<< snd) bindings)
    found <- signature selected
    guard (found.args == args && found.result == result)
    let lambda = abstractions selected
    guard (Array.length lambda.args == Array.length args)
    pure { self, args: lambda.args, body: lambda.body, initial: entry.args }
  _ -> Nothing

selfCall :: Ref -> NeutralExpr -> Maybe (Array NeutralExpr)
selfCall self value = do
  let call = application value
  guard (isLocal self call.head)
  pure call.args

reverseOf :: Shape -> Ref -> NeutralExpr -> Boolean
reverseOf shape source value = fromMaybe false do
  loop <- localLoop [ shape.ty, shape.ty ] shape.ty value
  Tuple input acc <- two loop.args
  Tuple initial seed <- two loop.initial
  guard (isLocal source initial && emptyList shape seed)
  cases <- listCases shape input loop.body
  guard (isLocal acc cases.empty)
  Tuple tail constructed <- selfCall loop.self cases.cons >>= two
  Tuple head rest <- consFields shape constructed
  pure (field shape input 1 tail && field shape input 0 head && isLocal acc rest)

recognizeFilter :: BackendModule -> NeutralExpr -> Maybe Shape
recognizeFilter mod value = do
  sig <- signature value
  ty <- case sig.args of
    [ Func [ Int ] Boolean, ty ] | ty == sig.result -> Just ty
    _ -> Nothing
  shape <- listShape mod ty
  let lambda = abstractions value
  Tuple predicate source <- two lambda.args
  loop <- localLoop [ ty, ty ] ty lambda.body
  Tuple input acc <- two loop.args
  Tuple initial seed <- two loop.initial
  guard (isLocal source initial && emptyList shape seed)
  cases <- listCases shape input loop.body
  guard (reverseOf shape acc cases.empty)
  branch <- case strip cases.cons of
    Branch choices no -> case NEA.toArray choices of
      [ Pair condition yes ] -> Just { condition, yes, no }
      _ -> Nothing
    _ -> Nothing
  let called = application branch.condition
  argument <- single called.args
  guard (isLocal predicate called.head && field shape input 0 argument)
  Tuple yesTail yesAcc <- selfCall loop.self branch.yes >>= two
  Tuple noTail noAcc <- selfCall loop.self branch.no >>= two
  Tuple head rest <- consFields shape yesAcc
  guard (field shape input 1 yesTail && field shape input 1 noTail)
  guard (field shape input 0 head && isLocal acc rest && isLocal acc noAcc)
  pure shape

recognizeSieve :: BackendModule -> Map.Map Ident Shape -> Ident -> NeutralExpr -> Maybe Shape
recognizeSieve mod filters name value = do
  sig <- signature value
  ty <- single sig.args
  guard (sig.result == ty)
  shape <- listShape mod ty
  let lambda = abstractions value
  input <- single lambda.args
  cases <- listCases shape input lambda.body
  guard (emptyList shape cases.empty)
  head <- case strip cases.cons of
    Let name_ level selected body | field shape input 0 selected -> Just { ref: Tuple name_ level, body }
    _ -> Nothing
  Tuple first rest <- consFields shape head.body
  guard (isLocal head.ref first)
  let recur = application rest
  guard (global mod.name recur.head == Just name)
  filtered <- single recur.args
  let call = application filtered
  filter <- global mod.name call.head >>= flip Map.lookup filters
  guard (filter == shape)
  Tuple predicate tail <- two call.args
  guard (field shape input 1 tail && excludesMultiples head.ref predicate)
  pure shape

-- The only discarded callback is this complete primitive expression. In the
-- guarded positive interval it is pure, total, and equivalent to marking the
-- arithmetic progression. Opaque calls and additional captures are rejected.
excludesMultiples :: Ref -> NeutralExpr -> Boolean
excludesMultiples head value = fromMaybe false do
  sig <- signature value
  guard (sig.args == [ Int ] && sig.result == Boolean)
  let lambda = abstractions value
  item <- single lambda.args
  case strip lambda.body of
    PrimOp (Op1 OpBooleanNot equal) -> case strip equal of
      PrimOp (Op2 (OpIntOrd OpEq) modulo zero) | literal 0 zero -> case strip modulo of
        PrimOp (Op2 (OpIntNum OpMod) x divisor) -> pure (isLocal item x && isLocal head divisor)
        _ -> Nothing
      _ -> Nothing
    _ -> Nothing

recognizeFolder :: BackendModule -> Ref -> NeutralExpr -> Maybe Shape
recognizeFolder mod self value = do
  sig <- signature value
  ty <- case sig.args, sig.result of
    [ ty, Int ], Int -> Just ty
    _, _ -> Nothing
  shape <- listShape mod ty
  let lambda = abstractions value
  Tuple input acc <- two lambda.args
  cases <- listCases shape input lambda.body
  guard (isLocal acc cases.empty)
  Tuple tail total <- selfCall self cases.cons >>= two
  guard (field shape input 1 tail)
  guard case strip total of
    PrimOp (Op2 (OpIntNum OpAdd) old item) -> isLocal acc old && field shape input 0 item
    _ -> false
  pure shape

rangeBounds :: Shape -> NeutralExpr -> Maybe { lower :: NeutralExpr, upper :: NeutralExpr }
rangeBounds shape value = do
  loop <- localLoop [ Int, shape.ty ] shape.ty value
  Tuple counter acc <- two loop.args
  Tuple upper seed <- two loop.initial
  guard (scalarValue upper && emptyList shape seed)
  branch <- case strip loop.body of
    Branch choices step -> case NEA.toArray choices of
      [ Pair condition result ] | isLocal acc result -> Just { condition, step }
      _ -> Nothing
    _ -> Nothing
  lower <- case strip branch.condition of
    PrimOp (Op2 (OpIntOrd OpLt) current lower) | isLocal counter current && scalarValue lower -> Just lower
    _ -> Nothing
  -- A bound referencing either loop argument is not invariant or available
  -- at its call site. Only already evaluated enclosing values may be captured.
  guard (not (isLocal counter lower || isLocal acc lower || isLocal loop.self lower))
  Tuple next built <- selfCall loop.self branch.step >>= two
  guard case strip next of
    PrimOp (Op2 (OpIntNum OpSubtract) current amount) -> isLocal counter current && literal 1 amount
    _ -> false
  Tuple head tail <- consFields shape built
  guard (isLocal counter head && isLocal acc tail)
  pure { lower, upper }

rewrite :: Context -> NeutralExpr -> State Boolean NeutralExpr
rewrite context original@(NeutralExpr syntax) = case fuse context original of
  Just replacement -> put true *> pure replacement
  Nothing -> case syntax of
    LetRec level bindings body -> do
      let locals = Map.union (Map.fromFoldable (map (\(Tuple name value) -> Tuple (Tuple (Just name) level) value) (NEA.toArray bindings))) context.locals
          nested = context { locals = locals }
      rewritten <- traverse (\(Tuple name value) -> Tuple name <$> rewrite nested value) bindings
      rest <- rewrite nested body
      pure (NeutralExpr (LetRec level rewritten rest))
    Abs args body ->
      let locals = foldl (flip Map.delete) context.locals args
      in NeutralExpr <<< Abs args <$> rewrite (context { locals = locals }) body
    UncurriedAbs args body ->
      let locals = foldl (flip Map.delete) context.locals args
      in NeutralExpr <<< UncurriedAbs args <$> rewrite (context { locals = locals }) body
    UncurriedEffectAbs args body ->
      let locals = foldl (flip Map.delete) context.locals args
      in NeutralExpr <<< UncurriedEffectAbs args <$> rewrite (context { locals = locals }) body
    Let name level value body -> do
      next <- rewrite context value
      rest <- rewrite (context { locals = Map.delete (Tuple name level) context.locals }) body
      pure (NeutralExpr (Let name level next rest))
    EffectBind name level value body -> do
      next <- rewrite context value
      rest <- rewrite (context { locals = Map.delete (Tuple name level) context.locals }) body
      pure (NeutralExpr (EffectBind name level next rest))
    _ -> NeutralExpr <$> traverse (rewrite context) syntax

fuse :: Context -> NeutralExpr -> Maybe NeutralExpr
fuse context original = do
  let call = application original
  self <- localRef call.head
  folder <- Map.lookup self context.locals
  shape <- recognizeFolder context.module_ self folder
  Tuple source seed <- two call.args
  guard (scalarValue seed)
  let recur = application source
  selected <- global context.module_.name recur.head >>= flip Map.lookup context.sieves
  guard (selected == shape)
  produced <- single recur.args
  bounds <- rangeBounds shape produced
  let lower = bounds.lower
      upper = bounds.upper
      integer n = typed Int (NeutralExpr (Lit (LitInt n)))
      binary ty op a b = typed ty (NeutralExpr (PrimOp (Op2 op a b)))
      compare op = binary Boolean (OpIntOrd op)
      both = binary Boolean OpBooleanAnd
      tests =
        [ compare OpGte lower (integer 2)
        , compare OpGte upper lower
        , compare OpLte upper (integer 2147483647)
        , compare OpLte (binary Int (OpIntNum OpSubtract) upper lower) (integer 1048575)
        , compare OpGte seed (integer (-2147483647 - 1))
        , compare OpLte seed (integer 2147483647)
        ]
      condition = Array.foldr both (typed Boolean (NeutralExpr (Lit (LitBoolean true)))) tests
      worker = typed (Func [ Int, Int, Int ] Int) (NeutralExpr (Var (Qualified (Just context.module_.name) context.helper)))
      fast = typed Int (NeutralExpr (App worker (NEA.cons' lower [ upper, seed ])))
  -- Keep the original application and its scopes verbatim in the fallback.
  -- All duplicated operands are already evaluated locals or literal values.
  pure (typed Int (NeutralExpr (Branch (NEA.singleton (Pair condition fast)) original)))

typed :: ExprType -> NeutralExpr -> NeutralExpr
typed ty = NeutralExpr <<< Typed ty
