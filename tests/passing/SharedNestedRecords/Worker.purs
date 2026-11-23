-- @inline export makePacket never
-- @inline export advance never
-- @inline export invoke never
-- @inline export keep never
-- @inline export score never
-- @inline export readColumn never
module Worker where

import Prelude

data Payload a = Vacant | Payload a | Tagged String a

type Position a = { column :: Int, line :: Int, active :: Boolean, note :: Payload a }
type Detail = { score :: Int, label :: String, active :: Boolean }
type Packet a =
  { payload :: Payload a
  , left :: Position a
  , right :: Detail
  , enabled :: Boolean
  }

makePacket :: Int -> Packet Int
makePacket seed =
  { payload: if seed < 0 then Vacant else if seed == 0 then Tagged "zero" 0 else Payload (seed * 2)
  , left: { column: seed, line: 3, active: true, note: Payload (seed + 1) }
  , right: { score: seed + 7, label: "keep", active: false }
  , enabled: true
  }

advance :: Int -> Packet Int -> Packet Int
advance amount packet = packet
  { payload = case packet.payload of
      Vacant -> Vacant
      Payload value -> Payload (value + amount)
      Tagged label value -> Tagged label (value + amount)
  , left = packet.left { column = packet.left.column + amount }
  , right = packet.right { score = packet.right.score * 2 }
  , enabled = not packet.enabled
  }

invoke :: forall @a @b. (a -> b) -> a -> b
invoke callback value = callback value

keep :: forall a. a -> Unit -> a
keep value _ = value

score :: Packet Int -> Int
score packet = packet.left.column + packet.left.line + packet.right.score
  + (if packet.left.active then 1 else 0)
  + (if packet.right.active then 2 else 0)
  + (if packet.enabled then 4 else 0)
  + case packet.payload of
      Vacant -> 0
      Payload value -> value
      Tagged _ value -> value

readColumn :: forall a r. { left :: Position a | r } -> Int
readColumn packet = packet.left.column
