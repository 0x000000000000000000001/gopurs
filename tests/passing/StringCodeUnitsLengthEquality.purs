-- @inline export sameLength never
-- @inline export differentLength never
-- @inline export shorter never
-- @inline export compareWith never
-- @inline export check never
module Main where

-- @dependencies: assert arrays strings prelude effect console refs

import Prelude

import Data.Array as Array
import Data.String.CodeUnits as SCU
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

sameLength :: String -> String -> Boolean
sameLength left right = SCU.length left == SCU.length right

differentLength :: String -> String -> Boolean
differentLength left right = SCU.length left /= SCU.length right

shorter :: String -> String -> Boolean
shorter left right = SCU.length left < SCU.length right

compareWith :: (String -> Int) -> String -> String -> Boolean
compareWith measure left right = measure left == measure right

check :: Boolean -> String -> String -> Effect Unit
check expected left right = do
  reference <- Ref.new { left, right }
  input <- Ref.read reference
  assertEqual { expected, actual: sameLength input.left input.right }
  assertEqual { expected: not expected, actual: differentLength input.left input.right }

main :: Effect Unit
main = do
  check true "" ""
  check false "" "a"
  check true "abc" "xyz"
  check false "abc" "ab"
  -- UTF-16 lengths can agree with different byte lengths, and conversely.
  check true "é" "a"
  check false "é" "ab"
  check true "😀" "ab"
  check false "😀" "abcd"
  check true "é😀" "abc"
  check true "\xD800" "a"
  check true "\xDFFF" "é"
  check true "\xD834\xDF06" "ab"
  check false "\xD800" "ab"

  reference <- Ref.new ("x" <> SCU.fromCharArray (Array.replicate 4096 'é') <> "😀")
  source <- Ref.read reference
  let
    suffix = SCU.drop 1 source
    copied = SCU.fromCharArray (SCU.toCharArray suffix)
  check true suffix suffix
  check true suffix copied
  check false source suffix
  assertEqual { expected: true, actual: shorter suffix source }
  assertEqual { expected: false, actual: shorter source suffix }

  partialRef <- Ref.new (sameLength suffix)
  partial <- Ref.read partialRef
  assertEqual { expected: true, actual: partial copied }
  assertEqual { expected: false, actual: partial source }

  -- The caller cannot assume that an opaque callback measures string length.
  measureRef <- Ref.new SCU.length
  measure <- Ref.read measureRef
  assertEqual { expected: true, actual: compareWith measure suffix copied }
  assertEqual { expected: false, actual: compareWith measure source suffix }
  assertEqual { expected: true, actual: compareWith (const 0) source suffix }
  log "Done"
