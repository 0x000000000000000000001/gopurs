module Main where

-- @dependencies: console

import Test
import Effect.Console

data Test = Test

main = log (runZ (Z "Done"))
