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
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Set as Set
import Data.Tuple (Tuple(..))
import Effect.Ref (Ref)
import Gopurs.CodegenState (CodegenMetadata, CodegenState)
import Gopurs.ExprContext (ExprResult, LocalEnv, StmtTree(..))
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..), goTypeToStr, sanitizeName, getStructName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr, unboxGoExpr, unboxableADTs)
import Gopurs.GoTypes (exprTypeToGoType, instantiateGenericGoType, structFieldGoType)
import Gopurs.ConstructorLayout (PreparedConstructor)
import Gopurs.ConstructorLayout as ConstructorLayout
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName(..))
import PureScript.Backend.Optimizer.FfiSupport (hashString)

type AdtExpr =
  { expr :: GoExpr
  , exprType :: GoType
  }

-- The caller supplies the constructor's result type after resolving annotations
-- and function wrappers; metadata determines its native representation.
definition :: CodegenMetadata -> Ref CodegenState -> String -> String -> Array String -> ExprType -> AdtExpr
definition metadata codegenStateRef modNameStr name fields ctorType =
  let
    prepared = ConstructorLayout.prepare ConstructorLayout.Definition metadata modNameStr modNameStr name ctorType
    { identity, fields: fieldInfo } = prepared.layout
    { structName, baseStructName } = identity
    typeArgs = prepared.typeArgs
    coercedFields = Array.mapWithIndex
      (\i f -> coerceGoExpr codegenStateRef modNameStr (GoVar (sanitizeName f)) TypeValue
        (ConstructorLayout.fieldType metadata modNameStr fieldInfo prepared.fieldTypeArgs i))
      fields
    isElided = Set.member structName metadata.elidedCtors
    isEnum = Set.member baseStructName metadata.enumCtors
    boxedCtor = boxGoExpr codegenStateRef modNameStr
      (GoConstructor (hashString baseStructName) structName typeArgs coercedFields) prepared.pointerType
    finalExprType =
      if isEnum then TypeUint32
      else case prepared.leafPointerType of
        Just pointer -> pointer
        Nothing -> if Array.null fields then prepared.pointerType else TypeValue

    funcExpr =
      if isElided then
        case Array.head fields of
          Just f -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoFuncLit [ Tuple (sanitizeName f) TypeValue ] [] (coerceGoExpr codegenStateRef modNameStr (GoVar (sanitizeName f)) TypeValue TypeValue) TypeValue ]
          Nothing -> Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoFuncLit [ Tuple (sanitizeName f) TypeValue ] [] inner TypeValue ]) boxedCtor fields
      else case prepared.leafPointerType of
        Just pointer -> rawGo ("(" <> goTypeToStr pointer <> ")(nil)")
        Nothing ->
          if isEnum then rawGo (hashString baseStructName)
          else if Array.null fields then GoConstructor (hashString baseStructName) structName typeArgs coercedFields
          else Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoFuncLit [ Tuple (sanitizeName f) TypeValue ] [] inner TypeValue ]) boxedCtor fields

  in
    { expr: funcExpr, exprType: finalExprType }

-- Opaque preparation shared by field typing and final construction. The
-- immutable metadata is shared while child expressions register conversions.
newtype SaturatedConstructor = SaturatedConstructor
  { metadata :: CodegenMetadata
  , modNameStr :: String
  , name :: String
  , prepared :: PreparedConstructor
  }

prepareSaturated :: CodegenMetadata -> String -> Maybe ModuleName -> String -> ExprType -> SaturatedConstructor
prepareSaturated metadata modNameStr mbMod name ctorType =
  let
    fallbackModule = case mbMod of
      Just (ModuleName mn) -> mn
      Nothing -> modNameStr
    prepared = ConstructorLayout.prepare ConstructorLayout.Saturated metadata modNameStr fallbackModule name ctorType
  in
    SaturatedConstructor { metadata, modNameStr, name, prepared }

