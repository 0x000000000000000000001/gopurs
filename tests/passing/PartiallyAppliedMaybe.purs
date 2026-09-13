-- @inline export addMaybe never
-- @inline export partialMaybe never
module Main where

-- @dependencies: assert prelude effect console refs maybe

import Prelude

import Data.Maybe (Maybe(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

addMaybe :: Int -> Int -> Maybe Int
addMaybe left right =
  if right < 0 then Nothing else Just (left + right)

-- The branch keeps one outer lambda whose result is another function.
-- Its final result type is Maybe, but its immediate result is a closure.
partialMaybe :: Int -> Int -> Maybe Int
partialMaybe left =
  if left < 0 then addMaybe (-left) else addMaybe left

main :: Effect Unit
main = do
  seedRef <- Ref.new 20
  seed <- Ref.read seedRef
  saved <- Ref.new (partialMaybe seed)
  partial <- Ref.read saved
  assertEqual { expected: Just 42, actual: partial 22 }
  assertEqual { expected: Just 20, actual: partial 0 }
  assertEqual { expected: Nothing, actual: partial (-1) }
  negativeSaved <- Ref.new (partialMaybe (-seed))
  negativePartial <- Ref.read negativeSaved
  assertEqual { expected: Just 23, actual: negativePartial 3 }
  log "Done"
