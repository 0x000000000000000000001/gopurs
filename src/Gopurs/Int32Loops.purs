module Gopurs.Int32Loops (loop) where

import Prelude
import Data.Array as Array
import Data.Enum (fromEnum)
import Data.Foldable (all, any, foldl)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.String (Pattern(..), contains)
import Data.String.CodePoints as CodePoints
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoFunctions (Parameters, iterationBindings)

-- Int's public representation remains int64: zshr and bottom / -1 can produce
-- values outside the signed 32-bit range. Version only scalar tail loops whose
-- selected slots are written by proven narrowing operations (including through
-- proven-normalized local branch temporaries). Check their initial values once,
-- then keep both slots and iteration aliases native inside the loop.
-- Also hoist loop-invariant Value array representation dispatches out of pure
-- tail loops so ArrayAccess/ArrayLength execute directly against typed slices.
loop :: String -> Parameters -> GoType -> Array GoExpr -> GoExpr
loop label params result statements =
  let
    paramNames = Set.fromFoldable (map fst params)
    originalLoop = hoistInvariantArrays label params [] result (iterationBindings params <> statements)
    { promoted, normTemps } = analyzeInt32Slots label params statements
  in
    case Array.uncons promoted of
      Just { head, tail } | result == TypeInt64 ->
        let
          fits name = GoBinOp "==" (GoVar (slot name)) (cast "int64" (cast "int32" (GoVar (slot name))))
          guard = foldl (GoBinOp "&&") (fits head) (map fits tail)
          initial = map (\name -> GoAssign (slot name) (cast "int32" (GoVar (slot name)))) promoted
          aliases = Array.concatMap (\param@(Tuple name _) ->
            if Array.elem name promoted then [ GoAssign name (GoVar (slot name)) ]
            else iterationBindings [ param ]) params
          rewritten = aliases <> map (rewrite promoted normTemps) statements
          fastBranch = hoistInvariantArraysReturn label params initial result rewritten
        in
          -- Go labels have function scope. An IIFE gives the copied branch/loop
          -- labels a separate scope without renaming opaque goto syntax.
          GoBlock [ GoIfElse guard fastBranch [], originalLoop ]
      _ ->
        let _ = paramNames
        in originalLoop

slot :: String -> String
slot name = name <> "_loop"

cast :: String -> GoExpr -> GoExpr
cast name value = GoCall (GoVar name) [ value ]

arithmetic :: GoExpr -> Maybe { operator :: String, left :: GoExpr, right :: GoExpr }
arithmetic = case _ of
  GoCall (GoSelector (GoVar "gopurs_runtime") name) [ left, right ]
    | name == "IntAdd" -> Just { operator: "+", left, right }
    | name == "IntSub" -> Just { operator: "-", left, right }
  _ -> Nothing

constPosMod :: GoExpr -> Maybe { left :: GoExpr, divisor :: Int }
constPosMod = case _ of
  GoCall (GoSelector (GoVar "gopurs_runtime") "IntModPos64") [ left, GoCall (GoVar "int64") [ GoInt d ] ]
    | d > 0 && signedLiteral d -> Just { left, divisor: d }
  GoCall (GoSelector (GoVar "gopurs_runtime") "IntModPos64") [ left, GoInt d ]
    | d > 0 && signedLiteral d -> Just { left, divisor: d }
  _ -> Nothing

constPosDiv :: GoExpr -> Maybe { left :: GoExpr, divisor :: Int }
constPosDiv = case _ of
  GoCall (GoSelector (GoVar "gopurs_runtime") "IntDivPos64") [ left, GoCall (GoVar "int64") [ GoInt d ] ]
    | d > 0 && signedLiteral d -> Just { left, divisor: d }
  GoCall (GoSelector (GoVar "gopurs_runtime") "IntDivPos64") [ left, GoInt d ]
    | d > 0 && signedLiteral d -> Just { left, divisor: d }
  _ -> Nothing

powerOfTwoMask :: GoExpr -> Maybe { left :: GoExpr, mask :: Int }
powerOfTwoMask = case _ of
  GoBinOp "&" left (GoInt mask) | mask >= 0 && signedLiteral mask -> Just { left, mask }
  _ -> Nothing

