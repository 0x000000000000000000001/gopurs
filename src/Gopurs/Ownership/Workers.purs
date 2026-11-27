module Gopurs.Ownership.Workers (declarations) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (StateT, evalStateT, get, put)
import Control.Monad.Trans.Class (lift)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (any, foldMap, foldM, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst)
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), rawGo, goTypeToStr, sanitizeName)
import Gopurs.GoSymbol (GoReference(..), generated)
import Gopurs.Ownership.Analysis (continuationPaths, disjoint, leaves, overlap, pathExpr, pathValue, prefix, prefixes, qualify, scalar, strip, treeTerm)
import Gopurs.Ownership.Continuations (FieldFact, GuardState, emptyGuardState, evalStaticGuard, extractFalseGuardFacts, extractGuardFacts, fieldIndexOf, isExactConstant, mergeGuardState, pathOfGoExpr)
import Gopurs.Ownership.Continuations as Continuations
import Gopurs.Ownership.Types (Argument(..), Candidate, Candidates, Context, Env, LocalRef, Path(..), TreeSpec, TreeTerm(..), Value(..))
import PureScript.Backend.Optimizer.CoreFn (Ident, ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Pair(..))

type KnownCell = { name :: String, origin :: Path }
type CellPool = { known :: Array KnownCell, nullable :: Array String }

data FieldOrigin
  = FromField Path Int
  | FromConst GoExpr
  | UnknownOrigin

derive instance eqFieldOrigin :: Eq FieldOrigin

type Generated =
  { stmts :: Array GoExpr
  , expr :: GoExpr
  , pool :: CellPool
  , origin :: FieldOrigin
  }

type TakenCell = { stmts :: Array GoExpr, expr :: GoExpr, pool :: CellPool }
type Gen = StateT Int Maybe

type InlineSpec =
  { callee :: Candidate
  , label :: String
  , nodeVar :: String
  , body :: Array GoExpr
  }

-- Emission is also the body proof: any unrecognised term, overlapping use or
-- missing consuming callee rejects the complete worker. Each attempt starts
-- with a fresh counter and publishes no partial declarations on failure.
declarations :: CodegenMetadata -> ModuleName -> Candidates -> Candidate -> Maybe (Array GoDecl)
declarations metadata moduleName candidates fn = do
  let context = { metadata, moduleName, candidates, candidate: fn }
      owner = String.replaceAll (Pattern ".") (Replacement "_") (unwrap moduleName)
      params = Array.mapWithIndex (\index ty -> Tuple ("__arg" <> show index) ty) fn.argTypes
      env = Map.fromFoldable $ Array.zipWith (\ref (Tuple name ty) -> Tuple ref
        (if ty == fn.spec.goType then Tree (Path name []) else Scalar { expr: GoVar name, goType: ty, reads: [] })) fn.args params
      inlineSpecs = collectInlineSpecs context
      terminal guards expression = evalStateT (emitBody context inlineSpecs env guards Map.empty expression) 0
      condition expression = do
        value <- scalar context env expression
        guard (value.goType == TypeBool)
        pure value.expr
      allParams = params <> [ Tuple "__donor" fn.spec.goType ]
  rawBody <- case Continuations.shareDetailed "__owned" condition terminal emptyGuardState fn.body of
    Just region ->
      let
        branchBlock (Tuple index stmts) =
          [ rawGo ("__owned_branch_" <> show index <> ": {") ] <> stmts <> [ rawGo "}" ]
        contBlock (Tuple index stmts) =
          [ rawGo ("__owned_continuation_" <> show index <> ": {") ] <> stmts <> [ rawGo "}" ]
      in
        Just (region.tree
          <> foldMap branchBlock (Array.mapWithIndex Tuple region.branches)
          <> foldMap contBlock (Array.mapWithIndex Tuple region.blocks))
    Nothing -> terminal emptyGuardState fn.body
  let usedLabels = collectGotoLabels rawBody
      activeInlines = Array.filter (\spec -> Set.member spec.label usedLabels) (Array.fromFoldable $ Map.values inlineSpecs)
      inlineDecls = map (\spec -> rawGo ("var " <> spec.nodeVar <> " " <> goTypeToStr fn.spec.goType)) activeInlines
      inlineBlocks = foldMap (\spec -> [ rawGo (spec.label <> ": {") ] <> spec.body <> [ rawGo "}" ]) activeInlines
      coreBody = if hasLoopContinue rawBody then [ GoFor "__owned_loop" rawBody ] else rawBody
      consumeBody = GoBlock (inlineDecls <> coreBody <> inlineBlocks)
  pure
    [ GoFunctionDecl { name: generated owner fn.consume, params: allParams
        , result: fn.spec.goType, body: consumeBody }
    , GoFunctionDecl { name: generated owner fn.native, params, result: fn.spec.goType
        , body: GoReturn (GoCall (GoRef (SymbolReference (generated owner fn.consume))) (map (GoVar <<< fst) params <> [ rawGo "nil" ])) }
    ]

