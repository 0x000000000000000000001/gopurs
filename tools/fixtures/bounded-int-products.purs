-- @inline Main.lowProduct never
-- @inline Main.highProduct never
-- @inline Main.exactBoundary never
-- @inline Main.roundedBoundary never
-- @inline Main.generalProduct never
-- @inline Main.assemble never
module Main where

-- @dependencies: prelude effect console refs arrays integers

import Prelude
import Data.Int.Bits as Bits
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref

lowProduct :: Int -> Int -> Int
lowProduct left right = Bits.and left 65535 * Bits.or right 0

highProduct :: Int -> Int -> Int
highProduct left right = Bits.shr left 16 * Bits.or right 0

exactBoundary :: Int -> Int
exactBoundary value = Bits.or value 0 * 4194304

roundedBoundary :: Int -> Int
roundedBoundary value = Bits.or value 0 * 4194305

generalProduct :: Int -> Int -> Int
generalProduct left right = left * right

identityStep :: (Int -> Int) -> Int -> Int
identityStep _ value = value

assemble :: Int -> (Int -> Int) -> Int -> Int
assemble 0 = identityStep
assemble remaining =
  let previous = assemble (remaining - 1)
  in \step value -> step (previous step value)

combine :: ((Int -> Int) -> Int -> Int) -> ((Int -> Int) -> Int -> Int) -> (Int -> Int) -> Int -> Int
combine left right step value = left (right step) value

square :: Int -> (Int -> Int) -> Int -> Int
square count = combine (assemble count) (assemble count)

fourth :: Int -> (Int -> Int) -> Int -> Int
fourth count = combine (square count) (square count)

consume :: Int -> Int
consume count = combine (fourth count) (assemble count) (\value -> value + 1) 0

main :: Effect Unit
main = do
  minimumRef <- Ref.new (-2147483648)
  maximumRef <- Ref.new 2147483647
  minimum <- Ref.read minimumRef
  maximum <- Ref.read maximumRef
  logShow [ lowProduct maximum maximum, lowProduct minimum maximum, lowProduct (-1) minimum ]
  logShow [ highProduct minimum maximum, highProduct maximum minimum, highProduct (-1) maximum ]
  logShow [ exactBoundary minimum, exactBoundary maximum ]
  logShow [ roundedBoundary maximum, generalProduct maximum maximum ]
  countRef <- Ref.new 5
  count <- Ref.read countRef
  logShow [ consume 0, consume 2, consume count ]
