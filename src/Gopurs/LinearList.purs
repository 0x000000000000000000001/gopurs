-- | Consuming TRMC workers for singly-linked Int pointer lists.
-- |
-- | Recognizes module-local singly-linked list ADTs (monomorphic or polymorphic
-- | instantiated at Int) and self-recursive TRMC list transformers of the form
-- |
-- |   f Nil = Nil
-- |   f (Cons h t) = Cons (headExpr h) (f (nextList))
-- |
-- | where nextList is either `t` or `filter (\x -> pred x h) t` for a proven
-- | module-local list filter and a pure total scalar predicate.
-- |
-- | For each admitted transformer, emits a consuming worker that runs as a flat
-- | loop, reuses the dead input Cons cell in place, and (when filtering) relinks
-- | surviving cells in place with zero heap allocations. Only call sites whose
-- | argument is proven freshly constructed (no external aliases) are rewritten
-- | to the consuming worker; all other callers retain the persistent worker.
module Gopurs.LinearList (prepare) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (State, get, put, runState)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Either (Either(..))
import Data.Foldable (all, any, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), snd)
import Gopurs.CodegenState (CodegenMetadata, FunctionInfo)
import Gopurs.ConstructorLayout as Layout
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), StructPointer, constructorNames, goTypeToStr, rawGo, sanitizeName)
import Gopurs.GoSymbol (GoReference(..), generated)
import Gopurs.GoTypes (exprTypeToGoType)
import Gopurs.Printer (printGoExpr)
import Gopurs.WorkerNames as WorkerNames
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendSyntax(..), Level, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type LocalRef = Tuple (Maybe Ident) Level

type ListShape =
  { ty :: ExprType
  , pointer :: StructPointer
  , goType :: GoType
  , empty :: Qualified Ident
  , cons :: Qualified Ident
  }

data TailStep
  = TailDirect
  | TailFiltered
      { predParam :: LocalRef
      , predBody :: NeutralExpr
      }

type LinearCandidate =
  { original :: Ident
  , worker :: Ident
  , native :: String
  , shape :: ListShape
  , input :: LocalRef
  , lets :: Array { ref :: LocalRef, rhs :: GoExpr, goType :: GoType }
  , headExpr :: GoExpr
  , headIsOriginal :: Boolean
  , tailStep :: TailStep
  , predExpr :: Maybe GoExpr
  , predExpr32 :: Maybe GoExpr
  , hoistedDivisors :: Array { varName :: String, divisorExpr :: GoExpr }
  }

type FreshProducer =
  { accIndex :: Int
  , arity :: Int
  , shape :: ListShape
  }

type RangeLoop =
  { currIndex :: Int
  , accIndex :: Int
  , shape :: ListShape
  , limitExpr :: NeutralExpr
  }

type TopRangeProducer =
  { startIndex :: Int
  , endIndex :: Int
  , shape :: ListShape
  }

type FoldLoop =
  { listIndex :: Int
  , accIndex :: Int
  , arity :: Int
  , shape :: ListShape
  , stepExpr :: GoExpr
  , isAddFold :: Boolean
  , stepKey :: String
  }

type TopFolder =
  { shape :: ListShape
  , listIndex :: Int
  , seedSource :: Either Int NeutralExpr
  , arity :: Int
  , requireHarmlessAt :: Maybe Int
  , stepExpr :: GoExpr
  , isAddFold :: Boolean
  , stepKey :: String
  }

type FusedWorker =
  { worker :: Ident
  , native :: String
  , candidate :: LinearCandidate
  , stepExpr :: GoExpr
  , isAddFold :: Boolean
  }

type RewriteState =
  { names :: WorkerNames.Names
  , fusedByKey :: Map.Map String FusedWorker
  , fusedItems :: Array FusedWorker
  }

prepare
  :: CodegenMetadata
  -> BackendModule
  -> { module :: BackendModule, declarations :: Array GoDecl, functions :: Map.Map String FunctionInfo }
prepare metadata mod =
  let
    bindings = Array.concatMap _.bindings mod.bindings
    shapes = collectListShapes metadata mod
    filters = collectFilters mod.name shapes bindings
    initialNames = WorkerNames.fromModule mod
    admitted = foldl (collectCandidate mod.name shapes filters) { names: initialNames, items: [] } mod.bindings
    candidates = Map.fromFoldable $ map (\c -> Tuple c.original c) admitted.items
  in
    if Map.isEmpty candidates then
      { module: mod, declarations: [], functions: Map.empty }
    else
      let
        unit = modulePrefix mod.name
        topProducers = collectTopFreshProducers mod.name shapes bindings
        topRanges = collectTopRangeProducers mod.name shapes bindings
        topFolders = collectTopFolders mod.name shapes bindings
        initState = { names: admitted.names, fusedByKey: Map.empty, fusedItems: [] }
        env =
          { current: mod.name
          , unit
          , shapes
          , topProducers
          , topRanges
          , topFolders
          , candidates
          , localProducers: Map.empty
          , localRanges: Map.empty
          , localFolders: Map.empty
          }
        Tuple rewrittenBindings finalState = runState
          (traverse (\group -> do
            bs <- traverse (\(Tuple name expr) -> Tuple name <$> rewriteExpr env expr) group.bindings
            pure (group { bindings = bs })) mod.bindings)
          initState
        workerDecls = map (emitWorker unit) admitted.items
        fusedDecls = map (emitFusedRangeFoldWorker unit) finalState.fusedItems
        workerFuncs = Map.fromFoldable $ map (\c -> Tuple (sanitizeName (unwrap c.worker))
          { arrayFold: Nothing
          , codePointFold: Nothing
          , reference: SymbolReference (generated unit c.native)
          , fArgs: [ c.shape.goType ]
          , fRet: c.shape.goType
          , arity: 1
          }) admitted.items
        fusedFuncs = Map.fromFoldable $ map (\f -> Tuple (sanitizeName (unwrap f.worker))
          { arrayFold: Nothing
          , codePointFold: Nothing
          , reference: SymbolReference (generated unit f.native)
          , fArgs: [ TypeInt64, TypeInt64, TypeInt64 ]
          , fRet: TypeInt64
          , arity: 3
          }) finalState.fusedItems
      in
        { module: mod { bindings = rewrittenBindings }
        , declarations: workerDecls <> fusedDecls
        , functions: Map.union workerFuncs fusedFuncs
        }

modulePrefix :: ModuleName -> String
modulePrefix = String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syn) = case syn of
  Typed _ inner -> strip inner
  Syn.TypeApp inner _ -> strip inner
  _ -> syn

annotation :: NeutralExpr -> Maybe ExprType
annotation (NeutralExpr syn) = case syn of
  Typed ty _ -> Just ty
  Syn.TypeApp inner _ -> annotation inner
  _ -> Nothing

arrow :: ExprType -> { args :: Array ExprType, result :: ExprType }
arrow (ForAll _ ty) = arrow ty
arrow (Func args result) = let next = arrow result in { args: args <> next.args, result: next.result }
arrow ty = { args: [], result: ty }

abstractions :: NeutralExpr -> { args :: Array LocalRef, body :: NeutralExpr }
abstractions expr = case strip expr of
  Abs args body -> let next = abstractions body in { args: NEA.toArray args <> next.args, body: next.body }
  UncurriedAbs args body | not (Array.null args) ->
    let next = abstractions body in { args: args <> next.args, body: next.body }
  _ -> { args: [], body: expr }

spine :: NeutralExpr -> { head :: NeutralExpr, args :: Array NeutralExpr }
spine expr = case strip expr of
  App fn args -> let prior = spine fn in prior { args = prior.args <> NEA.toArray args }
  UncurriedApp fn args -> let prior = spine fn in prior { args = prior.args <> args }
  _ -> { head: expr, args: [] }

qualify :: ModuleName -> Qualified Ident -> Qualified Ident
qualify current (Qualified mbMod ident) = Qualified (Just (fromMaybe current mbMod)) ident

sameLocal :: LocalRef -> NeutralExpr -> Boolean
sameLocal (Tuple _ lvl) expr = case strip expr of
  Local _ found -> found == lvl
  _ -> false

localGlobal :: ModuleName -> NeutralExpr -> Maybe Ident
localGlobal current expr = case strip expr of
  Var (Qualified mbMod name) | mbMod == Nothing || mbMod == Just current -> Just name
  _ -> Nothing

