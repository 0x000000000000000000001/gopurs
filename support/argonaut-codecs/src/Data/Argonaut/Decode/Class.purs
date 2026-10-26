module Data.Argonaut.Decode.Class where

import Data.Argonaut.Decode.Decoders
import Data.Argonaut.Decode.Decoders as Decoders

import Data.Argonaut.Core (Json, toObject)
import Data.Argonaut.Decode.Error (JsonDecodeError(..))
import Data.Argonaut.Decode.Internal.Record
  ( recordConsImpl
  , recordNilImpl
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
  , typedFieldId
  , typedFieldMaybe
  , fieldStep
  )
import Data.Array.NonEmpty (NonEmptyArray)
import Data.Either (Either(..), isRight)
import Data.Identity (Identity)
import Data.List (List)
import Data.List.NonEmpty (NonEmptyList)
import Data.String.NonEmpty (NonEmptyString)
import Data.Map as M
import Data.Maybe (Maybe(..))
import Data.NonEmpty (NonEmpty)
import Data.Set as S
import Data.String (CodePoint)
import Data.Symbol (class IsSymbol, reflectSymbol)
import Data.Tuple (Tuple)
import Foreign.Object as FO
import Partial.Unsafe (unsafeCrashWith)
import Prelude (class Ord, Unit, Void, bind, ($), (<$>))
import Prim.Row as Row
import Prim.RowList as RL
import Record as Record
import Type.Proxy (Proxy(..))

-- The public DecodeJson dictionaries stay unchanged. Standard instances carry
-- a schema tag so record plans can decode the parser's DOM directly. Ordinary
-- calls retain the decoder; unrecognized dictionaries keep their own method.
class DecodeJson a where
  decodeJson :: Json -> Either JsonDecodeError a

instance decodeIdentity :: DecodeJson a => DecodeJson (Identity a) where
  decodeJson = decodeIdentity decodeJson

instance decodeJsonMaybe :: DecodeJson a => DecodeJson (Maybe a) where
  decodeJson = typedMaybe recordErrorSupport
    (decodeJson :: Json -> Either JsonDecodeError a)
    (Decoders.decodeMaybe (decodeJson :: Json -> Either JsonDecodeError a))

instance decodeJsonTuple :: (DecodeJson a, DecodeJson b) => DecodeJson (Tuple a b) where
  decodeJson = decodeTuple decodeJson decodeJson

instance decodeJsonEither :: (DecodeJson a, DecodeJson b) => DecodeJson (Either a b) where
  decodeJson = decodeEither decodeJson decodeJson

instance decodeJsonNull :: DecodeJson Unit where
  decodeJson = decodeNull

instance decodeJsonBoolean :: DecodeJson Boolean where
  decodeJson = typedBoolean recordErrorSupport Decoders.decodeBoolean

instance decodeJsonNumber :: DecodeJson Number where
  decodeJson = typedNumber recordErrorSupport Decoders.decodeNumber

instance decodeJsonInt :: DecodeJson Int where
  decodeJson = typedInt recordErrorSupport Decoders.decodeInt

instance decodeJsonString :: DecodeJson String where
  decodeJson = typedString recordErrorSupport Decoders.decodeString

instance decodeJsonNonEmptyString :: DecodeJson NonEmptyString where
  decodeJson = decodeNonEmptyString

instance decodeJsonJson :: DecodeJson Json where
  decodeJson = typedJson recordErrorSupport Right

instance decodeJsonNonEmpty_Array :: (DecodeJson a) => DecodeJson (NonEmpty Array a) where
  decodeJson = decodeNonEmpty_Array decodeJson

instance decodeJsonNonEmptyArray :: (DecodeJson a) => DecodeJson (NonEmptyArray a) where
  decodeJson = decodeNonEmptyArray decodeJson

instance decodeJsonNonEmpty_List :: (DecodeJson a) => DecodeJson (NonEmpty List a) where
  decodeJson = decodeNonEmpty_List decodeJson

instance decodeJsonNonEmptyList :: (DecodeJson a) => DecodeJson (NonEmptyList a) where
  decodeJson = decodeNonEmptyList decodeJson

instance decodeJsonCodePoint :: DecodeJson CodePoint where
  decodeJson = decodeCodePoint

instance decodeForeignObject :: DecodeJson a => DecodeJson (FO.Object a) where
  decodeJson = typedObject
    (decodeJson :: Json -> Either JsonDecodeError a)
    (Decoders.decodeForeignObject (decodeJson :: Json -> Either JsonDecodeError a))
    Right