-- A non-recursive, constructor-shaped leaf consuming worker (e.g. a node
-- rebalancer whose parameters match the constructor fields) can be inlined
-- once at the end of its caller as a shared labeled block operating directly
-- on the caller's dead parent cell.
collectInlineSpecs :: Context -> Map.Map Ident InlineSpec
collectInlineSpecs context =
  Map.fromFoldable $ Array.mapMaybe buildSpec (Array.fromFoldable $ Map.values context.candidates)
  where
  buildSpec callee = do
    guard (callee.original /= context.candidate.original)
    guard (callee.spec.ty == context.candidate.spec.ty)
    guard (callee.argTypes == context.candidate.spec.fields)
    guard (not (hasCandidateCall context callee.body))
    let suffix = sanitizeName (unwrap callee.original)
        label = "__owned_inline_" <> suffix
        nodeVar = "__owned_inline_node_" <> suffix
        calleeCtx = context { candidate = callee }
        calleeEnv = Map.fromFoldable $ Array.mapWithIndex (\index ref ->
          let ty = fromMaybe TypeValue (Array.index callee.argTypes index)
          in Tuple ref $
            if ty == callee.spec.goType then Tree (Path nodeVar [ index ])
            else Scalar { expr: GoStructAccess (GoVar nodeVar) ("V" <> show index), goType: ty, reads: [ Path nodeVar [] ] })
          callee.args
        initialGuards = { facts: [], nonNullPaths: [ Path nodeVar [] ], nullPaths: [] }
        cond expression = do
          value <- scalar calleeCtx calleeEnv expression
          guard (value.goType == TypeBool)
          pure value.expr
        term guards expression = evalStateT (emitBody calleeCtx Map.empty calleeEnv guards Map.empty expression) 0
    body <- case Continuations.shareDetailed label cond term initialGuards callee.body of
      Just region ->
        let
          branchBlock (Tuple index stmts) =
            [ rawGo (label <> "_branch_" <> show index <> ": {") ] <> stmts <> [ rawGo "}" ]
          contBlock (Tuple index stmts) =
            [ rawGo (label <> "_continuation_" <> show index <> ": {") ] <> stmts <> [ rawGo "}" ]
        in
          Just (region.tree
            <> foldMap branchBlock (Array.mapWithIndex Tuple region.branches)
            <> foldMap contBlock (Array.mapWithIndex Tuple region.blocks))
      Nothing -> term initialGuards callee.body
    guard (not (hasLoopContinue body))
    pure (Tuple callee.original { callee, label, nodeVar, body })

hasCandidateCall :: Context -> NeutralExpr -> Boolean
hasCandidateCall context (NeutralExpr syn) = case syn of
  Var qualified -> case qualify context.moduleName qualified of
    Qualified (Just mod) name -> mod == context.moduleName && Map.member name context.candidates
    _ -> false
  _ -> any (hasCandidateCall context) syn

-- Check whether an expression contains any call to a known owned candidate.
-- Note: we check Var references directly against candidate names in the same module.
hasLoopContinue :: Array GoExpr -> Boolean
hasLoopContinue = any check
  where
  check = case _ of
    GoContinue "__owned_loop" -> true
    GoIfElse _ t f -> hasLoopContinue t || hasLoopContinue f
    GoFor _ stmts -> hasLoopContinue stmts
    GoForRange _ stmts -> hasLoopContinue stmts
    GoBlock stmts -> hasLoopContinue stmts
    _ -> false

