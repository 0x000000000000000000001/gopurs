-- @inline export accumulate never
-- @inline export choose never
-- @inline export shiftState never
-- @inline export divideState never
-- @inline export literalState never
-- @inline export literalStep never
-- @inline export literalCompare never
module Main where

-- @dependencies: assert prelude effect console refs integers

import Prelude
import Data.Int.Bits (zshr)
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

accumulate :: Int -> Int -> Int -> Int
accumulate 0 seed _ = seed
accumulate n seed step = accumulate (n - 1) (seed + step) step

choose :: Int -> Int -> Int -> Int
choose 0 seed _ = seed
choose n seed step =
  if seed < 0 then choose (n - 1) (seed - step) step
  else choose (n - 1) (seed + step) step

-- The count may narrow independently of an accumulator which becomes unsigned.
shiftState :: Int -> Int -> Int
shiftState 0 seed = seed
shiftState n seed = shiftState (n - 1) (zshr seed 0)

divideState :: Int -> Int -> Int
divideState 0 seed = seed
divideState n seed = divideState (n - 1) (seed / (-1))

-- PBO may fold these to unsigned literals; they are not signed-range proofs.
literalState :: Int -> Int -> Int
literalState 0 seed = seed
literalState n _ = literalState (n - 1) (zshr (-1) 0)

literalStep :: Int -> Int -> Int
literalStep 0 seed = seed
literalStep n seed = literalStep (n - 1) (seed + zshr (-1) 0)

literalCompare :: Int -> Int -> Int
literalCompare 0 seed = seed
literalCompare n seed =
  if seed < zshr (-1) 0 then literalCompare (n - 1) (seed + 1)
  else literalCompare (n - 1) (seed - 1)

localLoop :: Int -> Int -> Int
localLoop count initial =
  let
    go 0 seed = seed
    go n seed = go (n - 1) (seed + n)
  in go count initial

main :: Effect Unit
main = do
  input <- Ref.new { count: 3, high: top :: Int, low: bottom :: Int, negative: -1 }
  { count, high, low, negative } <- Ref.read input
  let unsigned = zshr negative 0
  let quotient = low / negative
  assertEqual { expected: low + 2, actual: accumulate count high 1 }
  assertEqual { expected: high - 2, actual: accumulate count low negative }
  assertEqual { expected: unsigned, actual: accumulate 0 unsigned 1 }
  assertEqual { expected: quotient, actual: accumulate 0 quotient 1 }
  assertEqual { expected: 2, actual: accumulate count unsigned 1 }
  assertEqual { expected: low + count, actual: accumulate count quotient 1 }
  assertEqual { expected: 0, actual: accumulate 1 1 unsigned }
  assertEqual { expected: low, actual: choose count high 1 }
  assertEqual { expected: unsigned, actual: shiftState count negative }
  assertEqual { expected: quotient, actual: divideState 1 low }
  assertEqual { expected: low, actual: divideState 2 low }
  assertEqual { expected: unsigned, actual: literalState count 0 }
  assertEqual { expected: high - count, actual: literalStep count high }
  assertEqual { expected: low + 2, actual: literalCompare count high }
  assertEqual { expected: low + 5, actual: localLoop count high }
  assertEqual { expected: unsigned, actual: localLoop 0 unsigned }
  log "Done"