-- Collect all module-local singly-linked Int list ADTs (monomorphic or
-- polymorphic instantiated at Int) whose native Go layout is a verified
-- 2-field pointer struct { Rc, V0 int64, V1 *Self }.
collectListShapes :: CodegenMetadata -> BackendModule -> Array ListShape
collectListShapes metadata mod = Array.mapMaybe make mod.dataDecls
  where
  unit = modulePrefix mod.name
  make decl = do
    let fullName = unwrap mod.name <> "." <> decl.name
        typeArgs = Array.replicate (Array.length decl.vars) Int
        ty = ADT fullName (String.split (Pattern ".") fullName) typeArgs
        variables = Map.fromFoldable (Array.zip decl.vars typeArgs)
        substitute = case _ of
          TypeVar name -> fromMaybe (TypeVar name) (Map.lookup name variables)
          ADT name path args -> ADT name path (map substitute args)
          other -> other
    Tuple emptyCtor consCtor <- case decl.constructors of
      [ a, b ] | Array.null a.fields -> Just (Tuple a b)
      [ a, b ] | Array.null b.fields -> Just (Tuple b a)
      _ -> Nothing
    guard (map substitute consCtor.fields == [ Int, ty ])
    let consNames = constructorNames unit consCtor.name
        emptyNames = constructorNames unit emptyCtor.name
    path <- Map.lookup fullName metadata.pointerAdtPaths
    guard (path.ctorName == consCtor.name && path.arity == Array.length decl.vars)
    leaf <- Map.lookup emptyNames.baseStructName metadata.pointerAdtLeaves
    guard (leaf.nodeBaseStruct == consNames.baseStructName && leaf.nodeCtor == consCtor.name)
    pointer <- case exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors unit ty of
      TypeStructPointer p -> Just p
      _ -> Nothing
    let goType = TypeStructPointer pointer
        layout = Layout.layout metadata unit consCtor.name (Just goType)
    fields <- layout.constructorFields
    guard (Array.length fields.vars == Array.length pointer.typeArgs && Array.length fields.fields == 2)
    guard (Layout.fieldType metadata unit fields pointer.typeArgs 0 == TypeInt64)
    guard (Layout.fieldType metadata unit fields pointer.typeArgs 1 == goType)
    pure
      { ty
      , pointer
      , goType
      , empty: Qualified (Just mod.name) (Ident emptyCtor.name)
      , cons: Qualified (Just mod.name) (Ident consCtor.name)
      }

isEmpty :: ModuleName -> ListShape -> NeutralExpr -> Boolean
isEmpty current shape expr = case strip expr of
  CtorSaturated ctor _ _ _ fields -> qualify current ctor == shape.empty && Array.null fields
  Var ctor -> qualify current ctor == shape.empty
  _ -> false

consFields :: ModuleName -> ListShape -> NeutralExpr -> Maybe (Tuple NeutralExpr NeutralExpr)
consFields current shape expr = case strip expr of
  CtorSaturated ctor _ _ _ [ Tuple _ h, Tuple _ t ] | qualify current ctor == shape.cons -> Just (Tuple h t)
  _ -> Nothing

isField :: ModuleName -> ListShape -> LocalRef -> Int -> NeutralExpr -> Boolean
isField current shape ref index expr = case strip expr of
  Accessor base (GetCtorField ctor _ _ _ _ found) ->
    qualify current ctor == shape.cons && found == index && sameLocal ref base
  _ -> false

isTag :: ModuleName -> Qualified Ident -> LocalRef -> NeutralExpr -> Boolean
isTag current expected ref expr = case strip expr of
  PrimOp (Op1 (OpIsTag ctor) target) -> qualify current ctor == expected && sameLocal ref target
  _ -> false

listCases :: ModuleName -> ListShape -> LocalRef -> NeutralExpr -> Maybe { empty :: NeutralExpr, cons :: NeutralExpr }
listCases current shape ref expr = case strip expr of
  Branch choices failure -> case NEA.toArray choices, strip failure of
    [ Pair emptyTest emptyBody, Pair consTest consBody ], Fail _
      | isTag current shape.empty ref emptyTest && isTag current shape.cons ref consTest ->
        Just { empty: emptyBody, cons: consBody }
    [ Pair consTest consBody, Pair emptyTest emptyBody ], Fail _
      | isTag current shape.empty ref emptyTest && isTag current shape.cons ref consTest ->
        Just { empty: emptyBody, cons: consBody }
    _, _ -> Nothing
  _ -> Nothing

-- Recognize module-local list filter functions (either tail-recursive
-- filter/prepend/reverse matching ListFilter.purs, or direct-recursive filter).
collectFilters :: ModuleName -> Array ListShape -> Array (Tuple Ident NeutralExpr) -> Map.Map Ident ListShape
collectFilters current shapes bindings =
  Map.fromFoldable $ Array.mapMaybe (\(Tuple name expr) -> map (Tuple name) (recognizeFilter current shapes name expr)) bindings

recognizeFilter :: ModuleName -> Array ListShape -> Ident -> NeutralExpr -> Maybe ListShape
recognizeFilter current shapes name expr = do
  sig <- arrow <$> annotation expr
  listTy <- case sig.args, sig.result of
    [ Func [ Int ] Boolean, ty ], ret | ty == ret -> Just ty
    _, _ -> Nothing
  shape <- Array.find (\s -> s.ty == listTy) shapes
  let lambda = abstractions expr
  Tuple predicate source <- case lambda.args of
    [ p, s ] -> Just (Tuple p s)
    _ -> Nothing
  if matchReverseFilter current shape predicate source lambda.body
    || matchDirectFilter current shape name predicate source lambda.body then
    Just shape
  else
    Nothing

type LocalLoop =
  { level :: Level
  , name :: Ident
  , input :: LocalRef
  , acc :: LocalRef
  , body :: NeutralExpr
  , initial :: Array NeutralExpr
  }

matchLocalLoop :: ListShape -> NeutralExpr -> Maybe LocalLoop
matchLocalLoop shape expr = case strip expr of
  LetRec level bindings body -> do
    let entry = spine body
    loopName <- case strip entry.head of
      Local (Just n) l | l == level -> Just n
      _ -> Nothing
    selected <- snd <$> Array.find (\(Tuple n _) -> n == loopName) (NEA.toArray bindings)
    guard (all (not <<< Array.null <<< _.args <<< abstractions <<< snd) bindings)
    sig <- arrow <$> annotation selected
    guard (sig.args == [ shape.ty, shape.ty ] && sig.result == shape.ty)
    let lambda = abstractions selected
    Tuple input acc <- case lambda.args of
      [ a, b ] -> Just (Tuple a b)
      _ -> Nothing
    guard (snd input /= snd acc && snd input /= level && snd acc /= level)
    pure { level, name: loopName, input, acc, body: lambda.body, initial: entry.args }
  _ -> Nothing

matchLoopSelfCall :: LocalLoop -> NeutralExpr -> Maybe (Tuple NeutralExpr NeutralExpr)
matchLoopSelfCall loop expr = do
  let call = spine expr
  guard case strip call.head of
    Local _ l -> l == loop.level
    _ -> false
  case call.args of
    [ a, b ] -> Just (Tuple a b)
    _ -> Nothing

matchReverseLoop :: ModuleName -> ListShape -> LocalRef -> NeutralExpr -> Boolean
matchReverseLoop current shape source expr = case check of
  Just _ -> true
  Nothing -> false
  where
  check = do
    loop <- matchLocalLoop shape expr
    Tuple initSrc initAcc <- case loop.initial of
      [ a, b ] -> Just (Tuple a b)
      _ -> Nothing
    guard (sameLocal source initSrc && isEmpty current shape initAcc)
    cases <- listCases current shape loop.input loop.body
    guard (sameLocal loop.acc cases.empty)
    Tuple nextTail nextAcc <- matchLoopSelfCall loop cases.cons
    Tuple head tail <- consFields current shape nextAcc
    guard (isField current shape loop.input 1 nextTail
      && isField current shape loop.input 0 head
      && sameLocal loop.acc tail)

matchReverseFilter :: ModuleName -> ListShape -> LocalRef -> LocalRef -> NeutralExpr -> Boolean
matchReverseFilter current shape predicate source body = case check of
  Just _ -> true
  Nothing -> false
  where
  check = do
    loop <- matchLocalLoop shape body
    Tuple initSrc initAcc <- case loop.initial of
      [ a, b ] -> Just (Tuple a b)
      _ -> Nothing
    guard (sameLocal source initSrc && isEmpty current shape initAcc)
    cases <- listCases current shape loop.input loop.body
    guard (matchReverseLoop current shape loop.acc cases.empty)
    { cond, yes, no } <- case strip cases.cons of
      Branch choices no -> case NEA.toArray choices of
        [ Pair cond yes ] -> Just { cond, yes, no }
        _ -> Nothing
      _ -> Nothing
    let call = spine cond
    arg <- case call.args of
      [ a ] -> Just a
      _ -> Nothing
    guard (sameLocal predicate call.head && isField current shape loop.input 0 arg)
    Tuple yesTail yesAcc <- matchLoopSelfCall loop yes
    Tuple noTail noAcc <- matchLoopSelfCall loop no
    Tuple head tail <- consFields current shape yesAcc
    guard (isField current shape loop.input 1 yesTail && isField current shape loop.input 1 noTail)
    guard (isField current shape loop.input 0 head && sameLocal loop.acc tail && sameLocal loop.acc noAcc)

matchDirectFilter :: ModuleName -> ListShape -> Ident -> LocalRef -> LocalRef -> NeutralExpr -> Boolean
matchDirectFilter current shape name predicate source body = case check of
  Just _ -> true
  Nothing -> false
  where
  check = do
    cases <- listCases current shape source body
    guard (isEmpty current shape cases.empty)
    { cond, yes, no } <- case strip cases.cons of
      Branch choices no -> case NEA.toArray choices of
        [ Pair cond yes ] -> Just { cond, yes, no }
        _ -> Nothing
      _ -> Nothing
    let call = spine cond
    arg <- case call.args of
      [ a ] -> Just a
      _ -> Nothing
    guard (sameLocal predicate call.head && isField current shape source 0 arg)
    Tuple head tailCall <- consFields current shape yes
    guard (isField current shape source 0 head)
    let isSelfTail e = case spine e of
          { head: h, args: [ pArg, tArg ] } ->
            localGlobal current h == Just name && sameLocal predicate pArg && isField current shape source 1 tArg
          _ -> false
    guard (isSelfTail tailCall && isSelfTail no)

