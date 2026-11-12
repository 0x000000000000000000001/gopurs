module Gopurs.FfiBridge.Structured
  ( needsConversion
  , problem
  , fromValue
  , toValue
  ) where

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..), isJust)
import Data.String as String
import Data.Tuple (Tuple(..))
import Gopurs.FfiBridge.TypeSupport (printTypeNode)
import Gopurs.FfiTypes (StructField, TypeNode(..))
import Gopurs.GoAst (capitalize)
import Gopurs.GoTypes (isClosedRowTail, visibleRecordFields)
import Gopurs.Printer (quoteGoString)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

type Convert = TypeNode -> ExprType -> String -> String

type MatchedField = { label :: String, field :: StructField, tast :: ExprType }

layout :: TypeNode -> Maybe (Array StructField)
layout (TStruct _ fields) = Just fields
layout (TIdentified _ _ (Just underlying)) = layout underlying
layout _ = Nothing

shape :: TypeNode -> TypeNode
shape (TIdentified _ _ (Just underlying)) = shape underlying
shape other = other

containsStruct :: TypeNode -> Boolean
containsStruct t | isJust (layout t) = true
containsStruct (TIdentified _ _ (Just t)) = containsStruct t
containsStruct (TArray t) = containsStruct t
containsStruct (TPointer t) = containsStruct t
containsStruct _ = false

containsRecord :: ExprType -> Boolean
containsRecord (Record _) = true
containsRecord (Array t) = containsRecord t
containsRecord (ADT "Data.Nullable.Nullable" _ [t]) = containsRecord t
containsRecord _ = false

-- Opaque handles stay native even when their local Go type happens to be a
-- struct. Only a structural PureScript contract requests field conversion.
needsConversion :: TypeNode -> ExprType -> Boolean
needsConversion go tast = containsStruct go && containsRecord tast

problem :: TypeNode -> ExprType -> Maybe String
problem go tast
  | isLegacy go = Nothing
  | containsStruct go || containsRecord tast = validate go tast
  | otherwise = Nothing

isLegacy :: TypeNode -> Boolean
isLegacy (TNamed name) = Array.elem name [ "gopurs_runtime.Value", "any", "interface{}" ]
isLegacy (TMap _ _) = true
isLegacy (TIdentified _ _ (Just underlying)) = isLegacy underlying
isLegacy _ = false

validate :: TypeNode -> ExprType -> Maybe String
validate go _ | isLegacy go = Nothing
validate go (Record (Row fields tail))
  | not (isClosedRowTail tail) = Just "Open PureScript record rows cannot be converted to a Go struct; use a closed record."
  | otherwise = case layout go of
      Just members -> validateFields members (visibleRecordFields fields)
      Nothing -> case go of
        TNamed "gopurs_runtime.Value" -> Nothing
        TNamed "any" -> Nothing
        TNamed "interface{}" -> Nothing
        TMap _ _ -> Nothing
        TPointer _ -> Just "A non-nullable PureScript record requires a Go struct value; use Nullable Record for a *Struct."
        _ -> Just "A PureScript record requires a locally declared Go struct whose fields can be checked."
validate (TArray item) (Array tast) = validate item tast
validate (TPointer item) (ADT "Data.Nullable.Nullable" _ [tast]) = validate item tast
validate (TIdentified _ _ (Just underlying)) tast = validate underlying tast
validate (TNamed "string") String = Nothing
validate (TNamed "bool") Boolean = Nothing
validate (TNamed "int") Int = Nothing
validate (TNamed "int32") Int = Nothing
validate (TNamed "int64") Int = Nothing
validate (TNamed "float64") Number = Nothing
validate (TNamed "gopurs_runtime.Value") _ = Nothing
validate (TNamed "any") Any = Nothing
validate (TNamed "interface{}") Any = Nothing
validate (TNamed _) (ADT _ _ _) = Nothing
validate (TPointer _) (ADT _ _ _) = Nothing
validate (TStruct _ _) (ADT _ _ _) = Nothing
validate _ _ = Just "The Go field type is incompatible with its PureScript field; supported structural fields are scalars, arrays, closed records, Nullable record pointers and opaque handles."

validateFields :: Array StructField -> Array (Tuple String ExprType) -> Maybe String
validateFields fields labels =
  let
    names = map (\(Tuple label _) -> capitalize label) labels
    inaccessible = Array.find (\field -> field.embedded || not field.exported) fields
    missing = Array.find (\name -> not (Array.any (\field -> field.name == name) fields)) names
    extra = Array.find (\field -> not (Array.elem field.name names)) fields
  in
    case inaccessible, missing, extra of
      Just field, _, _ -> Just ("Go struct field " <> field.name <> " must be exported and non-embedded for record conversion.")
      _, _, _ | Array.length (Array.nub names) /= Array.length names -> Just "PureScript record labels collide after Go field capitalization."
      _, Just name, _ -> Just ("Missing Go struct field " <> name <> "; fields use the capitalized PureScript label.")
      _, _, Just field -> Just ("Unexpected Go struct field " <> field.name <> "; the closed record must describe every field.")
      _, _, _ -> Array.head (Array.mapMaybe check (matchFields fields labels))
  where
  check item = map (\reason -> "Field " <> item.label <> ": " <> reason) (validateNativeField item.field.type item.tast)

