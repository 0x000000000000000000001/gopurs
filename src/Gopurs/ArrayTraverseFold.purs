module Gopurs.ArrayTraverseFold (emit, justMapper) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (evalState)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldM)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.IndexedAddLanes as IndexedAddLanes
import Gopurs.ArrayProducerFold (onlyCurried)
import Gopurs.ArrayTraverse (recognizeMaybe)
import Gopurs.ArrayTraverseTotal (totalIntegerMaybe)
import Gopurs.CallAnalysis (collectCurriedAbs, collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.ClosedIndexedFolds (closedPrimitiveCallback, validDescriptor)
import Gopurs.CodegenState (ArrayFoldDescriptor)
import Gopurs.ExprAnalysis (extractFuncType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, bindParameters, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.ImmediateApplications.Scope (freshen, maximumLevel)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), ModuleName(..), ProperName(..), Qualified(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendSyntax(..), Level, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syntax

maybeInt :: ExprType
maybeInt = ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ Int ]

-- The callback must construct Just directly. Its payload is then checked by
-- the existing closed scalar proof, including every nested Typed annotation.
-- Nothing/branches/captures/foreign calls retain the ordinary traversal.
justMapper :: Int -> TcoExpr -> Maybe NeutralExpr
justMapper arity callback = do
  signature <- extractFuncType callback
  guard (signature.fArgs == Array.replicate arity Int && signature.fRet == maybeInt)
  case unwrapTcoExpr callback of
    Abs parameters body -> do
      let grouped = collectCurriedAbs parameters body
      guard (NEA.length grouped.args == arity && totalIntegerMaybe (map (\(Tuple _ level) -> level) (NEA.toArray grouped.args)) grouped.body)
    _ -> Nothing
  mapped <- convert arity callback
  guard (closedPrimitiveCallback arity mapped)
  pure mapped
  where
  convert remaining value = do
    guard (annotated (Func (Array.replicate remaining Int) maybeInt) value)
    case unwrapTcoExpr value of
      Abs parameters body -> do
        let count = NEA.length parameters
        guard (count <= remaining)
        scalar <- if count < remaining then convert (remaining - count) body else do
          guard (annotated maybeInt body)
          case unwrapTcoExpr body of
            CtorSaturated (Qualified (Just (ModuleName "Data.Maybe")) (Ident "Just")) _ (ProperName "Maybe") (Ident "Just") [ Tuple _ payload ] ->
              pure (neutral payload)
            _ -> Nothing
        pure (NeutralExpr (Typed (Func (Array.replicate remaining Int) Int) (NeutralExpr (Abs parameters scalar))))
      _ -> Nothing

-- TAST can retain generic m (t b) around a canonical traversal. Refine those
-- transparent applications against the proved result, but reject any concrete
-- contradiction, including nested Maybe String under an outer Maybe Int.
annotated :: ExprType -> TcoExpr -> Boolean
annotated expected (TcoExpr _ syntax) = case syntax of
  Typed actual value -> compatible expected actual && annotated expected value
  Syntax.TypeApp value _ -> annotated expected value
  _ -> true

compatible :: ExprType -> ExprType -> Boolean
compatible expected actual = isJust (refine expected actual Map.empty)

refine :: ExprType -> ExprType -> Map.Map String ExprType -> Maybe (Map.Map String ExprType)
refine expected actual variables = case expected, actual of
  _, Any -> Just variables
  _, TypeVar name -> bindVariable name expected variables
  Array wanted, Array supplied -> refine wanted supplied variables
  Array wanted, ADT "Prim.Array" [ "Prim", "Array" ] [ supplied ] -> refine wanted supplied variables
  Array wanted, CoreFn.TypeApp (TypeVar name) [ supplied ] -> do
    known <- bindVariable name (ADT "Prim.Array" [ "Prim", "Array" ] []) variables
    refine wanted supplied known
  ADT name path arguments, ADT suppliedName suppliedPath suppliedArguments -> do
    guard (name == suppliedName && path == suppliedPath)
    sameArguments arguments suppliedArguments variables
  ADT name path [ wanted ], CoreFn.TypeApp (TypeVar variable) [ supplied ] -> do
    known <- bindVariable variable (ADT name path []) variables
    refine wanted supplied known
  Func arguments result, Func suppliedArguments suppliedResult -> do
    known <- sameArguments arguments suppliedArguments variables
    refine result suppliedResult known
  _, _ | expected == actual -> Just variables
  _, _ -> Nothing
  where
  sameArguments wanted supplied known = do
    guard (Array.length wanted == Array.length supplied)
    foldM (\current (Tuple left right) -> refine left right current) known (Array.zip wanted supplied)
  bindVariable name value known = case Map.lookup name known of
    Just previous | previous /= value -> Nothing
    _ -> Just (Map.insert name value known)

emit :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> Maybe ExprResult
emit translate context nextId mbIdent level producer continuation = do
  guard (Array.null context.recVars && onlyCurried producer)
  let Tuple function spine = collectGoSpine producer
  -- Start with direct class-method forms. FFI methods may contain their own
  -- application stages and are not flattened by this fusion.
  guard (directTraversal function)
  traversal <- recognizeMaybe function (getGoSpineArgs spine)
  guard (Array.length traversal.args == traversal.callbackOffset + 2)
  callback <- Array.index traversal.args traversal.callbackOffset
  input <- Array.index traversal.args (traversal.callbackOffset + 1)
  guard (annotated (Array Int) input)
  guard (annotated (ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ Array Int ]) producer)
  mapper <- justMapper (if traversal.indexed then 2 else 1) callback
  branch <- case unwrapTcoExpr continuation of
    EffectPure value -> do
      -- Do not move traversal work into or out of a deferred effect value.
      guard context.options.inEffectBlock
      guard (annotated (ADT "Effect.Effect" [ "Effect", "Effect" ] [ Int ]) continuation)
      pure value
    _ -> pure continuation
  descriptor <- consumedArray context level branch
  let whole = NeutralExpr (Let mbIdent level (neutral producer) (neutral continuation))
  pure (emitLoop translate context nextId traversal mapper descriptor (maximumLevel whole))

