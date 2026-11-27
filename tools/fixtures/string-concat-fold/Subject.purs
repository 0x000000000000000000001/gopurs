-- @inline Subject.consume never
-- @inline Subject.unknown never
-- @inline Subject.unknownFold never
module Subject (consume, unknownFold) where

import Prelude
import Data.Array as Array
import Data.Enum (fromEnum)
import Data.String.CodePoints as CP

consume :: String -> Int
consume input = Array.foldl (\total cp -> total * 33 + fromEnum cp) 17 (CP.toCodePointArray input)

unknown :: Int -> CP.CodePoint -> Int
unknown total cp = total + fromEnum cp

unknownFold :: String -> Int
unknownFold input = Array.foldl unknown 0 (CP.toCodePointArray input)
