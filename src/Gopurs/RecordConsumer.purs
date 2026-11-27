-- A module-local proof, captured by the translator, never published as ABI.
module Gopurs.RecordConsumer (Plan, Plans, prepare, emit) where

import Prelude
import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldl, all)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..), snd)
import Gopurs.AffineLoops as AffineLoops
import Gopurs.ArrayProducerFold (onlyCurried)
import Gopurs.BindingExprs as BindingExprs
import Gopurs.CallAnalysis (extractUncurriedAbs, collectGoSpine, getGoSpineArgs)
import Gopurs.CodePointFold (compatibleT)
import Gopurs.ExprAnalysis (extractFuncType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, ModuleFunctions, TranslateExpr, StmtTree(..), flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo, sanitizeName)
import Gopurs.GoFunctions (namedParameters)
import Gopurs.Printer (printGoExpr)
import Gopurs.RecordAffineLoops as RecordAffineLoops
import Gopurs.RecordShapeGuard as RecordShapeGuard
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Literal(..), Prop(..), Ident(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax as Syn
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendAccessor(..), BackendOperator(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), Level, Pair(..))

type Plan = { keys :: Array String, deltas :: Array Int, counterFirst :: Boolean, recordType :: ExprType }
type Plans = Map.Map String Plan

prepare :: String -> ModuleFunctions -> Array { recursive :: Boolean, bindings :: Array (Tuple Ident TcoExpr) } -> Plans
prepare owner functions groups = Map.fromFoldable (Array.mapMaybe describe groups)
  where
  describe group = do
    guard group.recursive
    Tuple (Ident rawName) value <- case group.bindings of
      [ binding ] -> Just binding
      _ -> Nothing
    let name = sanitizeName rawName
    info <- Map.lookup name functions
    guard (info.arity == 2 && Array.length info.fArgs == 2)
    abs <- extractUncurriedAbs value
    guard (Array.length abs.args == 2)
    signature <- extractFuncType value
    proof <- RecordAffineLoops.flatAdditive (AffineLoops.globalTarget owner name)
      (namedParameters abs.args info.fArgs) info.fRet abs.body
    recordType <- case signature.fArgs of
      [ Int, ty@(Record _) ] | proof.counterFirst -> Just ty
      [ ty@(Record _), Int ] | not proof.counterFirst -> Just ty
      _ -> Nothing
    guard (signature.fRet == recordType && validRecurrence owner name signature.fArgs recordType value)
    guard (not (Array.null proof.keys) && Array.length (Array.nub proof.keys) == Array.length proof.keys)
    pure (Tuple name { keys: proof.keys, deltas: proof.deltas, counterFirst: proof.counterFirst, recordType })

-- Require the complete additive observation, with no other use or intervening
-- computation. Inputs are existing locals: capture order and effects stay at
-- their original call sites. No usage annotation is treated as ownership.
emit :: Plans -> TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> Maybe ExprResult
emit plans translate context nextId ident level binding body = do
  guard (Array.null context.recVars && onlyCurried binding)
  let Tuple target segments = collectGoSpine binding
  name <- case unwrapTcoExpr target of
    Var (Qualified (Just (ModuleName _)) (Ident name))
      | AffineLoops.globalTarget context.modNameStr (sanitizeName name) target -> Just (sanitizeName name)
    _ -> Nothing
  plan <- Map.lookup name plans
  Tuple count input <- case getGoSpineArgs segments of
    [ first, second ] -> Just (if plan.counterFirst then Tuple first second else Tuple second first)
    _ -> Nothing
  inputId <- local input
  countId <- local count
  inputBound <- Map.lookup inputId context.bound
  countBound <- Map.lookup countId context.bound
  guard (inputBound.goType == TypeValue && (countBound.goType == TypeInt64 || countBound.goType == TypeValue))
  let args = if plan.counterFirst then [ Int, plan.recordType ] else [ plan.recordType, Int ]
  guard (compatibleT (Func args plan.recordType) target && compatibleT plan.recordType binding
    && compatibleT plan.recordType input && compatibleT Int count)
  scalar <- case unwrapTcoExpr body of
    EffectPure value | context.options.inEffectBlock && effectInt body -> Just value
    EffectPure _ -> Nothing
    _ -> Just body
  reads <- projections plan.recordType scalar
  let resultId = localId ident level
      updated = map snd (Array.filter (\(Tuple source _) -> source == resultId) reads)
      original = map snd (Array.filter (\(Tuple source _) -> source == inputId) reads)
      keys = Array.sort plan.keys
  guard (resultId /= inputId && Array.length reads == 2 * Array.length keys
    && Array.sort updated == keys && Array.sort original == keys)
  let
    -- The fallback is translated through the ordinary Let path, not this hook.
    fallback = BindingExprs.nonRecursive translate context nextId ident level binding body
    prefix = fresh ("recordConsumer_" <> show nextId) (map (\(Tuple _ value) -> value.name) (Map.toUnfoldable context.bound :: Array (Tuple String { name :: String, goType :: GoType })))
    record = prefix <> "_input"
    counter = if countBound.goType == TypeValue then GoSelector (GoVar countBound.name) "IntVal" else GoVar countBound.name
    fast values =
      let sum = foldl (GoBinOp "+") (GoInt 0) values
          subtotal = GoVar (prefix <> "_sum")
          total = GoBinOp "+" (GoBinOp "+" subtotal subtotal) (RecordAffineLoops.additiveTotal counter plan.deltas)
          result = "int64(int32(" <> printGoExpr total <> "))"
      in prefix <> "_sum := " <> printGoExpr sum <> "\nreturn " <> (if fallback.exprType == TypeValue then "gopurs_runtime.Int(" <> result <> ")" else result)
    admitted = GoBinOp "<" (GoBinOp "-" (GoCall (GoVar "uint64") [ counter ]) (GoInt 1)) (GoInt 2147483647)
    statements = [ GoAssign record (GoVar inputBound.name)
      , GoIfElse admitted [ rawGo (RecordShapeGuard.integerReads prefix record keys fast) ] []
      ] <> flattenStmts fallback.stmts <> [ GoReturn fallback.expr ]
  -- The proved primitive result is Int64 or its existing Value boxing. No
  -- admission check follows fallback translation, which may allocate IDs.
  pure { stmts: StmtEmpty, expr: GoCall (GoFuncBlock [] statements fallback.exprType) [], exprType: fallback.exprType, nextId: fallback.nextId }

