module Data.Argonaut.Decode.Parser where

import Prelude

import Data.Argonaut.Core (Json)
import Data.Argonaut.Decode.Class (class DecodeJson, decodeJson)
import Data.Argonaut.Decode.Error (JsonDecodeError(..))
import Data.Argonaut.Decode.Internal.Record (schemaDecoderABI1)
import Data.Argonaut.Parser (jsonParser)
import Data.Bifunctor (lmap)
import Data.Either (Either)

-- | Attempt to parse a string as `Json`, failing with a typed error if the
-- | JSON string is malformed.
parseJson :: String -> Either JsonDecodeError Json
parseJson = lmap (\_ -> TypeMismatch "JSON") <<< jsonParser

-- | Parse and decode a complete JSON document. Known native decoders can read
-- | validated text directly; arbitrary decoders keep the ordinary composition.
decodeJsonString :: forall a. DecodeJson a => String -> Either JsonDecodeError a
decodeJsonString = decodeJsonStringWith decodeJson

-- | The explicit-decoder version preserves custom decoding and error order.
decodeJsonStringWith :: forall a. (Json -> Either JsonDecodeError a) -> String -> Either JsonDecodeError a
decodeJsonStringWith decoder = decodeJsonStringImpl schemaDecoderABI1 (\text -> parseJson text >>= decoder) decoder

foreign import textDecoderABI1 :: Int

foreign import decodeJsonStringImpl
  :: forall a
   . Int
  -> (String -> Either JsonDecodeError a)
  -> (Json -> Either JsonDecodeError a)
  -> String
  -> Either JsonDecodeError a
