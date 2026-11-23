-- @inline export coordinateSum never
module Main where

-- @dependencies: assert lists tuples prelude effect console refs
-- @snapshot-ffi

import Prelude

import Data.List (List(..))
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)
import Worker (Phase(..), Stack, invoke, keep, mapStack, prepend, shift, summarize)

foreign import echoPayload :: forall a. a -> a

main :: Effect Unit
main = do
  seedRef <- Ref.new 5
  seed <- Ref.read seedRef
  let
    original =
      Cons (Tuple { byteOffset: seed, rowIndex: 2 } Begin)
        (Cons (Tuple { byteOffset: -4, rowIndex: 7 } Between)
          (Cons (Tuple { byteOffset: 9, rowIndex: -3 } End) Nil))
  originalRef <- Ref.new original
  callbackRef <- Ref.new (shift 10)
  callback <- Ref.read callbackRef
  let changed = invoke @Stack @Stack callback original
  assertEqual { expected: "15:2:B;-4:7:M;9:-3:E;", actual: summarize changed }
  assertEqual { expected: "5:2:B;-4:7:M;9:-3:E;", actual: summarize original }
  assertEqual { expected: "", actual: summarize (invoke @Stack @Stack callback Nil) }

  -- A generic foreign boundary must retain the complete list and enum payloads.
  let echoed = echoPayload changed
  assertEqual { expected: changed, actual: echoed }
  assertEqual { expected: Nil, actual: echoPayload (Nil :: Stack) }

  -- A callback return and a later invocation of a captured closure cross the
  -- generic ABI independently of the native producer and consumer functions.
  mappedRef <- Ref.new (mapStack (\(Tuple position phase) -> Tuple (position { rowIndex = position.rowIndex + 1 }) phase))
  mapped <- Ref.read mappedRef
  let moved = invoke @Stack @Stack mapped echoed
  assertEqual { expected: "15:3:B;-4:8:M;9:-2:E;", actual: summarize moved }
  savedRef <- Ref.new (keep moved)
  saved <- Ref.read savedRef
  assertEqual { expected: moved, actual: saved unit }

  -- Extending two versions must keep the shared tail and the original valid.
  let
    first = prepend { byteOffset: -8, rowIndex: 12 } End original
    second = prepend { byteOffset: 0, rowIndex: 0 } Between original
  assertEqual { expected: "-8:12:E;5:2:B;-4:7:M;9:-3:E;", actual: summarize first }
  assertEqual { expected: "0:0:M;5:2:B;-4:7:M;9:-3:E;", actual: summarize second }
  preserved <- Ref.read originalRef
  assertEqual { expected: original, actual: preserved }

  -- An open-row consumer and a richer payload must preserve their extra fields,
  -- including reference sharing, regardless of which representation is chosen.
  cell <- Ref.new 3
  let shared = { byteOffset: 4, rowIndex: 8, cell, label: "retained" }
  widerRef <- Ref.new (Cons (Tuple shared Begin) (Cons (Tuple shared End) Nil))
  wider <- Ref.read widerRef
  genericRef <- Ref.new identity
  generic <- Ref.read genericRef
  let roundtrip = echoPayload (invoke generic wider)
  case roundtrip of
    Cons (Tuple firstRow _) (Cons (Tuple secondRow _) Nil) -> do
      Ref.write 19 firstRow.cell
      observed <- Ref.read secondRow.cell
      assertEqual { expected: 19, actual: observed }
      assertEqual { expected: "retained", actual: firstRow.label }
      assertEqual { expected: 12, actual: coordinateSum secondRow }
    _ -> assertEqual { expected: true, actual: false }
  originalCell <- Ref.read cell
  assertEqual { expected: 19, actual: originalCell }
  log "Done"

coordinateSum :: forall r. { byteOffset :: Int, rowIndex :: Int | r } -> Int
coordinateSum row = row.byteOffset + row.rowIndex
