module Gopurs.ClosedDictionaries (cacheClosedDictionaries) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (State, get, put, runState)
import Data.Array as Array
import Data.Array.NonEmpty (toArray)
import Data.Foldable (class Foldable, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (CodegenMetadata)
import PureScript.Backend.Optimizer.Convert (BackendBindingGroup, BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType, Ident(..), Literal(..), ModuleName, Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendSyntax(..), Level(..), Pair(..))

-- Class dictionaries and other closed instance values are rebuilt at every
-- evaluation of the expression that constructs them. When such an expression
-- is closed (it captures no local variable), the value is the same on every
-- evaluation, so it belongs in a module-level cached getter exactly like a
-- top-level binding. Lifting the outermost closed dictionary application turns
-- a per-record plan reconstruction into a one-time construction and lets the
-- plan actually be shared by every invocation of the dictionary, which is what
-- the native record decoder already assumes.
--
-- Conditions are deliberately narrow: the expression must be an application
-- (not a bare reference or literal), free of local variables, and its
-- annotation must be a class type from the metadata. Effects are out of scope
-- because a class dictionary value is never an effect computation.

type LiftState =
  { next :: Int
  , names :: Set.Set String
  , moduleName :: ModuleName
  , lifted :: Array (Tuple Ident NeutralExpr)
  }

cacheClosedDictionaries :: CodegenMetadata -> BackendModule -> BackendModule
cacheClosedDictionaries metadata mod =
  let
    existing = foldl
      (\acc (Tuple (Ident name) _) -> Set.insert name acc)
      Set.empty
      (Array.concatMap _.bindings mod.bindings)

    initial :: LiftState
    initial = { next: 0, names: existing, moduleName: mod.name, lifted: [] }

    Tuple rewritten liftedState = runState (traverse (rewriteGroup metadata) mod.bindings) initial
  in
    if Array.null liftedState.lifted then mod
    else mod
      { bindings = rewritten <> [ { recursive: false, bindings: liftedState.lifted } ]
      }

rewriteGroup :: CodegenMetadata -> BackendBindingGroup Ident NeutralExpr -> State LiftState (BackendBindingGroup Ident NeutralExpr)
rewriteGroup metadata group = do
  bindings <- traverse
    (\(Tuple ident expr) -> Tuple ident <$> rewriteNested metadata expr)
    group.bindings
  pure group { bindings = bindings }

-- A top-level binding body already lives in a cached getter, so it is not
-- lifted a second time; only constructions nested below it are.
rewriteNested :: CodegenMetadata -> NeutralExpr -> State LiftState NeutralExpr
rewriteNested metadata (NeutralExpr syn) =
  NeutralExpr <$> traverse (rewriteExpr metadata) syn

rewriteExpr :: CodegenMetadata -> NeutralExpr -> State LiftState NeutralExpr
rewriteExpr metadata expr@(NeutralExpr syn) =
  case liftable metadata expr of
    Just ty -> lift ty expr
    Nothing ->
      NeutralExpr <$> traverse (rewriteExpr metadata) syn

liftable :: CodegenMetadata -> NeutralExpr -> Maybe ExprType
liftable metadata expr = do
  guard (isApplication expr)
  guard (Set.isEmpty (freeVars expr))
  guard (not (containsEffect expr))
  ty <- annotatedType expr
  className <- dictionaryClass ty
  guard (Map.member className metadata.classDeclsFields)
  pure ty

-- Effect syntax is excluded: constructing an effect computation can register
-- runtime work, and sharing such a value is out of scope for this pass.
containsEffect :: NeutralExpr -> Boolean
containsEffect (NeutralExpr syn) = case syn of
  PrimEffect _ -> true
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> true
  UncurriedEffectApp _ _ -> true
  UncurriedEffectAbs _ _ -> true
  _ -> foldl (\found child -> found || containsEffect child) false syn

isApplication :: NeutralExpr -> Boolean
isApplication expr = case strip expr of
  App _ _ -> true
  UncurriedApp _ _ -> true
  _ -> false

annotatedType :: NeutralExpr -> Maybe ExprType
annotatedType (NeutralExpr syn) = case syn of
  Typed ty _ -> Just ty
  _ -> Nothing

dictionaryClass :: ExprType -> Maybe String
dictionaryClass = case _ of
  CoreFn.ADT name _ _ -> Just name
  CoreFn.ForAll _ ty -> dictionaryClass ty
  CoreFn.ConstrainedType _ ty -> dictionaryClass ty
  CoreFn.TypeApp ty _ -> dictionaryClass ty
  _ -> Nothing

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syn) = case syn of
  Typed _ inner -> strip inner
  TypeApp inner _ -> strip inner
  _ -> syn

