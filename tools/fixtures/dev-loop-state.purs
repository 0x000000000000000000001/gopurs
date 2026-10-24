module Main where

-- @dependencies: prelude effect console lists tailrec foldable-traversable

import Prelude
import Control.Monad.Rec.Class (Step(..), tailRec)
import Data.Foldable (sum)
import Data.List (List(..))
import Data.List as List
import Effect (Effect)
import Effect.Console (logShow)

buildList :: Int -> List Int
buildList n = tailRec loop { index: 0, results: Nil }
  where
  loop :: { index :: Int, results :: List Int } -> Step { index :: Int, results :: List Int } (List Int)
  loop state
    | state.index == n = Done (List.reverse state.results)
    | otherwise = Loop { index: state.index + 1, results: Cons 1 state.results }

main :: Effect Unit
main = do
  let results = buildList 100000
  logShow (List.length results)
  logShow (sum results)
