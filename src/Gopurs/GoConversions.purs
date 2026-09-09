module Gopurs.GoConversions
  ( UnboxedADT
  , unboxableADTs
  , getUnboxedADT
  , coerceGoExpr
  , boxGoExpr
  , boxGoExprImpl
  , unboxGoExpr
  , ReboxFields
  , registerReboxPair
  , findReboxFields
  , renderReboxFunction
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
import Debug as Debug
import Effect (Effect)
import Effect.Console as Console
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CodegenState (CodegenState)
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, sanitizeName)
import Gopurs.GoTypes as GoTypes
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))
import PureScript.Backend.Optimizer.FfiSupport (hashString)

-- Native value layouts supported by the expression generator. These rules
-- also define how their payloads cross the runtime Value boundary.
type UnboxedADT =
  { signature :: Array GoType
  , mapConstructor :: String -> Array GoExpr -> Array GoExpr
  , boxExpr :: GoExpr -> GoExpr
  , unboxExpr :: GoExpr -> GoExpr
  }

unboxableADTs :: Map String UnboxedADT
unboxableADTs = Map.fromFoldable
  [ Tuple "Data.Maybe.Maybe"
      { signature: [ TypeValue, TypeBool ]
      , mapConstructor: \ctorName args ->
          case ctorName of
            "Just" -> [ fromMaybe (GoRaw "gopurs_runtime.Value{}") (Array.index args 0), GoRaw "true" ]
            "Nothing" -> [ GoRaw "gopurs_runtime.Value{}", GoRaw "false" ]
            _ -> args
      , boxExpr: \expr ->
          -- Maybe uses the Just constructor id for both cases, with a nil pointer for Nothing.
          GoRaw ("func() gopurs_runtime.Value {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.V1 {\n\t\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Maybe_Just" <> ", UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}\n\t\t\t\t}\n\t\t\t\treturn gopurs_runtime.Value{Type: 9, IntVal: " <> hashString "Data_Data_Maybe_Just" <> "}\n\t\t\t}()")
      , unboxExpr: \expr ->
          GoRaw ("func() struct{V0 gopurs_runtime.Value; V1 bool} {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.Type == 9 && _v.IntVal == " <> hashString "Data_Data_Maybe_Just" <> " && _v.UnsafePtr != nil {\n\t\t\t\t\treturn struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}\n\t\t\t\t}\n\t\t\t\treturn struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}\n\t\t\t}()")
      }
  , Tuple "Data.Tuple.Tuple"
      { signature: [ TypeValue, TypeValue ]
      , mapConstructor: \ctorName args -> args
      , boxExpr: \expr ->
          GoRaw ("func() gopurs_runtime.Value {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\treturn gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})\n\t\t\t}()")
      , unboxExpr: \expr ->
          GoRaw ("func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\t_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)\n\t\t\t\treturn struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}\n\t\t\t}()")
      }
  , Tuple "Data.Either.Either"
      { signature: [ TypeValue, TypeValue, TypeBool ]
      , mapConstructor: \ctor args -> case ctor of
          "Left" -> [ fromMaybe (GoRaw "gopurs_runtime.Value{}") (Array.head args), GoRaw "gopurs_runtime.Value{}", GoRaw "false" ]
          "Right" -> [ GoRaw "gopurs_runtime.Value{}", fromMaybe (GoRaw "gopurs_runtime.Value{}") (Array.head args), GoRaw "true" ]
          _ -> [ GoRaw "gopurs_runtime.Value{}", GoRaw "gopurs_runtime.Value{}", GoRaw "false" ]
      , boxExpr: \expr ->
          GoRaw ("func() gopurs_runtime.Value {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.V2 {\n\t\t\t\t\treturn gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})\n\t\t\t\t}\n\t\t\t\treturn gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})\n\t\t\t}()")
      , unboxExpr: \expr ->
          GoRaw ("func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {\n\t\t\t\t_v := " <> printGoExpr expr <> "\n\t\t\t\tif _v.Type == 9 && _v.IntVal == " <> hashString "Constructor_Data_Either_Right" <> " && _v.UnsafePtr != nil {\n\t\t\t\t\treturn struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}\n\t\t\t\t}\n\t\t\t\treturn struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}\n\t\t\t}()")
      }
  ]

