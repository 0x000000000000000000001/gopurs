module Gopurs.GoConversions
  ( UnboxedADT
  , unboxableADTs
  , getUnboxedADT
  , coerceGoExpr
  , boxGoExpr
  , unboxGoExpr
  , generateReboxFunctions
  ) where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..), fst)
import Effect (Effect)
import Effect.Console as Console
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CodegenState (CodegenMetadata, CodegenState)
import Gopurs.GoAst (rawGo, GoExpr(..), GoDecl(..), GoType(..), goTypeToStr, structPointer, sanitizeName)
import Gopurs.GoTypes as GoTypes
import Gopurs.Printer (printGoExpr)
import Gopurs.ReboxMetadata (ReboxFields)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))
import PureScript.Backend.Optimizer.FfiSupport (hashString)

-- Native value layouts supported by the expression generator. These rules
-- also define how their payloads cross the runtime Value boundary.
type UnboxedADT =
  { signature :: Array GoType
  , mapConstructor :: String -> Array GoExpr -> Array GoExpr
  , isConstructor :: String -> GoExpr -> GoExpr
  , boxExpr :: GoExpr -> GoExpr
  , unboxExpr :: GoExpr -> GoExpr
  }

unboxableADTs :: Map String UnboxedADT
unboxableADTs = Map.fromFoldable
  [ Tuple "Data.Maybe.Maybe"
      { signature: [ TypeValue, TypeBool ]
      , mapConstructor: \ctorName args ->
          case ctorName of
            "Just" -> [ fromMaybe (rawGo "gopurs_runtime.Value{}") (Array.index args 0), rawGo "true" ]
            "Nothing" -> [ rawGo "gopurs_runtime.Value{}", rawGo "false" ]
            _ -> args
      , isConstructor: \ctor expr -> case ctor of
          "Data_Data_Maybe_Just" -> GoSelector expr "V1"
          "Data_Data_Maybe_Nothing" -> rawGo ("(!" <> printGoExpr (GoSelector expr "V1") <> ")")
          _ -> rawGo "false"
      , boxExpr: \expr ->
          -- Maybe uses the Just constructor id for both cases, with a nil pointer for Nothing.
          rawGo ("func() gopurs_runtime.Value {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.V1 {\n\t\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Maybe_Just" <> ", UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}\n\t\t\t\t}\n\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Maybe_Just" <> "}\n\t\t\t}()")
      , unboxExpr: \expr ->
          rawGo ("func() struct{V0 gopurs_runtime.Value; V1 bool} {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.Type == 9 && _v.IntVal == " <> hashString "Data_Data_Maybe_Just" <> " && _v.UnsafePtr != nil {\n\t\t\t\t\treturn struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}\n\t\t\t\t}\n\t\t\t\treturn struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}\n\t\t\t}()")
      }
  , Tuple "Data.Tuple.Tuple"
      { signature: [ TypeValue, TypeValue ]
      , mapConstructor: \_ args -> args
      , isConstructor: \ctor _ -> rawGo (if ctor == "Data_Data_Tuple_Tuple" then "true" else "false")
      , boxExpr: \expr ->
          rawGo ("func() gopurs_runtime.Value {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Tuple_Tuple" <> ", UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}\n\t\t\t}()")
      , unboxExpr: \expr ->
          rawGo ("func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\t_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)\n\t\t\t\treturn struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}\n\t\t\t}()")
      }
  , Tuple "Data.Either.Either"
      { signature: [ TypeValue, TypeValue, TypeBool ]
      , mapConstructor: \ctor args -> case ctor of
          "Left" -> [ fromMaybe (rawGo "gopurs_runtime.Value{}") (Array.head args), rawGo "gopurs_runtime.Value{}", rawGo "false" ]
          "Right" -> [ rawGo "gopurs_runtime.Value{}", fromMaybe (rawGo "gopurs_runtime.Value{}") (Array.head args), rawGo "true" ]
          _ -> [ rawGo "gopurs_runtime.Value{}", rawGo "gopurs_runtime.Value{}", rawGo "false" ]
      , isConstructor: \ctor expr -> case ctor of
          "Data_Data_Either_Right" -> GoSelector expr "V2"
          "Data_Data_Either_Left" -> rawGo ("(!" <> printGoExpr (GoSelector expr "V2") <> ")")
          _ -> rawGo "false"
      , boxExpr: \expr ->
          rawGo ("func() gopurs_runtime.Value {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.V2 {\n\t\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Either_Right" <> ", UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}\n\t\t\t\t}\n\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Either_Left" <> ", UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}\n\t\t\t}()")
      , unboxExpr: \expr ->
          rawGo ("func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.Type == 9 && _v.IntVal == " <> hashString "Data_Data_Either_Right" <> " && _v.UnsafePtr != nil {\n\t\t\t\t\treturn struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}\n\t\t\t\t}\n\t\t\t\treturn struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}\n\t\t\t}()")
      }
  ]

