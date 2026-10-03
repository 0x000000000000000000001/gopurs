module Gopurs.GoConversions
  ( module NativeAdts
  , coerceGoExpr
  , boxGoExpr
  , unboxGoExpr
  , generateReboxFunctions
  ) where

import Prelude

import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Ref (Ref)
import Gopurs.CodegenState (CodegenMetadata, CodegenState)
import Gopurs.GoAst (rawGo, GoExpr(..), GoDecl, GoType(..), StructPointer, goTypeToStr, structPointer, recordFieldName)
import Gopurs.GoConversions.NativeAdts (UnboxedADT, unboxableADTs, getUnboxedADT, adtPayloadTypes) as NativeAdts
import Gopurs.GoConversions.NativeAdts (slotType, zeroGoExpr)
import Gopurs.GoConversions.Rebox as Rebox
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.FfiSupport (hashString)

-- Select a conversion before crossing Value. Record projection, native ADT
-- slots and same-constructor pointers have their own representation contracts.
coerceGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoType -> GoExpr
coerceGoExpr _ _ expr from to | from == to = expr
coerceGoExpr stateRef moduleName expr (TypeRecord sourceFields) (TypeRecord targetFields)
  | canProjectRecord sourceFields targetFields =
      projectRecord (coerceGoExpr stateRef moduleName) expr sourceFields targetFields
coerceGoExpr stateRef moduleName ctor@(GoConstructor _ structName typeArgs args) _ (TypeStructValue adtName fields) =
  let
    parts = String.split (Pattern "_") structName
    ctorName = fromMaybe "" (Array.last parts)
  in
    case Map.lookup adtName NativeAdts.unboxableADTs of
      Just adt ->
        let
          provided = Array.mapWithIndex (\i arg -> Tuple (adt.fieldIndex ctorName i) arg) args
          slotExpr slot = case Array.find (\(Tuple mbSlot _) -> mbSlot == Just slot) provided of
            Just (Tuple _ arg) ->
              coerceGoExpr stateRef moduleName arg (slotType typeArgs slot) (slotType fields slot)
            Nothing -> zeroGoExpr (slotType fields slot)
        in
          GoStructValue adtName fields (Array.mapWithIndex (\i _ -> slotExpr i) fields)
      Nothing -> ctor
coerceGoExpr stateRef moduleName expr source@(TypeStructValue _ _) target@(TypeStructValue _ _) =
  unboxGoExpr stateRef moduleName (boxGoExpr stateRef moduleName expr source) TypeValue target
-- Pointer identity ignores the PureScript fullName when the instantiated Go
-- layout and runtime tag are already identical.
coerceGoExpr _ _ expr (TypeStructPointer source) (TypeStructPointer target)
  | source.baseStructName == target.baseStructName && source.fullPath == target.fullPath && source.typeArgs == target.typeArgs = expr
coerceGoExpr stateRef moduleName expr (TypeStructPointer source) (TypeStructPointer target)
  | source.baseStructName == target.baseStructName = Rebox.request stateRef moduleName expr source target
coerceGoExpr stateRef moduleName expr source@(TypeStructPointer _) target@(TypeStructPointer _) =
  unboxGoExpr stateRef moduleName (boxGoExpr stateRef moduleName expr source) TypeValue target

-- Boxed generic pointers carry Value payloads. Rebox fields before boxing, or
-- read that erased layout first and Rebox afterwards when recovering a pointer.
coerceGoExpr stateRef moduleName expr source@(TypeStructPointer pointer) TypeValue | hasTypedArguments pointer =
  let target = boxedPointerType pointer
  in boxGoExpr stateRef moduleName (coerceGoExpr stateRef moduleName expr source target) target
coerceGoExpr stateRef moduleName expr TypeValue target@(TypeStructPointer pointer) | hasTypedArguments pointer =
  let source = boxedPointerType pointer
  in coerceGoExpr stateRef moduleName (unboxGoExpr stateRef moduleName expr TypeValue source) source target
coerceGoExpr stateRef moduleName expr source@(TypeStructValue sourceAdt _) target@(TypeStructPointer { fullName: targetAdt }) | sourceAdt == targetAdt =
  -- Native sums also box to Value payloads. Use the coercion above, rather
  -- than interpreting those fields directly as the target's typed arguments.
  coerceGoExpr stateRef moduleName (boxGoExpr stateRef moduleName expr source) TypeValue target
coerceGoExpr stateRef moduleName expr from TypeValue = boxGoExpr stateRef moduleName expr from
coerceGoExpr stateRef moduleName expr TypeValue to = unboxGoExpr stateRef moduleName expr TypeValue to
coerceGoExpr stateRef moduleName expr from to =
  unboxGoExpr stateRef moduleName (boxGoExpr stateRef moduleName expr from) TypeValue to

hasTypedArguments :: StructPointer -> Boolean
hasTypedArguments pointer = Array.any (_ /= TypeValue) pointer.typeArgs

boxedPointerType :: StructPointer -> GoType
boxedPointerType pointer = structPointer pointer (map (const TypeValue) pointer.typeArgs)

