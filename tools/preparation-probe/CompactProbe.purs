module PreparationCompactProbe (main) where

import Prelude

import Data.Array as Array
import Data.Either (Either(..))
import Data.Maybe (Maybe(..))
import Data.String.CodeUnits as CU
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Aff (Aff, runAff_, throwError)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Exception (error, message)
import Gopurs.Cache.Wire as Wire
import Gopurs.Metrics as Metrics
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import Node.Process as Process
import PreparationProbe (mapJobs, readIndex, restoreSet, writeSet)
import PureScript.Backend.Optimizer.Monomorphize (InstantiationMap, PreparedInstantiation)
import Type.Proxy (Proxy(..))

-- The enclosing snapshot must validate ALL preparation inputs. These digests
-- are audit witnesses, not selective inter-build dependency proofs. In
-- particular, an unchanged digest cannot authorize reuse after a TAST edit.
type CompactContribution =
  { infoDigest :: String
  , observationsDigest :: String
  , contribution :: InstantiationMap
  }

required :: String -> Aff String
required name = do
  value <- liftEffect (Process.lookupEnv name)
  case value of
    Just text -> pure text
    Nothing -> throwError (error ("Missing " <> name))

project :: String -> String -> Aff Unit
project source target = do
  entries <- Metrics.measure "compact/read-full" \_ -> do
    index <- readIndex (source <> "/contributions")
    mapJobs
      (\(Tuple name (Tuple digest bytes)) -> do
        text <- FS.readTextFile UTF8 (source <> "/contributions/" <> Wire.hash name <> ".json")
        unless (CU.length text == bytes && Wire.hash text == digest) $ throwError (error "Full contribution integrity")
        case Wire.read text :: Maybe PreparedInstantiation of
          Just value -> pure (Tuple name value)
          Nothing -> throwError (error ("Invalid full contribution: " <> name)))
      index
  compact <- Metrics.measure "compact/project" \_ -> mapJobs
    (\(Tuple name value) -> do
      infoDigest <- liftEffect (Wire.hash <$> Wire.write value.info)
      observationsDigest <- liftEffect (Wire.hash <$> Wire.write value.dependencies)
      pure (Tuple name ({ infoDigest, observationsDigest, contribution: value.contribution } :: CompactContribution)))
    entries
  writeSet target "compact" compact
  liftEffect $ Console.log ("[p01] projected=" <> show (Array.length compact))

run :: Aff Unit
run = do
  liftEffect (Metrics.setMemProfileRate 0)
  action <- required "P01_ACTION"
  target <- required "P00_OBJECTS"
  if action == "restore" then restoreSet (Proxy :: Proxy CompactContribution) target "compact"
  else if action == "restore-full" then restoreSet (Proxy :: Proxy PreparedInstantiation) target "contributions"
  else if action == "project" then do
    source <- required "P00_INPUT"
    project source target
  else throwError (error ("Unknown compact action: " <> action))

main :: Effect Unit
main = runAff_ (case _ of
  Right _ -> pure unit
  Left failure -> do
    Console.error (message failure)
    Process.exit' 1) run
