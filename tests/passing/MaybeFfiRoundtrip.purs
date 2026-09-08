-- @inline export makeMaybe never
module Main where

-- @dependencies: assert arrays prelude effect console refs maybe

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

makeMaybe :: Boolean -> Int -> Maybe Int
makeMaybe present value = if present then Just value else Nothing

checkMaybe :: String -> String -> Maybe Int -> Effect Unit
checkMaybe label expected value = do
  -- Cross a polymorphic FFI boundary before inspecting the constructor/payload.
  saved <- Ref.new value
  actual <- Ref.read saved
  assertEqual { expected, actual: show actual }
  log label

checkFindIndex :: String -> String -> (Int -> Boolean) -> Array Int -> Effect Unit
checkFindIndex label expected predicate values = do
  saved <- Ref.new values
  input <- Ref.read saved
  checkMaybe label expected (Array.findIndex predicate input)

checkSpan :: String -> (Int -> Boolean) -> Array Int -> Array Int -> Array Int -> Effect Unit
checkSpan label predicate values expectedInit expectedRest = do
  saved <- Ref.new values
  input <- Ref.read saved
  let actual = Array.span predicate input
  assertEqual { expected: expectedInit, actual: actual.init }
  assertEqual { expected: expectedRest, actual: actual.rest }
  log label

main :: Effect Unit
main = do
  checkMaybe "Nothing roundtrip" "Nothing" Nothing
  checkMaybe "Just zero roundtrip" "(Just 0)" (Just 0)
  checkMaybe "Just payload roundtrip" "(Just 7)" (Just 7)

  savedCases <- Ref.new { absent: false, present: true, zero: 0, payload: 7 }
  cases <- Ref.read savedCases
  checkMaybe "constructed Nothing" "Nothing" (makeMaybe cases.absent cases.payload)
  checkMaybe "constructed Just zero" "(Just 0)" (makeMaybe cases.present cases.zero)
  checkMaybe "constructed Just payload" "(Just 7)" (makeMaybe cases.present cases.payload)

  checkFindIndex "empty, false" "Nothing" (\_ -> false) []
  checkFindIndex "empty, true" "Nothing" (\_ -> true) []
  checkFindIndex "singleton, false" "Nothing" (\_ -> false) [0]
  checkFindIndex "singleton, true" "(Just 0)" (\_ -> true) [0]
  checkFindIndex "later match" "(Just 2)" (_ == 3) [1, 2, 3]
  checkFindIndex "no match" "Nothing" (_ == 4) [1, 2, 3]

  checkSpan "span empty" (\_ -> true) [] [] []
  checkSpan "span singleton all" (\_ -> true) [1] [1] []
  checkSpan "span singleton none" (\_ -> false) [1] [] [1]
  checkSpan "span all" (\_ -> true) [1, 2, 3] [1, 2, 3] []
  checkSpan "span prefix" (_ < 3) [1, 2, 3] [1, 2] [3]

  saved <- Ref.new (Array.range 1 10000)
  input <- Ref.read saved
  let result = Array.span (\_ -> true) input
  assertEqual { expected: 10000, actual: Array.length result.init }
  assertEqual { expected: 0, actual: Array.length result.rest }

  log "Done"
