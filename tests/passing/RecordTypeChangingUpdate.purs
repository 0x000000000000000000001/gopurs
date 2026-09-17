-- @inline Main.convert never
module Main where

-- @dependencies: assert prelude effect console refs maybe

import Prelude
import Data.Maybe (Maybe(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

-- Preserve both a Go keyword and a builtin identifier while changing another
-- field's representation. Their native field accesses must use escaped names.
convert
  :: { after :: Maybe String, limit :: Int, type :: String, close :: Boolean }
  -> { after :: Maybe { identifier :: String }, limit :: Int, type :: String, close :: Boolean }
convert options = options { after = options.after <#> \identifier -> { identifier } }

main :: Effect Unit
main = do
  input <- Ref.new { after: Just "book-42", limit: 10, type: "book", close: true }
  options <- Ref.read input
  assertEqual
    { expected: { after: Just { identifier: "book-42" }, limit: 10, type: "book", close: true }
    , actual: convert options
    }
  log "Done"
