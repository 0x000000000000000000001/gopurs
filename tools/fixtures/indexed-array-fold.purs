-- @inline Main.runUnknown never
-- @inline Main.callbackFactory never
-- @inline Main.seedFactory never
-- @inline Main.sourceFactory never
-- @inline Main.stagedCallback never
-- @inline Main.observedStep never
-- @inline Main.partialStep never
-- @inline Main.recordIndex never
module Main where

-- @dependencies: prelude effect console refs arrays foldable-traversable maybe either exceptions

import Prelude
import Data.Array as Array
import Data.Either (Either(..))
import Data.Foldable (for_)
import Data.FoldableWithIndex (foldlWithIndex)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Exception (throw, try)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

weighted :: Int -> Array Int -> Int
weighted seed values = foldlWithIndex
  (\index total value -> total + (1 + index `mod` 2) * value) seed values

ordered :: Int -> Array Int -> Int
ordered seed values = foldlWithIndex
  (\index total value -> total * 31 + index * 7 + value) seed values

fixedPartial :: Array Int -> Int
fixedPartial = foldlWithIndex (\index total value -> total * 7 + value + index) 23

captured :: Boolean -> Int -> Int -> Array Int -> Int
captured enabled delta seed values = foldlWithIndex
  (\index total value -> if enabled then total + value + delta + index else total - value) seed values

indices :: Int -> Array Int -> Int
indices seed values = foldlWithIndex (\index total _ -> total + index) seed values

constant :: Int -> Array Int -> Int
constant seed values = foldlWithIndex (\_ _ _ -> 7) seed values

divisionZero :: Int -> Array Int -> Int
divisionZero seed values = foldlWithIndex
  (\index total value -> total + value `div` 0 + index `mod` 0) seed values

runUnknown :: (Int -> Int -> Int -> Int) -> Int -> Array Int -> Int
runUnknown callback seed values = foldlWithIndex callback seed values

structured :: Int -> Array Int -> Array Int
structured seed values =
  let result = foldlWithIndex
        (\index acc value -> { total: acc.total + value, visited: Array.snoc acc.visited index })
        { total: seed, visited: [] } values
  in Array.cons result.total result.visited

recordElements :: Int -> Array { value :: Int } -> Int
recordElements seed values = foldlWithIndex (\index total value -> total + index + value.value) seed values

strings :: String -> Array String -> String
strings seed values = foldlWithIndex (\index total value -> total <> show index <> value) seed values

type Trace = Ref.Ref (Array (Array Int))

stagedCallback :: Trace -> Int -> Int -> Int -> Int
stagedCallback trace index = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 4, index ]) trace
  pure (\total -> unsafePerformEffect do
    Ref.modify_ (\events -> Array.snoc events [ 5, total ]) trace
    pure (\value -> unsafePerformEffect do
      Ref.modify_ (\events -> Array.snoc events [ 6, value ]) trace
      pure (total * 10 + index + value)))

callbackFactory :: Trace -> (Int -> Int -> Int -> Int)
callbackFactory trace = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 1 ]) trace
  pure (stagedCallback trace)

seedFactory :: Trace -> Int -> Int
seedFactory trace seed = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 2, seed ]) trace
  pure seed

sourceFactory :: Trace -> Array Int -> Array Int
sourceFactory trace values = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 3, Array.length values ]) trace
  pure values

observedStep :: Trace -> Int -> Int -> Int -> Int
observedStep trace index total value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 7, index, total, value ]) trace
  pure (total + index + value)

partialStep :: Trace -> Int -> Int -> Int -> Int
partialStep trace index total value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 8, index, total, value ]) trace
  if index == 2 then throw "third callback" else pure (total + value)

recordIndex :: Trace -> Int -> Int
recordIndex trace index = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 9, index ]) trace
  pure index

main :: Effect Unit
main = do
  inputs <- Ref.new
    [ { values: [], seed: 19, delta: 7, enabled: true }
    , { values: [ 1 ], seed: -7, delta: 1, enabled: false }
    , { values: [ 1, 2, 3, 4 ], seed: 0, delta: -3, enabled: true }
    , { values: [ -3, 0, -3, 5 ], seed: -19, delta: -7, enabled: false }
    , { values: [ 2147483647, -2147483648, -1, 0 ], seed: 2147483647, delta: 1, enabled: true }
    , { values: [ -2147483648, 2147483647 ], seed: -2147483648, delta: -1, enabled: false }
    , { values: [ -7, 19, 8 ], seed: 17, delta: 2147483647, enabled: true }
    , { values: [ 0, 2, 4, 6, 8, 10, 12 ], seed: -19, delta: -2147483648, enabled: true }
    ]
  cases <- Ref.read inputs
  for_ cases \c -> do
    let
      partial = weighted c.seed
      callback = \index total value -> total * 31 + index * 7 + value
    logShow
      [ c.values
      , [ weighted c.seed c.values, ordered c.seed c.values
        , captured c.enabled c.delta c.seed c.values
        , indices c.seed c.values, constant c.seed c.values
        , runUnknown callback c.seed c.values, partial c.values, partial []
        , recordElements c.seed (map (\value -> { value }) c.values)
        , fixedPartial c.values, fixedPartial []
        , divisionZero c.seed c.values
        ]
      , structured c.seed c.values
      , c.values
      ]
  logShow (strings "seed:" [])
  logShow (strings "seed:" [ "a", "", "bc" ])
  trace <- Ref.new []
  arrays <- Ref.new [ [ 1, 2 ], [], [ 3, -1, 4 ] ]
  sources <- Ref.read arrays
  let partial = foldlWithIndex (callbackFactory trace) (seedFactory trace 5)
  for_ sources \values -> do
    dynamic <- Ref.new values
    current <- Ref.read dynamic
    logShow (partial (sourceFactory trace current))
  events <- Ref.read trace
  logShow events
  inlineTrace <- Ref.new []
  input <- Ref.read arrays
  for_ input \values -> do
    dynamic <- Ref.new values
    current <- Ref.read dynamic
    logShow (foldlWithIndex (\index total value -> observedStep inlineTrace index total value) 7 current)
  inlineEvents <- Ref.read inlineTrace
  logShow inlineEvents
  letTrace <- Ref.new []
  letInput <- Ref.new [ 1, 2, 3 ]
  letValues <- Ref.read letInput
  logShow (foldlWithIndex
    (\index -> let recorded = recordIndex letTrace index
      in \total value -> total + recorded + value) 11 letValues)
  letEvents <- Ref.read letTrace
  logShow letEvents
  failingTrace <- Ref.new []
  failingInput <- Ref.new [ 1, 2, 3, 4 ]
  outcome <- try do
    values <- Ref.read failingInput
    pure (foldlWithIndex (partialStep failingTrace) 9 values)
  logShow (case outcome of
    Left _ -> true
    Right _ -> false)
  failureEvents <- Ref.read failingTrace
  logShow failureEvents
