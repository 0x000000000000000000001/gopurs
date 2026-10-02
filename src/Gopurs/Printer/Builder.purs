module Gopurs.Printer.Builder
  ( Out
  , emit
  , emitMany
  , withOut
  ) where

import Prelude

import Data.Array as Array
import Effect (Effect)
import Effect.Unsafe (unsafePerformEffect)

foreign import data Builder :: Type
foreign import newBuilderImpl :: Effect Builder
foreign import pushImpl :: Builder -> String -> Builder
foreign import toStringImpl :: Builder -> String

-- One private mutable handle per rendering. Thread the returned Out in order;
-- earlier handles alias the same buffer and must not be reused as snapshots.
-- The FFI operations stay direct pure calls, without boxing each fragment.
newtype Out = Out Builder

emit :: Out -> String -> Out
emit (Out builder) piece = Out (pushImpl builder piece)

emitMany :: Out -> Array String -> Out
emitMany = Array.foldl emit

finish :: Out -> String
finish (Out builder) = toStringImpl builder

-- The handle is confined to this synchronous callback. Independent renderings
-- own independent buffers; only the completed immutable string escapes.
withOut :: (Out -> Out) -> String
withOut write = unsafePerformEffect do
  builder <- newBuilderImpl
  pure (finish (write (Out builder)))
