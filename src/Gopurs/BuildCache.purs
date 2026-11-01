-- | Versioned storage protocol; see docs/build-cache.md. The driver supplies
-- | complete input recipes and private staging outputs through Driver.Cache.
module Gopurs.BuildCache (exchange, sessionExchange, sessionExchangeOnMiss) where

import Data.Unit (Unit)
import Effect (Effect)

-- | A JSON request always returns a JSON response, including protocol errors.
-- | Host/process I/O failures are Effect exceptions. Domain values cross this
-- | boundary through explicit codecs, never the host's ADT representation.
foreign import exchange :: String -> Effect String

foreign import sessionExchange :: String -> Effect String

-- | Optional native preload notification, scoped to this request. JS and Rust
-- | use the ordinary sequential exchange and do not invoke the callback.
foreign import sessionExchangeOnMiss :: String -> (String -> Effect Unit) -> Effect String