-- Recognize a self-recursive linear TRMC list transformer.
collectCandidate
  :: ModuleName
  -> Array ListShape
  -> Map.Map Ident ListShape
  -> { names :: WorkerNames.Names, items :: Array LinearCandidate }
  -> { recursive :: Boolean, bindings :: Array (Tuple Ident NeutralExpr) }
  -> { names :: WorkerNames.Names, items :: Array LinearCandidate }
collectCandidate current shapes filters acc group = case group.bindings of
  [ Tuple name expr ] | group.recursive -> case recognizeCandidate current shapes filters name expr of
    Just found ->
      let
        worker = WorkerNames.fresh "__gopurs_linear_" name acc.names
        native = "Call_" <> modulePrefix current <> "_" <> sanitizeName (unwrap worker)
        candidate = found { worker = worker, native = native }
      in
        { names: WorkerNames.reserve worker acc.names
        , items: Array.snoc acc.items candidate
        }
    Nothing -> acc
  _ -> acc

type ScalarVal = { expr :: GoExpr, goType :: GoType }
type ScalarEnv = Map.Map Level ScalarVal

recognizeCandidate
  :: ModuleName
  -> Array ListShape
  -> Map.Map Ident ListShape
  -> Ident
  -> NeutralExpr
  -> Maybe LinearCandidate
recognizeCandidate current shapes filters name expr = do
  sig <- arrow <$> annotation expr
  listTy <- case sig.args, sig.result of
    [ ty ], ret | ty == ret -> Just ty
    _, _ -> Nothing
  shape <- Array.find (\s -> s.ty == listTy) shapes
  let lambda = abstractions expr
  input <- case lambda.args of
    [ a ] -> Just a
    _ -> Nothing
  cases <- listCases current shape input lambda.body
  guard (isEmpty current shape cases.empty)
  -- Peel any leading scalar Let bindings (e.g. `p = cur.V0`) inside the Cons branch.
  peeled <- peelScalarLets current shape input Map.empty [] cases.cons
  Tuple headRaw tailCallRaw <- consFields current shape peeled.body
  headVal <- compileScalar current shape input "__curVal" Nothing peeled.env headRaw
  guard (headVal.goType == TypeInt64)
  let headIsOriginal = isField current shape input 0 headRaw || case strip headRaw of
        Local _ lvl -> Map.lookup lvl peeled.origHeadLocals == Just true
        _ -> false
      tailCall = spine tailCallRaw
  guard (localGlobal current tailCall.head == Just name)
  nextArg <- case tailCall.args of
    [ a ] -> Just a
    _ -> Nothing
  Tuple tailStep { predExpr, predExpr32, hoistedDivisors } <-
    if isField current shape input 1 nextArg then
      Just (Tuple TailDirect { predExpr: Nothing, predExpr32: Nothing, hoistedDivisors: [] })
    else do
      let fCall = spine nextArg
      filterName <- localGlobal current fCall.head
      filterShape <- Map.lookup filterName filters
      guard (filterShape.ty == shape.ty)
      Tuple predLambda tailArg <- case fCall.args of
        [ p, t ] -> Just (Tuple p t)
        _ -> Nothing
      guard (isField current shape input 1 tailArg)
      let pLam = abstractions predLambda
      predParam <- case pLam.args of
        [ x ] -> Just x
        _ -> Nothing
      guard (snd predParam /= snd input && not (Map.member (snd predParam) peeled.env))
      let predEnv = Map.insert (snd predParam) { expr: GoVar "__x", goType: TypeInt64 } peeled.env
      predVal <- compileScalar current shape input "__curVal" Nothing predEnv pLam.body
      guard (predVal.goType == TypeBool)
      let pred32Res = compileScalar32 current shape input (snd predParam) peeled.env predEnv pLam.body
      pure (Tuple (TailFiltered { predParam, predBody: pLam.body })
        { predExpr: Just predVal.expr
        , predExpr32: map _.expr pred32Res
        , hoistedDivisors: fromMaybe [] (map _.divisors pred32Res)
        })
  pure
    { original: name
    , worker: name
    , native: ""
    , shape
    , input
    , lets: peeled.lets
    , headExpr: headVal.expr
    , headIsOriginal
    , tailStep
    , predExpr
    , predExpr32
    , hoistedDivisors
    }

peelScalarLets
  :: ModuleName
  -> ListShape
  -> LocalRef
  -> ScalarEnv
  -> Array { ref :: LocalRef, rhs :: GoExpr, goType :: GoType }
  -> NeutralExpr
  -> Maybe
       { env :: ScalarEnv
       , origHeadLocals :: Map.Map Level Boolean
       , lets :: Array { ref :: LocalRef, rhs :: GoExpr, goType :: GoType }
       , body :: NeutralExpr
       }
peelScalarLets current shape input env lets expr = go env Map.empty lets expr
  where
  go curEnv origHeads curLets currentExpr = case strip currentExpr of
    Let mbIdent lvl rhs next -> do
      compiled <- compileScalar current shape input "__curVal" Nothing curEnv rhs
      let varName = "__s" <> show (Array.length curLets)
          isOrig = isField current shape input 0 rhs || case strip rhs of
            Local _ l -> Map.lookup l origHeads == Just true
            _ -> false
          nextEnv = Map.insert lvl { expr: GoVar varName, goType: compiled.goType } curEnv
          nextOrig = if isOrig then Map.insert lvl true origHeads else origHeads
          nextLets = Array.snoc curLets { ref: Tuple mbIdent lvl, rhs: compiled.expr, goType: compiled.goType }
      go nextEnv nextOrig nextLets next
    _ -> Just { env: curEnv, origHeadLocals: origHeads, lets: curLets, body: currentExpr }

type HoistedDivisor = { varName :: String, divisorExpr :: GoExpr }

compileScalar32
  :: ModuleName
  -> ListShape
  -> LocalRef
  -> Level
  -> ScalarEnv
  -> ScalarEnv
  -> NeutralExpr
  -> Maybe { expr :: GoExpr, divisors :: Array HoistedDivisor }
compileScalar32 current shape input predParamLvl outerEnv predEnv expr = do
  val <- compileScalar current shape input "__curVal" (Just { predParamLvl, outerEnv }) predEnv expr
  guard (val.goType == TypeBool)
  pure { expr: val.expr, divisors: val.divisors }

type ScalarValOut = { expr :: GoExpr, goType :: GoType, divisors :: Array HoistedDivisor }
type Mode32 = { predParamLvl :: Level, outerEnv :: ScalarEnv }

-- Compile pure total Int/Boolean scalar expressions. Any call, effect,
-- closure, or reference to the input list spine itself is rejected.
-- When mode32 is Just, input elements are proven to fit signed int32, so
-- loop-invariant modulo-zero checks hoist a Lemire Divisor32 reciprocal.
compileScalar :: ModuleName -> ListShape -> LocalRef -> String -> Maybe Mode32 -> ScalarEnv -> NeutralExpr -> Maybe ScalarValOut
compileScalar current shape input headVar mode32 env expr = case strip expr of
  Lit (LitInt n) -> Just { expr: GoCall (GoVar "int64") [ GoInt n ], goType: TypeInt64, divisors: [] }
  Lit (LitBoolean b) -> Just { expr: rawGo (if b then "true" else "false"), goType: TypeBool, divisors: [] }
  Local _ lvl -> map (\v -> { expr: v.expr, goType: v.goType, divisors: [] }) (Map.lookup lvl env)
  Accessor base (GetCtorField ctor _ _ _ _ 0)
    | qualify current ctor == shape.cons && sameLocal input base ->
      Just { expr: GoVar headVar, goType: TypeInt64, divisors: [] }
  PrimOp (Op1 OpBooleanNot child) -> do
    c <- compileScalar current shape input headVar mode32 env child
    guard (c.goType == TypeBool)
    pure { expr: GoPrefixOp "!" c.expr, goType: TypeBool, divisors: c.divisors }
  PrimOp (Op1 OpIntNegate child) -> do
    c <- compileScalar current shape input headVar mode32 env child
    guard (c.goType == TypeInt64)
    pure { expr: GoCall (GoVar "int64") [ GoPrefixOp "-" (GoCall (GoVar "int32") [ c.expr ]) ], goType: TypeInt64, divisors: c.divisors }
  PrimOp (Op1 OpIntBitNot child) -> do
    c <- compileScalar current shape input headVar mode32 env child
    guard (c.goType == TypeInt64)
    pure { expr: GoCall (GoVar "int64") [ GoPrefixOp "^" (GoCall (GoVar "int32") [ c.expr ]) ], goType: TypeInt64, divisors: c.divisors }
  PrimOp (Op2 op left right) -> compileBinOp current shape input headVar mode32 env op left right
  Branch branches fallback -> case NEA.toArray branches of
    [ Pair cond yes ] -> do
      c <- compileScalar current shape input headVar mode32 env cond
      guard (c.goType == TypeBool)
      case strip yes, strip fallback of
        Lit (LitBoolean false), Lit (LitBoolean true) ->
          Just { expr: GoPrefixOp "!" c.expr, goType: TypeBool, divisors: c.divisors }
        Lit (LitBoolean true), Lit (LitBoolean false) ->
          Just c
        _, _ -> Nothing
    _ -> Nothing
  _ -> Nothing

