module Test.Main where

import Prelude
import Effect.Console (logShow)

main = do
  let a = [1]
  let b = [1, 2]
  logShow (length a < length b)
