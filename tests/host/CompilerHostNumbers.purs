module Main where

-- @dependencies: assert prelude effect console refs

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

-- Folded literals must keep their IEEE sign through the compiler itself.
negativeZero :: Number
negativeZero = negate 0.0

positiveZero :: Number
positiveZero = negate negativeZero

checkZero :: Boolean -> Number -> Effect Unit
checkZero negative value = do
  ref <- Ref.new value
  actual <- Ref.read ref
  assertEqual { expected: 0.0, actual }
  assertEqual { expected: negative, actual: 1.0 / actual < 0.0 }

main :: Effect Unit
main = do
  checkZero true negativeZero
  checkZero false positiveZero
  ref <- Ref.new 0.0
  dynamic <- Ref.read ref
  checkZero true (negate dynamic)
  checkZero false (negate (negate dynamic))
  log "Done"
