module Gopurs.Ownership.Workers (declarations) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (StateT, evalStateT, get, put)
import Control.Monad.Trans.Class (lift)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (any, foldMap, foldM)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst)
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), rawGo, goTypeToStr)
import Gopurs.Ownership.Analysis (continuationPaths, disjoint, leaves, overlap, pathExpr, pathValue, prefix, prefixes, scalar, strip, treeTerm)
import Gopurs.Ownership.Types (Argument(..), Candidate, Candidates, Context, Env, LocalRef, Path(..), TreeSpec, TreeTerm(..), Value(..))
import PureScript.Backend.Optimizer.CoreFn (ModuleName)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Pair(..))

type CellPool = { known :: Array String, nullable :: Array String }
type Generated = { stmts :: Array GoExpr, expr :: GoExpr, pool :: CellPool }
type TakenCell = { stmts :: Array GoExpr, expr :: GoExpr, pool :: CellPool, nonNull :: Boolean }
type Gen = StateT Int Maybe

-- Emission is also the body proof: any unrecognised term, overlapping use or
-- missing consuming callee rejects the complete worker. Each attempt starts
-- with a fresh counter and publishes no partial declarations on failure.
declarations :: CodegenMetadata -> ModuleName -> Candidates -> Candidate -> Maybe (Array GoDecl)
declarations metadata moduleName candidates fn = do
  let context = { metadata, moduleName, candidates, candidate: fn }
      params = Array.mapWithIndex (\index ty -> Tuple ("__arg" <> show index) ty) fn.argTypes
      env = Map.fromFoldable $ Array.zipWith (\ref (Tuple name ty) -> Tuple ref
        (if ty == fn.spec.goType then Tree (Path name []) else Scalar { expr: GoVar name, goType: ty, reads: [] })) fn.args params
  body <- evalStateT (emitBody context env fn.body) 0
  pure
    [ GoFunctionDecl { name: fn.consume, params: params <> [ Tuple "__donor" fn.spec.goType ]
        , result: fn.spec.goType, body: GoFor "__owned_loop" body }
    , GoFunctionDecl { name: fn.native, params, result: fn.spec.goType
        , body: GoReturn (GoCall (GoVar fn.consume) (map (GoVar <<< fst) params <> [ rawGo "nil" ])) }
    ]

freshName :: String -> Gen String
freshName prefix_ = do
  next <- get
  put (next + 1)
  pure (prefix_ <> show next)

-- Capture every path and scalar before allocating/reusing any cell, including
-- reads in nested call arguments. A parent's reuse cannot change a later read
-- of one of its original children.
snapshot :: TreeTerm -> Gen { stmts :: Array GoExpr, term :: TreeTerm }
snapshot = case _ of
  Keep path -> do
    name <- freshName "__read_"
    pure { stmts: [ GoAssign name (pathExpr path) ], term: Existing (GoVar name) }
  Construct args -> do
    result <- traverse snapshotArgument args
    pure { stmts: foldMap _.stmts result, term: Construct (map _.arg result) }
  Call name args -> do
    result <- traverse snapshotArgument args
    pure { stmts: foldMap _.stmts result, term: Call name (map _.arg result) }
  term -> pure { stmts: [], term }

snapshotArgument :: Argument -> Gen { stmts :: Array GoExpr, arg :: Argument }
snapshotArgument = case _ of
  TreeArg value -> do
    result <- snapshot value
    pure { stmts: result.stmts, arg: TreeArg result.term }
  ScalarArg value -> do
    name <- freshName "__scalar_"
    pure { stmts: [ GoAssign name (GoCall (GoVar (goTypeToStr value.goType)) [ value.expr ]) ]
         , arg: ScalarArg (value { expr = GoVar name, reads = [] }) }

-- Only Keep snapshots necessarily dereference their prefixes. Scalar
-- short-circuit expressions may leave deeper paths unevaluated, so they cannot
-- contribute to the known-nonnull pool. The continuation protects every path
-- it can still observe, including through an ancestor or an alias.
reusablePaths :: Env -> Array Path -> Array Path -> { nonNull :: Array Path, dead :: Array Path }
reusablePaths env future retained =
  let roots = Array.mapMaybe (case _ of
        Tree (Path root _) -> Just (Path root [])
        _ -> Nothing) (Array.fromFoldable $ Map.values env)
      nonNull = Array.nub $ foldMap prefixes retained
      candidates = Array.nub $ roots <> nonNull
      dead = Array.filter (\path -> not (any (\kept -> prefix kept path) retained)
        && not (any (overlap path) future)) candidates
  in { nonNull, dead }

