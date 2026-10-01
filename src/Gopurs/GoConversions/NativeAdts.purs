module Gopurs.GoConversions.NativeAdts
  ( UnboxedADT
  , unboxableADTs
  , getUnboxedADT
  , adtPayloadTypes
  , slotType
  , zeroGoExpr
  ) where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..), goTypeToStr)
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))
import PureScript.Backend.Optimizer.FfiSupport (hashString)

-- Native value layouts supported by the expression generator. These rules
-- also define how their payloads cross the runtime Value boundary. A payload
-- slot whose instantiated type is a closed record stays native across direct
-- worker calls; every other payload remains a Value, so existing traversals,
-- array fusion and dynamic boundaries keep their representation.
type UnboxedADT =
  { signature :: Array GoType -> Array GoType
  , mapConstructor :: Array GoType -> String -> Array GoExpr -> Array GoExpr
  , isConstructor :: String -> GoExpr -> GoExpr
  , fieldIndex :: String -> Int -> Maybe Int
  , boxExpr :: (GoExpr -> GoType -> GoExpr) -> Array GoType -> GoExpr -> GoExpr
  , unboxExpr :: (GoExpr -> GoType -> GoExpr) -> Array GoType -> GoExpr -> GoExpr
  }

-- Only closed records avoid the box here. Arrays, ADT pointers and scalars
-- keep their existing Value payload so no traversal or coercion changes.
nativeSlot :: GoType -> GoType
nativeSlot t = case t of
  TypeRecord _ -> t
  _ -> TypeValue

slotType :: Array GoType -> Int -> GoType
slotType types index = fromMaybe TypeValue (Array.index types index)

zeroGoExpr :: GoType -> GoExpr
zeroGoExpr t = case t of
  TypeValue -> rawGo "gopurs_runtime.Value{}"
  TypeInt64 -> rawGo "int64(0)"
  TypeFloat64 -> rawGo "float64(0)"
  TypeString -> rawGo "\"\""
  TypeBool -> rawGo "false"
  TypeUint32 -> rawGo "uint32(0)"
  TypeStructPointer _ -> rawGo ("(" <> goTypeToStr t <> ")(nil)")
  TypeRecord _ -> rawGo (goTypeToStr t <> "{}")
  TypeStructValue _ _ -> rawGo (goTypeToStr t <> "{}")
  TypeNativeArray _ -> rawGo "nil"
  TypeInterface _ -> rawGo "nil"
  TypeFunc _ _ -> rawGo "nil"
  TypeGenericParam _ -> rawGo "gopurs_runtime.Value{}"

-- Payload argument types of a known Maybe/Either/Tuple instantiation.
adtPayloadTypes :: (ExprType -> GoType) -> ExprType -> Array GoType
adtPayloadTypes toGoType = case _ of
  ADT _ _ args -> map toGoType args
  TypeApp fn args -> adtPayloadTypes toGoType fn <> map toGoType args
  _ -> []

