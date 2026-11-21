module Main where

-- @dependencies: console prelude

import Prelude

import Effect.Console (log)

class Convert a b | a -> b where
  convert :: a -> b

type Words = String

instance convertSB :: Convert Int Words where
  convert 0 = "Nope"
  convert _ = "Done"

main = log $ convert 1
