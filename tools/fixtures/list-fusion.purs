module Main where

-- @dependencies: prelude effect console refs foldable-traversable

import Prelude
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref

-- Deliberately different names from the ListOps benchmark: admission must
-- depend on the producer/consumer bodies and their types.
data Chain a = End | Link a (Chain a)

fromTo :: Int -> Int -> Chain Int
fromTo lower upper = descend upper End
  where
  descend current result =
    if current < lower then result
    else descend (current - 1) (Link current result)

keepEven :: Chain Int -> Chain Int
keepEven input = select input End
  where
  select End result = result
  select (Link value rest) result =
    if value `mod` 2 == 0 then select rest (Link value result)
    else select rest result

keepOdd :: Chain Int -> Chain Int
keepOdd input = select input End
  where
  select End result = result
  select (Link value rest) result =
    if value `mod` 2 /= 0 then select rest (Link value result)
    else select rest result

keepThird :: Chain Int -> Chain Int
keepThird input = select input End
  where
  select End result = result
  select (Link value rest) result =
    if value `mod` 3 == 0 then select rest (Link value result)
    else select rest result

keepResidue :: Chain Int -> Chain Int
keepResidue input = select input End
  where
  select End result = result
  select (Link value rest) result =
    if value `mod` 5 == 3 then select rest (Link value result)
    else select rest result

reduce :: forall a b. (b -> a -> b) -> b -> Chain a -> b
reduce _ initial End = initial
reduce combine initial (Link value rest) =
  reduce combine (combine initial value) rest

transform :: (Int -> Int) -> Chain Int -> Chain Int
transform _ End = End
transform change (Link value rest) = Link (change value) (transform change rest)

score :: Int -> Int -> Int
score upper seed = reduce (+) seed (keepEven (fromTo 1 upper))

oddScore :: Int -> Int -> Int
oddScore upper seed = reduce (+) seed (keepOdd (fromTo (-8) upper))

thirdScore :: Int -> Int -> Int
thirdScore upper seed = reduce (+) seed (keepThird (fromTo (-5) upper))

residueScore :: Int -> Int -> Int
residueScore upper seed = reduce (+) seed (keepResidue (fromTo 1 upper))

-- The same recurrence is deliberately written without any list operations.
scalarResidueScore :: Int -> Int -> Int
scalarResidueScore current seed =
  if current < 1 then seed
  else scalarResidueScore (current - 1)
    (if current `mod` 5 == 3 then seed + current else seed)

-- The interval stays small even at MaxInt32, while the selected sum overflows.
nearMaxScore :: Int -> Int -> Int
nearMaxScore upper seed = reduce (+) seed (keepResidue (fromTo 2147483631 upper))

-- Both pipelines belong to one binding: their fresh local identities must
-- remain disjoint when the pass rewrites both children of the addition.
combinedScore :: Int -> Int -> Int
combinedScore upper seed =
  reduce (+) seed (keepEven (fromTo 1 upper))
    + reduce (+) (-17) (keepOdd (fromTo (-8) upper))

mappedScore :: Int -> Int -> Int -> Int
mappedScore lower upper seed =
  reduce (+) seed (transform (\value -> value * value - 3 * value + 2) (keepEven (fromTo lower upper)))

-- A noncommutative fold makes the filter's reversed order observable.
fingerprint :: Int -> Int -> Int
fingerprint accumulated value = accumulated * 31 + value

orderedScore :: Int -> Int -> Int
orderedScore upper seed = reduce fingerprint seed (keepEven (fromTo 1 upper))

probe :: Int -> Int -> Int -> Array Int
probe lower upper seed =
  let
    source = fromTo lower upper
    selected = keepEven source
    sharedScore = reduce (+) seed selected
  in
    [ score upper seed
    , oddScore upper seed
    , thirdScore upper seed
    , mappedScore lower upper seed
    , orderedScore upper seed
    , sharedScore
    -- Reuse both list values after their first consumption. A fast path must
    -- preserve all nodes visible through these aliases.
    , reduce fingerprint 0 source
    , reduce fingerprint 0 selected
    , reduce (+) 0 source
    , combinedScore upper seed
    , residueScore upper seed
    , scalarResidueScore upper seed
    ]

main :: Effect Unit
main = do
  for_
    [ { lower: 4, upper: 3, seed: 123 }
    , { lower: 2, upper: 2, seed: 0 }
    , { lower: 3, upper: 3, seed: 7 }
    , { lower: 0, upper: 0, seed: -7 }
    , { lower: -7, upper: -1, seed: 0 }
    , { lower: -8, upper: 8, seed: 17 }
    , { lower: 1, upper: 9, seed: 2147483647 }
    , { lower: 1, upper: 10, seed: -2147483648 }
    , { lower: 1, upper: 900, seed: 0 }
    , { lower: -7, upper: -1, seed: -2147483648 }
    , { lower: 1, upper: 900, seed: -17 }
    ] \input -> do
      -- Effectful reads keep the inputs dynamic through optimization.
      reference <- Ref.new input
      arguments <- Ref.read reference
      logShow (probe arguments.lower arguments.upper arguments.seed)
  for_
    [ { upper: 2147483630, seed: 17 }
    , { upper: 2147483631, seed: -2147483648 }
    , { upper: 2147483633, seed: 2147483647 }
    , { upper: 2147483640, seed: -2147483648 }
    , { upper: 2147483646, seed: 0 }
    , { upper: 2147483647, seed: 17 }
    , { upper: 2147483647, seed: -2147483648 }
    ] \input -> do
      reference <- Ref.new input
      arguments <- Ref.read reference
      logShow [ nearMaxScore arguments.upper arguments.seed ]
