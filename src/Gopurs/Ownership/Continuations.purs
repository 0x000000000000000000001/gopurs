module Gopurs.Ownership.Continuations
  ( FieldFact
  , GuardState
  , SharedRegion
  , emptyGuardState
  , extractGuardFacts
  , extractFalseGuardFacts
  , evalStaticGuard
  , mergeGuardState
  , pathOfGoExpr
  , fieldIndexOf
  , isExactConstant
  , share
  , shareDetailed
  ) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (State, StateT, get, put, runState, runStateT)
import Control.Monad.Trans.Class (lift)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldMap, foldl)
import Data.Int as Int
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.String.CodeUnits as SCU
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (GoExpr(..), rawGo)
import Gopurs.Ownership.Analysis (prefixes, strip)
import Gopurs.Ownership.Types (Path(..))
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Pair(..))

type FieldFact =
  { path :: Path
  , field :: Int
  , value :: GoExpr
  }

type GuardState =
  { facts :: Array FieldFact
  , nonNullPaths :: Array Path
  , nullPaths :: Array Path
  }

emptyGuardState :: GuardState
emptyGuardState = { facts: [], nonNullPaths: [], nullPaths: [] }

mergeGuardState :: GuardState -> GuardState -> GuardState
mergeGuardState a b =
  { facts: a.facts <> b.facts
  , nonNullPaths: Array.nub (a.nonNullPaths <> b.nonNullPaths)
  , nullPaths: Array.nub (a.nullPaths <> b.nullPaths)
  }

fieldIndexOf :: String -> Maybe Int
fieldIndexOf name = do
  suffix <- String.stripPrefix (String.Pattern "V") name
  idx <- Int.fromString suffix
  guard (idx >= 0 && show idx == suffix)
  pure idx

pathOfGoExpr :: GoExpr -> Maybe Path
pathOfGoExpr = case _ of
  GoVar name -> Just (Path name [])
  GoStructAccess base fieldName -> do
    Path root fields <- pathOfGoExpr base
    idx <- fieldIndexOf fieldName
    pure (Path root (Array.snoc fields idx))
  _ -> Nothing

-- Exact-comparable scalar/nil constants whose Go `==` guarantees bit-identical
-- storage (excludes float64 where -0.0 == +0.0).
isExactConstant :: GoExpr -> Boolean
isExactConstant = case _ of
  GoInt _ -> true
  GoRaw code ->
    code.text == "nil"
      || code.text == "true"
      || code.text == "false"
      || isUnsignedDigits code.text
  GoCall (GoVar fn) [ inner ] ->
    (fn == "int64" || fn == "uint32" || fn == "bool") && isExactConstant inner
  _ -> false
  where
  isUnsignedDigits str =
    let chars = SCU.toCharArray str
    in not (Array.null chars) && Array.all (\c -> c >= '0' && c <= '9') chars

extractGuardFacts :: GoExpr -> GuardState
extractGuardFacts = case _ of
  GoBinOp "&&" lhs rhs -> mergeGuardState (extractGuardFacts lhs) (extractGuardFacts rhs)
  GoBinOp "!=" lhs (GoRaw { text: "nil" }) -> nonNullOf lhs
  GoBinOp "!=" (GoRaw { text: "nil" }) rhs -> nonNullOf rhs
  GoBinOp "==" lhs (GoRaw { text: "nil" }) -> nullOf lhs
  GoBinOp "==" (GoRaw { text: "nil" }) rhs -> nullOf rhs
  GoBinOp "==" lhs rhs -> mergeGuardState (eqFact lhs rhs) (eqFact rhs lhs)
  _ -> emptyGuardState
  where
  nonNullOf expr = case pathOfGoExpr expr of
    Just path -> { facts: [], nonNullPaths: Array.snoc (prefixes path) path, nullPaths: [] }
    Nothing -> emptyGuardState

  nullOf expr = case pathOfGoExpr expr of
    Just path -> { facts: [], nonNullPaths: prefixes path, nullPaths: [ path ] }
    Nothing -> emptyGuardState

  eqFact lhs rhs = case lhs of
    GoStructAccess base fieldName | isExactConstant rhs ->
      case pathOfGoExpr base, fieldIndexOf fieldName of
        Just path, Just idx ->
          { facts: [ { path, field: idx, value: rhs } ]
          , nonNullPaths: Array.snoc (prefixes path) path
          , nullPaths: []
          }
        _, _ -> emptyGuardState
    _ -> emptyGuardState

