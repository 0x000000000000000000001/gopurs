-- @dependencies: prelude effect console refs arrays foldable-traversable maybe either integers exceptions strings enums
-- @inline Main.bounded never
-- @inline Main.shared never
-- @inline Main.captured never
-- @inline Main.withCallback never
-- @inline Main.ordered never
-- @inline Main.countEffect never
-- @inline Main.stringEffect never
module Main where

import Prelude
import Data.Array as Array
import Data.Either (Either(..))
import Data.Enum (fromEnum, toEnum)
import Data.Foldable (for_)
import Data.Maybe (fromMaybe)
import Data.String.CodePoints as CP
import Data.String.Common as String
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Exception (message, throw, try)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Subject as Subject

bounded :: Int -> String -> Int
bounded count input = Subject.consume (CP.take count input)

shared :: Int -> String -> Int
shared count input = let prefix = CP.take count input
  in Subject.consume prefix + CP.length prefix

captured :: Int -> Int -> String -> Int
captured bias count input = Array.foldl (\total cp -> total + fromEnum cp + bias) 7 (CP.toCodePointArray (CP.take count input))

withCallback :: (Int -> CP.CodePoint -> Int) -> Int -> String -> Int
withCallback callback count input = Array.foldl callback 7 (CP.toCodePointArray (CP.take count input))

point :: Int -> String
point value = CP.singleton (fromMaybe bottom (toEnum value))

type Trace = Ref.Ref (Array Int)

countEffect :: Trace -> Boolean -> Int -> Int
countEffect trace fail count = unsafePerformEffect do
  Ref.modify_ (\values -> Array.snoc values 1) trace
  if fail then throw "count failure" else pure count

stringEffect :: Trace -> Boolean -> String -> String
stringEffect trace fail input = unsafePerformEffect do
  Ref.modify_ (\values -> Array.snoc values 2) trace
  if fail then throw "input failure" else pure input

ordered :: Trace -> Boolean -> Boolean -> Int -> String -> Int
ordered trace failCount failInput count input =
  Subject.consume (CP.take (countEffect trace failCount count) (stringEffect trace failInput input))

main :: Effect Unit
main = do
  casesRef <- Ref.new [ [], [ 97, 98, 99 ], [ 233, 8364, 128512 ], [ 55296 ], [ 56320 ], [ 55296, 56320 ], [ 56320, 55296 ], [ 55296, 97, 56320 ] ]
  cases <- Ref.read casesRef
  countsRef <- Ref.new [ -2, 0, 1, 2, 3, 10, 2147483647 ]
  counts <- Ref.read countsRef
  functionRef <- Ref.new bounded
  indirect <- Ref.read functionRef
  callbackRef <- Ref.new (\total cp -> total + fromEnum cp)
  callback <- Ref.read callbackRef
  for_ cases \points -> do
    inputRef <- Ref.new (String.joinWith "" (map point points))
    input <- Ref.read inputRef
    for_ counts \count -> do
      current <- Ref.read inputRef
      logShow [ bounded count current, indirect count current, shared count current, captured 5 count current, withCallback callback count current ]
    logShow (map fromEnum (CP.toCodePointArray input))
  trace <- Ref.new []
  inputRef <- Ref.new "abc"
  for_ [ -1, 0, 2 ] \count -> do
    input <- Ref.read inputRef
    logShow (ordered trace false false count input)
    Ref.read trace >>= logShow
    Ref.write [] trace
  for_ [ true, false ] \failCount -> do
    result <- try do
      input <- Ref.read inputRef
      pure (ordered trace failCount (not failCount) 2 input)
    logShow (case result of
      Left error -> message error
      Right value -> show value)
    Ref.read trace >>= logShow
    Ref.write [] trace
