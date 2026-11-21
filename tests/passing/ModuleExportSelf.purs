module Main where

-- @dependencies: console prelude

import Effect.Console
import A

bar :: Foo
bar = true

main = do
  logShow (show bar)
  log "Done"
