module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

data X = X Int | Y

x :: X -> Int
x = case _ of
      Y -> 0
      X n | 1 <- n -> 1
          | otherwise -> 2

main = log "Done"