collectGotoLabels :: Array GoExpr -> Set.Set String
collectGotoLabels = foldl collect Set.empty
  where
  collect acc = case _ of
    GoRaw code -> case String.stripPrefix (Pattern "goto ") code.text of
      Just lbl -> Set.insert lbl acc
      Nothing -> acc
    GoIfElse _ t f -> foldl collect (foldl collect acc t) f
    GoFor _ stmts -> foldl collect acc stmts
    GoForRange _ stmts -> foldl collect acc stmts
    GoBlock stmts -> foldl collect acc stmts
    _ -> acc

freshName :: String -> Gen String
freshName prefix_ = do
  next <- get
  put (next + 1)
  pure (prefix_ <> show next)

-- Capture every path and scalar before allocating/reusing any cell, including
-- reads in nested call arguments. A parent's reuse cannot change a later read
-- of one of its original children.
snapshot :: Map.Map String FieldOrigin -> TreeTerm -> Gen { stmts :: Array GoExpr, term :: TreeTerm, origins :: Map.Map String FieldOrigin }
snapshot letOrigins = case _ of
  Keep path -> do
    name <- freshName "__read_"
    let orig = case path of
          Path root fields -> case Array.unsnoc fields of
            Just { init, last } -> FromField (Path root init) last
            Nothing -> UnknownOrigin
    pure { stmts: [ GoAssign name (pathExpr path) ], term: Existing (GoVar name), origins: Map.singleton name orig }
  Construct args -> do
    result <- traverse (snapshotArgument letOrigins) args
    pure
      { stmts: foldMap _.stmts result
      , term: Construct (map _.arg result)
      , origins: foldl Map.union Map.empty (map _.origins result)
      }
  Call name args -> do
    result <- traverse (snapshotArgument letOrigins) args
    pure
      { stmts: foldMap _.stmts result
      , term: Call name (map _.arg result)
      , origins: foldl Map.union Map.empty (map _.origins result)
      }
  term -> pure { stmts: [], term, origins: Map.empty }

snapshotArgument :: Map.Map String FieldOrigin -> Argument -> Gen { stmts :: Array GoExpr, arg :: Argument, origins :: Map.Map String FieldOrigin }
snapshotArgument letOrigins = case _ of
  TreeArg value -> do
    result <- snapshot letOrigins value
    pure { stmts: result.stmts, arg: TreeArg result.term, origins: result.origins }
  ScalarArg value -> do
    name <- freshName "__scalar_"
    let orig = scalarOrigin letOrigins value.expr
    pure
      { stmts: [ GoAssign name (GoCall (GoVar (goTypeToStr value.goType)) [ value.expr ]) ]
      , arg: ScalarArg (value { expr = GoVar name, reads = [] })
      , origins: Map.singleton name orig
      }

scalarOrigin :: Map.Map String FieldOrigin -> GoExpr -> FieldOrigin
scalarOrigin letOrigins expr = case expr of
  GoVar name -> fromMaybe UnknownOrigin (Map.lookup name letOrigins)
  GoStructAccess base fieldName ->
    case pathOfGoExpr base, fieldIndexOf fieldName of
      Just path, Just idx -> FromField path idx
      _, _ -> UnknownOrigin
  _ | isExactConstant expr -> FromConst expr
  _ -> UnknownOrigin

-- Only Keep snapshots and dominating branch guards necessarily prove a path
-- non-null. Scalar short-circuit expressions may leave deeper paths unevaluated,
-- so they cannot contribute to the known-nonnull pool. The continuation protects
-- every path it can still observe, including through an ancestor or an alias.
reusablePaths :: Env -> GuardState -> Array Path -> Array Path -> { nonNull :: Array Path, dead :: Array Path }
reusablePaths env guards future retained =
  let roots = Array.mapMaybe (case _ of
        Tree (Path root _) -> Just (Path root [])
        _ -> Nothing) (Array.fromFoldable $ Map.values env)
      guardRoots = map (\(Path root _) -> Path root []) guards.nonNullPaths
      nonNull = Array.nub (foldMap prefixes retained <> guards.nonNullPaths)
      candidates = Array.nub (roots <> guardRoots <> nonNull)
      dead = Array.filter (\path -> not (Array.elem path guards.nullPaths)
        && not (any (\kept -> prefix kept path) retained)
        && not (any (overlap path) future)) candidates
  in { nonNull, dead }

