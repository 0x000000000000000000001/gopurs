-- @inline export measure never
module Cache.Dictionary where

import Prelude

class Meter a where
  meter :: a -> Int

instance meterInt :: Meter Int where
  meter n = n + 1

measure :: forall a. Meter a => a -> Int
measure value = meter value
