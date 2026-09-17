-- Keep the guarded local recursion visible in the generated Go snapshot.
-- @inline export keepItems never
module Main where

-- @dependencies: assert lists prelude effect console refs

import Prelude

import Data.List (List(..))
import Data.List as List
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

type Item = { keep :: Boolean, value :: Int }

keepItems :: List Item -> List Item
keepItems = go Nil
  where
  go acc = case _ of
    Cons item rest | item.keep ->
      go (Cons item acc) rest
    Cons _ rest ->
      go acc rest
    Nil ->
      List.reverse acc

checkItems :: String -> List Item -> List Item -> Effect Unit
checkItems label expected original = do
  inputRef <- Ref.new original
  input <- Ref.read inputRef
  assertEqual { expected, actual: keepItems input }
  preserved <- Ref.read inputRef
  assertEqual { expected: original, actual: preserved }
  log label

checkRecordReturns :: Effect Unit
checkRecordReturns = do
  seedRef <- Ref.new 40
  seed <- Ref.read seedRef
  let
    makeItem :: Int -> Item
    makeItem offset = { keep: offset >= 0, value: seed + offset }

    makeOffset :: Int -> Int -> Item
    makeOffset first second = makeItem (first + second)

  assertEqual
    { expected: { keep: true, value: 42 }
    , actual: makeItem 2
    }

  -- The reference retains the local function as a value, exercising its
  -- wrapper as well as the direct call above.
  makerRef <- Ref.new makeItem
  maker <- Ref.read makerRef
  assertEqual
    { expected: { keep: true, value: 45 }
    , actual: maker 5
    }
  assertEqual
    { expected: { keep: false, value: 39 }
    , actual: maker (-1)
    }
  partialRef <- Ref.new (makeOffset 2)
  partial <- Ref.read partialRef
  assertEqual
    { expected: { keep: true, value: 45 }
    , actual: partial 3
    }
  log "local record returns"

main :: Effect Unit
main = do
  checkItems "guarded list preserves order"
    (Cons { keep: true, value: 7 } (Cons { keep: true, value: 3 } Nil))
    (Cons { keep: false, value: 0 }
      (Cons { keep: true, value: 7 }
        (Cons { keep: false, value: 2 }
          (Cons { keep: true, value: 3 } Nil))))
  checkItems "empty list" Nil Nil
  checkItems "all rejected" Nil
    (Cons { keep: false, value: 4 } (Cons { keep: false, value: 8 } Nil))
  checkRecordReturns
  log "Done"
