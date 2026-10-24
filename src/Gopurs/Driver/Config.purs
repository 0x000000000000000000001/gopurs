module Gopurs.Driver.Config (CompilerConfig, readConfig) where

import Prelude

import Data.Array as Array
import Data.Int as Int
import Data.Maybe (Maybe(..), fromMaybe)
import Effect (Effect)
import Effect.Exception (throw)
import Gopurs.Emission (EmissionOptions)
import Node.Process as Process
import PureScript.Backend.Optimizer.App (parseCLIArgs)

-- Only options consumed by the Go backend belong here. The shared PBO parser
-- also recognises options used by other backends.
type CompilerConfig =
  { mainModule :: Maybe String
  , ffiDirectory :: Maybe String
  , rewriteLimit :: Int
  , devMode :: Boolean
  , prepareJobs :: Int
  , optimizerJobs :: Int
  , emission :: EmissionOptions
  , allocationProfile :: Maybe String
  , forceBuild :: Boolean
  , goPackages :: Boolean
  , buildManifest :: Boolean
  }

readConfig :: Effect CompilerConfig
readConfig = do
  argv <- Process.argv
  let args = parseCLIArgs argv
  let devMode = Array.elem "--dev" argv
  when (devMode && Array.elem "--release" argv) $
    throw "Choose either --dev or --release, not both."
  prepareJobs <- readJobs "GOPURS_PREPARE_JOBS" 2
  emitJobs <- readJobs "GOPURS_EMIT_JOBS" 8
  optimizerJobs <- readJobs "GOPURS_PBO_JOBS" 1
  pipeline <- Process.lookupEnv "GOPURS_PIPELINE"
  allocationProfile <- Process.lookupEnv "GOPURS_ALLOC_PROFILE"
  forceBuild <- Process.lookupEnv "GOPURS_FORCE_BUILD"
  goPackages <- Process.lookupEnv "GOPURS_GO_PACKAGES"
  buildManifest <- Process.lookupEnv "GOPURS_BUILD_MANIFEST"
  pure
    { mainModule: args.mbMainModule
    , ffiDirectory: args.mbFfiDir
    , rewriteLimit: fromMaybe 10_000 args.mbRewriteLimit
    , devMode
    -- Preparation partitions its own tasks and caps concurrency at eight.
    , prepareJobs: max 1 (min 8 prepareJobs)
    , optimizerJobs: boundedJobs optimizerJobs
    , emission: { jobs: boundedJobs emitJobs, pipelined: pipeline /= Just "0" }
    , allocationProfile
    , forceBuild: forceBuild == Just "1" || Array.elem "--force" argv
    , goPackages: goPackages == Just "1"
    , buildManifest: buildManifest == Just "1"
    }

readJobs :: String -> Int -> Effect Int
readJobs variable fallback = do
  configured <- Process.lookupEnv variable
  pure (fromMaybe fallback (configured >>= Int.fromString))

boundedJobs :: Int -> Int
boundedJobs = max 1 <<< min 64
