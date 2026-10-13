module Main where

-- @dependencies: prelude effect console refs

import Prelude
import Cache.Body as Body
import Cache.Calls as Calls
import Cache.Dictionary as Dictionary
import Cache.Move as Move
import Cache.Shapes as Shapes
import Effect (Effect)
import Effect.Console (log, logShow)
import Effect.Ref as Ref

main :: Effect Unit
main = do
  cell <- Ref.new 5
  seed <- Ref.read cell
  logShow (Body.value seed)
  logShow (Shapes.unpack (Shapes.pack seed))
  logShow (Shapes.scorePacket (Shapes.packet seed))
  logShow (Dictionary.measure seed)
  logShow (Calls.value seed)
  logShow (Move.value seed)
  log "Done"