-- Compatibility for legacy top-level any/maps does not authorize boxed values
-- inside a newly generated native struct. Its dynamic fields need Foreign.
validateNativeField :: TypeNode -> ExprType -> Maybe String
validateNativeField (TIdentified _ _ (Just underlying)) tast = validateNativeField underlying tast
validateNativeField (TArray item) (Array tast) = validateNativeField item tast
validateNativeField (TPointer item) (ADT "Data.Nullable.Nullable" _ [tast]) = validateNativeField item tast
validateNativeField (TNamed name) tast | name == "any" || name == "interface{}" = case tast of
  ADT "Foreign.Foreign" _ _ -> Nothing
  _ -> Just "A Go any/interface{} struct field requires Foreign.Foreign; use a concrete Go type for a typed PureScript field."
validateNativeField (TMap _ _) _ = Just "Map fields in native structs require an explicit conversion; use a closed Go struct, or any with Foreign.Foreign for dynamic data."
validateNativeField go tast = validate go tast

dynamicForeign :: TypeNode -> ExprType -> Boolean
dynamicForeign go (ADT "Foreign.Foreign" _ _) = case shape go of
  TNamed name -> name == "any" || name == "interface{}"
  _ -> false
dynamicForeign _ _ = false

matchFields :: Array StructField -> Array (Tuple String ExprType) -> Array MatchedField
matchFields fields = Array.mapMaybe \(Tuple label tast) -> do
  field <- Array.find (\field -> field.name == capitalize label) fields
  pure { label, field, tast }

-- The caller validates before emission and supplies its normal conversion for
-- scalar/callback/opaque leaves. Each source expression is evaluated once.
fromValue :: Convert -> TypeNode -> ExprType -> String -> String
fromValue _ (TNamed "gopurs_runtime.Value") _ value = value
fromValue convert go tast value | dynamicForeign go tast = convert (shape go) tast value
fromValue convert go tast value = case shape go, tast of
  TArray item, Array itemTast ->
    "func() " <> printTypeNode go <> " { _source := " <> value
      <> "; _out := make(" <> printTypeNode go <> ", gopurs_runtime.ArrayLength(_source)); for _i := range _out { _out[_i] = "
      <> fromValue convert item itemTast "gopurs_runtime.ArrayAccess(_source, _i)" <> " }; return _out }()"
  TPointer item, ADT "Data.Nullable.Nullable" _ [itemTast] ->
    "func() " <> printTypeNode go <> " { _nullable := " <> value
      <> "; if _nullable.Type == gopurs_runtime.TypeAny { if _nullable.PtrVal() == nil { return nil }; if _pointer, _ok := _nullable.PtrVal().(" <> printTypeNode go <> "); _ok { return _pointer } }; _record := "
      <> fromValue convert item itemTast "_nullable" <> "; return &_record }()"
  _, Record (Row labels _) -> case layout go of
    Just fields ->
      let
        entries = map (\item -> item.field.name <> ": " <> fromValue convert item.field.type item.tast ("gopurs_runtime.RecordGet(_source, " <> quoteGoString item.label <> ")")) (matchFields fields (visibleRecordFields labels))
      in
        "func() " <> printTypeNode go <> " { _source := " <> value <> "; _ = _source; return " <> printTypeNode go <> "{" <> String.joinWith ", " entries <> "} }()"
    Nothing -> convert go tast value
  _, _ -> convert go tast value

toValue :: Convert -> TypeNode -> ExprType -> String -> String
toValue convert go tast value | dynamicForeign go tast = convert (shape go) tast value
toValue convert go tast value = case shape go, tast of
  TArray item, Array itemTast ->
    "func() gopurs_runtime.Value { _source := " <> value
      <> "; _out := make([]gopurs_runtime.Value, len(_source)); for _i, _item := range _source { _out[_i] = "
      <> toValue convert item itemTast "_item" <> " }; return gopurs_runtime.Array(_out) }()"
  TPointer item, ADT "Data.Nullable.Nullable" _ [itemTast] ->
    "func() gopurs_runtime.Value { _nullable := " <> value <> "; if _nullable == nil { return gopurs_runtime.Any(nil) }; return "
      <> toValue convert item itemTast "(*_nullable)" <> " }()"
  _, Record (Row labels _) -> case layout go of
    Just fields ->
      let
        matched = matchFields fields (visibleRecordFields labels)
        keys = map (quoteGoString <<< _.label) matched
        values = map (\item -> toValue convert item.field.type item.tast ("_source." <> item.field.name)) matched
        record = if Array.length keys <= 5 then
            "gopurs_runtime.RecordDict" <> show (Array.length keys) <> "(" <> String.joinWith ", " (keys <> values) <> ")"
          else "gopurs_runtime.RecordDict([]string{" <> String.joinWith ", " keys <> "}, []gopurs_runtime.Value{" <> String.joinWith ", " values <> "})"
      in
        "func() gopurs_runtime.Value { _source := " <> value <> "; _ = _source; return " <> record <> " }()"
    Nothing -> convert go tast value
  _, _ -> convert go tast value
