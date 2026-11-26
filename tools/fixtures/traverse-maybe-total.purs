-- @inline Main.runUnknown never
-- @inline Main.recordCallback never
-- @inline Main.callbackFactory never
-- @inline Main.sourceFactory never
-- @inline Main.partialCallback never
module Main where

-- @dependencies: prelude effect console refs arrays foldable-traversable maybe either exceptions

import Prelude
import Data.Array as Array
import Data.Either (Either(..))
import Data.Foldable (for_)
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Data.TraversableWithIndex (traverseWithIndex)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Exception (throw, try)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

-- Return the complete array to an arbitrary consumer: the optimization may
-- not assume that the result is consumed by a checksum or known benchmark.
selected :: Int -> Int -> Array Int -> Maybe (Array Int)
selected rejected delta values = traverse
  (\value -> if value == rejected then Nothing else Just (value + delta)) values

allSuccess :: Int -> Array Int -> Maybe (Array Int)
allSuccess delta values = traverse (\value -> Just (value + delta)) values

indexed :: Int -> Int -> Array Int -> Maybe (Array Int)
indexed rejected delta values = traverseWithIndex
  (\index value -> if index == rejected then Nothing else Just (value + delta + index)) values

constantSuccess :: Array Int -> Maybe (Array Int)
constantSuccess values = traverse (\_ -> Just 7) values

constantFailure :: Array Int -> Maybe (Array Int)
constantFailure values = traverse (\_ -> Nothing) values

capturedFlag :: Boolean -> Array Int -> Maybe (Array Int)
capturedFlag reject values = traverse
  (\value -> if reject then Nothing else Just (value + 1)) values

branched :: Array Int -> Maybe (Array Int)
branched values = traverse
  (\value -> if value < 0 then Just (-value)
    else if value == 2 then Nothing
    else Just (value * 3 - 7)) values

runUnknown :: (Int -> Maybe Int) -> Array Int -> Maybe (Array Int)
runUnknown callback values = traverse callback values

encode :: Maybe (Array Int) -> Array Int
encode Nothing = [ 0 ]
encode (Just values) = Array.cons 1 values

type Trace = Ref.Ref (Array (Array Int))

recordCallback :: Trace -> Int -> Int -> Maybe Int
recordCallback trace rejected value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 3, value ]) trace
  pure (if value == rejected then Nothing else Just (value + 1))

callbackFactory :: Trace -> Int -> (Int -> Maybe Int)
callbackFactory trace rejected = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 1, rejected ]) trace
  pure (recordCallback trace rejected)

sourceFactory :: Trace -> Array Int -> Array Int
sourceFactory trace values = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 2, Array.length values ]) trace
  pure values

partialCallback :: Trace -> Int -> Maybe Int
partialCallback trace value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 4, value ]) trace
  if value == 3 then throw "later callback" else pure Nothing

main :: Effect Unit
main = do
  inputs <- Ref.new
    [ { values: [], rejected: 99, delta: 1 }
    , { values: [ 1 ], rejected: 1, delta: 1 }
    , { values: [ 1 ], rejected: 99, delta: 1 }
    , { values: [ 1, 2, 3, 4 ], rejected: 1, delta: 7 }
    , { values: [ 1, 2, 3, 4 ], rejected: 2, delta: -3 }
    , { values: [ 1, 2, 3, 4 ], rejected: 4, delta: 1 }
    , { values: [ 1, 2, 3, 4 ], rejected: 99, delta: 1 }
    , { values: [ -3, 0, -3, 5 ], rejected: -3, delta: -7 }
    , { values: [ 2147483647, -2147483648, -1, 0 ], rejected: 19, delta: 1 }
    , { values: [ -2147483648, 2147483647 ], rejected: 19, delta: -1 }
    , { values: [ -7, 19, 8 ], rejected: 8, delta: 2147483647 }
    , { values: [ 0, 2, 4, 6, 8, 10, 12 ], rejected: 6, delta: -2147483648 }
    ]
  cases <- Ref.read inputs
  for_ cases \c -> do
    let
      result = selected c.rejected c.delta c.values
      partial = selected c.rejected c.delta
      callback = \value -> if value == c.rejected then Nothing else Just (value + c.delta)
    logShow
      [ c.values
      , encode result
      , encode (partial c.values)
      , encode (runUnknown callback c.values)
      , encode (allSuccess c.delta c.values)
      , encode (indexed c.rejected c.delta c.values)
      , encode (constantSuccess c.values)
      , encode (constantFailure c.values)
      , encode (capturedFlag (c.rejected == 1) c.values)
      , encode (branched c.values)
      , encode result
      , c.values
      ]
  for_ [ 1, 2, 4, 99 ] \rejected -> do
    trace <- Ref.new []
    values <- Ref.new [ 1, 2, 3, 4 ]
    input <- Ref.read values
    -- Callback creation, then input creation, then every callback left to
    -- right. A preceding Nothing cannot suppress later observable work.
    logShow (encode (traverse (callbackFactory trace rejected) (sourceFactory trace input)))
    events <- Ref.read trace
    logShow events
  partialTrace <- Ref.new []
  partialInput <- Ref.new [ 1, 2, 3, 4 ]
  outcome <- try do
    input <- Ref.read partialInput
    pure (traverse (partialCallback partialTrace) input)
  logShow (case outcome of
    Left _ -> true
    Right _ -> false)
  events <- Ref.read partialTrace
  logShow events