getUnboxedADT :: ExprType -> Maybe (Tuple String UnboxedADT)
getUnboxedADT (ADT fullName _ _) = do
  adt <- Map.lookup fullName unboxableADTs
  pure (Tuple fullName adt)
getUnboxedADT (TypeApp f a) = getUnboxedADT f
getUnboxedADT _ = Nothing

-- Coercions can request Rebox helpers through the state of the current
-- translation. Keep their registration and generation in this module.
coerceGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoType -> GoExpr
coerceGoExpr _ modNameStr expr from to | from == to = expr
coerceGoExpr _ modNameStr ctor@(GoConstructor _ structName _ args) _ (TypeStructValue adtName fields) =
  let
    parts = String.split (Pattern "_") structName
    ctorName = fromMaybe "" (Array.last parts)
  in
    case Map.lookup adtName unboxableADTs of
      Just adt -> GoStructValue adtName fields (adt.mapConstructor ctorName args)
      Nothing -> ctor -- fallback
coerceGoExpr _ modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) | b1 == b2 && s1 == s2 && a1 == a2 = expr

coerceGoExpr codegenStateRef modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) | b1 == b2 =
  let
    _register = unsafePerformEffect (registerReboxPair codegenStateRef srcT destT)
  in
    GoCall (GoVar ("Rebox_" <> modNameStr <> "_" <> hashString s1 <> "_" <> hashString s2)) [ expr ]

coerceGoExpr codegenStateRef modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) = Debug.trace ("MISMATCH B1 B2: " <> b1 <> " vs " <> b2 <> " from " <> goTypeToStr srcT <> " to " <> goTypeToStr destT) \_ ->
  unboxGoExpr codegenStateRef modNameStr (boxGoExpr codegenStateRef modNameStr expr srcT) TypeValue destT

coerceGoExpr codegenStateRef modNameStr expr srcT@(TypeStructPointer b1 f1 s1 a1) destT@(TypeStructPointer b2 f2 s2 a2) | b1 == b2 =
  let
    _register = unsafePerformEffect (registerReboxPair codegenStateRef srcT destT)
  in
    GoCall (GoVar ("Rebox_" <> modNameStr <> "_" <> hashString s1 <> "_" <> hashString s2)) [ expr ]

coerceGoExpr codegenStateRef modNameStr expr srcT@(TypeStructPointer b f s a) TypeValue | Array.any (_ /= TypeValue) a =
  let
    basePath = case String.indexOf (Pattern "[") s of
      Just i -> String.take i s
      Nothing -> s
    destT = TypeStructPointer b f (basePath <> if Array.length a > 0 then "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") a) <> "]" else "") (map (const TypeValue) a)
  in
    boxGoExpr codegenStateRef modNameStr (coerceGoExpr codegenStateRef modNameStr expr srcT destT) destT

coerceGoExpr codegenStateRef modNameStr expr TypeValue destT@(TypeStructPointer b f s a) | Array.any (_ /= TypeValue) a =
  let
    basePath = case String.indexOf (Pattern "[") s of
      Just i -> String.take i s
      Nothing -> s
    srcT = TypeStructPointer b f (basePath <> if Array.length a > 0 then "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") a) <> "]" else "") (map (const TypeValue) a)
  in
    coerceGoExpr codegenStateRef modNameStr (unboxGoExpr codegenStateRef modNameStr expr TypeValue srcT) srcT destT

coerceGoExpr codegenStateRef modNameStr expr srcT@(TypeStructValue "Data.Maybe.Maybe" _) destT@(TypeStructPointer _ "Data.Maybe.Maybe" _ _) =
  -- The boxed payload is Value; reuse the typed-pointer conversion to rebox it.
  coerceGoExpr codegenStateRef modNameStr (boxGoExpr codegenStateRef modNameStr expr srcT) TypeValue destT

coerceGoExpr codegenStateRef modNameStr expr from TypeValue = boxGoExpr codegenStateRef modNameStr expr from
coerceGoExpr codegenStateRef modNameStr expr TypeValue to = unboxGoExpr codegenStateRef modNameStr expr TypeValue to
coerceGoExpr codegenStateRef modNameStr expr from to = unboxGoExpr codegenStateRef modNameStr (boxGoExpr codegenStateRef modNameStr expr from) TypeValue to