compileBinOp :: ModuleName -> ListShape -> LocalRef -> String -> Maybe Mode32 -> ScalarEnv -> BackendOperator2 -> NeutralExpr -> NeutralExpr -> Maybe ScalarValOut
compileBinOp current shape input headVar mode32 env op left right =
  case op of
    OpIntOrd OpEq -> case modZeroCheck left right of
      Just val -> Just val
      Nothing -> binOrd "==" left right
    OpIntOrd OpNotEq -> case modZeroCheck left right of
      Just val -> Just (val { expr = GoPrefixOp "!" val.expr })
      Nothing -> binOrd "!=" left right
    OpIntOrd OpLt -> binOrd "<" left right
    OpIntOrd OpLte -> binOrd "<=" left right
    OpIntOrd OpGt -> binOrd ">" left right
    OpIntOrd OpGte -> binOrd ">=" left right
    OpIntNum OpAdd -> binRuntime "IntAdd" left right
    OpIntNum OpSubtract -> binRuntime "IntSub" left right
    OpIntNum OpMultiply -> binRuntime "IntMul" left right
    OpIntNum OpDivide -> binRuntime "IntDiv" left right
    OpIntNum OpMod -> binRuntime "IntMod" left right
    OpBooleanAnd -> binBool "&&" left right
    OpBooleanOr -> binBool "||" left right
    OpBooleanOrd OpEq -> binBool "==" left right
    OpBooleanOrd OpNotEq -> binBool "!=" left right
    _ -> Nothing
  where
  binOrd sym a b = do
    la <- compileScalar current shape input headVar mode32 env a
    rb <- compileScalar current shape input headVar mode32 env b
    guard (la.goType == TypeInt64 && rb.goType == TypeInt64)
    pure { expr: GoBinOp sym la.expr rb.expr, goType: TypeBool, divisors: la.divisors <> rb.divisors }
  binBool sym a b = do
    la <- compileScalar current shape input headVar mode32 env a
    rb <- compileScalar current shape input headVar mode32 env b
    guard (la.goType == TypeBool && rb.goType == TypeBool)
    pure { expr: GoBinOp sym la.expr rb.expr, goType: TypeBool, divisors: la.divisors <> rb.divisors }
  binRuntime fn a b = do
    la <- compileScalar current shape input headVar mode32 env a
    rb <- compileScalar current shape input headVar mode32 env b
    guard (la.goType == TypeInt64 && rb.goType == TypeInt64)
    pure { expr: GoCall (GoSelector (GoVar "gopurs_runtime") fn) [ la.expr, rb.expr ], goType: TypeInt64, divisors: la.divisors <> rb.divisors }
  modZeroCheck a b = case strip a, strip b of
    PrimOp (Op2 (OpIntNum OpMod) x y), Lit (LitInt 0) -> compileModZero x y
    Lit (LitInt 0), PrimOp (Op2 (OpIntNum OpMod) x y) -> compileModZero x y
    _, _ -> Nothing
  compileModZero x y = do
    lx <- compileScalar current shape input headVar mode32 env x
    ry <- compileScalar current shape input headVar mode32 env y
    guard (lx.goType == TypeInt64 && ry.goType == TypeInt64)
    case mode32 of
      Just m | not (mentionsLevel m.predParamLvl y) ->
        let
          divName = "__div0"
          -- Only hoist when no earlier divisor was already named __div0;
          -- otherwise fall back to direct int32 modulo.
        in if Array.null lx.divisors && Array.null ry.divisors then
          pure
            { expr: GoCall (GoSelector (GoVar divName) "Divides") [ GoCall (GoVar "int32") [ lx.expr ] ]
            , goType: TypeBool
            , divisors: [ { varName: divName, divisorExpr: ry.expr } ]
            }
        else
          pure
            { expr: GoBinOp "||"
                (GoBinOp "==" (GoCall (GoVar "int32") [ ry.expr ]) (GoCall (GoVar "int32") [ GoInt 0 ]))
                (GoBinOp "==" (GoBinOp "%" (GoCall (GoVar "int32") [ lx.expr ]) (GoCall (GoVar "int32") [ ry.expr ])) (GoCall (GoVar "int32") [ GoInt 0 ]))
            , goType: TypeBool
            , divisors: lx.divisors <> ry.divisors
            }
      Just _ ->
        pure
          { expr: GoBinOp "||"
              (GoBinOp "==" (GoCall (GoVar "int32") [ ry.expr ]) (GoCall (GoVar "int32") [ GoInt 0 ]))
              (GoBinOp "==" (GoBinOp "%" (GoCall (GoVar "int32") [ lx.expr ]) (GoCall (GoVar "int32") [ ry.expr ])) (GoCall (GoVar "int32") [ GoInt 0 ]))
          , goType: TypeBool
          , divisors: lx.divisors <> ry.divisors
          }
      Nothing ->
        pure
          { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "IntModZero") [ lx.expr, ry.expr ]
          , goType: TypeBool
          , divisors: []
          }

-- Emit the flat TRMC consuming worker.
-- - Direct tail (`map`): single-pass in-place update of existing Cons cells.
-- - Filtered tail (`sieve` / iterative list filter): compacts elements in a
--   contiguous stack-backed buffer (eliminating O(N^2) pointer chasing and GC
--   write barriers, hoisting loop-invariant int32 divisors via Lemire
--   Divisor32, and keeping the int64 fallback out-of-line to minimize stack
--   frame size), then writes surviving values back into the existing Cons
--   cells in place with zero heap allocations for len <= 512.
emitWorker :: String -> LinearCandidate -> GoDecl
emitWorker unit c =
  let
    ptrStr = goTypeToStr c.shape.goType
    nilExpr = rawGo ("(" <> ptrStr <> ")(nil)")
    letStmts = Array.mapWithIndex (\i entry ->
      GoAssign ("__s" <> show i) entry.rhs) c.lets
    body = case c.tailStep, c.predExpr of
      TailFiltered _, Just predGo ->
        let
          pred32 = fromMaybe predGo c.predExpr32
          hoistStmts = map (\d ->
            GoAssign d.varName (GoCall (GoSelector (GoVar "gopurs_runtime") "NewDivisor32")
              [ GoCall (GoVar "int32") [ d.divisorExpr ] ])) c.hoistedDivisors
          writeBack =
            [ GoAssign "__outCell" (GoVar "__input")
            , rawGo ("var __outPrev " <> ptrStr)
            , GoForRange "__i := 0; __i < __pos; __i++"
                [ GoMutate "__outCell.V0" (rawGo "__buf[__i]")
                , GoMutate "__outPrev" (GoVar "__outCell")
                , GoMutate "__outCell" (GoStructAccess (GoVar "__outCell") "V1")
                ]
            , GoIfElse (GoBinOp "==" (GoVar "__outPrev") nilExpr)
                [ GoReturn nilExpr ]
                [ GoIfElse (GoBinOp "!=" (GoStructAccess (GoVar "__outPrev") "V1") nilExpr)
                    [ GoMutate "__outPrev.V1" nilExpr ]
                    []
                , GoReturn (GoVar "__input")
                ]
            ]
          writeBack32 =
            [ GoAssign "__outCell" (GoVar "__input")
            , rawGo ("var __outPrev " <> ptrStr)
            , GoForRange "__i := 0; __i < __pos; __i++"
                [ GoMutate "__outCell.V0" (rawGo "int64(__buf32[__i])")
                , GoMutate "__outPrev" (GoVar "__outCell")
                , GoMutate "__outCell" (GoStructAccess (GoVar "__outCell") "V1")
                ]
            , GoIfElse (GoBinOp "==" (GoVar "__outPrev") nilExpr)
                [ GoReturn nilExpr ]
                [ GoIfElse (GoBinOp "!=" (GoStructAccess (GoVar "__outPrev") "V1") nilExpr)
                    [ GoMutate "__outPrev.V1" nilExpr ]
                    []
                , GoReturn (GoVar "__input")
                ]
            ]
          int32Loop =
            [ rawGo "var __stack32 [512]int32"
            , rawGo "__buf32 := __stack32[:0]"
            , rawGo "__all32 := true"
            , GoForRange "__scan := __input; __scan != nil; __scan = __scan.V1"
                [ rawGo "__v := __scan.V0"
                , rawGo "__v32 := int32(__v)"
                , GoIfElse (rawGo "int64(__v32) != __v")
                    [ rawGo "__all32 = false"
                    , rawGo "break"
                    ]
                    []
                , rawGo "__buf32 = append(__buf32, __v32)"
                ]
            , GoIfElse (rawGo "__all32")
                ([ rawGo "__pos := 0"
                 , GoForRange "__pos < len(__buf32)"
                     ([ rawGo "__curVal := int64(__buf32[__pos])"
                      , rawGo "_ = __curVal"
                      ] <> letStmts
                        <> hoistStmts
                        <> (if c.headIsOriginal then [] else [ rawGo ("__buf32[__pos] = int32(" <> headExprStr <> ")") ])
                        <> [ rawGo "__pos++"
                           , rawGo "__rest := __buf32[__pos:]"
                           , rawGo "__write := 0"
                           , GoForRange "_, __v32 := range __rest"
                               [ rawGo "__x := int64(__v32)"
                               , rawGo "_ = __x"
                               , GoIfElse pred32
                                   [ rawGo "__rest[__write] = __v32"
                                   , rawGo "__write++"
                                   ]
                                   []
                               ]
                           , rawGo "__buf32 = __buf32[:__pos+__write]"
                           ])
                 ] <> writeBack32)
                []
            ]
          -- Wide int64 fallback: use a heap slice instead of a stack array so the
          -- common int32 fast path never pays a 4 KB stack frame or zeroing loop.
          int64Fallback =
            [ rawGo "__buf := make([]int64, 0, 64)"
            , GoForRange "__scan := __input; __scan != nil; __scan = __scan.V1"
                [ rawGo "__buf = append(__buf, __scan.V0)" ]
            , rawGo "__pos := 0"
            , GoForRange "__pos < len(__buf)"
                ([ rawGo "__curVal := __buf[__pos]"
                 , rawGo "_ = __curVal"
                 ] <> letStmts
                   <> (if c.headIsOriginal then [] else [ GoMutate "__buf[__pos]" c.headExpr ])
                   <> [ rawGo "__pos++"
                      , rawGo "__rest := __buf[__pos:]"
                      , rawGo "__write := 0"
                      , GoForRange "_, __x := range __rest"
                          [ rawGo "_ = __x"
                          , GoIfElse predGo
                              [ rawGo "__rest[__write] = __x"
                              , rawGo "__write++"
                              ]
                              []
                          ]
                      , rawGo "__buf = __buf[:__pos+__write]"
                      ])
            ] <> writeBack
          fastPath = if c.headIsOriginal then int32Loop else []
        in
          GoBlock
            ([ GoIfElse (GoBinOp "==" (GoVar "__input") nilExpr) [ GoReturn nilExpr ] [] ]
              <> fastPath
              <> int64Fallback)
      _, _ ->
        GoBlock
          [ GoForRange "__cur := __input; __cur != nil; __cur = __cur.V1"
              ([ rawGo "__curVal := __cur.V0"
               , rawGo "_ = __curVal"
               ] <> letStmts <> (if c.headIsOriginal then [] else [ GoMutate "__cur.V0" c.headExpr ]))
          , GoReturn (GoVar "__input")
          ]
    headExprStr = "__curVal"
  in
    GoFunctionDecl
      { name: generated unit c.native
      , params: [ Tuple "__input" c.shape.goType ]
      , result: c.shape.goType
      , body
      }

