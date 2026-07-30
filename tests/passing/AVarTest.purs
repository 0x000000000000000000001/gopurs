module Main where

import Prelude
import Effect (Effect)
import Effect.Console (log, logShow)
import Effect.Aff (Aff, launchAff_, forkAff, joinFiber, delay, Milliseconds(..))
import Effect.Class (liftEffect)
import Effect.AVar as AVar
import Effect.Aff.AVar as AffAVar

producer :: AVar.AVar String -> Aff Unit
producer avar = do
  delay (Milliseconds 5.0)
  liftEffect $ log "Producer: Putting value..."
  AffAVar.put "Go + PureScript AVar" avar
  liftEffect $ log "Producer: Done!"

consumer :: AVar.AVar String -> Aff Unit
consumer avar = do
  liftEffect $ log "Consumer: Waiting for value..."
  val <- AffAVar.take avar
  liftEffect $ log "Consumer: Got value!"
  
  if val == "Go + PureScript AVar" then
    liftEffect $ log "Success: AVar value matches"
  else do
    liftEffect $ log "Fail: Wrong AVar value"
    liftEffect $ logShow val

main :: Effect Unit
main = launchAff_ do
  liftEffect $ log "Starting AVar test..."
  avar <- liftEffect AVar.empty
  
  f1 <- forkAff (consumer avar)
  f2 <- forkAff (producer avar)
  
  joinFiber f1
  joinFiber f2
  
  liftEffect $ log "Test completed."