getUnboxedADT :: ExprType -> Maybe (Tuple String UnboxedADT)
getUnboxedADT (ADT fullName _ _) = do
  adt <- Map.lookup fullName unboxableADTs
  pure (Tuple fullName adt)
getUnboxedADT (TypeApp f _) = getUnboxedADT f
getUnboxedADT _ = Nothing

-- Coercions can request Rebox helpers through the state of the current
-- translation. Keep their registration and generation in this module.
coerceGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoType -> GoExpr
coerceGoExpr _ _ expr from to | from == to = expr
-- A known record can supply a smaller native worker argument without boxing
-- every field first. Evaluate the source once, including unused extra fields.
coerceGoExpr codegenStateRef modNameStr expr source@(TypeRecord sourceFields) target@(TypeRecord targetFields)
  | canProjectRecord sourceFields targetFields =
      let sourceTypes = Map.fromFoldable sourceFields
      in
      GoCall
        (GoFuncLit [ Tuple "record" source ] []
          (GoRecordDict target (map
            (\(Tuple key targetType) -> Tuple key
              (coerceGoExpr codegenStateRef modNameStr
                (GoStructAccess (GoVar "record") (sanitizeName key))
                (fromMaybe TypeValue (Map.lookup key sourceTypes)) targetType))
            targetFields)) target)
        [ expr ]
coerceGoExpr _ _ ctor@(GoConstructor _ structName _ args) _ (TypeStructValue adtName fields) =
  let
    parts = String.split (Pattern "_") structName
    ctorName = fromMaybe "" (Array.last parts)
  in
    case Map.lookup adtName unboxableADTs of
      Just adt -> GoStructValue adtName fields (adt.mapConstructor ctorName args)
      Nothing -> ctor -- fallback
coerceGoExpr _ _ expr (TypeStructPointer { baseStructName: b1, fullPath: s1, typeArgs: a1 }) (TypeStructPointer { baseStructName: b2, fullPath: s2, typeArgs: a2 }) | b1 == b2 && s1 == s2 && a1 == a2 = expr

coerceGoExpr codegenStateRef modNameStr expr srcT@(TypeStructPointer { baseStructName: b1, fullPath: s1 }) destT@(TypeStructPointer { baseStructName: b2, fullPath: s2 }) | b1 == b2 =
  unsafePerformEffect do
    registerReboxPair codegenStateRef srcT destT
    pure $ GoCall (GoVar ("Rebox_" <> modNameStr <> "_" <> hashString s1 <> "_" <> hashString s2)) [ expr ]

coerceGoExpr codegenStateRef modNameStr expr srcT@(TypeStructPointer _) destT@(TypeStructPointer _) =
  unboxGoExpr codegenStateRef modNameStr (boxGoExpr codegenStateRef modNameStr expr srcT) TypeValue destT

coerceGoExpr codegenStateRef modNameStr expr srcT@(TypeStructPointer pointer@{ typeArgs: a }) TypeValue | Array.any (_ /= TypeValue) a =
  let
    destT = structPointer pointer (map (const TypeValue) a)
  in
    boxGoExpr codegenStateRef modNameStr (coerceGoExpr codegenStateRef modNameStr expr srcT destT) destT

coerceGoExpr codegenStateRef modNameStr expr TypeValue destT@(TypeStructPointer pointer@{ typeArgs: a }) | Array.any (_ /= TypeValue) a =
  let
    srcT = structPointer pointer (map (const TypeValue) a)
  in
    coerceGoExpr codegenStateRef modNameStr (unboxGoExpr codegenStateRef modNameStr expr TypeValue srcT) srcT destT

