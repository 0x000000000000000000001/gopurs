-- @inline export makePacket never
-- @inline export bump never
-- @inline export score never
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

-- Every call involving Packet is direct and monomorphic. In particular,
-- this is not a pointfree alias that would itself be a function value.
run :: Int -> Int -> Int
run count seed = go count (makePacket seed) 0
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
  assertEqual { expected: 35, actual: score (makePacket seed) }
  assertEqual { expected: 15, actual: score (makePacket 0) }
  assertEqual { expected: 109, actual: run 3 seed }
  assertEqual { expected: 0, actual: run 0 seed }
  log "Done"
