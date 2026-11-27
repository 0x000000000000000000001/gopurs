module Main where

-- @dependencies: prelude effect console refs exceptions either foldable-traversable
-- @inline Main.choose never
-- @inline Main.nested never
-- @inline Main.factory never
-- @inline Main.condition never
-- @inline Main.makeSaved never
-- @inline Main.returned never

import Prelude
import Data.Either (isLeft)
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (log, logShow)
import Effect.Exception (throw, try)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

tick :: forall a. Ref.Ref String -> String -> a -> Effect a
tick trace label value = do
  Ref.modify_ (_ <> label) trace
  pure value

condition :: Ref.Ref String -> String -> Boolean -> Boolean
condition trace label value = unsafePerformEffect (tick trace label value)

factory :: Ref.Ref String -> String -> Int -> Effect Int
factory trace label value = unsafePerformEffect do
  _ <- tick trace ("F" <> label) unit
  pure (tick trace ("E" <> label) value)

choose :: Ref.Ref String -> Ref.Ref Boolean -> Effect Int
choose trace flag = do
  selected <- Ref.read flag
  if condition trace "C" selected then pure 7 else factory trace "U" 19

nested :: Ref.Ref String -> Ref.Ref Boolean -> Effect Int
nested trace flag = do
  selected <- Ref.read flag
  if condition trace "A" selected then pure 11
  else if condition trace "B" (not selected) then tick trace "K" 22
  else factory trace "N" 33

makeSaved :: Ref.Ref String -> Ref.Ref Boolean -> Effect (Effect Int)
makeSaved trace flag = do
  selected <- Ref.read flag
  pending <- if condition trace "S" selected then do
    _ <- tick trace "FP" unit
    pure (tick trace "EP" 41)
    else do
      _ <- tick trace "FQ" unit
      pure (tick trace "EQ" 43)
  _ <- tick trace "D" unit
  pure pending

returned :: Ref.Ref String -> Ref.Ref Boolean -> Effect (Effect Int)
returned trace flag = do
  selected <- Ref.read flag
  if condition trace "R" selected then pure (tick trace "I" 51) else pure (factory trace "J" 53)

failing :: Ref.Ref String -> Ref.Ref Boolean -> Effect Int
failing trace flag = do
  selected <- Ref.read flag
  if unsafePerformEffect (do
    _ <- tick trace "X" unit
    if selected then throw "condition marker" else pure false)
    then pure 1
    else factory trace "Z" 2

main :: Effect Unit
main = do
  trace <- Ref.new ""
  for_ [ true, false ] \initial -> do
    flag <- Ref.new initial
    for_ [ 0, 1, 2 ] \_ -> do
      Ref.write "" trace
      a <- choose trace flag
      b <- nested trace flag
      logShow [ a, b ]
      Ref.read trace >>= log
    Ref.write "" trace
    saved <- makeSaved trace flag
    Ref.read trace >>= log
    logShow =<< saved
    logShow =<< saved
    Ref.read trace >>= log
    Ref.write "" trace
    inner <- returned trace flag
    Ref.read trace >>= log
    logShow =<< inner
    logShow =<< inner
    Ref.read trace >>= log
    Ref.write "" trace
    result <- try (failing trace flag)
    logShow (isLeft result)
    Ref.read trace >>= log
