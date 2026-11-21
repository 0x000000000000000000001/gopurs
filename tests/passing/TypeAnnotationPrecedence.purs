-- See #3554
module Main where

-- @dependencies: console effect prelude tuples

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Data.Tuple (Tuple(..), uncurry)

appendAndLog = log <<< uncurry append :: Tuple String String -> Effect Unit

main = appendAndLog (Tuple "Do" "ne")
