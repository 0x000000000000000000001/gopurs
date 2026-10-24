module Gopurs.Driver.Prepare.ProjectionCapsule (Fragment, create, open) where

import Prelude
import Data.Array as Array
import Data.Either (Either)
import Data.Foldable (foldM)
import Data.Lazy (Lazy)
import Data.List (List)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set (Set)
import Data.Set as Set
import Gopurs.Cache.LookupTable as LookupTable
import Gopurs.CodegenState (SharedMetadataRow)
import Gopurs.Driver.Prepare.Projection (Projection)
import Gopurs.Driver.Prepare.Projection as Projection
import PureScript.Backend.Optimizer.CoreFn (Ann, ExprType, Module)

-- This is a distinct wire record, never a Projection decoded with a partial
-- globalTypes Map. Its object SHA is referenced by the accepted source proof.
type Fragment =
  { schema :: Int, ordinal :: Int, count :: Int
  , modules :: List (Module Ann), interfaces :: Map String String
  , mainModules :: Array String, metadata :: { | SharedMetadataRow }
  , coverage :: Set String, values :: Map String ExprType
  }

create :: Set String -> Projection -> Array Fragment
create coverage projection = map fragment parts
  where
  full = LookupTable.completeOrFail projection.metadata.globalTypes
  parts = Projection.splitMaps (projection
    { metadata = projection.metadata { globalTypes = Map.filterKeys (flip Set.member coverage) full } })
  keys = Set.toUnfoldable coverage :: Array String
  -- Construct the exact row shape explicitly. A field insertion is not a
  -- same-shape record update in the native backend.
  fragment part =
    { schema: 1, ordinal: part.ordinal, count: part.count
    , modules: part.value.modules, interfaces: part.value.interfaces
    , mainModules: part.value.mainModules
    , metadata: { classDeclsFields: part.value.metadata.classDeclsFields
        , ctorTypes: part.value.metadata.ctorTypes
        , elidedCtors: part.value.metadata.elidedCtors
        , enumAdts: part.value.metadata.enumAdts
        , enumCtors: part.value.metadata.enumCtors
        , ffiFunctions: part.value.metadata.ffiFunctions
        , globalFunctions: part.value.metadata.globalFunctions
        , pointerAdtLeaves: part.value.metadata.pointerAdtLeaves
        , pointerAdtNodes: part.value.metadata.pointerAdtNodes
        , pointerAdtPaths: part.value.metadata.pointerAdtPaths
        , reboxFields: part.value.metadata.reboxFields
        }
    , values: part.value.metadata.globalTypes
    , coverage: Set.fromFoldable (Array.slice
        (Array.length keys * part.ordinal / part.count)
        (Array.length keys * (part.ordinal + 1) / part.count) keys)
    }

-- Both values and verified absences come exclusively from create's complete
-- source Map. Coverage can be incomplete; the shared Lazy answers every other
-- key from the original full fragments after their hashes have been checked.
open :: Array Fragment -> Lazy (Either String (Map String ExprType)) -> Maybe { projection :: Projection, coverage :: Set String }
open fragments delayed = do
  if Array.all (\fragment -> fragment.schema == 1) fragments then Just unit else Nothing
  coverage <- foldM mergeCoverage Set.empty fragments
  mapped <- Projection.mergeMaps (map toMapFragment fragments)
  globals <- LookupTable.covered coverage mapped.metadata.globalTypes delayed
  pure
    { projection: mapped { metadata = mapped.metadata { globalTypes = globals } }
    , coverage
    }
  where
  mergeCoverage keys fragment =
    let combined = Set.union keys fragment.coverage
    in if Set.size combined == Set.size keys + Set.size fragment.coverage then Just combined else Nothing
  toMapFragment fragment =
    { ordinal: fragment.ordinal, count: fragment.count
    , value:
        { modules: fragment.modules, interfaces: fragment.interfaces
        , mainModules: fragment.mainModules
        , metadata: { classDeclsFields: fragment.metadata.classDeclsFields
            , ctorTypes: fragment.metadata.ctorTypes
            , elidedCtors: fragment.metadata.elidedCtors
            , enumAdts: fragment.metadata.enumAdts
            , enumCtors: fragment.metadata.enumCtors
            , ffiFunctions: fragment.metadata.ffiFunctions
            , globalFunctions: fragment.metadata.globalFunctions
            , pointerAdtLeaves: fragment.metadata.pointerAdtLeaves
            , pointerAdtNodes: fragment.metadata.pointerAdtNodes
            , pointerAdtPaths: fragment.metadata.pointerAdtPaths
            , reboxFields: fragment.metadata.reboxFields
            , globalTypes: fragment.values
            }
        }
    }
