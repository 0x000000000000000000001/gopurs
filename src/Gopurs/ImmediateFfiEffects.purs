module Gopurs.ImmediateFfiEffects
  ( translate
  , direct
  , bindThunk
  , invokeBoundThunk
  , invokeBoundThunkExpr
  ) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (any, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as Arguments
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.ExprAnalysis (executeIfOpaque, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext)
import Gopurs.FfiBridge.ImmediateEffects as Workers
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoSymbol (GoReference(..), generated)
import Gopurs.ImmediateEffectBranches as ImmediateEffectBranches
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident, Literal(..), ModuleName)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Abs, App, EffectBind, Let, LetRec, Lit, Local, UncurriedAbs, UncurriedEffectAbs, Var), Level)

-- A directly adjacent bind is the only permitted consumer. In particular, a
-- captured, replayed or saved action keeps the existing boxed representation
-- unless bindThunk proves every use is an EffectBind action.
translate :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> Maybe ExprResult
translate translateExpr context nextId actionIdent actionLevel factory body = case unwrapTcoExpr body of
  EffectBind resultIdent resultLevel action continuation -> do
    guard context.options.inEffectBlock
    case unwrapTcoExpr action of
      Local ident level -> guard (ident == actionIdent && level == actionLevel)
      _ -> Nothing
    guard (not (usesLocal actionIdent actionLevel continuation))
    direct translateExpr context nextId resultIdent resultLevel factory continuation
  _ -> Nothing

-- A factory directly in EffectBind is also immediately executed. Unlike
-- the Let form it creates no reusable action binding; all of its arguments,
-- the native factory and then its thunk retain their original order. Direct
-- factories inline the call and skip Box when the continuation discards the
-- result; complex adapters retain their certified _runEffectWorker.
direct :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> Maybe ExprResult
direct translateExpr context nextId resultIdent resultLevel factory continuation = do
  guard context.options.inEffectBlock
  guard (onlyCurriedApplications factory)
  let Tuple target spine = collectGoSpine factory
  foreignTarget <- qualifiedTarget target
  let args = getGoSpineArgs spine
  case directFactoryInfo context foreignTarget (Array.length args) of
    Just { info, retTy } ->
      let
        captured = captureArguments translateExpr context info (nextId + 1) args
        rawCall = GoCall (GoCall (GoRef info.reference) captured.exprs) []
      in
        Just (finishEffectCall translateExpr context nextId resultIdent resultLevel captured.stmts captured.nextId rawCall retTy continuation)
    Nothing -> do
      info <- lookupEffectInfo context Workers.key foreignTarget
      guard (Array.length args == info.arity)
      let
        captured = captureArguments translateExpr context info (nextId + 1) args
        originalName = localId resultIdent resultLevel
        resultName = originalName <> "_" <> show nextId
        bound = Map.insert originalName { name: resultName, goType: TypeValue } context.bound
        bodyContext = context { depth = context.depth + 1, bound = bound, tcoIdent = Nothing }
        cont = translateContinuation translateExpr bodyContext captured.nextId continuation
      pure
        { stmts: captured.stmts <> StmtLeaf (GoAssign resultName (GoCall (GoRef info.reference) captured.exprs)) <> cont.stmts
        , expr: cont.expr
        , exprType: cont.exprType
        , nextId: cont.nextId
        }

-- A Let-bound FFI Effect factory whose local is consumed strictly as an
-- EffectBind action (never escaped, returned, aliased or stored) can keep its
-- native Go thunk `func() R`. The factory still runs once at the Let site, and
-- every EffectBind replays the native thunk without runtime Apply/Func wrappers.
bindThunk :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> Maybe ExprResult
bindThunk translateExpr context nextId actionIdent actionLevel factory body = do
  guard (onlyCurriedApplications factory)
  let Tuple target spine = collectGoSpine factory
  foreignTarget <- qualifiedTarget target
  let args = getGoSpineArgs spine
  { info, retTy } <- directFactoryInfo context foreignTarget (Array.length args)
  guard (usedOnlyAsEffectBindAction actionIdent actionLevel body)
  let
    captured = captureArguments translateExpr context info (nextId + 1) args
    originalName = localId actionIdent actionLevel
    actionName = originalName <> "_" <> show nextId
    bound = Map.insert originalName { name: actionName, goType: TypeFunc [] retTy } context.bound
    resBody = translateExpr (context { depth = context.depth + 1, bound = bound, tcoIdent = Nothing }) captured.nextId body
  pure
    { stmts: captured.stmts <> StmtLeaf (GoAssign actionName (GoCall (GoRef info.reference) captured.exprs)) <> resBody.stmts
    , expr: resBody.expr
    , exprType: resBody.exprType
    , nextId: resBody.nextId
    }

