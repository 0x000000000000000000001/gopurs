module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

shout = log <<< (_ <> "!") <<< show

main = do
  shout "Test"
  log "Done"
