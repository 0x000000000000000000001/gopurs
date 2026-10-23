module Gopurs.Driver.Prepare (PreparedModules, PreparedCache, prepareModules, prepareModulesWithCache, prepareModulesDelta, prepareLoadedModules, sourceMetadata, selectMainModules) where

import Prelude

import Data.Array as Array
import Data.List (List)
import Data.Map as Map
import Data.Set (Set)
import Data.Set as Set
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Data.String as Data.String
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Effect.Aff (Aff, throwError)
import Effect.Class (liftEffect)
import Effect.Exception (error)
import Effect.Ref as Ref
import Gopurs.AdtMetadata (buildEnumAdtMetadata, buildPointerAdtMetadata)
import Gopurs.ClassMetadata (addClassDataDeclarations, buildClassFields)
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.ConstructorMetadata (buildConstructorTypes, collectElidedConstructors)
import Gopurs.Driver.Cache (SourceModule)
import Gopurs.Driver.Prepare.Proof as Proof
import Gopurs.GlobalTypes (buildGlobalTypes)
import Gopurs.Metrics as Metrics
import Gopurs.Cache.Wire as Wire
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import Node.Process as Process
import Gopurs.Monomorphization (MonoCache, monomorphizeModulesWithCache)
import Gopurs.Preparation (runPreparationJobs)
import Gopurs.ReboxMetadata (buildReboxFieldIndex)
import PureScript.Backend.Optimizer.App (coreFnModulesFromOutput)
import PureScript.Backend.Optimizer.CoreFn (Ann, Ident(..), Module(..))
import PureScript.Backend.Optimizer.Monomorphize (transitiveCollectWithReads)
import PureScript.Backend.Optimizer.Semantics (InlineDirectiveMap)

type PreparedModules =
  { directives :: InlineDirectiveMap
  -- Specialized CoreFn for PBO; metadata below still describes source layouts.
  , modules :: List (Module Ann)
  , mainModules :: Array String
  , metadata :: CodegenMetadata
  -- A committed preparation proof covers every source input of sourceMetadata.
  -- Consumers still bind a reused metadata fingerprint to the committed memos.
  , metadataUnchanged :: Boolean
  , sources :: Array SourceModule
  }

type PreparedCache =
  { prepared :: PreparedModules
  , sourceHashes :: Map.Map String String
  , sourceModules :: List (Module Ann)
  , sourceInterfaces :: Map.Map String String
  , mono :: MonoCache
  , reusable :: Set String
  , proof :: Maybe Proof.State
  , materialize :: Maybe (String -> Aff (Maybe (Module Ann)))
  }

sourceFileHash :: String -> Module Ann -> Aff (Tuple String String)
sourceFileHash input (Module m) = do
  text <- FS.readTextFile UTF8 (input <> "/" <> unwrap m.name <> "/corefn.json")
  pure (Tuple (unwrap m.name) (Wire.hash text))

prepareModules :: String -> Int -> Int -> Maybe String -> InlineDirectiveMap -> Aff PreparedModules
prepareModules input expectedCount jobs mainModule directives =
  _.prepared <$> prepareModulesWithCache input expectedCount jobs mainModule directives

prepareModulesWithCache :: String -> Int -> Int -> Maybe String -> InlineDirectiveMap -> Aff PreparedCache
prepareModulesWithCache input expectedCount jobs mainModule directives = do
  sortedModules <- Metrics.measure "load TAST + sort" \_ -> coreFnModulesFromOutput input
  when (Array.length (Array.fromFoldable sortedModules) /= expectedCount) $
    throwError (error "TAST decoding did not produce every inventoried module")
  prepareLoadedModules input jobs mainModule directives sortedModules

