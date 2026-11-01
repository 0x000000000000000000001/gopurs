module Gopurs.Driver (compile, compileWithConfig, compileWithConfigCache) where

import Prelude

import Data.Either (Either(..))
import Data.Map as Map
import Data.Newtype (unwrap)
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..))
import Effect.Aff (Aff, attempt, forkAff, joinFiber, throwError)
import Effect.Class (liftEffect)
import Effect.Ref as Ref
import Effect.Exception (message)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (Module(..))
import Effect.Console as Console
import Gopurs.Driver.Build (buildWithHeaders)
import Gopurs.Cache.HeaderPreload as HeaderPreload
import Gopurs.Driver.Cache as BuildCache
import Gopurs.Driver.Config (CompilerConfig, readConfig)
import Gopurs.Driver.Output (writeEntryPoints, writeRuntime, writePackageFiles)
import Gopurs.Driver.Prepare (PreparedCache, prepareModulesDeltaForProfile, prepareModulesWithCacheForProfile)
import Gopurs.Driver.Prepare.Persistent (preloadProjection, preparePersistentWithProjectionForProfile, publishPersistent)
import Gopurs.Driver.Manifest as Manifest
import Gopurs.Cache.Modules as ModuleCache
import Gopurs.Host.Metrics as Metrics
import Data.Array as Array
import Gopurs.Host.Diagnostics as Diagnostics
import PureScript.Backend.Optimizer.App (loadDirectives)
import PureScript.Backend.Optimizer.Cache as Cache

-- The driver owns phase ordering; each phase owns its inputs and mutable state.
compile :: Aff Unit
compile = do
  config <- liftEffect readConfig
  compileWithConfig config

compileWithConfig :: CompilerConfig -> Aff Unit
compileWithConfig config = do
  cache <- liftEffect (Ref.new Nothing)
  backends <- liftEffect (Ref.new Map.empty)
  ffiDigests <- liftEffect (Ref.new Map.empty)
  compileWithConfigCache config cache backends ffiDigests

compileWithConfigCache :: CompilerConfig -> Ref.Ref (Maybe PreparedCache) -> Ref.Ref (Map.Map String BackendModule) -> Ref.Ref (Map.Map String String) -> Aff Unit
compileWithConfigCache config cacheRef backends ffiDigestsRef = do
  outcome <- attempt (compileCached config cacheRef backends ffiDigestsRef)
  case outcome of
    Right _ -> pure unit
    Left err -> do
      -- A failed build may have emitted only part of its backend state. None
      -- of that uncommitted memory can justify a later disk-cache hit.
      liftEffect do
        Ref.write Nothing cacheRef
        Ref.write Map.empty backends
        Ref.write Map.empty ffiDigestsRef
      if not config.forceBuild && String.contains (Pattern "incremental cache:") (message err) then do
        liftEffect $ Console.error ("[gopurs] full fallback: " <> message err)
        compileWithConfigCache (config { forceBuild = true }) cacheRef backends ffiDigestsRef
      else throwError err