-- Deliberately exclude general division, unsigned shifts and arbitrary
-- Int-returning calls. Types alone do not prove a signed 32-bit result.
normalized :: GoExpr -> Boolean
normalized expr = case expr of
  GoInt value -> signedLiteral value
  GoCall (GoVar "int64") [ GoInt value ] -> signedLiteral value
  _ -> case arithmetic expr of
    Just _ -> true
    Nothing -> case constPosMod expr of
      Just _ -> true
      Nothing -> case powerOfTwoMask expr of
        Just _ -> true
        Nothing -> false

-- PBO can fold zshr or bottom / -1 to an Int literal outside this range.
signedLiteral :: Int -> Boolean
signedLiteral value = value >= bottom && value <= top

isSimpleIdent :: String -> Boolean
isSimpleIdent text =
  let codes = map fromEnum (CodePoints.toCodePointArray text)
  in case Array.uncons codes of
    Just { head, tail } -> isStart head && all isRest tail
    Nothing -> false
  where
  isStart c = c == 95 || (c >= 65 && c <= 90) || (c >= 97 && c <= 122)
  isRest c = isStart c || (c >= 48 && c <= 57)

parseVarInt64Decl :: GoExpr -> Maybe String
parseVarInt64Decl = case _ of
  GoRaw { text, imports } | Array.null imports -> do
    rest <- String.stripPrefix (Pattern "var ") text
    name <- String.stripSuffix (Pattern " int64") rest
    if isSimpleIdent name then Just name else Nothing
  _ -> Nothing

parseDiscardRaw :: GoExpr -> Maybe String
parseDiscardRaw = case _ of
  GoRaw { text, imports } | Array.null imports -> do
    name <- String.stripPrefix (Pattern "_ = ") text
    if isSimpleIdent name then Just name else Nothing
  _ -> Nothing

-- Collect locally declared int64 temporaries (e.g. ControlExprs branch temps
-- `var __tN int64`) with strict scope dominance: a variable is only admitted
-- when declared once inside the loop body, never referenced outside its active
-- declaration scope, and never mentioned in arbitrary GoRaw code.
type TempScan =
  { scopes :: Array (Set String)
  , declared :: Set String
  , tainted :: Set String
  , writes :: Map String (Array GoExpr)
  , raws :: Array String
  }

collectLocalTemps :: Set String -> Array GoExpr -> Map String (Array GoExpr)
collectLocalTemps reserved stmts =
  let
    emptyState :: TempScan
    emptyState =
      { scopes: [ Set.empty ]
      , declared: Set.empty
      , tainted: reserved
      , writes: Map.empty
      , raws: []
      }
    inScope name state = any (Set.member name) state.scopes
    pushScope state = state { scopes = Array.snoc state.scopes Set.empty }
    popScope state = case Array.unsnoc state.scopes of
      Just { init } | not (Array.null init) -> state { scopes = init }
      _ -> state
    declare name state =
      if Set.member name state.declared || Set.member name state.tainted then
        state { tainted = Set.insert name state.tainted }
      else
        let
          scopes' = case Array.unsnoc state.scopes of
            Just { init, last } -> Array.snoc init (Set.insert name last)
            Nothing -> [ Set.singleton name ]
        in
          state
            { scopes = scopes'
            , declared = Set.insert name state.declared
            , writes = Map.insertWith (<>) name [] state.writes
            }
    recordWrite name rhs state =
      if not (inScope name state) then state { tainted = Set.insert name state.tainted }
      else state { writes = Map.insertWith (<>) name [ rhs ] state.writes }
    useVar name state =
      if Set.member name state.declared && not (inScope name state) then
        state { tainted = Set.insert name state.tainted }
      else state
    walkExpr expr state = case expr of
      GoVar name -> useVar name state
      GoString _ -> state
      GoInt _ -> state
      GoCall fn args -> foldl (flip walkExpr) (walkExpr fn state) args
      GoSelector obj _ -> walkExpr obj state
      GoBinOp _ l r -> walkExpr r (walkExpr l state)
      GoPrefixOp _ v -> walkExpr v state
      GoTypeAssertion v _ -> walkExpr v state
      GoIndex obj idx -> walkExpr idx (walkExpr obj state)
      GoRaw { text } -> state { raws = Array.snoc state.raws text }
      _ -> state
    walkStmt stmt state = case stmt of
      GoRaw { text, imports } | Array.null imports ->
        case parseVarInt64Decl stmt of
          Just name -> declare name state
          Nothing -> case parseDiscardRaw stmt of
            Just _ -> state
            Nothing ->
              if text == "{" then pushScope state
              else if text == "}" then popScope state
              else state { raws = Array.snoc state.raws text }
      GoRaw { text } -> state { raws = Array.snoc state.raws text }
      GoAssign name rhs ->
        let s1 = walkExpr rhs state
            s2 = declare name s1
        in recordWrite name rhs s2
      GoMutate name rhs ->
        let s1 = walkExpr rhs state
        in recordWrite name rhs s1
      GoReturn value -> walkExpr value state
      GoContinue _ -> state
      GoBlock inner -> popScope (foldl (flip walkStmt) (pushScope state) inner)
      GoIfElse cond yes no ->
        let sCond = walkExpr cond state
            sYes = popScope (foldl (flip walkStmt) (pushScope sCond) yes)
        in popScope (foldl (flip walkStmt) (pushScope sYes) no)
      other -> walkExpr other state
    finalState = foldl (flip walkStmt) emptyState stmts
    mentionedInRaw name = any (\text -> contains (Pattern name) text) finalState.raws
  in
    Map.filterWithKey (\name _ -> not (Set.member name finalState.tainted) && not (mentionedInRaw name)) finalState.writes

