-- @dependencies: prelude effect console refs arrays foldable-traversable either exceptions
-- @inline Main.mapped never
-- @inline Main.appended never
-- @inline Main.mapper never
-- @inline Main.factory never
-- @inline Main.withUnknown never
-- @inline Main.withFactories never
-- @inline Main.captured never
-- @inline Main.shared never
module Main where

import Prelude
import Data.Array as Array
import Data.Either (Either(..))
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Exception (message, throw, try)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Subject as Subject

type Trace = Ref.Ref (Array Int)

mapped :: Ref.Ref (Array Int) -> Effect Int
mapped reference = do
  values <- Ref.read reference
  pure (Subject.consume (map (_ + 1) values))

appended :: Ref.Ref (Array Int) -> Ref.Ref (Array Int) -> Effect Int
appended left right = do
  xs <- Ref.read left
  ys <- Ref.read right
  pure (Subject.consume (xs <> ys))

captured :: Int -> Array Int -> Int
captured delta values = Subject.consume (map (\value -> value + delta) values)

shared :: Array Int -> Int
shared values =
  let produced = map (_ + 1) values
  in Subject.consume produced + Array.length produced

mapper :: Trace -> Boolean -> Int -> Int
mapper trace failing value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events value) trace
  if failing && value == 2 then throw "mapper failure" else pure (value + 1)

withUnknown :: Trace -> Boolean -> Array Int -> Int
withUnknown trace failing values = Subject.consume (map (mapper trace failing) values)

factory :: Trace -> Int -> Array Int -> Array Int
factory trace tag values = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events tag) trace
  pure values

withFactories :: Trace -> Array Int -> Array Int -> Int
withFactories trace xs ys = Subject.consume (factory trace 1 xs <> factory trace 2 ys)

main :: Effect Unit
main = do
  for_ [ [], [ 1, 2, 3 ], [ 2147483647, -2147483648, -1, 0 ] ] \input -> do
    left <- Ref.new input
    right <- Ref.new (Array.reverse input)
    a <- mapped left
    b <- appended left right
    values <- Ref.read left
    logShow [ a, b, captured 3 values, shared values ]
    Ref.read left >>= logShow
    Ref.read right >>= logShow
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
  logShow (withFactories trace values [ 4, 5 ])
  Ref.read trace >>= logShow
  Ref.write [] trace
  empty <- Ref.new []
  current <- Ref.read empty
  logShow (withUnknown trace true current)
  Ref.read trace >>= logShow
