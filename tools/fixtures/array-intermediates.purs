-- @inline Main.observeMap never
-- @inline Main.observeFold never
-- @inline Main.mapFactory never
-- @inline Main.foldFactory never
-- @inline Main.seedFactory never
-- @inline Main.inputFactory never
-- @inline Main.failingMap never
-- @inline Main.failingFold never
-- @inline Main.failingTraverse never
-- @inline Main.failingRight never
-- @inline Main.applyConsumer never
module Main where

-- @dependencies: prelude effect console refs arrays foldable-traversable maybe either exceptions

import Prelude
import Consumer as Consumer
import Data.Array as Array
import Data.Either (Either(..))
import Data.Foldable (for_)
import Data.FoldableWithIndex (foldlWithIndex)
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Exception (message, throw, try)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

mapScore :: Array Int -> Int
mapScore values = Consumer.consume (map (\value -> value + 1) values)

appendScore :: Array Int -> Array Int -> Int
appendScore left right = Consumer.consume (left <> right)

traverseSuccess :: Array Int -> Int
traverseSuccess values = case traverse (\value -> Just (value + 1)) values of
  Nothing -> -1
  Just mapped -> Consumer.consume mapped

traverseRejected :: Int -> Array Int -> Int
traverseRejected rejected values = case traverse (\value -> if value == rejected then Nothing else Just (value + 1)) values of
  Nothing -> -1
  Just mapped -> Consumer.consume mapped

mapCaptured :: Int -> Int -> Array Int -> Int
mapCaptured delta seed values = foldlWithIndex
  (\index total value -> total + (1 + index `mod` 2) * value) seed (map (\value -> value + delta) values)

appendSeeded :: Int -> Array Int -> Array Int -> Int
appendSeeded seed left right = foldlWithIndex
  (\index total value -> total + (1 + index `mod` 2) * value) seed (left <> right)

applyConsumer :: (Array Int -> Int) -> Array Int -> Int
applyConsumer callback values = callback values

encode :: Maybe (Array Int) -> Array Int
encode Nothing = [ 0 ]
encode (Just values) = Array.cons 1 values

type Trace = Ref.Ref (Array (Array Int))

observeMap :: Trace -> Int -> Int
observeMap trace value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 1, value ]) trace
  pure (value + 1)

observeFold :: Trace -> Int -> Int -> Int -> Int
observeFold trace index total value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 2, index, total, value ]) trace
  pure (total + (1 + index `mod` 2) * value)

mapFactory :: Trace -> (Int -> Int)
mapFactory trace = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 10 ]) trace
  pure (observeMap trace)

foldFactory :: Trace -> (Int -> Int -> Int -> Int)
foldFactory trace = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 11 ]) trace
  pure (observeFold trace)

seedFactory :: Trace -> Int -> Int
seedFactory trace seed = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 12, seed ]) trace
  pure seed

inputFactory :: Trace -> Int -> Array Int -> Array Int
inputFactory trace tag values = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ tag, Array.length values ]) trace
  pure values

failingMap :: Trace -> Int -> Int
failingMap trace value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 3, value ]) trace
  if value == 3 then throw "producer failure" else pure (value + 1)

failingFold :: Trace -> Int -> Int -> Int -> Int
failingFold trace index _ _ = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 4, index ]) trace
  throw "consumer failure"

failingTraverse :: Trace -> Int -> Maybe Int
failingTraverse trace value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 5, value ]) trace
  if value == 3 then throw "later traversal failure"
    else pure (if value == 1 then Nothing else Just (value + 1))

failingRight :: Trace -> Array Int -> Array Int
failingRight trace _ = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 15 ]) trace
  throw "right input failure"

