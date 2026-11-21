module Main where

-- @dependencies: console

import Effect.Console (log)

main = log ({hasOwnProperty: "Hi"} {hasOwnProperty = "Done"}).hasOwnProperty
