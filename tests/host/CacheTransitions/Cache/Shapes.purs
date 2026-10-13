-- @inline export pack never
-- @inline export unpack never
-- @inline export packet never
-- @inline export scorePacket never
module Cache.Shapes where

import Prelude

data Payload = Empty | Payload Int

pack :: Int -> Payload
pack n = Payload n

unpack :: Payload -> Int
unpack Empty = 0
unpack (Payload n) = n

type Packet = { left :: Int, right :: Int }

packet :: Int -> Packet
packet n = { left: n, right: 3 }

scorePacket :: Packet -> Int
scorePacket p = p.left + p.right
