module Main where

import Prelude
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref
import Subject as Subject

main :: Effect Unit
main = do
  source <- Ref.new { total: 2147483647, borrowed: [ 7, 11 ], stamp: "retained" }
  counts <- Ref.new [ 0, 1, 2, 7, 1000 ]
  values <- Ref.read counts
  for_ values \count -> do
    input <- Ref.read source
    logShow [ (Subject.step count input).total, (Subject.normalize count input).total, (Subject.down count input).total ]
  input <- Ref.read source
  logShow input.total
  logShow input.borrowed
  logShow input.stamp
