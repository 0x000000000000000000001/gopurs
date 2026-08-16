module TestArrayMatch where
import Prelude
import Effect (Effect)
import Effect.Console (log)

testMatch :: Array String -> String
testMatch arr = case arr of
  [_] -> "Matched one element"
  _ -> "Matched something else"

main :: Effect Unit
main = do
  log (testMatch ["bool"])
  log (testMatch ["bool", "bool"])