-- Fail closed on opaque references, shadowing, closures, nested loops and
-- unrecognized AST forms. Collects all RHS expressions assigned to `slot name`.
inspectParam :: String -> String -> Array GoExpr -> Maybe (Array GoExpr)
inspectParam label name statements = map Array.concat (traverse go statements)
  where
  reserved value = value == name || value == slot name
  many values = map Array.concat (traverse go values)
  go = case _ of
    GoVar value | value == slot name -> Nothing
    GoVar _ -> Just []
    GoString _ -> Just []
    GoInt _ -> Just []
    GoRaw code | contains (Pattern name) code.text -> Nothing
    GoRaw _ -> Just []
    GoAssign target _ | reserved target -> Nothing
    GoAssign _ value -> go value
    GoMutate target value
      | target == name -> Nothing
      | target == slot name -> map (Array.cons value) (go value)
      | otherwise -> go value
    GoContinue target -> if target == label then Just [] else Nothing
    GoCall fn args -> many (Array.cons fn args)
    GoSelector obj _ -> go obj
    GoBinOp _ left right -> many [ left, right ]
    GoPrefixOp _ value -> go value
    GoTypeAssertion value _ -> go value
    GoIndex obj index -> many [ obj, index ]
    GoReturn value -> go value
    GoBlock values -> many values
    GoIfElse cond yes no -> many (Array.cons cond (yes <> no))
    _ -> Nothing

analyzeInt32Slots :: String -> Parameters -> Array GoExpr -> { promoted :: Array String, normTemps :: Set String }
analyzeInt32Slots label params statements =
  let
    reserved = Set.fromFoldable (Array.concatMap (\(Tuple n _) -> [ n, slot n ]) params)
    tempWrites = collectLocalTemps reserved statements
    paramWrites :: Map String (Array GoExpr)
    paramWrites = Map.fromFoldable (Array.mapMaybe (\(Tuple name ty) ->
      if ty /= TypeInt64 then Nothing
      else case inspectParam label name statements of
        Just writes | not (Array.null writes) -> Just (Tuple name writes)
        _ -> Nothing) params)
    initialParams = Set.fromFoldable (Map.keys paramWrites)
    initialTemps = Set.fromFoldable (Map.keys tempWrites)
    exprNorm pSet tSet expr = case expr of
      GoVar v -> Set.member v pSet || Set.member v tSet
      GoCall (GoVar "int64") [ inner ] -> normalized expr || exprNorm pSet tSet inner
      _ -> normalized expr
    fixpoint pSet tSet =
      let
        tSet' = Set.filter (\t -> all (exprNorm pSet tSet) (fromMaybe [] (Map.lookup t tempWrites))) tSet
        pSet' = Set.filter (\p -> all (exprNorm pSet tSet') (fromMaybe [] (Map.lookup p paramWrites))) pSet
      in
        if pSet' == pSet && tSet' == tSet then { pSet, tSet }
        else fixpoint pSet' tSet'
    { pSet, tSet } = fixpoint initialParams initialTemps
    promoted = Array.mapMaybe (\(Tuple name _) -> if Set.member name pSet then Just name else Nothing) params
  in
    { promoted, normTemps: tSet }