-- Before boxing a generic pointer, Rebox converts its type arguments to
-- Value. boxGoExprImpl then emits the box without repeating that conversion.
boxGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoExpr
boxGoExpr codegenStateRef modNameStr expr srcT@(TypeStructPointer baseStructName fullName fullPath typeArgs) =
  if Array.any (\t -> t /= TypeValue) typeArgs then
    coerceGoExpr codegenStateRef modNameStr expr srcT TypeValue
  else
    boxGoExprImpl codegenStateRef modNameStr expr srcT
boxGoExpr codegenStateRef modNameStr expr t = boxGoExprImpl codegenStateRef modNameStr expr t

-- Records capture their source once and box each field. Int arrays have a
-- dedicated helper; other typed array elements are converted recursively.
boxGoExprImpl :: Ref CodegenState -> String -> GoExpr -> GoType -> GoExpr
boxGoExprImpl _ modNameStr expr TypeValue = expr
boxGoExprImpl _ modNameStr expr TypeInt64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ expr ]
boxGoExprImpl _ modNameStr expr TypeFloat64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Float") [ expr ]
boxGoExprImpl _ modNameStr expr TypeString = GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ expr ]
boxGoExprImpl _ modNameStr expr TypeBool = GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ expr ]
boxGoExprImpl _ modNameStr expr (TypeStructPointer baseStructName _ _ _) = GoRaw ("gopurs_runtime.Value{Type: 9, IntVal: " <> hashString baseStructName <> ", UnsafePtr: unsafe.Pointer(" <> printGoExpr expr <> ")}")
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
    GoRaw ("func() gopurs_runtime.Value {\n\t\t\t\torig := " <> printGoExpr expr <> "\n\t\t\t\t_ = orig\n\t\t\t\treturn " <> boxedRecord <> "\n\t\t\t\t}()")
boxGoExprImpl _ modNameStr expr (TypeInterface _) = expr
boxGoExprImpl _ modNameStr expr (TypeNativeArray TypeValue) = GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ expr ]
boxGoExprImpl _ _ expr (TypeNativeArray TypeInt64) = GoBoxIntArray expr
boxGoExprImpl codegenStateRef modNameStr expr (TypeNativeArray inner) = GoRaw ("func() gopurs_runtime.Value {\n\t\t\t\t\tarr := " <> printGoExpr expr <> "\n\t\t\t\t\tboxed := make([]gopurs_runtime.Value, len(arr))\n\t\t\t\t\tfor i, v := range arr { boxed[i] = " <> printGoExpr (boxGoExpr codegenStateRef modNameStr (GoVar "v") inner) <> " }\n\t\t\t\t\treturn gopurs_runtime.Array(boxed)\n\t\t\t\t}()")
boxGoExprImpl _ modNameStr expr TypeUint32 = GoRaw ("gopurs_runtime.Value{Type: 9, IntVal: int64(" <> printGoExpr expr <> "), UnsafePtr: nil}")
boxGoExprImpl _ modNameStr expr (TypeGenericParam _) = expr
boxGoExprImpl _ modNameStr expr (TypeFunc _ _) = expr
boxGoExprImpl _ modNameStr expr (TypeStructValue adtName fields) =
  case Map.lookup adtName unboxableADTs of
    Just adt -> adt.boxExpr expr
    Nothing -> GoRaw ("func() gopurs_runtime.Value {\n\t\t\t\t_ = " <> printGoExpr expr <> "\n\t\t\t\tpanic(\"boxTypeStructValue not implemented yet for " <> adtName <> "\")\n\t\t\t}()")