planCells :: Env -> GuardState -> Map.Map String FieldOrigin -> Array Path -> TreeTerm -> Gen
  { stmts :: Array GoExpr, term :: TreeTerm, pool :: CellPool, retired :: Array Path, origins :: Map.Map String FieldOrigin }
planCells env guards letOrigins future term = do
  let retained = leaves term
  lift $ guard (disjoint retained)
  captured <- snapshot letOrigins term
  let available = reusablePaths env guards future retained
  slots <- traverse (\path -> do
    name <- freshName "__dead_"
    pure { name, path, nonNull: Array.elem path available.nonNull, stmt: GoAssign name (pathExpr path) }) available.dead
  donor <- freshName "__donor_slot_"
  pure { stmts: captured.stmts <> [ GoAssign donor (GoVar "__donor") ] <> map _.stmt slots
       , term: captured.term
       , pool: { known: map (\s -> { name: s.name, origin: s.path }) (Array.filter _.nonNull slots)
               , nullable: [ donor ] <> map _.name (Array.filter (not <<< _.nonNull) slots) }
       , retired: retained <> available.dead
       , origins: captured.origins }

-- Known cells were dereferenced by a completed snapshot or dominating guard and
-- are consumed once during emission. Nullable slots still require selection and
-- clearing at runtime. Neither stock contains a cell present in the other one.
takeAnyCell :: CellPool -> Gen TakenCell
takeAnyCell pool = case Array.uncons pool.known of
  Just { head, tail } -> pure
    { stmts: [], expr: GoVar head.name, pool: pool { known = tail } }
  Nothing -> takeNullableCell pool

takeNullableCell :: CellPool -> Gen TakenCell
takeNullableCell pool = case pool.nullable of
  [ slot ] -> do
    name <- freshName "__cell_"
    pure
      { stmts: [ GoAssign name (GoVar slot), GoMutate slot (rawGo "nil") ]
      , expr: GoVar name
      , pool
      }
  _ -> do
    name <- freshName "__cell_"
    let takeOne slot rest = [ GoIfElse (GoBinOp "!=" (GoVar slot) (rawGo "nil"))
          [ GoMutate name (GoVar slot), GoMutate slot (rawGo "nil") ] rest ]
    pure { stmts: [ rawGo ("var " <> name <> " " <> "__TREE_TYPE__") ] <> Array.foldr takeOne [] pool.nullable
         , expr: GoVar name, pool }

-- Substitute only this internal declaration placeholder, never user syntax.
cellStatements :: TreeSpec -> Array GoExpr -> Array GoExpr
cellStatements spec = map replace
  where
  replace (GoRaw code) = rawGo $ String.replaceAll (Pattern "__TREE_TYPE__") (Replacement $ goTypeToStr spec.goType) code.text
  replace expr = expr

hasFact :: Array FieldFact -> Path -> Int -> GoExpr -> Boolean
hasFact facts path idx val = any (\f -> f.path == path && f.field == idx && f.value == val) facts

fieldUnchanged :: Array FieldFact -> Path -> Int -> FieldOrigin -> Boolean
fieldUnchanged facts cellPath idx = case _ of
  FromField p k -> p == cellPath && k == idx
  FromConst c -> hasFact facts cellPath idx c
  UnknownOrigin -> false

pickBestKnownCell
  :: Array GoType
  -> Array FieldFact
  -> Array FieldOrigin
  -> Array KnownCell
  -> Maybe { chosen :: KnownCell, rest :: Array KnownCell }
pickBestKnownCell fieldTypes facts argOrigins known =
  case Array.uncons known of
    Nothing -> Nothing
    Just { head, tail } ->
      let
        scoreCell cell = foldl (+) 0 $ Array.mapWithIndex (\idx orig ->
          if fieldUnchanged facts cell.origin idx orig then
            case Array.index fieldTypes idx of
              Just (TypeStructPointer _) -> 3
              _ -> 2
          else 0) argOrigins
        step best (Tuple idx cell) =
          let s = scoreCell cell
          in if s > best.score then { idx, cell, score: s } else best
        winner = foldl step { idx: 0, cell: head, score: scoreCell head }
          (Array.mapWithIndex (\i c -> Tuple (i + 1) c) tail)
        rest = fromMaybe tail (Array.deleteAt winner.idx known)
      in
        Just { chosen: winner.cell, rest }