compileCached :: CompilerConfig -> Ref.Ref (Maybe PreparedCache) -> Ref.Ref (Map.Map String BackendModule) -> Ref.Ref (Map.Map String String) -> Aff Unit
compileCached config cacheRef backends ffiDigestsRef = Metrics.measure "backend total" \_ -> do
  liftEffect $ Console.error $
    "[gopurs] workers: prepare=" <> show config.prepareJobs
      <> ", pbo=" <> show config.optimizerJobs
      <> ", emit=" <> show config.emission.jobs
      <> ", pipeline=" <> show (config.emission.pipelined && config.emission.jobs > 1)
  -- Keep allocation sampling disabled unless explicitly requested. The native
  -- compiler allocates heavily, and sampling itself has a measurable cost.
  case config.allocationProfile of
    Nothing -> liftEffect (Metrics.setMemProfileRate 0)
    Just _ -> pure unit

  -- Directive diagnostics are inexpensive and retain their normal behavior on
  -- hits too. Other semantic diagnostics conservatively disable reuse.
  directives <- loadDirectives
  let preload = if not config.forceBuild && not (isJust config.allocationProfile)
        && config.buildManifest && config.optimizerJobs > 1
        then Just (preloadProjection config.prepareJobs directives)
        else Nothing
  BuildCache.withSession config preload \session projected -> do
    if session.status == "hit" then do
      -- Another process may have published this generation while the resident
      -- was idle. Its unkeyed memory must not be paired with these disk keys.
      liftEffect do
        Ref.write Nothing cacheRef
        Ref.write Map.empty backends
        Ref.write Map.empty ffiDigestsRef
      liftEffect $ Console.error "[gopurs] cache hit: inputs and outputs verified; 0 modules optimized/emitted"
    else HeaderPreload.withPreload session (isJust preload && isJust projected) \rawHeaders -> do
      liftEffect (Ref.write Map.empty backends)
      liftEffect $ Console.error ("[gopurs] cache miss: " <> session.reason)
      cached <- liftEffect (Ref.read cacheRef)
      preparedBase <- Metrics.measure "prepare delta" \_ ->
        if config.forceBuild then do
          BuildCache.materializeAllInputs session
          prepareModulesWithCacheForProfile config.devMode session.input session.count config.prepareJobs config.mainModule directives
        else if config.buildManifest && config.optimizerJobs > 1 then
          preparePersistentWithProjectionForProfile config.devMode projected session config.prepareJobs config.mainModule directives
        else case cached of
          Just previous -> do
            BuildCache.materializeAllInputs session
            prepareModulesDeltaForProfile config.devMode session.input session.count config.prepareJobs config.mainModule directives previous
          Nothing -> do
            BuildCache.materializeAllInputs session
            prepareModulesWithCacheForProfile config.devMode session.input session.count config.prepareJobs config.mainModule directives
      -- The anchored compact proof already establishes this exact source
      -- delta. The optional manifest adds no evidence on that path; leave it
      -- for the next full preparation, where an older sidecar is conservative.
      let useManifest = config.buildManifest && case preparedBase.compactSourceDelta of
            Nothing -> true
            Just _ -> false
      previousManifest <- if useManifest
        then Metrics.measure "build manifest read" \_ -> Manifest.readManifest session.cacheRoot
        else pure Nothing
      manifest <- if useManifest
        then Just <$> Metrics.measure "build manifest encode" \_ ->
          liftEffect (Manifest.buildManifestWithInterfaces preparedBase.sourceInterfaces previousManifest session.root session.context directives (Array.fromFoldable preparedBase.sourceModules) preparedBase.sourceHashes)
        else pure Nothing
      let
        manifestChanged = case manifest, previousManifest of
          Just current, Just saved
            | saved.context == current.context && saved.directives == current.directives ->
                Manifest.diffManifest saved current
          Just _, _ -> Map.singleton "manifest" "context-changed"
          Nothing, _ -> case preparedBase.compactSourceDelta of
            Just name -> Map.singleton name "source"
            Nothing -> Map.empty
        -- A caller body can change another module's specialized bindings even
        -- when every source interface is identical. Only preparation itself
        -- can prove reuse of a prepared fingerprint.
        withCache = preparedBase
      case manifest of
        Just current -> do
          let
            interfaceChanged = Array.length (Array.filter (\(Tuple _ reason) -> reason == "interface" || reason == "source+interface") (Map.toUnfoldable manifestChanged :: Array _))
          liftEffect $ Console.error ("[gopurs] build manifest: " <> show (Map.size manifestChanged) <> " changed, " <> show interfaceChanged <> " interface, " <> show (Set.size withCache.reusable) <> " reusable")
          Metrics.measure "build manifest write" \_ -> Manifest.writeManifest session.stage session.cacheRoot current
        Nothing -> case withCache.compactSourceDelta of
          Just _ -> liftEffect $ Console.error ("[gopurs] build manifest: 1 changed, 0 interface, " <> show (Set.size withCache.reusable) <> " reusable (committed preparation proof)")
          Nothing -> pure unit
      let prepared = withCache.prepared
      HeaderPreload.withDecodedPreload config.prepareJobs rawHeaders \headers -> do
        frozen <- Metrics.measure "freeze FFI" \_ -> BuildCache.plan session prepared.sources
        let ffi = Map.fromFoldable (map (\source -> Tuple source.name source) frozen)
        let currentFfi = Map.fromFoldable (map (\source -> Tuple source.name source.digest) frozen)
        memoryFfi <- liftEffect (Ref.read ffiDigestsRef)
        diskFfi <- if config.buildManifest then Manifest.readFfiDigests session.cacheRoot else pure Nothing
        let priorFfi = if Map.isEmpty memoryFfi then fromMaybe Map.empty diskFfi else memoryFfi
        let
          ffiChanged = Array.mapMaybe
            (\(Tuple name digest) -> if Map.lookup name priorFfi == Just digest then Nothing else Just name)
            (Map.toUnfoldable currentFfi :: Array (Tuple String String))
          changedSource = Array.mapMaybe
            (\(Tuple name reason) -> if reason == "source" then Just name else Nothing)
            (Map.toUnfoldable manifestChanged :: Array (Tuple String String))
          onlySource = Array.all (\(Tuple _ reason) -> reason == "source") (Map.toUnfoldable manifestChanged :: Array (Tuple String String))
        let moduleNames = map (\(Module value) -> unwrap value.name) (Array.fromFoldable withCache.prepared.modules)
        let
          headerCandidates
            | not (Array.null ffiChanged) = Set.empty
            | otherwise = case withCache.compactSourceDelta of
                Just _ -> Set.fromFoldable moduleNames
                Nothing -> case manifest, previousManifest of
                  Just _, Just _ | onlySource -> case changedSource of
                    [ _ ] -> Set.fromFoldable moduleNames
                    _ -> Set.empty
                  _, _ -> Set.empty
        Metrics.measure "runtime" \_ -> writeRuntime session.output
        reused <- Metrics.measure "optimize + emit" \_ -> buildWithHeaders headers session ffi config prepared withCache.reusable withCache.materialize headerCandidates backends priorFfi
        -- Preparation proof serialization reads immutable prepared data and
        -- writes its own cache objects. Start it alongside Go package planning
        -- and publication, then join before the session closes.
        proofFiber <- if config.buildManifest
          then Just <$> forkAff (Metrics.measure "prepare proof publish" \_ -> publishPersistent session config.prepareJobs withCache)
          else pure Nothing
        if config.goPackages then do
          request <- Metrics.measure "Go exports" \_ -> BuildCache.packages session prepared.mainModules
          Metrics.measure "Go packages" \_ -> writePackageFiles session.output session.cacheRoot prepared.mainModules request (session.stage <> "/package-result.json") (session.stage <> "/package-overlay.json")
        else
          Metrics.measure "entry points" \_ -> writeEntryPoints session.output prepared.mainModules
        diagnostics <- liftEffect Diagnostics.recorded
        Metrics.measure "cache publish" \_ -> BuildCache.publish session prepared.mainModules diagnostics reused
        when config.buildManifest $ do
          case proofFiber of
            Just fiber -> void (joinFiber fiber)
            Nothing -> pure unit
          let restoredNames = Set.fromFoldable (map _.name reused)
          Metrics.measure "reads publish" \_ -> ModuleCache.publishReads session.stage session.cacheRoot (Array.filter (not <<< flip Set.member restoredNames) moduleNames)
          Metrics.measure "ffi digests write" \_ -> Manifest.writeFfiDigests session.stage session.cacheRoot currentFfi
        liftEffect do
          Ref.write (Just withCache) cacheRef
          Ref.write currentFfi ffiDigestsRef

  case config.allocationProfile of
    Just path -> liftEffect (Cache.writeAllocProfile path)
    Nothing -> pure unit
