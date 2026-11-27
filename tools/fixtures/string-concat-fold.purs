-- @dependencies: prelude effect console refs arrays foldable-traversable maybe strings enums
-- @inline Main.combined never
-- @inline Main.shared never
-- @inline Main.unknown never
module Main where

import Prelude
import Data.Array as Array
import Data.Enum (toEnum)
import Data.Foldable (for_)
import Data.Maybe (fromMaybe)
import Data.String.CodePoints as CP
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref
import Subject as Subject

combined :: String -> Array String -> Int
combined initial parts = Subject.consume (Array.foldl (<>) initial parts)

shared :: String -> Array String -> Int
shared initial parts = let joined = Array.foldl (<>) initial parts
  in Subject.consume joined + CP.length joined

unknown :: String -> Array String -> Int
unknown initial parts = Subject.unknownFold (Array.foldl (<>) initial parts)

point :: Int -> String
point value = CP.singleton (fromMaybe bottom (toEnum value))

main :: Effect Unit
main = do
  casesRef <- Ref.new [ [], [ 97, 98 ], [ 55296, 56320 ], [ 56320, 55296 ], [ 55296, 97, 56320 ], [ 233, 128512 ], [ 97, 55296, 56320, 98 ] ]
  cases <- Ref.read casesRef
  initialRef <- Ref.new [ "", "prefix", point 55296 ]
  initials <- Ref.read initialRef
  functionRef <- Ref.new combined
  indirect <- Ref.read functionRef
  for_ initials \initial -> for_ cases \values -> do
    partsRef <- Ref.new (Array.concatMap (\value -> [ "", point value ]) values)
    parts <- Ref.read partsRef
    logShow [ combined initial parts, indirect initial parts, shared initial parts, unknown initial parts ]