coerceGoExpr codegenStateRef modNameStr expr srcT@(TypeStructValue srcAdt _) destT@(TypeStructPointer { fullName: destAdt }) | srcAdt == destAdt =
  -- Boxed native ADTs have Value payloads; convert each field before using
  -- a typed pointer instead of reinterpreting the generic payload layout.
  coerceGoExpr codegenStateRef modNameStr (boxGoExpr codegenStateRef modNameStr expr srcT) TypeValue destT

coerceGoExpr codegenStateRef modNameStr expr from TypeValue = boxGoExpr codegenStateRef modNameStr expr from
coerceGoExpr codegenStateRef modNameStr expr TypeValue to = unboxGoExpr codegenStateRef modNameStr expr TypeValue to
coerceGoExpr codegenStateRef modNameStr expr from to = unboxGoExpr codegenStateRef modNameStr (boxGoExpr codegenStateRef modNameStr expr from) TypeValue to

canProjectRecord :: Array (Tuple String GoType) -> Array (Tuple String GoType) -> Boolean
canProjectRecord sourceFields targetFields =
  let sourceTypes = Map.fromFoldable sourceFields
  in Array.all (\(Tuple key _) -> Map.member key sourceTypes) targetFields

-- Before boxing a generic pointer, Rebox converts its type arguments to
-- Value. boxGoExprImpl then emits the box without repeating that conversion.
boxGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoExpr
boxGoExpr codegenStateRef modNameStr expr srcT@(TypeStructPointer { typeArgs }) =
  if Array.any (\t -> t /= TypeValue) typeArgs then
    coerceGoExpr codegenStateRef modNameStr expr srcT TypeValue
  else
    boxGoExprImpl codegenStateRef modNameStr expr srcT
boxGoExpr codegenStateRef modNameStr expr t = boxGoExprImpl codegenStateRef modNameStr expr t

-- Records capture their source once and box each field. Int arrays have a
-- dedicated helper; other typed array elements are converted recursively.
boxGoExprImpl :: Ref CodegenState -> String -> GoExpr -> GoType -> GoExpr
boxGoExprImpl _ _ expr TypeValue = expr
boxGoExprImpl _ _ expr TypeInt64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ expr ]
boxGoExprImpl _ _ expr TypeFloat64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Float") [ expr ]
boxGoExprImpl _ _ expr TypeString = GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ expr ]
boxGoExprImpl _ _ expr TypeBool = GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ expr ]
boxGoExprImpl _ _ expr (TypeStructPointer { baseStructName }) = GoBoxStructPointer (hashString baseStructName) expr
boxGoExprImpl codegenStateRef modNameStr expr (TypeRecord fields) =
  let
    keys = map (\(Tuple k _) -> k) fields
    keysStr = String.joinWith ", " (map (\k -> "\"" <> k <> "\"") keys)
    valsStr = String.joinWith ", " (map (\(Tuple k v) -> printGoExpr (boxGoExpr codegenStateRef modNameStr (GoStructAccess (GoVar "orig") (sanitizeName k)) v)) fields)
    boxedRecord =
      case Array.length fields of
        0 -> "gopurs_runtime.RecordDict0()"
        size | size <= 5 ->
          "gopurs_runtime.RecordDict" <> show size <> "(" <> keysStr <> ", " <> valsStr <> ")"
        _ ->
          "gopurs_runtime.RecordDict([]string{" <> keysStr <> "}, []gopurs_runtime.Value{" <> valsStr <> "})"
  in
    rawGo ("func() gopurs_runtime.Value {\n\t\t\t\torig := " <> printGoExpr expr <> "\n\t\t\t\t_ = orig\n\t\t\t\treturn " <> boxedRecord <> "\n\t\t\t\t}()")
