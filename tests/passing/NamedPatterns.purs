module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

foo = \x -> case x of
  y@{ foo: "Foo" } -> y
  y -> y

main = log "Done"
