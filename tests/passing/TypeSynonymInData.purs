module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

type A a = Array a

data Foo a = Foo (A a) | Bar

foo (Foo []) = Bar

main = log "Done"
