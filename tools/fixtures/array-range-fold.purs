-- @inline Main.recordInt never
-- @inline Main.predicateFactory never
-- @inline Main.recordPredicate never
-- @inline Main.recordFold never
module Main where

-- @dependencies: prelude effect console refs arrays foldable-traversable

import Prelude
import Data.Array as Array
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

makeRange :: Int -> Int -> Array Int
makeRange first last = Array.range first last

sumRange :: Int -> Int -> Int -> Int -> Int
sumRange first last seed divisor =
  Array.foldl (+) seed (Array.filter (\value -> value `mod` divisor == 0) (Array.range first last))

sumEven :: Int -> Int -> Int -> Int
sumEven first last seed =
  Array.foldl (+) seed (Array.filter (\value -> value `mod` 2 == 0) (Array.range first last))

sumResidue :: Int -> Int -> Int -> Int
sumResidue first last seed =
  Array.foldl (+) seed (Array.filter (\value -> value `mod` 5 == 3) (Array.range first last))

-- The array parameter keeps this path independent of range fusion. The Go
-- probe also uses it as a reference for noncanonical Int64 boundary values.
sumBorrowed :: Int -> Int -> Array Int -> Int
sumBorrowed seed divisor input =
  Array.foldl (+) seed (Array.filter (\value -> value `mod` divisor == 0) input)

sumBorrowedResidue :: Int -> Array Int -> Int
sumBorrowedResidue seed input =
  Array.foldl (+) seed (Array.filter (\value -> value `mod` 5 == 3) input)

orderedScore :: Int -> Int -> Int -> Int -> Int
orderedScore first last seed divisor =
  Array.foldl (\acc value -> acc * 31 + value) seed
    (Array.filter (\value -> value `mod` divisor == 0) (Array.range first last))

type Trace = Ref.Ref (Array (Array Int))

recordInt :: Trace -> Int -> Int -> Int
recordInt trace tag value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ tag, value ]) trace
  pure value

recordPredicate :: Trace -> Int -> Int -> Boolean
recordPredicate trace divisor value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 4, value ]) trace
  pure (value `mod` divisor == 0)

predicateFactory :: Trace -> Int -> (Int -> Boolean)
predicateFactory trace divisor = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 1, divisor ]) trace
  pure (recordPredicate trace divisor)

recordFold :: Trace -> Int -> Int -> Int
recordFold trace accumulated value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ 5, value ]) trace
  pure (accumulated * 31 + value)

observedSum :: Trace -> Int -> Int -> Int -> Int -> Int
observedSum trace first last seed divisor =
  Array.foldl (+) (recordInt trace 0 seed)
    (Array.filter (predicateFactory trace divisor)
      (Array.range (recordInt trace 2 first) (recordInt trace 3 last)))

-- A pure predicate permits fewer invocations, but seed and bound evaluation
-- must still occur once, in the original order, before the arithmetic sum.
observedPureSum :: Trace -> Int -> Int -> Int -> Int
observedPureSum trace first last seed =
  Array.foldl (+) (recordInt trace 0 seed)
    (Array.filter (\value -> value `mod` 5 == 3)
      (Array.range (recordInt trace 2 first) (recordInt trace 3 last)))

-- An effectful fold must preserve two phases: every predicate first, then
-- every fold callback. Interleaving them would change the recorded trace.
twoPhaseFold :: Trace -> Int -> Int -> Int -> Int -> Int
twoPhaseFold trace first last seed divisor =
  Array.foldl (recordFold trace) (recordInt trace 0 seed)
    (Array.filter (predicateFactory trace divisor)
      (Array.range (recordInt trace 2 first) (recordInt trace 3 last)))

main :: Effect Unit
main = for_
  [ { first: 1, last: 9, seed: 0, divisor: 2 }
  , { first: 9, last: 1, seed: 17, divisor: 2 }
  , { first: -4, last: 4, seed: 0, divisor: 3 }
  , { first: 4, last: -4, seed: -17, divisor: -3 }
  , { first: 3, last: 3, seed: 2147483647, divisor: 2 }
  , { first: 4, last: 4, seed: 2147483647, divisor: 2 }
  , { first: 0, last: 0, seed: -7, divisor: 0 }
  , { first: -2147483648, last: -2147483646, seed: -1, divisor: 2 }
  , { first: -2147483646, last: -2147483648, seed: -2147483648, divisor: -3 }
  , { first: 2147483645, last: 2147483647, seed: 2147483647, divisor: 2 }
  , { first: 2147483647, last: 2147483645, seed: -2147483648, divisor: 3 }
  , { first: 1, last: 7, seed: 13, divisor: 20 }
  , { first: 1, last: 8, seed: -2147483648, divisor: 0 }
  , { first: -8, last: -1, seed: 2147483647, divisor: 2 }
  , { first: 2, last: 2, seed: 0, divisor: -2147483648 }
  ] \input -> do
    dynamic <- Ref.new input
    arguments <- Ref.read dynamic
    let
      source = Array.range arguments.first arguments.last
      selected = Array.filter (\value -> value `mod` arguments.divisor == 0) source
      positive = Array.filter (_ > 0) selected
    logShow
      [ source, selected, positive, source, selected
      , [ sumRange arguments.first arguments.last arguments.seed arguments.divisor
        , sumBorrowed arguments.seed arguments.divisor source
        , orderedScore arguments.first arguments.last arguments.seed arguments.divisor
        ]
      ]
    observed <- Ref.new []
    logShow (observedSum observed arguments.first arguments.last arguments.seed arguments.divisor)
    firstTrace <- Ref.read observed
    logShow firstTrace
    phased <- Ref.new []
    logShow (twoPhaseFold phased arguments.first arguments.last arguments.seed arguments.divisor)
    secondTrace <- Ref.read phased
    logShow secondTrace
    logShow
      [ sumEven arguments.first arguments.last arguments.seed
      , sumResidue arguments.first arguments.last arguments.seed
      ]
    pureObserved <- Ref.new []
    logShow (observedPureSum pureObserved arguments.first arguments.last arguments.seed)
    pureTrace <- Ref.read pureObserved
    logShow pureTrace
