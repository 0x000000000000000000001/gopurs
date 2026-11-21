-- No failure here as the export `thing` only refers to Main.thing
module Main (thing, main) where

-- @dependencies: console

import A
import Effect.Console (log)

thing :: Int
thing = 2

main = log "Done"
