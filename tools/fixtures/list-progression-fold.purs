-- @inline Main.recordValue never
-- @inline Main.recordDecision never
-- @inline Main.sumBorrowed never
module Main where

-- @dependencies: prelude effect console refs arrays foldable-traversable

import Prelude
import Data.Array as Array
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

-- Deliberately unrelated to benchmark names and constructor spelling.
data Trail a = Stop | Cell a (Trail a)

makeTrail :: Int -> Int -> Trail Int
makeTrail lower upper = collect upper Stop
  where
  collect current result =
    if current < lower then result
    else collect (current - 1) (Cell current result)

turnAround :: forall a. Trail a -> Trail a
turnAround input = collect input Stop
  where
  collect Stop result = result
  collect (Cell value rest) result = collect rest (Cell value result)

selectStable :: forall a. (a -> Boolean) -> Trail a -> Trail a
selectStable predicate input = collect input Stop
  where
  collect Stop result = turnAround result
  collect (Cell value rest) result =
    if predicate value then collect rest (Cell value result)
    else collect rest result

eliminate :: Trail Int -> Trail Int
eliminate Stop = Stop
eliminate (Cell leader rest) =
  Cell leader (eliminate (selectStable (\value -> not (value `mod` leader == 0)) rest))

eliminateDirect :: Trail Int -> Trail Int
eliminateDirect Stop = Stop
eliminateDirect (Cell leader rest) =
  Cell leader (eliminateDirect (selectStable (\value -> value `mod` leader /= 0) rest))

eliminateUsing :: (Int -> Int -> Boolean) -> Trail Int -> Trail Int
eliminateUsing _ Stop = Stop
eliminateUsing predicate (Cell leader rest) =
  Cell leader (eliminateUsing predicate (selectStable (predicate leader) rest))

sumFrom :: Int -> Trail Int -> Int
sumFrom seed input = collect input seed
  where
  collect Stop result = result
  collect (Cell value rest) result = collect rest (result + value)

sumTrail :: Trail Int -> Int
sumTrail input = sumFrom 0 input

sumRange :: Int -> Int -> Int -> Int
sumRange lower upper seed = sumFrom seed (eliminate (makeTrail lower upper))

score :: Int -> Int -> Int
score lower upper = sumTrail (eliminate (makeTrail lower upper))

scoreDirect :: Int -> Int -> Int
scoreDirect lower upper = sumTrail (eliminateDirect (makeTrail lower upper))

-- Preserve a materialized reference with a genuinely borrowed list argument.
sumBorrowed :: Int -> Trail Int -> Int
sumBorrowed seed input = sumFrom seed (eliminate input)

fingerprint :: Trail Int -> Int
fingerprint input = collect input 0
  where
  collect Stop result = result
  collect (Cell value rest) result = collect rest (result * 31 + value)

orderedRange :: Int -> Int -> Int
orderedRange lower upper = fingerprint (eliminate (makeTrail lower upper))

toArray :: forall a. Trail a -> Array a
toArray Stop = []
toArray (Cell value rest) = Array.cons value (toArray rest)

type Trace = Ref.Ref (Array (Array Int))

recordValue :: Trace -> Int -> Int -> Int
recordValue reference tag value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ tag, value ]) reference
  pure value

recordDecision :: Trace -> Trail Int -> Int -> Int -> Boolean
recordDecision reference original leader value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events [ leader, value, fingerprint original ]) reference
  pure (value `mod` leader /= 0)

observedPureRange :: Trace -> Int -> Int -> Int -> Int
observedPureRange reference lower upper seed =
  sumFrom (recordValue reference 0 seed)
    (eliminate (makeTrail (recordValue reference 1 lower) (recordValue reference 2 upper)))

recordArgument :: Trace -> Int -> Int -> Effect Int
recordArgument reference tag value = do
  Ref.modify_ (\events -> Array.snoc events [ tag, value ]) reference
  pure value

-- Explicit Effect sequencing defines the observable order of captures. The
-- pure factory above remains available to check that opaque calls are refused.
sequencedRange :: Trace -> Int -> Int -> Int -> Effect Int
sequencedRange reference lower upper seed = do
  capturedSeed <- recordArgument reference 0 seed
  capturedLower <- recordArgument reference 1 lower
  capturedUpper <- recordArgument reference 2 upper
  pure (sumFrom capturedSeed (eliminate (makeTrail capturedLower capturedUpper)))

main :: Effect Unit
main = for_
  [ { lower: 4, upper: 3, seed: 123 }
  , { lower: 2, upper: 2, seed: 0 }
  , { lower: 2, upper: 3, seed: 17 }
  , { lower: 2, upper: 30, seed: 0 }
  , { lower: 4, upper: 10, seed: 7 }
  , { lower: 6, upper: 18, seed: -17 }
  , { lower: 10, upper: 30, seed: 2147483647 }
  , { lower: 1, upper: 10, seed: -2147483648 }
  , { lower: 0, upper: 8, seed: 23 }
  , { lower: -4, upper: 4, seed: 17 }
  , { lower: -8, upper: -1, seed: -2147483648 }
  -- MinInt32 itself is excluded: this producer would wrap before terminating.
  , { lower: -2147483647, upper: -2147483645, seed: 0 }
  , { lower: 2147483644, upper: 2147483647, seed: 2147483647 }
  , { lower: 2147483647, upper: 2147483647, seed: -2147483648 }
  , { lower: 2, upper: 500, seed: 0 }
  ] \input -> do
    dynamic <- Ref.new input
    arguments <- Ref.read dynamic
    let
      source = makeTrail arguments.lower arguments.upper
      selected = eliminate source
    logShow [ toArray source, toArray selected, toArray source ]
    logShow
      [ sumRange arguments.lower arguments.upper arguments.seed
      , score arguments.lower arguments.upper
      , scoreDirect arguments.lower arguments.upper
      , sumBorrowed arguments.seed source
      , orderedRange arguments.lower arguments.upper
      ]
    visits <- Ref.new []
    predicateReference <- Ref.new (recordDecision visits source)
    predicate <- Ref.read predicateReference
    logShow (sumFrom arguments.seed (eliminateUsing predicate source))
    calls <- Ref.read visits
    logShow calls
    -- The original remains available after all opaque predicate invocations.
    logShow (toArray source)
    captures <- Ref.new []
    capturedResult <- sequencedRange captures arguments.lower arguments.upper arguments.seed
    logShow capturedResult
    captured <- Ref.read captures
    logShow captured
