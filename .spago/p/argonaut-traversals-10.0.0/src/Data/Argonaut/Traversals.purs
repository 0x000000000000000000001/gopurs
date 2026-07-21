module Data.Argonaut.Traversals where

import Prelude ((<<<), identity)
import Data.Argonaut.Core
import Data.Lens (Traversal', filtered)

_JsonNull :: Traversal' Json Json
_JsonNull = identity <<< filtered isNull

_JsonBoolean :: Traversal' Json Json
_JsonBoolean = identity <<< filtered isBoolean

_JsonNumber :: Traversal' Json Json
_JsonNumber = identity <<< filtered isNumber

_JsonString :: Traversal' Json Json
_JsonString = identity <<< filtered isString

_JsonArray :: Traversal' Json Json
_JsonArray = identity <<< filtered isArray

_JsonObject :: Traversal' Json Json
_JsonObject = identity <<< filtered isObject