boxGoExprImpl _ _ expr (TypeInterface _) = expr
boxGoExprImpl _ _ expr (TypeNativeArray TypeValue) = GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ expr ]
boxGoExprImpl _ _ expr (TypeNativeArray TypeInt64) = GoBoxIntArray expr
boxGoExprImpl codegenStateRef modNameStr expr (TypeNativeArray inner) = rawGo ("func() gopurs_runtime.Value {\n\t\t\t\t\tarr := " <> printGoExpr expr <> "\n\t\t\t\t\tboxed := make([]gopurs_runtime.Value, len(arr))\n\t\t\t\t\tfor i, v := range arr { boxed[i] = " <> printGoExpr (boxGoExpr codegenStateRef modNameStr (GoVar "v") inner) <> " }\n\t\t\t\t\treturn gopurs_runtime.Array(boxed)\n\t\t\t\t}()")
boxGoExprImpl _ _ expr TypeUint32 = rawGo ("gopurs_runtime.Value{Type: 9, IntVal: int64(" <> printGoExpr expr <> "), UnsafePtr: nil}")
boxGoExprImpl _ _ expr (TypeGenericParam _) = expr
boxGoExprImpl _ _ expr (TypeFunc _ _) = expr
boxGoExprImpl _ _ expr (TypeStructValue adtName _) =
  case Map.lookup adtName unboxableADTs of
    Just adt -> adt.boxExpr expr
    Nothing -> rawGo ("func() gopurs_runtime.Value {\n\t\t\t\t_ = " <> printGoExpr expr <> "\n\t\t\t\tpanic(\"boxTypeStructValue not implemented yet for " <> adtName <> "\")\n\t\t\t}()")

-- The destination GoType determines how to read Value: native record fields,
-- array elements, primitive payloads, or an ADT pointer.
unboxGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoType -> GoExpr
unboxGoExpr codegenStateRef modNameStr expr currentType desiredType =
  if currentType == desiredType then expr
  else if currentType /= TypeValue then
    unboxGoExpr codegenStateRef modNameStr (boxGoExpr codegenStateRef modNameStr expr currentType) TypeValue desiredType
  else case desiredType of
    TypeValue -> boxGoExpr codegenStateRef modNameStr expr currentType
    (TypeRecord fields) ->
      let
        assignments = String.joinWith "\n" (map (\(Tuple k v) -> "\t\t\t\t\tclone." <> sanitizeName k <> " = " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr (GoCall (GoSelector (GoVar "gopurs_runtime") "RecordGet") [ GoVar "orig", GoString k ]) TypeValue v)) fields)
      in
        rawGo ("func() " <> goTypeToStr desiredType <> " {\n\t\t\t\t\torig := " <> printGoExpr expr <> "\n\t\t\t\t\t_ = orig\n\t\t\t\t\tclone := " <> goTypeToStr desiredType <> "{}\n" <> assignments <> "\n\t\t\t\t\treturn clone\n\t\t\t\t}()")
    TypeInt64 -> GoSelector expr "IntVal"
    TypeFloat64 -> GoCall (GoSelector expr "FloatVal") []
    TypeString -> GoCall (GoSelector expr "StrVal") []
    TypeBool -> GoBinOp "!=" (GoSelector expr "IntVal") (GoInt 0)
    TypeUint32 -> rawGo ("uint32(" <> printGoExpr (GoSelector expr "IntVal") <> ")")
    (TypeStructPointer { fullPath }) -> GoCall (rawGo ("gopurs_runtime.CoerceToStruct[" <> fullPath <> "]")) [ expr ]
    (TypeInterface _) -> expr
    -- Array Value already has this representation. Borrow its immutable slice,
    -- just as boxing a native []Value shares the backing storage.
    (TypeNativeArray TypeValue) ->
      rawGo ("(*(*[]gopurs_runtime.Value)((" <> printGoExpr expr <> ").UnsafePtr))")
    (TypeNativeArray TypeInt64) -> GoUnboxIntArray expr
    (TypeNativeArray inner) -> case currentType of
      TypeNativeArray currentInner ->
        rawGo ("func() " <> goTypeToStr desiredType <> " {\n\t\t\t\t\tarr := " <> printGoExpr expr <> "\n\t\t\t\t\tunboxed := make(" <> goTypeToStr desiredType <> ", len(arr))\n\t\t\t\t\tfor i, v := range arr { unboxed[i] = " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr (GoVar "v") currentInner inner) <> " }\n\t\t\t\t\treturn unboxed\n\t\t\t\t}()")
      _ ->
        rawGo ("func() " <> goTypeToStr desiredType <> " {\n\t\t\t\t\tarr := *(*[]gopurs_runtime.Value)(" <> printGoExpr expr <> ".UnsafePtr)\n\t\t\t\t\tunboxed := make(" <> goTypeToStr desiredType <> ", len(arr))\n\t\t\t\t\tfor i, v := range arr { unboxed[i] = " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr (GoVar "v") TypeValue inner) <> " }\n\t\t\t\t\treturn unboxed\n\t\t\t\t}()")
    (TypeGenericParam _) -> expr
    (TypeFunc _ _) -> expr
    (TypeStructValue adtName fields) ->
      case Map.lookup adtName unboxableADTs of
        Just adt -> adt.unboxExpr expr
        Nothing -> rawGo ("func() " <> goTypeToStr (TypeStructValue adtName fields) <> " {\n\t\t\t\t_ = " <> printGoExpr expr <> "\n\t\t\t\tpanic(\"unboxTypeStructValue not implemented yet for " <> adtName <> "\")\n\t\t\t}()")