-- Reuse the decoded inventory on a resident fallback. A source edit can change
-- another module's specializations even when every interface remains identical.
prepareLoadedModules :: String -> Int -> Maybe String -> InlineDirectiveMap -> List (Module Ann) -> Aff PreparedCache
prepareLoadedModules input jobs mainModule directives sortedModules =
  Metrics.measure "prepare + monomorphize" \_ -> do
    let
      sourceModules = Array.fromFoldable sortedModules
      modulesWithClasses = map addClassDataDeclarations sortedModules
      metadata = sourceMetadata sourceModules modulesWithClasses
      globalTypes = metadata.globalTypes

    -- The final per-module rewrite joins the same preparation pool as the
    -- transitive rounds: pure thunks over immutable snapshots, results in
    -- caller order, evaluated only while this Aff runs (so the phase measure
    -- below still covers it).
    transitiveReads <- liftEffect (Ref.new Set.empty)
    Tuple monoBase modules <- monomorphizeModulesWithCache
      (\ast instantiations -> Metrics.measure "transitive specializations" \_ ->
        do
          collected <- transitiveCollectWithReads (runPreparationJobs jobs) ast instantiations
          liftEffect (Ref.write collected.bodyReads transitiveReads)
          pure collected.instantiations)
      (runPreparationJobs jobs)
      globalTypes
      modulesWithClasses

    transitiveBodyReads <- liftEffect (Ref.read transitiveReads)
    let mono = monoBase { bodyReads = Set.union monoBase.bodyReads transitiveBodyReads }
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
    sourceHashes <- Map.fromFoldable <$> traverse (sourceFileHash input) (Array.fromFoldable sourceModules)
    let prepared =
          { directives, modules, metadata, metadataUnchanged: false, mainModules: selectMainModules mainModule sourceModules
          , sources: map (\(Module m) -> { name: unwrap m.name, path: m.path, hasForeign: not (Map.isEmpty m.foreign) }) sourceModules
          }
    pure { prepared, sourceHashes, sourceModules: sortedModules, sourceInterfaces: Map.empty, mono, reusable: Set.empty, proof: Nothing, materialize: Nothing }

-- Metadata always comes from source declarations, independently of the compact
-- specialization proof. Rebuilding these tables does not traverse mono bodies.
sourceMetadata :: Array (Module Ann) -> List (Module Ann) -> CodegenMetadata
sourceMetadata sourceModules modulesWithClasses =
  let
    elidedCtors = collectElidedConstructors sourceModules
    ctorTypes = buildConstructorTypes sourceModules
    globalTypes = buildGlobalTypes sourceModules
    classDeclsFields = buildClassFields sourceModules
    reboxFields = buildReboxFieldIndex ctorTypes classDeclsFields
    { pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves } = buildPointerAdtMetadata (Array.fromFoldable modulesWithClasses)
    { enumAdts, enumCtors } = buildEnumAdtMetadata sourceModules
  in
    { elidedCtors, ctorTypes, globalTypes, globalFunctions: Map.empty
    , ffiFunctions: Map.empty, classDeclsFields, reboxFields
    , pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves, enumAdts, enumCtors
    }

prepareModulesDelta :: String -> Int -> Int -> Maybe String -> InlineDirectiveMap -> PreparedCache -> Aff PreparedCache
prepareModulesDelta input expectedCount jobs mainModule directives cache = do
  loaded <- Metrics.measure "load TAST + sort" \_ -> coreFnModulesFromOutput input
  when (Array.length (Array.fromFoldable loaded) /= expectedCount) $
    throwError (error "TAST decoding did not produce every inventoried module")
  hashes <- Map.fromFoldable <$> traverse (sourceFileHash input) (Array.fromFoldable loaded)
  let changed = Array.filter (\(Module m) -> Map.lookup (unwrap m.name) cache.sourceHashes /= Just (fromMaybe "" (Map.lookup (unwrap m.name) hashes))) (Array.fromFoldable loaded)
  if Array.null changed && Map.size hashes == Map.size cache.sourceHashes then
    pure (cache
      { prepared = cache.prepared
          { directives = directives
          , mainModules = selectMainModules mainModule (Array.fromFoldable loaded)
          }
      , reusable = Map.keys hashes
      })
  else
    -- No durable contribution/AST-read proof currently covers this edit. Keep
    -- the daemon protocol and backend caches, but rebuild preparation before
    -- deciding which prepared modules and PBO results are actually unchanged.
    prepareLoadedModules input jobs mainModule directives loaded

selectMainModules :: Maybe String -> Array (Module Ann) -> Array String
selectMainModules requested modules = case requested of
  Just name -> [ name ]
  Nothing -> Array.mapMaybe exportedMain modules
  where
  exportedMain (Module m)
    | Array.elem (Ident "main") m.exports = Just (unwrap m.name)
    | otherwise = Nothing
