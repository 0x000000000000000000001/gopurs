module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

data Foo r = Foo { | r }

test :: Foo ()
test = Foo {}

main = log "Done"