-- Emit a zero-allocation deforested worker for closed pipelines of the form
-- `fold (linearTransform (range lower upper)) seed`.
-- Eliminates all intermediate Cons cell allocations completely for ranges
-- fitting in the stack int32 buffer (and falls back seamlessly to a contiguous
-- int32/int64 slice otherwise while preserving exact PureScript semantics).
emitFusedRangeFoldWorker :: String -> FusedWorker -> GoDecl
emitFusedRangeFoldWorker unit fw =
  let
    c = fw.candidate
    letStmts = Array.mapWithIndex (\i entry ->
      GoAssign ("__s" <> show i) entry.rhs) c.lets
    accInit32 = if fw.isAddFold then [ rawGo "__acc32 := int32(__seed)" ] else [ rawGo "__acc := __seed" ]
    accStep32 =
      if fw.isAddFold then
        if c.headIsOriginal then
          [ rawGo "__acc32 += __buf32[__pos]" ]
        else
          [ GoAssign "__elem" c.headExpr
          , rawGo "__acc32 += int32(__elem)"
          ]
      else
        [ GoAssign "__elem" c.headExpr
        , rawGo "_ = __elem"
        , GoMutate "__acc" fw.stepExpr
        ]
    accRet32 = if fw.isAddFold then GoReturn (rawGo "int64(__acc32)") else GoReturn (GoVar "__acc")
    accStep64 =
      [ GoAssign "__elem" c.headExpr
      , rawGo "_ = __elem"
      , GoMutate "__acc" fw.stepExpr
      ]
    body = case c.tailStep, c.predExpr of
      TailFiltered _, Just predGo ->
        let
          pred32 = fromMaybe predGo c.predExpr32
          hoistStmts = map (\d ->
            GoAssign d.varName (GoCall (GoSelector (GoVar "gopurs_runtime") "NewDivisor32")
              [ GoCall (GoVar "int32") [ d.divisorExpr ] ])) c.hoistedDivisors
        in
          GoBlock
            [ GoIfElse (rawGo "__upper < __lower")
                [ GoReturn (GoVar "__seed") ]
                []
            , GoIfElse (rawGo "__lower > -2147483648 && __upper <= 2147483647 && uint64(__upper - __lower) < (1 << 30)")
                ([ rawGo "__count := int(__upper - __lower + 1)"
                 , rawGo "var __stack32 [512]int32"
                 , rawGo "var __buf32 []int32"
                 , GoIfElse (rawGo "__count <= len(__stack32)")
                     [ rawGo "__buf32 = __stack32[:__count]" ]
                     [ rawGo "__buf32 = make([]int32, __count)" ]
                 , rawGo "__v := int32(__lower)"
                 , GoForRange "__i := range __buf32"
                     [ rawGo "__buf32[__i] = __v"
                     , rawGo "__v++"
                     ]
                 ] <> accInit32 <>
                 [ rawGo "__pos := 0"
                 , GoForRange "__pos < len(__buf32)"
                     ([ rawGo "__curVal := int64(__buf32[__pos])"
                      , rawGo "_ = __curVal"
                      ] <> letStmts
                        <> hoistStmts
                        <> accStep32
                        <> [ rawGo "__pos++"
                           , rawGo "__rest := __buf32[__pos:]"
                           , rawGo "__write := 0"
                           , GoForRange "_, __v32 := range __rest"
                               [ rawGo "__x := int64(__v32)"
                               , rawGo "_ = __x"
                               , GoIfElse pred32
                                   [ rawGo "__rest[__write] = __v32"
                                   , rawGo "__write++"
                                   ]
                                   []
                               ]
                           , rawGo "__buf32 = __buf32[:__pos+__write]"
                           ])
                 , accRet32
                 ])
                []
            , rawGo "__buf := make([]int64, 0, 64)"
            , GoForRange "__curr := __upper; !(__curr < __lower); __curr = gopurs_runtime.IntSub(__curr, int64(1))"
                [ rawGo "__buf = append(__buf, __curr)" ]
            , GoForRange "__l, __r := 0, len(__buf)-1; __l < __r; __l, __r = __l+1, __r-1"
                [ rawGo "__buf[__l], __buf[__r] = __buf[__r], __buf[__l]" ]
            , rawGo "__acc := __seed"
            , rawGo "__pos := 0"
            , GoForRange "__pos < len(__buf)"
                ([ rawGo "__curVal := __buf[__pos]"
                 , rawGo "_ = __curVal"
                 ] <> letStmts
                   <> accStep64
                   <> [ rawGo "__pos++"
                      , rawGo "__rest := __buf[__pos:]"
                      , rawGo "__write := 0"
                      , GoForRange "_, __x := range __rest"
                          [ rawGo "_ = __x"
                          , GoIfElse predGo
                              [ rawGo "__rest[__write] = __x"
                              , rawGo "__write++"
                              ]
                              []
                          ]
                      , rawGo "__buf = __buf[:__pos+__write]"
                      ])
            , GoReturn (GoVar "__acc")
            ]
      _, _ ->
        let
          directAccStep32 =
            if fw.isAddFold then
              [ GoAssign "__elem" c.headExpr
              , rawGo "__acc32 += int32(__elem)"
              ]
            else
              accStep64
        in
          GoBlock
            [ GoIfElse (rawGo "__upper < __lower")
                [ GoReturn (GoVar "__seed") ]
                []
            , GoIfElse (rawGo "__lower > -2147483648 && __upper <= 2147483647")
                (accInit32 <>
                 [ GoForRange "__curVal := __lower; __curVal <= __upper; __curVal++"
                     ([ rawGo "_ = __curVal" ] <> letStmts <> directAccStep32)
                 , accRet32
                 ])
                []
            , rawGo "__buf := make([]int64, 0, 64)"
            , GoForRange "__curr := __upper; !(__curr < __lower); __curr = gopurs_runtime.IntSub(__curr, int64(1))"
                [ rawGo "__buf = append(__buf, __curr)" ]
            , GoForRange "__l, __r := 0, len(__buf)-1; __l < __r; __l, __r = __l+1, __r-1"
                [ rawGo "__buf[__l], __buf[__r] = __buf[__r], __buf[__l]" ]
            , rawGo "__acc := __seed"
            , GoForRange "_, __curVal := range __buf"
                ([ rawGo "_ = __curVal" ] <> letStmts <> accStep64)
            , GoReturn (GoVar "__acc")
            ]
  in
    GoFunctionDecl
      { name: generated unit fw.native
      , params: [ Tuple "__lower" TypeInt64, Tuple "__upper" TypeInt64, Tuple "__seed" TypeInt64 ]
      , result: TypeInt64
      , body
      }

-- ---- Range producer & list folder analysis for closed-pipeline deforestation ----

isSimpleInt :: NeutralExpr -> Boolean
isSimpleInt expr = case strip expr of
  Lit (LitInt _) -> true
  Local _ _ -> annotation expr == Just Int
  _ -> false

isLitInt :: Int -> NeutralExpr -> Boolean
isLitInt expected expr = case strip expr of
  Lit (LitInt n) -> n == expected
  _ -> false

isHarmlessIntAdd :: NeutralExpr -> Boolean
isHarmlessIntAdd expr = case strip expr of
  Var (Qualified (Just owner) (Ident name)) -> unwrap owner == "Data.Semiring" && name == "intAdd"
  _ -> false

