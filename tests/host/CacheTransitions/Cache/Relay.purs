-- @inline export relay never
module Cache.Relay where

import Cache.Poly as Poly

relay :: forall a. a -> a
relay value = Poly.echo value
