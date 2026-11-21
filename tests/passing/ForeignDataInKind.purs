module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

foreign import data A :: Type
data B (x :: A)

main = log "Done"