children :: GoExpr -> Array GoExpr
children = case _ of
  GoCall fn args -> Array.cons fn args
  GoSelector obj _ -> [ obj ]
  GoBinOp _ left right -> [ left, right ]
  GoPrefixOp _ value -> [ value ]
  GoTypeAssertion value _ -> [ value ]
  GoIndex obj index -> [ obj, index ]
  _ -> []

readsNative :: Array String -> GoExpr -> Boolean
readsNative names = go
  where
  go (GoVar name) = Array.elem name names
  go expr = Array.any go (children expr)

known32 :: Array String -> Set String -> GoExpr -> Boolean
known32 names normTemps (GoVar name) = Array.elem name names || Set.member name normTemps
known32 _ _ expr = normalized expr

isRuntimeExpr :: GoExpr -> Boolean
isRuntimeExpr = case _ of
  GoVar _ -> true
  GoIndex _ _ -> true
  GoSelector _ _ -> true
  GoCall (GoVar name) [ arg ]
    | name == "int64" || name == "int32" || name == "int" || name == "uint64" || name == "uint32" -> isRuntimeExpr arg
  GoCall _ _ -> true
  GoBinOp _ l r -> isRuntimeExpr l || isRuntimeExpr r
  GoPrefixOp _ v -> isRuntimeExpr v
  _ -> false

comparison :: String -> Boolean
comparison op = Array.elem op [ "==", "!=", "<", "<=", ">", ">=" ]

rewrite :: Array String -> Set String -> GoExpr -> GoExpr
rewrite names normTemps = wide
  where
  narrow expr = case expr of
    GoVar name
      | Array.elem name names -> expr
      | Set.member name normTemps -> cast "int32" expr
    GoInt value | signedLiteral value -> cast "int32" expr
    GoCall (GoVar "int64") [ GoInt value ] -> narrow (GoInt value)
    _ -> case arithmetic expr of
      -- A native operand also prevents Go constant expressions overflowing at
      -- type checking time. Constant-only additions retain the runtime helper.
      Just { operator, left, right } | readsNative names expr ->
        GoBinOp operator (narrow left) (narrow right)
      _ -> case constPosMod expr of
        Just { left, divisor } | known32 names normTemps left ->
          GoCall (GoSelector (GoVar "gopurs_runtime") "IntModPos32") [ narrow left, cast "int32" (GoInt divisor) ]
        _ -> case constPosDiv expr of
          Just { left, divisor } | known32 names normTemps left ->
            GoCall (GoSelector (GoVar "gopurs_runtime") "IntDivPos32") [ narrow left, cast "int32" (GoInt divisor) ]
          _ -> case powerOfTwoMask expr of
            Just { left, mask } | known32 names normTemps left ->
              GoBinOp "&" (narrow left) (cast "int32" (GoInt mask))
            _ ->
              let widened = wide expr
              in
                if isRuntimeExpr widened then cast "int32" widened
                else
                  -- Constant unsigned expressions cannot convert directly to int32 in Go.
                  cast "int32" (GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd")
                    [ widened, cast "int64" (GoInt 0) ])

  wide expr = case expr of
    GoVar name | Array.elem name names -> cast "int64" expr
    GoCall fn args -> case arithmetic expr of
      Just _ | readsNative names expr -> cast "int64" (narrow expr)
      _ -> case constPosMod expr of
        Just { left } | known32 names normTemps left -> cast "int64" (narrow expr)
        _ -> case constPosDiv expr of
          Just { left } | known32 names normTemps left -> cast "int64" (narrow expr)
          _ -> GoCall (wide fn) (map wide args)
    GoBinOp op left right
      | comparison op && known32 names normTemps left && known32 names normTemps right -> GoBinOp op (narrow left) (narrow right)
      | otherwise -> case powerOfTwoMask expr of
          Just { left: inner } | known32 names normTemps inner -> cast "int64" (narrow expr)
          _ -> GoBinOp op (wide left) (wide right)
    GoSelector obj field -> GoSelector (wide obj) field
    GoPrefixOp op value -> GoPrefixOp op (wide value)
    GoTypeAssertion value ty -> GoTypeAssertion (wide value) ty
    GoIndex obj index -> GoIndex (wide obj) (wide index)
    GoAssign target value -> GoAssign target (wide value)
    GoMutate target value
      | Array.any (\name -> target == slot name) names -> GoMutate target (narrow value)
      | otherwise -> GoMutate target (wide value)
    GoReturn value -> GoReturn (wide value)
    GoBlock values -> GoBlock (map wide values)
    GoIfElse cond yes no -> GoIfElse (wide cond) (map wide yes) (map wide no)
    _ -> expr

