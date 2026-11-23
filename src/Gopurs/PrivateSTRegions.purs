module Gopurs.PrivateSTRegions (admit, emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, any, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), isJust)
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.ExprAnalysis (executeIfOpaque, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, LocalBinding, StmtTree(..), TranslateExpr, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident, ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendEffect(..), BackendSyntax(..), Level, Pair(..))

data Handle = RefHandle | ArrayHandle
derive instance Eq Handle

type Scope = Map.Map Level Handle

-- A plan certifies every use of every private handle before code generation.
-- Ordinary values retain their original TAST and use the usual translator.
data Plan
  = Pure TcoExpr
  | Opaque TcoExpr
  | Bind (Maybe Ident) Level Plan Plan
  | LetValue (Maybe Ident) Level TcoExpr Plan
  | NewRef Level TcoExpr Plan
  | NewArray Level Plan
  | ReadRef Level
  | WriteRef Level TcoExpr
  | Push Level TcoExpr
  | Freeze Level
  | While Plan Plan
  | Branches (Array (Tuple TcoExpr Plan)) Plan

admit :: TcoExpr -> Boolean
admit expression = isJust do
  args <- arguments "Control.Monad.ST.Internal" "run" expression
  case args of
    [ action ] -> checked action
    _ -> Nothing

emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
emit translate context nextId fn args = do
  guard (named "Control.Monad.ST.Internal" "run" fn)
  action <- case args of
    [ action ] -> Just action
    _ -> Nothing
  plan <- checked action
  pure (generate translate (childContext context Nothing) Map.empty nextId plan)

checked :: TcoExpr -> Maybe Plan
checked action = do
  plan <- inspect Map.empty action
  guard (allocates plan)
  pure plan

-- These are synchronous ST operations. Unknown calls may run only when none
-- of the private handles occurs anywhere in their arguments or closures.
-- Local usage annotations are deliberately not treated as ownership evidence.
inspect :: Scope -> TcoExpr -> Maybe Plan
inspect scope expression = case unwrapTcoExpr expression of
  EffectDefer inner -> do
    -- A defer returns the result of its body, without an extra opaque-action
    -- call. Only unfold bodies which the ordinary emitter executes directly.
    guard (directEffectBody inner)
    inspect scope inner
  EffectPure value -> Pure value <$ guard (independent scope value)
  EffectBind ident level value body -> do
    guard (not (Map.member level scope))
    case allocation value of
      Just (Just initial) -> do
        guard (independent scope initial)
        NewRef level initial <$> inspect (Map.insert level RefHandle scope) body
      Just Nothing -> NewArray level <$> inspect (Map.insert level ArrayHandle scope) body
      Nothing -> Bind ident level <$> inspect scope value <*> inspect scope body
  Let ident level value body -> do
    guard (not (Map.member level scope) && independent scope value)
    LetValue ident level value <$> inspect scope body
  Branch branches def -> Branches
    <$> traverse (\(Pair cond body) -> do
      guard (independent scope cond)
      Tuple cond <$> inspect scope body) (NEA.toArray branches)
    <*> inspect scope def
  Fail _ -> Pure expression <$ guard (independent scope expression)
  PrimEffect (EffectRefRead ref) -> ReadRef <$> handle scope RefHandle ref
  PrimEffect (EffectRefWrite ref value) -> do
    level <- handle scope RefHandle ref
    guard (independent scope value)
    pure (WriteRef level value)
  UncurriedEffectApp fn args | named "Data.Array.ST" "pushImpl" fn -> case args of
    [ value, array ] -> do
      level <- handle scope ArrayHandle array
      guard (independent scope value)
      pure (Push level value)
    _ -> Nothing
  UncurriedEffectApp fn args | named "Data.Array.ST" "unsafeFreezeImpl" fn -> case args of
    [ array ] -> Freeze <$> handle scope ArrayHandle array
    _ -> Nothing
  _ | Just args <- arguments "Control.Monad.ST.Internal" "while" expression -> case args of
    [ condition, body ] -> do
      -- Preparing an action outside its suspension is a separate evaluation
      -- step. Do not move such a Let into the repeated loop computation.
      guard (repeatable condition && repeatable body)
      While <$> inspect scope condition <*> inspect scope body
    _ -> Nothing
  _ | Just [ ref ] <- arguments "Control.Monad.ST.Internal" "read" expression ->
    ReadRef <$> handle scope RefHandle ref
  _ | Just [ value, ref ] <- arguments "Control.Monad.ST.Internal" "write" expression -> do
    level <- handle scope RefHandle ref
    guard (independent scope value)
    pure (WriteRef level value)
  _ -> Opaque expression <$ guard (independent scope expression)