saturatedFieldType :: SaturatedConstructor -> Int -> { exprType :: ExprType, goType :: GoType }
saturatedFieldType (SaturatedConstructor { metadata, modNameStr, prepared }) fieldIdx =
  let
    expectedExprType = fromMaybe Any (Array.index prepared.layout.fields.fields fieldIdx)
    expectedType = case prepared.adtFullName >>= \fn -> Map.lookup fn unboxableADTs of
      Just adt -> fromMaybe TypeValue (Array.index adt.signature fieldIdx)
      Nothing -> ConstructorLayout.fieldType metadata modNameStr prepared.layout.fields prepared.fieldTypeArgs fieldIdx
  in
    { exprType: expectedExprType, goType: expectedType }

-- Fields arrive in source order, already translated and immediately coerced.
-- Constructor reuse and statement assembly remain with the caller.
saturated :: Ref CodegenState -> SaturatedConstructor -> { exprs :: Array GoExpr, exprTypes :: Array GoType } -> AdtExpr
saturated codegenStateRef (SaturatedConstructor { metadata, modNameStr, name, prepared }) accProps =
  let
    { structName, baseStructName } = prepared.layout.identity
    native =
      { expr: GoConstructor (hashString baseStructName) structName prepared.typeArgs accProps.exprs
      , exprType: prepared.pointerType
      }
  in
    if Set.member structName metadata.elidedCtors then
      case Array.head accProps.exprs of
        -- Fields have already been coerced to their instantiated representation.
        Just expr -> { expr: boxGoExpr codegenStateRef modNameStr expr (fromMaybe TypeValue (Array.head accProps.exprTypes)), exprType: TypeValue }
        Nothing -> native
    else case prepared.adtFullName >>= \fn -> Map.lookup fn unboxableADTs >>= \adt -> Just (Tuple fn adt) of
      Just (Tuple fn adt) ->
        { expr: GoStructValue fn adt.signature (adt.mapConstructor name accProps.exprs)
        , exprType: TypeStructValue fn adt.signature
        }
      Nothing -> case prepared.leafPointerType of
        Just pointer -> { expr: rawGo ("(" <> goTypeToStr pointer <> ")(nil)"), exprType: pointer }
        Nothing ->
          if Set.member baseStructName metadata.enumCtors then { expr: rawGo (hashString baseStructName), exprType: TypeUint32 }
          else native

-- The object's statements and identifiers remain with the caller.
getField
  :: CodegenMetadata
  -> Ref CodegenState
  -> String
  -> { moduleName :: Maybe ModuleName, ctorName :: String, index :: Int }
  -> { expr :: GoExpr, exprType :: GoType, sourceType :: ExprType }
  -> AdtExpr
getField metadata codegenStateRef modNameStr { moduleName: mbMod, ctorName, index: idx } object =
  let
    defMod = case mbMod of
      Just (ModuleName mod) -> mod
      Nothing -> modNameStr
    ctorLayout = ConstructorLayout.layout metadata defMod ctorName Nothing
    { structName } = ctorLayout.identity
  in
    if Set.member structName metadata.elidedCtors then
      { expr: coerceGoExpr codegenStateRef modNameStr object.expr object.exprType TypeValue, exprType: TypeValue }
    else
      let
        fields = fromMaybe [] (map _.fields ctorLayout.constructorFields)
        monoStructName = structName

        expectedType = case Array.index fields idx of
          Just ty -> exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr ty
          Nothing -> TypeValue

        typeArgs = case object.exprType of
          TypeStructPointer { typeArgs: tArgs } -> tArgs
          _ -> ConstructorLayout.typeArguments metadata modNameStr Nothing
            (fromMaybe [] (map _.vars ctorLayout.constructorFields)) object.sourceType

        isNative = case object.exprType of
          TypeStructPointer _ -> true
          _ -> false

        actualFieldType = case object.exprType of
          TypeStructPointer { typeArgs: tArgs } ->
            case ctorLayout.constructorFields of
              Just ctorInfo ->
                let
                  env = Map.fromFoldable (Array.zip ctorInfo.vars tArgs)
                  genericTy = structFieldGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors ctorInfo.vars modNameStr (fromMaybe (TypeVar "") (Array.index fields idx))
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
  TypeStructPointer _, GoConstructor _ ctor typeArgs fields ->
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
          , condition: GoBinOp "&&" (GoBinOp "!=" source (rawGo "nil"))
              (GoBinOp "==" (projection changedIndex) replacement)
          }
      _ -> Nothing
  _, _ -> Nothing

