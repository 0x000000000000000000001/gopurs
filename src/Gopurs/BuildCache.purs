-- | Versioned storage protocol; see docs/build-cache.md. The driver will supply
-- | complete input recipes and private staging outputs when cache reuse is wired.
module Gopurs.BuildCache (exchange) where

import Effect (Effect)

-- | A JSON request always returns a JSON response, including protocol errors.
-- | Host/process I/O failures are Effect exceptions. Domain values cross this
-- | boundary through explicit codecs, never the host's ADT representation.
foreign import exchange :: String -> Effect String