-- ---------------------------------------------------------------------------
-- Loop-Invariant Value Array Representation Hoisting
-- ---------------------------------------------------------------------------
-- When a pure tail loop indexes an invariant gopurs_runtime.Value via
-- ArrayAccess / ArrayLength (directly or through a local alias), hoist the
-- runtime tag switch once before the loop so the inner loop indexes a Go slice
-- held directly in registers with zero per-iteration tag checks or slice-header
-- reloads. Unrecognized tags fall back to the unmodified loop without ever
-- dereferencing UnsafePtr.

data HoistKind
  = HoistInt32
  | HoistInt64
  | HoistArray
  | HoistString

hoistInvariantArrays :: String -> Parameters -> Array GoExpr -> GoType -> Array GoExpr -> GoExpr
hoistInvariantArrays label params setup result stmts =
  case findHoistCandidate params stmts of
    Nothing -> GoFor label stmts
    Just candidate ->
      GoBlock (emitHoistedBranches label setup result stmts candidate <> [ GoFor label stmts ])

hoistInvariantArraysReturn :: String -> Parameters -> Array GoExpr -> GoType -> Array GoExpr -> Array GoExpr
hoistInvariantArraysReturn label params setup result stmts =
  let fallback = GoReturn (GoCall (GoFuncBlock [] (setup <> [ GoFor label stmts ]) result) [])
  in case findHoistCandidate params stmts of
    Nothing -> [ fallback ]
    Just candidate ->
      emitHoistedBranches label setup result stmts candidate <> [ fallback ]

type HoistCandidate =
  { root :: String
  , sourceExprName :: String
  , targets :: Set String
  , intValOnly :: Boolean
  }

emitHoistedBranches :: String -> Array GoExpr -> GoType -> Array GoExpr -> HoistCandidate -> Array GoExpr
emitHoistedBranches label setup result stmts { root, sourceExprName, targets, intValOnly } =
  let
    sliceVar = "__gopurs_hoisted_" <> root
    kinds =
      if intValOnly then [ HoistInt32, HoistInt64, HoistArray ]
      else [ HoistInt32, HoistInt64, HoistArray, HoistString ]
    branch kind =
      let
        tagName = case kind of
          HoistInt32 -> "TypeInt32Array"
          HoistInt64 -> "TypeIntArray"
          HoistArray -> "TypeArray"
          HoistString -> "TypeStringArray"
        sliceType = case kind of
          HoistInt32 -> "[]int32"
          HoistInt64 -> "[]int64"
          HoistArray -> "[]gopurs_runtime.Value"
          HoistString -> "[]string"
        cond = GoBinOp "=="
          (GoSelector (GoVar sourceExprName) "Type")
          (GoSelector (GoVar "gopurs_runtime") tagName)
        sliceDecl = GoAssign sliceVar (rawGo ("*(*" <> sliceType <> ")(" <> sourceExprName <> ".UnsafePtr)"))
        rewritten = map (rewriteArrayAccess kind targets sliceVar) stmts
        iife = GoCall (GoFuncBlock [] (setup <> [ sliceDecl, GoFor label rewritten ]) result) []
      in
        GoIfElse cond [ GoReturn iife ] []
  in
    map branch kinds

findHoistCandidate :: Parameters -> Array GoExpr -> Maybe HoistCandidate
findHoistCandidate params stmts =
  if not (all isPureLoopStmt stmts) then Nothing
  else
    let
      paramTypes = Map.fromFoldable params
      -- Track alias chains `alias := v` (single assign, never mutated, not in opaque GoRaw).
      aliasScan = scanAliases stmts
      resolve name = go 0 name
        where
        go depth cur
          | depth > 8 = cur
          | otherwise = case Map.lookup cur aliasScan.aliases of
              Just parent -> go (depth + 1) parent
              Nothing -> cur
      accessInfo = collectArrayAccesses stmts
      roots = Array.nub (map resolve accessInfo.targets)
      isInvariant root =
        not (Set.member root aliasScan.mutated)
          && not (Set.member root aliasScan.opaqueRawVars)
          && case Map.lookup root paramTypes of
            Nothing -> not (Set.member root aliasScan.assigned)
            Just TypeValue ->
              -- Loop parameter: must only be assigned by iterationBindings (`var root ... = root_loop`)
              -- and its slot must only ever be written with `root` (or never written).
              all (\rhs -> rhs == GoVar root) (fromMaybe [] (Map.lookup (slot root) aliasScan.slotWrites))
            Just _ -> false
    in
      case Array.find isInvariant roots of
        Nothing -> Nothing
        Just root ->
          let
            targets = Set.fromFoldable (Array.filter (\t -> resolve t == root) accessInfo.targets)
            intValOnly = all (\t -> not (Set.member t accessInfo.nonIntValTargets)) (Array.fromFoldable targets)
            sourceExprName = if Map.member root paramTypes then slot root else root
          in
            Just { root, sourceExprName, targets, intValOnly }