-- The operand has already been translated. A temporary keeps non-variable
-- operands evaluated exactly once, including constant native tag tests.
isTag :: CodegenMetadata -> Ref CodegenState -> String -> Maybe ModuleName -> String -> ExprResult -> ExprResult
isTag metadata codegenStateRef modNameStr mbMod tag resE =
  let
    baseStructName = getStructName modNameStr mbMod tag
    hashStr = hashString baseStructName
    nativeTagTest = case resE.exprType of
      TypeStructValue adtName _ -> map (\adt -> adt.isConstructor baseStructName) (Map.lookup adtName unboxableADTs)
      _ -> Nothing

    isNativePointer = case resE.exprType of
      TypeStructPointer { baseStructName: typedBaseStructName } ->
        typedBaseStructName == baseStructName ||
          ( case Map.lookup baseStructName metadata.pointerAdtLeaves of
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
                case Map.lookup baseStructName metadata.pointerAdtLeaves of
                  Just _ -> "(" <> printGoExpr resE.expr <> " == nil)"
                  Nothing -> "(" <> printGoExpr resE.expr <> " != nil)"
              else case Map.lookup baseStructName metadata.pointerAdtLeaves of
                Just nodeInfo -> "(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".IntVal == " <> hashString nodeInfo.nodeBaseStruct <> " && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".UnsafePtr == nil)"
                Nothing ->
                  if Set.member baseStructName metadata.pointerAdtNodes then
                    "(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> " && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".UnsafePtr != nil)"
                  else if Set.member baseStructName metadata.enumCtors then
                    "(" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeUint32) <> " == " <> hashStr <> ")"
                  else
                    "(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> ")"
        in
          { stmts: resE.stmts, expr: rawGo exprStr, exprType: TypeBool, nextId: resE.nextId }
      _ ->
        let
          tmpVar = "__t_tag_" <> show resE.nextId
          declTmp =
            if isNativePointer || resE.exprType /= TypeValue then
              StmtLeaf (rawGo ("var " <> tmpVar <> " " <> goTypeToStr resE.exprType <> " = " <> printGoExpr resE.expr))
            else
              StmtLeaf (rawGo ("var " <> tmpVar <> " gopurs_runtime.Value = " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType)))

          exprStr = case nativeTagTest of
            Just test -> printGoExpr (test (GoVar tmpVar))
            Nothing ->
              if isNativePointer then
                case Map.lookup baseStructName metadata.pointerAdtLeaves of
                  Just _ -> "(" <> tmpVar <> " == nil)"
                  Nothing -> "(" <> tmpVar <> " != nil)"
              else if resE.exprType /= TypeValue then
                "(uint32(" <> tmpVar <> ") == " <> hashStr <> ")"
              else case Map.lookup baseStructName metadata.pointerAdtLeaves of
                Just nodeInfo -> "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashString nodeInfo.nodeBaseStruct <> " && " <> tmpVar <> ".UnsafePtr == nil)"
                Nothing ->
                  if Set.member baseStructName metadata.pointerAdtNodes then
                    "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashStr <> " && " <> tmpVar <> ".UnsafePtr != nil)"
                  else if Set.member baseStructName metadata.enumCtors then
                    "(" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr (GoVar tmpVar) TypeValue TypeUint32) <> " == " <> hashStr <> ")"
                  else
                    "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashStr <> ")"
        in
          -- A single-constructor native value has a constant test,
          -- but its operand must still be evaluated exactly once.
          { stmts: resE.stmts <> declTmp <> StmtLeaf (rawGo ("_ = " <> tmpVar)), expr: rawGo exprStr, exprType: TypeBool, nextId: resE.nextId + 1 }
