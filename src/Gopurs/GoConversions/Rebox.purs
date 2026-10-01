module Gopurs.GoConversions.Rebox
  ( request
  , generate
  ) where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.String as String
import Data.Tuple (Tuple(..), fst)
import Effect (Effect)
import Effect.Console as Console
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CodegenState (CodegenMetadata, CodegenState)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), StructPointer, rawGo)
import Gopurs.GoTypes as GoTypes
import Gopurs.Printer (printGoExpr)
import Gopurs.ReboxMetadata (ReboxFields)
import PureScript.Backend.Optimizer.FfiSupport (hashString)

type CoerceField = GoExpr -> GoType -> GoType -> GoExpr

functionName :: String -> StructPointer -> StructPointer -> String
functionName moduleName source target =
  "Rebox_" <> moduleName <> "_" <> hashString source.fullPath <> "_" <> hashString target.fullPath

-- The converter has already established equal runtime constructor identities.
-- Requesting a helper records the directed pair; it does not render fields yet.
request :: Ref CodegenState -> String -> GoExpr -> StructPointer -> StructPointer -> GoExpr
request stateRef moduleName expr source target = unsafePerformEffect do
  registerPair stateRef (TypeStructPointer source) (TypeStructPointer target)
  pure (GoCall (GoVar (functionName moduleName source target)) [ expr ])

registerPair :: Ref CodegenState -> GoType -> GoType -> Effect Unit
registerPair stateRef source target = do
  state <- Ref.read stateRef
  let pairs = state.reboxPairs
  if Set.member (Tuple source target) pairs then pure unit
  else Ref.modify_ (\s -> s { reboxPairs = Set.insert (Tuple source target) pairs }) stateRef

-- Keep the existing diagnostic and omission when no field metadata is known.
findFields :: CodegenMetadata -> String -> Maybe ReboxFields
findFields metadata baseStructName =
  case Map.lookup baseStructName metadata.reboxFields of
    Just info -> Just info
    Nothing -> unsafePerformEffect do
      Console.log ("ERROR: Rebox missing! b1=" <> baseStructName <> " keysCtor: " <> String.joinWith ", " (map fst (Map.toUnfoldable metadata.ctorTypes :: Array (Tuple String _))))
      pure Nothing

-- Instantiating a field in both environments chooses its conversion. Calling
-- coerceField may request another helper, including this same recursive pair.
convertFields :: CoerceField -> CodegenMetadata -> String -> ReboxFields -> Array GoType -> Array GoType -> Array (Tuple String GoExpr)
convertFields coerceField metadata moduleName info sourceArgs targetArgs =
  let
    sourceEnv = Map.fromFoldable (Array.zip info.vars sourceArgs)
    targetEnv = Map.fromFoldable (Array.zip info.vars targetArgs)
  in
    Array.mapWithIndex
      (\index fieldType ->
        let
          fieldName = "V" <> show index
          genericType = GoTypes.structFieldGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors info.vars moduleName fieldType
          sourceType = GoTypes.instantiateGenericGoType sourceEnv genericType
          targetType = GoTypes.instantiateGenericGoType targetEnv genericType
        in
          Tuple fieldName (coerceField (GoStructAccess (GoVar "in") fieldName) sourceType targetType))
      info.fields

renderBody :: String -> Array (Tuple String GoExpr) -> String
renderBody targetPath fields =
  let
    -- Preserve the printed-identity check: phantom parameters can share the
    -- same immutable layout. Rc is not read by these conversions.
    allIdentity = Array.all (\(Tuple fieldName expr) -> printGoExpr expr == "in." <> fieldName) fields
  in
    if allIdentity then
      "\tif in == nil { return nil }\n\treturn (*" <> targetPath <> ")(unsafe.Pointer(in))"
    else
      let assignments = String.joinWith "\n" (map (\(Tuple fieldName expr) -> "\t\tout." <> fieldName <> " = " <> printGoExpr expr) fields)
      in "\tif in == nil { return nil }\n\tout := &" <> targetPath <> "{}\n" <> assignments <> "\n\treturn out"

renderFunction :: CoerceField -> CodegenMetadata -> String -> Map String GoDecl -> Tuple GoType GoType -> Maybe (Tuple String GoDecl)
renderFunction coerceField metadata moduleName generated (Tuple sourceType targetType) =
  case sourceType, targetType of
    TypeStructPointer source, TypeStructPointer target | source.baseStructName == target.baseStructName ->
      let name = functionName moduleName source target
      in
        if Map.member name generated then Nothing
        else case findFields metadata source.baseStructName of
          Just info ->
            let
              fields = convertFields coerceField metadata moduleName info source.typeArgs target.typeArgs
              declaration = GoFunctionDecl
                { name, params: [ Tuple "in" sourceType ], result: targetType
                , body: rawGo (renderBody target.fullPath fields)
                }
            in
              Just (Tuple name declaration)
          Nothing -> Nothing
    _, _ -> Nothing

-- Each wave reads the ordered set of requested pairs. Field conversions can
-- extend it; keep rendering until a wave adds no declaration. Deduplicate and
-- return by helper name, preserving the existing Map order and collision rules.
generate :: CodegenMetadata -> Ref CodegenState -> String -> CoerceField -> Effect (Array GoDecl)
generate metadata stateRef moduleName coerceField = loop Map.empty
  where
  loop generated = do
    state <- Ref.read stateRef
    let
      newFunctions = Map.fromFoldable
        (Array.mapMaybe (renderFunction coerceField metadata moduleName generated) (Array.fromFoldable state.reboxPairs))
      nextGenerated = Map.union generated newFunctions
    if Map.isEmpty newFunctions then
      pure (Array.fromFoldable (Map.values nextGenerated))
    else
      loop nextGenerated
