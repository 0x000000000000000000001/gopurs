-- @dependencies: prelude effect console refs arrays foldable-traversable maybe either integers exceptions
-- @inline Main.weighted never
-- @inline Main.seven never
-- @inline Main.scalar never
-- @inline Main.observe never
-- @inline Main.failLeft never
module Main where

import Prelude
import Data.Array as Array
import Data.Either (Either(..))
import Data.Foldable (for_)
import Data.FoldableWithIndex (foldlWithIndex)
import Data.Int.Bits as Bits
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Traversable (traverse)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Exception (message, throw, try)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

weighted :: Array Int -> Int
weighted = foldlWithIndex (\index total value -> total + (1 + index `mod` 2) * value) 0

score :: Maybe (Array Int) -> Int
score Nothing = -1
score (Just values) = weighted values

seven :: Array Int -> Array Int -> Array Int
seven left right =
  [ weighted (map (_ + 1) left)
  , weighted (left <> right)
  , score (traverse (\value -> Just (value + 1)) left)
  , score (traverse (\value -> if value == fromMaybe 19 (Array.head left) then Nothing else Just (value + 1)) left)
  , score (traverse (\value -> if value == fromMaybe 19 (Array.last left) then Nothing else Just (value + 1)) left)
  , weighted (Array.filter (\value -> value `mod` 2 == 0) left)
  , weighted (Array.sort left)
  ]

scalar :: Int -> Array Int
scalar value =
  [ value `mod` 2, value `mod` 4, value `mod` 8
  , value `mod` (-2), value `mod` 0, value * value
  , (1 + value `mod` 2) * value, value * (1 + value `mod` 2)
  , (value `mod` 2) * value
  , (value `mod` 4 - 2) * value, value * (value `mod` 4 - 2)
  ]

observe :: Ref.Ref (Array Int) -> Int -> Int -> Int
observe trace tag value = unsafePerformEffect do
  Ref.modify_ (\items -> Array.snoc items tag) trace
  pure value

failLeft :: Ref.Ref (Array Int) -> Int -> Int
failLeft trace _ = unsafePerformEffect do
  Ref.modify_ (\items -> Array.snoc items 3) trace
  throw "left operand failure"

main :: Effect Unit
main = do
  minimumRef <- Ref.new (-2147483648)
  maximumRef <- Ref.new 2147483647
  minimum <- Ref.read minimumRef
  maximum <- Ref.read maximumRef
  negativeRef <- Ref.new (-1)
  negative <- Ref.read negativeRef
  shiftRef <- Ref.new Bits.zshr
  shift <- Ref.read shiftRef
  let unsigned = shift negative 0
  let quotient = minimum `div` (-1)
  casesRef <- Ref.new
    [ { left: [], right: [] }
    , { left: [ 1 ], right: [] }
    , { left: [ 1, 2, 3, 4, 5 ], right: [ -1, 0, 2 ] }
    , { left: [ -7, 0, -7, 4, -2 ], right: [ 1, 1 ] }
    , { left: [ minimum, maximum, unsigned, quotient, 0, -1 ], right: [ unsigned, minimum ] }
    , { left: Array.range (-17) 17, right: Array.range 7 1 }
    ]
  cases <- Ref.read casesRef
  for_ cases \sample -> do
    logShow (seven sample.left sample.right)
    logShow (seven sample.left sample.right)
    logShow [ sample.left, sample.right ]
  scalarsRef <- Ref.new [ minimum, minimum + 1, -65537, -9, -2, -1, 0, 1, 2, 3, 7, 65535, maximum, unsigned, quotient ]
  scalars <- Ref.read scalarsRef
  for_ scalars \value -> logShow (scalar value)
  trace <- Ref.new []
  logShow ((1 + observe trace 1 (-3) `mod` 2) * observe trace 2 unsigned)
  Ref.read trace >>= logShow
  failed <- try do
    current <- Ref.read negativeRef
    pure ((1 + failLeft trace current `mod` 2) * observe trace 4 unsigned)
  logShow (case failed of
    Left error -> message error
    Right value -> show value)
  Ref.read trace >>= logShow
