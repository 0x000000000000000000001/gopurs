module Main where

import Prelude
import Effect.Console (log)
import Data.Int (fromNumber)
import Data.Maybe (Maybe(..))

main = do
  case fromNumber (-2147483648.0) of
    Just i -> log ("Just " <> show i)
    Nothing -> log "Nothing"
