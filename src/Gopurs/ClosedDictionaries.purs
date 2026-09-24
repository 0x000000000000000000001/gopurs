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
import Data.String as String
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (CodegenMetadata)
import PureScript.Backend.Optimizer.Convert (BackendBindingGroup, BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType, Ident(..), Literal(..), ModuleName, Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendSyntax(..), Level(..), Pair(..))
import PureScript.Backend.Optimizer.Substitute (unify)
import PureScript.Backend.Optimizer.TypeSubstitution (substitute)

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
-- annotation (or a verified shared binding) must have a class type. Effects are out of scope
-- because a class dictionary value is never an effect computation.

type LiftState =
  { next :: Int
  , names :: Set.Set String
  , moduleName :: ModuleName
  , lifted :: Array (Tuple Ident NeutralExpr)
  , shared :: Map.Map ApplicationKey { ident :: Ident, type :: ExprType }
  }

-- PBO can inline an instance application's body while losing the result's
-- annotation. Reuse an existing typed top-level dictionary for an identical
-- application of globals. Keep curried/uncurried spines distinct and reject
-- locals, type applications and annotated children: erased syntax alone must
-- not conflate different instantiations or coercions. Shared candidates must
-- use imported globals only (no new local-module initialization dependency),
-- and their result type must be uniquely determined by monomorphic arguments.
data ApplicationKey
  = GlobalKey (Qualified Ident)
  | CurriedKey ApplicationKey (Array ApplicationKey)
  | UncurriedKey ApplicationKey (Array ApplicationKey)

derive instance eqApplicationKey :: Eq ApplicationKey
derive instance ordApplicationKey :: Ord ApplicationKey

applicationKey :: NeutralExpr -> Maybe ApplicationKey
applicationKey (NeutralExpr syn) = case syn of
  Var name -> Just (GlobalKey name)
  App fn args -> CurriedKey <$> applicationKey fn <*> traverse applicationKey (toArray args)
  UncurriedApp fn args -> UncurriedKey <$> applicationKey fn <*> traverse applicationKey args
  _ -> Nothing

sharedDictionaries :: CodegenMetadata -> BackendModule -> Map.Map ApplicationKey { ident :: Ident, type :: ExprType }
sharedDictionaries metadata mod = foldl addGroup Map.empty mod.bindings
  where
  addGroup entries group
    | group.recursive = entries
    | otherwise = foldl addBinding entries group.bindings

  addBinding entries (Tuple ident expr) = case expr of
    NeutralExpr (Typed ty body) -> case dictionaryClass ty, applicationKey body of
      Just name, Just key | Map.member name metadata.classDeclsFields && isApplication body
        && monomorphic ty && importedApplicationType metadata mod.name key == Just ty ->
        -- Preserve the first source binding when equivalent names exist.
        if Map.member key entries then entries
        else Map.insert key { ident, type: ty } entries
      _, _ -> entries
    _ -> entries

-- This is intentionally not general inference. Missing types, local globals,
-- higher-rank arguments, unresolved variables and dynamic Any all decline reuse.
importedApplicationType :: CodegenMetadata -> ModuleName -> ApplicationKey -> Maybe ExprType
importedApplicationType metadata current = case _ of
  GlobalKey (Qualified (Just moduleName@(CoreFn.ModuleName name)) (Ident ident)) -> do
    guard (moduleName /= current)
    Map.lookup (name <> "." <> ident) metadata.globalTypes
  CurriedKey fn args -> infer fn args
  UncurriedKey fn args -> infer fn args
  _ -> Nothing
  where
  infer fn args = do
    fnType <- importedApplicationType metadata current fn
    argTypes <- traverse (importedApplicationType metadata current) args
    guard (Array.all monomorphic argTypes)
    applicationResult fnType argTypes

applicationResult :: ExprType -> Array ExprType -> Maybe ExprType
applicationResult ty args = case Array.uncons args of
  Nothing -> Just ty
  Just { head: arg, tail: rest } -> case ty of
    CoreFn.ForAll _ body -> applicationResult body args
    CoreFn.ConstrainedType constraints body -> applicationResult
      (CoreFn.Func (map (\(Tuple path types) -> CoreFn.ADT (String.joinWith "." path) path types) constraints) body) args
    CoreFn.Func params result -> do
      { head: param, tail: remaining } <- Array.uncons params
      let substitution = unify param arg Map.empty
      guard (substitute substitution param == arg)
      let next = if Array.null remaining then result else CoreFn.Func remaining result
      applicationResult (substitute substitution next) rest
    _ -> Nothing