-- Slot callbacks belong to GoConversions. Keeping them explicit lets native
-- layouts recurse through records and pointers without importing the converter.
unboxableADTs :: Map String UnboxedADT
unboxableADTs = Map.fromFoldable
  [ Tuple "Data.Maybe.Maybe"
      { signature: \types -> [ nativeSlot (slotType types 0), TypeBool ]
      , mapConstructor: \types ctorName args ->
          case ctorName of
            "Just" -> [ fromMaybe (zeroGoExpr (nativeSlot (slotType types 0))) (Array.index args 0), rawGo "true" ]
            _ -> [ zeroGoExpr (nativeSlot (slotType types 0)), rawGo "false" ]
      , isConstructor: \ctor expr -> case ctor of
          "Data_Data_Maybe_Just" -> GoSelector expr "V1"
          "Data_Data_Maybe_Nothing" -> rawGo ("(!" <> printGoExpr (GoSelector expr "V1") <> ")")
          _ -> rawGo "false"
      , fieldIndex: \ctor index -> if ctor == "Data_Data_Maybe_Just" && index == 0 then Just 0 else Nothing
      , boxExpr: \boxSlot types expr ->
          let
            slot = nativeSlot (slotType types 0)
            payload = boxSlot (GoSelector (GoVar "_v") "V0") slot
          in
            rawGo ("func() gopurs_runtime.Value {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.V1 {\n\t\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Maybe_Just" <> ", UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: " <> printGoExpr payload <> "})}\n\t\t\t\t}\n\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Maybe_Just" <> "}\n\t\t\t}()")
      , unboxExpr: \unboxSlot types expr ->
          let
            slot = nativeSlot (slotType types 0)
            signature = [ slot, TypeBool ]
            structType = goTypeToStr (TypeStructValue "Data.Maybe.Maybe" signature)
            payload = unboxSlot (rawGo "(*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0") slot
          in
            rawGo ("func() " <> structType <> " {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.Type == 9 && _v.IntVal == " <> hashString "Data_Data_Maybe_Just" <> " && _v.UnsafePtr != nil {\n\t\t\t\t\treturn " <> structType <> "{V0: " <> printGoExpr payload <> ", V1: true}\n\t\t\t\t}\n\t\t\t\treturn " <> structType <> "{V0: " <> printGoExpr (zeroGoExpr slot) <> ", V1: false}\n\t\t\t}()")
      }
  , Tuple "Data.Tuple.Tuple"
      { signature: \types -> [ nativeSlot (slotType types 0), nativeSlot (slotType types 1) ]
      , mapConstructor: \_ _ args -> args
      , isConstructor: \ctor _ -> rawGo (if ctor == "Data_Data_Tuple_Tuple" then "true" else "false")
      , fieldIndex: \ctor index -> if ctor == "Data_Data_Tuple_Tuple" && index >= 0 && index < 2 then Just index else Nothing
      , boxExpr: \boxSlot types expr ->
          let
            signature = [ nativeSlot (slotType types 0), nativeSlot (slotType types 1) ]
            v0 = boxSlot (GoSelector (GoVar "_v") "V0") (slotType signature 0)
            v1 = boxSlot (GoSelector (GoVar "_v") "V1") (slotType signature 1)
          in
            rawGo ("func() gopurs_runtime.Value {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Tuple_Tuple" <> ", UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: " <> printGoExpr v0 <> ", V1: " <> printGoExpr v1 <> "})}\n\t\t\t}()")
      , unboxExpr: \unboxSlot types expr ->
          let
            signature = [ nativeSlot (slotType types 0), nativeSlot (slotType types 1) ]
            structType = goTypeToStr (TypeStructValue "Data.Tuple.Tuple" signature)
            v0 = unboxSlot (rawGo "(*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0") (slotType signature 0)
            v1 = unboxSlot (rawGo "(*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1") (slotType signature 1)
          in
            rawGo ("func() " <> structType <> " {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\treturn " <> structType <> "{V0: " <> printGoExpr v0 <> ", V1: " <> printGoExpr v1 <> "}\n\t\t\t}()")
      }
  , Tuple "Data.Either.Either"
      { signature: \types -> [ nativeSlot (slotType types 0), nativeSlot (slotType types 1), TypeBool ]
      , mapConstructor: \types ctorName args ->
          let
            left = nativeSlot (slotType types 0)
            right = nativeSlot (slotType types 1)
          in
            case ctorName of
              "Left" -> [ fromMaybe (zeroGoExpr left) (Array.index args 0), zeroGoExpr right, rawGo "false" ]
              "Right" -> [ zeroGoExpr left, fromMaybe (zeroGoExpr right) (Array.index args 0), rawGo "true" ]
              _ -> [ zeroGoExpr left, zeroGoExpr right, rawGo "false" ]
      , isConstructor: \ctor expr -> case ctor of
          "Data_Data_Either_Right" -> GoSelector expr "V2"
          "Data_Data_Either_Left" -> rawGo ("(!" <> printGoExpr (GoSelector expr "V2") <> ")")
          _ -> rawGo "false"
      , fieldIndex: \ctor index -> if index /= 0 then Nothing else case ctor of
          "Data_Data_Either_Left" -> Just 0
          "Data_Data_Either_Right" -> Just 1
          _ -> Nothing
      , boxExpr: \boxSlot types expr ->
          let
            left = nativeSlot (slotType types 0)
            right = nativeSlot (slotType types 1)
            v0 = boxSlot (GoSelector (GoVar "_v") "V0") left
            v1 = boxSlot (GoSelector (GoVar "_v") "V1") right
          in
            rawGo ("func() gopurs_runtime.Value {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.V2 {\n\t\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Either_Right" <> ", UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: " <> printGoExpr v1 <> "})}\n\t\t\t\t}\n\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Either_Left" <> ", UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: " <> printGoExpr v0 <> "})}\n\t\t\t}()")
      , unboxExpr: \unboxSlot types expr ->
          let
            left = nativeSlot (slotType types 0)
            right = nativeSlot (slotType types 1)
            structType = goTypeToStr (TypeStructValue "Data.Either.Either" [ left, right, TypeBool ])
            leftValue = unboxSlot (rawGo "(*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0") left
            rightValue = unboxSlot (rawGo "(*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0") right
          in
            rawGo ("func() " <> structType <> " {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.Type == 9 && _v.IntVal == " <> hashString "Data_Data_Either_Right" <> " && _v.UnsafePtr != nil {\n\t\t\t\t\treturn " <> structType <> "{V0: " <> printGoExpr (zeroGoExpr left) <> ", V1: " <> printGoExpr rightValue <> ", V2: true}\n\t\t\t\t}\n\t\t\t\treturn " <> structType <> "{V0: " <> printGoExpr leftValue <> ", V1: " <> printGoExpr (zeroGoExpr right) <> ", V2: false}\n\t\t\t}()")
      }
  ]

getUnboxedADT :: ExprType -> Maybe (Tuple String UnboxedADT)
getUnboxedADT (ADT fullName _ _) = do
  adt <- Map.lookup fullName unboxableADTs
  pure (Tuple fullName adt)
getUnboxedADT (TypeApp f _) = getUnboxedADT f
getUnboxedADT _ = Nothing