extractFalseGuardFacts :: GoExpr -> GuardState
extractFalseGuardFacts = case _ of
  GoBinOp "==" lhs (GoRaw { text: "nil" }) -> nonNullOf lhs
  GoBinOp "==" (GoRaw { text: "nil" }) rhs -> nonNullOf rhs
  GoBinOp "!=" lhs (GoRaw { text: "nil" }) -> nullOf lhs
  GoBinOp "!=" (GoRaw { text: "nil" }) rhs -> nullOf rhs
  _ -> emptyGuardState
  where
  nonNullOf expr = case pathOfGoExpr expr of
    Just path -> { facts: [], nonNullPaths: Array.snoc (prefixes path) path, nullPaths: [] }
    Nothing -> emptyGuardState

  nullOf expr = case pathOfGoExpr expr of
    Just path -> { facts: [], nonNullPaths: prefixes path, nullPaths: [ path ] }
    Nothing -> emptyGuardState

evalStaticGuard :: GuardState -> GoExpr -> Maybe Boolean
evalStaticGuard guards = case _ of
  GoBinOp "!=" lhs (GoRaw { text: "nil" }) -> checkNonNull lhs
  GoBinOp "!=" (GoRaw { text: "nil" }) rhs -> checkNonNull rhs
  GoBinOp "==" lhs (GoRaw { text: "nil" }) -> map not (checkNonNull lhs)
  GoBinOp "==" (GoRaw { text: "nil" }) rhs -> map not (checkNonNull rhs)
  _ -> Nothing
  where
  checkNonNull expr = case pathOfGoExpr expr of
    Just path
      | Array.elem path guards.nonNullPaths -> Just true
      | Array.elem path guards.nullPaths -> Just false
    _ -> Nothing

type Shared = { blocks :: Array (Array GoExpr), leaves :: Int, saved :: Int }
type Sharing = StateT Shared Maybe

type SharedRegion =
  { tree :: Array GoExpr
  , branches :: Array (Array GoExpr)
  , blocks :: Array (Array GoExpr)
  }

-- Only split a root decision region: its internal nodes contain conditions,
-- never bindings, writes or calls which can donate cells. The supplied terminal
-- emitter proves every leaf independently and starts its temporary counter at
-- zero, so structurally equal bodies already have identical local names.
-- A label starts a separate block; its variables cannot enter another leaf's
-- scope. No test, read, write, allocation or consuming call moves in time.
share
  :: (NeutralExpr -> Maybe GoExpr)
  -> (NeutralExpr -> Maybe (Array GoExpr))
  -> NeutralExpr
  -> Maybe (Array GoExpr)
share condition terminal original = do
  region <- shareDetailed "__owned" condition (\_ -> terminal) emptyGuardState original
  pure (region.tree
    <> foldMap branchBlock (Array.mapWithIndex Tuple region.branches)
    <> foldMap contBlock (Array.mapWithIndex Tuple region.blocks))
  where
  branchBlock (Tuple index body) = [ rawGo ("__owned_branch_" <> show index <> ": {") ] <> body <> [ rawGo "}" ]
  contBlock (Tuple index body) = [ rawGo ("__owned_continuation_" <> show index <> ": {") ] <> body <> [ rawGo "}" ]

shareDetailed
  :: String
  -> (NeutralExpr -> Maybe GoExpr)
  -> (GuardState -> NeutralExpr -> Maybe (Array GoExpr))
  -> GuardState
  -> NeutralExpr
  -> Maybe SharedRegion