main :: Effect Unit
main = do
  inputs <- Ref.new
    [ { left: [], right: [], rejected: 99, delta: 7, seed: 19 }
    , { left: [ 1 ], right: [], rejected: 1, delta: 1, seed: -7 }
    , { left: [ 1, 2, 3, 4 ], right: [ 5, 6 ], rejected: 1, delta: -3, seed: 0 }
    , { left: [ 1, 2, 3, 4 ], right: [ 5 ], rejected: 3, delta: 7, seed: 5 }
    , { left: [ 1, 2, 3, 4 ], right: [ -2, -3 ], rejected: 4, delta: -1, seed: 2147483647 }
    , { left: [ -3, 0, -3, 5 ], right: [ 0 ], rejected: 99, delta: -2147483648, seed: -19 }
    , { left: [ 2147483647, -2147483648, -1, 0 ], right: [ 1, 2147483647 ], rejected: 19, delta: 1, seed: 2147483647 }
    , { left: [ -2147483648, 2147483647 ], right: [], rejected: -2147483648, delta: -1, seed: -2147483648 }
    , { left: [], right: [ 3, -1, 4 ], rejected: 3, delta: 1, seed: 7 }
    , { left: [ 4, -8, 9 ], right: [], rejected: 9, delta: 2147483647, seed: -7 }
    ]
  cases <- Ref.read inputs
  for_ cases \c -> do
    let
      mapped = map (_ + 1) c.left
      appended = c.left <> c.right
      traversed = traverse (\value -> if value == c.rejected then Nothing else Just (value + 1)) c.left
    stored <- Ref.new traversed
    replayed <- Ref.read stored
    logShow [ mapScore c.left, appendScore c.left c.right, traverseSuccess c.left
      , traverseRejected c.rejected c.left, mapCaptured c.delta c.seed c.left
      , appendSeeded c.seed c.left c.right, mapScore c.left ]
    logShow [ c.left, c.right, mapped, appended, encode traversed, encode replayed, mapped, c.left, c.right ]
  consumerRef <- Ref.new Consumer.consume
  callback <- Ref.read consumerRef
  logShow [ callback [ 1, 2, 3 ], callback [], callback [ 4, -1 ] ]
  logShow (applyConsumer callback [ 2, 0, -3 ])
  input <- Ref.new [ 1, 2, 3, 4 ]
  values <- Ref.read input
  trace <- Ref.new []
  -- Sequence these observable factories through Effect. The separate archived
  -- unsequenced diagnostic records a pre-existing generic argument-order issue.
  foldRef <- Ref.new (foldFactory trace)
  fold <- Ref.read foldRef
  seedRef <- Ref.new (seedFactory trace 7)
  seed <- Ref.read seedRef
  logShow (foldlWithIndex fold seed
    (map (mapFactory trace) (inputFactory trace 13 values)))
  Ref.read trace >>= logShow
  partialTrace <- Ref.new []
  partialRef <- Ref.new (foldlWithIndex (foldFactory partialTrace) (seedFactory partialTrace 7))
  partial <- Ref.read partialRef
  logShow [ partial [ 1, 2 ], partial [], partial [ 3 ] ]
  Ref.read partialTrace >>= logShow
  mapTrace <- Ref.new []
  mapOutcome <- try do
    current <- Ref.read input
    pure (foldlWithIndex (failingFold mapTrace) 0 (map (failingMap mapTrace) current))
  logShow (case mapOutcome of
    Left error -> message error
    Right value -> show value)
  Ref.read mapTrace >>= logShow
  traverseTrace <- Ref.new []
  traverseOutcome <- try do
    current <- Ref.read input
    pure (case traverse (failingTraverse traverseTrace) current of
      Nothing -> -1
      Just result -> Consumer.consume result)
  logShow (case traverseOutcome of
    Left error -> message error
    Right value -> show value)
  Ref.read traverseTrace >>= logShow
  appendTrace <- Ref.new []
  appendOutcome <- try do
    current <- Ref.read input
    pure (foldlWithIndex (failingFold appendTrace) 0
      (inputFactory appendTrace 14 current <> failingRight appendTrace current))
  logShow (case appendOutcome of
    Left error -> message error
    Right value -> show value)
  Ref.read appendTrace >>= logShow
  Consumer.seedRunsCount >>= logShow
