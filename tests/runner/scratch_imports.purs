module ScratchImports where

import Prelude
import Data.String (Pattern(..))
import Data.String as String
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Effect.Console (log)

findPkgs :: String -> Array String
findPkgs str =
  let
    parts = String.split (Pattern "pkg_") str
  in
    Array.nub $ Array.mapMaybe
      ( \part ->
          let
            subParts = String.split (Pattern ".") part
          in
            Array.head subParts
      )
      (fromMaybe [] (Array.tail parts))

main = do
  log $ show $ findPkgs "gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), pkg_Partial.Get__crashWith())"