canProjectRecord :: Array (Tuple String GoType) -> Array (Tuple String GoType) -> Boolean
canProjectRecord sourceFields targetFields =
  let sourceTypes = Map.fromFoldable sourceFields
  in Array.all (\(Tuple key _) -> Map.member key sourceTypes) targetFields

-- Evaluate the whole source once, including extra fields the worker drops.
projectRecord :: (GoExpr -> GoType -> GoType -> GoExpr) -> GoExpr -> Array (Tuple String GoType) -> Array (Tuple String GoType) -> GoExpr
projectRecord coerceField expr sourceFields targetFields =
  let
    sourceTypes = Map.fromFoldable sourceFields
    target = TypeRecord targetFields
    fields = map
      (\(Tuple key targetType) -> Tuple key
        (coerceField (GoStructAccess (GoVar "record") (recordFieldName key))
          (fromMaybe TypeValue (Map.lookup key sourceTypes)) targetType))
      targetFields
  in
    GoCall (GoFuncLit [ Tuple "record" (TypeRecord sourceFields) ] [] (GoRecordDict target fields) target) [ expr ]

-- Normalize generic pointer payloads before emitting the box itself. Keeping
-- this step outside boxNativeExpr prevents repeated Rebox requests.
boxGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoExpr
boxGoExpr stateRef moduleName expr source@(TypeStructPointer pointer) | hasTypedArguments pointer =
  coerceGoExpr stateRef moduleName expr source TypeValue
boxGoExpr stateRef moduleName expr source = boxNativeExpr stateRef moduleName expr source

boxNativeExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoExpr
boxNativeExpr _ _ expr TypeValue = expr
boxNativeExpr _ _ expr TypeInt64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ expr ]
boxNativeExpr _ _ expr TypeFloat64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Float") [ expr ]
boxNativeExpr _ _ expr TypeString = GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ expr ]
boxNativeExpr _ _ expr TypeBool = GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ expr ]
boxNativeExpr _ _ expr (TypeStructPointer { baseStructName }) = GoBoxStructPointer (hashString baseStructName) expr
boxNativeExpr stateRef moduleName expr (TypeRecord fields) = boxRecord (boxGoExpr stateRef moduleName) expr fields
boxNativeExpr _ _ expr (TypeInterface _) = expr
boxNativeExpr _ _ expr (TypeNativeArray TypeValue) = GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ expr ]
boxNativeExpr _ _ expr (TypeNativeArray TypeInt64) = GoBoxIntArray expr
boxNativeExpr stateRef moduleName expr (TypeNativeArray inner) = boxArray (boxGoExpr stateRef moduleName) expr inner
boxNativeExpr _ _ expr TypeUint32 = rawGo ("gopurs_runtime.Value{Type: 9, IntVal: int64(" <> printGoExpr expr <> "), UnsafePtr: nil}")
boxNativeExpr _ _ expr (TypeGenericParam _) = expr
boxNativeExpr _ _ expr (TypeFunc _ _) = expr
boxNativeExpr stateRef moduleName expr (TypeStructValue adtName fields) =
  case Map.lookup adtName NativeAdts.unboxableADTs of
    Just adt -> adt.boxExpr (boxGoExpr stateRef moduleName) fields expr
    Nothing -> rawGo ("func() gopurs_runtime.Value {\n\t\t\t\t_ = " <> printGoExpr expr <> "\n\t\t\t\tpanic(\"boxTypeStructValue not implemented yet for " <> adtName <> "\")\n\t\t\t}()")

-- Normalize the source to Value once. The destination reader below never sees
-- a native source; typed pointer payload adaptation belongs to coerceGoExpr.
unboxGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoType -> GoExpr
unboxGoExpr stateRef moduleName expr currentType desiredType
  | currentType == desiredType = expr
  | currentType /= TypeValue =
      unboxValue stateRef moduleName (boxGoExpr stateRef moduleName expr currentType) desiredType
  | otherwise = unboxValue stateRef moduleName expr desiredType

