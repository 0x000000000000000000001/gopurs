module Main where

-- @dependencies: console

-- No error as we never force `thing` to be resolved in `Main`
import A
import B
import Effect.Console (log)

main = log "Done"