shareDetailed prefix condition terminal initialGuards original = do
  guard (case strip original of
    Branch _ _ -> true
    _ -> false)
  Tuple tree state <- runStateT (visit initialGuards original) { blocks: [], leaves: 0, saved: 0 }
  -- Avoid extra jumps for tiny regions. This is a code-size threshold, not an
  -- ownership shortcut: every candidate leaf was proved before this decision.
  guard (state.leaves >= 4)
  if state.saved >= 32 then do
    let deduped = dedupBranches (prefix <> "_branch_") tree
    pure { tree: deduped.tree, branches: deduped.branches, blocks: state.blocks }
  else do
    let inlinedTree = inlineContinuations state.blocks tree
        deduped = dedupBranches (prefix <> "_branch_") inlinedTree
    guard (not (Array.null deduped.branches))
    pure { tree: deduped.tree, branches: deduped.branches, blocks: [] }
  where
  label index = prefix <> "_continuation_" <> show index

  visit :: GuardState -> NeutralExpr -> Sharing (Array GoExpr)
  visit guards expression = case strip expression of
    Branch branches fallback -> visitBranches guards (NEA.toArray branches) fallback
    _ -> do
      body <- lift (terminal guards expression)
      state <- get
      index <- case Array.elemIndex body state.blocks of
        Just existing -> do
          put (state { leaves = state.leaves + 1, saved = state.saved + Array.length body })
          pure existing
        Nothing -> do
          put (state { leaves = state.leaves + 1, blocks = Array.snoc state.blocks body })
          pure (Array.length state.blocks)
      pure [ rawGo ("goto " <> label index) ]

  visitBranches :: GuardState -> Array (Pair NeutralExpr) -> NeutralExpr -> Sharing (Array GoExpr)
  visitBranches guards branches fallback = case Array.uncons branches of
    Nothing -> visit guards fallback
    Just { head: Pair test body, tail } -> do
      testExpr <- lift (condition test)
      case evalStaticGuard guards testExpr of
        Just true -> visit (mergeGuardState guards (extractGuardFacts testExpr)) body
        Just false -> visitBranches (mergeGuardState guards (extractFalseGuardFacts testExpr)) tail fallback
        Nothing -> do
          yes <- visit (mergeGuardState guards (extractGuardFacts testExpr)) body
          no <- visitBranches (mergeGuardState guards (extractFalseGuardFacts testExpr)) tail fallback
          pure [ GoIfElse testExpr yes no ]

  inlineContinuations :: Array (Array GoExpr) -> Array GoExpr -> Array GoExpr
  inlineContinuations blocks = Array.concatMap replace
    where
    replace = case _ of
      GoIfElse c t f -> [ GoIfElse c (inlineContinuations blocks t) (inlineContinuations blocks f) ]
      GoRaw code ->
        case String.stripPrefix (String.Pattern ("goto " <> prefix <> "_continuation_")) code.text >>= Int.fromString of
          Just idx -> fromMaybe [ GoRaw code ] (Array.index blocks idx)
          Nothing -> [ GoRaw code ]
      other -> [ other ]

-- Factor repeated internal decision subtrees (such as redundant pattern-match
-- fallback trees) into labeled blocks. Every internal node terminates in a goto
-- or return and introduces no bindings, so hoisting preserves exact evaluation order.
dedupBranches :: String -> Array GoExpr -> { tree :: Array GoExpr, branches :: Array (Array GoExpr) }
dedupBranches branchPrefix root =
  let
    counts = countSubtrees root Map.empty
    Tuple tree finalState = runState (rewriteStmts counts root) { indexMap: Map.empty, blocks: [] }
  in
    { tree, branches: finalState.blocks }
  where
  countIfs :: GoExpr -> Int
  countIfs = case _ of
    GoIfElse _ t f -> 1 + foldl (\acc s -> acc + countIfs s) 0 t + foldl (\acc s -> acc + countIfs s) 0 f
    _ -> 0

  countSubtrees :: Array GoExpr -> Map.Map String Int -> Map.Map String Int
  countSubtrees stmts acc = foldl countExpr acc stmts

  countExpr :: Map.Map String Int -> GoExpr -> Map.Map String Int
  countExpr acc expr = case expr of
    GoIfElse _ t f ->
      let
        updated = if countIfs expr >= 2 then
          let key = printGoExpr expr
          in Map.insert key (fromMaybe 0 (Map.lookup key acc) + 1) acc
          else acc
      in countSubtrees f (countSubtrees t updated)
    _ -> acc

  rewriteStmts :: Map.Map String Int -> Array GoExpr -> State { indexMap :: Map.Map String Int, blocks :: Array (Array GoExpr) } (Array GoExpr)
  rewriteStmts counts stmts = Array.concat <$> traverse (rewriteExpr counts) stmts

  rewriteExpr :: Map.Map String Int -> GoExpr -> State { indexMap :: Map.Map String Int, blocks :: Array (Array GoExpr) } (Array GoExpr)
  rewriteExpr counts expr = case expr of
    GoIfElse cond t f ->
      let key = printGoExpr expr
          occurrences = fromMaybe 0 (Map.lookup key counts)
      in if countIfs expr >= 2 && occurrences >= 2 then do
        st <- get
        case Map.lookup key st.indexMap of
          Just idx -> pure [ rawGo ("goto " <> branchPrefix <> show idx) ]
          Nothing -> do
            let idx = Array.length st.blocks
            put (st { indexMap = Map.insert key idx st.indexMap, blocks = Array.snoc st.blocks [] })
            t' <- rewriteStmts counts t
            f' <- rewriteStmts counts f
            st2 <- get
            let updatedBlocks = fromMaybe st2.blocks (Array.updateAt idx [ GoIfElse cond t' f' ] st2.blocks)
            put (st2 { blocks = updatedBlocks })
            pure [ rawGo ("goto " <> branchPrefix <> show idx) ]
        else do
          t' <- rewriteStmts counts t
          f' <- rewriteStmts counts f
          pure [ GoIfElse cond t' f' ]
    other -> pure [ other ]
