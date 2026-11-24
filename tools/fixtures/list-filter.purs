-- @inline Main.recordVisit never
-- @inline Main.predicateFactory never
-- @inline Main.sourceFactory never
module Main where

-- @dependencies: prelude effect console refs arrays foldable-traversable

import Prelude
import Data.Array as Array
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import FilterLibrary (Trail(..))
import FilterLibrary as FilterLibrary

makeValues :: Int -> Int -> Trail Int
makeValues lower upper = collect upper Stop
  where
  collect value result =
    if value < lower then result
    else collect (value - 1) (Cell value result)

turnAround :: forall a. Trail a -> Trail a
turnAround input = collect input Stop
  where
  collect Stop result = result
  collect (Cell value rest) result = collect rest (Cell value result)

stableSelect :: forall a. (a -> Boolean) -> Trail a -> Trail a
stableSelect predicate input = collect input Stop
  where
  collect Stop result = turnAround result
  collect (Cell value rest) result =
    if predicate value then collect rest (Cell value result)
    else collect rest result

keepDivisible :: Int -> Trail Int -> Trail Int
keepDivisible divisor input = stableSelect (\value -> value `mod` divisor == 0) input

keepOthers :: Int -> Trail Int -> Trail Int
keepOthers divisor input = stableSelect (\value -> value `mod` divisor /= 0) input

countInts :: Trail Int -> Int
countInts input = collect input 0
  where
  collect Stop count = count
  collect (Cell _ rest) count = collect rest (count + 1)

fingerprintInts :: Trail Int -> Int
fingerprintInts input = collect input 0
  where
  collect Stop result = result
  collect (Cell value rest) result = collect rest (result * 31 + value)

toArray :: forall a. Trail a -> Array a
toArray Stop = []
toArray (Cell value rest) = Array.cons value (toArray rest)

transform :: forall a b. (a -> b) -> Trail a -> Trail b
transform _ Stop = Stop
transform change (Cell value rest) = Cell (change value) (transform change rest)

type Entry = { key :: Int, checksum :: Int }

keepEntries :: Int -> Trail Entry -> Trail Entry
keepEntries divisor input = stableSelect (\entry -> entry.key `mod` divisor == 0) input

recordVisit :: Ref.Ref (Array (Array Int)) -> Trail Int -> Int -> Int -> Boolean
recordVisit reference original divisor value = unsafePerformEffect do
  Ref.modify_ (\visited -> Array.snoc visited [ value, fingerprintInts original ]) reference
  pure (value `mod` divisor == 0)

predicateFactory :: Ref.Ref (Array (Array Int)) -> Trail Int -> Int -> (Int -> Boolean)
predicateFactory reference original divisor = unsafePerformEffect do
  Ref.modify_ (\visited -> Array.snoc visited [ -1001, divisor ]) reference
  pure (recordVisit reference original divisor)

sourceFactory :: Ref.Ref (Array (Array Int)) -> Trail Int -> Trail Int
sourceFactory reference original = unsafePerformEffect do
  Ref.modify_ (\visited -> Array.snoc visited [ -1002, fingerprintInts original ]) reference
  pure original

observedSelect :: Ref.Ref (Array (Array Int)) -> Int -> Trail Int -> Trail Int
observedSelect reference divisor source =
  stableSelect (predicateFactory reference source divisor) (sourceFactory reference source)

main :: Effect Unit
main = for_
  [ { lower: 4, upper: 3, divisor: 2 }
  , { lower: 0, upper: 0, divisor: 0 }
  , { lower: -4, upper: 4, divisor: 2 }
  , { lower: -4, upper: 4, divisor: -3 }
  , { lower: -4, upper: 4, divisor: 0 }
  , { lower: -4, upper: 4, divisor: 1 }
  , { lower: -4, upper: 4, divisor: -1 }
  , { lower: 1, upper: 9, divisor: 2 }
  , { lower: 1, upper: 9, divisor: 20 }
  , { lower: 1, upper: 12, divisor: 3 }
  , { lower: 8, upper: 8, divisor: 3 }
  ] \input -> do
    dynamic <- Ref.new input
    arguments <- Ref.read dynamic
    let
      source = makeValues arguments.lower arguments.upper
      selected = keepDivisible arguments.divisor source
      rejected = keepOthers arguments.divisor source
      selectedAgain = stableSelect (_ > 0) selected
      entries = transform (\key -> { key, checksum: key * key + 3 }) source
      selectedEntries = keepEntries arguments.divisor entries
    logShow
      [ toArray selected
      , toArray rejected
      , toArray selectedAgain
      -- Consume both original lists again after filtering them.
      , toArray source
      , toArray selected
      , toArray (transform _.key selectedEntries)
      , toArray (transform _.checksum selectedEntries)
      , toArray (transform _.checksum entries)
      ]
    visited <- Ref.new []
    -- Logging the selected result forces all callbacks before the Ref read.
    -- The explicit lambda can use the native ABI with all three captures.
    logShow (toArray (stableSelect (\value -> recordVisit visited source arguments.divisor value) source))
    calls <- Ref.read visited
    logShow calls
    importedVisits <- Ref.new []
    logShow (toArray (FilterLibrary.selectInts
      (\value -> recordVisit importedVisits source arguments.divisor value) source))
    importedCalls <- Ref.read importedVisits
    logShow importedCalls
    factories <- Ref.new []
    logShow (toArray (observedSelect factories arguments.divisor source))
    factoryCalls <- Ref.read factories
    logShow factoryCalls
    opaqueVisits <- Ref.new []
    predicateReference <- Ref.new (recordVisit opaqueVisits source arguments.divisor)
    opaquePredicate <- Ref.read predicateReference
    logShow (toArray (stableSelect opaquePredicate source))
    opaqueCalls <- Ref.read opaqueVisits
    logShow opaqueCalls
    partialVisits <- Ref.new []
    -- A first-class imported filter and its saved partial application both
    -- retain the public boxed ABI. Creating the partial must not traverse.
    filterReference <- Ref.new FilterLibrary.selectInts
    opaqueFilter <- Ref.read filterReference
    partialReference <- Ref.new (opaqueFilter (predicateFactory partialVisits source arguments.divisor))
    partial <- Ref.read partialReference
    beforePartial <- Ref.read partialVisits
    logShow beforePartial
    logShow (toArray (partial source))
    logShow (toArray (partial source))
    partialCalls <- Ref.read partialVisits
    logShow partialCalls