instance decodeArray :: DecodeJson a => DecodeJson (Array a) where
  decodeJson = typedArray recordErrorSupport
    (decodeJson :: Json -> Either JsonDecodeError a)
    (Decoders.decodeArray (decodeJson :: Json -> Either JsonDecodeError a))

instance decodeList :: DecodeJson a => DecodeJson (List a) where
  decodeJson = decodeList decodeJson

instance decodeSet :: (Ord a, DecodeJson a) => DecodeJson (S.Set a) where
  decodeJson = decodeSet decodeJson

instance decodeMap :: (Ord a, DecodeJson a, DecodeJson b) => DecodeJson (M.Map a b) where
  decodeJson = decodeMap decodeJson decodeJson

instance decodeVoid :: DecodeJson Void where
  decodeJson = decodeVoid

instance decodeRecord ::
  ( GDecodeJson row list
  , RL.RowToList row list
  ) =>
  DecodeJson (Record row) where
  decodeJson = typedRecord recordErrorSupport
    (gDecodeJson :: FO.Object Json -> Proxy list -> Either JsonDecodeError (Record row))
    \json -> case toObject json of
      Just object -> gDecodeJson object (Proxy :: Proxy list)
      Nothing -> Left $ TypeMismatch "Object"

class GDecodeJson (row :: Row Type) (list :: RL.RowList Type) | list -> row where
  gDecodeJson :: forall proxy. FO.Object Json -> proxy list -> Either JsonDecodeError (Record row)

-- Constructed once; plans and tags read its fields when a decoder is built.
-- Values are never inspected while documents are decoded.
recordErrorSupport :: RecordErrorSupport
recordErrorSupport =
  { atKey: AtKey
  , missingValue: MissingValue
  , typeMismatch: TypeMismatch
  , named: Named
  , atIndex: AtIndex
  , leftOf: \either -> case either of
      Left err -> err
      Right _ -> unsafeCrashWith "record plan expected a failing Either"
  , nothing: Nothing
  , just: Just
  , isRight: isRight
  , rightValue: \either -> case either of
      Right value -> value
      Left _ -> unsafeCrashWith "record plan expected a successful field"
  , left: Left
  , right: Right
  }

instance gDecodeJsonNil :: GDecodeJson () RL.Nil where
  gDecodeJson = recordNilImpl Left recordErrorSupport Right

instance gDecodeJsonCons ::
  ( DecodeJsonField value
  , GDecodeJson rowTail tail
  , IsSymbol field
  , Row.Cons field value rowTail row
  , Row.Lacks field rowTail
  ) =>
  GDecodeJson row (RL.Cons field value tail) where
  gDecodeJson = recordConsImpl reflectName (fieldStep decodeJsonField step)
    (gDecodeJson :: FO.Object Json -> Proxy tail -> Either JsonDecodeError (Record rowTail))
    fallback isRight rightValue Right
    Left recordErrorSupport Nothing Just
    where
    _field = Proxy :: Proxy field

    reflectName _ = reflectSymbol _field

    step
      :: String
      -> FO.Object Json
      -> Either JsonDecodeError value
    step fieldName object =
      case (decodeJsonField (FO.lookup fieldName object) :: Maybe (Either JsonDecodeError value)) of
        Just fieldValue -> case fieldValue of
          Left err -> Left (AtKey fieldName err)
          Right _ -> fieldValue
        Nothing -> Left (AtKey fieldName MissingValue)

    fallback object _ = do
      let
        fieldName = reflectSymbol _field
        fieldValue = FO.lookup fieldName object

      case decodeJsonField fieldValue of
        Just fieldVal -> do
          val <- case fieldVal of
            Left err -> Left (AtKey fieldName err)
            Right value -> Right value
          rest <- gDecodeJson object (Proxy :: Proxy tail)
          Right $ Record.insert _field val rest

        Nothing ->
          Left $ AtKey fieldName MissingValue

class DecodeJsonField a where
  decodeJsonField :: Maybe Json -> Maybe (Either JsonDecodeError a)

instance decodeFieldMaybe ::
  DecodeJson a =>
  DecodeJsonField (Maybe a) where
  decodeJsonField = typedFieldMaybe (decodeJson :: Json -> Either JsonDecodeError a)
    \j -> case j of
      Nothing -> Just $ Right Nothing
      Just value -> Just $ decodeJson value

else instance decodeFieldId ::
  DecodeJson a =>
  DecodeJsonField a where
  decodeJsonField = typedFieldId (decodeJson :: Json -> Either JsonDecodeError a)
    (\j -> decodeJson <$> j)
