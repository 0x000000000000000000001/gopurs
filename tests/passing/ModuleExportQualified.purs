module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log, logShow)
import A as B

main = do
  logShow (B.show 1.0)
  log "Done"