monomorphic :: ExprType -> Boolean
monomorphic = case _ of
  CoreFn.Any -> false
  CoreFn.TypeVar _ -> false
  CoreFn.ForAll _ _ -> false
  CoreFn.ConstrainedType _ _ -> false
  CoreFn.ADT _ _ args -> Array.all monomorphic args
  CoreFn.Array item -> monomorphic item
  CoreFn.Func args result -> Array.all monomorphic args && monomorphic result
  CoreFn.TypeApp fn args -> monomorphic fn && Array.all monomorphic args
  CoreFn.Record row -> monomorphic row
  CoreFn.Row fields tail -> Array.all (\(Tuple _ field) -> monomorphic field) fields
    && case tail of
      Nothing -> true
      Just row -> monomorphic row
  _ -> true

cacheClosedDictionaries :: CodegenMetadata -> BackendModule -> BackendModule
cacheClosedDictionaries metadata mod =
  let
    existing = foldl
      (\acc (Tuple (Ident name) _) -> Set.insert name acc)
      Set.empty
      (Array.concatMap _.bindings mod.bindings)

    initial :: LiftState
    initial = { next: 0, names: existing, moduleName: mod.name, lifted: [], shared: sharedDictionaries metadata mod }

    Tuple rewritten liftedState = runState (traverse (rewriteGroup metadata) mod.bindings) initial
  in
    if Array.null liftedState.lifted then mod { bindings = rewritten }
    else mod
      { bindings = rewritten <> [ { recursive: false, bindings: liftedState.lifted } ]
      }

rewriteGroup :: CodegenMetadata -> BackendBindingGroup Ident NeutralExpr -> State LiftState (BackendBindingGroup Ident NeutralExpr)
rewriteGroup metadata group
  | group.recursive = pure group
  | otherwise = do
      bindings <- traverse
        (\(Tuple ident expr) -> Tuple ident <$> rewriteNested metadata expr)
        group.bindings
      pure group { bindings = bindings }

-- A top-level binding body already lives in a cached getter, so it is not
-- lifted a second time; only constructions nested below it are.
rewriteNested :: CodegenMetadata -> NeutralExpr -> State LiftState NeutralExpr
rewriteNested metadata (NeutralExpr syn) =
  case syn of
    LetRec _ _ _ -> pure (NeutralExpr syn)
    Typed ty body -> NeutralExpr <<< Typed ty <$> rewriteNested metadata body
    TypeApp body ty -> NeutralExpr <<< flip TypeApp ty <$> rewriteNested metadata body
    _ -> NeutralExpr <$> traverse (rewriteExpr metadata) syn

rewriteExpr :: CodegenMetadata -> NeutralExpr -> State LiftState NeutralExpr
rewriteExpr metadata expr@(NeutralExpr syn) = do
  state <- get
  case applicationKey expr >>= flip Map.lookup state.shared of
    Just known -> pure (reference state.moduleName known.ident known.type)
    Nothing -> case liftable metadata expr of
      Just ty -> lift ty expr
      Nothing -> case syn of
        -- A recursive local scope can include early reads of otherwise global
        -- values; do not introduce shared-getter dependencies inside it.
        LetRec _ _ _ -> pure expr
        Typed ty body -> NeutralExpr <<< Typed ty <$> rewriteNested metadata body
        TypeApp body ty -> NeutralExpr <<< flip TypeApp ty <$> rewriteNested metadata body
        _ -> NeutralExpr <$> traverse (rewriteExpr metadata) syn

reference :: ModuleName -> Ident -> ExprType -> NeutralExpr
reference moduleName ident ty = NeutralExpr (Typed ty (NeutralExpr (Var (Qualified (Just moduleName) ident))))

liftable :: CodegenMetadata -> NeutralExpr -> Maybe ExprType
liftable metadata expr = do
  guard (isApplication expr)
  guard (Set.isEmpty (freeVars expr))
  guard (not (containsEffect expr))
  guard (not (containsRecursion expr))
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

containsRecursion :: NeutralExpr -> Boolean
containsRecursion (NeutralExpr syn) = case syn of
  LetRec _ _ _ -> true
  _ -> foldl (\found child -> found || containsRecursion child) false syn

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
  pure (reference st.moduleName ident ty)

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