repeatable :: TcoExpr -> Boolean
repeatable expression = case unwrapTcoExpr expression of
  EffectDefer _ -> true
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  PrimEffect _ -> true
  UncurriedEffectApp _ _ -> true
  Local _ _ -> true
  Var _ -> true
  -- In particular, constructing an opaque action with an App can allocate
  -- state captured by that action. It must not be reconstructed on each turn.
  _ -> false

directEffectBody :: TcoExpr -> Boolean
directEffectBody expression = case unwrapTcoExpr expression of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  PrimEffect _ -> true
  UncurriedEffectApp _ _ -> true
  Fail _ -> true
  Let _ _ _ body -> directEffectBody body
  _ -> false

allocation :: TcoExpr -> Maybe (Maybe TcoExpr)
allocation expression = case unwrapTcoExpr expression of
  PrimEffect (EffectRefNew initial) -> Just (Just initial)
  _ | Just [ initial ] <- arguments "Control.Monad.ST.Internal" "newImpl" expression -> Just (Just initial)
  _ | named "Data.Array.ST" "new" expression || named "Data.Array.ST" "newImpl" expression -> Just Nothing
  _ -> Nothing

handle :: Scope -> Handle -> TcoExpr -> Maybe Level
handle scope kind expression = case unwrapTcoExpr expression of
  Local _ level | Map.lookup level scope == Just kind -> Just level
  _ -> Nothing

independent :: Scope -> TcoExpr -> Boolean
independent scope (TcoExpr _ syntax) = case syntax of
  Local _ level -> not (Map.member level scope)
  _ -> all (independent scope) syntax

named :: String -> String -> TcoExpr -> Boolean
named owner name expression = case qualifiedTarget expression of
  Just target -> target.mbMod == Just (ModuleName owner) && target.name == name
  _ -> false

arguments :: String -> String -> TcoExpr -> Maybe (Array TcoExpr)
arguments owner name expression = do
  let Tuple fn spine = collectGoSpine expression
  guard (named owner name fn)
  pure (getGoSpineArgs spine)

allocates :: Plan -> Boolean
allocates = case _ of
  NewRef _ _ _ -> true
  NewArray _ _ -> true
  Bind _ _ a b -> allocates a || allocates b
  LetValue _ _ _ b -> allocates b
  While a b -> allocates a || allocates b
  Branches bs def -> any (\(Tuple _ body) -> allocates body) bs || allocates def
  _ -> false

type Locals = Map.Map Level LocalBinding

