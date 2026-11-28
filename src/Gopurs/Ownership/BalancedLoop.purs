-- | Proven-invariant iterative + small-slab specialization for balanced binary
-- | tree builders starting from Empty (`nil`).
-- |
-- | Mechanically verifies a 4-state regular tree automaton (B_OK / R_OK / NO_RR /
-- | ROOT_RR) on any 2-child pointer ADT with an enum tag field, constructor-shaped
-- | rebalancer, single-descent inserter, root-normalizing wrapper, and tail-loop
-- | builder. When invoked with an Empty root (provably B_OK), emits an internal
-- | iterative path-stack + side-only rebalance + early-exit + 512-node small-slab
-- | worker. Public functions and non-Empty callers retain exact general semantics.
module Gopurs.Ownership.BalancedLoop
  ( FastEntry
  , prepare
  ) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, any, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst)
import Gopurs.CodegenState (CodegenMetadata, FunctionInfo)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), goTypeToStr, rawGo, sanitizeName)
import Gopurs.GoSymbol (GoReference(..), generated)
import Gopurs.Ownership.Analysis (leaves, modulePrefix, pathExpr, pathValue, qualify, scalar, strip, treeTerm)
import Gopurs.Ownership.Continuations (GuardState, evalStaticGuard, extractFalseGuardFacts, extractGuardFacts, fieldIndexOf, mergeGuardState, pathOfGoExpr)
import Gopurs.Ownership.Types (Argument(..), Candidate, Candidates, Context, Env, Path(..), TreeSpec, TreeTerm(..), Value(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Pair(..))

type FastEntry =
  { fastWorker :: Ident
  , fastNative :: String
  , treeArgIdx :: Int
  , spec :: TreeSpec
  }

type RotationCase =
  { childSlot :: Int
  , grandSlot :: Int
  , leftChild :: Array Argument
  , rightChild :: Array Argument
  , rootScalars :: Array (Tuple Int GoExpr)
  }

type BalancerSpec =
  { fn :: Candidate
  , leftIdx :: Int
  , rightIdx :: Int
  , tagIdx :: Int
  , guardTag :: GoExpr
  , triggerTag :: GoExpr
  , rotations :: Array RotationCase
  }

data DescentAction
  = DescendLeft
  | DescendRight
  | KeepUnchanged

data DescentTree
  = DescentLeaf DescentAction
  | DescentBranch GoExpr DescentTree DescentTree

type InserterSpec =
  { fn :: Candidate
  , treeArgIdx :: Int
  , scalarIndices :: Array Int
  , leafScalars :: Array (Tuple Int GoExpr)
  , descent :: DescentTree
  }

type NormalizerSpec =
  { fn :: Candidate
  , ins :: InserterSpec
  }

data BuildStep
  = BuildDone Boolean -- true = return Keep tree, false = return nil
  | BuildRecurse (Array GoExpr) (Array GoExpr) -- next build scalar args, insert scalar args

data BuildTree
  = BuildLeaf BuildStep
  | BuildBranch GoExpr BuildTree BuildTree

type BuilderSpec =
  { fn :: Candidate
  , bal :: BalancerSpec
  , norm :: NormalizerSpec
  , treeArgIdx :: Int
  , scalarIndices :: Array Int
  , body :: BuildTree
  }

prepare
  :: CodegenMetadata
  -> ModuleName
  -> Candidates
  -> { declarations :: Array GoDecl
     , functions :: Map.Map String FunctionInfo
     , fastEntries :: Map.Map Ident FastEntry
     }
prepare metadata moduleName accepted =
  let
    owner = modulePrefix moduleName
    allCands = Array.fromFoldable (Map.values accepted)
    balancers = Array.mapMaybe (recognizeBalancer metadata moduleName accepted) allCands
    findFamily buildFn = Array.findMap (\bal -> do
      guard (bal.fn.spec.ty == buildFn.spec.ty)
      ins <- Array.findMap (\c ->
        if c.spec.ty == bal.fn.spec.ty && c.original /= bal.fn.original then
          recognizeInserter metadata moduleName accepted bal c
        else Nothing) allCands
      norm <- Array.findMap (\c ->
        if c.spec.ty == bal.fn.spec.ty && c.original /= bal.fn.original && c.original /= ins.fn.original then
          recognizeNormalizer metadata moduleName accepted bal ins c
        else Nothing) allCands
      guard (buildFn.original /= bal.fn.original && buildFn.original /= ins.fn.original && buildFn.original /= norm.fn.original)
      recognizeBuilder metadata moduleName accepted bal norm buildFn) balancers
    builders = Array.mapMaybe findFamily allCands
    usedNames = Set.fromFoldable $ Array.concatMap (\c -> [ c.native, c.consume ]) allCands
    emitAll = foldl (\acc b ->
      let
        insFastName = freshSymbolName (b.norm.ins.fn.native <> "_fast") acc.used
        buildFastWorker = Ident (unwrap b.fn.worker <> "_fast")
        buildFastName = freshSymbolName ("Call_" <> owner <> "_" <> sanitizeName (unwrap buildFastWorker)) (Set.insert insFastName acc.used)
        nextUsed = Set.insert buildFastName (Set.insert insFastName acc.used)
        decls = emitBuilderFamily owner insFastName buildFastName b
        entry = { fastWorker: buildFastWorker, fastNative: buildFastName, treeArgIdx: b.treeArgIdx, spec: b.fn.spec }
        fnInfo =
          { arrayFold: Nothing
          , codePointFold: Nothing
          , reference: SymbolReference (generated owner buildFastName)
          , fArgs: b.fn.argTypes
          , fRet: b.fn.spec.goType
          , arity: Array.length b.fn.args
          }
      in
        { used: nextUsed
        , declarations: acc.declarations <> decls
        , functions: Map.insert (sanitizeName (unwrap buildFastWorker)) fnInfo acc.functions
        , fastEntries: Map.insert b.fn.original entry acc.fastEntries
        }) { used: usedNames, declarations: [], functions: Map.empty, fastEntries: Map.empty } builders
  in
    { declarations: emitAll.declarations
    , functions: emitAll.functions
    , fastEntries: emitAll.fastEntries
    }

freshSymbolName :: String -> Set.Set String -> String
freshSymbolName base used =
  if not (Set.member base used) then base
  else go 0
  where
  go n =
    let candidate = base <> "_" <> show n
    in if Set.member candidate used then go (n + 1) else candidate

hasCandidateCall :: Context -> NeutralExpr -> Boolean
hasCandidateCall context (NeutralExpr syn) = case syn of
  Var qualified -> case qualify context.moduleName qualified of
    Qualified (Just mod) name -> mod == context.moduleName && Map.member name context.candidates
    _ -> false
  _ -> any (hasCandidateCall context) syn

-- ============================================================================
-- Step A: Recognize a 2-color binary tree rebalancer
-- ============================================================================

recognizeBalancer :: CodegenMetadata -> ModuleName -> Candidates -> Candidate -> Maybe BalancerSpec
recognizeBalancer metadata moduleName candidates fn = do
  let spec = fn.spec
      context = { metadata, moduleName, candidates, candidate: fn }
  guard (isJust spec.leaf)
  guard (fn.argTypes == spec.fields)
  guard (not (hasCandidateCall context fn.body))
  case spec.goType of
    TypeStructPointer _ -> pure unit
    _ -> Nothing
  let treeIndices = Array.mapMaybe (\(Tuple i ty) -> if ty == spec.goType then Just i else Nothing)
        (Array.mapWithIndex Tuple spec.fields)
  Tuple leftIdx rightIdx <- case treeIndices of
    [ l, r ] -> Just (Tuple l r)
    _ -> Nothing
  let env = Map.fromFoldable $ Array.mapWithIndex (\index ref ->
        let ty = fromMaybe TypeValue (Array.index fn.argTypes index)
        in Tuple ref $
          if ty == spec.goType then Tree (Path "p" [ index ])
          else Scalar { expr: GoStructAccess (GoVar "p") ("V" <> show index), goType: ty, reads: [ Path "p" [] ] })
        fn.args
      initialGuards = { facts: [], nonNullPaths: [ Path "p" [] ], nullPaths: [] }
  collected <- collectDecisionLeaves context env initialGuards fn.body
  guard (Array.length collected >= 3)
  let
    -- Classify every leaf as either Identity or a candidate Rotation
    classifyLeaf leaf =
      if isIdentityConstruct context leaf.env spec leaf.expr then
        Just Nothing
      else
        Just <$> extractRotationLeaf context leaf.env spec leftIdx rightIdx leaf.guards leaf.expr
  classified <- traverse classifyLeaf collected
  let rawRotations = Array.mapMaybe identity classified
  guard (not (Array.null rawRotations))
  firstRot <- Array.head rawRotations
  let tagIdx = firstRot.tagIdx
      guardTag = firstRot.guardTag
      triggerTag = firstRot.triggerTag
  guard (all (\r -> r.tagIdx == tagIdx && r.guardTag == guardTag && r.triggerTag == triggerTag) rawRotations)
  -- Multiple fallback branches in a sequential decision tree can reach the same
  -- (childSlot, grandSlot) rotation; require identical rotation bodies and dedup.
  rotations <- dedupRotations (map _.rot rawRotations)
  -- Verify side-independence: every rotation must still fire when the unmodified
  -- sibling is either Empty (`nil`) or a valid trigger-colored node (`R_OK` with
  -- nil children). Rejects shadowed decision trees where a trigger-colored
  -- sibling blocks checking the modified side!
  guard (all (verifySideIndependent context env spec leftIdx rightIdx tagIdx guardTag triggerTag fn.body) rotations)
  pure { fn, leftIdx, rightIdx, tagIdx, guardTag, triggerTag, rotations }

verifySideIndependent
  :: Context
  -> Env
  -> TreeSpec
  -> Int
  -> Int
  -> Int
  -> GoExpr
  -> GoExpr
  -> NeutralExpr
  -> RotationCase
  -> Boolean
verifySideIndependent context env spec leftIdx rightIdx tagIdx guardTag triggerTag body rot =
  let
    otherSlot = if rot.childSlot == leftIdx then rightIdx else leftIdx
    otherGrand = if rot.grandSlot == leftIdx then rightIdx else leftIdx
    baseFacts =
      [ { path: Path "p" [], field: tagIdx, value: guardTag }
      , { path: Path "p" [ rot.childSlot ], field: tagIdx, value: triggerTag }
      , { path: Path "p" [ rot.childSlot, rot.grandSlot ], field: tagIdx, value: triggerTag }
      ]
    baseNonNull = [ Path "p" [], Path "p" [ rot.childSlot ], Path "p" [ rot.childSlot, rot.grandSlot ] ]
    baseNull = [ Path "p" [ rot.childSlot, otherGrand ] ]
    -- Case 1: unmodified sibling is Empty (nil)
    gEmpty = { facts: baseFacts, nonNullPaths: baseNonNull, nullPaths: Array.snoc baseNull (Path "p" [ otherSlot ]) }
    -- Case 2a: unmodified sibling is R_OK (triggerTag with nil children)
    gRedOk =
      { facts: Array.snoc baseFacts { path: Path "p" [ otherSlot ], field: tagIdx, value: triggerTag }
      , nonNullPaths: Array.snoc baseNonNull (Path "p" [ otherSlot ])
      , nullPaths: baseNull <> [ Path "p" [ otherSlot, leftIdx ], Path "p" [ otherSlot, rightIdx ] ]
      }
    -- Case 2b: unmodified sibling is R_OK (triggerTag with guardTag children)
    gRedOkBlackKids =
      { facts: baseFacts <>
          [ { path: Path "p" [ otherSlot ], field: tagIdx, value: triggerTag }
          , { path: Path "p" [ otherSlot, leftIdx ], field: tagIdx, value: guardTag }
          , { path: Path "p" [ otherSlot, rightIdx ], field: tagIdx, value: guardTag }
          ]
      , nonNullPaths: baseNonNull <>
          [ Path "p" [ otherSlot ]
          , Path "p" [ otherSlot, leftIdx ]
          , Path "p" [ otherSlot, rightIdx ]
          ]
      , nullPaths: baseNull
      }
    -- Case 3: unmodified sibling is B_OK (guardTag)
    gBlackOk =
      { facts: Array.snoc baseFacts { path: Path "p" [ otherSlot ], field: tagIdx, value: guardTag }
      , nonNullPaths: Array.snoc baseNonNull (Path "p" [ otherSlot ])
      , nullPaths: baseNull
      }
    checkCase g = case evalBalancedWithFacts context env g body of
      Just leaf -> case extractRotationLeaf context leaf.env spec leftIdx rightIdx g leaf.expr of
        Just found -> sameRotation found.rot rot
        Nothing -> false
      Nothing -> false
  in
    checkCase gEmpty && checkCase gRedOk && checkCase gRedOkBlackKids && checkCase gBlackOk

evalFactGuard :: GuardState -> GoExpr -> Maybe Boolean
evalFactGuard guards expr = case evalStaticGuard guards expr of
  Just b -> Just b
  Nothing -> case expr of
    GoBinOp "&&" l r -> case evalFactGuard guards l of
      Just false -> Just false
      Just true -> evalFactGuard guards r
      Nothing -> case evalFactGuard guards r of
        Just false -> Just false
        _ -> Nothing
    GoBinOp "==" lhs rhs -> checkEq lhs rhs
    _ -> Nothing
  where
  checkEq lhs rhs = case lhs of
    GoStructAccess base fieldName ->
      case pathOfGoExpr base, fieldIndexOf fieldName of
        Just path, Just idx ->
          case Array.find (\f -> f.path == path && f.field == idx) guards.facts of
            Just f -> Just (f.value == rhs)
            Nothing -> Nothing
        _, _ -> Nothing
    _ -> Nothing

evalBalancedWithFacts :: Context -> Env -> GuardState -> NeutralExpr -> Maybe DecisionLeaf
evalBalancedWithFacts context env guards expr = case strip expr of
  Branch branches fallback -> go (NEA.toArray branches) fallback
  Let name level val body ->
    case pathValue context env val of
      Just path -> evalBalancedWithFacts context (Map.insert (Tuple name level) (Tree path) env) guards body
      Nothing -> case scalar context env val of
        Just sc -> evalBalancedWithFacts context (Map.insert (Tuple name level) (Scalar sc) env) guards body
        Nothing -> Nothing
  Fail _ -> Nothing
  _ -> Just { env, guards, expr }
  where
  go branches fallback = case Array.uncons branches of
    Nothing -> evalBalancedWithFacts context env guards fallback
    Just { head: Pair cond yes, tail } -> do
      sc <- scalar context env cond
      case evalFactGuard guards sc.expr of
        Just true -> evalBalancedWithFacts context env guards yes
        Just false -> go tail fallback
        Nothing -> Nothing

dedupRotations :: Array RotationCase -> Maybe (Array RotationCase)
dedupRotations = foldl step (Just [])
  where
  step Nothing _ = Nothing
  step (Just acc) rot =
    case Array.find (\r -> r.childSlot == rot.childSlot && r.grandSlot == rot.grandSlot) acc of
      Nothing -> Just (Array.snoc acc rot)
      Just existing ->
        if sameRotation existing rot then Just acc else Nothing

sameRotation :: RotationCase -> RotationCase -> Boolean
sameRotation a b =
  a.childSlot == b.childSlot
    && a.grandSlot == b.grandSlot
    && a.rootScalars == b.rootScalars
    && Array.length a.leftChild == Array.length b.leftChild
    && Array.length a.rightChild == Array.length b.rightChild
    && all identity (Array.zipWith sameArg a.leftChild b.leftChild)
    && all identity (Array.zipWith sameArg a.rightChild b.rightChild)
  where
  sameArg (TreeArg (Keep p1)) (TreeArg (Keep p2)) = p1 == p2
  sameArg (ScalarArg s1) (ScalarArg s2) = s1.expr == s2.expr && s1.goType == s2.goType
  sameArg _ _ = false

type DecisionLeaf = { env :: Env, guards :: GuardState, expr :: NeutralExpr }

collectDecisionLeaves :: Context -> Env -> GuardState -> NeutralExpr -> Maybe (Array DecisionLeaf)
collectDecisionLeaves context env guards expr = case strip expr of
  Branch branches fallback -> collectBranches (NEA.toArray branches) fallback
  Let name level val body ->
    case pathValue context env val of
      Just path -> collectDecisionLeaves context (Map.insert (Tuple name level) (Tree path) env) guards body
      Nothing -> case scalar context env val of
        Just sc -> collectDecisionLeaves context (Map.insert (Tuple name level) (Scalar sc) env) guards body
        Nothing -> Nothing
  Fail _ -> Nothing
  _ -> Just [ { env, guards, expr } ]
  where
  collectBranches branches fallback = case Array.uncons branches of
    Nothing -> collectDecisionLeaves context env guards fallback
    Just { head: Pair cond body, tail } -> do
      sc <- scalar context env cond
      guard (sc.goType == TypeBool)
      case evalStaticGuard guards sc.expr of
        Just true -> collectDecisionLeaves context env (mergeGuardState guards (extractGuardFacts sc.expr)) body
        Just false -> collectBranches tail fallback
        Nothing -> do
          yes <- collectDecisionLeaves context env (mergeGuardState guards (extractGuardFacts sc.expr)) body
          no <- collectDecisionLeaves context env (mergeGuardState guards (extractFalseGuardFacts sc.expr)) (case NEA.fromArray tail of
            Just rest -> NeutralExpr (Branch rest fallback)
            Nothing -> fallback)
          pure (yes <> no)

isIdentityConstruct :: Context -> Env -> TreeSpec -> NeutralExpr -> Boolean
isIdentityConstruct = isIdentityConstructAt "p"

isIdentityConstructAt :: String -> Context -> Env -> TreeSpec -> NeutralExpr -> Boolean
isIdentityConstructAt rootName context env spec expr = case treeTerm context env expr of
  Just (Construct args) ->
    Array.length args == Array.length spec.fields && all identity (Array.mapWithIndex checkArg args)
  _ -> false
  where
  checkArg idx = case _ of
    TreeArg (Keep p) -> p == Path rootName [ idx ]
    ScalarArg sc -> sc.expr == GoStructAccess (GoVar rootName) ("V" <> show idx)
    _ -> false

unwindLets :: Context -> Env -> NeutralExpr -> Maybe { env :: Env, expr :: NeutralExpr }
unwindLets context env expr = case strip expr of
  Let name level val body ->
    case pathValue context env val of
      Just path -> unwindLets context (Map.insert (Tuple name level) (Tree path) env) body
      Nothing -> case scalar context env val of
        Just sc -> unwindLets context (Map.insert (Tuple name level) (Scalar sc) env) body
        Nothing -> Nothing
  _ -> Just { env, expr }

extractRotationLeaf
  :: Context
  -> Env
  -> TreeSpec
  -> Int
  -> Int
  -> GuardState
  -> NeutralExpr
  -> Maybe { tagIdx :: Int, guardTag :: GoExpr, triggerTag :: GoExpr, rot :: RotationCase }
extractRotationLeaf context env spec leftIdx rightIdx guards expr = do
  -- Find root tag fact: Path "p" [] with field tagIdx of TypeUint32
  rootFact <- Array.find (\f -> f.path == Path "p" [] && Array.index spec.fields f.field == Just TypeUint32) guards.facts
  let tagIdx = rootFact.field
      guardTag = rootFact.value
  -- Find grandchild tag fact: Path "p" [childSlot, grandSlot] with same tagIdx and value /= guardTag.
  -- Using the grandchild fact uniquely selects the active side even when an
  -- earlier failed branch in a sequential decision tree tested the other child.
  grandFact <- Array.find (\f -> f.field == tagIdx && f.value /= guardTag && case f.path of
    Path "p" [ s, g ] -> (s == leftIdx || s == rightIdx) && (g == leftIdx || g == rightIdx)
    _ -> false) guards.facts
  Tuple childSlot grandSlot <- case grandFact.path of
    Path "p" [ s, g ] -> Just (Tuple s g)
    _ -> Nothing
  let triggerTag = grandFact.value
      otherSlot = if childSlot == leftIdx then rightIdx else leftIdx
  -- Confirm child on that same side also has tag == triggerTag
  guard (any (\f -> f.path == Path "p" [ childSlot ] && f.field == tagIdx && f.value == triggerTag) guards.facts)
  -- Verify output Construct is: C(tag=triggerTag, left=C(tag=guardTag, ...), right=C(tag=guardTag, ...))
  rootArgs <- case treeTerm context env expr of
    Just (Construct args) -> Just args
    _ -> Nothing
  guard (argMatchesTag rootArgs tagIdx triggerTag)
  leftArgs <- case Array.index rootArgs leftIdx of
    Just (TreeArg (Construct args)) -> Just args
    _ -> Nothing
  rightArgs <- case Array.index rootArgs rightIdx of
    Just (TreeArg (Construct args)) -> Just args
    _ -> Nothing
  guard (argMatchesTag leftArgs tagIdx guardTag)
  guard (argMatchesTag rightArgs tagIdx guardTag)
  -- Check that every subtree leaf in leftArgs and rightArgs is an opaque Keep
  -- of one of the 4 unmodified subtrees (and all 4 are distinct).
  let otherGrandSlot = if grandSlot == leftIdx then rightIdx else leftIdx
      allowedLeaves =
        [ Path "p" [ otherSlot ]
        , Path "p" [ childSlot, otherGrandSlot ]
        , Path "p" [ childSlot, grandSlot, leftIdx ]
        , Path "p" [ childSlot, grandSlot, rightIdx ]
        ]
      checkChildTree idx arg =
        if idx == leftIdx || idx == rightIdx then
          case arg of
            TreeArg (Keep p) -> Array.elem p allowedLeaves
            _ -> false
        else case arg of
          ScalarArg sc -> all (\rp -> allowedReadPath childSlot grandSlot rp) sc.reads
          _ -> false
  guard (all identity (Array.mapWithIndex checkChildTree leftArgs))
  guard (all identity (Array.mapWithIndex checkChildTree rightArgs))
  let keptLeaves = leaves (Construct rootArgs)
  guard (Array.length keptLeaves == 4 && Array.length (Array.nub keptLeaves) == 4)
  -- Extract root scalar fields (excluding leftIdx, rightIdx)
  rootScalars <- traverse (\(Tuple i arg) -> case arg of
    ScalarArg sc -> do
      guard (all (\rp -> allowedReadPath childSlot grandSlot rp) sc.reads)
      pure (Tuple i sc.expr)
    _ -> Nothing) (Array.filter (\(Tuple i _) -> i /= leftIdx && i /= rightIdx) (Array.mapWithIndex Tuple rootArgs))
  pure
    { tagIdx
    , guardTag
    , triggerTag
    , rot: { childSlot, grandSlot, leftChild: leftArgs, rightChild: rightArgs, rootScalars }
    }

allowedReadPath :: Int -> Int -> Path -> Boolean
allowedReadPath childSlot grandSlot = case _ of
  Path "p" [] -> true
  Path "p" [ s ] -> s == childSlot
  Path "p" [ s, g ] -> s == childSlot && g == grandSlot
  _ -> false

argMatchesTag :: Array Argument -> Int -> GoExpr -> Boolean
argMatchesTag args tagIdx expected = case Array.index args tagIdx of
  Just (ScalarArg sc) -> sc.expr == expected
  _ -> false

-- ============================================================================
-- Step B: Recognize a single-descent recursive inserter calling `bal`
-- ============================================================================

recognizeInserter :: CodegenMetadata -> ModuleName -> Candidates -> BalancerSpec -> Candidate -> Maybe InserterSpec
recognizeInserter metadata moduleName candidates bal fn = do
  let spec = fn.spec
      context = { metadata, moduleName, candidates, candidate: fn }
      treeParamIndices = Array.mapMaybe (\(Tuple i ty) -> if ty == spec.goType then Just i else Nothing)
        (Array.mapWithIndex Tuple fn.argTypes)
  treeArgIdx <- case treeParamIndices of
    [ idx ] -> Just idx
    _ -> Nothing
  let scalarIndices = Array.filter (_ /= treeArgIdx) (Array.range 0 (Array.length fn.argTypes - 1))
      env = Map.fromFoldable $ Array.mapWithIndex (\index ref ->
        let ty = fromMaybe TypeValue (Array.index fn.argTypes index)
        in Tuple ref $
          if index == treeArgIdx then Tree (Path "__cur" [])
          else Scalar { expr: GoVar ("__arg" <> show index), goType: ty, reads: [] })
        fn.args
  -- Body must split on `__cur == nil` vs `__cur != nil`
  { nilExpr, nonNilExpr } <- matchRootNullSplit context env fn.body
  nilUnwound <- unwindLets context env nilExpr
  -- Nil branch: Construct(tag=triggerTag, left=Empty, right=Empty, pure scalars)
  leafArgs <- case treeTerm context nilUnwound.env nilUnwound.expr of
    Just (Construct args) -> Just args
    _ -> Nothing
  guard (argMatchesTag leafArgs bal.tagIdx bal.triggerTag)
  guard (case Array.index leafArgs bal.leftIdx of
    Just (TreeArg Empty) -> true
    _ -> false)
  guard (case Array.index leafArgs bal.rightIdx of
    Just (TreeArg Empty) -> true
    _ -> false)
  leafScalars <- traverse (\(Tuple i arg) -> case arg of
    ScalarArg sc -> do
      guard (Array.null sc.reads)
      pure (Tuple i sc.expr)
    _ -> Nothing) (Array.filter (\(Tuple i _) -> i /= bal.leftIdx && i /= bal.rightIdx) (Array.mapWithIndex Tuple leafArgs))
  descent <- parseDescentTree context env bal fn.original treeArgIdx scalarIndices nonNilExpr
  guard (hasDescentAction DescendLeft descent && hasDescentAction DescendRight descent)
  pure { fn, treeArgIdx, scalarIndices, leafScalars, descent }

isNullTest :: String -> GoExpr -> Maybe Boolean
isNullTest op = case _ of
  GoBinOp found (GoVar "__cur") (GoRaw { text: "nil" }) | found == op -> Just true
  GoBinOp found (GoRaw { text: "nil" }) (GoVar "__cur") | found == op -> Just true
  _ -> Nothing

matchRootNullSplit :: Context -> Env -> NeutralExpr -> Maybe { nilExpr :: NeutralExpr, nonNilExpr :: NeutralExpr }
matchRootNullSplit context env expr = case strip expr of
  Branch branches fallback -> case NEA.toArray branches of
    [ Pair cond yes ] -> do
      sc <- scalar context env cond
      if isJust (isNullTest "==" sc.expr) then Just { nilExpr: yes, nonNilExpr: fallback }
      else if isJust (isNullTest "!=" sc.expr) then Just { nilExpr: fallback, nonNilExpr: yes }
      else Nothing
    [ Pair c1 e1, Pair c2 e2 ] -> do
      sc1 <- scalar context env c1
      sc2 <- scalar context env c2
      if isJust (isNullTest "==" sc1.expr) && isJust (isNullTest "!=" sc2.expr) then
        Just { nilExpr: e1, nonNilExpr: e2 }
      else if isJust (isNullTest "!=" sc1.expr) && isJust (isNullTest "==" sc2.expr) then
        Just { nilExpr: e2, nonNilExpr: e1 }
      else Nothing
    _ -> Nothing
  _ -> Nothing

parseDescentTree
  :: Context
  -> Env
  -> BalancerSpec
  -> Ident
  -> Int
  -> Array Int
  -> NeutralExpr
  -> Maybe DescentTree
parseDescentTree context env bal insName treeArgIdx scalarIndices expr = case strip expr of
  Branch branches fallback -> goBranches (NEA.toArray branches) fallback
  Let name level val body ->
    case pathValue context env val of
      Just path -> parseDescentTree context (Map.insert (Tuple name level) (Tree path) env) bal insName treeArgIdx scalarIndices body
      Nothing -> case scalar context env val of
        Just sc -> parseDescentTree context (Map.insert (Tuple name level) (Scalar sc) env) bal insName treeArgIdx scalarIndices body
        Nothing -> Nothing
  _ -> DescentLeaf <$> parseDescentLeaf
  where
  goBranches branches fallback = case Array.uncons branches of
    Nothing -> parseDescentTree context env bal insName treeArgIdx scalarIndices fallback
    Just { head: Pair cond yes, tail } -> do
      sc <- scalar context env cond
      guard (sc.goType == TypeBool)
      -- Condition may only read __cur's own scalar fields (not dereference children)
      guard (all (\p -> p == Path "__cur" []) sc.reads)
      yesTree <- parseDescentTree context env bal insName treeArgIdx scalarIndices yes
      noTree <- goBranches tail fallback
      pure (DescentBranch sc.expr yesTree noTree)

  parseDescentLeaf =
    if isIdentityConstructAt "__cur" context env bal.fn.spec expr then
      Just KeepUnchanged
    else case treeTerm context env expr of
      Just (Keep (Path "__cur" [])) -> Just KeepUnchanged
      Just (Call callee balArgs) | callee == bal.fn.original && Array.length balArgs == Array.length bal.fn.spec.fields ->
        if isDescendThrough bal.leftIdx balArgs then Just DescendLeft
        else if isDescendThrough bal.rightIdx balArgs then Just DescendRight
        else Nothing
      _ -> Nothing

  isDescendThrough dirIdx balArgs =
    all identity (Array.mapWithIndex (\idx arg ->
      if idx == dirIdx then
        case arg of
          TreeArg (Call recName recArgs) ->
            recName == insName
              && Array.length recArgs == Array.length scalarIndices + 1
              && all identity (Array.mapWithIndex (\ai rArg ->
                if ai == treeArgIdx then
                  case rArg of
                    TreeArg (Keep p) -> p == Path "__cur" [ dirIdx ]
                    _ -> false
                else
                  case rArg of
                    ScalarArg sc -> sc.expr == GoVar ("__arg" <> show ai)
                    _ -> false) recArgs)
          _ -> false
      else if idx == bal.leftIdx || idx == bal.rightIdx then
        case arg of
          TreeArg (Keep p) -> p == Path "__cur" [ idx ]
          _ -> false
      else
        case arg of
          ScalarArg sc -> sc.expr == GoStructAccess (GoVar "__cur") ("V" <> show idx)
          _ -> false) balArgs)

hasDescentAction :: DescentAction -> DescentTree -> Boolean
hasDescentAction target = go
  where
  go (DescentLeaf a) = case target, a of
    DescendLeft, DescendLeft -> true
    DescendRight, DescendRight -> true
    KeepUnchanged, KeepUnchanged -> true
    _, _ -> false
  go (DescentBranch _ y n) = go y || go n

-- ============================================================================
-- Step C: Recognize root-normalizing wrapper (`insert` = `makeBlack . ins`)
-- ============================================================================

recognizeNormalizer
  :: CodegenMetadata
  -> ModuleName
  -> Candidates
  -> BalancerSpec
  -> InserterSpec
  -> Candidate
  -> Maybe NormalizerSpec
recognizeNormalizer metadata moduleName candidates bal ins fn = do
  let context = { metadata, moduleName, candidates, candidate: fn }
  guard (fn.argTypes == ins.fn.argTypes)
  let env = Map.fromFoldable $ Array.mapWithIndex (\index ref ->
        let ty = fromMaybe TypeValue (Array.index fn.argTypes index)
        in Tuple ref $
          if index == ins.treeArgIdx then Tree (Path "__tree" [])
          else Scalar { expr: GoVar ("__arg" <> show index), goType: ty, reads: [] })
        fn.args
      isSameInsCall val = case treeTerm context env val of
        Just (Call name args) | name == ins.fn.original ->
          Array.length args == Array.length fn.argTypes && all identity (Array.mapWithIndex (\i arg ->
            if i == ins.treeArgIdx then
              case arg of
                TreeArg (Keep (Path "__tree" [])) -> true
                _ -> false
            else
              case arg of
                ScalarArg sc -> sc.expr == GoVar ("__arg" <> show i)
                _ -> false) args)
        _ -> false
  case strip fn.body of
    Let name level val body -> do
      guard (isSameInsCall val)
      let env2 = Map.insert (Tuple name level) (Tree (Path "__cur" [])) env
      guard (isMakeGuardTag context env2 bal body)
      pure { fn, ins }
    _ -> case treeTerm context env fn.body of
      Just (Call mbName [ TreeArg (Call insName args) ]) | insName == ins.fn.original -> do
        guard (Array.length args == Array.length fn.argTypes && all identity (Array.mapWithIndex (\i arg ->
          if i == ins.treeArgIdx then
            case arg of
              TreeArg (Keep (Path "__tree" [])) -> true
              _ -> false
          else
            case arg of
              ScalarArg sc -> sc.expr == GoVar ("__arg" <> show i)
              _ -> false) args))
        mbFn <- Map.lookup mbName candidates
        guard (mbFn.argTypes == [ fn.spec.goType ])
        mbArg <- Array.head mbFn.args
        let mbCtx = context { candidate = mbFn }
            mbEnv = Map.singleton mbArg (Tree (Path "__cur" []))
        guard (isMakeGuardTag mbCtx mbEnv bal mbFn.body)
        pure { fn, ins }
      _ -> Nothing

isMakeGuardTag :: Context -> Env -> BalancerSpec -> NeutralExpr -> Boolean
isMakeGuardTag context env bal body = case matchRootNullSplit context env body of
  Just { nilExpr, nonNilExpr } ->
    let
      nilIsEmpty = case unwindLets context env nilExpr of
        Just u -> case treeTerm context u.env u.expr of
          Just Empty -> true
          _ -> false
        Nothing -> false
    in nilIsEmpty && case unwindLets context env nonNilExpr of
      Just u -> case treeTerm context u.env u.expr of
        Just (Construct args) ->
          Array.length args == Array.length bal.fn.spec.fields && all identity (Array.mapWithIndex (\i arg ->
            if i == bal.tagIdx then
              case arg of
                ScalarArg sc -> sc.expr == bal.guardTag
                _ -> false
            else if i == bal.leftIdx || i == bal.rightIdx then
              case arg of
                TreeArg (Keep p) -> p == Path "__cur" [ i ]
                _ -> false
            else
              case arg of
                ScalarArg sc -> sc.expr == GoStructAccess (GoVar "__cur") ("V" <> show i)
                _ -> false) args)
        _ -> false
      Nothing -> false
  Nothing -> false

-- ============================================================================
-- Step D: Recognize tail-loop builder calling `norm` (`buildTree`)
-- ============================================================================

recognizeBuilder
  :: CodegenMetadata
  -> ModuleName
  -> Candidates
  -> BalancerSpec
  -> NormalizerSpec
  -> Candidate
  -> Maybe BuilderSpec
recognizeBuilder metadata moduleName candidates bal norm fn = do
  let spec = fn.spec
      context = { metadata, moduleName, candidates, candidate: fn }
      treeParamIndices = Array.mapMaybe (\(Tuple i ty) -> if ty == spec.goType then Just i else Nothing)
        (Array.mapWithIndex Tuple fn.argTypes)
  treeArgIdx <- case treeParamIndices of
    [ idx ] -> Just idx
    _ -> Nothing
  let scalarIndices = Array.filter (_ /= treeArgIdx) (Array.range 0 (Array.length fn.argTypes - 1))
      env = Map.fromFoldable $ Array.mapWithIndex (\index ref ->
        let ty = fromMaybe TypeValue (Array.index fn.argTypes index)
        in Tuple ref $
          if index == treeArgIdx then Tree (Path "__tree" [])
          else Scalar { expr: GoVar ("__arg" <> show index), goType: ty, reads: [] })
        fn.args
  body <- parseBuildTree context env norm fn.original treeArgIdx scalarIndices fn.body
  guard (hasBuildStep true body && hasBuildStep false body)
  pure { fn, bal, norm, treeArgIdx, scalarIndices, body }

parseBuildTree
  :: Context
  -> Env
  -> NormalizerSpec
  -> Ident
  -> Int
  -> Array Int
  -> NeutralExpr
  -> Maybe BuildTree
parseBuildTree context env norm buildName treeArgIdx scalarIndices expr = case strip expr of
  Branch branches fallback -> goBranches (NEA.toArray branches) fallback
  Let name level val body ->
    case pathValue context env val of
      Just path -> parseBuildTree context (Map.insert (Tuple name level) (Tree path) env) norm buildName treeArgIdx scalarIndices body
      Nothing -> case scalar context env val of
        Just sc | Array.null sc.reads -> parseBuildTree context (Map.insert (Tuple name level) (Scalar sc) env) norm buildName treeArgIdx scalarIndices body
        _ -> Nothing
  _ -> BuildLeaf <$> parseBuildLeaf
  where
  goBranches branches fallback = case Array.uncons branches of
    Nothing -> parseBuildTree context env norm buildName treeArgIdx scalarIndices fallback
    Just { head: Pair cond yes, tail } -> do
      sc <- scalar context env cond
      guard (sc.goType == TypeBool && Array.null sc.reads)
      yesTree <- parseBuildTree context env norm buildName treeArgIdx scalarIndices yes
      noTree <- goBranches tail fallback
      pure (BuildBranch sc.expr yesTree noTree)

  parseBuildLeaf = case treeTerm context env expr of
    Just (Keep (Path "__tree" [])) -> Just (BuildDone true)
    Just Empty -> Just (BuildDone false)
    Just (Call callee args) | callee == buildName && Array.length args == Array.length scalarIndices + 1 -> do
      treeArg <- Array.index args treeArgIdx
      insArgs <- case treeArg of
        TreeArg (Call normName iArgs) | normName == norm.fn.original -> Just iArgs
        _ -> Nothing
      guard (Array.length insArgs == Array.length norm.ins.fn.argTypes)
      guard (case Array.index insArgs norm.ins.treeArgIdx of
        Just (TreeArg (Keep (Path "__tree" []))) -> true
        _ -> false)
      nextScalars <- traverse (\idx -> case Array.index args idx of
        Just (ScalarArg sc) | Array.null sc.reads -> Just sc.expr
        _ -> Nothing) scalarIndices
      insScalars <- traverse (\idx -> case Array.index insArgs idx of
        Just (ScalarArg sc) | Array.null sc.reads -> Just sc.expr
        _ -> Nothing) norm.ins.scalarIndices
      pure (BuildRecurse nextScalars insScalars)
    _ -> Nothing

hasBuildStep :: Boolean -> BuildTree -> Boolean
hasBuildStep wantRecurse = go
  where
  go (BuildLeaf (BuildRecurse _ _)) = wantRecurse
  go (BuildLeaf (BuildDone _)) = not wantRecurse
  go (BuildBranch _ y n) = go y || go n

-- ============================================================================
-- Code Emission: Iterative Path-Stack + Side-Only Balance + Small Slab
-- ============================================================================

emitBuilderFamily :: String -> String -> String -> BuilderSpec -> Array GoDecl
emitBuilderFamily owner insFastName buildFastName b =
  let overflowName = insFastName <> "_overflow"
  in
    [ emitOverflowHelper owner overflowName b.bal b.norm.ins
    , emitFastInserter owner overflowName insFastName b.bal b.norm.ins
    , emitFastBuilder owner overflowName buildFastName b
    ]

emitOverflowHelper :: String -> String -> BalancerSpec -> InserterSpec -> GoDecl
emitOverflowHelper owner overflowName bal ins =
  let
    spec = ins.fn.spec
    ptrTypeStr = goTypeToStr spec.goType
    scalarParams = map (\i -> Tuple ("__arg" <> show i) (fromMaybe TypeValue (Array.index ins.fn.argTypes i))) ins.scalarIndices
    params = scalarParams <>
      [ Tuple "__cur" spec.goType
      , Tuple "__path" (TypeInterface ("[64]" <> ptrTypeStr))
      , Tuple "__dirMask" (TypeInterface "uint64")
      , Tuple "__depth" (TypeInterface "int")
      ]
    insConsumeArgs = Array.mapWithIndex (\i _ ->
      if i == ins.treeArgIdx then GoVar "__cur" else GoVar ("__arg" <> show i)) ins.fn.argTypes
    balConsumeArgs = Array.mapWithIndex (\i _ -> GoStructAccess (GoVar "__p") ("V" <> show i)) spec.fields
    body = GoBlock
      [ GoAssign "__r" (GoCall (GoRef (SymbolReference (generated owner ins.fn.consume))) (insConsumeArgs <> [ rawGo "nil" ]))
      , GoForRange "__depth > 0"
          [ rawGo "__depth--"
          , GoAssign "__p" (rawGo "__path[__depth]")
          , GoIfElse (rawGo "(__dirMask & 1) == 0")
              [ GoIfElse (GoBinOp "!=" (GoStructAccess (GoVar "__p") ("V" <> show bal.leftIdx)) (GoVar "__r"))
                  [ GoMutate ("__p.V" <> show bal.leftIdx) (GoVar "__r") ] []
              ]
              [ GoIfElse (GoBinOp "!=" (GoStructAccess (GoVar "__p") ("V" <> show bal.rightIdx)) (GoVar "__r"))
                  [ GoMutate ("__p.V" <> show bal.rightIdx) (GoVar "__r") ] []
              ]
          , rawGo "__dirMask >>= 1"
          , GoMutate "__r" (GoCall (GoRef (SymbolReference (generated owner bal.fn.consume))) (balConsumeArgs <> [ GoVar "__p" ]))
          ]
      , GoReturn (GoVar "__r")
      ]
  in
    GoFunctionDecl
      { name: generated owner overflowName
      , params
      , result: spec.goType
      , body
      }

emitFastInserter :: String -> String -> String -> BalancerSpec -> InserterSpec -> GoDecl
emitFastInserter owner overflowName insFastName bal ins =
  let
    spec = ins.fn.spec
    ptrTypeStr = goTypeToStr spec.goType
    structTypeStr = fromMaybe ptrTypeStr (String.stripPrefix (Pattern "*") ptrTypeStr)
    scalarParams = map (\i -> Tuple ("__arg" <> show i) (fromMaybe TypeValue (Array.index ins.fn.argTypes i))) ins.scalarIndices
    params = scalarParams <>
      [ Tuple "__root" spec.goType
      , Tuple "__slab" (TypeInterface ("*[]" <> structTypeStr))
      , Tuple "__pos" (TypeInterface "*int")
      ]
    overflowArgs = map (\i -> GoVar ("__arg" <> show i)) ins.scalarIndices
      <> [ GoVar "__cur", GoVar "__path", GoVar "__dirMask", GoVar "__depth" ]
    overflowFallback =
      [ GoReturn (GoCall (GoRef (SymbolReference (generated owner overflowName))) overflowArgs) ]
    emitDescent = case _ of
      DescentLeaf DescendLeft ->
        [ GoMutate "__cur" (GoStructAccess (GoVar "__cur") ("V" <> show bal.leftIdx)) ]
      DescentLeaf DescendRight ->
        [ rawGo "__dirMask |= 1"
        , GoMutate "__cur" (GoStructAccess (GoVar "__cur") ("V" <> show bal.rightIdx))
        ]
      DescentLeaf KeepUnchanged ->
        [ GoReturn (GoVar "__root") ]
      DescentBranch cond yes no ->
        [ GoIfElse cond (emitDescent yes) (emitDescent no) ]
    slabAlloc =
      [ GoIfElse (rawGo "*__pos == len(*__slab)")
          [ rawGo ("*__slab = make([]" <> structTypeStr <> ", 512)")
          , rawGo "*__pos = 0"
          ] []
      , GoAssign "__r" (rawGo "&(*__slab)[*__pos]")
      , rawGo "*__pos++"
      , GoMutate "__r.Rc" (GoInt 1)
      ] <> map (\(Tuple idx valExpr) -> GoMutate ("__r.V" <> show idx) valExpr) ins.leafScalars
    leftRots = Array.filter (\r -> r.childSlot == bal.leftIdx) bal.rotations
    rightRots = Array.filter (\r -> r.childSlot == bal.rightIdx) bal.rotations
    emitSideAscent sideIdx rots =
      [ GoIfElse (GoBinOp "!=" (GoStructAccess (GoVar "__p") ("V" <> show sideIdx)) (GoVar "__r"))
          [ GoMutate ("__p.V" <> show sideIdx) (GoVar "__r") ] []
      , GoIfElse (GoBinOp "==" (GoStructAccess (GoVar "__p") ("V" <> show bal.tagIdx)) bal.guardTag)
          [ GoIfElse (GoBinOp "==" (GoStructAccess (GoVar "__r") ("V" <> show bal.tagIdx)) bal.triggerTag)
              (Array.concatMap (emitRotationCheck bal) rots)
              []
          , GoReturn (GoVar "__root")
          ] []
      , GoMutate "__r" (GoVar "__p")
      ]
    -- Fresh leaf has triggerTag and nil children: attach directly to parent;
    -- if parent has guardTag no rotation is possible so return immediately.
    leafAttach =
      [ GoIfElse (rawGo "__depth == 0")
          [ GoReturn (GoVar "__r") ] []
      , rawGo "__depth--"
      , GoAssign "__p" (rawGo "__path[__depth]")
      , GoIfElse (rawGo "(__dirMask & 1) == 0")
          [ GoMutate ("__p.V" <> show bal.leftIdx) (GoVar "__r") ]
          [ GoMutate ("__p.V" <> show bal.rightIdx) (GoVar "__r") ]
      , rawGo "__dirMask >>= 1"
      , GoIfElse (GoBinOp "==" (GoStructAccess (GoVar "__p") ("V" <> show bal.tagIdx)) bal.guardTag)
          [ GoReturn (GoVar "__root") ] []
      , GoMutate "__r" (GoVar "__p")
      ]
    body = GoBlock
      ([ rawGo ("var __path [64]" <> ptrTypeStr)
       , rawGo "var __dirMask uint64"
       , GoAssign "__depth" (GoInt 0)
       , GoAssign "__cur" (GoVar "__root")
       , GoForRange "__cur != nil"
           ([ GoIfElse (rawGo "uint(__depth) >= uint(len(__path))") overflowFallback []
            , rawGo "__path[__depth] = __cur"
            , rawGo "__depth++"
            , rawGo "__dirMask <<= 1"
            ]
             <> emitDescent ins.descent)
       ] <> slabAlloc <> leafAttach <>
          [ GoForRange "__depth > 0"
              [ rawGo "__depth--"
              , GoMutate "__p" (rawGo "__path[__depth]")
              , GoAssign "__d" (rawGo "__dirMask & 1")
              , rawGo "__dirMask >>= 1"
              , GoIfElse (rawGo "__d == 0")
                  (emitSideAscent bal.leftIdx leftRots)
                  (emitSideAscent bal.rightIdx rightRots)
              ]
          , GoReturn (GoVar "__r")
          ])
  in
    GoFunctionDecl
      { name: generated owner insFastName
      , params
      , result: spec.goType
      , body
      }

emitRotationCheck :: BalancerSpec -> RotationCase -> Array GoExpr
emitRotationCheck bal rot =
  let
    gField = "V" <> show rot.grandSlot
    tagField = "V" <> show bal.tagIdx
    cond = GoBinOp "&&"
      (GoBinOp "!=" (GoStructAccess (GoVar "__r") gField) (rawGo "nil"))
      (GoBinOp "==" (GoStructAccess (GoStructAccess (GoVar "__r") gField) tagField) bal.triggerTag)
  in
    [ GoIfElse cond (emitRotationBody bal rot) [] ]

-- Emit an in-place 3-cell rotation reusing {__p, __r, __g} with affinity
-- matching (minimal field mutations, zero allocations) and `continue` ascent.
emitRotationBody :: BalancerSpec -> RotationCase -> Array GoExpr
emitRotationBody bal rot =
  let
    -- Map paths rooted at "p" to local variables __p, __r, __g:
    -- Path "p" []                         -> __p
    -- Path "p" [rot.childSlot]            -> __r
    -- Path "p" [rot.childSlot, rot.grandSlot] -> __g
    rewriteExpr = case _ of
      GoStructAccess base field ->
        case pathOfGoExpr (GoStructAccess base field) of
          Just (Path "p" [ s ]) | s == rot.childSlot -> GoVar "__r"
          Just (Path "p" [ s, g ]) | s == rot.childSlot && g == rot.grandSlot -> GoVar "__g"
          _ -> GoStructAccess (rewriteExpr base) field
      GoVar "p" -> GoVar "__p"
      GoBinOp op l r -> GoBinOp op (rewriteExpr l) (rewriteExpr r)
      GoPrefixOp op x -> GoPrefixOp op (rewriteExpr x)
      GoCall fn args -> GoCall (rewriteExpr fn) (map rewriteExpr args)
      other -> other
    rewritePath (Path "p" fields) = case Array.uncons fields of
      Just { head: s, tail } | s == rot.childSlot -> case Array.uncons tail of
        Just { head: g, tail: rest } | g == rot.grandSlot -> Path "__g" rest
        _ -> Path "__r" tail
      _ -> Path "__p" fields
    rewritePath p = p
    argExpr = case _ of
      TreeArg (Keep p) -> pathExpr (rewritePath p)
      ScalarArg sc -> rewriteExpr sc.expr
      _ -> rawGo "nil"
    -- Known static facts on the 3 dead cells before mutation
    knownFact varName idx =
      if idx == bal.tagIdx then
        if varName == "__p" then Just bal.guardTag else Just bal.triggerTag
      else if varName == "__p" && idx == rot.childSlot then Just (GoVar "__r")
      else if varName == "__r" && idx == rot.grandSlot then Just (GoVar "__g")
      else Nothing
    -- Score how many fields of a target node already hold their desired value
    -- on cell `cellName` (so we pick the cell needing the fewest writes!)
    cellIsAlready cellName idx expr =
      case expr of
        GoStructAccess (GoVar v) f | v == cellName && f == "V" <> show idx -> true
        _ -> knownFact cellName idx == Just expr
    scoreCell cellName fieldExprs =
      foldl (+) 0 (map (\(Tuple idx e) -> if cellIsAlready cellName idx e then 1 else 0) fieldExprs)
    leftExprs = Array.mapWithIndex (\i a -> Tuple i (argExpr a)) rot.leftChild
    rightExprs = Array.mapWithIndex (\i a -> Tuple i (argExpr a)) rot.rightChild
    -- Pick distinct assignments of {__p, __r, __g} -> {rootCell, leftCell, rightCell}
    -- maximizing unchanged fields (among all 6 permutations of 3 cells!).
    rootBaseExprs = map (\(Tuple i e) -> Tuple i (rewriteExpr e)) rot.rootScalars
    perms =
      [ { root: "__r", left: "__g", right: "__p" }
      , { root: "__r", left: "__p", right: "__g" }
      , { root: "__g", left: "__r", right: "__p" }
      , { root: "__g", left: "__p", right: "__r" }
      , { root: "__p", left: "__r", right: "__g" }
      , { root: "__p", left: "__g", right: "__r" }
      ]
    scorePerm perm =
      let
        rootAll = rootBaseExprs <>
          [ Tuple bal.leftIdx (GoVar perm.left), Tuple bal.rightIdx (GoVar perm.right) ]
      in scoreCell perm.root rootAll + scoreCell perm.left leftExprs + scoreCell perm.right rightExprs
    best = foldl (\acc p -> if scorePerm p > scorePerm acc then p else acc)
      { root: "__r", left: "__g", right: "__p" } perms
    rootExprs = rootBaseExprs <>
      [ Tuple bal.leftIdx (GoVar best.left), Tuple bal.rightIdx (GoVar best.right) ]
    -- Snapshot any value read from a cell field that will be overwritten before
    -- it is read, by snapshotting all non-constant changed field values first.
    plannedWrites cellName exprs =
      Array.mapMaybe (\(Tuple idx e) ->
        if cellIsAlready cellName idx e then Nothing
        else Just { cell: cellName, idx, expr: e }) exprs
    allWrites = plannedWrites best.left leftExprs <> plannedWrites best.right rightExprs <> plannedWrites best.root rootExprs
    -- Snapshot every non-constant write RHS into a temp var before performing writes
    snapsAndWrites = foldl (\acc (Tuple wi w) ->
      case w.expr of
        GoVar _ -> acc { writes = Array.snoc acc.writes (GoMutate (w.cell <> ".V" <> show w.idx) w.expr) }
        GoInt _ -> acc { writes = Array.snoc acc.writes (GoMutate (w.cell <> ".V" <> show w.idx) w.expr) }
        GoRaw _ -> acc { writes = Array.snoc acc.writes (GoMutate (w.cell <> ".V" <> show w.idx) w.expr) }
        _ ->
          let tmp = "__rot_v" <> show wi
          in { snaps: Array.snoc acc.snaps (GoAssign tmp w.expr)
             , writes: Array.snoc acc.writes (GoMutate (w.cell <> ".V" <> show w.idx) (GoVar tmp))
             }) { snaps: [], writes: [] } (Array.mapWithIndex Tuple allWrites)
    assignRoot =
      if best.root == "__r" then []
      else [ GoMutate "__r" (GoVar best.root) ]
  in
    [ GoAssign "__g" (GoStructAccess (GoVar "__r") ("V" <> show rot.grandSlot)) ]
      <> snapsAndWrites.snaps
      <> snapsAndWrites.writes
      <> assignRoot
      <> [ rawGo "continue" ]

renameInsArgs :: GoExpr -> GoExpr
renameInsArgs = case _ of
  GoVar v -> case String.stripPrefix (Pattern "__arg") v of
    Just rest -> GoVar ("__ins_arg_" <> rest)
    Nothing -> GoVar v
  GoStructAccess base f -> GoStructAccess (renameInsArgs base) f
  GoBinOp op l r -> GoBinOp op (renameInsArgs l) (renameInsArgs r)
  GoPrefixOp op x -> GoPrefixOp op (renameInsArgs x)
  GoCall fn args -> GoCall (renameInsArgs fn) (map renameInsArgs args)
  other -> other

emitFastBuilder :: String -> String -> String -> BuilderSpec -> GoDecl
emitFastBuilder owner overflowName buildFastName b =
  let
    spec = b.fn.spec
    bal = b.bal
    ins = b.norm.ins
    ptrTypeStr = goTypeToStr spec.goType
    structTypeStr = fromMaybe ptrTypeStr (String.stripPrefix (Pattern "*") ptrTypeStr)
    params = Array.mapWithIndex (\i ty -> Tuple (if i == b.treeArgIdx then "__tree" else "__arg" <> show i) ty) b.fn.argTypes
    fallbackArgs = map (GoVar <<< fst) params
    insArgDecls = map (\i ->
      let ty = fromMaybe TypeValue (Array.index ins.fn.argTypes i)
      in rawGo ("var __ins_arg_" <> show i <> " " <> goTypeToStr ty)) ins.scalarIndices
    overflowArgs = map (\i -> GoVar ("__ins_arg_" <> show i)) ins.scalarIndices
      <> [ GoVar "__cur", GoVar "__path", GoVar "__dirMask", GoVar "__depth" ]
    leftRots = Array.filter (\r -> r.childSlot == bal.leftIdx) bal.rotations
    rightRots = Array.filter (\r -> r.childSlot == bal.rightIdx) bal.rotations
    emitInlineInsert doneLabel skipNormLabel =
      let
        overflowFallback =
          [ GoMutate "__tree" (GoCall (GoRef (SymbolReference (generated owner overflowName))) overflowArgs)
          , rawGo ("goto " <> doneLabel)
          ]
        emitDescent = case _ of
          DescentLeaf DescendLeft ->
            [ GoMutate "__cur" (GoStructAccess (GoVar "__cur") ("V" <> show bal.leftIdx)) ]
          DescentLeaf DescendRight ->
            [ rawGo "__dirMask |= 1"
            , GoMutate "__cur" (GoStructAccess (GoVar "__cur") ("V" <> show bal.rightIdx))
            ]
          DescentLeaf KeepUnchanged ->
            [ rawGo ("goto " <> skipNormLabel) ]
          DescentBranch cond yes no ->
            [ GoIfElse (renameInsArgs cond) (emitDescent yes) (emitDescent no) ]
        slabAlloc =
          [ GoIfElse (rawGo "__pos == len(__slab)")
              [ rawGo ("__slab = make([]" <> structTypeStr <> ", 512)")
              , rawGo "__pos = 0"
              ] []
          , GoMutate "__r" (rawGo "&__slab[__pos]")
          , rawGo "__pos++"
          , GoMutate "__r.Rc" (GoInt 1)
          ] <> map (\(Tuple idx valExpr) -> GoMutate ("__r.V" <> show idx) (renameInsArgs valExpr)) ins.leafScalars
        emitSideAscent sideIdx rots =
          [ GoIfElse (GoBinOp "!=" (GoStructAccess (GoVar "__p") ("V" <> show sideIdx)) (GoVar "__r"))
              [ GoMutate ("__p.V" <> show sideIdx) (GoVar "__r") ] []
          , GoIfElse (GoBinOp "==" (GoStructAccess (GoVar "__p") ("V" <> show bal.tagIdx)) bal.guardTag)
              [ GoIfElse (GoBinOp "==" (GoStructAccess (GoVar "__r") ("V" <> show bal.tagIdx)) bal.triggerTag)
                  (Array.concatMap (emitRotationCheck bal) rots)
                  []
              , rawGo ("goto " <> skipNormLabel)
              ] []
          , GoMutate "__r" (GoVar "__p")
          ]
        leafAttach =
          [ GoIfElse (rawGo "__depth == 0")
              [ GoMutate "__tree" (GoVar "__r")
              , GoMutate ("__tree.V" <> show bal.tagIdx) bal.guardTag
              , rawGo ("goto " <> skipNormLabel)
              ] []
          , rawGo "__depth--"
          , GoMutate "__p" (rawGo "__path[__depth]")
          , GoIfElse (rawGo "(__dirMask & 1) == 0")
              [ GoMutate ("__p.V" <> show bal.leftIdx) (GoVar "__r") ]
              [ GoMutate ("__p.V" <> show bal.rightIdx) (GoVar "__r") ]
          , rawGo "__dirMask >>= 1"
          , GoIfElse (GoBinOp "==" (GoStructAccess (GoVar "__p") ("V" <> show bal.tagIdx)) bal.guardTag)
              [ rawGo ("goto " <> skipNormLabel) ] []
          , GoMutate "__r" (GoVar "__p")
          ]
      in
        [ GoMutate "__depth" (GoInt 0)
        , GoMutate "__dirMask" (GoInt 0)
        , GoMutate "__cur" (GoVar "__tree")
        , GoForRange "__cur != nil"
            ([ GoIfElse (rawGo "uint(__depth) >= uint(len(__path))") overflowFallback []
             , rawGo "__path[__depth] = __cur"
             , rawGo "__depth++"
             , rawGo "__dirMask <<= 1"
             ]
              <> emitDescent ins.descent)
        ] <> slabAlloc <> leafAttach <>
          [ GoForRange "__depth > 0"
              [ rawGo "__depth--"
              , GoMutate "__p" (rawGo "__path[__depth]")
              , GoAssign "__d" (rawGo "__dirMask & 1")
              , rawGo "__dirMask >>= 1"
              , GoIfElse (rawGo "__d == 0")
                  (emitSideAscent bal.leftIdx leftRots)
                  (emitSideAscent bal.rightIdx rightRots)
              ]
          , GoMutate "__tree" (GoVar "__r")
          , rawGo (doneLabel <> ":")
          ]
    emitStep labelIdx = case _ of
      BuildLeaf (BuildDone true) -> Tuple labelIdx [ GoReturn (GoVar "__tree") ]
      BuildLeaf (BuildDone false) -> Tuple labelIdx [ GoReturn (rawGo "nil") ]
      BuildLeaf (BuildRecurse nextScalars insScalars) ->
        let
          doneLabel = "__done_ins_" <> show labelIdx
          skipNormLabel = "__skip_norm_" <> show labelIdx
          assignInsArgs = Array.zipWith (\idx e -> GoMutate ("__ins_arg_" <> show idx) e) ins.scalarIndices insScalars
          nextBindings = Array.mapWithIndex (\i e -> Tuple ("__next_" <> show i) e) nextScalars
          assignNexts = map (\(Tuple n e) -> GoAssign n e) nextBindings
          updateScalars = Array.zipWith (\idx (Tuple n _) -> GoMutate ("__arg" <> show idx) (GoVar n)) b.scalarIndices nextBindings
        in
          Tuple (labelIdx + 1)
            ( assignInsArgs
                <> emitInlineInsert doneLabel skipNormLabel
                <> [ GoIfElse (GoBinOp "!=" (GoVar "__tree") (rawGo "nil"))
                       [ GoMutate ("__tree.V" <> show bal.tagIdx) bal.guardTag ] []
                   , rawGo (skipNormLabel <> ":")
                   ]
                <> assignNexts
                <> updateScalars
            )
      BuildBranch cond yes no ->
        let
          Tuple idx1 yesStmts = emitStep labelIdx yes
          Tuple idx2 noStmts = emitStep idx1 no
        in Tuple idx2 [ GoIfElse cond yesStmts noStmts ]
    Tuple _ stepStmts = emitStep 0 b.body
  in
    GoFunctionDecl
      { name: generated owner buildFastName
      , params
      , result: spec.goType
      , body: GoBlock
          ([ GoIfElse (GoBinOp "!=" (GoVar "__tree") (rawGo "nil"))
               [ GoReturn (GoCall (GoRef (SymbolReference (generated owner b.fn.native))) fallbackArgs) ] []
           , rawGo ("var __slab []" <> structTypeStr)
           , rawGo "var __pos int"
           , rawGo ("var __path [64]" <> ptrTypeStr)
           , rawGo "var __dirMask uint64"
           , rawGo "var __depth int"
           , rawGo ("var __cur " <> ptrTypeStr)
           , rawGo ("var __r " <> ptrTypeStr)
           , rawGo ("var __p " <> ptrTypeStr)
           ] <> insArgDecls <> [ GoForRange "" stepStmts ])
      }
