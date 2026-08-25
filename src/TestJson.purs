module TestJson where

import Prelude
import Data.Either (Either(..))
import Data.Argonaut.Parser (jsonParser)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Effect (Effect)
import Data.Bifunctor (lmap)
import Effect.Console as Console
import Node.FS.Sync as FS
import Node.Encoding (Encoding(..))
import PureScript.Backend.Optimizer.CoreFn.Json (decodeModule)

main :: Effect Unit
main = do
  content <- FS.readTextFile UTF8 "output/Main/corefn.json"
  case jsonParser content >>= (lmap printJsonDecodeError <<< decodeModule) of
    Left err -> Console.log err
    Right _ -> Console.log "Success"
