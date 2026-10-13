-- @inline export value never
module Cache.Calls where

import Prelude
import Cache.Relay as Relay

value :: Int -> Int
value n = Relay.relay n