-- Recognize a descending range builder loop:
--   go curr acc = if curr < limit then acc else go (curr - 1) (Cons curr acc)
-- (with either parameter order for curr/acc and either `curr < limit` or `limit > curr`).
analyzeRangeLoop :: ModuleName -> Array ListShape -> Level -> Ident -> NeutralExpr -> Maybe RangeLoop
analyzeRangeLoop current shapes selfLevel selfName expr = do
  sig <- arrow <$> annotation expr
  shape <- Array.find (\s -> s.ty == sig.result) shapes
  let lambda = abstractions expr
  guard (Array.length lambda.args == 2 && Array.length sig.args == 2)
  { currIndex, accIndex, currRef, accRef } <- case sig.args, lambda.args of
    [ Int, ty ], [ a, b ] | ty == shape.ty ->
      Just { currIndex: 0, accIndex: 1, currRef: a, accRef: b }
    [ ty, Int ], [ a, b ] | ty == shape.ty ->
      Just { currIndex: 1, accIndex: 0, currRef: b, accRef: a }
    _, _ -> Nothing
  guard (snd currRef /= snd accRef && snd currRef /= selfLevel && snd accRef /= selfLevel)
  { cond, step } <- case strip lambda.body of
    Branch choices step -> case NEA.toArray choices of
      [ Pair cond done ] | sameLocal accRef done -> Just { cond, step }
      _ -> Nothing
    _ -> Nothing
  limitExpr <- case strip cond of
    PrimOp (Op2 (OpIntOrd OpLt) lhs rhs) | sameLocal currRef lhs -> Just rhs
    PrimOp (Op2 (OpIntOrd OpGt) lhs rhs) | sameLocal currRef rhs -> Just lhs
    _ -> Nothing
  guard (isSimpleInt limitExpr && not (mentionsLevel selfLevel limitExpr) && not (mentionsLevel (snd currRef) limitExpr) && not (mentionsLevel (snd accRef) limitExpr))
  let recur = spine step
  guard case strip recur.head of
    Local mbName l -> l == selfLevel && (mbName == Nothing || mbName == Just selfName)
    _ -> false
  guard (Array.length recur.args == 2)
  nextCurr <- Array.index recur.args currIndex
  nextAcc <- Array.index recur.args accIndex
  guard case strip nextCurr of
    PrimOp (Op2 (OpIntNum OpSubtract) a b) -> sameLocal currRef a && isLitInt 1 b
    PrimOp (Op2 (OpIntNum OpAdd) a b) -> (sameLocal currRef a && isLitInt (-1) b) || (isLitInt (-1) a && sameLocal currRef b)
    _ -> false
  Tuple head tail <- consFields current shape nextAcc
  guard (sameLocal currRef head && sameLocal accRef tail)
  pure { currIndex, accIndex, shape, limitExpr }

collectTopRangeProducers :: ModuleName -> Array ListShape -> Array (Tuple Ident NeutralExpr) -> Map.Map Ident TopRangeProducer
collectTopRangeProducers current shapes bindings =
  Map.fromFoldable $ Array.mapMaybe (\(Tuple name expr) -> map (Tuple name) (analyzeTopRange current shapes expr)) bindings

analyzeTopRange :: ModuleName -> Array ListShape -> NeutralExpr -> Maybe TopRangeProducer
analyzeTopRange current shapes expr = do
  sig <- arrow <$> annotation expr
  shape <- Array.find (\s -> s.ty == sig.result) shapes
  guard (sig.args == [ Int, Int ])
  let lambda = abstractions expr
  Tuple p0 p1 <- case lambda.args of
    [ a, b ] -> Just (Tuple a b)
    _ -> Nothing
  case strip lambda.body of
    LetRec level bindings body -> do
      let entries = NEA.toArray bindings
      guard (all (not <<< Array.null <<< _.args <<< abstractions <<< snd) entries)
      let call = spine body
      loopName <- case strip call.head of
        Local (Just n) l | l == level -> Just n
        _ -> Nothing
      rhs <- snd <$> Array.find (\(Tuple n _) -> n == loopName) entries
      rl <- analyzeRangeLoop current [ shape ] level loopName rhs
      guard (Array.length call.args == 2)
      initCurr <- Array.index call.args rl.currIndex
      initAcc <- Array.index call.args rl.accIndex
      guard (isEmpty current shape initAcc)
      if sameLocal p0 rl.limitExpr && sameLocal p1 initCurr then
        Just { startIndex: 0, endIndex: 1, shape }
      else if sameLocal p1 rl.limitExpr && sameLocal p0 initCurr then
        Just { startIndex: 1, endIndex: 0, shape }
      else
        Nothing
    _ -> Nothing

-- Recognize a tail-recursive scalar fold loop over a singly-linked Int list:
--   go Nil acc = acc
--   go (Cons x xs) acc = go xs (step acc x)
analyzeFoldLoop :: ModuleName -> Array ListShape -> Level -> Ident -> NeutralExpr -> Maybe FoldLoop
analyzeFoldLoop current shapes selfLevel selfName expr = do
  sig <- arrow <$> annotation expr
  guard (sig.result == Int)
  let lambda = abstractions expr
      arity = Array.length lambda.args
  guard (arity == 2 && Array.length sig.args == 2)
  { shape, listIndex, accIndex, listRef, accRef } <- case sig.args, lambda.args of
    [ ty, Int ], [ a, b ] -> do
      s <- Array.find (\sh -> sh.ty == ty) shapes
      pure { shape: s, listIndex: 0, accIndex: 1, listRef: a, accRef: b }
    [ Int, ty ], [ a, b ] -> do
      s <- Array.find (\sh -> sh.ty == ty) shapes
      pure { shape: s, listIndex: 1, accIndex: 0, listRef: b, accRef: a }
    _, _ -> Nothing
  guard (snd listRef /= snd accRef && snd listRef /= selfLevel && snd accRef /= selfLevel)
  compileFoldBody current shape (\h -> case strip h of
    Local mbName l -> l == selfLevel && (mbName == Nothing || mbName == Just selfName)
    _ -> false) listIndex accIndex 2 listRef accRef lambda.body

compileFoldBody
  :: ModuleName
  -> ListShape
  -> (NeutralExpr -> Boolean)
  -> Int
  -> Int
  -> Int
  -> LocalRef
  -> LocalRef
  -> NeutralExpr
  -> Maybe FoldLoop
compileFoldBody current shape isSelfHead listIndex accIndex arity listRef accRef body = do
  cases <- listCases current shape listRef body
  guard (sameLocal accRef cases.empty)
  let recur = spine cases.cons
  guard (isSelfHead recur.head && Array.length recur.args == arity)
  nextList <- Array.index recur.args listIndex
  nextAcc <- Array.index recur.args accIndex
  guard (isField current shape listRef 1 nextList)
  let stepEnv = Map.singleton (snd accRef) { expr: GoVar "__acc", goType: TypeInt64 }
  stepVal <- compileScalar current shape listRef "__elem" Nothing stepEnv nextAcc
  guard (stepVal.goType == TypeInt64)
  let isAddFold = case strip nextAcc of
        PrimOp (Op2 (OpIntNum OpAdd) a b) ->
          (sameLocal accRef a && isField current shape listRef 0 b)
            || (isField current shape listRef 0 a && sameLocal accRef b)
        _ -> false
      stepKey = printGoExpr stepVal.expr
  pure { listIndex, accIndex, arity, shape, stepExpr: stepVal.expr, isAddFold, stepKey }

collectTopFolders :: ModuleName -> Array ListShape -> Array (Tuple Ident NeutralExpr) -> Map.Map Ident TopFolder
collectTopFolders current shapes bindings =
  Map.fromFoldable $ Array.mapMaybe (\(Tuple name expr) -> map (Tuple name) (analyzeTopFolder current shapes name expr)) bindings

