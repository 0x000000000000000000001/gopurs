module Main where

-- @dependencies: assert prelude effect console refs

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

identityStep :: (Int -> Int) -> Int -> Int
identityStep _ value = value

repeatStep :: Int -> (Int -> Int) -> Int -> Int
repeatStep 0 = identityStep
repeatStep count =
  let previous = repeatStep (count - 1)
  in \step value -> step (previous step value)

-- This producer captures the changing counter and must keep its original path.
repeatWithCounter :: Int -> (Int -> Int) -> Int -> Int
repeatWithCounter 0 = identityStep
repeatWithCounter count =
  let previous = repeatWithCounter (count - 1)
  in \step value -> step (previous step value) + count

main :: Effect Unit
main = do
  countRef <- Ref.new 4
  count <- Ref.read countRef
  let saved = repeatStep count
  let applied = saved (\value -> value * 3 - 7)
  assertEqual { expected: 611, actual: applied 11 }
  assertEqual { expected: 611, actual: applied 11 }
  assertEqual { expected: 935, actual: applied 15 }
  assertEqual { expected: 19, actual: saved (\value -> value + 2) 11 }
  assertEqual { expected: 11, actual: repeatStep (count - 4) (\_ -> 999) 11 }
  assertEqual { expected: 0, actual: repeatStep count (\value -> negate value) 0 }
  assertEqual { expected: 29, actual: repeatWithCounter count (\value -> value + 2) 11 }
  log "Done"
