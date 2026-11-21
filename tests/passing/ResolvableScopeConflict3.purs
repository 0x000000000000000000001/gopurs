module Main (thing, main, module A) where

-- @dependencies: console

import A
import Effect.Console (log)

thing :: Int
thing = 2

main = log "Done"