emitTree :: Context -> GuardState -> Map.Map String FieldOrigin -> CellPool -> Boolean -> TreeTerm -> Gen Generated
emitTree context guards origins pool outer = case _ of
  Existing expr ->
    let orig = case expr of
          GoVar name -> fromMaybe UnknownOrigin (Map.lookup name origins)
          _ -> UnknownOrigin
    in pure { stmts: [], expr, pool, origin: orig }
  Empty -> pure { stmts: [], expr: rawGo "nil", pool, origin: FromConst (rawGo "nil") }
  Keep _ -> lift Nothing
  Construct args -> do
    values <- emitArguments context guards origins pool args
    pointer <- case context.candidate.spec.goType of
      TypeStructPointer pointer -> pure pointer
      _ -> lift Nothing
    let fieldTypes = context.candidate.spec.fields
    case pickBestKnownCell fieldTypes guards.facts values.origins values.pool.known of
      Just { chosen, rest } ->
        let
          nextPool = values.pool { known = rest }
          fieldWrites = Array.mapMaybe identity $ Array.mapWithIndex (\index (Tuple value orig) ->
            if fieldUnchanged guards.facts chosen.origin index orig then Nothing
            else Just (GoMutate (chosen.name <> ".V" <> show index) value))
            (Array.zip values.exprs values.origins)
          resultOrigin = case chosen.origin of
            Path root fields -> case Array.unsnoc fields of
              Just { init, last } -> FromField (Path root init) last
              Nothing -> UnknownOrigin
        in
          pure { stmts: values.stmts <> fieldWrites
               , expr: GoVar chosen.name
               , pool: nextPool
               , origin: resultOrigin }
      Nothing -> do
        cell <- takeNullableCell values.pool
        name <- case cell.expr of
          GoVar name -> pure name
          _ -> lift Nothing
        let freshLiteral =
              GoMutate name (GoConstructor pointer.baseStructName pointer.symbol pointer.typeArgs values.exprs)
            reusedAssign =
              Array.mapWithIndex (\index value -> GoMutate (name <> ".V" <> show index) value) values.exprs
            populate = [ GoIfElse (GoBinOp "==" cell.expr (rawGo "nil")) [ freshLiteral ] reusedAssign ]
        pure { stmts: values.stmts <> cellStatements context.candidate.spec cell.stmts <> populate
             , expr: cell.expr, pool: cell.pool, origin: UnknownOrigin }
  Call name args -> do
    fn <- lift $ Map.lookup name context.candidates
    values <- emitArguments context guards origins pool args
    donor <- if outer then takeAnyCell values.pool
      else pure { stmts: [], expr: rawGo "nil", pool: values.pool }
    result <- freshName "__result_"
    pure { stmts: values.stmts <> cellStatements context.candidate.spec donor.stmts
             <> [ GoAssign result (GoCall (GoVar fn.consume) (values.exprs <> [ donor.expr ])) ]
         , expr: GoVar result, pool: donor.pool, origin: UnknownOrigin }

-- Thread the remaining stock across siblings, including constructions inside
-- call arguments. Starting each argument with the original stock would alias cells.
emitArguments :: Context -> GuardState -> Map.Map String FieldOrigin -> CellPool -> Array Argument -> Gen
  { stmts :: Array GoExpr, exprs :: Array GoExpr, origins :: Array FieldOrigin, pool :: CellPool }
emitArguments context guards origins pool = foldM step { stmts: [], exprs: [], origins: [], pool }
  where
  step result arg = do
    value <- emitArgument context guards origins result.pool arg
    pure
      { stmts: result.stmts <> value.stmts
      , exprs: Array.snoc result.exprs value.expr
      , origins: Array.snoc result.origins value.origin
      , pool: value.pool
      }

emitArgument :: Context -> GuardState -> Map.Map String FieldOrigin -> CellPool -> Argument -> Gen Generated
emitArgument context guards origins pool = case _ of
  TreeArg tree -> emitTree context guards origins pool false tree
  ScalarArg value ->
    let orig = scalarOrigin origins value.expr
    in pure { stmts: [], expr: value.expr, pool, origin: orig }

retirePaths :: Array Path -> Env -> Env
retirePaths retired = Map.filter (case _ of
  Tree path -> not (any (overlap path) retired)
  Scalar value -> not (any (\path -> any (overlap path) retired) value.reads))