generate :: TranslateExpr -> ExprContext -> Locals -> Int -> Plan -> ExprResult
generate translate context@{ codegenStateRef, modNameStr } locals nextId plan =
  let
    pureValue = translate (childContext context Nothing)
    boxed result = boxGoExpr codegenStateRef modNameStr result.expr result.exprType
    cell level = case Map.lookup level locals of
      Just value -> value
      Nothing -> { name: "__invalid_private_st_cell", goType: TypeValue }
    unit = { stmts: StmtEmpty, expr: rawGo "gopurs_runtime.Value{}", exprType: TypeValue, nextId }
    localResult level = let value = cell level
      in { stmts: StmtEmpty, expr: GoVar value.name, exprType: value.goType, nextId }
    continue ident level result body =
      let
        name = "__st_value" <> show result.nextId
        bound = Map.insert (localId ident level) { name, goType: result.exprType } context.bound
        rest = generate translate (context { bound = bound }) locals (result.nextId + 1) body
      in rest { stmts = result.stmts <> StmtLeaf (GoAssign name (scalarBinding result.exprType result.expr)) <> rest.stmts }
  in case plan of
    Pure expression -> pureValue nextId expression
    Opaque expression ->
      let result = translate (context { options = { isTail: false, inEffectBlock: true } }) nextId expression
      in result { expr = executeIfOpaque expression (boxed result), exprType = TypeValue }
    Bind ident level value body -> continue ident level (generate translate context locals nextId value) body
    LetValue ident level value body -> continue ident level (pureValue nextId value) body
    NewRef level initial body ->
      let
        result = pureValue nextId initial
        ty = case result.exprType of
          TypeString -> TypeString
          TypeBool -> TypeBool
          TypeInt64 -> TypeInt64
          TypeFloat64 -> TypeFloat64
          _ -> TypeValue
        name = "__st_ref" <> show result.nextId
        value = coerceGoExpr codegenStateRef modNameStr result.expr result.exprType ty
        -- Native scalar literals need an explicit Go type at their binding.
        initialValue = scalarBinding ty value
        rest = generate translate context (Map.insert level { name, goType: ty } locals) (result.nextId + 1) body
      in rest { stmts = result.stmts <> StmtLeaf (GoAssign name initialValue) <> rest.stmts }
    NewArray level body ->
      let
        name = "__st_array" <> show nextId
        rest = generate translate context (Map.insert level { name, goType: TypeNativeArray TypeValue } locals) (nextId + 1) body
      in rest { stmts = StmtLeaf (GoAssign name (rawGo "[]gopurs_runtime.Value{}")) <> rest.stmts }
    ReadRef level -> localResult level
    WriteRef level expression ->
      let target = cell level
          result = pureValue nextId expression
      in { stmts: result.stmts <> StmtLeaf (GoMutate target.name
            (coerceGoExpr codegenStateRef modNameStr result.expr result.exprType target.goType))
         , expr: GoVar target.name, exprType: target.goType, nextId: result.nextId }
    Push level expression ->
      let target = cell level
          result = pureValue nextId expression
      in { stmts: result.stmts <> StmtLeaf (GoMutate target.name
            (GoCall (GoVar "append") [ GoVar target.name, boxed result ]))
         , expr: GoCall (GoVar "int64") [ GoCall (GoVar "len") [ GoVar target.name ] ]
         , exprType: TypeInt64, nextId: result.nextId }
    Freeze level ->
      let target = cell level
          name = "__st_frozen" <> show nextId
      -- The ordinary bridge returns a snapshot. Keep that guarantee even when
      -- the admitted ST program continues pushing after this publication.
      in { stmts: StmtLeaf (GoAssign name (GoCall (GoVar "make")
              [ rawGo "[]gopurs_runtime.Value", GoCall (GoVar "len") [ GoVar target.name ] ]))
              <> StmtLeaf (GoCall (GoVar "copy") [ GoVar name, GoVar target.name ])
         , expr: GoVar name, exprType: TypeNativeArray TypeValue, nextId: nextId + 1 }
    While condition body ->
      let
        -- Atomic opaque actions are values prepared before while is called.
        -- Resolve even a global getter once, including when the first test is
        -- false, and repeatedly invoke that same action value.
        prepare start (Opaque expression) =
          let result = pureValue start expression
              name = "__st_action" <> show result.nextId
          in { before: result.stmts <> StmtLeaf (GoAssign name (boxed result))
             , result: { stmts: StmtEmpty
               , expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Apply")
                   [ GoVar name, rawGo "gopurs_runtime.Value{}" ]
               , exprType: TypeValue, nextId: result.nextId + 1 } }
        prepare start action = { before: StmtEmpty, result: generate translate context locals start action }
        testPlan = prepare nextId condition
        bodyPlan = prepare testPlan.result.nextId body
        test = testPlan.result
        action = bodyPlan.result
        conditionValue = coerceGoExpr codegenStateRef modNameStr test.expr test.exprType TypeBool
        statements = flattenStmts test.stmts
          <> [ GoIfElse (GoPrefixOp "!" conditionValue) [ rawGo "break" ] [] ]
          <> flattenStmts action.stmts <> [ GoMutate "_" action.expr ]
      in unit { stmts = testPlan.before <> bodyPlan.before <> StmtLeaf (GoForRange "" statements), nextId = action.nextId }
    Branches branches def ->
      let
        resultName = "__st_branch" <> show nextId
        fallback = generate translate context locals (nextId + 1) def
        addBranch acc (Tuple condition body) =
          let test = pureValue acc.nextId condition
              action = generate translate context locals test.nextId body
              statements = flattenStmts test.stmts <> [ GoIfElse
                (coerceGoExpr codegenStateRef modNameStr test.expr test.exprType TypeBool)
                (flattenStmts action.stmts <> [ GoMutate resultName (boxed action) ]) acc.statements ]
          in { statements, nextId: action.nextId }
        result = foldl addBranch
          { statements: flattenStmts fallback.stmts <> [ GoMutate resultName (boxed fallback) ], nextId: fallback.nextId }
          (Array.reverse branches)
      in { stmts: StmtLeaf (rawGo ("var " <> resultName <> " gopurs_runtime.Value")) <> StmtLeaf (GoBlock result.statements)
         , expr: GoVar resultName, exprType: TypeValue, nextId: result.nextId }

scalarBinding :: GoType -> GoExpr -> GoExpr
scalarBinding TypeInt64 value = GoCall (GoVar "int64") [ value ]
scalarBinding TypeFloat64 value = GoCall (GoVar "float64") [ value ]
scalarBinding TypeUint32 value = GoCall (GoVar "uint32") [ value ]
scalarBinding _ value = value
