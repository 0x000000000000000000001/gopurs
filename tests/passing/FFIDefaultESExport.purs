module Main where

-- @dependencies: console

import Effect.Console (log)

foreign import default :: String

main = log default
