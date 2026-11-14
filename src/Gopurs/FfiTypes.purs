module Gopurs.FfiTypes
  ( TypeNode(..)
  , FfiDecl
  , StructField
  ) where

import Prelude

import Data.Argonaut.Core (toObject)
import Data.Argonaut.Decode (class DecodeJson, (.:), (.:?))
import Data.Argonaut.Decode.Error (JsonDecodeError(TypeMismatch))
import Data.Either (note)
import Data.Maybe (Maybe, fromMaybe, isNothing)
import Data.Array as Array

data TypeNode
  = TNamed String
  | TFunc (Array TypeNode) (Maybe TypeNode)
  | TArray TypeNode
  | TMap TypeNode TypeNode
  | TIdentified String String (Maybe TypeNode)
  | TPointer TypeNode
  | TFixedArray TypeNode String
  | TVariadic TypeNode
  | TInterface String
  | TMultiFunc (Array TypeNode) (Array TypeNode)
  | TStruct String (Array StructField)
  | TUnknown String

type StructField = { name :: String, type :: TypeNode, embedded :: Boolean, exported :: Boolean, tag :: String }

derive instance Eq TypeNode
derive instance Ord TypeNode

instance Show TypeNode where
  show (TNamed n) = "(TNamed " <> show n <> ")"
  show (TFunc args ret) = "(TFunc " <> show args <> " " <> show ret <> ")"
  show (TArray elem) = "(TArray " <> show elem <> ")"
  show (TMap k v) = "(TMap " <> show k <> " " <> show v <> ")"
  show (TIdentified n i u) = "(TIdentified " <> show n <> " " <> show i <> " " <> show u <> ")"
  show (TPointer t) = "(TPointer " <> show t <> ")"
  show (TFixedArray t n) = "(TFixedArray " <> show t <> " " <> show n <> ")"
  show (TVariadic t) = "(TVariadic " <> show t <> ")"
  show (TInterface s) = "(TInterface " <> show s <> ")"
  show (TMultiFunc a r) = "(TMultiFunc " <> show a <> " " <> show r <> ")"
  show (TStruct source fields) = "(TStruct " <> show source <> " " <> show fields <> ")"
  show (TUnknown s) = "(TUnknown " <> show s <> ")"

instance DecodeJson TypeNode where
  decodeJson json = do
    obj <- note (TypeMismatch "Object") (toObject json)
    typStr <- obj .: "type"
    case typStr of
      "Named" -> do
        name <- obj .: "name"
        identity <- obj .:? "identity" <#> fromMaybe ""
        underlying <- obj .:? "underlying"
        pure (if identity == "" && isNothing underlying then TNamed name else TIdentified name identity underlying)
      "Func" -> do
        args <- obj .:? "args" <#> fromMaybe []
        ret <- obj .:? "ret"
        results <- obj .:? "results" <#> fromMaybe (Array.fromFoldable ret)
        pure (if Array.length results > 1 then TMultiFunc args results else TFunc args (Array.head results))
      "Array" -> TArray <$> obj .: "elem"
      "Pointer" -> TPointer <$> obj .: "elem"
      "FixedArray" -> TFixedArray <$> obj .: "elem" <*> obj .: "length"
      "Variadic" -> TVariadic <$> obj .: "elem"
      "Interface" -> TInterface <$> obj .: "name"
      "Struct" -> TStruct <$> obj .: "name" <*> (obj .:? "fields" <#> fromMaybe [])
      "Map" -> do
        k <- obj .: "key"
        v <- obj .: "val"
        pure (TMap k v)
      _ -> TUnknown <$> obj .: "name"

type FfiDecl =
  { name :: String
  , isVar :: Boolean
  , typeParams :: Array String
  , args :: Array TypeNode
  , ret :: Maybe TypeNode
  , results :: Array TypeNode
  , schemaVersion :: Int
  , line :: Int
  , column :: Int
  }