-- The destination GoType determines how to read Value: native record fields,
-- array elements, primitive payloads, or an ADT pointer.
unboxGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoType -> GoExpr
unboxGoExpr codegenStateRef modNameStr expr currentType desiredType =
  if currentType == desiredType then expr
  else if goTypeToStr currentType == goTypeToStr desiredType && String.contains (Pattern "Constructor_Test_RBTree_T") (goTypeToStr currentType) then
    let
      cArgs = case currentType of
        TypeStructPointer b1 k1 f1 args1 -> "cBase=" <> b1 <> ", cKey=" <> k1 <> ", cFull=" <> f1 <> ", cLen=" <> show (Array.length args1)
        _ -> "none"
      dArgs = case desiredType of
        TypeStructPointer b2 k2 f2 args2 -> "dBase=" <> b2 <> ", dKey=" <> k2 <> ", dFull=" <> f2 <> ", dLen=" <> show (Array.length args2)
        _ -> "none"
    in
      Debug.trace ("MISMATCH AGAIN: " <> cArgs <> " vs " <> dArgs <> ". Structurally equal arrays? " <> show (currentType == desiredType)) \_ ->
        unboxGoExpr codegenStateRef modNameStr (boxGoExpr codegenStateRef modNameStr expr currentType) TypeValue desiredType
  else if currentType /= TypeValue then
    unboxGoExpr codegenStateRef modNameStr (boxGoExpr codegenStateRef modNameStr expr currentType) TypeValue desiredType
  else case desiredType of
    TypeValue -> boxGoExpr codegenStateRef modNameStr expr currentType
    (TypeRecord fields) ->
      let
        assignments = String.joinWith "\n" (map (\(Tuple k v) -> "\t\t\t\t\tclone." <> sanitizeName k <> " = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr (GoCall (GoSelector (GoVar "gopurs_runtime") "RecordGet") [ GoVar "orig", GoString k ]) TypeValue v)) fields)
      in
        GoRaw ("func() " <> goTypeToStr desiredType <> " {\n\t\t\t\t\torig := " <> printGoExpr expr <> "\n\t\t\t\t\t_ = orig\n\t\t\t\t\tclone := " <> goTypeToStr desiredType <> "{}\n" <> assignments <> "\n\t\t\t\t\treturn clone\n\t\t\t\t}()")
    TypeInt64 -> GoSelector expr "IntVal"
    TypeFloat64 -> GoCall (GoSelector expr "FloatVal") []
    TypeString -> GoCall (GoSelector expr "StrVal") []
    TypeBool -> GoBinOp "!=" (GoSelector expr "IntVal") (GoInt 0)
    TypeUint32 -> GoRaw ("uint32(" <> printGoExpr (GoSelector expr "IntVal") <> ")")
    (TypeStructPointer _ _ fullPath _) -> GoCall (GoRaw ("gopurs_runtime.CoerceToStruct[" <> fullPath <> "]")) [ expr ]
    (TypeInterface _) -> expr
    (TypeNativeArray TypeInt64) -> GoUnboxIntArray expr
    (TypeNativeArray inner) -> case currentType of
      TypeNativeArray currentInner ->
        GoRaw ("func() " <> goTypeToStr desiredType <> " {\n\t\t\t\t\tarr := " <> printGoExpr expr <> "\n\t\t\t\t\tunboxed := make(" <> goTypeToStr desiredType <> ", len(arr))\n\t\t\t\t\tfor i, v := range arr { unboxed[i] = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr (GoVar "v") currentInner inner) <> " }\n\t\t\t\t\treturn unboxed\n\t\t\t\t}()")
      _ ->
        GoRaw ("func() " <> goTypeToStr desiredType <> " {\n\t\t\t\t\tarr := *(*[]gopurs_runtime.Value)(" <> printGoExpr expr <> ".UnsafePtr)\n\t\t\t\t\tunboxed := make(" <> goTypeToStr desiredType <> ", len(arr))\n\t\t\t\t\tfor i, v := range arr { unboxed[i] = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr (GoVar "v") TypeValue inner) <> " }\n\t\t\t\t\treturn unboxed\n\t\t\t\t}()")
    (TypeGenericParam _) -> expr
    (TypeFunc _ _) -> expr
    (TypeStructValue adtName fields) ->
      case Map.lookup adtName unboxableADTs of
        Just adt -> adt.unboxExpr expr
        Nothing -> GoRaw ("func() " <> goTypeToStr (TypeStructValue adtName fields) <> " {\n\t\t\t\t_ = " <> printGoExpr expr <> "\n\t\t\t\tpanic(\"unboxTypeStructValue not implemented yet for " <> adtName <> "\")\n\t\t\t}()")

type ReboxFields =
  { vars :: Array String
  , fields :: Array ExprType
  }