retireGuards :: Array Path -> GuardState -> GuardState
retireGuards retired guards =
  { facts: Array.filter (\f -> not (any (overlap f.path) retired)) guards.facts
  , nonNullPaths: Array.filter (\p -> not (any (overlap p) retired)) guards.nonNullPaths
  , nullPaths: Array.filter (\p -> not (any (overlap p) retired)) guards.nullPaths
  }

retireOrigins :: Array Path -> Map.Map String FieldOrigin -> Map.Map String FieldOrigin
retireOrigins retired = Map.filter (case _ of
  FromField path _ -> not (any (overlap path) retired)
  _ -> true)

-- Drop unused snapshot/dead-slot bindings whose variable name never appears in
-- the remaining statements. Safe whenever the pruned binding only reads a path
-- already proven non-null (or a local variable).
pruneUnusedBindings :: Array GoExpr -> Array GoExpr -> Array GoExpr
pruneUnusedBindings preamble body =
  foldl keep [] (Array.reverse preamble) # \kept -> kept <> body
  where
  keep acc stmt = case stmt of
    GoAssign name rhs ->
      let usedInTail = Set.member name (collectReferencedVars (acc <> body))
      in if usedInTail then Array.cons stmt acc
         else if isSafeToDrop rhs then acc
         else Array.cons stmt acc
    _ -> Array.cons stmt acc

  isSafeToDrop = case _ of
    GoVar _ -> true
    GoCall (GoVar fn) [ inner ] | fn == "int64" || fn == "uint32" || fn == "bool" || fn == "float64" || fn == "string" ->
      isSafeToDrop inner
    GoRaw _ -> true
    GoInt _ -> true
    GoStructAccess base _ -> isSafeToDrop base
    _ -> false

collectReferencedVars :: Array GoExpr -> Set.Set String
collectReferencedVars = foldl go Set.empty
  where
  go acc = case _ of
    GoVar name -> Set.insert name acc
    GoCall fn args -> foldl go (go acc fn) args
    GoSelector obj _ -> go acc obj
    GoStructAccess obj _ -> go acc obj
    GoRecordAccess obj _ -> go acc obj
    GoRecordFieldAccess obj _ -> go acc obj
    GoRecordAccessStatic obj _ _ -> go acc obj
    GoConstructor _ _ _ args -> foldl go acc args
    GoConstructorDict _ args -> foldl go acc args
    GoConstructorAccess obj _ _ _ _ -> go acc obj
    GoBinOp _ l r -> go (go acc l) r
    GoPrefixOp _ inner -> go acc inner
    GoTypeAssertion inner _ -> go acc inner
    GoIndex a i -> go (go acc a) i
    GoReturn inner -> go acc inner
    GoAssign _ rhs -> go acc rhs
    GoMutate target rhs ->
      let baseTarget = fromMaybe target (Array.head (String.split (Pattern ".") target))
      in go (Set.insert baseTarget acc) rhs
    GoIfElse c t f -> foldl go (foldl go (go acc c) t) f
    GoFor _ stmts -> foldl go acc stmts
    GoForRange _ stmts -> foldl go acc stmts
    GoBlock stmts -> foldl go acc stmts
    GoBoxStructPointer _ inner -> go acc inner
    GoBoxIntArray inner -> go acc inner
    GoUnboxIntArray inner -> go acc inner
    GoFreshFilterArray inner -> go acc inner
    GoStructValue _ _ args -> foldl go acc args
    _ -> acc

emitBody :: Context -> Map.Map Ident InlineSpec -> Env -> GuardState -> Map.Map String FieldOrigin -> NeutralExpr -> Gen (Array GoExpr)
emitBody context inlines env guards letOrigins expr = case strip expr of
  Branch branches fallback ->
    emitBranches context inlines env guards letOrigins (NEA.toArray branches) fallback
  Fail message -> pure [ GoCall (GoVar "panic") [ GoString message ] ]
  Let name level binding body -> emitLet context inlines env guards letOrigins (Tuple name level) binding body
  _ -> emitTerminal context inlines env guards letOrigins expr