registerReboxPair :: Ref CodegenState -> GoType -> GoType -> Effect Unit
registerReboxPair codegenStateRef srcT destT = do
  state <- Ref.read codegenStateRef
  let pairs = state.reboxPairs
  if Set.member (Tuple srcT destT) pairs then pure unit
  else Ref.modify_ (\s -> s { reboxPairs = Set.insert (Tuple srcT destT) pairs }) codegenStateRef

findReboxFields :: CodegenMetadata -> String -> Maybe ReboxFields
findReboxFields metadata baseStructName =
  case Map.lookup baseStructName metadata.reboxFields of
    Just info -> Just info
    Nothing -> unsafePerformEffect do
      Console.log ("ERROR: Rebox missing! b1=" <> baseStructName <> " keysCtor: " <> String.joinWith ", " (map fst (Map.toUnfoldable metadata.ctorTypes :: Array (Tuple String _))))
      pure Nothing

renderReboxFunction :: Ref CodegenState -> CodegenMetadata -> String -> Map String GoDecl -> Tuple GoType GoType -> Maybe (Tuple String GoDecl)
renderReboxFunction codegenStateRef metadata modNameStr generatedFuncs (Tuple srcT destT) =
  case srcT, destT of
    TypeStructPointer { baseStructName: b1, fullPath: s1, typeArgs: a1 }, TypeStructPointer { baseStructName: b2, fullPath: s2, typeArgs: a2 } | b1 == b2 ->
      let
        funcName = "Rebox_" <> modNameStr <> "_" <> hashString s1 <> "_" <> hashString s2
      in
        if Map.member funcName generatedFuncs then Nothing
        else case findReboxFields metadata b1 of
          Just info ->
            let
              env1 = Map.fromFoldable (Array.zip info.vars a1)
              env2 = Map.fromFoldable (Array.zip info.vars a2)
              assignments = String.joinWith "\n" (Array.mapWithIndex
                (\i fieldExprType ->
                  let
                    genericTy = GoTypes.structFieldGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors info.vars modNameStr fieldExprType
                    t1 = GoTypes.instantiateGenericGoType env1 genericTy
                    t2 = GoTypes.instantiateGenericGoType env2 genericTy
                  in
                    "\t\tout.V" <> show i <> " = " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)
                )
                info.fields)
              funcBody = GoFunctionDecl
                { name: funcName, params: [ Tuple "in" srcT ], result: destT
                , body: rawGo ("\tif in == nil { return nil }\n\tout := &" <> s2 <> "{}\n" <> assignments <> "\n\treturn out")
                }
            in
              Just (Tuple funcName funcBody)
          Nothing -> Nothing
    _, _ -> Nothing

generateReboxFunctions :: CodegenMetadata -> Ref CodegenState -> String -> Effect (Array GoDecl)
generateReboxFunctions metadata codegenStateRef modNameStr = loop Map.empty
  where
  -- Rendering fields can register more conversions; collect until none remain.
  loop generatedFuncs = do
    state <- Ref.read codegenStateRef
    let
      reboxPairs = state.reboxPairs
      newFuncs = Map.fromFoldable $
        Array.mapMaybe (renderReboxFunction codegenStateRef metadata modNameStr generatedFuncs) (Array.fromFoldable reboxPairs)
      nextGeneratedFuncs = Map.union generatedFuncs newFuncs
    if Map.isEmpty newFuncs then
      pure $ Array.fromFoldable (Map.values nextGeneratedFuncs)
    else
      loop nextGeneratedFuncs
