module Gopurs.Resident (run) where

import Prelude

import Data.Either (Either(..))
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.Time.Duration (Milliseconds(..))
import Effect.Aff (Aff, attempt, delay)
import Effect.Class (liftEffect)
import Effect.Ref as Ref
import Data.Map as Map
import Effect.Console as Console
import Effect.Exception (message)
import Gopurs.Driver (compileWithConfigCache)
import Gopurs.Driver.Config (CompilerConfig)
import Gopurs.Driver.Prepare (PreparedCache)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import Node.Process as Process

-- C2-2 lifecycle only: the process and config stay alive, and a trigger file
-- requests a build. The same loop serves the opt-in one-shot daemon: the
-- launcher writes a token to the trigger and waits for that token in the
-- status file. No filesystem watching is involved.
run :: CompilerConfig -> Aff Unit
run config = do
  trigger <- liftEffect (Process.lookupEnv "GOPURS_RESIDENT_TRIGGER")
  status <- liftEffect (Process.lookupEnv "GOPURS_RESIDENT_STATUS")
  case trigger of
    Nothing -> liftEffect $ Console.error "[gopurs] resident: GOPURS_RESIDENT_TRIGGER is required"
    Just path -> do
      liftEffect $ Console.error ("[gopurs] resident: watching " <> path)
      cache <- liftEffect (Ref.new Nothing)
      backends <- liftEffect (Ref.new Map.empty)
      ffiDigests <- liftEffect (Ref.new Map.empty)
      loop config cache backends ffiDigests status path ""

loop :: CompilerConfig -> Ref.Ref (Maybe PreparedCache) -> Ref.Ref (Map.Map String BackendModule) -> Ref.Ref (Map.Map String String) -> Maybe String -> String -> String -> Aff Unit
loop config cache backends ffiDigests status path previous = do
  delay (Milliseconds 100.0)
  current <- attempt (FS.readTextFile UTF8 path)
  case current of
    Left _ -> loop config cache backends ffiDigests status path previous
    Right text
      | text == previous -> loop config cache backends ffiDigests status path previous
      | otherwise -> do
          result <- attempt (compileWithConfigCache config cache backends ffiDigests)
          case result of
            Left err -> do
              liftEffect $ Console.error ("[gopurs] resident build failed: " <> message err)
              publishStatus status text ("error\n" <> message err)
            Right _ -> do
              liftEffect $ Console.error "[gopurs] resident build complete"
              publishStatus status text "ok"
          loop config cache backends ffiDigests status path text

publishStatus :: Maybe String -> String -> String -> Aff Unit
publishStatus Nothing _ _ = pure unit
publishStatus (Just file) trigger outcome = do
  let token = String.trim trigger
  void $ attempt do
    FS.writeTextFile UTF8 (file <> ".next") (token <> "\n" <> outcome <> "\n")
    FS.rename (file <> ".next") file
