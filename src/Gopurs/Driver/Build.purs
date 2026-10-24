module Gopurs.Driver.Build (build, buildWithHeaders) where

import Prelude

import Control.Parallel (parTraverse)
import Data.Array as Array
import Data.Either (either)
import Data.Foldable (foldl, for_)
import Data.Int as Int
import Data.List (List)
import Data.List as List
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Effect.Aff (Aff, attempt, forkAff, throwError)
import Effect.Exception (error)
import Effect.Aff.AVar as Avar
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Ref as Ref
import Gopurs.Driver.Config (CompilerConfig)
import Gopurs.Driver.Cache (FrozenFFI, Session, ReusedMemo)
import Gopurs.Cache.HeaderPreload as HeaderPreload
import Gopurs.Cache.Modules as ModuleCache
import Gopurs.Cache.Wire as Wire
import Gopurs.Driver.Prepare (PreparedModules)
import Gopurs.Emission (EmissionEntry, withEmitter)
import Gopurs.Metrics as Metrics
import PureScript.Backend.Optimizer.Builder (BuildEnv, BuildOptions, BuildObservation, JobScheduler, ParallelStats, buildModulesWithObservationAndProfile, buildModulesParallelWithBindingLookupAndProfile)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (Ann, Module(..))
import PureScript.Backend.Optimizer.Semantics.Foreign (coreForeignSemantics)
import PureScript.Backend.Optimizer.Semantics (OptimizationProfile(..))

type OptimizedModule =
  { coreFnModule :: Module Ann
  , backendMod :: BackendModule
  , optimizedCached :: Boolean
  }

build :: Session -> Map.Map String FrozenFFI -> CompilerConfig -> PreparedModules -> Set.Set String -> Maybe (String -> Aff (Maybe (Module Ann))) -> Set.Set String -> Ref.Ref (Map.Map String BackendModule) -> Map.Map String String -> Aff (Array ReusedMemo)
build = buildWithHeaders Nothing

