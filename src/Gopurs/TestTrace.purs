module Gopurs.TestTrace where

import Prelude
import Data.Maybe
import Data.String as String
import Data.String.Pattern (Pattern(..))

fullName :: Maybe String -> String -> String
fullName mbMod name = fromMaybe "" (map (\m -> String.joinWith "_" (String.split (Pattern ".") m) <> "_") mbMod) <> name

main = fullName (Just "Test.TCO") "deepTailRec"
