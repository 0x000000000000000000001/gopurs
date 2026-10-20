module Gopurs.Resident (run) where

import Prelude

import Data.Either (Either(..))
import Data.Maybe (Maybe(..))
import Data.Time.Duration (Milliseconds(..))
import Effect.Aff (Aff, attempt, delay)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Exception (message)
import Gopurs.Driver (compileWithConfig)
import Gopurs.Driver.Config (CompilerConfig)
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import Node.Process as Process

-- C2-2 lifecycle only: the process and config stay alive, and a trigger file
-- requests a build. Backend state reuse is deliberately not implemented here.
run :: CompilerConfig -> Aff Unit
run config = do
  trigger <- liftEffect (Process.lookupEnv "GOPURS_RESIDENT_TRIGGER")
  case trigger of
    Nothing -> liftEffect $ Console.error "[gopurs] resident: GOPURS_RESIDENT_TRIGGER is required"
    Just path -> do
      liftEffect $ Console.error ("[gopurs] resident: watching " <> path)
      loop config path ""

loop :: CompilerConfig -> String -> String -> Aff Unit
loop config path previous = do
  delay (Milliseconds 100.0)
  current <- attempt (FS.readTextFile UTF8 path)
  case current of
    Left _ -> loop config path previous
    Right text
      | text == previous -> loop config path previous
      | otherwise -> do
          result <- attempt (compileWithConfig config)
          case result of
            Left err -> liftEffect $ Console.error ("[gopurs] resident build failed: " <> message err)
            Right _ -> liftEffect $ Console.error "[gopurs] resident build complete"
          loop config path text
