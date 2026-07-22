module ScratchDebug where

import Prelude
import Effect.Console (log)

main = do
  log $ show 4.0
  log $ show { a: 1, b: true, c: 'd', e: 4.0 }
