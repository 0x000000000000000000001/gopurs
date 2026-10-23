module Gopurs.Driver.Prepare.Projection (Projection, Fragment, fragmentCount, split, merge) where

import Prelude
import Data.Array as Array
import Data.Foldable (foldM)
import Data.List (List)
import Data.List as List
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (Tuple)
import Gopurs.CodegenState (CodegenMetadata)
import PureScript.Backend.Optimizer.CoreFn (Ann, Module)

-- Only inert headers are stored here. Exact source types and layouts remain in
-- metadata; every fresh conversion reloads and prepares the original TAST.
type Projection =
  { modules :: List (Module Ann)
  , metadata :: CodegenMetadata
  , interfaces :: Map String String
  , mainModules :: Array String
  }

type Fragment = { ordinal :: Int, count :: Int, value :: Projection }

fragmentCount :: Int -> Int
fragmentCount modules = max 1 (min 8 modules)

-- Contiguous partitions preserve module and main selection order. Each map/set
-- partitions its ordered entries independently, without duplicating any key.
split :: Projection -> Array Fragment
split projection = map fragment (Array.range 0 (count - 1))
  where
  modules = Array.fromFoldable projection.modules
  count = fragmentCount (Array.length modules)
  slice :: forall a. Int -> Array a -> Array a
  slice index values = Array.slice (Array.length values * index / count) (Array.length values * (index + 1) / count) values
  table :: forall a. Int -> Map String a -> Map String a
  table index values = Map.fromFoldable (slice index (Map.toUnfoldable values :: Array (Tuple String a)))
  keys :: Int -> Set String -> Set String
  keys index values = Set.fromFoldable (slice index (Set.toUnfoldable values :: Array String))
  fragment ordinal =
    { ordinal, count
    , value:
        { modules: List.fromFoldable (slice ordinal modules)
        , interfaces: table ordinal projection.interfaces
        , mainModules: slice ordinal projection.mainModules
        , metadata:
            { elidedCtors: keys ordinal projection.metadata.elidedCtors
            , ctorTypes: table ordinal projection.metadata.ctorTypes
            , pointerAdtPaths: table ordinal projection.metadata.pointerAdtPaths
            , pointerAdtNodes: keys ordinal projection.metadata.pointerAdtNodes
            , pointerAdtLeaves: table ordinal projection.metadata.pointerAdtLeaves
            , enumAdts: keys ordinal projection.metadata.enumAdts
            , enumCtors: keys ordinal projection.metadata.enumCtors
            , globalTypes: table ordinal projection.metadata.globalTypes
            , globalFunctions: table ordinal projection.metadata.globalFunctions
            , ffiFunctions: table ordinal projection.metadata.ffiFunctions
            , classDeclsFields: table ordinal projection.metadata.classDeclsFields
            , reboxFields: table ordinal projection.metadata.reboxFields
            }
        }
    }

mergeTable :: forall a. Map String a -> Map String a -> Maybe (Map String a)
mergeTable left right =
  let merged = Map.union left right
  in if Map.size merged == Map.size left + Map.size right then Just merged else Nothing

mergeKeys :: Set String -> Set String -> Maybe (Set String)
mergeKeys left right =
  let merged = Set.union left right
  in if Set.size merged == Set.size left + Set.size right then Just merged else Nothing

mergeMetadata :: CodegenMetadata -> CodegenMetadata -> Maybe CodegenMetadata
mergeMetadata left right = do
  elidedCtors <- mergeKeys left.elidedCtors right.elidedCtors
  ctorTypes <- mergeTable left.ctorTypes right.ctorTypes
  pointerAdtPaths <- mergeTable left.pointerAdtPaths right.pointerAdtPaths
  pointerAdtNodes <- mergeKeys left.pointerAdtNodes right.pointerAdtNodes
  pointerAdtLeaves <- mergeTable left.pointerAdtLeaves right.pointerAdtLeaves
  enumAdts <- mergeKeys left.enumAdts right.enumAdts
  enumCtors <- mergeKeys left.enumCtors right.enumCtors
  globalTypes <- mergeTable left.globalTypes right.globalTypes
  globalFunctions <- mergeTable left.globalFunctions right.globalFunctions
  ffiFunctions <- mergeTable left.ffiFunctions right.ffiFunctions
  classDeclsFields <- mergeTable left.classDeclsFields right.classDeclsFields
  reboxFields <- mergeTable left.reboxFields right.reboxFields
  pure { elidedCtors, ctorTypes, pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves
       , enumAdts, enumCtors, globalTypes, globalFunctions, ffiFunctions, classDeclsFields, reboxFields }

merge :: Array Fragment -> Maybe Projection
merge fragments = do
  first <- Array.head fragments
  rest <- Array.tail fragments
  let count = Array.length fragments
  if count < 1 || count > 8 || Array.mapWithIndex (\index fragment -> fragment.ordinal == index && fragment.count == count) fragments /= Array.replicate count true
    then Nothing
    else foldM append first.value rest
  where
  append left fragment = do
    interfaces <- mergeTable left.interfaces fragment.value.interfaces
    metadata <- mergeMetadata left.metadata fragment.value.metadata
    pure { modules: left.modules <> fragment.value.modules
         , interfaces, metadata, mainModules: left.mainModules <> fragment.value.mainModules }