local :: TcoExpr -> Maybe String
local value = case unwrapTcoExpr value of
  Local ident level -> Just (localId ident level)
  _ -> Nothing

projections :: ExprType -> TcoExpr -> Maybe (Array (Tuple String String))
projections recordType value = do
  guard (compatibleT Int value)
  case unwrapTcoExpr value of
    PrimOp (Op2 (OpIntNum OpAdd) left right) -> (<>) <$> projections recordType left <*> projections recordType right
    Accessor source (GetProp key) -> do
      guard (compatibleT recordType source)
      name <- local source
      pure [ Tuple name key ]
    _ -> Nothing

fresh :: String -> Array String -> String
fresh prefix names = if Array.any (String.contains (Pattern prefix)) names
  then fresh (prefix <> "_") names else prefix

-- Recheck every annotation inside the proved recurrence. The structural loop
-- proof strips wrappers for matching, but wrappers cannot contradict its facts.
validRecurrence :: String -> String -> Array ExprType -> ExprType -> TcoExpr -> Boolean
validRecurrence owner name arguments recordType = lambdas Map.empty arguments
  where
  lambdas locals remaining value = compatibleT (Func remaining recordType) value && case unwrapTcoExpr value of
    Abs parameters body ->
      let params = NEA.toArray parameters
          count = Array.length params
          bound = Map.union (Map.fromFoldable (Array.zipWith (\(Tuple ident level) ty -> Tuple (localId ident level) ty) params (Array.take count remaining))) locals
          rest = Array.drop count remaining
      in count <= Array.length remaining && if Array.null rest then check bound recordType body else lambdas bound rest body
    _ -> false
  check locals expected value = compatibleT expected value && case unwrapTcoExpr value of
    Local ident level -> Map.lookup (localId ident level) locals == Just expected
    Lit (LitInt _) -> expected == Int
    Accessor source (GetProp _) -> expected == Int && check locals recordType source
    Update source props -> expected == recordType && check locals recordType source && all (\(Prop _ field) -> check locals Int field) props
    PrimOp (Op2 (OpIntNum operator) left right) -> expected == Int && Array.elem operator [ OpAdd, OpSubtract ] && check locals Int left && check locals Int right
    PrimOp (Op2 (OpIntOrd OpEq) left right) -> expected == Boolean && check locals Int left && check locals Int right
    Branch branches final -> check locals expected final && all (\(Pair condition branch) -> check locals Boolean condition && check locals expected branch) branches
    App function args -> expected == recordType && AffineLoops.globalTarget owner name function && compatibleT (Func arguments recordType) function
      && Array.length arguments == NEA.length args && all identity (Array.zipWith (check locals) arguments (NEA.toArray args))
    _ -> false

-- PBO keeps an inner Effect a annotation under the instantiated Effect Int.
-- Only this canonical constructor and an unconstrained scalar slot refine;
-- concrete contradictory effect payloads retain the ordinary path.
effectInt :: TcoExpr -> Boolean
effectInt (TcoExpr _ syntax) = case syntax of
  Typed ty value -> valid ty && effectInt value
  Syn.TypeApp value _ -> effectInt value
  _ -> true
  where
  valid = case _ of
    Any -> true
    TypeVar _ -> true
    ADT "Effect.Effect" [ "Effect", "Effect" ] [ Int ] -> true
    ADT "Effect.Effect" [ "Effect", "Effect" ] [ Any ] -> true
    ADT "Effect.Effect" [ "Effect", "Effect" ] [ TypeVar _ ] -> true
    _ -> false
