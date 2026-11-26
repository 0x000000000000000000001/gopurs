module Gopurs.ArrayTraverseTotal (emit, totalIntegerMaybe) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, foldr)
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.AdtExprs as AdtExprs
import Gopurs.CallAnalysis (collectCurriedAbs)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (extractFuncType, getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, bindParameters, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendSyntax(..), Level, Pair(..))

-- A pure function type alone does not permit short circuiting strict traverse:
-- callbacks can still throw, diverge or call foreign code. This deliberately
-- small proof accepts only total integer/boolean expressions and Maybe leaves.
-- Captured locals are already evaluated immutable values, never calls/getters.
totalIntegerMaybe :: Array Level -> TcoExpr -> Boolean
totalIntegerMaybe parameters expression = case unwrapTcoExpr expression of
  CtorSaturated (Qualified (Just (ModuleName "Data.Maybe")) _) _ _ (Ident "Nothing") [] -> true
  CtorSaturated (Qualified (Just (ModuleName "Data.Maybe")) _) _ _ (Ident "Just") [ Tuple _ value ] -> totalInt parameters value
  Branch branches otherwise -> all (\(Pair condition result) -> totalBoolean parameters condition && totalIntegerMaybe parameters result) branches
    && totalIntegerMaybe parameters otherwise
  _ -> false

totalInt :: Array Level -> TcoExpr -> Boolean
totalInt parameters expression = case unwrapTcoExpr expression of
  Local _ level -> Array.elem level parameters || getExprType expression == Int
  Lit (LitInt _) -> true
  PrimOp (Op1 OpIntNegate value) -> totalInt parameters value
  PrimOp (Op1 OpIntBitNot value) -> totalInt parameters value
  PrimOp (Op2 (OpIntNum operator) left right)
    | operator == OpAdd || operator == OpSubtract || operator == OpMultiply -> totalInt parameters left && totalInt parameters right
  _ -> false

totalBoolean :: Array Level -> TcoExpr -> Boolean
totalBoolean parameters expression = case unwrapTcoExpr expression of
  Local _ _ -> getExprType expression == Boolean
  Lit (LitBoolean _) -> true
  PrimOp (Op1 OpBooleanNot value) -> totalBoolean parameters value
  PrimOp (Op2 (OpIntOrd _) left right) -> totalInt parameters left && totalInt parameters right
  PrimOp (Op2 (OpBooleanOrd _) left right) -> totalBoolean parameters left && totalBoolean parameters right
  PrimOp (Op2 OpBooleanAnd left right) -> totalBoolean parameters left && totalBoolean parameters right
  PrimOp (Op2 OpBooleanOr left right) -> totalBoolean parameters left && totalBoolean parameters right
  _ -> false

emit :: TranslateExpr -> ExprContext -> Int -> Boolean -> Int -> Array TcoExpr -> Maybe ExprResult
emit translate context nextId indexed callbackOffset args = do
  guard (Array.length args == callbackOffset + 2)
  -- A recursive value can be represented by a delayed cell. Its local read is
  -- not the already evaluated scalar assumed by the totality proof.
  guard (Array.null context.recVars)
  callback <- Array.index args callbackOffset
  signature <- extractFuncType callback
  guard (signature.fArgs == (if indexed then [ Int, Int ] else [ Int ]))
  guard (signature.fRet == ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ Int ])
  case unwrapTcoExpr callback of
    Abs parameters result -> do
      let grouped = collectCurriedAbs parameters result
      let parametersArray = NEA.toArray grouped.args
      guard (Array.length parametersArray == Array.length signature.fArgs)
      guard (totalIntegerMaybe (map (\(Tuple _ level) -> level) parametersArray) grouped.body)
      pure (emitTotal translate context nextId callbackOffset args parametersArray grouped.body)
    _ -> Nothing

