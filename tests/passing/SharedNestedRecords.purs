module Main where

-- @dependencies: assert prelude effect console refs
-- @snapshot-ffi

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)
import Worker (Packet, Payload(..), advance, invoke, keep, makePacket, readColumn, score)

foreign import echoPayload :: forall a. a -> a

main :: Effect Unit
main = do
  seedRef <- Ref.new 5
  seed <- Ref.read seedRef
  let original = makePacket seed
  originalRef <- Ref.new original
  callbackRef <- Ref.new (advance 3)
  callback <- Ref.read callbackRef
  let changed = invoke @(Packet Int) @(Packet Int) callback original
  assertEqual { expected: 35, actual: score original }
  assertEqual { expected: 15, actual: score (makePacket 0) }
  assertEqual { expected: 49, actual: score changed }
  assertEqual { expected: 5, actual: original.left.column }
  assertEqual { expected: 12, actual: original.right.score }
  assertEqual { expected: true, actual: original.enabled }
  assertEqual { expected: 8, actual: changed.left.column }
  assertEqual { expected: 24, actual: changed.right.score }
  assertEqual { expected: false, actual: changed.enabled }
  assertEqual { expected: "keep", actual: changed.right.label }

  -- Retaining a function return and crossing a foreign boundary must keep
  -- the complete record, including both nested records and the ADT payload.
  savedRef <- Ref.new (keep changed)
  saved <- Ref.read savedRef
  let echoed = echoPayload (saved unit)
  assertEqual { expected: 49, actual: score echoed }
  assertEqual { expected: 8, actual: readColumn echoed }
  assertEqual { expected: true, actual: echoed.left.active }
  assertEqual { expected: false, actual: echoed.right.active }
  case echoed.left.note of
    Payload value -> assertEqual { expected: 6, actual: value }
    _ -> assertEqual { expected: true, actual: false }

  -- Two updates of the same value preserve the previously retained version.
  let other = invoke @(Packet Int) @(Packet Int) (advance (-2)) original
  assertEqual { expected: 39, actual: score other }
  assertEqual { expected: 49, actual: score changed }
  preserved <- Ref.read originalRef
  assertEqual { expected: 35, actual: score preserved }
  assertEqual { expected: 3, actual: other.left.column }

  -- The same polymorphic family can hold a reference inside its opaque ADT.
  -- Boxing and unboxing must preserve that reference's identity.
  cell <- Ref.new 7
  let withCell =
        { payload: Payload cell
        , left:
            { column: original.left.column
            , line: original.left.line
            , active: original.left.active
            , note: Payload cell
            }
        , right: original.right
        , enabled: original.enabled
        }
  retainedRef <- Ref.new (keep withCell)
  retained <- Ref.read retainedRef
  let returned = echoPayload (retained unit)
  case returned.payload of
    Payload shared -> Ref.write 23 shared
    _ -> assertEqual { expected: true, actual: false }
  case returned.left.note of
    Payload shared -> do
      sharedValue <- Ref.read shared
      assertEqual { expected: 23, actual: sharedValue }
    _ -> assertEqual { expected: true, actual: false }
  observed <- Ref.read cell
  assertEqual { expected: 23, actual: observed }

  -- An open-row consumer must also accept a richer record without dropping
  -- its extra field at a generic foreign boundary.
  let wider =
        { payload: original.payload
        , left: original.left
        , right: original.right
        , enabled: original.enabled
        , extra: "retained"
        }
  genericRef <- Ref.new identity
  generic <- Ref.read genericRef
  let roundtrip = echoPayload (invoke generic wider)
  assertEqual { expected: 5, actual: readColumn roundtrip }
  assertEqual { expected: "retained", actual: roundtrip.extra }
  log "Done"
