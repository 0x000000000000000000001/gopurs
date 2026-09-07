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

squarePlusOne :: Int -> Int
squarePlusOne value = value * value + 1

checkPartialApplications :: Int -> Effect Unit
checkPartialApplications count = do
  assertEqual { expected: 26, actual: repeatApply count squarePlusOne 1 }
  assertEqual { expected: 677, actual: repeatApply count squarePlusOne 2 }
  -- Ref keeps the partial application from being inlined at each call site.
  partialRef <- Ref.new (repeatApply count squarePlusOne)
  partial <- Ref.read partialRef
  assertEqual { expected: 26, actual: partial 1 }
  assertEqual { expected: 677, actual: partial 2 }
  assertEqual { expected: 5, actual: partial 0 }

type EffectRepeater = (Int -> Effect Int) -> Int -> Effect Int

repeatEffects :: Int -> EffectRepeater
repeatEffects 0 = \_ seed -> pure seed
repeatEffects count =
  let previous = repeatEffects (count - 1)
  in (\step -> ((\seed -> previous step seed >>= step) :: Int -> Effect Int)) :: EffectRepeater

checkEffectOrder :: Int -> Effect Unit
checkEffectOrder count = do
  trace <- Ref.new ""
  let step value = do
        Ref.modify_ (\seen -> seen <> show value <> ",") trace
        pure (squarePlusOne value)
  saturated <- repeatEffects count step 1
  saturatedTrace <- Ref.read trace
  assertEqual { expected: 26, actual: saturated }
  assertEqual { expected: "1,2,5,", actual: saturatedTrace }
  Ref.write "" trace

  partialRef <- Ref.new (repeatEffects count step)
  partial <- Ref.read partialRef
  before <- Ref.read trace
  assertEqual { expected: "", actual: before }
  -- Constructing an Effect value must not run it, even after all Int arguments.
  pendingRef <- Ref.new (partial 1)
  pending <- Ref.read pendingRef
  beforeRun <- Ref.read trace
  assertEqual { expected: "", actual: beforeRun }
  first <- pending
  firstTrace <- Ref.read trace
  assertEqual { expected: 26, actual: first }
  assertEqual { expected: "1,2,5,", actual: firstTrace }
  second <- partial 2
  secondTrace <- Ref.read trace
  assertEqual { expected: 677, actual: second }
  assertEqual { expected: "1,2,5,2,5,26,", actual: secondTrace }

type ThreeArgs = Int -> Int -> Int -> Int
type SixArgs = Int -> Int -> Int -> Int -> Int -> Int -> Int

-- Build the captured repeater before returning the argument chain, so PBO
-- retains the returned closures. Distinct weights detect swapped arguments.
makeThree :: Int -> ThreeArgs
makeThree count =
  let previous = repeatApply count
  in (\a b -> ((\c -> previous (\value -> value + 1) (a + 10 * b + 100 * c)) :: Int -> Int)) :: ThreeArgs

makeSix :: Int -> SixArgs
makeSix count =
  let previous = repeatApply count
  in (\a b c d e -> ((\f -> previous (\value -> value + 1)
    (a + 10 * b + 100 * c + 1000 * d + 10000 * e + 100000 * f)) :: Int -> Int)) :: SixArgs

checkArities :: Int -> Effect Unit
checkArities count = do
  threeRef <- Ref.new (makeThree count)
  three <- Ref.read threeRef
  assertEqual { expected: count + 321, actual: three 1 2 3 }
  threePartialRef <- Ref.new (three 4)
  threePartial <- Ref.read threePartialRef
  assertEqual { expected: count + 654, actual: threePartial 5 6 }
  assertEqual { expected: count + 34, actual: threePartial 3 0 }

  sixRef <- Ref.new (makeSix count)
  six <- Ref.read sixRef
  assertEqual { expected: count + 654321, actual: six 1 2 3 4 5 6 }
  -- Exercise partial application before the Func5 boundary, then across it.
  twoRef <- Ref.new (six 1 2)
  two <- Ref.read twoRef
  assertEqual { expected: count + 654321, actual: two 3 4 5 6 }
  assertEqual { expected: count + 345621, actual: two 6 5 4 3 }
  -- Applying five arguments returns a reusable unary function for the sixth.
  fiveRef <- Ref.new (six 1 2 3 4 5)
  five <- Ref.read fiveRef
  assertEqual { expected: count + 654321, actual: five 6 }
  assertEqual { expected: count + 954321, actual: five 9 }

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
  smallDepthRef <- Ref.new 3
  smallDepth <- Ref.read smallDepthRef
  checkPartialApplications smallDepth
  checkEffectOrder smallDepth
  checkArities smallDepth
  log "Done"