isPureLoopStmt :: GoExpr -> Boolean
isPureLoopStmt = go
  where
  allowedRuntime =
    [ "Int", "Bool", "Float", "Str"
    , "IntAdd", "IntSub", "IntMul", "IntDiv", "IntMod"
    , "IntDivPos32", "IntDivPos64", "IntModPos32", "IntModPos64", "IntModZero"
    , "IntNegate", "IntBitNot", "IntBitAnd", "IntBitOr", "IntBitXor"
    , "IntShl", "IntShr", "Zshr"
    , "ArrayAccess", "ArrayLength"
    ]
  allowedBuiltins = [ "int", "int32", "int64", "uint32", "uint64", "float64", "len" ]
  isPureRaw text =
    text == "{" || text == "}"
      || String.indexOf (Pattern "var ") text == Just 0
      || String.indexOf (Pattern "_ = ") text == Just 0
      || String.indexOf (Pattern "goto ") text == Just 0
      || String.indexOf (Pattern "end_branch_") text == Just 0
      || (String.indexOf (Pattern "func() ") text == Just 0 && contains (Pattern "panic(") text)
  go = case _ of
    GoVar _ -> true
    GoString _ -> true
    GoInt _ -> true
    GoContinue _ -> true
    GoRaw { text, imports } -> Array.null imports && isPureRaw text
    GoCall (GoVar fn) args -> Array.elem fn allowedBuiltins && all go args
    GoCall (GoSelector (GoVar "gopurs_runtime") fn) args -> Array.elem fn allowedRuntime && all go args
    GoSelector obj _ -> go obj
    GoBinOp _ l r -> go l && go r
    GoPrefixOp _ v -> go v
    GoTypeAssertion v _ -> go v
    GoIndex obj idx -> go obj && go idx
    GoAssign _ v -> go v
    GoMutate _ v -> go v
    GoReturn v -> go v
    GoBlock inner -> all go inner
    GoIfElse c y n -> go c && all go y && all go n
    _ -> false

scanAliases :: Array GoExpr -> { aliases :: Map String String, assigned :: Set String, mutated :: Set String, slotWrites :: Map String (Array GoExpr), opaqueRawVars :: Set String }
scanAliases stmts =
  let
    empty = { aliases: Map.empty, assignCount: Map.empty, assigned: Set.empty, mutated: Set.empty, slotWrites: Map.empty, raws: [] }
    walk stmt st = case stmt of
      GoAssign target rhs ->
        let
          count = 1 + fromMaybe 0 (Map.lookup target st.assignCount)
          aliases' = case rhs of
            GoVar src | count == 1 -> Map.insert target src st.aliases
            _ -> Map.delete target st.aliases
        in
          st
            { aliases = aliases'
            , assignCount = Map.insert target count st.assignCount
            , assigned = Set.insert target st.assigned
            }
      GoMutate target rhs ->
        st
          { aliases = Map.delete target st.aliases
          , mutated = Set.insert target st.mutated
          , slotWrites = Map.insertWith (<>) target [ rhs ] st.slotWrites
          }
      GoBlock inner -> foldl (flip walk) st inner
      GoIfElse _ y n -> foldl (flip walk) (foldl (flip walk) st y) n
      GoRaw { text } ->
        case parseDiscardRaw stmt of
          Just _ -> st
          Nothing ->
            if text == "{" || text == "}"
              || String.indexOf (Pattern "goto ") text == Just 0
              || String.indexOf (Pattern "end_branch_") text == Just 0
              || String.indexOf (Pattern "var ") text == Just 0
              || (String.indexOf (Pattern "func() ") text == Just 0 && contains (Pattern "panic(") text)
            then st
            else st { raws = Array.snoc st.raws text }
      _ -> st
    res = foldl (flip walk) empty stmts
    cleanAliases = Map.filterWithKey (\k _ -> not (Set.member k res.mutated) && fromMaybe 0 (Map.lookup k res.assignCount) == 1) res.aliases
  in
    { aliases: cleanAliases, assigned: res.assigned, mutated: res.mutated, slotWrites: res.slotWrites, opaqueRawVars: Set.empty }

