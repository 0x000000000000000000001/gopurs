-- @inline Main.sequenceTerm never
-- @inline Main.weighted never
-- @inline Main.permuted never
-- @inline Main.recordValue never
-- @inline Main.observed never
module Main where

-- @dependencies: prelude effect console refs arrays foldable-traversable

import Prelude
import Data.Array as Array
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

-- Supply an actual noncanonical Int without frontend integer folding.
foreign import wideBase :: Effect Int

sequenceTerm :: Int -> Int
sequenceTerm n =
  if n == 0 then 0
  else if n == 1 then 1
  else sequenceTerm (n - 1) + sequenceTerm (n - 2)

weighted :: Int -> Int -> Int
weighted count seed =
  if count <= 0 then seed
  else (count + seed) + weighted (count - 1) (seed + 7)

permuted :: Int -> Int -> Int -> Int
permuted count first second =
  if count == 0 then first
  else second + permuted (count - 1) second (first + 1)

recordValue :: Ref.Ref (Array Int) -> Int -> Int -> Int
recordValue reference tag value = unsafePerformEffect do
  Ref.modify_ (\previous -> Array.snoc previous tag) reference
  pure value

observed :: Ref.Ref (Array Int) -> Int -> Int -> Int
observed events count seed =
  if count == 0 then recordValue events 300 seed
  else recordValue events (100 + count) 2147483647 +
    observed events (recordValue events (200 + count - 1) (count - 1)) seed

main :: Effect Unit
main = do
  wide <- wideBase
  arguments <- Ref.new [ 0, 1, 2, 3, 8, 12 ]
  values <- Ref.read arguments
  logShow (map sequenceTerm values)
  for_ [ 0, 7, 2147483647, -2147483648, wide ] \seed -> do
    reference <- Ref.new seed
    input <- Ref.read reference
    logShow (map (\count -> weighted count input) values)
    logShow (map (\count -> permuted count input (-17)) values)
    for_ [ 0, 1, 3 ] \count -> do
      events <- Ref.new []
      dynamic <- Ref.new count
      limit <- Ref.read dynamic
      logShow (observed events limit input)
      Ref.read events >>= logShow