planCells :: Env -> Array Path -> TreeTerm -> Gen
  { stmts :: Array GoExpr, term :: TreeTerm, pool :: CellPool, retired :: Array Path }
planCells env future term = do
  let retained = leaves term
  lift $ guard (disjoint retained)
  captured <- snapshot term
  let available = reusablePaths env future retained
  slots <- traverse (\path -> do
    name <- freshName "__dead_"
    pure { name, nonNull: Array.elem path available.nonNull, stmt: GoAssign name (pathExpr path) }) available.dead
  donor <- freshName "__donor_slot_"
  pure { stmts: captured.stmts <> [ GoAssign donor (GoVar "__donor") ] <> map _.stmt slots
       , term: captured.term
       , pool: { known: map _.name (Array.filter _.nonNull slots)
               , nullable: [ donor ] <> map _.name (Array.filter (not <<< _.nonNull) slots) }
       , retired: retained <> available.dead }

-- Known cells were dereferenced by a completed snapshot and are consumed once
-- during emission. Nullable slots still require selection and clearing at runtime.
-- Neither stock contains a cell present in the other one.
takeCell :: CellPool -> Gen TakenCell
takeCell pool = case Array.uncons pool.known of
  Just { head, tail } -> pure
    { stmts: [], expr: GoVar head, pool: pool { known = tail }, nonNull: true }
  Nothing -> do
    name <- freshName "__cell_"
    let takeOne slot rest = [ GoIfElse (GoBinOp "!=" (GoVar slot) (rawGo "nil"))
          [ GoMutate name (GoVar slot), GoMutate slot (rawGo "nil") ] rest ]
    pure { stmts: [ rawGo ("var " <> name <> " " <> "__TREE_TYPE__") ] <> Array.foldr takeOne [] pool.nullable
         , expr: GoVar name, pool, nonNull: false }

-- Substitute only this internal declaration placeholder, never user syntax.
cellStatements :: TreeSpec -> Array GoExpr -> Array GoExpr
cellStatements spec = map replace
  where
  replace (GoRaw code) = rawGo $ String.replaceAll (Pattern "__TREE_TYPE__") (Replacement $ goTypeToStr spec.goType) code.text
  replace expr = expr

emitTree :: Context -> CellPool -> Boolean -> TreeTerm -> Gen Generated
emitTree context pool outer = case _ of
  Existing expr -> pure { stmts: [], expr, pool }
  Empty -> pure { stmts: [], expr: rawGo "nil", pool }
  Keep _ -> lift Nothing
  Construct args -> do
    values <- emitArguments context pool args
    cell <- takeCell values.pool
    name <- case cell.expr of
      GoVar name -> pure name
      _ -> lift Nothing
    pointer <- case context.candidate.spec.goType of
      TypeStructPointer pointer -> pure pointer
      _ -> lift Nothing
    let initialize = if cell.nonNull then [] else
          [ GoIfElse (GoBinOp "==" cell.expr (rawGo "nil"))
              [ GoMutate name (GoCall (GoVar "new") [ GoVar pointer.fullPath ]) ] [] ]
        assign = Array.mapWithIndex (\index value -> GoMutate (name <> ".V" <> show index) value) values.exprs
    pure { stmts: values.stmts <> cellStatements context.candidate.spec cell.stmts
             <> initialize <> [ GoMutate (name <> ".Rc") (GoInt 1) ] <> assign
         , expr: cell.expr, pool: cell.pool }
  Call name args -> do
    fn <- lift $ Map.lookup name context.candidates
    values <- emitArguments context pool args
    donor <- if outer then takeCell values.pool
      else pure { stmts: [], expr: rawGo "nil", pool: values.pool, nonNull: false }
    result <- freshName "__result_"
    pure { stmts: values.stmts <> cellStatements context.candidate.spec donor.stmts
             <> [ GoAssign result (GoCall (GoVar fn.consume) (values.exprs <> [ donor.expr ])) ]
         , expr: GoVar result, pool: donor.pool }

