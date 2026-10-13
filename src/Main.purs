module Main (main) where

import Prelude

import Data.Either (Either(..))
import Effect (Effect)
import Effect.Aff (runAff_)
import Effect.Console as Console
import Effect.Exception (message)
import Gopurs.Driver (compile)
import Gopurs.CompilerSnapshot as CompilerSnapshot

foreign import exitFailure :: Effect Unit

main :: Effect Unit
main = CompilerSnapshot.withSnapshot $ runAff_ onComplete compile
  where
  -- Observe the result only after the driver's brackets and supervision have
  -- finished. A discarded Aff fiber does not report failure to the native CLI.
  onComplete result = do
    CompilerSnapshot.release
    case result of
      Right _ -> pure unit
      Left err -> do
        Console.error ("[gopurs] error: " <> message err)
        exitFailure
