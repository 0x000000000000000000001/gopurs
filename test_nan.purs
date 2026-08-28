module TestNaN where
import Prelude
import Effect.Console (log)

main = do
  log (show (0.0 / 0.0))
  log (show (1.0 / 0.0))
  log (show (-1.0 / 0.0))
