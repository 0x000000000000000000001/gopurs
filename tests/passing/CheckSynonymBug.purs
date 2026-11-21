module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

length :: forall a. Array a -> Int
length _ = 0

type Foo a = Array a

foo _ = length ([] :: Foo Number)

main = log "Done"
