module Gopurs.ConstructorMetadata
  ( ConstructorTypes
  , buildConstructorTypes
  , collectElidedConstructors
  ) where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Gopurs.CodeGen (getStructName)
import PureScript.Backend.Optimizer.CoreFn (Ann, DataConstructor, DataDecl, ExprType, Module(..))

type ConstructorTypes = Map String { vars :: Array String, fields :: Array ExprType }

-- Read the original TAST dataDecls, before adding synthetic class declarations.
buildConstructorTypes :: Array (Module Ann) -> ConstructorTypes
buildConstructorTypes = Array.foldl addModuleConstructorTypes Map.empty

addModuleConstructorTypes :: ConstructorTypes -> Module Ann -> ConstructorTypes
addModuleConstructorTypes types (Module mod) =
  let
    moduleKey = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
  in
    Array.foldl (addDataDeclarationTypes moduleKey) types mod.dataDecls

addDataDeclarationTypes :: String -> ConstructorTypes -> DataDecl -> ConstructorTypes
addDataDeclarationTypes moduleKey types declaration =
  Array.foldl (addConstructorType moduleKey declaration.vars) types declaration.constructors

addConstructorType :: String -> Array String -> ConstructorTypes -> DataConstructor -> ConstructorTypes
addConstructorType moduleKey vars types constructor =
  Map.insert (moduleKey <> "." <> constructor.name) { vars, fields: constructor.fields } types

-- A single constructor with a single field can omit its wrapper in Go.
-- As with the type table, only inspect the original TAST dataDecls.
collectElidedConstructors :: Array (Module Ann) -> Set String
collectElidedConstructors = Array.foldl addModuleElidedConstructors Set.empty

addModuleElidedConstructors :: Set String -> Module Ann -> Set String
addModuleElidedConstructors constructors (Module mod) =
  Array.foldl (addElidedConstructor (unwrap mod.name)) constructors mod.dataDecls

addElidedConstructor :: String -> Set String -> DataDecl -> Set String
addElidedConstructor moduleName constructors declaration =
  case declaration.constructors of
    [ constructor ] | Array.length constructor.fields == 1 ->
      Set.insert (elidedConstructorName moduleName constructor.name) constructors
    _ -> constructors

elidedConstructorName :: String -> String -> String
elidedConstructorName moduleName constructorName =
  let
    -- Preserve the existing struct naming convention, including module dots.
    structName = getStructName moduleName Nothing constructorName
  in
    "Constructor_" <> String.drop 5 structName
