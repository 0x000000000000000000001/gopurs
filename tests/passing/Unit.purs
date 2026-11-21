module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (logShow, log)

main = do
  logShow (const unit $ "Hello world")
  log "Done"
