module Main where

-- @dependencies: console prelude

  import Prelude
  import Middle
  import Effect.Console

  main = do
    logShow (middle unit)
    log "Done"
