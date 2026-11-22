-- @inline export appendInt never
-- @inline export appendItem never
-- @inline export applyAppend never
module Main where

-- @dependencies: assert arrays lists prelude effect console refs

import Prelude

import Data.Array as Array
import Data.List (List(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

type Position = { column :: Int, line :: Int }
type Entry = { label :: String, position :: Position, values :: Array Int }

data Item = Empty | Item Entry (List Int)

describe :: Item -> String
describe Empty = "empty"
describe (Item entry stack) =
  entry.label <> ":" <> show entry.position <> ":" <> show entry.values <> ":" <> show stack

appendInt :: Array Int -> Int -> Array Int
appendInt = Array.snoc

appendItem :: Array Item -> Item -> Array Item
appendItem = Array.snoc

-- Keep an opaque higher-order route as well as specialized direct calls.
applyAppend :: forall a. (Array a -> a -> Array a) -> Array a -> a -> Array a
applyAppend append values value = append values value

main :: Effect Unit
main = do
  inputsRef <- Ref.new
    { ints: [ -2147483648, 0, 2147483647 ]
    , numbers: [ 1.25, -2.5 ]
    , strings: [ "", "λ😀" ]
    , booleans: [ true, false ]
    , arrays: [ [], [ 1, 2 ] ]
    , lists: [ Nil, Cons 3 (Cons 5 Nil) ]
    , entry: { label: "kept", position: { column: 7, line: 11 }, values: [ 13, 17 ] }
    , functions: [ \n -> n + 1, \n -> n * 2 ]
    }
  input <- Ref.read inputsRef

  assertEqual { expected: [ -2147483648, 0, 2147483647, -7 ], actual: appendInt input.ints (-7) }
  assertEqual { expected: [ 42 ], actual: appendInt [] 42 }
  assertEqual { expected: [ 1.25, -2.5, 3.75 ], actual: Array.snoc input.numbers 3.75 }
  assertEqual { expected: [ "", "λ😀", "last" ], actual: Array.snoc input.strings "last" }
  assertEqual { expected: [ true, false, true ], actual: Array.snoc input.booleans true }
  assertEqual { expected: [ [], [ 1, 2 ], [ 3 ] ], actual: Array.snoc input.arrays [ 3 ] }
  assertEqual
    { expected: [ Nil, Cons 3 (Cons 5 Nil), Cons 7 Nil ]
    , actual: Array.snoc input.lists (Cons 7 Nil)
    }
  assertEqual
    { expected: [ 11, 20, 7 ]
    , actual: map (\f -> f 10) (Array.snoc input.functions (\n -> n - 3))
    }

  let
    extra = { label: "new", position: { column: 19, line: 23 }, values: [ 29 ] }
    records = Array.snoc [ input.entry ] extra
    originalItems = [ Empty, Item input.entry (Cons 31 Nil) ]
    appendedItems = appendItem originalItems (Item extra Nil)
  assertEqual { expected: [ input.entry, extra ], actual: records }
  assertEqual
    { expected: map describe [ Empty, Item input.entry (Cons 31 Nil), Item extra Nil ]
    , actual: map describe appendedItems
    }
  assertEqual { expected: [ "empty" ], actual: map describe (appendItem [] Empty) }
  assertEqual { expected: 2, actual: Array.length originalItems }

  -- A saved partial application can be reused without accumulating prior appends.
  partialRef <- Ref.new (appendInt input.ints)
  partial <- Ref.read partialRef
  assertEqual { expected: [ -2147483648, 0, 2147483647, 5 ], actual: partial 5 }
  assertEqual { expected: [ -2147483648, 0, 2147483647, 9 ], actual: partial 9 }
  aliasRef <- Ref.new (Array.snoc :: Array String -> String -> Array String)
  alias <- Ref.read aliasRef
  assertEqual { expected: [ "", "λ😀", "alias" ], actual: applyAppend alias input.strings "alias" }
  assertEqual { expected: [ "only" ], actual: applyAppend alias [] "only" }

  -- Retain one version while extending it twice. The old versions stay intact.
  firstRef <- Ref.new (appendInt input.ints 5)
  first <- Ref.read firstRef
  assertEqual { expected: [ -2147483648, 0, 2147483647, 5, 6 ], actual: appendInt first 6 }
  assertEqual { expected: [ -2147483648, 0, 2147483647, 5, 8 ], actual: appendInt first 8 }
  assertEqual { expected: [ -2147483648, 0, 2147483647, 5 ], actual: first }
  preserved <- Ref.read inputsRef
  assertEqual { expected: [ -2147483648, 0, 2147483647 ], actual: preserved.ints }
  assertEqual { expected: [ [], [ 1, 2 ] ], actual: preserved.arrays }
  assertEqual { expected: [ "", "λ😀" ], actual: preserved.strings }

  -- Appending copies the array container while retaining reference-valued fields.
  shared <- Ref.new 10
  other <- Ref.new 20
  cellsRef <- Ref.new [ { label: "shared", cell: shared } ]
  cells <- Ref.read cellsRef
  let withOther = Array.snoc cells { label: "other", cell: other }
  case withOther of
    [ firstCell, secondCell ] -> do
      Ref.write 41 firstCell.cell
      Ref.write 42 secondCell.cell
    _ -> assertEqual { expected: 2, actual: Array.length withOther }
  sharedValue <- Ref.read shared
  otherValue <- Ref.read other
  assertEqual { expected: 41, actual: sharedValue }
  assertEqual { expected: 42, actual: otherValue }
  assertEqual { expected: 1, actual: Array.length cells }
  log "Done"
