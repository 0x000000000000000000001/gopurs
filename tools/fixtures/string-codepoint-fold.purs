-- @dependencies: prelude effect console refs arrays foldable-traversable maybe either integers exceptions strings enums
-- @inline Main.checksum never
-- @inline Main.explicit never
-- @inline Main.seeded never
-- @inline Main.weighted never
-- @inline Main.captured never
-- @inline Main.withCallback never
-- @inline Main.recordPoint never
-- @inline Main.ordered never
-- @inline Main.seedEffect never
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

checksum :: String -> Int
checksum = Array.foldl (\total cp -> total + fromEnum cp) 0 <<< CP.toCodePointArray

explicit :: String -> Int
explicit input = Array.foldl (\total cp -> total + fromEnum cp) 0 (CP.toCodePointArray input)

seeded :: Int -> String -> Int
seeded initial input = Array.foldl (\total cp -> total + fromEnum cp) initial (CP.toCodePointArray input)

weighted :: String -> Int
weighted = Array.foldl (\total cp -> total * 3 + fromEnum cp) 1 <<< CP.toCodePointArray

captured :: Int -> String -> Int
captured bias input = Array.foldl (\total cp -> total + fromEnum cp + bias) 0 (CP.toCodePointArray input)

withCallback :: (Int -> CP.CodePoint -> Int) -> Int -> String -> Int
withCallback callback initial input = Array.foldl callback initial (CP.toCodePointArray input)

point :: Int -> String
point value = CP.singleton (fromMaybe bottom (toEnum value))

recordPoint :: Ref.Ref (Array Int) -> Boolean -> Int -> CP.CodePoint -> Int
recordPoint trace fail total cp = unsafePerformEffect do
  Ref.modify_ (\values -> Array.snoc values (fromEnum cp)) trace
  if fail && fromEnum cp == 98 then throw "callback failure" else pure (total + fromEnum cp)

seedEffect :: Ref.Ref (Array Int) -> Boolean -> Int
seedEffect trace fail = unsafePerformEffect do
  Ref.modify_ (\values -> Array.snoc values 1) trace
  if fail then throw "seed failure" else pure 7

stringEffect :: Ref.Ref (Array Int) -> String -> String
stringEffect trace input = unsafePerformEffect do
  Ref.modify_ (\values -> Array.snoc values 2) trace
  pure input

ordered :: Ref.Ref (Array Int) -> Boolean -> String -> Int
ordered trace fail input = Array.foldl (\total cp -> total + fromEnum cp)
  (seedEffect trace fail) (CP.toCodePointArray (stringEffect trace input))

main :: Effect Unit
main = do
  casesRef <- Ref.new
    [ [], [ 0 ], [ 97, 98, 99 ], [ 233, 8364, 65535 ]
    , [ 128512, 1114111 ], [ 55296 ], [ 56320 ], [ 55296, 56320 ]
    , [ 56320, 55296 ], [ 55296, 97, 56320 ]
    , [ 55296, 55297, 56320 ], [ 128512, 55296, 56320, 0, 233 ]
    , Array.replicate 3000 1114111
    ]
  cases <- Ref.read casesRef
  checksumRef <- Ref.new checksum
  indirect <- Ref.read checksumRef
  maximumRef <- Ref.new 2147483647
  maximum <- Ref.read maximumRef
  for_ cases \values -> do
    inputRef <- Ref.new (String.joinWith "" (map point values))
    input <- Ref.read inputRef
    logShow [ checksum input, explicit input, indirect input, checksum input, seeded maximum input, weighted input, captured 3 input, Subject.checksum input ]
    logShow (map fromEnum (CP.toCodePointArray input))
    logShow [ checksum (CP.take 0 input), checksum (CP.take 1 input), checksum (CP.take 2 input), checksum (CP.take 10000 input) ]
  trace <- Ref.new []
  inputRef <- Ref.new "abc"
  input <- Ref.read inputRef
  callbackRef <- Ref.new (recordPoint trace false)
  callback <- Ref.read callbackRef
  logShow (withCallback callback 7 input)
  Ref.read trace >>= logShow
  Ref.write [] trace
  failingRef <- Ref.new (recordPoint trace true)
  failing <- Ref.read failingRef
  failed <- try do
    current <- Ref.read inputRef
    pure (withCallback failing 7 current)
  logShow (case failed of
    Left error -> message error
    Right value -> show value)
  Ref.read trace >>= logShow
  Ref.write [] trace
  emptyRef <- Ref.new ""
  empty <- Ref.read emptyRef
  logShow (withCallback failing maximum empty)
  Ref.read trace >>= logShow
  logShow (ordered trace false input)
  Ref.read trace >>= logShow
  Ref.write [] trace
  seedFailed <- try do
    current <- Ref.read inputRef
    pure (ordered trace true current)
  logShow (case seedFailed of
    Left error -> message error
    Right value -> show value)
  Ref.read trace >>= logShow
