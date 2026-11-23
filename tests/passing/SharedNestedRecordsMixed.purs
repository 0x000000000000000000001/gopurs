-- @inline export makePacket never
-- @inline export bump never
-- @inline export score never
-- @inline export rareIdentity never
-- @inline export run never
module Main where

-- @dependencies: assert prelude effect console refs

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

data Payload a = Vacant | Payload a | Tagged String a

type Packet a =
  { payload :: Payload a
  , left :: { column :: Int, line :: Int, active :: Boolean, note :: Payload a }
  , right :: { score :: Int, label :: String, active :: Boolean }
  , enabled :: Boolean
  }

makePacket :: Int -> Packet Int
makePacket seed =
  { payload: if seed < 0 then Vacant else if seed == 0 then Tagged "zero" 0 else Payload (seed * 2)
  , left: { column: seed, line: 3, active: true, note: Payload (seed + 1) }
  , right: { score: seed + 7, label: "keep", active: false }
  , enabled: true
  }

bump :: Int -> Packet Int -> Packet Int
bump amount packet = packet
  { left = packet.left { column = packet.left.column + amount }
  , right = packet.right { score = packet.right.score + amount }
  , enabled = not packet.enabled
  }

score :: Packet Int -> Int
score packet = packet.left.column + packet.left.line + packet.right.score
  + (if packet.left.active then 1 else 0)
  + (if packet.right.active then 2 else 0)
  + (if packet.enabled then 4 else 0)
  + case packet.payload of
      Vacant -> 0
      Payload value -> value
      Tagged _ value -> value

rareIdentity :: Packet Int -> Packet Int
rareIdentity packet = packet

-- The dynamic callback runs once. The repeated work then consists entirely
-- of direct calls; a global layout decision must account for this mixed case.
run :: (Packet Int -> Packet Int) -> Int -> Int -> Int
run callback count seed = go count (callback (makePacket seed)) 0
  where
  go n packet total
    | n <= 0 = total
    | otherwise =
        let next = bump 1 packet
        in go (n - 1) next (total + score next)

main :: Effect Unit
main = do
  seedRef <- Ref.new 5
  seed <- Ref.read seedRef
  callbackRef <- Ref.new rareIdentity
  callback <- Ref.read callbackRef
  assertEqual { expected: 109, actual: run callback 3 seed }
  assertEqual { expected: 0, actual: run callback 0 seed }
  log "Done"
