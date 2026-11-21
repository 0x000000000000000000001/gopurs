module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

data Foo a = Foo

foo = \f -> case f of Foo -> "foo"

main = log "Done"
