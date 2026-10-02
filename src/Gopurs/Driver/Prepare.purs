module Gopurs.Driver.Prepare (PreparedModules, prepareModules) where

import Prelude

import Data.Array as Array
import Data.List (List)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Effect.Aff (Aff)
import Gopurs.AdtMetadata (buildEnumAdtMetadata, buildPointerAdtMetadata)
import Gopurs.ClassMetadata (addClassDataDeclarations, buildClassFields)
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.ConstructorMetadata (buildConstructorTypes, collectElidedConstructors)
import Gopurs.GlobalTypes (buildGlobalTypes)
import Gopurs.Metrics as Metrics
import Gopurs.Monomorphization (monomorphizeModulesWith)
import Gopurs.Preparation (runPreparationJobs)
import Gopurs.ReboxMetadata (buildReboxFieldIndex)
import PureScript.Backend.Optimizer.App (coreFnModulesFromOutput, loadDirectives)
import PureScript.Backend.Optimizer.CoreFn (Ann, Ident(..), Module(..))
import PureScript.Backend.Optimizer.Monomorphize (transitiveCollectWith)
import PureScript.Backend.Optimizer.Semantics (InlineDirectiveMap)

type PreparedModules =
  { directives :: InlineDirectiveMap
  -- Specialized CoreFn for PBO; metadata below still describes source layouts.
  , modules :: List (Module Ann)
  , mainModules :: Array String
  , metadata :: CodegenMetadata
  }

prepareModules :: Int -> Maybe String -> Aff PreparedModules
prepareModules jobs mainModule = do
  sortedModules <- Metrics.measure "load TAST + sort" \_ -> coreFnModulesFromOutput "output"
  Metrics.measure "prepare + monomorphize" \_ -> do
    let
      sourceModules = Array.fromFoldable sortedModules
      elidedCtors = collectElidedConstructors sourceModules

    directives <- loadDirectives

    -- Freeze source metadata before any binding copies or substitutions. The
    -- constructor/class tables and their rebox index must describe declarations,
    -- not one caller's specialized instantiation. Explicit Any stays explicit.
    let
      ctorTypes = buildConstructorTypes sourceModules
      globalTypes = buildGlobalTypes sourceModules
      classDeclsFields = buildClassFields sourceModules
      reboxFields = buildReboxFieldIndex ctorTypes classDeclsFields
      modulesWithClasses = map addClassDataDeclarations sortedModules

    modules <- monomorphizeModulesWith
      (\ast instantiations -> Metrics.measure "transitive specializations" \_ ->
        transitiveCollectWith (runPreparationJobs jobs) ast instantiations)
      globalTypes
      modulesWithClasses

    let
      -- Deliberately use the enriched source list, not `modules`: dictionary
      -- layouts need synthetic class declarations, while enum/elision policy
      -- only sees original ADTs. Specialization does not redefine either ABI.
      { pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves } =
        buildPointerAdtMetadata (Array.fromFoldable modulesWithClasses)
      { enumAdts, enumCtors } = buildEnumAdtMetadata sourceModules
      metadata =
        { elidedCtors
        , ctorTypes
        , globalTypes
        , globalFunctions: Map.empty
        , ffiFunctions: Map.empty
        , classDeclsFields
        , reboxFields
        , pointerAdtPaths
        , pointerAdtNodes
        , pointerAdtLeaves
        , enumAdts
        , enumCtors
        }

    -- Entry selection is also source-based. Generated bindings do not export
    -- additional mains; an explicit request keeps its original CLI semantics.
    pure { directives, modules, metadata, mainModules: selectMainModules mainModule sourceModules }

selectMainModules :: Maybe String -> Array (Module Ann) -> Array String
selectMainModules requested modules = case requested of
  Just name -> [ name ]
  Nothing -> Array.mapMaybe exportedMain modules
  where
  exportedMain (Module m)
    | Array.elem (Ident "main") m.exports = Just (unwrap m.name)
    | otherwise = Nothing
