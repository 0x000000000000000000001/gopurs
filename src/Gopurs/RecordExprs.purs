module Gopurs.RecordExprs
  ( RecordExpr
  , prepareLiteral
  , coerceLiteralField
  , literal
  , getProp
  , update
  ) where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CodegenState (CodegenState)
import Gopurs.GoAst (GoExpr(..), GoType(..), sanitizeName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr, unboxGoExpr)
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

type RecordExpr =
  { expr :: GoExpr
  , exprType :: GoType
  }

prepareLiteral :: Ref CodegenState -> String -> ExprType -> Maybe ExprType -> { recordType :: GoType, fields :: Map String ExprType }
prepareLiteral codegenStateRef modNameStr baseExprType mbExpectedExprType =
  let
    exprType = case baseExprType of
      Record _ -> baseExprType
      _ -> fromMaybe baseExprType mbExpectedExprType
    mbRecordType = case exprType of
      Record (Row fields _) -> Just fields
      _ -> Nothing
    recordType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr exprType
    recordFields = case mbRecordType of
      Just fields -> Map.fromFoldable fields
      Nothing -> Map.empty
  in
    { recordType, fields: recordFields }

-- CodeGen calls this immediately after translating each field. Delaying these
-- conversions until all fields are translated would change reboxing state.
coerceLiteralField :: Ref CodegenState -> String -> String -> GoType -> RecordExpr -> Tuple String GoExpr
coerceLiteralField codegenStateRef modNameStr key recordType value =
  let
    expectedGoType = case recordType of
      TypeRecord fields -> fromMaybe TypeValue (Map.lookup key (Map.fromFoldable fields))
      _ -> TypeValue
    coercedVal = coerceGoExpr codegenStateRef modNameStr value.expr value.exprType expectedGoType
  in
    Tuple key coercedVal

literal :: GoType -> Array (Tuple String GoExpr) -> RecordExpr
literal recordType fields = { expr: GoRecordDict recordType fields, exprType: recordType }

getProp :: Ref CodegenState -> String -> String -> RecordExpr -> RecordExpr
getProp codegenStateRef modNameStr prop obj = case obj.exprType of
  TypeRecord fields ->
    let
      fieldGoType = fromMaybe TypeValue (Map.lookup prop (Map.fromFoldable fields))
    in
      { expr: GoStructAccess obj.expr (sanitizeName prop), exprType: fieldGoType }
  TypeStructPointer _ fullName _ _ ->
    let
      h = unsafePerformEffect (Ref.read codegenStateRef)
    in
      case Map.lookup fullName h.classDeclsFields of
        Just info ->
          case Array.findIndex (\f -> f.name == prop) info.fields of
            Just idx ->
              let
                unboxedObj = unboxGoExpr codegenStateRef modNameStr obj.expr obj.exprType obj.exprType
                fieldExpr = GoStructAccess unboxedObj ("V" <> show idx)
                boxedFieldExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Box") [ fieldExpr ]
              in
                { expr: boxedFieldExpr, exprType: TypeValue }
            Nothing -> genericGetProp codegenStateRef modNameStr prop obj
        Nothing -> genericGetProp codegenStateRef modNameStr prop obj
  _ -> genericGetProp codegenStateRef modNameStr prop obj

genericGetProp :: Ref CodegenState -> String -> String -> RecordExpr -> RecordExpr
genericGetProp codegenStateRef modNameStr prop obj =
  { expr: GoRecordAccess (boxGoExpr codegenStateRef modNameStr obj.expr obj.exprType) prop, exprType: TypeValue }

-- Updates reach this emitter only after CodeGen translates the object and every
-- new value. Keep the supplied order, including duplicate labels.
update :: Ref CodegenState -> String -> RecordExpr -> Array { key :: String, expr :: GoExpr, goType :: GoType } -> RecordExpr
update codegenStateRef modNameStr obj props = case obj.exprType of
  TypeRecord fields ->
    let
      coercedUpdates = map
        ( \p ->
            let
              expectedGoType = fromMaybe TypeValue (Map.lookup p.key (Map.fromFoldable fields))
              coercedVal = coerceGoExpr codegenStateRef modNameStr p.expr p.goType expectedGoType
            in
              Tuple p.key coercedVal
        )
        props
    in
      { expr: GoRecordUpdateNative obj.exprType obj.expr coercedUpdates, exprType: obj.exprType }
  _ ->
    let
      boxedExprs = map (\p -> Tuple p.key (boxGoExpr codegenStateRef modNameStr p.expr p.goType)) props
    in
      { expr: GoRecordUpdateDict (boxGoExpr codegenStateRef modNameStr obj.expr obj.exprType) boxedExprs, exprType: TypeValue }