directTraversal :: TcoExpr -> Boolean
directTraversal function = case unwrapTcoExpr function of
  Accessor _ (GetProp "traverse") -> true
  _ -> case qualifiedTarget function of
    Just { mbMod: Just (ModuleName "Data.Traversable"), name: "traverse" } -> true
    Just { mbMod: Just (ModuleName "Data.TraversableWithIndex"), name } -> name == "traverseWithIndex" || name == "traverseWithIndexDefault"
    _ -> false

consumedArray :: ExprContext -> Level -> TcoExpr -> Maybe ArrayFoldDescriptor
consumedArray context level expression = do
  guard (annotated Int expression)
  success <- case unwrapTcoExpr expression of
    Branch branches otherwise -> do
      guard (case unwrapTcoExpr otherwise of
        Fail _ -> true
        _ -> false)
      case NEA.toArray branches of
        [ Pair first firstResult, Pair second secondResult ]
          | tagged "Nothing" first && tagged "Just" second -> Just secondResult
          | tagged "Just" first && tagged "Nothing" second -> Just firstResult
        _ -> Nothing
    _ -> Nothing
  guard (annotated Int success && onlyCurried success)
  let Tuple fn spine = collectGoSpine success
  guard (annotated (Func [ Array Int ] Int) fn)
  case getGoSpineArgs spine of
    [ value ] -> do
      guard (annotated (Array Int) value)
      guard (case unwrapTcoExpr value of
        Accessor object (GetCtorField (Qualified (Just (ModuleName "Data.Maybe")) (Ident "Just")) _ (ProperName "Maybe") (Ident "Just") _ 0) -> sameLocal object
        _ -> false)
    _ -> Nothing
  target <- qualifiedTarget fn
  let local = target.mbMod == Nothing || map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) target.mbMod == Just context.modNameStr
  info <- if local then Map.lookup target.name context.moduleFunctions else do
    owner <- target.mbMod
    Map.lookup (unwrap owner <> "." <> target.name) context.metadata.globalFunctions
  guard (info.arity == 1 && info.fArgs == [ TypeNativeArray TypeInt64 ] && info.fRet == TypeInt64)
  descriptor <- info.arrayFold
  guard (validDescriptor descriptor)
  pure descriptor
  where
  sameLocal object = annotated (ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ Array Int ]) object
    && case unwrapTcoExpr object of
      Local _ actual -> actual == level
      _ -> false
  tagged tag condition = annotated Boolean condition && case unwrapTcoExpr condition of
    PrimOp (Op1 (OpIsTag (Qualified (Just (ModuleName "Data.Maybe")) (Ident name))) object) -> name == tag && sameLocal object
    _ -> false

