module Gopurs.AdtExprs
  ( AdtExpr
  , definition
  , SaturatedConstructor
  , prepareSaturated
  , saturatedFieldType
  , saturated
  , getField
  , isTag
  , constructorReuse
  ) where

import Prelude

import Control.Alternative (guard)

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CodegenState (CodegenState)
import Gopurs.ExprContext (ExprResult, LocalEnv, StmtTree(..))
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, sanitizeName, getStructName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr, unboxGoExpr, unboxableADTs)
import Gopurs.GoTypes (exprTypeToGoType, exprTypeToGenericGoType, instantiateGenericGoType, structFieldGoType)
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName(..))
import PureScript.Backend.Optimizer.FfiSupport (hashString)

type AdtExpr =
  { expr :: GoExpr
  , exprType :: GoType
  }

-- The caller supplies the constructor's result type after resolving annotations
-- and function wrappers; metadata determines its native representation.
definition :: Ref CodegenState -> String -> String -> Array String -> ExprType -> AdtExpr
definition codegenStateRef modNameStr name fields ctorType =
  let
    helpers = unsafePerformEffect (Ref.read codegenStateRef)
    expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr ctorType

    trueModPart = case expectedGoType of
      TypeStructPointer _ fn _ _ ->
        let
          parts = String.split (Pattern ".") fn
        in
          String.joinWith "." (Array.slice 0 (Array.length parts - 1) parts)
      _ -> modNameStr

    trueModPartUnderscores = String.replaceAll (Pattern ".") (Replacement "_") trueModPart

    structName = "Constructor_" <> trueModPartUnderscores <> "_" <> sanitizeName name
    baseStructName = "Data_" <> trueModPartUnderscores <> "_" <> sanitizeName name
    key = trueModPartUnderscores <> "." <> name

    fullName = case expectedGoType of
      TypeStructPointer _ fn _ _ -> fn
      _ -> if key == "Test_RBTree.E" then "Test.RBTree.Tree" else key

    ctorInfo = Map.lookup key helpers.ctorTypes
    classInfo = Map.lookup fullName helpers.classDeclsFields

    fields' = case ctorInfo of
      Just info -> info.fields
      Nothing -> case classInfo of
        Just info -> map _."type" info.fields
        Nothing -> []

    vars' = case ctorInfo of
      Just info -> info.vars
      Nothing -> case classInfo of
        Just info -> info.vars
        Nothing -> []

    typeArgs = case ctorType of
      ADT adtName _ tArgs ->
        let
          mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) tArgs
          arity = case Map.lookup adtName (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths of
            Just info -> info.arity
            Nothing -> Array.length vars'
        in
          Array.take arity mapped
      TypeApp _ _ ->
        let
          unwrapTypeApp (TypeApp f a) acc = unwrapTypeApp f (a <> acc)
          unwrapTypeApp other acc = Tuple other acc
        in
          case unwrapTypeApp ctorType [] of
            Tuple (ADT _ _ tArgs) allArgs ->
              let
                mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) (tArgs <> allArgs)
              in
                Array.take (Array.length vars') mapped
            _ -> map (const TypeValue) vars'
      _ -> map (const TypeValue) vars'
    typeArgsStr = if Array.length typeArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr typeArgs) <> "]" else ""
    instMap = Map.fromFoldable (Array.zip vars' typeArgs)
    coercedFields = Array.mapWithIndex
      ( \i f ->
          let
            expectedType = case Array.index fields' i of
              Just ty ->
                let
                  genericGoType = exprTypeToGenericGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors vars' modNameStr ty
                in
                  instantiateGenericGoType instMap genericGoType
              Nothing -> TypeValue
          in
            coerceGoExpr codegenStateRef modNameStr (GoVar (sanitizeName f)) TypeValue expectedType
      )
      fields
    isElided = Set.member structName helpers.elidedCtors
    isPointerAdtLeaf = Map.member baseStructName helpers.pointerAdtLeaves
    isEnum = Set.member baseStructName helpers.enumCtors
    boxedCtor = boxGoExpr codegenStateRef modNameStr (GoConstructor (hashString baseStructName) structName typeArgs coercedFields) (TypeStructPointer baseStructName fullName (structName <> typeArgsStr) typeArgs)

    finalExprType =
      if isEnum then TypeUint32
      else if isPointerAdtLeaf then
        let
          nodeInfo = Map.lookup baseStructName helpers.pointerAdtLeaves
          nodeBaseStruct = case nodeInfo of
            Just info -> info.nodeBaseStruct
            Nothing -> ""
          nodeCtorName = case nodeInfo of
            Just info -> info.nodeCtor
            Nothing -> ""
          nodeStruct = "Constructor_" <> modNameStr <> "_" <> sanitizeName nodeCtorName
          nodeFullPath = nodeStruct <> typeArgsStr
        in
          TypeStructPointer nodeBaseStruct fullName nodeFullPath typeArgs
      else if Array.length fields == 0 then
        TypeStructPointer baseStructName fullName (structName <> typeArgsStr) typeArgs
      else TypeValue

    funcExpr =
      if isElided then
        case Array.head fields of
          Just f -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr (GoVar (sanitizeName f)) TypeValue TypeValue) <> "\n}") ]
          Nothing -> Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr inner <> "\n}") ]) boxedCtor fields
      else if isPointerAdtLeaf then
        let
          nodeCtorName = case Map.lookup baseStructName helpers.pointerAdtLeaves of
            Just info -> info.nodeCtor
            Nothing -> ""
          nodeStruct = "Constructor_" <> modNameStr <> "_" <> sanitizeName nodeCtorName <> typeArgsStr
        in
          GoRaw ("(*" <> nodeStruct <> ")(nil)")
      else if isEnum then GoRaw (hashString baseStructName)
      else if Array.length fields == 0 then
        GoConstructor (hashString baseStructName) structName typeArgs coercedFields
      else
        Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr inner <> "\n}") ]) boxedCtor fields
  in
    { expr: funcExpr, exprType: finalExprType }

