module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

data Foo = Foo

instance eqFoo :: Eq Foo where
  eq _ _ = true

main = log "Done"
