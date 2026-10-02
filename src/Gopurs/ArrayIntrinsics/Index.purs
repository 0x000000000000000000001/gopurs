module Gopurs.ArrayIntrinsics.Index (safeIndex, unsafeIndex) where

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.CallAnalysis (CallTarget)
import Gopurs.CallArguments (applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (getExprType)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Typed))

type IndexCall =
  { justArg :: TcoExpr, nothingArg :: TcoExpr, arrayArg :: TcoExpr, indexArg :: TcoExpr }

-- indexImpl's FFI bridge converts its entire array to []any. Retain the
-- source representation and convert only the element passed to Just.
safeIndex :: TranslateExpr -> ExprContext -> Int -> Maybe CallTarget -> Array TcoExpr -> Maybe ExprResult
safeIndex translate context nextId target args =
  emitSafeIndex translate context nextId <$> recognizeSafeIndex context.modNameStr target args

recognizeSafeIndex :: String -> Maybe CallTarget -> Array TcoExpr -> Maybe IndexCall
recognizeSafeIndex modNameStr target args = case target, args of
  Just { mbMod, name: "indexImpl" }, [ justArg, nothingArg, arrayArg, indexArg ]
    | mbMod == Just (ModuleName "Data.Array") || (mbMod == Nothing && modNameStr == "Data_Array") ->
      Just { justArg, nothingArg, arrayArg, indexArg }
  _, _ -> Nothing

emitSafeIndex :: TranslateExpr -> ExprContext -> Int -> IndexCall -> ExprResult
emitSafeIndex translate context@{ codegenStateRef, modNameStr } nextId { justArg, nothingArg, arrayArg, indexArg } =
  let
    captured = CallArguments.capture { prefix: "arrayIndex_arg_", markUsed: false }
      (\_ -> translate (childContext context Nothing)) nextId
      [ justArg, nothingArg, withoutArrayAnnotation arrayArg, indexArg ]
    arg n = fromMaybe (rawGo "nil") (Array.index captured.exprs n)
    argType n = fromMaybe TypeValue (Array.index captured.exprTypes n)
    boxed n = boxGoExpr codegenStateRef modNameStr (arg n) (argType n)
    arrayName = "arrayIndex_source_" <> show captured.nextId
    -- The captured argument is a variable, so this binding cannot repeat
    -- evaluation or traverse the array.
    source = arraySource "arrayIndex_value" arrayName (argType 2)
    elementType = selectedType context arrayArg source.elementType
    index = coerceGoExpr codegenStateRef modNameStr (arg 3) (argType 3) TypeInt64
    element = GoIndex (rawGo ("(" <> source.target <> ")")) index
    converted = coerceGoExpr codegenStateRef modNameStr element source.elementType elementType
    selected = boxGoExpr codegenStateRef modNameStr converted elementType
    onFound = case argType 0 of
      TypeFunc [ inputType ] outputType -> boxGoExpr codegenStateRef modNameStr
        (GoCall (arg 0) [ coerceGoExpr codegenStateRef modNameStr converted elementType inputType ]) outputType
      _ -> applyBoxed (boxed 0) [ selected ]
    inBounds = GoBinOp "&&"
      (GoBinOp ">=" index (rawGo "0"))
      (GoBinOp "<" index (GoCall (GoVar "int64") [ GoCall (GoVar "len") [ rawGo source.target ] ]))
    expr = GoCall (GoFuncLit []
      [ GoAssign "arrayIndex_value" (arg 2)
      , source.assignment
      , GoIfElse inBounds [ GoReturn onFound ] []
      ] (boxed 1) TypeValue) []
  in
    { stmts: captured.stmts, expr, exprType: TypeValue, nextId: captured.nextId }

-- OpArrayIndex requires an in-bounds index. Its source binding precedes the
-- index's statements; unlike safeIndex, it does not defer both into an IIFE.
unsafeIndex :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> TcoExpr -> ExprResult
unsafeIndex translate context@{ codegenStateRef, modNameStr } nextId arrayArg indexArg =
  let
    child = childContext context Nothing
    array = translate child nextId (withoutArrayAnnotation arrayArg)
    arrayName = "arrayUnsafe_value_" <> show array.nextId
    sourceName = "arrayUnsafe_source_" <> show array.nextId
    source = arraySource arrayName sourceName array.exprType
    index = translate child (array.nextId + 1) indexArg
    indexName = "arrayUnsafe_index_" <> show index.nextId
    elementType = selectedType context arrayArg source.elementType
    selected = GoIndex (rawGo ("(" <> source.target <> ")")) (GoVar indexName)
    converted = coerceGoExpr codegenStateRef modNameStr selected source.elementType elementType
  in
    { stmts: array.stmts
        <> StmtLeaf (GoAssign arrayName array.expr)
        <> StmtLeaf source.assignment
        <> index.stmts
        <> StmtLeaf (GoAssign indexName (coerceGoExpr codegenStateRef modNameStr index.expr index.exprType TypeInt64))
    , expr: boxGoExpr codegenStateRef modNameStr converted elementType
    , exprType: TypeValue
    , nextId: index.nextId + 1
    }

-- Typed array coercions are element-wise. Moving this coercion onto the
-- selected element preserves its TAST type without copying the container.
withoutArrayAnnotation :: TcoExpr -> TcoExpr
withoutArrayAnnotation (TcoExpr _ (Typed (Array _) inner)) = withoutArrayAnnotation inner
withoutArrayAnnotation arg = arg

selectedType :: ExprContext -> TcoExpr -> GoType -> GoType
selectedType { metadata, modNameStr } arrayArg fallback = case getExprType arrayArg of
  Array inner -> exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr inner
  _ -> fallback
