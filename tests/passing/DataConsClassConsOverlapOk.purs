module Main where

-- @dependencies: console

import Effect.Console (log)
import Prim.Row (class Cons)

data Cons = Cons

main = log "Done"
