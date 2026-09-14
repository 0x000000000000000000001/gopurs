module Main where

-- @dependencies: assert prelude effect console refs

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

-- The thunk is hidden behind a polymorphic newtype, as in LazyEvaluation.
-- ThunkFusion.purs covers the remaining guards for bare Unit -> Int functions.
newtype Lazy a = Lazy (Unit -> a)

defer :: forall a. (Unit -> a) -> Lazy a
defer = Lazy

force :: forall a. Lazy a -> a
force (Lazy thunk) = thunk unit

wrapIncrements :: Int -> Lazy Int -> Lazy Int
wrapIncrements 0 acc = acc
wrapIncrements n acc = wrapIncrements (n - 1) (defer \_ -> force acc + 1)

runIncrements :: Int -> Int -> Int
runIncrements depth seed = force (wrapIncrements depth (defer \_ -> seed))

wrapAdds :: Int -> Int -> Lazy Int -> Lazy Int
wrapAdds 0 _ acc = acc
wrapAdds n step acc = wrapAdds (n - 1) step (defer \_ -> force acc + step)

runAdds :: Int -> Int -> Int -> Int
runAdds depth seed step = force (wrapAdds depth step (defer \_ -> seed))

checkAdds :: Int -> Int -> Int -> Effect Unit
checkAdds depth seed step =
  assertEqual { expected: seed + depth * step, actual: runAdds depth seed step }

-- The captured countdown values must be applied in their original order.
-- For depth 3 and seed 7, reversing them gives 67 instead of 73.
wrapOrder :: Int -> Lazy Int -> Lazy Int
wrapOrder 0 acc = acc
wrapOrder n acc = wrapOrder (n - 1) (defer \_ -> 2 * force acc + n)

runOrder :: Int -> Int -> Int
runOrder depth seed = force (wrapOrder depth (defer \_ -> seed))

keepWrapped :: Int -> Int -> Lazy Int
keepWrapped depth seed = wrapIncrements depth (defer \_ -> seed)

-- A thunk yielding an Effect remains outside the pure Int fusion rule.
wrapEffects :: Int -> Lazy (Effect Int) -> Lazy (Effect Int)
wrapEffects 0 acc = acc
wrapEffects n acc = wrapEffects (n - 1) (defer \_ -> do
  value <- force acc
  pure (value + 1))

main :: Effect Unit
main = do
  depthRef <- Ref.new 1000
  seedRef <- Ref.new 7
  stepRef <- Ref.new (-3)
  depth <- Ref.read depthRef
  seed <- Ref.read seedRef
  step <- Ref.read stepRef
  assertEqual { expected: seed, actual: runIncrements 0 seed }
  assertEqual { expected: seed + 1, actual: runIncrements 1 seed }
  assertEqual { expected: seed + depth, actual: runIncrements depth seed }
  checkAdds 0 seed step
  checkAdds 1 seed step
  checkAdds 7 seed step
  checkAdds depth seed step
  checkAdds 7 step seed
  assertEqual { expected: seed, actual: runOrder 0 seed }
  assertEqual { expected: 15, actual: runOrder 1 seed }
  assertEqual { expected: 73, actual: runOrder 3 seed }

  -- Storing the producer's result makes the Lazy value escape its consumer.
  escapedRef <- Ref.new (keepWrapped depth seed)
  escaped <- Ref.read escapedRef
  assertEqual { expected: 2 * (seed + depth), actual: force escaped + force escaped }

  calls <- Ref.new 0
  let pending = wrapEffects 3 (defer \_ -> do
        Ref.modify_ (\count -> count + 1) calls
        pure seed)
  before <- Ref.read calls
  assertEqual { expected: 0, actual: before }
  first <- force pending
  second <- force pending
  after <- Ref.read calls
  assertEqual { expected: 10, actual: first }
  assertEqual { expected: 10, actual: second }
  assertEqual { expected: 2, actual: after }
  log "Done"
