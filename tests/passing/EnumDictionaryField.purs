-- @inline Main.priority never
-- @inline Main.isSafe never
module Main where

-- @dependencies: assert prelude effect console

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Test.Assert (assertEqual)

data Priority = Fast | Safe
derive instance Eq Priority

class Protected a where
  priority :: Priority

instance protectedInt :: Protected Int where
  priority = Safe

instance protectedString :: Protected String where
  priority = Fast

isSafe :: forall @a. Protected a => Boolean
isSafe = priority @a == Safe

main :: Effect Unit
main = do
  assertEqual { expected: true, actual: isSafe @Int }
  assertEqual { expected: false, actual: isSafe @String }
  log "Done"
