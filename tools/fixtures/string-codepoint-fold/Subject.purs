-- @inline Subject.checksum never
module Subject (checksum) where

import Prelude
import Data.Array as Array
import Data.Enum (fromEnum)
import Data.String.CodePoints as CP

checksum :: String -> Int
checksum = Array.foldl (\total cp -> total + fromEnum cp) 0 <<< CP.toCodePointArray