emitTotal :: TranslateExpr -> ExprContext -> Int -> Int -> Array TcoExpr -> Array (Tuple (Maybe Ident) Level) -> TcoExpr -> ExprResult
emitTotal translate context nextId callbackOffset args parameters result =
  let
    -- The admitted callback is an adjacent lambda, so constructing it has no
    -- evaluation. Other arguments still run in their original order and stage.
    captured = CallArguments.capture { prefix: "traverseTotal_arg_", markUsed: true }
      (\index currentId arg -> if index == callbackOffset then
        { stmts: StmtEmpty, expr: rawGo "true", exprType: TypeBool, nextId: currentId }
        else translate (childContext context Nothing) currentId arg) nextId args
    suffix = show captured.nextId
    valuesName = "traverseTotal_values_" <> suffix
    outputName = "traverseTotal_output_" <> suffix
    indexName = "traverseTotal_index_" <> suffix
    valueName = "traverseTotal_value_" <> suffix
    input = fromMaybe (rawGo "nil") (Array.index captured.exprs (callbackOffset + 1))
    inputType = fromMaybe TypeValue (Array.index captured.exprTypes (callbackOffset + 1))
    source = case inputType of
      TypeNativeArray elementType -> { expr: input, elementType }
      _ -> { expr: GoPrefixOp "*" (GoCall (rawGo "(*[]gopurs_runtime.Value)") [ GoSelector input "UnsafePtr" ]), elementType: TypeValue }
    parameterNames = map (\(Tuple name level) -> localId name level) parameters
    parameterTypes = map (const TypeInt64) parameters
    parameterValues = if Array.length parameters == 2 then
      [ GoCall (GoVar "int64") [ GoVar indexName ], integer (GoVar valueName) source.elementType ]
      else [ integer (GoVar valueName) source.elementType ]
    bindings = Array.zipWith GoAssign parameterNames parameterValues
    boundContext = (childContext context (Just Int))
      { bound = bindParameters (Array.zip parameterNames parameterTypes) context.bound }
    maybeType = ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ Any ]
    constructor name fields =
      let built = AdtExprs.saturated context.codegenStateRef
            (AdtExprs.prepareSaturated context.metadata context.modNameStr (Just (ModuleName "Data.Maybe")) name maybeType)
            { exprs: fields, exprTypes: map (const TypeValue) fields }
      in boxGoExpr context.codegenStateRef context.modNameStr built.expr built.exprType
    failure = constructor "Nothing" []
    success = constructor "Just" [ boxGoExpr context.codegenStateRef context.modNameStr (GoVar outputName) (TypeNativeArray TypeValue) ]
    loop = emitResult translate boundContext captured.nextId failure outputName indexName valuesName result
    body =
      [ GoAssign valuesName source.expr
      , GoAssign outputName (rawGo "[]gopurs_runtime.Value{}")
      , GoForRange (indexName <> ", " <> valueName <> " := range " <> valuesName)
          ([ GoMutate "_" (GoVar indexName) ] <> bindings <> loop.stmts)
      ]
  in
    { stmts: captured.stmts
    , expr: GoCall (GoFuncLit [] body success TypeValue) []
    , exprType: TypeValue
    , nextId: loop.nextId + 1
    }
  where
  integer expr from = coerceGoExpr context.codegenStateRef context.modNameStr expr from TypeInt64

-- Emit the callback's primitive branches directly into the loop. No per-element
-- closure/Maybe object is needed. On failure we may omit later callbacks only
-- because the complete callback body passed the totality proof above.
emitResult :: TranslateExpr -> ExprContext -> Int -> GoExpr -> String -> String -> String -> TcoExpr -> { stmts :: Array GoExpr, nextId :: Int }
emitResult translate context nextId failure outputName indexName valuesName expression = case unwrapTcoExpr expression of
  CtorSaturated _ _ _ (Ident "Nothing") [] -> { stmts: [ GoReturn failure ], nextId }
  CtorSaturated _ _ _ (Ident "Just") [ Tuple _ value ] ->
    let
      translated = translate context nextId value
      boxed = boxGoExpr context.codegenStateRef context.modNameStr translated.expr translated.exprType
    in
      { stmts: flattenStmts translated.stmts <>
          [ GoIfElse (GoBinOp "==" (GoVar indexName) (GoInt 0))
              [ GoMutate outputName (GoCall (GoVar "make") [ rawGo "[]gopurs_runtime.Value", GoCall (GoVar "len") [ GoVar valuesName ] ]) ] []
          , GoMutate (outputName <> "[" <> indexName <> "]") boxed
          ]
      , nextId: translated.nextId
      }
  Branch branches otherwise ->
    foldr (\(Pair condition result) remaining ->
      let
        test = translate (childContext context (Just Boolean)) remaining.nextId condition
        yes = emitResult translate context test.nextId failure outputName indexName valuesName result
      in
        { stmts: flattenStmts test.stmts <> [ GoIfElse test.expr yes.stmts remaining.stmts ], nextId: yes.nextId })
      (emitResult translate context nextId failure outputName indexName valuesName otherwise) branches
  -- Unreachable: the complete tree is admitted before any code is emitted.
  _ -> { stmts: [], nextId }
