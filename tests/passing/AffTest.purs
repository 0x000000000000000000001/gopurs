module Main where

import Prelude

import Effect (Effect)
import Effect.Console (log, logShow)
import Effect.Aff (Aff, delay, forkAff, joinFiber, launchAff_)
import Effect.Class (liftEffect)
import Effect.Ref as Ref
import Data.Time.Duration (Milliseconds(..))

incrementTask :: Ref.Ref Int -> Aff Unit
incrementTask ref = do
  delay (Milliseconds 1.0)
  liftEffect $ Ref.modify_ (\x -> x + 1) ref

loop :: Int -> Aff Unit -> Aff Unit
loop 0 _ = pure unit
loop n task = do
  task
  loop (n - 1) task

main :: Effect Unit
main = launchAff_ do
  liftEffect $ log "Starting Aff test..."
  ref <- liftEffect $ Ref.new 0
  
  f1 <- forkAff (loop 100 (incrementTask ref))
  f2 <- forkAff (loop 100 (incrementTask ref))
  
  joinFiber f1
  joinFiber f2
  
  finalVal <- liftEffect $ Ref.read ref
  
  if finalVal == 200 then
    liftEffect $ log "Success: 200"
  else do
    liftEffect $ log "Fail: Wrong value"
    liftEffect $ logShow finalVal
