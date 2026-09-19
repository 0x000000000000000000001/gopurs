module Main where

-- @dependencies: assert prelude effect console refs

import Prelude
import Effect
import Effect.Console
import Effect.Ref as Ref
import Test.Assert

checkOrdering :: forall a. Ord a => a -> a -> Ordering -> Effect Unit
checkOrdering left right expected = do
  -- Keep the operands dynamic so the assertions exercise native comparisons.
  leftRef <- Ref.new left
  rightRef <- Ref.new right
  x <- Ref.read leftRef
  y <- Ref.read rightRef
  assertEqual { expected, actual: compare x y }
  assertEqual { expected: expected == LT, actual: x < y }
  assertEqual { expected: expected /= GT, actual: x <= y }
  assertEqual { expected: expected == EQ, actual: x == y }
  assertEqual { expected: expected /= EQ, actual: x /= y }
  assertEqual { expected: expected /= LT, actual: x >= y }
  assertEqual { expected: expected == GT, actual: x > y }

  comparatorRef <- Ref.new compare
  comparator <- Ref.read comparatorRef
  assertEqual { expected, actual: comparator x y }

  partialRef <- Ref.new (comparator x)
  partial <- Ref.read partialRef
  assertEqual { expected, actual: partial y }
  assertEqual { expected: EQ, actual: partial x }
  assertEqual { expected, actual: partial y }

  lessThanRef <- Ref.new ((<) x)
  lessThan <- Ref.read lessThanRef
  assertEqual { expected: expected == LT, actual: lessThan y }
  assertEqual { expected: false, actual: lessThan x }
  assertEqual { expected: expected == LT, actual: lessThan y }

main :: Effect Unit
main = do
  assert (1.0 < 2.0)
  assert (2.0 == 2.0)
  assert (3.0 > 1.0)
  assert ("a" < "b")
  assert ("a" == "a")
  assert ("z" > "a")
  checkOrdering 'a' 'b' LT
  checkOrdering 'é' 'é' EQ
  checkOrdering 'Ā' 'ÿ' GT
  checkOrdering 'ÿ' 'Ā' LT
  checkOrdering "a" "b" LT
  checkOrdering "é" "é" EQ
  checkOrdering "éa" "é" GT
  checkOrdering "é" "éa" LT
  checkOrdering "😀" "😁" LT
  checkOrdering "😁" "😀" GT
  log "Done"
