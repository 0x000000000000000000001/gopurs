module Gopurs.Host.Metrics (measure, accumulate, setMemProfileRate) where

import Prelude

import Data.Either (Either(..), either)
import Data.Int as Int
import Effect (Effect)
import Effect.Aff (Aff, attempt, throwError)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Ref as Ref

-- Milliseconds from a monotonic clock; only the clock is host-specific.
foreign import now :: Effect Number

-- | Set Go memory profiling's sampling rate (`runtime.MemProfileRate`).
-- | A compiler allocates hundreds of GB: the default rate (512 KiB)
-- | wastes ~15% of CPU time when no profile is requested.
foreign import setMemProfileRate :: Int -> Effect Unit

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

-- Sum batch wall times without logging from concurrent emission workers.
-- This overlaps the producer clock in pipelined mode; it is not an additive
-- component of backend total. Construction remains inside the timed action.
accumulate :: forall a. Ref.Ref Number -> (Unit -> Aff a) -> Aff a
accumulate counter action = do
  started <- liftEffect now
  result <- attempt (pure unit >>= action)
  ended <- liftEffect now
  liftEffect $ Ref.modify_ (\total -> total + (ended - started)) counter
  either throwError pure result