analyzeTopFolder :: ModuleName -> Array ListShape -> Ident -> NeutralExpr -> Maybe TopFolder
analyzeTopFolder current shapes name expr = do
  sig <- arrow <$> annotation expr
  guard (sig.result == Int)
  let lambda = abstractions expr
  guard (Array.length lambda.args == Array.length sig.args)
  case sig.args, lambda.args of
    -- 1-arg wrapper: `sumList lst = let go ... in go lst seed`
    [ ty ], [ listParam ] -> do
      shape <- Array.find (\s -> s.ty == ty) shapes
      case strip lambda.body of
        LetRec level bindings body -> do
          let entries = NEA.toArray bindings
          guard (all (not <<< Array.null <<< _.args <<< abstractions <<< snd) entries)
          let call = spine body
          loopName <- case strip call.head of
            Local (Just n) l | l == level -> Just n
            _ -> Nothing
          rhs <- snd <$> Array.find (\(Tuple n _) -> n == loopName) entries
          fl <- analyzeFoldLoop current [ shape ] level loopName rhs
          guard (Array.length call.args == fl.arity)
          listArg <- Array.index call.args fl.listIndex
          seedArg <- Array.index call.args fl.accIndex
          guard (sameLocal listParam listArg)
          guard case strip seedArg of
            Lit (LitInt _) -> true
            _ -> false
          pure
            { shape
            , listIndex: 0
            , seedSource: Right seedArg
            , arity: 1
            , requireHarmlessAt: Nothing
            , stepExpr: fl.stepExpr
            , isAddFold: fl.isAddFold
            , stepKey: fl.stepKey
            }
        _ -> Nothing
    -- 2-arg direct tail-recursive folder OR 2-arg wrapper around local LetRec folder:
    [ ty, Int ], [ a, b ] -> do
      shape <- Array.find (\s -> s.ty == ty) shapes
      fl <- matchTwoArgFolder current shape name 0 1 a b lambda.body
      pure
        { shape
        , listIndex: 0
        , seedSource: Left 1
        , arity: 2
        , requireHarmlessAt: Nothing
        , stepExpr: fl.stepExpr
        , isAddFold: fl.isAddFold
        , stepKey: fl.stepKey
        }
    [ Int, ty ], [ a, b ] -> do
      shape <- Array.find (\s -> s.ty == ty) shapes
      fl <- matchTwoArgFolder current shape name 1 0 b a lambda.body
      pure
        { shape
        , listIndex: 1
        , seedSource: Left 0
        , arity: 2
        , requireHarmlessAt: Nothing
        , stepExpr: fl.stepExpr
        , isAddFold: fl.isAddFold
        , stepKey: fl.stepKey
        }
    -- 3-arg specialized foldl: `foldl intAdd acc lst = ...`
    [ Func [ Int, Int ] Int, Int, ty ], [ _, accParam, listParam ] -> do
      shape <- Array.find (\s -> s.ty == ty) shapes
      cases <- listCases current shape listParam lambda.body
      guard (sameLocal accParam cases.empty)
      let recur = spine cases.cons
      guard (localGlobal current recur.head == Just name)
      case recur.args of
        [ fnArg, _, nextList ] | isHarmlessIntAdd fnArg && isField current shape listParam 1 nextList -> do
          fl <- compileFoldBody current shape (\h -> localGlobal current h == Just name) 2 1 3 listParam accParam lambda.body
          pure
            { shape
            , listIndex: 2
            , seedSource: Left 1
            , arity: 3
            , requireHarmlessAt: Just 0
            , stepExpr: fl.stepExpr
            , isAddFold: fl.isAddFold
            , stepKey: fl.stepKey
            }
        _ -> Nothing
    _, _ -> Nothing

matchTwoArgFolder
  :: ModuleName
  -> ListShape
  -> Ident
  -> Int
  -> Int
  -> LocalRef
  -> LocalRef
  -> NeutralExpr
  -> Maybe FoldLoop
matchTwoArgFolder current shape name listIndex accIndex listParam seedParam body =
  case compileFoldBody current shape (\h -> localGlobal current h == Just name) listIndex accIndex 2 listParam seedParam body of
    Just fl -> Just fl
    Nothing -> case strip body of
      LetRec level bindings innerBody -> do
        let entries = NEA.toArray bindings
        guard (all (not <<< Array.null <<< _.args <<< abstractions <<< snd) entries)
        let call = spine innerBody
        loopName <- case strip call.head of
          Local (Just n) l | l == level -> Just n
          _ -> Nothing
        rhs <- snd <$> Array.find (\(Tuple n _) -> n == loopName) entries
        fl <- analyzeFoldLoop current [ shape ] level loopName rhs
        guard (Array.length call.args == fl.arity)
        listArg <- Array.index call.args fl.listIndex
        seedArg <- Array.index call.args fl.accIndex
        guard (sameLocal listParam listArg && sameLocal seedParam seedArg)
        pure fl
      _ -> Nothing

-- ---- Fresh list producer analysis & safe call-site rewriting ----

-- Collect top-level functions proven to return a fresh, unaliased list.
collectTopFreshProducers :: ModuleName -> Array ListShape -> Array (Tuple Ident NeutralExpr) -> Map.Map Ident ListShape
collectTopFreshProducers current shapes bindings =
  Map.fromFoldable $ Array.mapMaybe (\(Tuple name expr) -> map (Tuple name) (topFreshProducer current shapes expr)) bindings

topFreshProducer :: ModuleName -> Array ListShape -> NeutralExpr -> Maybe ListShape
topFreshProducer current shapes expr = do
  sig <- arrow <$> annotation expr
  shape <- Array.find (\s -> s.ty == sig.result) shapes
  -- Conservative: no input parameter may itself be a list of this type.
  guard (not (Array.elem shape.ty sig.args))
  let lambda = abstractions expr
  guard (isFreshListExpr current shape Map.empty Map.empty Map.empty lambda.body)
  pure shape

-- Check if a LetRec binding is an accumulator-style fresh list builder loop:
-- all parameters except `accIndex` are non-list types, and every return path
-- returns only `acc`, `Nil`, or newly allocated `Cons` chains ending in `acc`/`Nil`.
analyzeFreshLoop :: ModuleName -> Array ListShape -> Level -> Ident -> NeutralExpr -> Maybe FreshProducer
analyzeFreshLoop current shapes level name expr = do
  sig <- arrow <$> annotation expr
  shape <- Array.find (\s -> s.ty == sig.result) shapes
  let lambda = abstractions expr
      arity = Array.length lambda.args
  guard (arity >= 1 && Array.length sig.args == arity)
  let listIndices = Array.mapMaybe (\(Tuple i ty) -> if ty == shape.ty then Just i else Nothing) (Array.mapWithIndex Tuple sig.args)
  accIndex <- case listIndices of
    [ idx ] -> Just idx
    _ -> Nothing
  accRef <- Array.index lambda.args accIndex
  guard (snd accRef /= level)
  guard (allFreshLoopReturns current shape level name accIndex accRef lambda.body)
  pure { accIndex, arity, shape }

allFreshLoopReturns :: ModuleName -> ListShape -> Level -> Ident -> Int -> LocalRef -> NeutralExpr -> Boolean
allFreshLoopReturns current shape selfLevel selfName accIndex accRef = go
  where
  go expr = case strip expr of
    Branch branches fallback ->
      all (\(Pair cond branch) -> not (mentionsLevel (snd accRef) cond) && go branch) branches
        && go fallback
    Let _ lvl rhs body ->
      not (mentionsLevel (snd accRef) rhs) && lvl /= snd accRef && go body
    _ ->
      isFreshTermFromAcc expr || isFreshSelfCall expr

  isFreshTermFromAcc expr =
    sameLocal accRef expr
      || isEmpty current shape expr
      || case consFields current shape expr of
           Just (Tuple h t) -> not (mentionsLevel (snd accRef) h) && isFreshTermFromAcc t
           Nothing -> false

  isFreshSelfCall expr =
    let call = spine expr
    in case strip call.head of
         Local mbName l | l == selfLevel && (mbName == Nothing || mbName == Just selfName) ->
           case Array.index call.args accIndex of
             Just accArg ->
               isFreshTermFromAcc accArg
                 && all (\(Tuple i a) -> i == accIndex || not (mentionsLevel (snd accRef) a))
                      (Array.mapWithIndex Tuple call.args)
             Nothing -> false
         _ -> false

mentionsLevel :: Level -> NeutralExpr -> Boolean
mentionsLevel target (NeutralExpr syn) = case syn of
  Local _ lvl -> lvl == target
  _ -> any (mentionsLevel target) syn

isFreshListExpr
  :: ModuleName
  -> ListShape
  -> Map.Map Ident ListShape
  -> Map.Map Ident LinearCandidate
  -> Map.Map LocalRef FreshProducer
  -> NeutralExpr
  -> Boolean
isFreshListExpr current shape topProducers candidates localProducers expr =
  if isEmpty current shape expr then true
  else case consFields current shape expr of
    Just (Tuple _ t) -> isFreshListExpr current shape topProducers candidates localProducers t
    Nothing -> case strip expr of
      LetRec level bindings body ->
        let nextLocals = extendLocalProducers current [ shape ] level bindings localProducers
        in isFreshListExpr current shape topProducers candidates nextLocals body
      Let _ _ _ _ -> false
      _ ->
        let call = spine expr
        in case strip call.head of
          Local mbName lvl -> case Map.lookup (Tuple mbName lvl) localProducers of
            Just prod | prod.shape.ty == shape.ty && Array.length call.args == prod.arity ->
              case Array.index call.args prod.accIndex of
                Just accInit -> isFreshListExpr current shape topProducers candidates localProducers accInit
                Nothing -> false
            _ -> false
          Var (Qualified mbMod name) | mbMod == Nothing || mbMod == Just current ->
            case Map.lookup name topProducers of
              Just s | s.ty == shape.ty -> true
              _ -> case Map.lookup name candidates of
                Just cand | cand.shape.ty == shape.ty -> case call.args of
                  [ arg ] -> isFreshListExpr current shape topProducers candidates localProducers arg
                  _ -> false
                _ -> false
          _ -> false

extendLocalProducers
  :: ModuleName
  -> Array ListShape
  -> Level
  -> NEA.NonEmptyArray (Tuple Ident NeutralExpr)
  -> Map.Map LocalRef FreshProducer
  -> Map.Map LocalRef FreshProducer
extendLocalProducers current shapes level bindings localProducers =
  let entries = NEA.toArray bindings
  in if all (not <<< Array.null <<< _.args <<< abstractions <<< snd) entries then
    foldl (\acc (Tuple name rhs) -> case analyzeFreshLoop current shapes level name rhs of
      Just prod -> Map.insert (Tuple (Just name) level) prod acc
      Nothing -> acc) localProducers entries
  else localProducers

extendLocalRanges
  :: ModuleName
  -> Array ListShape
  -> Level
  -> NEA.NonEmptyArray (Tuple Ident NeutralExpr)
  -> Map.Map LocalRef RangeLoop
  -> Map.Map LocalRef RangeLoop
