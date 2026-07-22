module ScratchFFI where

import Prelude
import Effect.Console (log)

main = do
  log $ showFFI { a: 1, b: true, c: 'd', e: 4.0 }

showFFI :: forall a. Show a => a -> String
showFFI = showImpl show

foreign import showImpl :: forall a. (a -> String) -> a -> String
