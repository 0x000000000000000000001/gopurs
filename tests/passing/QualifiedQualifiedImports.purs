module Main where

-- @dependencies: console

-- qualified import with qualified imported names
import Effect.Console (log) as Console

main = Console.log "Done"
