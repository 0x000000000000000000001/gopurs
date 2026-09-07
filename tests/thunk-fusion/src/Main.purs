module Main where

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Partial.Unsafe (unsafeCrashWith)
import Test.Assert (assertEqual)

-- These producers deliberately use names and inputs unrelated to the benchmark.
-- Run through tests/thunk-fusion/run with the typed PureScript compiler.
suspendAdds :: Int -> Int -> (Unit -> Int) -> Unit -> Int
suspendAdds 0 _ acc = acc
suspendAdds n step acc = suspendAdds (n - 1) step (\_ -> acc unit + step)

runAdds :: Int -> Int -> Int -> Int
runAdds depth seed step = suspendAdds depth step (\_ -> seed) unit

checkAdds :: Int -> Int -> Int -> Effect Unit
checkAdds depth seed step =
  assertEqual { expected: seed + depth * step, actual: runAdds depth seed step }

-- Reversing the updates would give 67 instead of 73 for depth 3, seed 7.
suspendOrder :: Int -> (Unit -> Int) -> Unit -> Int
suspendOrder 0 acc = acc
suspendOrder n acc = suspendOrder (n - 1) (\_ -> 2 * acc unit + n)

runOrder :: Int -> Int -> Int
runOrder depth seed = suspendOrder depth (\_ -> seed) unit

-- All right-hand sides must see the previous iteration's n and step.
suspendVary :: Int -> Int -> (Unit -> Int) -> Unit -> Int
suspendVary 0 _ acc = acc
suspendVary n step acc =
  suspendVary (n - 1) (step + n) (\_ -> acc unit + step * n)

runVary :: Int -> Int -> Int -> Int
runVary depth seed step = suspendVary depth step (\_ -> seed) unit

-- The worker name must be fresh after Go identifier sanitization as well.
suspendClash' :: Int -> (Unit -> Int) -> Unit -> Int
suspendClash' 0 acc = acc
suspendClash' n acc = suspendClash' (n - 1) (\_ -> acc unit + 1)

suspendClash_prime___gopurs_strict_thunk_0 :: Int -> Int
suspendClash_prime___gopurs_strict_thunk_0 x = x + 100

runClash :: Int -> Int -> Int
runClash depth seed = suspendClash' depth (\_ -> seed) unit

keepSuspended :: Int -> Int -> Unit -> Int
keepSuspended depth seed = suspendAdds depth 2 (\_ -> seed)

-- Forcing the result need not demand its predecessor or the initial seed.
suspendConditional :: Int -> Boolean -> (Unit -> Int) -> Unit -> Int
suspendConditional 0 _ acc = acc
suspendConditional n demand acc =
  suspendConditional (n - 1) demand (\_ -> if demand then acc unit + 1 else n)

runConditional :: Int -> Boolean -> Int
runConditional depth demand =
  suspendConditional depth demand (\_ -> unsafeCrashWith "Unused seed was forced") unit

suspendOverwrite :: Int -> (Unit -> Int) -> Unit -> Int
suspendOverwrite 0 acc = acc
suspendOverwrite n _ = suspendOverwrite (n - 1) (\_ -> n)

runOverwrite :: Int -> Int
runOverwrite depth =
  suspendOverwrite depth (\_ -> unsafeCrashWith "Discarded seed was forced") unit

-- Isolate the force-count guards: Int parameters and total seeds would
-- otherwise qualify. The predecessor is demanded conditionally, zero times,
-- or twice respectively.
suspendConditionalInt :: Int -> (Unit -> Int) -> Unit -> Int
suspendConditionalInt 0 acc = acc
suspendConditionalInt n acc =
  suspendConditionalInt (n - 1) (\_ -> if n == 2 then n else acc unit + 1)

runConditionalInt :: Int -> Int -> Int
runConditionalInt depth seed = suspendConditionalInt depth (\_ -> seed) unit

runOverwriteTotal :: Int -> Int -> Int
runOverwriteTotal depth seed = suspendOverwrite depth (\_ -> seed) unit

suspendTwice :: Int -> (Unit -> Int) -> Unit -> Int
suspendTwice 0 acc = acc
suspendTwice n acc = suspendTwice (n - 1) (\_ -> acc unit + acc unit)

runTwice :: Int -> Int -> Int
runTwice depth seed = suspendTwice depth (\_ -> seed) unit

-- Producing or forcing a thunk returning Effect must preserve effect timing.
scheduleEffects :: Int -> (Unit -> Effect Int) -> Unit -> Effect Int
scheduleEffects 0 acc = acc
scheduleEffects n acc = scheduleEffects (n - 1) \_ -> do
  value <- acc unit
  pure (value + 1)

main :: Effect Unit
main = do
  depthRef <- Ref.new 1000
  seedRef <- Ref.new 7
  stepRef <- Ref.new (-3)
  depth <- Ref.read depthRef
  seed <- Ref.read seedRef
  step <- Ref.read stepRef
  checkAdds 0 seed step
  checkAdds 1 seed step
  checkAdds 2 seed step
  checkAdds 7 seed step
  checkAdds depth seed step
  checkAdds 7 step seed
  assertEqual { expected: seed, actual: runOrder 0 seed }
  assertEqual { expected: 73, actual: runOrder 3 seed }
  assertEqual { expected: 30, actual: runVary 3 seed 2 }
  assertEqual { expected: 10, actual: runClash 3 seed }
  assertEqual { expected: 107, actual: suspendClash_prime___gopurs_strict_thunk_0 seed }

  let escaped = keepSuspended 7 seed
  assertEqual { expected: 42, actual: escaped unit + escaped unit }
  assertEqual { expected: 1, actual: runConditional 3 false }
  assertEqual { expected: 1, actual: runOverwrite 3 }
  assertEqual { expected: seed, actual: runConditionalInt 0 seed }
  assertEqual { expected: 8, actual: runConditionalInt 1 seed }
  assertEqual { expected: 3, actual: runConditionalInt 3 seed }
  assertEqual { expected: seed, actual: runOverwriteTotal 0 seed }
  assertEqual { expected: 1, actual: runOverwriteTotal 3 seed }
  assertEqual { expected: seed, actual: runTwice 0 seed }
  assertEqual { expected: 56, actual: runTwice 3 seed }

  calls <- Ref.new 0
  let pending = scheduleEffects 3 \_ -> do
        Ref.modify_ (\count -> count + 1) calls
        pure seed
  before <- Ref.read calls
  assertEqual { expected: 0, actual: before }
  first <- pending unit
  second <- pending unit
  after <- Ref.read calls
  assertEqual { expected: 10, actual: first }
  assertEqual { expected: 10, actual: second }
  assertEqual { expected: 2, actual: after }
  log "Done"
