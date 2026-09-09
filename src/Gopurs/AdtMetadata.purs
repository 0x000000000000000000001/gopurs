module Gopurs.AdtMetadata
  ( PointerAdtMetadata
  , EnumAdtMetadata
  , buildPointerAdtMetadata
  , buildEnumAdtMetadata
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
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (sanitizeName)
import PureScript.Backend.Optimizer.CoreFn (Ann, DataConstructor, DataDecl, Module(..))

type PointerAdtMetadata =
  { pointerAdtPaths :: Map String { ctorName :: String, arity :: Int }
  , pointerAdtNodes :: Set String
  , pointerAdtLeaves :: Map String { nodeBaseStruct :: String, nodeCtor :: String }
  }

type EnumAdtMetadata =
  { enumAdts :: Set String
  , enumCtors :: Set String
  }

type PointerAdtInfo =
  { adtPath :: String
  , nodeCtor :: String
  , nodeBaseStruct :: String
  , leafBaseStruct :: Maybe String
  , arity :: Int
  }

type EnumAdtInfo =
  { adtPath :: String
  , constructorStructs :: Array String
  }

-- Inspect dataDecls after synthetic class declarations have been added.
buildPointerAdtMetadata :: Array (Module Ann) -> PointerAdtMetadata
buildPointerAdtMetadata modules =
  let
    declarations = Array.concatMap modulePointerAdts modules
  in
    { pointerAdtPaths: Map.fromFoldable (map (\info -> Tuple info.adtPath { ctorName: info.nodeCtor, arity: info.arity }) declarations)
    , pointerAdtNodes: Set.fromFoldable (map _.nodeBaseStruct declarations)
    , pointerAdtLeaves: Map.fromFoldable (Array.mapMaybe pointerLeafEntry declarations)
    }

modulePointerAdts :: Module Ann -> Array PointerAdtInfo
modulePointerAdts (Module mod) =
  Array.mapMaybe (pointerAdtInfo (unwrap mod.name)) mod.dataDecls

-- Exactly one constructor carries fields; any number of nullary constructors is allowed.
pointerAdtInfo :: String -> DataDecl -> Maybe PointerAdtInfo
pointerAdtInfo moduleName declaration =
  case Array.filter (\constructor -> not (Array.null constructor.fields)) declaration.constructors of
    [ node ] -> Just
      { adtPath: moduleName <> "." <> declaration.name
      , nodeCtor: node.name
      , nodeBaseStruct: constructorStructName moduleName node.name
      , leafBaseStruct: uniqueLeafStructName moduleName declaration.constructors
      , arity: Array.length declaration.vars
      }
    _ -> Nothing

uniqueLeafStructName :: String -> Array DataConstructor -> Maybe String
uniqueLeafStructName moduleName constructors =
  case Array.filter (\constructor -> Array.null constructor.fields) constructors of
    [ leaf ] | leaf.name /= "" -> Just (constructorStructName moduleName leaf.name)
    _ -> Nothing

pointerLeafEntry :: PointerAdtInfo -> Maybe (Tuple String { nodeBaseStruct :: String, nodeCtor :: String })
pointerLeafEntry info =
  map (\name -> Tuple name { nodeBaseStruct: info.nodeBaseStruct, nodeCtor: info.nodeCtor }) info.leafBaseStruct

-- Use the same enriched dataDecls as the pointer metadata.
buildEnumAdtMetadata :: Array (Module Ann) -> EnumAdtMetadata
buildEnumAdtMetadata modules =
  let
    declarations = Array.concatMap moduleEnumAdts modules
  in
    { enumAdts: Set.fromFoldable (map _.adtPath declarations)
    , enumCtors: Set.fromFoldable (Array.concatMap _.constructorStructs declarations)
    }

moduleEnumAdts :: Module Ann -> Array EnumAdtInfo
moduleEnumAdts (Module mod) =
  Array.mapMaybe (enumAdtInfo (unwrap mod.name)) mod.dataDecls

enumAdtInfo :: String -> DataDecl -> Maybe EnumAdtInfo
enumAdtInfo moduleName declaration =
  if not (Array.null declaration.constructors) && Array.all (\constructor -> Array.null constructor.fields) declaration.constructors then
    Just
      { adtPath: moduleName <> "." <> declaration.name
      , constructorStructs: map (\constructor -> constructorStructName moduleName constructor.name) declaration.constructors
      }
  else Nothing

constructorStructName :: String -> String -> String
constructorStructName moduleName constructorName =
  let
    packageName = String.replaceAll (Pattern ".") (Replacement "_") moduleName
  in
    "Data_" <> packageName <> "_" <> sanitizeName constructorName
