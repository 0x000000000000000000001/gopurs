module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

foo :: forall a. {b :: Number | a} -> {b :: Number | _}
foo f = f

main = log "Done"
