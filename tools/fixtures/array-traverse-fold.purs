-- @dependencies: prelude effect console refs arrays foldable-traversable either exceptions maybe
-- @inline Main.direct never
-- @inline Main.indexed never
-- @inline Main.mapper never
-- @inline Main.factory never
-- @inline Main.withUnknown never
-- @inline Main.withFactory never
-- @inline Main.withBarrier never
-- @inline Main.captured never
-- @inline Main.shared never
-- @inline Main.mayFail never
module Main where

import Prelude
import Data.Array as Array
import Data.Either (Either(..))
import Data.Foldable (for_)
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Data.TraversableWithIndex (traverseWithIndex)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Exception (message, throw, try)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Subject as Subject

type Trace = Ref.Ref (Array Int)

direct :: Ref.Ref (Array Int) -> Effect Int
direct reference = do
  values <- Ref.read reference
  pure (case traverse (\value -> Just (value + 1)) values of
    Nothing -> -1
    Just produced -> Subject.consume produced)

indexed :: Ref.Ref (Array Int) -> Effect Int
indexed reference = do
  values <- Ref.read reference
  pure (case traverseWithIndex (\index value -> Just (value + index)) values of
    Nothing -> -1
    Just produced -> Subject.consume produced)

captured :: Int -> Array Int -> Int
captured delta values = case traverse (\value -> Just (value + delta)) values of
  Nothing -> -1
  Just produced -> Subject.consume produced

shared :: Array Int -> Int
shared values = case traverse (\value -> Just (value + 1)) values of
  Nothing -> -1
  Just produced -> Subject.consume produced + Array.length produced

mayFail :: Array Int -> Int
mayFail values = case traverse (\value -> if value == 2 then Nothing else Just (value + 1)) values of
  Nothing -> -1
  Just produced -> Subject.consume produced

mapper :: Trace -> Boolean -> Int -> Maybe Int
mapper trace failing value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events value) trace
  if failing && value == 2 then throw "mapper failure" else pure (Just (value + 1))

withUnknown :: Trace -> Boolean -> Array Int -> Int
withUnknown trace failing values = case traverse (mapper trace failing) values of
  Nothing -> -1
  Just produced -> Subject.consume produced

factory :: Trace -> Int -> Array Int -> Array Int
factory trace tag values = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events tag) trace
  pure values

withFactory :: Trace -> Array Int -> Int
withFactory trace values = case traverse (\value -> Just (value + 1)) (factory trace 1 values) of
  Nothing -> -1
  Just produced -> Subject.consume produced

withBarrier :: Trace -> Array Int -> Effect Int
withBarrier trace values = do
  let result = traverse (\value -> Just (value + 1)) values
  Ref.modify_ (\events -> Array.snoc events 9) trace
  pure (case result of
    Nothing -> -1
    Just produced -> Subject.consume produced)

main :: Effect Unit
main = do
  for_ [ [], [ 1, 2, 3 ], [ 2147483647, -2147483648, -1, 0 ] ] \input -> do
    reference <- Ref.new input
    a <- direct reference
    b <- indexed reference
    values <- Ref.read reference
    logShow [ a, b, captured 3 values, shared values, mayFail values ]
    Ref.read reference >>= logShow
    logShow (Array.reverse input)
  trace <- Ref.new []
  input <- Ref.new [ 1, 2, 3 ]
  values <- Ref.read input
  logShow (withUnknown trace false values)
  Ref.read trace >>= logShow
  Ref.write [] trace
  failure <- try do
    current <- Ref.read input
    pure (withUnknown trace true current)
  logShow (case failure of
    Left error -> message error
    Right value -> show value)
  Ref.read trace >>= logShow
  Ref.write [] trace
  logShow (withFactory trace values)
  Ref.read trace >>= logShow
  Ref.write [] trace
  empty <- Ref.new []
  current <- Ref.read empty
  logShow (withUnknown trace true current)
  Ref.read trace >>= logShow
  withBarrier trace values >>= logShow
  Ref.read trace >>= logShow
