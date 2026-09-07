module Main where

-- @dependencies: assert prelude effect console refs functions

import Prelude

import Data.Function.Uncurried (Fn2, mkFn2, runFn2)
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

-- Keep these boundaries inside returned closures, where collectCurriedAbs runs.
-- Each producer captures a computed repeater before returning its first lambda.
-- The snapshot must keep the calculation, branch, recursive bindings or call
-- in that unary closure, before constructing or returning the next function.
makeLetBoundary :: Int -> Int -> Int -> Int
makeLetBoundary count =
  let previous = repeatApply count
  in \a ->
    let cached = previous squarePlusOne a
    in \b -> cached + cached + b

makeBranchBoundary :: Int -> Int -> Int -> Int
makeBranchBoundary count =
  let previous = repeatApply count
  in \a ->
    if a < 0 then
      \b -> previous (\value -> value + 1) (a - b)
    else
      \b -> previous (\value -> value + 1) (a + 2 * b)

makeRecursiveBoundary :: Int -> Int -> Int -> Int
makeRecursiveBoundary count =
  let previous = repeatApply count
  in \a ->
    let
      visit n =
        if n == 0 then previous (\value -> value + 1) a
        else visit (n - 1) + 1
    in \b -> visit b + visit (b + 1)

makeCallBoundary :: Int -> (Int -> Int -> Int) -> Int -> Int -> Int
makeCallBoundary count factory =
  let previous = repeatApply count
  in \a -> factory (previous (\value -> value + 1) a)

makeUncurriedBoundary :: Int -> Int -> Fn2 Int Int Int
makeUncurriedBoundary count =
  let previous = repeatApply count
  -- The snapshot must keep Func returning Func2, not a single Func3.
  in \a -> mkFn2 (\b c -> previous (\value -> value + 1) (a + 10 * b + 100 * c))

-- A visible type application on the returned polymorphic lambda preserves
-- TypeApp between the two Abs nodes; the snapshot must keep two unary Funcs.
makeTypeAppBoundary :: Int -> Int -> Int -> { left :: Int, right :: Int }
makeTypeAppBoundary count =
  let previous = repeatApply count
  in \a ->
    ((\b -> { left: previous (\value -> value + 1) a, right: b })
      :: forall @value. value -> { left :: Int, right :: value }) @Int

checkBoundaries :: Int -> Effect Unit
checkBoundaries count = do
  letRef <- Ref.new (makeLetBoundary count)
  withLet <- Ref.read letRef
  cachedRef <- Ref.new (withLet 1)
  cached <- Ref.read cachedRef
  assertEqual { expected: 54, actual: cached 2 }
  assertEqual { expected: 59, actual: cached 7 }
  assertEqual { expected: 1353, actual: withLet 2 (-1) }

  branchRef <- Ref.new (makeBranchBoundary count)
  branch <- Ref.read branchRef
  negativeRef <- Ref.new (branch (-2))
  negative <- Ref.read negativeRef
  assertEqual { expected: -4, actual: negative 5 }
  assertEqual { expected: 1, actual: negative 0 }
  assertEqual { expected: 15, actual: branch 2 5 }

  recursiveRef <- Ref.new (makeRecursiveBoundary count)
  recursive <- Ref.read recursiveRef
  visitRef <- Ref.new (recursive 7)
  visit <- Ref.read visitRef
  assertEqual { expected: 25, actual: visit 2 }
  assertEqual { expected: 21, actual: visit 0 }

  factoryRef <- Ref.new (\a b -> 100 * a + b)
  factory <- Ref.read factoryRef
  callRef <- Ref.new (makeCallBoundary count factory)
  call <- Ref.read callRef
  calledRef <- Ref.new (call 2)
  called <- Ref.read calledRef
  assertEqual { expected: 507, actual: called 7 }
  assertEqual { expected: 499, actual: called (-1) }

  uncurriedRef <- Ref.new (makeUncurriedBoundary count)
  uncurried <- Ref.read uncurriedRef
  pairRef <- Ref.new (uncurried 1)
  pair <- Ref.read pairRef
  assertEqual { expected: 324, actual: runFn2 pair 2 3 }
  assertEqual { expected: 234, actual: runFn2 pair 3 2 }
  assertEqual { expected: -176, actual: runFn2 (uncurried 1) 2 (-2) }

  typeAppRef <- Ref.new (makeTypeAppBoundary count)
  typeApp <- Ref.read typeAppRef
  recordRef <- Ref.new (typeApp 7)
  record <- Ref.read recordRef
  assertEqual { expected: { left: 10, right: 2 }, actual: record 2 }
  assertEqual { expected: { left: 10, right: -5 }, actual: record (-5) }

makeEffectBoundary :: Int -> Ref.Ref Int -> Int -> Effect (Int -> Int)
makeEffectBoundary count state =
  let previous = repeatApply count
  in \a -> do
    offset <- Ref.read state
    Ref.write (offset + 1) state
    pure (\b -> previous (\value -> value + 1) (a + offset) + b)

checkEffectBoundary :: Int -> Effect Unit
checkEffectBoundary count = do
  state <- Ref.new 10
  makerRef <- Ref.new (makeEffectBoundary count state)
  maker <- Ref.read makerRef
  pendingRef <- Ref.new (maker 1)
  pending <- Ref.read pendingRef
  beforeRun <- Ref.read state
  assertEqual { expected: 10, actual: beforeRun }
  -- The read happens when the Effect runs, and its result is then captured.
  Ref.write 20 state
  first <- pending
  afterRun <- Ref.read state
  assertEqual { expected: 21, actual: afterRun }
  Ref.write 100 state
  assertEqual { expected: 26, actual: first 2 }
  assertEqual { expected: 29, actual: first 5 }
  afterCalls <- Ref.read state
  assertEqual { expected: 100, actual: afterCalls }
  second <- pending
  afterSecondRun <- Ref.read state
  assertEqual { expected: 101, actual: afterSecondRun }
  assertEqual { expected: 106, actual: second 2 }

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
  checkBoundaries smallDepth
  checkEffectBoundary smallDepth
  log "Done"
