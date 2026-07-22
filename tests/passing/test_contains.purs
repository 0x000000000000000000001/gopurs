module Main where
import Prelude
import Effect (Effect)
import Effect.Console (log)
import Test.Assert (assert)
import Data.String (contains)
import Data.String.Pattern (Pattern(..))

main :: Effect Unit
main = do
  let err = "interface conversion: interface {} is nil, not map[string]gopurs_runtime.Value"
  log "Testing contains"
  assert $ contains (Pattern "interface conversion") err
  log "Success"
