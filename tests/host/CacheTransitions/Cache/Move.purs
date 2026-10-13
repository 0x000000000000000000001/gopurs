module Cache.Move where

import Prelude
import Effect (Effect)
import Effect.Console (log)

foreign import echo :: Int -> Int

value :: Int -> Int
value n = echo n

main :: Effect Unit
main = log "secondary"
