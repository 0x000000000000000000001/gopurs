module Gopurs.Host.Diagnostics (log, recorded) where

import Prelude
import Effect (Effect)
import Effect.Console as Console

foreign import mark :: Effect Unit
foreign import recorded :: Effect Boolean

-- Diagnostics without a replay codec conservatively require a real build.
log :: String -> Effect Unit
log message = mark *> Console.log message
