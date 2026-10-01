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
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.ConstructorMetadata as ConstructorMetadata
import Gopurs.GoAst (GoType(..), constructorNames, structPointer)
import Gopurs.GoTypes (appliedAdt, exprTypeToGoType, exprTypeToGenericGoType, instantiateGenericGoType)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

type ConstructorIdentity =
  { moduleName :: String
  , structName :: String
  , baseStructName :: String
  , key :: String
  , fullName :: String
  }

type ConstructorFields = ConstructorMetadata.ConstructorFields

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
    identity = constructorIdentity fallbackModule name expectedType
    constructorFields = Map.lookup identity.key metadata.ctorTypes
    -- Original constructor fields take precedence over dictionary metadata.
    fields = case constructorFields of
      Just info -> info
      Nothing -> case Map.lookup identity.fullName metadata.classDeclsFields of
        Just info -> { vars: info.vars, fields: map _."type" info.fields }
        Nothing -> { vars: [], fields: [] }
  in
    { identity, constructorFields, fields }

constructorIdentity :: String -> String -> Maybe GoType -> ConstructorIdentity
constructorIdentity fallbackModule name expectedType =
  let
    moduleName = String.replaceAll (Pattern ".") (Replacement "_") case expectedType of
      Just (TypeStructPointer { fullName: typeName }) ->
        let parts = String.split (Pattern ".") typeName
        in String.joinWith "." (Array.take (Array.length parts - 1) parts)
      _ -> fallbackModule
    key = moduleName <> "." <> name
    fullName = case expectedType of
      Just (TypeStructPointer pointer) -> pointer.fullName
      -- Historical type hint for the empty RBTree constructor.
      _ -> if key == "Test_RBTree.E" then "Test.RBTree.Tree" else key
    names = constructorNames moduleName name
  in
    { moduleName, key, fullName, structName: names.structName, baseStructName: names.baseStructName }

-- The layout's explicit arity wins over pointer metadata, then the annotation's
-- own argument count. Only plain ADTs use those arguments; other annotations
-- erase the supplied variables. Definition TypeApps are handled separately.
typeArguments :: CodegenMetadata -> String -> Maybe Int -> Array String -> ExprType -> Array GoType
typeArguments metadata modNameStr layoutArity vars = case _ of
  ADT name _ args ->
    let
      mapped = map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) args
      arity = fromMaybe
        (case Map.lookup name metadata.pointerAdtPaths of
          Just info -> info.arity
          Nothing -> Array.length mapped)
        layoutArity
    in
      -- An opaque type introduced by unsafeCoerce can hide constructor
      -- parameters. Its shorter argument list cannot instantiate this layout.
      if Array.length mapped < arity then Array.replicate arity TypeValue
      else Array.take arity mapped
  _ -> map (const TypeValue) vars

data ConstructionForm = Definition | Saturated

constructionTypeArguments :: ConstructionForm -> CodegenMetadata -> String -> ConstructorFields -> ExprType -> Array GoType
constructionTypeArguments form metadata modNameStr fields ctorType = case form, ctorType of
  -- Definitions retain the available prefix of a TypeApp, without padding.
  -- Saturated TypeApps use the ordinary erased fallback in typeArguments.
  Definition, TypeApp _ _ -> case appliedAdt ctorType of
    Just adt -> Array.take (Array.length fields.vars)
      (map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) adt.args)
    Nothing -> map (const TypeValue) fields.vars
  _, _ -> typeArguments metadata modNameStr (Just (Array.length fields.vars)) fields.vars ctorType

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
    typeArgs = constructionTypeArguments form metadata modNameStr fields ctorType
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
            , structName: (constructorNames leafModule leaf.nodeCtor).structName
            }
            typeArgs)
      (Map.lookup identity.baseStructName metadata.pointerAdtLeaves)
    adtFullName = case ctorType of
      ADT fullName _ _ -> Just fullName
      _ -> Nothing
  in
    { layout: ctorLayout, fieldTypeArgs: typeArgs, typeArgs, pointerType, leafPointerType, adtFullName }

fieldType :: CodegenMetadata -> String -> ConstructorFields -> Array GoType -> Int -> GoType
fieldType metadata modNameStr info args index = case Array.index info.fields index of
  Just ty ->
    let
      genericType = exprTypeToGenericGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors info.vars modNameStr ty
      env = Map.fromFoldable (Array.zip info.vars args)
    in
      instantiateGenericGoType env genericType
  Nothing -> TypeValue