-- Opaque preparation shared by field typing and final construction. The
-- metadata snapshot is retained while child expressions register conversions.
newtype SaturatedConstructor = SaturatedConstructor
  { helpers :: CodegenState
  , modNameStr :: String
  , name :: String
  , ctorType :: ExprType
  , trueModPartUnderscores :: String
  , baseStructName :: String
  , adtFullName :: Maybe String
  , structName :: String
  , key :: String
  , fullName :: String
  , fields :: Array ExprType
  , vars :: Array String
  , instMap :: Map String GoType
  }

prepareSaturated :: Ref CodegenState -> String -> Maybe ModuleName -> String -> ExprType -> SaturatedConstructor
prepareSaturated codegenStateRef modNameStr mbMod name ctorType =
  let
    helpers = unsafePerformEffect (Ref.read codegenStateRef)
    expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr ctorType

    modPart = case mbMod of
      Just (ModuleName mn) -> mn
      Nothing -> modNameStr

    trueModPart = case expectedGoType of
      TypeStructPointer _ fn _ _ ->
        let
          parts = String.split (Pattern ".") fn
        in
          String.joinWith "." (Array.slice 0 (Array.length parts - 1) parts)
      _ -> modPart

    trueModPartUnderscores = String.replaceAll (Pattern ".") (Replacement "_") trueModPart

    baseStructName = "Data_" <> trueModPartUnderscores <> "_" <> sanitizeName name
    adtFullName = case ctorType of
      ADT fn _ _ -> Just fn
      _ -> Nothing

    structName = "Constructor_" <> trueModPartUnderscores <> "_" <> sanitizeName name
    key = trueModPartUnderscores <> "." <> name

    fullName = case expectedGoType of
      TypeStructPointer _ fn _ _ -> fn
      _ -> if key == "Test_RBTree.E" then "Test.RBTree.Tree" else key

    ctorInfo = Map.lookup key helpers.ctorTypes
    classInfo = Map.lookup fullName helpers.classDeclsFields

    fields = case ctorInfo of
      Just info -> info.fields
      Nothing -> case classInfo of
        Just info -> map _."type" info.fields
        Nothing -> []

    vars = case ctorInfo of
      Just info -> info.vars
      Nothing -> case classInfo of
        Just info -> info.vars
        Nothing -> []

    typeArgs = case ctorType of
      ADT adtName _ tArgs ->
        let
          mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) tArgs
          arity = case Map.lookup adtName (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths of
            Just info -> info.arity
            Nothing -> Array.length vars
        in
          Array.take arity mapped
      _ -> map (const TypeValue) vars
    instMap = Map.fromFoldable (Array.zip vars typeArgs)
  in
    SaturatedConstructor
      { helpers
      , modNameStr
      , name
      , ctorType
      , trueModPartUnderscores
      , baseStructName
      , adtFullName
      , structName
      , key
      , fullName
      , fields
      , vars
      , instMap
      }

saturatedFieldType :: SaturatedConstructor -> Int -> { exprType :: ExprType, goType :: GoType }
saturatedFieldType (SaturatedConstructor { helpers, modNameStr, adtFullName, fields, vars, instMap }) fieldIdx =
  let
    expectedExprType = case Array.index fields fieldIdx of
      Just ty -> ty
      Nothing -> Any
    expectedType = case adtFullName >>= \fn -> Map.lookup fn unboxableADTs of
      Just adt -> fromMaybe TypeValue (Array.index adt.signature fieldIdx)
      Nothing -> case Array.index fields fieldIdx of
        Just ty ->
          let
            genericGoType = exprTypeToGenericGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors vars modNameStr ty
          in
            instantiateGenericGoType instMap genericGoType
        Nothing -> TypeValue
  in
    { exprType: expectedExprType, goType: expectedType }

-- Fields arrive in source order, already translated and immediately coerced.
-- Constructor reuse and statement assembly remain with the caller.
saturated :: Ref CodegenState -> SaturatedConstructor -> { exprs :: Array GoExpr, exprTypes :: Array GoType } -> AdtExpr
saturated codegenStateRef (SaturatedConstructor { helpers, modNameStr, name, ctorType, trueModPartUnderscores, baseStructName, adtFullName, structName, key, fullName }) accProps =
  let
    isElided = Set.member structName helpers.elidedCtors
    isPointerAdtLeaf = Map.member baseStructName helpers.pointerAdtLeaves

    modPart' = trueModPartUnderscores
    monoStructName = "Constructor_" <> modPart' <> "_" <> sanitizeName name

    typeArgsCtor = case ctorType of
      ADT adtName _ tArgs ->
        let
          mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) tArgs
          arity = case Map.lookup adtName (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths of
            Just info -> info.arity
            Nothing -> Array.length mapped
        in
          Array.take arity mapped
      _ -> case Map.lookup key helpers.ctorTypes of
        Just ctorInfo -> map (const TypeValue) ctorInfo.vars
        Nothing -> []
    typeArgsStr = if Array.length typeArgsCtor > 0 then "[" <> String.joinWith ", " (map goTypeToStr typeArgsCtor) <> "]" else ""
    fullPath = monoStructName <> typeArgsStr
    isEnum = Set.member baseStructName helpers.enumCtors
    res =
      if isElided then
        case Array.head accProps.exprs of
          -- The field has already been coerced to its instantiated
          -- representation; its declaration may still contain a type variable.
          Just expr -> { expr: boxGoExpr codegenStateRef modNameStr expr (fromMaybe TypeValue (Array.head accProps.exprTypes)), exprType: TypeValue }
          Nothing -> { expr: GoConstructor (hashString baseStructName) monoStructName typeArgsCtor accProps.exprs, exprType: TypeStructPointer baseStructName fullName fullPath typeArgsCtor }
      else case adtFullName >>= \fn -> Map.lookup fn unboxableADTs >>= \adt -> Just (Tuple fn adt) of
        Just (Tuple fn adt) ->
          { expr: GoStructValue fn adt.signature (adt.mapConstructor name accProps.exprs)
          , exprType: TypeStructValue fn adt.signature
          }
        Nothing ->
          if isPointerAdtLeaf then
            let
              nodeInfo = Map.lookup baseStructName helpers.pointerAdtLeaves
              nodeCtorName = case nodeInfo of
                Just info -> info.nodeCtor
                Nothing -> ""
              nodeBaseStruct = case nodeInfo of
                Just info -> info.nodeBaseStruct
                Nothing -> ""
              nodeStruct = "Constructor_" <> modPart' <> "_" <> sanitizeName nodeCtorName
              nodeFullPath = nodeStruct <> typeArgsStr
            in
              { expr: GoRaw ("(*" <> nodeFullPath <> ")(nil)"), exprType: TypeStructPointer nodeBaseStruct fullName nodeFullPath typeArgsCtor }
          else if isEnum then { expr: GoRaw (hashString baseStructName), exprType: TypeUint32 }
          else { expr: GoConstructor (hashString baseStructName) monoStructName typeArgsCtor accProps.exprs, exprType: TypeStructPointer baseStructName fullName fullPath typeArgsCtor }
  in
    res

-- The erased-constructor set is the caller's snapshot from before translating
-- the object. Its statements and identifiers remain with the caller.
getField
  :: Ref CodegenState
  -> String
  -> Set String
  -> { moduleName :: Maybe ModuleName, ctorName :: String, index :: Int }
  -> { expr :: GoExpr, exprType :: GoType, sourceType :: ExprType }
  -> AdtExpr
getField codegenStateRef modNameStr elidedCtors { moduleName: mbMod, ctorName, index: idx } object =
  let
    defMod = case mbMod of
      Just (ModuleName mod) -> String.replaceAll (Pattern ".") (Replacement "_") mod
      Nothing -> modNameStr
    structName = "Constructor_" <> defMod <> "_" <> sanitizeName ctorName
    key = defMod <> "." <> ctorName
    helpers = unsafePerformEffect (Ref.read codegenStateRef)
  in
    if Set.member structName elidedCtors then
      { expr: coerceGoExpr codegenStateRef modNameStr object.expr object.exprType TypeValue, exprType: TypeValue }
    else
      let
        fields = fromMaybe [] (map _.fields (Map.lookup key helpers.ctorTypes))
        monoStructName = structName

        expectedType = case Array.index fields idx of
          Just ty -> exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr ty
          Nothing -> TypeValue

        typeArgs = case object.exprType of
          TypeStructPointer _ _ _ tArgs -> tArgs
          _ -> case object.sourceType of
            ADT fullName _ tArgs ->
              let
                mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) tArgs
                arity = case Map.lookup fullName (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths of
                  Just info -> info.arity
                  Nothing -> Array.length mapped
              in
                Array.take arity mapped
            _ -> case Map.lookup key helpers.ctorTypes of
              Just ctorInfo -> map (const TypeValue) ctorInfo.vars
              Nothing -> []

        isNative = case object.exprType of
          TypeStructPointer _ _ _ _ -> true
          _ -> false

        actualFieldType = case object.exprType of
          TypeStructPointer _ _ _ tArgs ->
            case Map.lookup key helpers.ctorTypes of
              Just ctorInfo ->
                let
                  env = Map.fromFoldable (Array.zip ctorInfo.vars tArgs)
                  genericTy = structFieldGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors ctorInfo.vars modNameStr (fromMaybe (TypeVar "") (Array.index fields idx))
                in
                  instantiateGenericGoType env genericTy
              Nothing -> expectedType
          _ -> expectedType

        exprAccess =
          if isNative then
            GoConstructorAccess object.expr monoStructName typeArgs idx true
          else
            GoConstructorAccess (boxGoExpr codegenStateRef modNameStr object.expr object.exprType) monoStructName typeArgs idx false
      in
        { expr: exprAccess, exprType: actualFieldType }

-- Reuse an unchanged constructor without mutating it. Restrict this to direct
-- projections of one typed local, with exactly one constant field replaced.
-- Comparing Number fields would be unsound for observable signed zero.
constructorReuse :: LocalEnv -> GoType -> Array Boolean -> GoExpr -> Maybe { source :: GoExpr, condition :: GoExpr }
constructorReuse bound resultType constants constructor = case resultType, constructor of
  TypeStructPointer _ _ _ _, GoConstructor _ ctor typeArgs fields ->
    case Array.catMaybes (Array.mapWithIndex (\index constant -> if constant then Just index else Nothing) constants) of
      [ changedIndex ] -> do
        guard (Array.length constants == Array.length fields)
        sourceName <- Array.head
          ( Array.mapMaybe
              ( \(Tuple index field) -> case field of
                  GoConstructorAccess (GoVar name) sourceCtor sourceTypeArgs sourceIndex true
                    | index /= changedIndex && sourceCtor == ctor
                        && sourceTypeArgs == typeArgs
                        && sourceIndex == index -> Just name
                  _ -> Nothing
              )
              (Array.mapWithIndex Tuple fields)
          )
        sourceBinding <- Array.find (\binding -> binding.name == sourceName)
          (Array.fromFoldable (Map.values bound))
        guard (sourceBinding.goType == resultType)
        let
          source = GoVar sourceName
          projection index = GoConstructorAccess source ctor typeArgs index true
        guard
          ( Array.all identity
              ( Array.mapWithIndex
                  (\index field -> index == changedIndex || field == projection index)
                  fields
              )
          )
        replacement <- Array.index fields changedIndex
        pure
          { source
          , condition: GoBinOp "&&" (GoBinOp "!=" source (GoRaw "nil"))
              (GoBinOp "==" (projection changedIndex) replacement)
          }
      _ -> Nothing
  _, _ -> Nothing

-- The operand has already been translated. A temporary keeps non-variable
-- operands evaluated exactly once, including constant native tag tests.
isTag :: Ref CodegenState -> String -> Maybe ModuleName -> String -> ExprResult -> ExprResult
isTag codegenStateRef modNameStr mbMod tag resE =
  let
    baseStructName = getStructName modNameStr mbMod tag
    hashStr = hashString baseStructName
    helpers = unsafePerformEffect (Ref.read codegenStateRef)
    nativeTagTest = case resE.exprType of
      TypeStructValue adtName _ -> map (\adt -> adt.isConstructor baseStructName) (Map.lookup adtName unboxableADTs)
      _ -> Nothing

    isNativePointer = case resE.exprType of
      TypeStructPointer typedBaseStructName _ _ _ ->
        typedBaseStructName == baseStructName ||
          ( case Map.lookup baseStructName helpers.pointerAdtLeaves of
              Just nodeInfo -> typedBaseStructName == nodeInfo.nodeBaseStruct
              Nothing -> false
          )
      _ -> false
  in
    case resE.expr of
      GoVar _ ->
        let
          exprStr = case nativeTagTest of
            Just test -> printGoExpr (test resE.expr)
            Nothing ->
              if isNativePointer then
                case Map.lookup baseStructName helpers.pointerAdtLeaves of
                  Just _ -> "(" <> printGoExpr resE.expr <> " == nil)"
                  Nothing -> "(" <> printGoExpr resE.expr <> " != nil)"
              else case Map.lookup baseStructName helpers.pointerAdtLeaves of
                Just nodeInfo -> "(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".IntVal == " <> hashString nodeInfo.nodeBaseStruct <> " && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".UnsafePtr == nil)"
                Nothing ->
                  if Set.member baseStructName helpers.pointerAdtNodes then
                    "(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> " && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".UnsafePtr != nil)"
                  else if Set.member baseStructName helpers.enumCtors then
                    "(" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeUint32) <> " == " <> hashStr <> ")"
                  else
                    "(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> ")"
        in
          { stmts: resE.stmts, expr: GoRaw exprStr, exprType: TypeBool, nextId: resE.nextId }
      _ ->
        let
          tmpVar = "__t_tag_" <> show resE.nextId
          declTmp =
            if isNativePointer || resE.exprType /= TypeValue then
              StmtLeaf (GoRaw ("var " <> tmpVar <> " " <> goTypeToStr resE.exprType <> " = " <> printGoExpr resE.expr))
            else
              StmtLeaf (GoRaw ("var " <> tmpVar <> " gopurs_runtime.Value = " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType)))

          exprStr = case nativeTagTest of
            Just test -> printGoExpr (test (GoVar tmpVar))
            Nothing ->
              if isNativePointer then
                case Map.lookup baseStructName helpers.pointerAdtLeaves of
                  Just _ -> "(" <> tmpVar <> " == nil)"
                  Nothing -> "(" <> tmpVar <> " != nil)"
              else if resE.exprType /= TypeValue then
                "(uint32(" <> tmpVar <> ") == " <> hashStr <> ")"
              else case Map.lookup baseStructName helpers.pointerAdtLeaves of
                Just nodeInfo -> "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashString nodeInfo.nodeBaseStruct <> " && " <> tmpVar <> ".UnsafePtr == nil)"
                Nothing ->
                  if Set.member baseStructName helpers.pointerAdtNodes then
                    "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashStr <> " && " <> tmpVar <> ".UnsafePtr != nil)"
                  else if Set.member baseStructName helpers.enumCtors then
                    "(" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr (GoVar tmpVar) TypeValue TypeUint32) <> " == " <> hashStr <> ")"
                  else
                    "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashStr <> ")"
        in
          -- A single-constructor native value has a constant test,
          -- but its operand must still be evaluated exactly once.
          { stmts: resE.stmts <> declTmp <> StmtLeaf (GoRaw ("_ = " <> tmpVar)), expr: GoRaw exprStr, exprType: TypeBool, nextId: resE.nextId + 1 }