invokeBoundThunk :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> Maybe ExprResult
invokeBoundThunk translateExpr context nextId resultIdent resultLevel action continuation = do
  guard context.options.inEffectBlock
  invoked <- invokeBoundThunkExpr context action
  pure (finishEffectCall translateExpr context nextId resultIdent resultLevel StmtEmpty (nextId + 1) invoked.rawCall invoked.retTy continuation)

invokeBoundThunkExpr :: ExprContext -> TcoExpr -> Maybe { rawCall :: GoExpr, retTy :: GoType, adapted :: { expr :: GoExpr, exprType :: GoType } }
invokeBoundThunkExpr context action = case unwrapTcoExpr action of
  Local ident level -> case Map.lookup (localId ident level) context.bound of
    Just { name, goType: TypeFunc [] retTy } ->
      let rawCall = GoCall (GoVar name) []
      in Just { rawCall, retTy, adapted: adaptThunkCall rawCall retTy }
    _ -> Nothing
  _ -> Nothing

finishEffectCall
  :: TranslateExpr
  -> ExprContext
  -> Int
  -> Maybe Ident
  -> Level
  -> StmtTree
  -> Int
  -> GoExpr
  -> GoType
  -> TcoExpr
  -> ExprResult
finishEffectCall translateExpr context nextId resultIdent resultLevel prefixStmts continuationNextId rawCall retTy continuation =
  let
    used = usesLocal resultIdent resultLevel continuation
    originalName = localId resultIdent resultLevel
    resultName = originalName <> "_" <> show nextId
    adapted = adaptThunkCall rawCall retTy
    callStmt =
      if used then StmtLeaf (GoAssign resultName adapted.expr)
      else StmtLeaf (GoMutate "_" rawCall)
    bound =
      if used then Map.insert originalName { name: resultName, goType: adapted.exprType } context.bound
      else context.bound
    bodyContext = context { depth = context.depth + 1, bound = bound, tcoIdent = Nothing }
    cont = translateContinuation translateExpr bodyContext continuationNextId continuation
  in
    { stmts: prefixStmts <> callStmt <> cont.stmts
    , expr: cont.expr
    , exprType: cont.exprType
    , nextId: cont.nextId
    }

translateContinuation :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
translateContinuation translateExpr bodyContext nextId continuation =
  case ImmediateEffectBranches.emit translateExpr bodyContext nextId continuation of
    Just branchResult -> branchResult
    Nothing ->
      let res = translateExpr bodyContext nextId continuation
      in res { expr = executeIfOpaque continuation res.expr }

adaptThunkCall :: GoExpr -> GoType -> { expr :: GoExpr, exprType :: GoType }
adaptThunkCall rawCall = case _ of
  TypeInterface _ -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Box") [ rawCall ], exprType: TypeValue }
  retTy -> { expr: rawCall, exprType: retTy }

directFactoryInfo :: ExprContext -> { mbMod :: Maybe ModuleName, name :: String } -> Int -> Maybe { info :: FunctionInfo, retTy :: GoType }
directFactoryInfo context foreignTarget argCount = do
  info <- lookupEffectInfo context Workers.factoryKey foreignTarget
  guard (argCount == info.arity)
  case info.fRet of
    TypeFunc [] retTy -> Just { info, retTy }
    _ -> Nothing

lookupEffectInfo :: ExprContext -> (String -> String) -> { mbMod :: Maybe ModuleName, name :: String } -> Maybe FunctionInfo
lookupEffectInfo context keyFn foreignTarget =
  let
    workerKey = keyFn foreignTarget.name
    localOwner owner = String.replaceAll (Pattern ".") (Replacement "_") (unwrap owner) == context.modNameStr
  in
    case foreignTarget.mbMod of
      Just owner | not (localOwner owner) -> Map.lookup (unwrap owner <> "." <> workerKey) context.metadata.globalFunctions
      _ -> Map.lookup workerKey context.metadata.ffiFunctions

captureArguments :: TranslateExpr -> ExprContext -> FunctionInfo -> Int -> Array TcoExpr -> Arguments.Arguments
captureArguments translateExpr context info startId args =
  Arguments.capture { prefix: "__ffi_effect_arg_", markUsed: false }
    (\index argumentId argument ->
      let argumentResult = translateExpr (childContext context Nothing) argumentId argument
      in prepareArgument context (fromMaybe TypeValue (Array.index info.fArgs index)) argument argumentResult)
    startId args

