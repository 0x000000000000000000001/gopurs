module Gopurs.Driver.Prepare (PreparedModules, PreparedCache, prepareModules, prepareModulesWithCache, prepareModulesDelta) where

import Prelude

import Data.Array as Array
import Data.List (List)
import Data.Map as Map
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Data.String as Data.String
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Effect.Aff (Aff, throwError)
import Effect.Class (liftEffect)
import Effect.Exception (error)
import Gopurs.AdtMetadata (buildEnumAdtMetadata, buildPointerAdtMetadata)
import Gopurs.ClassMetadata (addClassDataDeclarations, buildClassFields)
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.ConstructorMetadata (buildConstructorTypes, collectElidedConstructors)
import Gopurs.Driver.Cache (SourceModule)
import Gopurs.GlobalTypes (buildGlobalTypes)
import Gopurs.Metrics as Metrics
import Gopurs.Cache.Wire as Wire
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import Node.Process as Process
import Gopurs.Monomorphization (monomorphizeModulesWithDispatch)
import Gopurs.Preparation (runPreparationJobs)
import Gopurs.ReboxMetadata (buildReboxFieldIndex)
import PureScript.Backend.Optimizer.App (coreFnModulesFromOutput)
import PureScript.Backend.Optimizer.CoreFn (Ann, Ident(..), Module(..))
import PureScript.Backend.Optimizer.Monomorphize (transitiveCollectWith)
import PureScript.Backend.Optimizer.Semantics (InlineDirectiveMap)

type PreparedModules =
  { directives :: InlineDirectiveMap
  -- Specialized CoreFn for PBO; metadata below still describes source layouts.
  , modules :: List (Module Ann)
  , mainModules :: Array String
  , metadata :: CodegenMetadata
  , sources :: Array SourceModule
  }

type PreparedCache =
  { prepared :: PreparedModules
  , sourceHashes :: Map.Map String String
  , sourceModules :: List (Module Ann)
  }

prepareModules :: String -> Int -> Int -> Maybe String -> InlineDirectiveMap -> Aff PreparedModules
prepareModules input expectedCount jobs mainModule directives =
  _.prepared <$> prepareModulesWithCache input expectedCount jobs mainModule directives

prepareModulesWithCache :: String -> Int -> Int -> Maybe String -> InlineDirectiveMap -> Aff PreparedCache
prepareModulesWithCache input expectedCount jobs mainModule directives = do
  sortedModules <- Metrics.measure "load TAST + sort" \_ -> coreFnModulesFromOutput input
  when (Array.length (Array.fromFoldable sortedModules) /= expectedCount) $
    throwError (error "TAST decoding did not produce every inventoried module")
  Metrics.measure "prepare + monomorphize" \_ -> do
    let
      sourceModules = Array.fromFoldable sortedModules
      elidedCtors = collectElidedConstructors sourceModules

    -- Freeze source metadata before any binding copies or substitutions. The
    -- constructor/class tables and their rebox index must describe declarations,
    -- not one caller's specialized instantiation. Explicit Any stays explicit.
    let
      ctorTypes = buildConstructorTypes sourceModules
      globalTypes = buildGlobalTypes sourceModules
      classDeclsFields = buildClassFields sourceModules
      reboxFields = buildReboxFieldIndex ctorTypes classDeclsFields
      modulesWithClasses = map addClassDataDeclarations sortedModules

    -- The final per-module rewrite joins the same preparation pool as the
    -- transitive rounds: pure thunks over immutable snapshots, results in
    -- caller order, evaluated only while this Aff runs (so the phase measure
    -- below still covers it).
    modules <- monomorphizeModulesWithDispatch
      (\ast instantiations -> Metrics.measure "transitive specializations" \_ ->
        transitiveCollectWith (runPreparationJobs jobs) ast instantiations)
      (runPreparationJobs jobs)
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
    dump <- liftEffect (Process.lookupEnv "GOPURS_DUMP_PREPARED")
    case dump of
      Nothing -> pure unit
      Just file -> do
        hashes <- liftEffect $ traverse (\(Module m) -> do
          digest <- Wire.hash <$> Wire.write (Module m)
          pure (unwrap m.name <> " " <> digest)) (Array.fromFoldable modules)
        FS.writeTextFile UTF8 file (Data.String.joinWith "\n" hashes <> "\n")
    sourceHashes <- liftEffect $ Map.fromFoldable <$> traverse (\(Module m) -> do
      digest <- Wire.hash <$> Wire.write (Module m)
      pure (Tuple (unwrap m.name) digest)) (Array.fromFoldable sourceModules)
    let prepared =
          { directives, modules, metadata, mainModules: selectMainModules mainModule sourceModules
          , sources: map (\(Module m) -> { name: unwrap m.name, path: m.path, hasForeign: not (Map.isEmpty m.foreign) }) sourceModules
          }
    pure { prepared, sourceHashes, sourceModules: sortedModules }

prepareModulesDelta :: String -> Int -> Int -> Maybe String -> InlineDirectiveMap -> PreparedCache -> Aff (Either PreparedCache PreparedModules)
prepareModulesDelta input expectedCount jobs mainModule directives cache = do
  loaded <- Metrics.measure "load TAST + sort" \_ -> coreFnModulesFromOutput input
  when (Array.length (Array.fromFoldable loaded) /= expectedCount) $
    throwError (error "TAST decoding did not produce every inventoried module")
  hashes <- liftEffect $ Map.fromFoldable <$> traverse (\(Module m) -> do
    digest <- Wire.hash <$> Wire.write (Module m)
    pure (Tuple (unwrap m.name) digest)) (Array.fromFoldable loaded)
  let changed = Array.filter (\(Module m) -> Map.lookup (unwrap m.name) cache.sourceHashes /= Just (fromMaybe "" (Map.lookup (unwrap m.name) hashes))) (Array.fromFoldable loaded)
  if Array.length changed /= 1 then do
    full <- prepareModulesWithCache input expectedCount jobs mainModule directives
    pure (Left full)
  else do
    let changedName = case changed of
          [ Module m ] -> unwrap m.name
          _ -> ""
    let enriched = Array.fromFoldable (map addClassDataDeclarations loaded)
    let previous = Map.fromFoldable (map (\(Module m) -> Tuple (unwrap m.name) (Module m)) (Array.fromFoldable cache.prepared.modules))
    let selected = map (\\(Module m) -> case Map.lookup (unwrap m.name) previous of
          Just old | unwrap m.name /= changedName -> old
          _ -> addClassDataDeclarations (Module m)) (Array.fromFoldable loaded)
    let prepared =
          { directives
          , modules: List.fromFoldable selected
          , metadata: cache.prepared.metadata
          , mainModules: selectMainModules mainModule (Array.fromFoldable loaded)
          , sources: map (\(Module m) -> { name: unwrap m.name, path: m.path, hasForeign: not (Map.isEmpty m.foreign) }) (Array.fromFoldable loaded)
          }
    pure (Right prepared)

selectMainModules :: Maybe String -> Array (Module Ann) -> Array String
selectMainModules requested modules = case requested of
  Just name -> [ name ]
  Nothing -> Array.mapMaybe exportedMain modules
  where
  exportedMain (Module m)
    | Array.elem (Ident "main") m.exports = Just (unwrap m.name)
    | otherwise = Nothing
