module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

data A = A B

type B = A

main = log "Done"