buildWithHeaders :: Maybe (Aff (Maybe HeaderPreload.Reservoir)) -> Session -> Map.Map String FrozenFFI -> CompilerConfig -> PreparedModules -> Set.Set String -> Maybe (String -> Aff (Maybe (Module Ann))) -> Set.Set String -> Ref.Ref (Map.Map String BackendModule) -> Map.Map String String -> Aff (Array ReusedMemo)
buildWithHeaders pendingHeaders session ffi config preparedInput reusableModules materializePrepared headerSafe backends priorFfi = do
  initialized <- Metrics.measure "module cache keys" \_ -> ModuleCache.initialize config.prepareJobs session preparedInput reusableModules backends (Map.fromFoldable (map (\(Tuple name value) -> Tuple name value.digest) (Map.toUnfoldable ffi :: Array (Tuple String FrozenFFI)))) priorFfi headerSafe
  headerValues <- case pendingHeaders of
    Nothing -> pure Nothing
    Just pending -> Metrics.measure "module header preload join" \_ -> pending
  cache <- Metrics.measure "module header bytes" \_ ->
    ModuleCache.loadHeaderTextsWithPreload headerValues (initialized { lazyContributions = config.optimizerJobs > 1 })
  preparedReaders <- liftEffect (Ref.new Map.empty)
  readerMillis <- liftEffect (Ref.new 0.0)
  let prepared = preparedInput
  let
    moduleRanks = Map.fromFoldable (Array.mapWithIndex (\rank (Module m) -> Tuple (unwrap m.name) rank) (Array.fromFoldable prepared.modules))
    prepareFresh mod@(Module m) = do
      ModuleCache.warmImplementationInputs config.prepareJobs cache moduleRanks (unwrap m.name)
      materializeFresh mod
    materializeFresh mod@(Module m) = case materializePrepared of
      Nothing -> pure mod
      Just materialize | Set.member (unwrap m.name) reusableModules -> do
        let name = unwrap m.name
        previous <- liftEffect (Map.lookup name <$> Ref.read preparedReaders)
        case previous of
          Just value -> pure value
          Nothing -> Metrics.accumulate readerMillis \_ -> do
            unless (Map.member name cache.candidates) $ throwError (error ("incremental cache: missing prepared candidate: " <> name))
            restored <- materialize name
            case restored of
              Nothing -> throwError (error ("incremental cache: cannot prepare reader: " <> name))
              Just value -> do
                fingerprint <- liftEffect (Wire.hash <$> Wire.write value)
                case Map.lookup name cache.keys of
                  Just keys | keys.prepared == fingerprint -> do
                    liftEffect (Ref.modify_ (Map.insert name value) preparedReaders)
                    pure value
                  _ -> throwError (error ("incremental cache: reader preparation changed: " <> name))
      _ -> pure mod
    checkedLookup (Module m) lookup = do
      value <- lookup
      case materializePrepared, value of
        Just _, Nothing | Set.member (unwrap m.name) reusableModules && Set.member (unwrap m.name) cache.headerSafe ->
          throwError (error ("incremental cache: unavailable prepared placeholder: " <> unwrap m.name))
        _, _ -> pure value
  reusable <- Metrics.measure "module cache all-hit lookup" \_ -> ModuleCache.lookupAll config.optimizerJobs cache prepared.modules
  restoredHeaders <- if config.optimizerJobs <= 1 then pure Map.empty else case reusable of
    Just _ -> pure Map.empty
    Nothing -> Metrics.measure "module header decode" \_ -> ModuleCache.lookupHeaders config.optimizerJobs cache prepared.modules
  workerOwnersRef <- liftEffect (Ref.new Map.empty)
  workersDigestRef <- liftEffect (Ref.new (Wire.hash "workers-v1"))
  attemptsRef <- liftEffect (Ref.new 0)
  codegenRef <- liftEffect (Ref.new 0)
  emissionMillisRef <- liftEffect (Ref.new 0.0)
  let
    emitBatch :: Array OptimizedModule -> Aff Unit
    emitBatch batch = Metrics.accumulate emissionMillisRef \_ -> do
      workerOwners <- liftEffect (Ref.read workerOwnersRef)
      visibleWorkers <- liftEffect (Ref.read workersDigestRef)
      ModuleCache.warmWorkerSnapshot config.prepareJobs cache workerOwners (map _.backendMod (Array.filter (not <<< _.optimizedCached) batch))
      let
        resolve = ModuleCache.resolveWorkerSnapshot (ModuleCache.readWorkerContribution cache) workerOwners
        emit entry = ModuleCache.emitWithWorkerSnapshot resolve cache ffi prepared.metadata visibleWorkers entry.optimizedCached entry.coreFnModule entry.backendMod
      results <- if config.emission.jobs == 1 then traverse emit batch else parTraverse emit batch
      -- Read one immutable snapshot per batch and publish in module order.
      -- Later batches must see all signatures produced by this batch.
      liftEffect (Ref.write (foldl (\acc (Tuple entry result) ->
        Map.insert (unwrap entry.backendMod.name) (if result.deferred then Nothing else Just result.functions) acc)
        workerOwners (Array.zip batch results)) workerOwnersRef)
      liftEffect (Ref.write (foldl (\acc result -> Wire.hash (acc <> "/" <> result.workers)) visibleWorkers results) workersDigestRef)

    buildOptions :: (EmissionEntry OptimizedModule -> Aff Unit) -> BuildOptions Aff
    buildOptions enqueue =
      { directives: prepared.directives
      , analyzeCustom: \_ _ -> Nothing
      , foreignSemantics: coreForeignSemantics
      , traceIdents: Set.empty
      , rewriteLimit: config.rewriteLimit
      -- Attempts can be replayed; only finalised modules advance progress.
      , onPrepareModule: \_ m -> do
          liftEffect (Ref.modify_ (_ + 1) attemptsRef)
          pure m
       , onSkipModule: \_ coreFnModule -> checkedLookup coreFnModule (ModuleCache.lookupBackend cache coreFnModule)
      , onCodegenModule: \env coreFnModule backendMod _ -> do
          liftEffect (Ref.modify_ (_ + 1) codegenRef)
          when (env.moduleIndex `mod` 100 == 0) $ liftEffect $ Console.error $
            "[gopurs] optimize + emit: module " <> show (env.moduleIndex + 1)
              <> "/" <> show env.moduleCount <> " (" <> unwrap backendMod.name <> ")"
          enqueue
            { name: backendMod.name
            , imports: backendMod.imports
             , value: { coreFnModule, backendMod, optimizedCached: false }
            }
      }

  withEmitter config.emission emitBatch \enqueue -> do
    let enqueueCached coreFnModule backendMod = enqueue
           { name: backendMod.name, imports: backendMod.imports, value: { coreFnModule, backendMod, optimizedCached: true } }
    case reusable of
      Just previousBackends -> do
        liftEffect $ Console.error $ "[gopurs] PBO cache: all " <> show (Array.length previousBackends) <> " backends validated; optimizer not started"
        Metrics.measure "cached module producer" \_ ->
          for_ (Array.zip (Array.fromFoldable prepared.modules) previousBackends) \(Tuple coreFnModule backendMod) -> enqueueCached coreFnModule backendMod
      Nothing -> Metrics.measure "PBO producer" \_ ->
        runBuilder (if config.devMode then Dev else Release) config.optimizerJobs cache restoredHeaders prepareFresh
          { onSkipModule: \_ mod view -> checkedLookup mod (ModuleCache.lookupBackendWithReads cache mod view)
          , onReads: \_ mod reads -> ModuleCache.captureReads cache mod reads
          }
          (\_ -> enqueueCached) (buildOptions enqueue) prepared.modules

  attempts <- liftEffect (Ref.read attemptsRef)
  codegen <- liftEffect (Ref.read codegenRef)
  ModuleCache.report cache (List.length prepared.modules) codegen
  emissionMillis <- liftEffect (Ref.read emissionMillisRef)
  liftEffect $ Console.error $
    "[gopurs] generation + writes (cumulative batches): " <> show (Int.round emissionMillis) <> " ms"
  liftEffect $ Console.error $
    "[gopurs] pbo module attempts: " <> show attempts <> ", codegen: " <> show codegen
  workerOwnerCount <- liftEffect (Map.size <$> Ref.read cache.priorWorkers)
  liftEffect $ Console.error $ "[gopurs] worker contributions: decoded=" <> show workerOwnerCount <> "/" <> show (List.length prepared.modules)
  readerCount <- liftEffect (Map.size <$> Ref.read preparedReaders)
  preparedMillis <- liftEffect (Ref.read readerMillis)
  liftEffect $ Console.error $ "[gopurs] prepare PBO readers: " <> show readerCount <> " modules, " <> show (Int.round preparedMillis) <> " ms"
  capsules <- liftEffect (Ref.read cache.priorCapsules)
  capsuleHits <- liftEffect (Ref.read cache.capsuleHits)
  capsuleCreated <- liftEffect (Ref.read cache.capsuleCreated)
  liftEffect $ Console.error $ "[gopurs] binding capsules: loaded=" <> show (Map.size (Map.mapMaybe identity capsules))
    <> ", attempted=" <> show (Map.size capsules) <> ", hits=" <> show capsuleHits <> ", created=" <> show capsuleCreated
  liftEffect (Array.fromFoldable <<< Map.values <$> Ref.read cache.reused)

