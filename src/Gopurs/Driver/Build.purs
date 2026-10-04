module Gopurs.Driver.Build (build) where

import Prelude

import Control.Parallel (parTraverse)
import Data.Either (either)
import Data.Foldable (foldl)
import Data.Int as Int
import Data.List (List)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.Traversable (traverse)
import Effect.Aff (Aff, attempt, forkAff, throwError)
import Effect.Aff.AVar as Avar
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Ref as Ref
import Gopurs.Driver.Config (CompilerConfig)
import Gopurs.Driver.Output (emitModule)
import Gopurs.Driver.Prepare (PreparedModules)
import Gopurs.Emission (EmissionEntry, withEmitter)
import Gopurs.Metrics as Metrics
import PureScript.Backend.Optimizer.Builder (BuildOptions, JobScheduler, ParallelStats, buildModules, buildModulesParallel)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (Ann, Module)
import PureScript.Backend.Optimizer.Semantics.Foreign (coreForeignSemantics)

type OptimizedModule =
  { coreFnModule :: Module Ann
  , backendMod :: BackendModule
  }

build :: CompilerConfig -> PreparedModules -> Aff Unit
build config prepared = do
  globalFunctionsRef <- liftEffect (Ref.new Map.empty)
  attemptsRef <- liftEffect (Ref.new 0)
  codegenRef <- liftEffect (Ref.new 0)
  emissionMillisRef <- liftEffect (Ref.new 0.0)
  let
    emitBatch :: Array OptimizedModule -> Aff Unit
    emitBatch batch = Metrics.accumulate emissionMillisRef \_ -> do
      globalFunctions <- liftEffect (Ref.read globalFunctionsRef)
      let
        metadata = prepared.metadata { globalFunctions = globalFunctions }
        emit entry = emitModule metadata config.ffiDirectory entry.coreFnModule entry.backendMod
      functions <- if config.emission.jobs == 1 then traverse emit batch else parTraverse emit batch
      -- Read one immutable snapshot per batch and publish in module order.
      -- Later batches must see all signatures produced by this batch.
      liftEffect (Ref.write (foldl (flip Map.union) globalFunctions functions) globalFunctionsRef)

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
      -- Regenerate every module and its FFI output on each invocation.
      , onSkipModule: \_ _ -> pure Nothing
      , onCodegenModule: \env coreFnModule backendMod _ -> do
          liftEffect (Ref.modify_ (_ + 1) codegenRef)
          when (env.moduleIndex `mod` 100 == 0) $ liftEffect $ Console.error $
            "[gopurs] optimize + emit: module " <> show (env.moduleIndex + 1)
              <> "/" <> show env.moduleCount <> " (" <> unwrap backendMod.name <> ")"
          enqueue
            { name: backendMod.name
            , imports: backendMod.imports
            , value: { coreFnModule, backendMod }
            }
      }

  withEmitter config.emission emitBatch \enqueue ->
    Metrics.measure "PBO producer" \_ ->
      runBuilder config.optimizerJobs (buildOptions enqueue) prepared.modules

  attempts <- liftEffect (Ref.read attemptsRef)
  codegen <- liftEffect (Ref.read codegenRef)
  emissionMillis <- liftEffect (Ref.read emissionMillisRef)
  liftEffect $ Console.error $
    "[gopurs] generation + writes (cumulative batches): " <> show (Int.round emissionMillis) <> " ms"
  liftEffect $ Console.error $
    "[gopurs] pbo module attempts: " <> show attempts <> ", codegen: " <> show codegen

runBuilder :: Int -> BuildOptions Aff -> List (Module Ann) -> Aff Unit
runBuilder jobs options modules
  | jobs <= 1 = buildModules options modules
  | otherwise = do
      scheduler <- createJobScheduler
      buildModulesParallel { jobs, scheduler, onStats: Just reportParallelStats } options modules

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
