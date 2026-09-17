module Gopurs.Metrics (measure) where

import Prelude

import Data.Either (Either(..), either)
import Data.Int as Int
import Effect (Effect)
import Effect.Aff (Aff, attempt, throwError)
import Effect.Class (liftEffect)
import Effect.Console as Console

-- Milliseconds from a monotonic clock; only the clock is host-specific.
foreign import now :: Effect Number

-- Delay construction of the action until after the first clock read.
-- Nested phases are included in the outer total; they are not additive to it.
measure :: forall a. String -> (Unit -> Aff a) -> Aff a
measure label action = do
  started <- liftEffect now
  result <- attempt (pure unit >>= action)
  ended <- liftEffect now
  let status = case result of
        Left _ -> " (failed)"
        Right _ -> ""
  liftEffect $ Console.error $ "[gopurs] " <> label <> ": "
    <> show (Int.round (ended - started)) <> " ms" <> status
  either throwError pure result
