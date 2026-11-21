module Test.Cancellation where

import Prelude

import Data.Either (Either(..))
import Effect (Effect)
import Effect.Aff (bracket, forkAff, killFiber, runAff_)
import Effect.Aff.AVar as AVar
import Effect.Aff.Compat (EffectFnAff, fromEffectFnAff)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Exception (error, message)
import Example.HTTP (serve)
import Node.Process as Process

-- A deterministic gate controlled by the test runner keeps the application
-- operation pending until cancellation and optional client disconnection.
foreign import awaitHandlerRelease :: EffectFnAff Unit

main :: Effect Unit
main = runAff_ finish do
  started <- AVar.empty
  fiber <- forkAff $ bracket
    (liftEffect $ Console.log "RESOURCE ACQUIRED")
    (\_ -> liftEffect $ Console.log "RESOURCE RELEASED")
    \_ -> serve "127.0.0.1:0" \_ reply -> runAff_ (respond reply) do
      liftEffect $ Console.log "HANDLER STARTED"
      AVar.put unit started
      fromEffectFnAff awaitHandlerRelease
      liftEffect $ Console.log "HANDLER DONE"
      pure { status: 200, body: "{\"done\":true}" }
  AVar.take started
  liftEffect $ Console.log "CANCEL REQUESTED"
  killFiber (error "test caller cancellation") fiber
  liftEffect $ Console.log "CANCELED"
  where
  finish (Left err) = Console.error (message err) *> Process.exit' 1
  finish (Right _) = pure unit
  respond reply (Right response) = reply response
  respond reply (Left err) = do
    Console.error (message err)
    reply { status: 500, body: "{\"error\":\"handler failed\"}" }