collectArrayAccesses :: Array GoExpr -> { targets :: Array String, nonIntValTargets :: Set String }
collectArrayAccesses stmts = foldl (flip (walk false)) { targets: [], nonIntValTargets: Set.empty } stmts
  where
  walk underIntVal expr st = case expr of
    GoSelector inner "IntVal" -> walk true inner st
    GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ GoVar target, idx ] ->
      let st1 = walk false idx st
      in st1
        { targets = Array.snoc st1.targets target
        , nonIntValTargets = if underIntVal then st1.nonIntValTargets else Set.insert target st1.nonIntValTargets
        }
    GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayLength") [ GoVar target ] ->
      st { targets = Array.snoc st.targets target }
    GoCall fn args -> foldl (flip (walk false)) (walk false fn st) args
    GoSelector obj _ -> walk false obj st
    GoBinOp _ l r -> walk false r (walk false l st)
    GoPrefixOp _ v -> walk false v st
    GoTypeAssertion v _ -> walk false v st
    GoIndex obj idx -> walk false idx (walk false obj st)
    GoAssign _ v -> walk false v st
    GoMutate _ v -> walk false v st
    GoReturn v -> walk false v st
    GoBlock inner -> foldl (flip (walk false)) st inner
    GoIfElse c y n -> foldl (flip (walk false)) (foldl (flip (walk false)) (walk false c st) y) n
    _ -> st

rewriteArrayAccess :: HoistKind -> Set String -> String -> GoExpr -> GoExpr
rewriteArrayAccess kind targets sliceVar = go
  where
  isTarget (GoVar name) = Set.member name targets
  isTarget _ = false
  go = case _ of
    GoCall (GoVar "int32") [ GoSelector (GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ target, idx ]) "IntVal" ]
      | isTarget target ->
          let idx' = go idx
          in case kind of
            HoistInt32 -> GoIndex (GoVar sliceVar) idx'
            HoistInt64 -> cast "int32" (GoIndex (GoVar sliceVar) idx')
            HoistArray -> cast "int32" (GoSelector (GoIndex (GoVar sliceVar) idx') "IntVal")
            HoistString -> cast "int32" (GoSelector (GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoIndex (GoVar sliceVar) idx' ]) "IntVal")
    GoSelector (GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ target, idx ]) "IntVal"
      | isTarget target ->
          let idx' = go idx
          in case kind of
            HoistInt32 -> cast "int64" (GoIndex (GoVar sliceVar) idx')
            HoistInt64 -> GoIndex (GoVar sliceVar) idx'
            HoistArray -> GoSelector (GoIndex (GoVar sliceVar) idx') "IntVal"
            HoistString -> GoSelector (GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoIndex (GoVar sliceVar) idx' ]) "IntVal"
    GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ target, idx ]
      | isTarget target ->
          let idx' = go idx
          in case kind of
            HoistInt32 -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ cast "int64" (GoIndex (GoVar sliceVar) idx') ]
            HoistInt64 -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoIndex (GoVar sliceVar) idx' ]
            HoistArray -> GoIndex (GoVar sliceVar) idx'
            HoistString -> GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoIndex (GoVar sliceVar) idx' ]
    GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayLength") [ target ]
      | isTarget target -> GoCall (GoVar "len") [ GoVar sliceVar ]
    GoCall fn args -> GoCall (go fn) (map go args)
    GoSelector obj field -> GoSelector (go obj) field
    GoBinOp op l r -> GoBinOp op (go l) (go r)
    GoPrefixOp op v -> GoPrefixOp op (go v)
    GoTypeAssertion v ty -> GoTypeAssertion (go v) ty
    GoIndex obj idx -> GoIndex (go obj) (go idx)
    GoAssign target v -> GoAssign target (go v)
    GoMutate target v -> GoMutate target (go v)
    GoReturn v -> GoReturn (go v)
    GoBlock inner -> GoBlock (map go inner)
    GoIfElse c y n -> GoIfElse (go c) (map go y) (map go n)
    other -> other
