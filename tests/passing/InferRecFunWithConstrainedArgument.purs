module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log, logShow)

test 100 = 100
test n = test(1 + n)

main = do
  logShow (test 0)
  log "Done"