-- Thread the remaining stock across siblings, including constructions inside
-- call arguments. Starting each argument with the original stock would alias cells.
emitArguments :: Context -> CellPool -> Array Argument -> Gen
  { stmts :: Array GoExpr, exprs :: Array GoExpr, pool :: CellPool }
emitArguments context pool = foldM step { stmts: [], exprs: [], pool }
  where
  step result arg = do
    value <- emitArgument context result.pool arg
    pure { stmts: result.stmts <> value.stmts, exprs: Array.snoc result.exprs value.expr, pool: value.pool }

emitArgument :: Context -> CellPool -> Argument -> Gen Generated
emitArgument context pool = case _ of
  TreeArg tree -> emitTree context pool false tree
  ScalarArg value -> pure { stmts: [], expr: value.expr, pool }

retirePaths :: Array Path -> Env -> Env
retirePaths retired = Map.filter (case _ of
  Tree path -> not (any (overlap path) retired)
  Scalar value -> not (any (\path -> any (overlap path) retired) value.reads))

emitBody :: Context -> Env -> NeutralExpr -> Gen (Array GoExpr)
emitBody context env expr = case strip expr of
  Branch branches fallback -> do
    -- Keep the fallback-first name allocation order. Each branch plans its own
    -- cells; a successful guard does not donate cells to another branch.
    def <- emitBody context env fallback
    cases <- traverse (\(Pair condition body) -> do
      cond <- lift $ scalar context env condition
      lift $ guard (cond.goType == TypeBool)
      result <- emitBody context env body
      pure { condition: cond.expr, body: result }) (NEA.toArray branches)
    pure $ Array.foldr (\branch rest -> [ GoIfElse branch.condition branch.body rest ]) def cases
  Fail message -> pure [ GoCall (GoVar "panic") [ GoString message ] ]
  Let name level binding body -> emitLet context env (Tuple name level) binding body
  _ -> emitTerminal context env expr

emitLet :: Context -> Env -> LocalRef -> NeutralExpr -> NeutralExpr -> Gen (Array GoExpr)
emitLet context env ref binding body = case pathValue context env binding of
  Just path -> emitBody context (Map.insert ref (Tree path) env) body
  Nothing -> case scalar context env binding of
    Just value -> do
      temporary <- freshName "__let_scalar_"
      rest <- emitBody context (Map.insert ref
        (Scalar (value { expr = GoVar temporary, reads = [] })) env) body
      pure ([ GoAssign temporary (GoCall (GoVar (goTypeToStr value.goType)) [ value.expr ]) ] <> rest)
    Nothing -> do
      term <- lift $ treeTerm context env binding
      let consumed = leaves term
          future = continuationPaths env body
      lift $ guard (not (any (\path -> any (overlap path) future) consumed))
      prepared <- planCells env future term
      result <- emitTree context prepared.pool true prepared.term
      temporary <- freshName "__let_tree_"
      let nextEnv = Map.insert ref (Tree $ Path temporary []) (retirePaths prepared.retired env)
      rest <- emitBody context nextEnv body
      -- The donated cell may now belong to this result. Retire every old alias
      -- and clear the incoming donor before the continuation builds another plan.
      pure (prepared.stmts <> result.stmts
        <> [ rawGo ("var " <> temporary <> " " <> goTypeToStr context.candidate.spec.goType)
           , GoMutate temporary result.expr, GoMutate "_" (GoVar temporary)
           , GoMutate "__donor" (rawGo "nil") ] <> rest)

emitTerminal :: Context -> Env -> NeutralExpr -> Gen (Array GoExpr)
emitTerminal context env expr = do
  term <- lift $ treeTerm context env expr
  prepared <- planCells env [] term
  case prepared.term of
    Call name args | name == context.candidate.original -> do
      values <- emitArguments context prepared.pool args
      donor <- takeCell values.pool
      let assign = Array.mapWithIndex (\index value -> GoMutate ("__arg" <> show index) value) values.exprs
      pure (prepared.stmts <> values.stmts <> cellStatements context.candidate.spec donor.stmts
        <> assign <> [ GoMutate "__donor" donor.expr, GoContinue "__owned_loop" ])
    _ -> do
      result <- emitTree context prepared.pool true prepared.term
      pure (prepared.stmts <> result.stmts <> [ GoReturn result.expr ])
