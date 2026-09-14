module Gopurs.FfiTypes
  ( TypeNode(..)
  , FfiDecl
  ) where

import Prelude

import Data.Argonaut.Core (toObject)
import Data.Argonaut.Decode (class DecodeJson, (.:), (.:?))
import Data.Argonaut.Decode.Error (JsonDecodeError(TypeMismatch))
import Data.Either (note)
import Data.Maybe (Maybe, fromMaybe)

data TypeNode
  = TNamed String
  | TFunc (Array TypeNode) (Maybe TypeNode)
  | TArray TypeNode
  | TMap TypeNode TypeNode
  | TUnknown String

derive instance Eq TypeNode
derive instance Ord TypeNode

instance Show TypeNode where
  show (TNamed n) = "(TNamed " <> show n <> ")"
  show (TFunc args ret) = "(TFunc " <> show args <> " " <> show ret <> ")"
  show (TArray elem) = "(TArray " <> show elem <> ")"
  show (TMap k v) = "(TMap " <> show k <> " " <> show v <> ")"
  show (TUnknown s) = "(TUnknown " <> show s <> ")"

instance DecodeJson TypeNode where
  decodeJson json = do
    obj <- note (TypeMismatch "Object") (toObject json)
    typStr <- obj .: "type"
    case typStr of
      "Named" -> TNamed <$> obj .: "name"
      "Func" -> do
        args <- obj .:? "args" <#> fromMaybe []
        ret <- obj .:? "ret"
        pure (TFunc args ret)
      "Array" -> TArray <$> obj .: "elem"
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
  }
