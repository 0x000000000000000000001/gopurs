module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console

id' = (\x -> x) <$> \y -> y

main = log (id' "Done")
