module Gopurs.ConstructorLayout
  ( ConstructorIdentity
  , ConstructorFields
  , ConstructorLayout
  , ConstructionForm(..)
  , PreparedConstructor
  , layout
  , typeArguments
  , prepare
  , fieldType
  ) where

import Prelude
import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.GoAst (GoType(..), sanitizeName, structPointer)
import Gopurs.GoTypes (exprTypeToGoType, exprTypeToGenericGoType, instantiateGenericGoType)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

type ConstructorIdentity =
  { moduleName :: String
  , structName :: String
  , baseStructName :: String
  , key :: String
  , fullName :: String
  }

type ConstructorFields = { vars :: Array String, fields :: Array ExprType }

type ConstructorLayout =
  { identity :: ConstructorIdentity
  , constructorFields :: Maybe ConstructorFields
  , fields :: ConstructorFields
  }

-- Construction can recover the defining module from a native result type.
-- Field access supplies Nothing and keeps the explicitly qualified constructor.
layout :: CodegenMetadata -> String -> String -> Maybe GoType -> ConstructorLayout
layout metadata fallbackModule name expectedType =
  let
    moduleName = String.replaceAll (Pattern ".") (Replacement "_") case expectedType of
      Just (TypeStructPointer { fullName: typeName }) ->
        let parts = String.split (Pattern ".") typeName
        in String.joinWith "." (Array.take (Array.length parts - 1) parts)
      _ -> fallbackModule
    key = moduleName <> "." <> name
    fullName = case expectedType of
      Just (TypeStructPointer pointer) -> pointer.fullName
      _ -> if key == "Test_RBTree.E" then "Test.RBTree.Tree" else key
    identity =
      { moduleName, key, fullName
      , structName: "Constructor_" <> moduleName <> "_" <> sanitizeName name
      , baseStructName: "Data_" <> moduleName <> "_" <> sanitizeName name
      }
    constructorFields = Map.lookup key metadata.ctorTypes
    fields = case constructorFields of
      Just info -> info
      Nothing -> case Map.lookup fullName metadata.classDeclsFields of
        Just info -> { vars: info.vars, fields: map _."type" info.fields }
        Nothing -> { vars: [], fields: [] }
  in
    { identity, constructorFields, fields }

-- Metadata fixes an ADT's arity. The caller specifies the fallback because
-- field instantiation and pointer construction have different information.
typeArguments :: CodegenMetadata -> String -> Maybe Int -> Array String -> ExprType -> Array GoType
typeArguments metadata modNameStr fallbackArity vars = case _ of
  ADT name _ args ->
    let
      mapped = map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) args
      arity = fromMaybe
        (case Map.lookup name metadata.pointerAdtPaths of
          Just info -> info.arity
          Nothing -> Array.length mapped)
        fallbackArity
    in
      -- An opaque type introduced by unsafeCoerce can hide constructor
      -- parameters. Its shorter argument list cannot instantiate this layout.
      if Array.length mapped < arity then Array.replicate arity TypeValue
      else Array.take arity mapped
  _ -> map (const TypeValue) vars

data ConstructionForm = Definition | Saturated

type PreparedConstructor =
  { layout :: ConstructorLayout
  , fieldTypeArgs :: Array GoType
  , typeArgs :: Array GoType
  , pointerType :: GoType
  , leafPointerType :: Maybe GoType
  , adtFullName :: Maybe String
  }

prepare :: ConstructionForm -> CodegenMetadata -> String -> String -> String -> ExprType -> PreparedConstructor
prepare form metadata modNameStr fallbackModule name ctorType =
  let
    expectedType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr ctorType
    ctorLayout = layout metadata fallbackModule name (Just expectedType)
    { identity, fields } = ctorLayout
    fieldTypeArgs = case form, ctorType of
      Definition, TypeApp _ _ ->
        let
          unwrap (TypeApp fn args) acc = unwrap fn (args <> acc)
          unwrap other acc = Tuple other acc
        in
          case unwrap ctorType [] of
            Tuple (ADT _ _ args) applied ->
              Array.take (Array.length fields.vars)
                (map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) (args <> applied))
            _ -> map (const TypeValue) fields.vars
      _, _ -> typeArguments metadata modNameStr (Just (Array.length fields.vars)) fields.vars ctorType
    typeArgs = case form of
      Definition -> fieldTypeArgs
      Saturated -> typeArguments metadata modNameStr (Just (Array.length fields.vars)) fields.vars ctorType
    pointerType = structPointer identity typeArgs
    leafPointerType = map
      (\leaf ->
        let
          -- Definitions historically resolve leaf workers in the current
          -- module; saturated constructors use the resolved defining module.
          leafModule = case form of
            Definition -> modNameStr
            Saturated -> identity.moduleName
        in
          structPointer
            { baseStructName: leaf.nodeBaseStruct
            , fullName: identity.fullName
            , structName: "Constructor_" <> leafModule <> "_" <> sanitizeName leaf.nodeCtor
            }
            typeArgs)
      (Map.lookup identity.baseStructName metadata.pointerAdtLeaves)
    adtFullName = case ctorType of
      ADT fullName _ _ -> Just fullName
      _ -> Nothing
  in
    { layout: ctorLayout, fieldTypeArgs, typeArgs, pointerType, leafPointerType, adtFullName }

fieldType :: CodegenMetadata -> String -> ConstructorFields -> Array GoType -> Int -> GoType
fieldType metadata modNameStr info args index = case Array.index info.fields index of
  Just ty ->
    let
      genericType = exprTypeToGenericGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors info.vars modNameStr ty
      env = Map.fromFoldable (Array.zip info.vars args)
    in
      instantiateGenericGoType env genericType
  Nothing -> TypeValue
