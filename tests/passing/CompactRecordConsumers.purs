-- @inline export makeEntry never
-- @inline export produceEntry never
-- @inline export inspectProduced never
module Main where

-- @dependencies: assert arrays prelude effect console refs
-- @snapshot-ffi

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

type Entry = { count :: Int, label :: String }
type Payload = { child :: Entry, values :: Array Int }
type Scalars = { flag :: Boolean, number :: Number }

-- The Go FFI supplies actual RecordDict2 values, including reversed keys.
-- An effect keeps these inputs opaque to the PureScript optimizer.
foreign import compactEntry :: Boolean -> Effect Entry
foreign import compactPayload :: Effect Payload
foreign import compactScalars :: Effect Scalars

-- Test-only representation probe: the Go version rejects a noncompact input.
-- Use only on the original FFI values, not on native records reboxed by Go.
foreign import compactKeys :: forall a. a -> String

-- These Go functions take map[string]any, exercising the generated bridge.
-- Its field values must remain boxed, including nested records and arrays.
foreign import describeEntryMap :: Entry -> String
foreign import describePayloadMap :: Payload -> String
foreign import describeScalarsMap :: Scalars -> String

-- Keep the produced value polymorphic until the FFI consumer receives it.
-- Passing a closed record directly currently unboxes and reboxes it first.
inspectProduced :: forall a. (a -> String) -> Effect a -> Effect String
inspectProduced inspect produce = do
  value <- produce
  pure (inspect value)

check :: String -> String -> String -> Effect Unit
check label expected actual = do
  assertEqual { expected, actual }
  log (label <> ": " <> actual)

checkEntry :: Boolean -> Effect Unit
checkEntry reverse = do
  entry <- compactEntry reverse
  retained <- Ref.new entry
  let
    prefix = if reverse then "reversed " else "ordered "
    keys = if reverse then "label|count" else "count|label"
  actualKeys <- inspectProduced compactKeys (compactEntry reverse)
  check (prefix <> "compact keys") keys actualKeys
  check (prefix <> "count access") "5" (show entry.count)
  check (prefix <> "label access") "alpha" entry.label
  check (prefix <> "FFI map") "5:alpha" (describeEntryMap entry)
  directMap <- inspectProduced describeEntryMap (compactEntry reverse)
  check (prefix <> "compact FFI map") "5:alpha" directMap

  -- Each update starts from the compact input, rather than another update
  -- whose result may already have become a native Go struct.
  let
    changedCount = entry { count = -7 }
    changedLabel = entry { label = "beta" }
    changedBoth = entry { label = "", count = 0 }
  savedCount <- Ref.new changedCount
  check (prefix <> "count update") "-7:alpha" (describeEntryMap changedCount)
  check (prefix <> "label update") "5:beta" (describeEntryMap changedLabel)
  check (prefix <> "both updates") "0:" (describeEntryMap changedBoth)
  check (prefix <> "original after updates") "5:alpha" (describeEntryMap entry)

  originalAgain <- Ref.read retained
  countAgain <- Ref.read savedCount
  check (prefix <> "retained original") "5:alpha" (describeEntryMap originalAgain)
  check (prefix <> "retained count update") "-7:alpha" (describeEntryMap countAgain)

checkPayload :: Effect Unit
checkPayload = do
  payload <- compactPayload
  retained <- Ref.new payload
  actualKeys <- inspectProduced compactKeys compactPayload
  check "payload compact keys" "values|child" actualKeys
  check "nested count access" "2" (show payload.child.count)
  check "nested label access" "nested" payload.child.label
  check "array access" "[-3,0,7]" (show payload.values)
  check "payload FFI map" "2:nested:[-3 0 7]" (describePayloadMap payload)
  directMap <- inspectProduced describePayloadMap compactPayload
  check "compact payload FFI map" "2:nested:[-3 0 7]" directMap

  let
    emptyValues = payload { values = [] }
    changedChild = payload { child = { count: 9, label: "replacement" } }
    changedNested = payload { child { label = "updated" } }
  check "empty array update" "2:nested:[]" (describePayloadMap emptyValues)
  check "child update" "9:replacement:[-3 0 7]" (describePayloadMap changedChild)
  check "nested field update" "2:updated:[-3 0 7]" (describePayloadMap changedNested)
  check "payload after updates" "2:nested:[-3 0 7]" (describePayloadMap payload)
  originalAgain <- Ref.read retained
  check "retained payload" "2:nested:[-3 0 7]" (describePayloadMap originalAgain)

checkScalars :: Effect Unit
checkScalars = do
  scalars <- compactScalars
  actualKeys <- inspectProduced compactKeys compactScalars
  check "scalar compact keys" "number|flag" actualKeys
  check "boolean access" "true" (show scalars.flag)
  check "number access" "-1.25" (show scalars.number)
  check "scalar FFI map" "true:-1.25" (describeScalarsMap scalars)
  directMap <- inspectProduced describeScalarsMap compactScalars
  check "compact scalar FFI map" "true:-1.25" directMap
  let changed = scalars { flag = false, number = 2.5 }
  check "scalar updates" "false:2.5" (describeScalarsMap changed)
  check "retained scalars" "true:-1.25" (describeScalarsMap scalars)

makeEntry :: Int -> String -> Entry
makeEntry count label = { count, label }

produceEntry :: Ref.Ref (Array String) -> Effect Entry
produceEntry trace = do
  Ref.modify_ (_ <> [ "produce" ]) trace
  -- Keep this native expression visible: boxing must evaluate it only once.
  pure (makeEntry 9 "created")

checkEffects :: Effect Unit
checkEffects = do
  trace <- Ref.new []
  consumerRef <- Ref.new (\entry ->
    Ref.modify_ (_ <> [ "consume:" <> describeEntryMap entry ]) trace)
  consume <- Ref.read consumerRef
  Ref.modify_ (_ <> [ "before" ]) trace
  entry <- produceEntry trace
  consume entry
  Ref.modify_ (_ <> [ "after" ]) trace
  actual <- Ref.read trace
  assertEqual
    { expected: [ "before", "produce", "consume:9:created", "after" ]
    , actual
    }
  log ("effect order: " <> show actual)

main :: Effect Unit
main = do
  checkEntry false
  checkEntry true
  checkPayload
  checkScalars
  checkEffects
  log "Done"