emitBranches :: Context -> Map.Map Ident InlineSpec -> Env -> GuardState -> Map.Map String FieldOrigin -> Array (Pair NeutralExpr) -> NeutralExpr -> Gen (Array GoExpr)
emitBranches context inlines env guards letOrigins branches fallback = case Array.uncons branches of
  Nothing -> emitBody context inlines env guards letOrigins fallback
  Just { head: Pair condition body, tail } -> do
    cond <- lift $ scalar context env condition
    lift $ guard (cond.goType == TypeBool)
    case evalStaticGuard guards cond.expr of
      Just true ->
        emitBody context inlines env (mergeGuardState guards (extractGuardFacts cond.expr)) letOrigins body
      Just false ->
        emitBranches context inlines env (mergeGuardState guards (extractFalseGuardFacts cond.expr)) letOrigins tail fallback
      Nothing -> do
        rest <- emitBranches context inlines env (mergeGuardState guards (extractFalseGuardFacts cond.expr)) letOrigins tail fallback
        yes <- emitBody context inlines env (mergeGuardState guards (extractGuardFacts cond.expr)) letOrigins body
        pure [ GoIfElse cond.expr yes rest ]

emitLet :: Context -> Map.Map Ident InlineSpec -> Env -> GuardState -> Map.Map String FieldOrigin -> LocalRef -> NeutralExpr -> NeutralExpr -> Gen (Array GoExpr)
emitLet context inlines env guards letOrigins ref binding body = case pathValue context env binding of
  Just path -> emitBody context inlines (Map.insert ref (Tree path) env) guards letOrigins body
  Nothing -> case scalar context env binding of
    Just value -> do
      temporary <- freshName "__let_scalar_"
      let nextOrigins = Map.insert temporary (scalarOrigin letOrigins value.expr) letOrigins
      rest <- emitBody context inlines (Map.insert ref
        (Scalar (value { expr = GoVar temporary, reads = [] })) env) guards nextOrigins body
      let assign = GoAssign temporary (GoCall (GoVar (goTypeToStr value.goType)) [ value.expr ])
      pure (pruneUnusedBindings [ assign ] rest)
    Nothing -> do
      term <- lift $ treeTerm context env binding
      let consumed = leaves term
          future = continuationPaths env body
      lift $ guard (not (any (\path -> any (overlap path) future) consumed))
      prepared <- planCells env guards letOrigins future term
      result <- emitTree context guards prepared.origins prepared.pool true prepared.term
      temporary <- freshName "__let_tree_"
      let nextEnv = Map.insert ref (Tree $ Path temporary []) (retirePaths prepared.retired env)
          nextGuards = retireGuards prepared.retired guards
          nextOrigins = retireOrigins prepared.retired letOrigins
      rest <- emitBody context inlines nextEnv nextGuards nextOrigins body
      -- The donated cell may now belong to this result. Retire every old alias
      -- and clear the incoming donor before the continuation builds another plan.
      let bindTemp = case result.expr of
            GoRaw { text: "nil" } ->
              [ rawGo ("var " <> temporary <> " " <> goTypeToStr context.candidate.spec.goType)
              , GoMutate "_" (GoVar temporary) ]
            _ -> [ GoAssign temporary result.expr ]
          bodyStmts = result.stmts <> bindTemp
            <> [ GoMutate "__donor" (rawGo "nil") ] <> rest
      pure (pruneUnusedBindings prepared.stmts bodyStmts)

emitTerminal :: Context -> Map.Map Ident InlineSpec -> Env -> GuardState -> Map.Map String FieldOrigin -> NeutralExpr -> Gen (Array GoExpr)
emitTerminal context inlines env guards letOrigins expr = do
  term <- lift $ treeTerm context env expr
  case term of
    Call name args ->
      case tryInlineLeafCall context inlines env guards letOrigins name args of
        Just gen -> gen
        Nothing -> emitStandardTerminal context env guards letOrigins term
    _ -> emitStandardTerminal context env guards letOrigins term

-- When a tail call targets a non-recursive constructor-shaped leaf worker and
-- we have a known non-null dead parent cell whose fields already match at least
-- one argument, update ONLY the modified fields on that cell (without spilling
-- unchanged fields across inner recursive calls) and jump to the shared inlined
-- copy of the leaf worker.
tryInlineLeafCall
  :: Context
  -> Map.Map Ident InlineSpec
  -> Env
  -> GuardState
  -> Map.Map String FieldOrigin
  -> Ident
  -> Array Argument
  -> Maybe (Gen (Array GoExpr))
