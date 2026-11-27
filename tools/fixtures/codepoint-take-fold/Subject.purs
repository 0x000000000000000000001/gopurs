-- @inline Subject.consume never
module Subject (consume) where

import Prelude
import Data.Array as Array
import Data.Enum (fromEnum)
import Data.String.CodePoints as CP

consume :: String -> Int
consume input = Array.foldl (\total cp -> total * 3 + fromEnum cp) 7 (CP.toCodePointArray input)
