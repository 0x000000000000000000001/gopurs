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
import Gopurs.CodegenState (CodegenMetadata, CodegenState)
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, rawGo, sanitizeName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr, unboxGoExpr)
import Gopurs.GoTypes (exprTypeToGoType, instantiateGenericGoType, structFieldGoType, visibleRecordFields)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

type RecordExpr =
  { expr :: GoExpr
  , exprType :: GoType
  }

prepareLiteral :: CodegenMetadata -> String -> ExprType -> Maybe ExprType -> { recordType :: GoType, fields :: Map String ExprType }
prepareLiteral metadata modNameStr baseExprType mbExpectedExprType =
  let
    exprType = case baseExprType of
      Record _ -> baseExprType
      _ -> fromMaybe baseExprType mbExpectedExprType
    mbRecordType = case exprType of
      Record (Row fields _) -> Just fields
      _ -> Nothing
    recordType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr exprType
    recordFields = case mbRecordType of
      Just fields -> Map.fromFoldable (visibleRecordFields fields)
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

getProp :: CodegenMetadata -> Ref CodegenState -> String -> String -> RecordExpr -> RecordExpr
getProp metadata codegenStateRef modNameStr prop obj = case obj.exprType of
  TypeRecord fields ->
    let
      fieldGoType = fromMaybe TypeValue (Map.lookup prop (Map.fromFoldable fields))
    in
      { expr: GoStructAccess obj.expr (sanitizeName prop), exprType: fieldGoType }
  TypeStructPointer { fullName, typeArgs } ->
    case Map.lookup fullName metadata.classDeclsFields of
      Just info ->
        case Array.find (\(Tuple _ field) -> field.name == prop) (Array.mapWithIndex Tuple info.fields) of
          Just (Tuple idx field) ->
            let
              typeEnv = Map.fromFoldable (Array.zip info.vars typeArgs)
              genericFieldType = structFieldGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors info.vars modNameStr field."type"
              fieldGoType = instantiateGenericGoType typeEnv genericFieldType
              unboxedObj = unboxGoExpr codegenStateRef modNameStr obj.expr obj.exprType obj.exprType
              fieldExpr = GoStructAccess unboxedObj ("V" <> show idx)
              boxedFieldExpr = boxGoExpr codegenStateRef modNameStr fieldExpr fieldGoType
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
update :: Ref CodegenState -> String -> Maybe GoType -> RecordExpr -> Array { key :: String, expr :: GoExpr, goType :: GoType } -> RecordExpr
update codegenStateRef modNameStr expectedType obj props = case obj.exprType of
  TypeRecord fields ->
    let
      resultFields = case expectedType of
        Just (TypeRecord targetFields) -> targetFields
        _ -> map (\(Tuple key ty) -> Tuple key
          (fromMaybe ty (map _.goType (Array.last (Array.filter (\p -> p.key == key) props))))) fields
      resultType = TypeRecord resultFields
      coercedUpdates = map
        ( \p ->
            let
              expectedGoType = fromMaybe TypeValue (Map.lookup p.key (Map.fromFoldable resultFields))
              coercedVal = coerceGoExpr codegenStateRef modNameStr p.expr p.goType expectedGoType
            in
              Tuple p.key coercedVal
        )
        props
      unchangedFields = Array.filter (\(Tuple key _) -> not (Array.any (\p -> p.key == key) props)) resultFields
      copyField (Tuple key targetType) =
        let sourceType = fromMaybe TypeValue (Map.lookup key (Map.fromFoldable fields))
        in GoMutate ("clone." <> sanitizeName key)
          (coerceGoExpr codegenStateRef modNameStr (GoStructAccess (GoVar "originalRecord") key) sourceType targetType)
      -- A type-changing update needs a new native layout. Do not convert the
      -- overwritten fields through their old types while copying the record.
      changedLayout = GoCall (GoFuncLit []
        ([ GoAssign "originalRecord" obj.expr
         , rawGo "_ = originalRecord"
         , rawGo ("var clone " <> goTypeToStr resultType)
         ] <> map copyField unchangedFields
           <> map (\(Tuple key value) -> GoMutate ("clone." <> sanitizeName key) value) coercedUpdates)
        (GoVar "clone") resultType) []
    in
      { expr: if resultType == obj.exprType then GoRecordUpdateNative resultType obj.expr coercedUpdates else changedLayout
      , exprType: resultType
      }
  _ ->
    let
      boxedExprs = map (\p -> Tuple p.key (boxGoExpr codegenStateRef modNameStr p.expr p.goType)) props
    in
      { expr: GoRecordUpdateDict (boxGoExpr codegenStateRef modNameStr obj.expr obj.exprType) boxedExprs, exprType: TypeValue }
