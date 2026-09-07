module Main where

-- @dependencies: assert prelude effect console refs

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

type Repeater = (Int -> Int) -> Int -> Int

-- Construct the previous function before returning the two annotated lambdas.
-- This keeps the returned closures visible after PBO optimization.
repeatApply :: Int -> Repeater
repeatApply 0 = \_ seed -> seed
repeatApply count =
  let previous = repeatApply (count - 1)
  in (\step -> ((\seed -> step (previous step seed)) :: Int -> Int)) :: Repeater

composeRepeats :: Int -> Int -> Repeater
composeRepeats leftCount rightCount =
  let
    left = repeatApply leftCount
    right = repeatApply rightCount
  in (\step -> ((\seed -> left step (right step seed)) :: Int -> Int)) :: Repeater

checkRepeat :: Int -> Int -> Int -> Effect Unit
checkRepeat count step seed =
  assertEqual
    { expected: seed + count * step
    , actual: repeatApply count (\value -> value + step) seed
    }

main :: Effect Unit
main = do
  -- Runtime inputs prevent the assertions from becoming constant expressions.
  depthRef <- Ref.new 5
  seedRef <- Ref.new 7
  stepRef <- Ref.new (-3)
  depth <- Ref.read depthRef
  seed <- Ref.read seedRef
  step <- Ref.read stepRef
  checkRepeat 0 step seed
  checkRepeat 1 step seed
  checkRepeat 2 step seed
  checkRepeat depth step seed
  checkRepeat depth 0 seed
  checkRepeat depth 2 0
  checkRepeat 3 2 (-4)
  assertEqual
    { expected: seed + (depth + 2) * step
    , actual: composeRepeats depth 2 (\value -> value + step) seed
    }
  assertEqual
    { expected: seed
    , actual: composeRepeats 0 0 (\value -> value + step) seed
    }
  log "Done"