tryInlineLeafCall context inlines env guards letOrigins calleeName args = do
  spec <- Map.lookup calleeName inlines
  let retained = leaves (Call calleeName args)
  guard (disjoint retained)
  let available = reusablePaths env guards [] retained
      isUnchanged targetPath idx = case _ of
        TreeArg (Keep p) -> p == case targetPath of Path r fs -> Path r (Array.snoc fs idx)
        ScalarArg s -> scalarOrigin letOrigins s.expr == FromField targetPath idx
        _ -> false
      nonNullDead = Array.filter (\p -> Array.elem p available.nonNull) available.dead
      matchesTarget targetPath =
        any (\(Tuple idx arg) -> isUnchanged targetPath idx arg) (Array.mapWithIndex Tuple args)
  targetPath <- Array.find matchesTarget nonNullDead
  pure do
    -- Reserve targetPath out of the pool so updated arguments cannot recycle it.
    let otherDead = Array.filter (_ /= targetPath) available.dead
    slots <- traverse (\path -> do
      name <- freshName "__dead_"
      pure { name, path, nonNull: Array.elem path available.nonNull, stmt: GoAssign name (pathExpr path) }) otherDead
    donor <- freshName "__donor_slot_"
    let initialPool =
          { known: map (\s -> { name: s.name, origin: s.path }) (Array.filter _.nonNull slots)
          , nullable: [ donor ] <> map _.name (Array.filter (not <<< _.nonNull) slots)
          }
        indexedArgs = Array.mapWithIndex Tuple args
        updatedArgs = Array.filter (\(Tuple idx arg) -> not (isUnchanged targetPath idx arg)) indexedArgs
    -- Snapshot ONLY the updated arguments (unchanged fields stay in targetPath!).
    snapped <- traverse (\(Tuple idx arg) -> do
      s <- snapshotArgument letOrigins arg
      pure { idx, snapped: s }) updatedArgs
    let snapStmts = foldMap (_.snapped >>> _.stmts) snapped
        snapOrigins = foldl Map.union Map.empty (map (_.snapped >>> _.origins) snapped)
        slotStmts = [ GoAssign donor (GoVar "__donor") ] <> map _.stmt slots
    emitted <- foldM (\acc item -> do
      val <- emitArgument context guards snapOrigins acc.pool item.snapped.arg
      pure
        { stmts: acc.stmts <> val.stmts
        , updates: Array.snoc acc.updates (Tuple item.idx val.expr)
        , pool: val.pool
        }) { stmts: [], updates: [], pool: initialPool } snapped
    let targetExpr = pathExpr targetPath
        assignTarget = GoMutate spec.nodeVar targetExpr
        writeField (Tuple idx valExpr) =
          let fieldPath = spec.nodeVar <> ".V" <> show idx
          in case Array.index context.candidate.spec.fields idx of
            Just (TypeStructPointer _) ->
              GoIfElse (GoBinOp "!=" (rawGo fieldPath) valExpr) [ GoMutate fieldPath valExpr ] []
            _ ->
              GoMutate fieldPath valExpr
        bodyStmts = emitted.stmts <> [ assignTarget ] <> map writeField emitted.updates <> [ rawGo ("goto " <> spec.label) ]
    pure (pruneUnusedBindings (snapStmts <> slotStmts) bodyStmts)

emitStandardTerminal :: Context -> Env -> GuardState -> Map.Map String FieldOrigin -> TreeTerm -> Gen (Array GoExpr)
emitStandardTerminal context env guards letOrigins term = do
  prepared <- planCells env guards letOrigins [] term
  case prepared.term of
    Call name args | name == context.candidate.original -> do
      values <- emitArguments context guards prepared.origins prepared.pool args
      donor <- takeAnyCell values.pool
      let assign = Array.mapWithIndex (\index value -> GoMutate ("__arg" <> show index) value) values.exprs
          bodyStmts = values.stmts <> cellStatements context.candidate.spec donor.stmts
            <> assign <> [ GoMutate "__donor" donor.expr, GoContinue "__owned_loop" ]
      pure (pruneUnusedBindings prepared.stmts bodyStmts)
    _ -> do
      result <- emitTree context guards prepared.origins prepared.pool true prepared.term
      pure (pruneUnusedBindings prepared.stmts (result.stmts <> [ GoReturn result.expr ]))
