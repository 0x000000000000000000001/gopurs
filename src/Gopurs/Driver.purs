module Gopurs.Driver (compile) where

import Prelude

import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Effect.Aff (Aff, bracket)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Gopurs.Driver.Build (build)
import Gopurs.Driver.Cache as BuildCache
import Gopurs.Driver.Config (readConfig)
import Gopurs.Driver.Output (writeEntryPoints, writeRuntime, writePackageFiles)
import Gopurs.Driver.Prepare (prepareModules)
import Gopurs.Metrics as Metrics
import Gopurs.Diagnostics as Diagnostics
import PureScript.Backend.Optimizer.App (loadDirectives)
import PureScript.Backend.Optimizer.Cache as Cache

-- The driver owns phase ordering; each phase owns its inputs and mutable state.
compile :: Aff Unit
compile = Metrics.measure "backend total" \_ -> do
  config <- liftEffect readConfig
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
  bracket (Metrics.measure "cache lookup + freeze" \_ -> BuildCache.begin config) BuildCache.close \session -> do
    if session.status == "hit" then
      liftEffect $ Console.error "[gopurs] cache hit: inputs and outputs verified; 0 modules optimized/emitted"
    else do
      liftEffect $ Console.error ("[gopurs] cache miss: " <> session.reason)
      prepared <- prepareModules session.input session.count config.prepareJobs config.mainModule directives
      frozen <- Metrics.measure "freeze FFI" \_ -> BuildCache.plan session prepared.sources
      let ffi = Map.fromFoldable (map (\source -> Tuple source.name source) frozen)
      Metrics.measure "runtime" \_ -> writeRuntime session.output
      Metrics.measure "optimize + emit" \_ -> build session ffi config prepared
      if config.goPackages then do
        request <- Metrics.measure "Go exports" \_ -> BuildCache.packages session prepared.mainModules
        Metrics.measure "Go packages" \_ -> writePackageFiles session.output prepared.mainModules request (session.stage <> "/package-result.json")
      else
        Metrics.measure "entry points" \_ -> writeEntryPoints session.output prepared.mainModules
      diagnostics <- liftEffect Diagnostics.recorded
      Metrics.measure "cache publish" \_ -> BuildCache.publish session prepared.mainModules diagnostics

  case config.allocationProfile of
    Just path -> liftEffect (Cache.writeAllocProfile path)
    Nothing -> pure unit
