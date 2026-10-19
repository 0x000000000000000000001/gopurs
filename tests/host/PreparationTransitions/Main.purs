module Main where

-- @dependencies: prelude effect console refs

import Prelude
import Effect (Effect)
import Effect.Console (log, logShow)
import Effect.Ref as Ref
import Preparation.Calls as Calls

foreign import tweak :: Int -> Int

main :: Effect Unit
main = do
  cell <- Ref.new 5
  seed <- Ref.read cell
  logShow (Calls.run seed)
  logShow (Calls.provided seed)
  logShow (tweak seed)
  log "Done"
