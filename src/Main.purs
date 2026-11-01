module Main (main) where

import Prelude

import Data.Either (Either(..))
import Data.Maybe (Maybe(..))
import Effect (Effect)
import Effect.Aff (runAff_)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Exception (message)
import Gopurs.Driver (compileWithConfig)
import Gopurs.Driver.Config (readConfig)
import Gopurs.Host.Resident as Resident
import Gopurs.Host.Snapshot as Snapshot
import Node.Process as Process

foreign import exitFailure :: Effect Unit

main :: Effect Unit
main = Snapshot.withSnapshot $ runAff_ onComplete do
  config <- liftEffect readConfig
  resident <- liftEffect (Process.lookupEnv "GOPURS_RESIDENT")
  if resident == Just "1" then Resident.run config else compileWithConfig config
  where
  -- Observe the result only after the driver's brackets and supervision have
  -- finished. A discarded Aff fiber does not report failure to the native CLI.
  onComplete result = do
    Snapshot.release
    case result of
      Right _ -> pure unit
      Left err -> do
        Console.error ("[gopurs] error: " <> message err)
        exitFailure
