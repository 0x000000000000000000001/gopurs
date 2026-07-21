module Data.Argonaut.Prisms where

import Data.Argonaut.Core
import Data.Lens (Prism', prism')
import Foreign.Object as FO
import Prelude (Unit, const)

_Null :: Prism' Json Unit
_Null = prism' (const jsonNull) toNull

_Boolean :: Prism' Json Boolean
_Boolean = prism' fromBoolean toBoolean

_Number :: Prism' Json Number
_Number = prism' fromNumber toNumber

_String :: Prism' Json String
_String = prism' fromString toString

_Array :: Prism' Json (Array Json)
_Array = prism' fromArray toArray

_Object :: Prism' Json (FO.Object Json)
_Object = prism' fromObject toObject
