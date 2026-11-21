module Main where

-- @dependencies: console

import Effect.Console (log)

test :: forall a. a -> a
test = \(x :: a) -> x

main = log "Done"
