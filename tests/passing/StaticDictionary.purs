-- @inline export callRun never
-- @inline export callUse never
module Main where

-- @dependencies: assert prelude effect console refs either
import Prelude
import Data.Either (Either(..))
import Effect (Effect)
import Effect.Console (log)
import Test.Assert (assertEqual)
import Worker (runSummary, useSummary)

callRun :: Int -> Int
callRun n = case runSummary n of
  Left _ -> -1
  Right s -> s.id

callUse :: Int -> Int
callUse n = case useSummary n of
  Left _ -> -1
  Right s -> s.id

main :: Effect Unit
main = do
  assertEqual { expected: 8, actual: callRun 7 }
  assertEqual { expected: 9, actual: callUse 7 }
  let runName = case runSummary 7 of
        Right s -> s.name
        Left _ -> ""
  assertEqual { expected: "alpha", actual: runName }
  assertEqual { expected: -1, actual: callRun (-1) }
  assertEqual { expected: -1, actual: callUse (-1) }
  log "Done"
