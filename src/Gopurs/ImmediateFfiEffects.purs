module Gopurs.ImmediateFfiEffects (translate) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Foldable (any)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as Arguments
import Gopurs.ExprAnalysis (executeIfOpaque, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext)
import Gopurs.FfiBridge.ImmediateEffects as Workers
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(App, EffectBind, Local, Var), Level)

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
        (\_ argumentId argument ->
          let argumentResult = translateExpr (childContext context Nothing) argumentId argument
          in argumentResult { expr = boxGoExpr context.codegenStateRef context.modNameStr argumentResult.expr argumentResult.exprType, exprType = TypeValue })
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
  _ -> Nothing

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
