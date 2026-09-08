-- Keep the native producer and the boxing at the opaque call visible in Go.
-- @inline export makeEntry never
-- @inline export consumeEntry never
module Main where

-- @dependencies: assert prelude effect console refs

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

type Entry = { count :: Int, label :: String }

makeEntry :: Int -> String -> Entry
makeEntry count label = { count, label }

-- The closed record parameter must stay native, while the callback is opaque.
-- Its call currently boxes via RecordDict([]string{...}, []Value{...}), using
-- CodeGen.boxGoExprImpl (TypeRecord fields), rather than Printer's RecordDict2.
consumeEntry :: (Entry -> String) -> Entry -> String
consumeEntry consume entry = consume entry

describeEntry :: Entry -> String
describeEntry entry = show entry.count <> ":" <> entry.label

check :: String -> String -> String -> Effect Unit
check label expected actual = do
  assertEqual { expected, actual }
  log (label <> ": " <> actual)

main :: Effect Unit
main = do
  consumerRef <- Ref.new describeEntry
  consume <- Ref.read consumerRef

  originalRef <- Ref.new (makeEntry 5 "alpha")
  original <- Ref.read originalRef
  check "original" "5:alpha" (consumeEntry consume original)

  let changedCount = original { count = -7 }
  countRef <- Ref.new changedCount
  check "count updated" "-7:alpha" (consumeEntry consume changedCount)
  check "original after count update" "5:alpha" (consumeEntry consume original)

  let changedLabel = changedCount { label = "beta" }
  labelRef <- Ref.new changedLabel
  check "label updated" "-7:beta" (consumeEntry consume changedLabel)
  check "count version after label update" "-7:alpha" (consumeEntry consume changedCount)

  let changedBoth = changedLabel { count = 0, label = "" }
  check "both fields updated" "0:" (consumeEntry consume changedBoth)
  check "label version after both updates" "-7:beta" (consumeEntry consume changedLabel)

  -- Read the retained versions after all updates and opaque consumptions.
  savedOriginal <- Ref.read originalRef
  savedCount <- Ref.read countRef
  savedLabel <- Ref.read labelRef
  check "retained original" "5:alpha" (consumeEntry consume savedOriginal)
  check "retained count version" "-7:alpha" (consumeEntry consume savedCount)
  check "retained label version" "-7:beta" (consumeEntry consume savedLabel)

  log "Done"