extendLocalRanges current shapes level bindings localRanges =
  let entries = NEA.toArray bindings
  in if all (not <<< Array.null <<< _.args <<< abstractions <<< snd) entries then
    foldl (\acc (Tuple name rhs) -> case analyzeRangeLoop current shapes level name rhs of
      Just rl -> Map.insert (Tuple (Just name) level) rl acc
      Nothing -> acc) localRanges entries
  else localRanges

extendLocalFolders
  :: ModuleName
  -> Array ListShape
  -> Level
  -> NEA.NonEmptyArray (Tuple Ident NeutralExpr)
  -> Map.Map LocalRef FoldLoop
  -> Map.Map LocalRef FoldLoop
extendLocalFolders current shapes level bindings localFolders =
  let entries = NEA.toArray bindings
  in if all (not <<< Array.null <<< _.args <<< abstractions <<< snd) entries then
    foldl (\acc (Tuple name rhs) -> case analyzeFoldLoop current shapes level name rhs of
      Just fl -> Map.insert (Tuple (Just name) level) fl acc
      Nothing -> acc) localFolders entries
  else localFolders

type RewriteEnv =
  { current :: ModuleName
  , unit :: String
  , shapes :: Array ListShape
  , topProducers :: Map.Map Ident ListShape
  , topRanges :: Map.Map Ident TopRangeProducer
  , topFolders :: Map.Map Ident TopFolder
  , candidates :: Map.Map Ident LinearCandidate
  , localProducers :: Map.Map LocalRef FreshProducer
  , localRanges :: Map.Map LocalRef RangeLoop
  , localFolders :: Map.Map LocalRef FoldLoop
  }

matchRangeExpr :: RewriteEnv -> ListShape -> NeutralExpr -> Maybe { lower :: NeutralExpr, upper :: NeutralExpr }
matchRangeExpr env shape expr = case strip expr of
  LetRec level bindings body ->
    let nextRanges = extendLocalRanges env.current [ shape ] level bindings env.localRanges
    in do
      res <- matchRangeExpr (env { localRanges = nextRanges }) shape body
      -- Ensure extracted lower/upper do not reference the dropped LetRec level.
      guard (not (mentionsLevel level res.lower) && not (mentionsLevel level res.upper))
      pure res
  _ ->
    let call = spine expr
    in case strip call.head of
      Local mbName lvl -> do
        rl <- Map.lookup (Tuple mbName lvl) env.localRanges
        guard (rl.shape.ty == shape.ty && Array.length call.args == 2)
        upper <- Array.index call.args rl.currIndex
        accInit <- Array.index call.args rl.accIndex
        guard (isEmpty env.current shape accInit && isSimpleInt upper && isSimpleInt rl.limitExpr)
        pure { lower: rl.limitExpr, upper }
      Var (Qualified mbMod name) | mbMod == Nothing || mbMod == Just env.current -> do
        tr <- Map.lookup name env.topRanges
        guard (tr.shape.ty == shape.ty && Array.length call.args == 2)
        lower <- Array.index call.args tr.startIndex
        upper <- Array.index call.args tr.endIndex
        guard (isSimpleInt lower && isSimpleInt upper)
        pure { lower, upper }
      _ -> Nothing

matchFoldCall
  :: RewriteEnv
  -> NeutralExpr
  -> Maybe
       { shape :: ListShape
       , listExpr :: NeutralExpr
       , seedExpr :: NeutralExpr
       , stepExpr :: GoExpr
       , isAddFold :: Boolean
       , stepKey :: String
       }
matchFoldCall env expr =
  let call = spine expr
  in case strip call.head of
    Local mbName lvl -> do
      fl <- Map.lookup (Tuple mbName lvl) env.localFolders
      guard (Array.length call.args == fl.arity)
      listExpr <- Array.index call.args fl.listIndex
      seedExpr <- Array.index call.args fl.accIndex
      guard (isSimpleInt seedExpr)
      pure { shape: fl.shape, listExpr, seedExpr, stepExpr: fl.stepExpr, isAddFold: fl.isAddFold, stepKey: fl.stepKey }
    Var (Qualified mbMod name) | mbMod == Nothing || mbMod == Just env.current -> do
      tf <- Map.lookup name env.topFolders
      guard (Array.length call.args == tf.arity)
      case tf.requireHarmlessAt of
        Just idx -> do
          fnArg <- Array.index call.args idx
          guard (isHarmlessIntAdd fnArg)
        Nothing -> pure unit
      listExpr <- Array.index call.args tf.listIndex
      seedExpr <- case tf.seedSource of
        Left idx -> Array.index call.args idx
        Right literalSeed -> Just literalSeed
      guard (isSimpleInt seedExpr)
      pure { shape: tf.shape, listExpr, seedExpr, stepExpr: tf.stepExpr, isAddFold: tf.isAddFold, stepKey: tf.stepKey }
    _ -> Nothing

getOrCreateFusedWorker
  :: RewriteEnv
  -> LinearCandidate
  -> GoExpr
  -> Boolean
  -> String
  -> State RewriteState FusedWorker
getOrCreateFusedWorker env cand stepExpr isAddFold stepKey = do
  st <- get
  let key = unwrap cand.original <> "\x00" <> stepKey <> "\x00" <> show isAddFold
  case Map.lookup key st.fusedByKey of
    Just existing -> pure existing
    Nothing -> do
      let worker = WorkerNames.fresh "__gopurs_linear_fold_" cand.original st.names
          native = "Call_" <> env.unit <> "_" <> sanitizeName (unwrap worker)
          fw = { worker, native, candidate: cand, stepExpr, isAddFold }
      put
        { names: WorkerNames.reserve worker st.names
        , fusedByKey: Map.insert key fw st.fusedByKey
        , fusedItems: Array.snoc st.fusedItems fw
        }
      pure fw

rewriteExpr :: RewriteEnv -> NeutralExpr -> State RewriteState NeutralExpr
rewriteExpr env original@(NeutralExpr syn) = case syn of
  LetRec level bindings body -> do
    let entries = NEA.toArray bindings
        allPureLambdas = all (not <<< Array.null <<< _.args <<< abstractions <<< snd) entries
        nextEnv = env
          { localProducers = extendLocalProducers env.current env.shapes level bindings env.localProducers
          , localRanges = extendLocalRanges env.current env.shapes level bindings env.localRanges
          , localFolders = extendLocalFolders env.current env.shapes level bindings env.localFolders
          }
    bindings' <- traverse (\(Tuple n rhs) -> Tuple n <$> rewriteExpr env rhs) bindings
    body' <- rewriteExpr nextEnv body
    if allPureLambdas && mentionsLevel level body && not (mentionsLevel level body') then
      pure body'
    else
      pure (NeutralExpr (LetRec level bindings' body'))
  _ ->
    case matchFusedPipeline env original of
      Just matched -> do
        fw <- getOrCreateFusedWorker env matched.candidate matched.stepExpr matched.isAddFold matched.stepKey
        let sig = Func [ Int, Int, Int ] Int
        pure $ NeutralExpr $ Typed Int $ NeutralExpr $ App
          (NeutralExpr $ Typed sig $ NeutralExpr $ Var $ Qualified (Just env.current) fw.worker)
          (NEA.cons' matched.lower [ matched.upper, matched.seed ])
      Nothing ->
        case matchFreshLinearCall env original of
          Just { candidate, arg } -> do
            arg' <- rewriteExpr env arg
            let sig = Func [ candidate.shape.ty ] candidate.shape.ty
            pure $ NeutralExpr $ Typed candidate.shape.ty $ NeutralExpr $ App
              (NeutralExpr $ Typed sig $ NeutralExpr $ Var $ Qualified (Just env.current) candidate.worker)
              (NEA.singleton arg')
          Nothing ->
            NeutralExpr <$> traverse (rewriteExpr env) syn

matchFusedPipeline
  :: RewriteEnv
  -> NeutralExpr
  -> Maybe
       { candidate :: LinearCandidate
       , lower :: NeutralExpr
       , upper :: NeutralExpr
       , seed :: NeutralExpr
       , stepExpr :: GoExpr
       , isAddFold :: Boolean
       , stepKey :: String
       }
matchFusedPipeline env expr = do
  foldInfo <- matchFoldCall env expr
  let transCall = spine foldInfo.listExpr
  target <- localGlobal env.current transCall.head
  cand <- Map.lookup target env.candidates
  guard (cand.shape.ty == foldInfo.shape.ty)
  rangeArg <- case transCall.args of
    [ a ] -> Just a
    _ -> Nothing
  { lower, upper } <- matchRangeExpr env cand.shape rangeArg
  pure
    { candidate: cand
    , lower
    , upper
    , seed: foldInfo.seedExpr
    , stepExpr: foldInfo.stepExpr
    , isAddFold: foldInfo.isAddFold
    , stepKey: foldInfo.stepKey
    }

matchFreshLinearCall :: RewriteEnv -> NeutralExpr -> Maybe { candidate :: LinearCandidate, arg :: NeutralExpr }
matchFreshLinearCall env expr = do
  let call = spine expr
  target <- localGlobal env.current call.head
  cand <- Map.lookup target env.candidates
  arg <- case call.args of
    [ a ] -> Just a
    _ -> Nothing
  guard (isFreshListExpr env.current cand.shape env.topProducers env.candidates env.localProducers arg)
  pure { candidate: cand, arg }
