module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

foreign import value :: Int

main = log "Done"