emitLoop :: TranslateExpr -> ExprContext -> Int -> { indexed :: Boolean, callbackOffset :: Int, args :: Array TcoExpr } -> NeutralExpr -> ArrayFoldDescriptor -> Int -> ExprResult
emitLoop translate context nextId traversal mapper descriptor lastLevel =
  let
    start = max lastLevel (maximumLevel descriptor.callback) + nextId + 1
    folded = Tco.analyze [] (evalState (freshen descriptor.callback) start)
    mapped = Tco.analyze [] mapper
    native currentId callback = case unwrapTcoExpr callback of
      Abs parameters callbackBody ->
        let
          grouped = collectCurriedAbs parameters callbackBody
          params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) (NEA.toArray grouped.args)
          callbackContext = context { depth = context.depth + 1, bound = bindParameters params context.bound
            , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
            , mbExpectedExprType = Just Int, mbExpectedGoType = Just TypeInt64 }
          result = translate callbackContext currentId grouped.body
        in { expr: GoFuncLit params (flattenStmts result.stmts) (integer result.expr result.exprType) TypeInt64, nextId: result.nextId }
      _ -> { expr: rawGo "nil", nextId: currentId }
    captured = CallArguments.capture { prefix: "traverseFold_arg_", markUsed: true }
      (\position currentId argument -> if position == traversal.callbackOffset then
        { stmts: StmtEmpty, expr: rawGo "true", exprType: TypeBool, nextId: currentId }
        else translate (childContext context Nothing) currentId argument) nextId traversal.args
    mapResult = native captured.nextId mapped
    foldResult = native mapResult.nextId folded
    suffix = show foldResult.nextId
    variable label = "traverseFold_" <> label <> "_" <> suffix
    arg position = fromMaybe (rawGo "nil") (Array.index captured.exprs position)
    source = arraySource (case arg (traversal.callbackOffset + 1) of
      GoVar name -> name
      _ -> "") (variable "source") (fromMaybe TypeValue (Array.index captured.exprTypes (traversal.callbackOffset + 1)))
    index = GoCall (GoVar "int64") [ GoVar (variable "index") ]
    value = integer (GoVar (variable "value")) source.elementType
    mappedValue = GoCall (GoVar (variable "map")) (if traversal.indexed then [ index, value ] else [ value ])
    ordinary = GoForRange (variable "index" <> ", " <> variable "value" <> " := range " <> source.target)
      [ GoMutate (variable "result") (GoCall (GoVar (variable "fold")) [ index, GoVar (variable "result"), mappedValue ]) ]
    loop = if IndexedAddLanes.eligible mapped folded then
      IndexedAddLanes.emit { source: source.target, prefix: variable "lane", result: variable "result" }
        (\i accumulator item ->
          let current = GoCall (GoVar "int64") [ i ]
              produced = GoCall (GoVar (variable "map")) (if traversal.indexed then [ current, integer item source.elementType ] else [ integer item source.elementType ])
          in GoCall (GoVar (variable "fold")) [ current, accumulator, produced ]) ordinary
      else ordinary
    body =
      [ GoAssign (variable "map") mapResult.expr, GoAssign (variable "fold") foldResult.expr
      , GoAssign (variable "result") (GoCall (GoVar "int64") [ GoInt descriptor.seed ])
      , source.assignment
      , loop
      ]
  in { stmts: captured.stmts, expr: GoCall (GoFuncLit [] body (GoVar (variable "result")) TypeInt64) []
     , exprType: TypeInt64, nextId: foldResult.nextId + 1 }
  where
  integer value from = coerceGoExpr context.codegenStateRef context.modNameStr value from TypeInt64

neutral :: TcoExpr -> NeutralExpr
neutral (TcoExpr _ syntax) = NeutralExpr (map neutral syntax)