registerReboxPair :: Ref CodegenState -> GoType -> GoType -> Effect Unit
registerReboxPair codegenStateRef srcT destT = do
  state <- Ref.read codegenStateRef
  let pairs = state.reboxPairs
  if Set.member (Tuple srcT destT) pairs then pure unit
  else Ref.modify_ (\s -> s { reboxPairs = Set.insert (Tuple srcT destT) pairs }) codegenStateRef

findReboxFields :: CodegenState -> String -> Maybe ReboxFields
findReboxFields helpers baseStructName =
  let
    matchesConstructor key =
      let
        parts = String.split (Pattern ".") key
      in
        if Array.length parts >= 2 then
          let
            ctorName = fromMaybe "" (Array.last parts)
            pkgName = String.joinWith "_" (Array.slice 0 (Array.length parts - 1) parts)
            constructorName = "Constructor_" <> pkgName <> "_" <> sanitizeName ctorName
            dataName = "Data_" <> pkgName <> "_" <> sanitizeName ctorName
          in
            constructorName == baseStructName || dataName == baseStructName
        else false

    mbCtor = Array.find (\(Tuple key _) -> matchesConstructor key)
      (Map.toUnfoldable helpers.ctorTypes :: Array (Tuple String ReboxFields))
    mbClass = Array.find (\(Tuple key _) -> matchesConstructor key)
      (Map.toUnfoldable helpers.classDeclsFields :: Array (Tuple String { vars :: Array String, fields :: Array { name :: String, "type" :: ExprType } }))
  in
    case mbCtor of
      Just (Tuple _ info) -> Just info
      Nothing -> case mbClass of
        Just (Tuple _ classInfo) ->
          Just { vars: classInfo.vars, fields: map (\field -> field."type") classInfo.fields }
        Nothing ->
          let
            _trace = unsafePerformEffect (Console.log ("ERROR: Rebox missing! b1=" <> baseStructName <> " keysCtor: " <> String.joinWith ", " (map fst (Map.toUnfoldable helpers.ctorTypes :: Array (Tuple String _)))))
          in
            Nothing

renderReboxFunction :: Ref CodegenState -> CodegenState -> String -> Map String String -> Tuple GoType GoType -> Maybe (Tuple String String)
renderReboxFunction codegenStateRef helpers modNameStr generatedFuncs (Tuple srcT destT) =
  case srcT, destT of
    TypeStructPointer b1 _ s1 a1, TypeStructPointer b2 _ s2 a2 | b1 == b2 ->
      let
        funcName = "Rebox_" <> modNameStr <> "_" <> hashString s1 <> "_" <> hashString s2
      in
        if Map.member funcName generatedFuncs then Nothing
        else case findReboxFields helpers b1 of
          Just info ->
            let
              env1 = Map.fromFoldable (Array.zip info.vars a1)
              env2 = Map.fromFoldable (Array.zip info.vars a2)
              assignments = String.joinWith "\n" (Array.mapWithIndex
                (\i fieldExprType ->
                  let
                    genericTy = GoTypes.structFieldGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors info.vars modNameStr fieldExprType
                    t1 = GoTypes.instantiateGenericGoType env1 genericTy
                    t2 = GoTypes.instantiateGenericGoType env2 genericTy
                  in
                    "\t\tout.V" <> show i <> " = " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)
                )
                info.fields)
              funcBody = "func " <> funcName <> "(in *" <> s1 <> ") *" <> s2 <> " {\n\tif in == nil { return nil }\n\tout := &" <> s2 <> "{}\n" <> assignments <> "\n\treturn out\n}"
            in
              Just (Tuple funcName funcBody)
          Nothing -> Nothing
    _, _ -> Nothing

generateReboxFunctions :: Ref CodegenState -> String -> Effect (Array String)
generateReboxFunctions codegenStateRef modNameStr = loop Map.empty
  where
  -- Rendering fields can register more conversions; collect until none remain.
  loop generatedFuncs = do
    helpers <- Ref.read codegenStateRef
    let
      reboxPairs = helpers.reboxPairs
      newFuncs = Map.fromFoldable $
        Array.mapMaybe (renderReboxFunction codegenStateRef helpers modNameStr generatedFuncs) (Array.fromFoldable reboxPairs)
      nextGeneratedFuncs = Map.union generatedFuncs newFuncs
    if Map.isEmpty newFuncs then
      pure $ Array.fromFoldable (Map.values nextGeneratedFuncs)
    else
      loop nextGeneratedFuncs