lift :: ExprType -> NeutralExpr -> State LiftState NeutralExpr
lift ty expr = do
  st <- get
  let
    allocated = allocate st.names st.next
    ident = Ident allocated.name
  put st
    { next = allocated.next
    , names = Set.insert allocated.name st.names
    , lifted = Array.snoc st.lifted (Tuple ident expr)
    }
  pure (NeutralExpr (Typed ty (NeutralExpr (Var (Qualified (Just st.moduleName) ident)))))

allocate :: Set.Set String -> Int -> { name :: String, next :: Int }
allocate names index =
  let
    candidate = "__cached_dict_" <> show index
  in
    if Set.member candidate names then allocate names (index + 1)
    else { name: candidate, next: index + 1 }

freeVars :: NeutralExpr -> Set.Set String
freeVars (NeutralExpr syn) = case syn of
  Var _ -> Set.empty
  Local mbIdent lvl -> Set.singleton (localId mbIdent lvl)
  Lit lit -> case lit of
    LitArray arr -> foldl (\acc e -> Set.union acc (freeVars e)) Set.empty arr
    LitRecord rec -> foldl (\acc (Prop _ e) -> Set.union acc (freeVars e)) Set.empty rec
    _ -> Set.empty
  App fn args ->
    foldl (\acc e -> Set.union acc (freeVars e)) (freeVars fn) (toArray args)
  TypeApp fn _ -> freeVars fn
  Abs args body -> differenceBound args body
  UncurriedApp fn args ->
    foldl (\acc e -> Set.union acc (freeVars e)) (freeVars fn) args
  UncurriedAbs args body -> differenceBound args body
  UncurriedEffectApp fn args ->
    foldl (\acc e -> Set.union acc (freeVars e)) (freeVars fn) args
  UncurriedEffectAbs args body -> differenceBound args body
  Accessor e _ -> freeVars e
  Update e props ->
    foldl (\acc (Prop _ val) -> Set.union acc (freeVars val)) (freeVars e) props
  CtorSaturated _ _ _ _ args ->
    foldl (\acc (Tuple _ e) -> Set.union acc (freeVars e)) Set.empty args
  CtorDef _ _ _ _ -> Set.empty
  LetRec lvl binds body ->
    let
      bindsSet = foldl (\acc (Tuple ident _) -> Set.insert (localId (Just ident) lvl) acc) Set.empty (toArray binds)
      bodyVars = freeVars body
      bindsVars = foldl (\acc (Tuple _ e) -> Set.union acc (freeVars e)) Set.empty (toArray binds)
    in
      Set.difference (Set.union bodyVars bindsVars) bindsSet
  Let mbIdent lvl val body ->
    Set.union (freeVars val) (Set.difference (freeVars body) (Set.singleton (localId mbIdent lvl)))
  EffectBind mbIdent lvl val body ->
    Set.union (freeVars val) (Set.difference (freeVars body) (Set.singleton (localId mbIdent lvl)))
  EffectPure e -> freeVars e
  EffectDefer e -> freeVars e
  Branch pairs def ->
    let
      defVars = freeVars def
      pairsVars = foldl (\acc (Pair cond body) -> Set.union acc (Set.union (freeVars cond) (freeVars body))) Set.empty (toArray pairs)
    in
      Set.union defVars pairsVars
  PrimOp op -> case op of
    Op1 _ e -> freeVars e
    Op2 _ e1 e2 -> Set.union (freeVars e1) (freeVars e2)
  PrimEffect _ -> Set.empty
  PrimUndefined -> Set.empty
  Fail _ -> Set.empty
  Typed _ a -> freeVars a

differenceBound :: forall f. Foldable f => f (Tuple (Maybe Ident) Level) -> NeutralExpr -> Set.Set String
differenceBound args body =
  let
    argsSet = foldl (\acc (Tuple mbIdent lvl) -> Set.insert (localId mbIdent lvl) acc) Set.empty args
  in
    Set.difference (freeVars body) argsSet

localId :: Maybe Ident -> Level -> String
localId (Just (Ident i)) (Level l) = i <> "_" <> show l
localId Nothing (Level l) = "__local_var_" <> show l
