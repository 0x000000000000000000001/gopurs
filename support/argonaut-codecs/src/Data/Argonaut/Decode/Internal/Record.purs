module Data.Argonaut.Decode.Internal.Record
  ( recordNilImpl
  , recordConsImpl
  , rightValue
  , RecordErrorSupport
  , typedJson
  , typedInt
  , typedNumber
  , typedString
  , typedBoolean
  , typedMaybe
  , typedArray
  , typedRecord
  , typedObject
  , borrowObject
  , schemaDecoderABI1
  , typedFieldId
  , typedFieldMaybe
  , fieldStep
  , decodeFieldFast
  , getFieldImpl
  , getFieldOptionalImpl
  , getFieldOptionalNullableImpl
  ) where

import Data.Argonaut.Core (Json)
import Data.Argonaut.Decode.Error (JsonDecodeError)
import Data.Either (Either(..))
import Foreign.Object (Object)
import Data.Maybe (Maybe)
import Partial.Unsafe (unsafeCrashWith)
import Prelude (Unit)
import Prim.RowList as RL

-- The public GDecodeJson dictionary remains unchanged. Native backends can
-- attach an immutable construction plan to its ordinary method value.
--
-- Error constructors travel in one record so the native plan can build the
-- exact JsonDecodeError values the generic path builds, including when it
-- decodes the parser's DOM directly. Left stays separate because it is
-- polymorphic in the error type.
type RecordErrorSupport =
  { atKey :: String -> JsonDecodeError -> JsonDecodeError
  , missingValue :: JsonDecodeError
  , typeMismatch :: String -> JsonDecodeError
  , named :: String -> JsonDecodeError -> JsonDecodeError
  , atIndex :: Int -> JsonDecodeError -> JsonDecodeError
  , leftOf :: Either JsonDecodeError Json -> JsonDecodeError
  , nothing :: Maybe Json
  , just :: Json -> Maybe Json
  , isRight :: Either JsonDecodeError Json -> Boolean
  , rightValue :: Either JsonDecodeError Json -> Json
  , left :: JsonDecodeError -> Either JsonDecodeError Json
  , right :: Json -> Either JsonDecodeError Json
  }

foreign import recordNilImpl
  :: forall proxy
   . (forall a. JsonDecodeError -> Either JsonDecodeError a)
  -> RecordErrorSupport
  -> (Record () -> Either JsonDecodeError (Record ()))
  -> Object Json
  -> proxy RL.Nil
  -> Either JsonDecodeError (Record ())

-- Keep ADT inspection in PureScript, and use this extractor only after isRight
-- has accepted the same result. The FFI never assumes a constructor layout.
rightValue :: forall a. Either JsonDecodeError a -> a
rightValue (Right value) = value
rightValue (Left _) = unsafeCrashWith "Record decoder expected a successful field"

-- A successful step returns its existing Either envelope unchanged. An unknown
-- tail uses fallback, retaining the behavior of arbitrary GDecodeJson methods.
foreign import recordConsImpl
  :: forall value rowTail tail row list proxyTail proxy
   . (Unit -> String)
  -> (String -> Object Json -> Either JsonDecodeError value)
  -> (Object Json -> proxyTail tail -> Either JsonDecodeError (Record rowTail))
  -> (Object Json -> proxy list -> Either JsonDecodeError (Record row))
  -> (forall a. Either JsonDecodeError a -> Boolean)
  -> (forall a. Either JsonDecodeError a -> a)
  -> (Record row -> Either JsonDecodeError (Record row))
  -> (forall a. JsonDecodeError -> Either JsonDecodeError a)
  -> RecordErrorSupport
  -> Maybe value
  -> (value -> Maybe value)
  -> Object Json
  -> proxy list
  -> Either JsonDecodeError (Record row)

-- ---------------------------------------------------------------------------
-- Schema tags
--
-- Primitive and container decoders carry an immutable tag so record plans and
-- field accessors can decode the parser's DOM directly. Ordinary application
-- keeps the passed decoder; only consumers of a recognized tag bypass it.

foreign import typedInt
  :: RecordErrorSupport
  -> (Json -> Either JsonDecodeError Int)
  -> (Json -> Either JsonDecodeError Int)

foreign import typedJson
  :: RecordErrorSupport
  -> (Json -> Either JsonDecodeError Json)
  -> (Json -> Either JsonDecodeError Json)

