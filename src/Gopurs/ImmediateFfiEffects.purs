module Gopurs.ImmediateFfiEffects (translate, direct) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Foldable (any)
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
import Gopurs.ExprAnalysis (executeIfOpaque, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext)
import Gopurs.FfiBridge.ImmediateEffects as Workers
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr)
import Gopurs.GoSymbol (GoReference(..), generated)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident, Literal(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(App, EffectBind, Lit, Local, Var), Level)

-- A directly adjacent bind is the only permitted consumer. In particular, a
-- captured, replayed or saved action keeps the existing boxed representation.
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
-- the native factory and then its thunk retain their original order. Worker
-- eligibility and ABI conversion are exactly the same as for the Let form.
direct :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> Maybe ExprResult
direct translateExpr context nextId resultIdent resultLevel factory continuation = do
  guard context.options.inEffectBlock
  guard (onlyCurriedApplications factory)
  let Tuple target spine = collectGoSpine factory
  foreignTarget <- qualifiedTarget target
  let
    args = getGoSpineArgs spine
    workerKey = Workers.key foreignTarget.name
    localOwner owner = String.replaceAll (Pattern ".") (Replacement "_") (unwrap owner) == context.modNameStr
  info <- case foreignTarget.mbMod of
    Just owner | not (localOwner owner) -> Map.lookup (unwrap owner <> "." <> workerKey) context.metadata.globalFunctions
    _ -> Map.lookup workerKey context.metadata.ffiFunctions
  guard (Array.length args == info.arity)
  let
    captured = Arguments.capture { prefix: "__ffi_effect_arg_", markUsed: false }
      (\index argumentId argument ->
        let argumentResult = translateExpr (childContext context Nothing) argumentId argument
        in prepareArgument context (fromMaybe TypeValue (Array.index info.fArgs index)) argument argumentResult)
      (nextId + 1) args
    originalName = localId resultIdent resultLevel
    resultName = originalName <> "_" <> show nextId
    bound = Map.insert originalName { name: resultName, goType: TypeValue } context.bound
    result = translateExpr (context { depth = context.depth + 1, bound = bound, tcoIdent = Nothing }) captured.nextId continuation
  pure
    { stmts: captured.stmts <> StmtLeaf (GoAssign resultName (GoCall (GoRef info.reference) captured.exprs)) <> result.stmts
    , expr: executeIfOpaque continuation result.expr
    , exprType: result.exprType
    , nextId: result.nextId
    }

prepareArgument :: ExprContext -> GoType -> TcoExpr -> ExprResult -> ExprResult
prepareArgument context expected argument result =
  let
    boxed = boxGoExpr context.codegenStateRef context.modNameStr result.expr result.exprType
    prepared = case expected of
      TypeInterface name -> case result.stmts of
        StmtEmpty | scalarLiteral argument result.exprType -> cacheArgument context expected boxed
        _ -> GoCall (GoVar name) [ boxed ]
      _ -> boxed
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

-- collectGoSpine also flattens UncurriedApp. Its zero-argument calls are real
-- evaluation steps, so no uncurried boundary may disappear into this worker.
onlyCurriedApplications :: TcoExpr -> Boolean
onlyCurriedApplications expression = case unwrapTcoExpr expression of
  App target _ -> onlyCurriedApplications target
  Var _ -> true
  _ -> false
