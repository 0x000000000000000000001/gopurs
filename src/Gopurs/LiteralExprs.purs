module Gopurs.LiteralExprs
  ( array
  , record
  ) where

import Prelude

import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Gopurs.ExprAnalysis (bindFieldFunctionParameters)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext)
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, rawGo)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoTypes (exprTypeToGenericGoType, exprTypeToGoType)
import Gopurs.Printer (printGoExpr)
import Gopurs.RecordExprs as RecordExprs
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Prop(..))

-- Array elements are translated in source order without an expected child type.
-- Choose their common representation only afterwards, then adapt every element.
array :: TranslateExpr -> ExprContext -> Int -> Array TcoExpr -> ExprResult
array translateExpr context@{ codegenStateRef, modNameStr } nextId elements =
  let
    child = childContext context Nothing
    values = Array.foldl
      (\acc element ->
        let result = translateExpr child acc.nextId element
        in
          { stmts: acc.stmts <> result.stmts
          , exprs: Array.snoc acc.exprs result.expr
          , exprTypes: Array.snoc acc.exprTypes result.exprType
          , nextId: result.nextId
          })
      { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
      elements
  in
    case nativeElementType context values.exprTypes of
      Just elementType ->
        let
          arrayType = TypeNativeArray elementType
          converted = Array.zipWith (\expr ty -> coerceGoExpr codegenStateRef modNameStr expr ty elementType)
            values.exprs values.exprTypes
        in
          { stmts: values.stmts
          , expr: rawGo (goTypeToStr arrayType <> "{" <> String.joinWith ", " (map printGoExpr converted) <> "}")
          , exprType: arrayType
          , nextId: values.nextId
          }
      Nothing ->
        let boxed = Array.zipWith (boxGoExpr codegenStateRef modNameStr) values.exprs values.exprTypes
        in
          { stmts: values.stmts
          , expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Array")
              [ rawGo ("[]gopurs_runtime.Value{" <> String.joinWith ", " (map printGoExpr boxed) <> "}") ]
          , exprType: TypeValue
          , nextId: values.nextId
          }

-- An annotation may specialize a uniform layout, including an empty array.
-- Mixed source representations still use Value even with an expected type.
nativeElementType :: ExprContext -> Array GoType -> Maybe GoType
nativeElementType { metadata, modNameStr, mbExpectedExprType } elementTypes =
  let
    first = Array.head elementTypes
    uniform = Array.all (\ty -> Just ty == first) elementTypes
    expected = case mbExpectedExprType of
      Just ty -> case exprTypeToGenericGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors [] modNameStr ty of
        TypeNativeArray elementType -> Just elementType
        _ -> Nothing
      Nothing -> Nothing
    selected = case expected of
      Just ty -> Just ty
      Nothing -> if uniform then first else Nothing
  in
    case selected of
      Just ty | uniform && ty /= TypeValue -> Just ty
      _ -> Nothing

-- Record labels determine translation order. Bind callback parameters from the
-- field annotation, then translate and coerce that field before the next one:
-- coercion may register Rebox helpers in the shared state.
record :: TranslateExpr -> ExprContext -> Int -> ExprType -> Array (Prop TcoExpr) -> ExprResult
record translateExpr context@{ metadata, codegenStateRef, modNameStr, bound, mbExpectedExprType } nextId baseType props =
  let
    child = childContext context Nothing
    toGoType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr
    info = RecordExprs.prepareLiteral metadata modNameStr baseType mbExpectedExprType
    fields = Array.foldl
      (\acc (Prop key value) ->
        let
          expected = fromMaybe Any (Map.lookup key info.fields)
          fieldBound = bindFieldFunctionParameters toGoType bound expected value
          valueResult = translateExpr (child { bound = fieldBound }) acc.nextId value
          field = RecordExprs.coerceLiteralField codegenStateRef modNameStr key info.recordType
            { expr: valueResult.expr, exprType: valueResult.exprType }
        in
          { stmts: acc.stmts <> valueResult.stmts, exprs: Array.snoc acc.exprs field, nextId: valueResult.nextId })
      { stmts: StmtEmpty, exprs: [], nextId }
      (Array.sortBy (comparing \(Prop key _) -> key) props)
    result = RecordExprs.literal info.recordType fields.exprs
  in
    { stmts: fields.stmts, expr: result.expr, exprType: result.exprType, nextId: fields.nextId }