foreign import typedNumber
  :: RecordErrorSupport
  -> (Json -> Either JsonDecodeError Number)
  -> (Json -> Either JsonDecodeError Number)

foreign import typedString
  :: RecordErrorSupport
  -> (Json -> Either JsonDecodeError String)
  -> (Json -> Either JsonDecodeError String)

foreign import typedBoolean
  :: RecordErrorSupport
  -> (Json -> Either JsonDecodeError Boolean)
  -> (Json -> Either JsonDecodeError Boolean)

foreign import typedMaybe
  :: forall a
   . RecordErrorSupport
  -> (Json -> Either JsonDecodeError a)
  -> (Json -> Either JsonDecodeError (Maybe a))
  -> (Json -> Either JsonDecodeError (Maybe a))

foreign import typedArray
  :: forall a
   . RecordErrorSupport
  -> (Json -> Either JsonDecodeError a)
  -> (Json -> Either JsonDecodeError (Array a))
  -> (Json -> Either JsonDecodeError (Array a))

-- The Json wrapper must carry the row method's plan too. Arbitrary GDecodeJson
-- methods have no recognized plan and keep the ordinary wrapper.
foreign import typedRecord
  :: forall row list proxy
   . RecordErrorSupport
  -> (Object Json -> proxy list -> Either JsonDecodeError (Record row))
  -> (Json -> Either JsonDecodeError (Record row))
  -> (Json -> Either JsonDecodeError (Record row))

-- Only the standard Json identity decoder enables the native object shortcut.
-- The result keeps its own object storage; custom element decoders traverse.
foreign import typedObject
  :: forall a
   . (Json -> Either JsonDecodeError a)
  -> (Json -> Either JsonDecodeError (Object a))
  -> (Object a -> Either JsonDecodeError (Object a))
  -> (Json -> Either JsonDecodeError (Object a))

-- Compiler-only call site: the result object must be consumed synchronously
-- by read-only field accessors, never returned, stored, captured or mutated.
-- The ordinary typedObject decoder always returns independent storage.
foreign import borrowObject
  :: (Json -> Either JsonDecodeError (Object Json))
  -> Json
  -> Either JsonDecodeError (Object Json)

-- Versioned compiler/library handshake for emitted native schema workers.
foreign import schemaDecoderABI1 :: Int

foreign import typedFieldId
  :: forall a
   . (Json -> Either JsonDecodeError a)
  -> (Maybe Json -> Maybe (Either JsonDecodeError a))
  -> (Maybe Json -> Maybe (Either JsonDecodeError a))

foreign import typedFieldMaybe
  :: forall a
   . (Json -> Either JsonDecodeError a)
  -> (Maybe Json -> Maybe (Either JsonDecodeError (Maybe a)))
  -> (Maybe Json -> Maybe (Either JsonDecodeError (Maybe a)))

foreign import fieldStep
  :: forall value
   . (Maybe Json -> Maybe (Either JsonDecodeError value))
  -> (String -> Object Json -> Either JsonDecodeError value)
  -> (String -> Object Json -> Either JsonDecodeError value)

-- Direct field decoding for tagged decoders. Untagged decoders are applied
-- unchanged, so this is the ordinary decoder call plus one metadata lookup.
foreign import decodeFieldFast
  :: forall a
   . (Json -> Either JsonDecodeError a)
  -> Json
  -> Either JsonDecodeError a

-- Fuse lookup and a tagged decoder into one result. Unknown decoders retain
-- the complete ordinary accessor, including its missing/null behavior.
foreign import getFieldImpl
  :: forall a
   . ((Json -> Either JsonDecodeError a) -> Object Json -> String -> Either JsonDecodeError a)
  -> (Json -> Either JsonDecodeError a)
  -> Object Json
  -> String
  -> Either JsonDecodeError a

foreign import getFieldOptionalImpl
  :: forall a
   . ((Json -> Either JsonDecodeError a) -> Object Json -> String -> Either JsonDecodeError (Maybe a))
  -> (Json -> Either JsonDecodeError a)
  -> Object Json
  -> String
  -> Either JsonDecodeError (Maybe a)

foreign import getFieldOptionalNullableImpl
  :: forall a
   . ((Json -> Either JsonDecodeError a) -> Object Json -> String -> Either JsonDecodeError (Maybe a))
  -> (Json -> Either JsonDecodeError a)
  -> Object Json
  -> String
  -> Either JsonDecodeError (Maybe a)
