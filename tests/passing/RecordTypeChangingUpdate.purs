-- @inline Main.convert never
module Main where

-- @dependencies: assert prelude effect console refs maybe

import Prelude
import Data.Maybe (Maybe(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

convert :: { after :: Maybe String, limit :: Int } -> { after :: Maybe { identifier :: String }, limit :: Int }
convert options = options { after = options.after <#> \identifier -> { identifier } }

main :: Effect Unit
main = do
  input <- Ref.new { after: Just "book-42", limit: 10 }
  options <- Ref.read input
  assertEqual
    { expected: { after: Just { identifier: "book-42" }, limit: 10 }
    , actual: convert options
    }
  log "Done"