unboxValue :: Ref CodegenState -> String -> GoExpr -> GoType -> GoExpr
unboxValue stateRef moduleName expr desiredType = case desiredType of
  TypeValue -> expr
  TypeRecord fields -> unboxRecord (coerceGoExpr stateRef moduleName) expr fields
  TypeInt64 -> GoSelector expr "IntVal"
  TypeFloat64 -> GoCall (GoSelector expr "FloatVal") []
  TypeString -> GoCall (GoSelector expr "StrVal") []
  TypeBool -> GoBinOp "!=" (GoSelector expr "IntVal") (GoInt 0)
  TypeUint32 -> rawGo ("uint32(" <> printGoExpr (GoSelector expr "IntVal") <> ")")
  TypeStructPointer { fullPath } -> GoCall (rawGo ("gopurs_runtime.CoerceToStruct[" <> fullPath <> "]")) [ expr ]
  TypeInterface _ -> expr
  -- Value arrays borrow their immutable slice; boxing []Value also shares it.
  TypeNativeArray TypeValue -> rawGo ("(*(*[]gopurs_runtime.Value)((" <> printGoExpr expr <> ").UnsafePtr))")
  TypeNativeArray TypeInt64 -> GoUnboxIntArray expr
  TypeNativeArray inner -> unboxArray (coerceGoExpr stateRef moduleName) expr inner
  TypeGenericParam _ -> expr
  TypeFunc _ _ -> expr
  TypeStructValue adtName fields ->
    case Map.lookup adtName NativeAdts.unboxableADTs of
      Just adt -> adt.unboxExpr (\value fieldType -> unboxGoExpr stateRef moduleName value TypeValue fieldType) fields expr
      Nothing -> rawGo ("func() " <> goTypeToStr desiredType <> " {\n\t\t\t\t_ = " <> printGoExpr expr <> "\n\t\t\t\tpanic(\"unboxTypeStructValue not implemented yet for " <> adtName <> "\")\n\t\t\t}()")

-- Records and arrays capture their operand once, then adapt fields/elements in
-- source order. These callbacks can register Rebox dependencies recursively.
boxRecord :: (GoExpr -> GoType -> GoExpr) -> GoExpr -> Array (Tuple String GoType) -> GoExpr
boxRecord boxField expr fields =
  let
    keys = map (\(Tuple key _) -> key) fields
    keysStr = String.joinWith ", " (map (printGoExpr <<< GoString) keys)
    valsStr = String.joinWith ", " (map (\(Tuple key fieldType) -> printGoExpr (boxField (GoStructAccess (GoVar "orig") (recordFieldName key)) fieldType)) fields)
    boxedRecord = case Array.length fields of
      0 -> "gopurs_runtime.RecordDict0()"
      size | size <= 5 -> "gopurs_runtime.RecordDict" <> show size <> "(" <> keysStr <> ", " <> valsStr <> ")"
      _ -> "gopurs_runtime.RecordDict([]string{" <> keysStr <> "}, []gopurs_runtime.Value{" <> valsStr <> "})"
  in
    rawGo ("func() gopurs_runtime.Value {\n\t\t\t\torig := " <> printGoExpr expr <> "\n\t\t\t\t_ = orig\n\t\t\t\treturn " <> boxedRecord <> "\n\t\t\t\t}()")

unboxRecord :: (GoExpr -> GoType -> GoType -> GoExpr) -> GoExpr -> Array (Tuple String GoType) -> GoExpr
unboxRecord coerceField expr fields =
  let
    recordType = goTypeToStr (TypeRecord fields)
    assignments = String.joinWith "\n" (map (\(Tuple key fieldType) -> "\t\t\t\t\tclone." <> recordFieldName key <> " = " <> printGoExpr (coerceField (GoCall (GoSelector (GoVar "gopurs_runtime") "RecordGet") [ GoVar "orig", GoString key ]) TypeValue fieldType)) fields)
  in
    rawGo ("func() " <> recordType <> " {\n\t\t\t\t\torig := " <> printGoExpr expr <> "\n\t\t\t\t\t_ = orig\n\t\t\t\t\tclone := " <> recordType <> "{}\n" <> assignments <> "\n\t\t\t\t\treturn clone\n\t\t\t\t}()")

boxArray :: (GoExpr -> GoType -> GoExpr) -> GoExpr -> GoType -> GoExpr
boxArray boxElement expr inner =
  rawGo ("func() gopurs_runtime.Value {\n\t\t\t\t\tarr := " <> printGoExpr expr <> "\n\t\t\t\t\tboxed := make([]gopurs_runtime.Value, len(arr))\n\t\t\t\t\tfor i, v := range arr { boxed[i] = " <> printGoExpr (boxElement (GoVar "v") inner) <> " }\n\t\t\t\t\treturn gopurs_runtime.Array(boxed)\n\t\t\t\t}()")

unboxArray :: (GoExpr -> GoType -> GoType -> GoExpr) -> GoExpr -> GoType -> GoExpr
unboxArray coerceElement expr inner =
  let arrayType = goTypeToStr (TypeNativeArray inner)
  in rawGo ("func() " <> arrayType <> " {\n\t\t\t\t\tarr := *(*[]gopurs_runtime.Value)(" <> printGoExpr expr <> ".UnsafePtr)\n\t\t\t\t\tunboxed := make(" <> arrayType <> ", len(arr))\n\t\t\t\t\tfor i, v := range arr { unboxed[i] = " <> printGoExpr (coerceElement (GoVar "v") TypeValue inner) <> " }\n\t\t\t\t\treturn unboxed\n\t\t\t\t}()")

-- Field rendering re-enters this converter with the same module-local state.
-- Rebox owns request tracking, naming and the transitive emission loop.
generateReboxFunctions :: CodegenMetadata -> Ref CodegenState -> String -> Effect (Array GoDecl)
generateReboxFunctions metadata stateRef moduleName =
  Rebox.generate metadata stateRef moduleName (coerceGoExpr stateRef moduleName)
