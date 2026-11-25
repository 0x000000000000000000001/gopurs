-- @inline Main.keepFirst never
-- @inline Main.projectKept never
-- @inline Main.observeValue never
-- @inline Main.invoke never
-- @inline Main.accumulateWide never
module Main where

-- @dependencies: prelude effect console refs arrays typelevel-prelude

import Prelude
import Data.Array as Array
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref
import Prim.RowList (class RowToList, RowList, Cons, Nil)
import Type.Proxy (Proxy(..))

-- Independent names and several row shapes exercise the general dictionary
-- interpreter without a benchmark-specific result or record layout.
class Cardinality (fields :: RowList Type) where
  cardinalityImpl :: Proxy fields -> Int

instance cardinalityNil :: Cardinality Nil where
  cardinalityImpl _ = 0

instance cardinalityCons :: Cardinality tail => Cardinality (Cons label field tail) where
  cardinalityImpl _ = 1 + cardinalityImpl (Proxy :: Proxy tail)

-- | @inline always
cardinality :: forall row fields. RowToList row fields => Cardinality fields => Record row -> Int
cardinality _ = cardinalityImpl (Proxy :: Proxy fields)

keepFirst :: forall a b. a -> b -> a
keepFirst value _ = value

projectKept :: forall a. { kept :: Int, discarded :: a } -> Int
projectKept value = value.kept

invoke :: (Int -> Int) -> Int -> Int
invoke callback value = callback value

observeValue :: Ref.Ref (Array Int) -> Int -> Int -> Effect Int
observeValue events tag value = do
  Ref.modify_ (\previous -> Array.snoc previous tag) events
  pure value

lexicalCapture :: Int -> Int
lexicalCapture initial =
  let captured = initial + 1
      addCaptured argument = captured + argument
      shadow captured = addCaptured captured
  in shadow 5

curried :: Int -> Int -> Int -> Int
curried first second third = first + second + third

moduleBase :: Int
moduleBase = 23

addModuleBase :: Int -> Int
addModuleBase value = moduleBase + value

-- A closed recurrence exceeds the scalar interpreter's ordinary fuel. Its
-- repeated Int additions must keep wrapping exactly, including negative seeds.
accumulateWide :: Int -> Int -> Int
accumulateWide remaining total
  | remaining == 0 = total
  | otherwise = accumulateWide (remaining - 1) (total + 2147483647)

main :: Effect Unit
main = do
  events <- Ref.new []
  Ref.modify_ (\previous -> Array.snoc previous 1) events
  logShow
    [ cardinality {}
    , cardinality { z: 7 }
    , cardinality { first: false, last: "x" }
    , cardinality { one: 1, two: 2, three: 3, four: 4 }
    ]
  Ref.modify_ (\previous -> Array.snoc previous 2) events
  dynamic <- Ref.new 10
  initial <- Ref.read dynamic
  logShow [ lexicalCapture 10, lexicalCapture initial, curried 3 4 5, (curried initial 2) 3, addModuleBase initial ]
  maximum <- Ref.new 2147483647
  maxInt <- Ref.read maximum
  logShow [ invoke (\value -> value + 1) maxInt, invoke (\value -> value + (-1)) (-2147483648) ]
  repetitions <- Ref.new 5000
  runtimeCount <- Ref.read repetitions
  logShow
    [ accumulateWide 5000 0
    , accumulateWide 5000 (-3)
    , accumulateWide 70000 3
    , accumulateWide 0 2147483647
    , accumulateWide 0 (-17)
    , accumulateWide runtimeCount (-19)
    ]
  ignored <- observeValue events 10 99
  logShow (keepFirst 41 ignored)
  discarded <- observeValue events 20 99
  logShow (projectKept { kept: 7, discarded })
  -- Constructing an unused closure must not execute its body; applying the
  -- callback below must retain exactly one call to the opaque effect.
  logShow (keepFirst 19 (\_ -> observeValue events 99 0))
  callbacks <- Ref.new (observeValue events 30)
  callback <- Ref.read callbacks
  result <- callback 5
  logShow result
  Ref.read events >>= logShow
