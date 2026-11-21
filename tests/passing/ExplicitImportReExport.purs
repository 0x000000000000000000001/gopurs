-- from #1244
module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)
import Bar (foo)

baz :: Int
baz = foo

main = log "Done"
