module Main where

-- @dependencies: console effect prelude

import Prelude
import Effect
import M1
import Effect.Console as C

main = C.log (log "Done")
