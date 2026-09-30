module Gopurs.Driver (compile) where

import Prelude

import Data.Maybe (Maybe(..))
import Effect.Aff (Aff)
import Effect.Class (liftEffect)
import Gopurs.Driver.Build (build)
import Gopurs.Driver.Config (readConfig)
import Gopurs.Driver.Output (writeEntryPoints, writeRuntime)
import Gopurs.Driver.Prepare (prepareModules)
import Gopurs.Metrics as Metrics
import PureScript.Backend.Optimizer.Cache as Cache

-- The driver owns phase ordering; each phase owns its inputs and mutable state.
compile :: Aff Unit
compile = Metrics.measure "backend total" \_ -> do
  config <- liftEffect readConfig
  prepared <- prepareModules config.prepareJobs config.mainModule

  -- Keep allocation sampling disabled unless explicitly requested. The native
  -- compiler allocates heavily, and sampling itself has a measurable cost.
  case config.allocationProfile of
    Nothing -> liftEffect (Metrics.setMemProfileRate 0)
    Just _ -> pure unit

  Metrics.measure "runtime" \_ -> writeRuntime
  Metrics.measure "optimize + emit" \_ -> build config prepared
  Metrics.measure "entry points" \_ -> writeEntryPoints prepared.mainModules

  case config.allocationProfile of
    Just path -> liftEffect (Cache.writeAllocProfile path)
    Nothing -> pure unit