prepareArgument :: ExprContext -> GoType -> TcoExpr -> ExprResult -> ExprResult
prepareArgument context expected argument result =
  let
    prepared = case expected of
      TypeInterface name ->
        let boxed = boxGoExpr context.codegenStateRef context.modNameStr result.expr result.exprType
        in case result.stmts of
          StmtEmpty | scalarLiteral argument result.exprType -> cacheArgument context expected boxed
          _ -> GoCall (GoVar name) [ boxed ]
      TypeValue -> boxGoExpr context.codegenStateRef context.modNameStr result.expr result.exprType
      _ -> coerceGoExpr context.codegenStateRef context.modNameStr result.expr result.exprType expected
  in result { expr = prepared, exprType = expected }

-- Cache the immutable Value-to-interface conversion, never an expression,
-- FFI factory or effect. Keep the emitted literal intact, including -0/NaN.
scalarLiteral :: TcoExpr -> GoType -> Boolean
scalarLiteral argument native = case unwrapTcoExpr argument, native of
  Lit (LitInt _), TypeInt64 -> true
  Lit (LitNumber _), TypeFloat64 -> true
  Lit (LitBoolean _), TypeBool -> true
  Lit (LitString _), TypeString -> true
  Lit (LitChar _), TypeString -> true
  _, _ -> false

cacheArgument :: ExprContext -> GoType -> GoExpr -> GoExpr
cacheArgument context native value = unsafePerformEffect do
  state <- Ref.read context.codegenStateRef
  let
    name = "__gopurs_ffi_argument_" <> context.modNameStr <> "_" <> show state.globalId
    binding = { identifier: generated context.modNameStr name, expression: value, goType: native }
  Ref.modify_ (\current -> current
    { globalId = current.globalId + 1
    , declarations = current.declarations <> [ GoCachedValue binding ]
    }) context.codegenStateRef
  pure (GoCall (GoRef (SymbolReference (generated context.modNameStr ("Get_" <> name)))) [])

usesLocal :: Maybe Ident -> Level -> TcoExpr -> Boolean
usesLocal ident level (TcoExpr _ syntax) = case syntax of
  Local otherIdent otherLevel -> ident == otherIdent && level == otherLevel
  _ -> any (usesLocal ident level) syntax

usedOnlyAsEffectBindAction :: Maybe Ident -> Level -> TcoExpr -> Boolean
usedOnlyAsEffectBindAction actionIdent actionLevel body =
  let usage = inspect body
  in usage.valid && usage.count > 0
  where
  same ident level = ident == actionIdent && level == actionLevel
  empty = { valid: true, count: 0 }
  invalid = { valid: false, count: 0 }
  combine a b = { valid: a.valid && b.valid, count: a.count + b.count }
  inspect (TcoExpr _ syntax) = case syntax of
    Local ident level | same ident level -> invalid
    EffectBind resIdent resLevel action continuation ->
      let
        actionUsage = case unwrapTcoExpr action of
          Local ident level | same ident level -> { valid: true, count: 1 }
          _ -> inspect action
        continuationUsage = if same resIdent resLevel then empty else inspect continuation
      in combine actionUsage continuationUsage
    Let ident level value inner ->
      let
        valueUsage = inspect value
        innerUsage = if same ident level then empty else inspect inner
      in combine valueUsage innerUsage
    Abs args inner ->
      if Array.any (\(Tuple ident level) -> same ident level) (NEA.toArray args) then empty else inspect inner
    UncurriedAbs args inner ->
      if Array.any (\(Tuple ident level) -> same ident level) args then empty else inspect inner
    UncurriedEffectAbs args inner ->
      if Array.any (\(Tuple ident level) -> same ident level) args then empty else inspect inner
    LetRec level bindings inner ->
      if Array.any (\(Tuple ident _) -> same (Just ident) level) (NEA.toArray bindings) then empty
      else foldl (\acc (Tuple _ value) -> combine acc (inspect value)) (inspect inner) bindings
    _ -> foldl (\acc child -> combine acc (inspect child)) empty syntax

-- collectGoSpine also flattens UncurriedApp. Its zero-argument calls are real
-- evaluation steps, so no uncurried boundary may disappear into this worker.
onlyCurriedApplications :: TcoExpr -> Boolean
onlyCurriedApplications expression = case unwrapTcoExpr expression of
  App target _ -> onlyCurriedApplications target
  Var _ -> true
  _ -> false
