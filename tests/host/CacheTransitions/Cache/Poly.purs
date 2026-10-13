-- @inline export echo never
module Cache.Poly where

echo :: forall a. a -> a
echo value = value
