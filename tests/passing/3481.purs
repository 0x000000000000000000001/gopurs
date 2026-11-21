module Main where

-- @dependencies: console

import Effect.Console (log)

message = { "0": { "1": "Done" }}

main = log message."0"."1"
