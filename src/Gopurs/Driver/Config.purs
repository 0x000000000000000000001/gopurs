module Gopurs.Driver.Config (CompilerConfig, readConfig) where

import Prelude

import Data.Int as Int
import Data.Maybe (Maybe(..), fromMaybe)
import Effect (Effect)
import Gopurs.Emission (EmissionOptions)
import Node.Process as Process
import PureScript.Backend.Optimizer.App (parseCLIArgs)

-- Only options consumed by the Go backend belong here. The shared PBO parser
-- also recognises options used by other backends.
type CompilerConfig =
  { mainModule :: Maybe String
  , ffiDirectory :: Maybe String
  , rewriteLimit :: Int
  , prepareJobs :: Int
  , optimizerJobs :: Int
  , emission :: EmissionOptions
  , allocationProfile :: Maybe String
  }

readConfig :: Effect CompilerConfig
readConfig = do
  args <- parseCLIArgs <$> Process.argv
  prepareJobs <- readJobs "GOPURS_PREPARE_JOBS" 2
  emitJobs <- readJobs "GOPURS_EMIT_JOBS" 8
  optimizerJobs <- readJobs "GOPURS_PBO_JOBS" 1
  pipeline <- Process.lookupEnv "GOPURS_PIPELINE"
  allocationProfile <- Process.lookupEnv "GOPURS_ALLOC_PROFILE"
  pure
    { mainModule: args.mbMainModule
    , ffiDirectory: args.mbFfiDir
    , rewriteLimit: fromMaybe 10_000 args.mbRewriteLimit
    -- Preparation partitions its own tasks and caps concurrency at eight.
    , prepareJobs
    , optimizerJobs: boundedJobs optimizerJobs
    , emission: { jobs: boundedJobs emitJobs, pipelined: pipeline /= Just "0" }
    , allocationProfile
    }

readJobs :: String -> Int -> Effect Int
readJobs variable fallback = do
  configured <- Process.lookupEnv variable
  pure (fromMaybe fallback (configured >>= Int.fromString))

boundedJobs :: Int -> Int
boundedJobs = max 1 <<< min 64
