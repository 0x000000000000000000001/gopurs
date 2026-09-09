module Gopurs.ClassMetadata
  ( ClassFields
  , buildClassFields
  , addClassDataDeclarations
  ) where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (fromMaybe)
import Data.Newtype (unwrap)
import Data.Tuple (Tuple(..), fst, snd)
import PureScript.Backend.Optimizer.CoreFn (Ann, ClassDecl, DataDecl, ExprType(..), Module(..))

type ClassFields = Map String
  { vars :: Array String
  , fields :: Array { name :: String, "type" :: ExprType }
  }

buildClassFields :: Array (Module Ann) -> ClassFields
buildClassFields = Array.foldl addModuleClassFields Map.empty

addModuleClassFields :: ClassFields -> Module Ann -> ClassFields
addModuleClassFields classes (Module mod) =
  Array.foldl (addClassFields (unwrap mod.name)) classes mod.classDecls

addClassFields :: String -> ClassFields -> ClassDecl -> ClassFields
addClassFields moduleName classes declaration =
  let
    fields = map (\(Tuple name ty) -> { name, "type": ty }) (sortedClassFields declaration)
  in
    Map.insert (moduleName <> "." <> declaration.name) { vars: declaration.vars, fields } classes

-- Keep the original ADTs first, then append one dictionary declaration per class.
addClassDataDeclarations :: Module Ann -> Module Ann
addClassDataDeclarations (Module mod) =
  let
    declarations = map classDataDeclaration mod.classDecls
  in
    Module (mod { dataDecls = mod.dataDecls <> declarations })

classDataDeclaration :: ClassDecl -> DataDecl
classDataDeclaration declaration =
  { name: declaration.name
  , vars: declaration.vars
  , constructors:
      [ { name: declaration.name
        , fields: map snd (sortedClassFields declaration)
        }
      ]
  }

-- The named field table and the synthetic constructor must share the same order.
sortedClassFields :: ClassDecl -> Array (Tuple String ExprType)
sortedClassFields declaration =
  let
    superclassFields = Array.mapWithIndex superclassField declaration.superclasses
  in
    Array.sortBy (comparing fst) (superclassFields <> declaration.methods)

superclassField :: Int -> Tuple (Array String) (Array ExprType) -> Tuple String ExprType
superclassField index (Tuple path _) =
  -- Preserve the existing superclass name and Any representation.
  Tuple (fromMaybe "" (Array.last path) <> show index) Any