runBuilder :: OptimizationProfile -> Int -> ModuleCache.ModuleCache -> Map.Map String BackendModule -> (Module Ann -> Aff (Module Ann)) -> BuildObservation Aff -> (BuildEnv -> Module Ann -> BackendModule -> Aff Unit) -> BuildOptions Aff -> List (Module Ann) -> Aff Unit
runBuilder profile jobs cache restoredHeaders prepareFresh observer onReuse options modules
  | jobs <= 1 = buildModulesWithObservationAndProfile profile (Just observer) onReuse options modules
  | otherwise = do
      let edges = foldl (\count owners -> count + Array.length owners) 0 (Map.values cache.schedulingReads)
      liftEffect $ Console.error $ "[gopurs] PBO scheduling hints: owners=" <> show (Map.size cache.schedulingReads) <> ", edges=" <> show edges
      liftEffect $ Console.error $ "[gopurs] PBO restored headers: " <> show (Map.size restoredHeaders)
      scheduler <- createJobScheduler
      buildModulesParallelWithBindingLookupAndProfile profile (Just (ModuleCache.readCachedBinding cache)) restoredHeaders cache.schedulingReads prepareFresh (Just (ModuleCache.readImplementationsSync cache)) (Just observer) onReuse { jobs, scheduler, onStats: Just reportParallelStats } options modules

-- An AVar has a single result slot. A worker's put can remain suspended if the
-- coordinator fails before taking its result. Run this scheduler inside the
-- withEmitter supervision scope, which cancels and joins those workers too.
createJobScheduler :: Aff (JobScheduler Aff)
createJobScheduler = do
  results <- Avar.empty
  let
    fork job = void $ forkAff do
      outcome <- attempt (job unit)
      Avar.put outcome results
    await = do
      outcome <- Avar.take results
      either throwError pure outcome
  pure { fork, await }

reportParallelStats :: ParallelStats -> Aff Unit
reportParallelStats stats = liftEffect $ Console.error $
  "[gopurs] pbo stats: dispatched=" <> show stats.dispatched
    <> ", fallbackDispatched=" <> show stats.fallbackDispatched
    <> ", maxReady=" <> show stats.maxReady
    <> ", deferredAttempts=" <> show stats.deferredAttempts
    <> ", wakeups=" <> show stats.wakeups
    <> ", waitingPeak=" <> show stats.waitingPeak
    <> ", attemptMillis=" <> show (Int.round stats.attemptMillis)
    <> ", attemptMaxMillis=" <> show (Int.round stats.attemptMaxMillis)
    <> ", coordinatorMillis=" <> show (Int.round stats.coordinatorMillis)
    <> ", awaitMillis=" <> show (Int.round stats.awaitMillis)
    <> ", emitMillis=" <> show (Int.round stats.emitMillis)
