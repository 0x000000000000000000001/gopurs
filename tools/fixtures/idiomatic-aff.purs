module Main where

-- @dependencies: prelude effect console aff either refs

import Prelude
import Data.Either (Either(..))
import Effect (Effect)
import Effect.Aff (Aff, Milliseconds(..), attempt, catchError, delay, forkAff, killFiber, launchAff_)
import Effect.Aff.Compat (EffectFnAff, fromEffectFnAff)
import Effect.Class (liftEffect)
import Effect.Console (log, logShow)
import Effect.Exception (error, message)
import Effect.Ref as Ref
import Effect.Uncurried (EffectFn1, mkEffectFn1)

foreign import hello :: String -> EffectFnAff String
foreign import notify :: EffectFnAff Unit
foreign import noError :: EffectFnAff String
foreign import fire :: EffectFnAff Unit
foreign import readCalls :: Effect Int
foreign import label :: Effect String
foreign import runCallback :: EffectFn1 String Unit -> Effect Unit
foreign import pending :: EffectFnAff String
foreign import isStarted :: Effect Boolean
foreign import isStopped :: Effect Boolean

waitFor :: Effect Boolean -> Aff Unit
waitFor check = do
  ready <- liftEffect check
  unless ready do
    delay (Milliseconds 1.0)
    waitFor check

main :: Effect Unit
main = launchAff_ $ catchError program \err ->
  liftEffect (log ("unexpected Aff error: " <> message err))

program :: Aff Unit
program = do
  let action = fromEffectFnAff (hello "Go")
  liftEffect (readCalls >>= logShow)
  first <- action
  second <- action
  liftEffect do
    log first
    log second
    readCalls >>= logShow
  failure <- attempt (fromEffectFnAff (hello "error"))
  liftEffect $ case failure of
    Left err -> log (message err)
    Right value -> log ("unexpected: " <> value)
  fromEffectFnAff notify
  fromEffectFnAff fire
  result <- fromEffectFnAff noError
  liftEffect do
    log result
    label >>= log
    runCallback (mkEffectFn1 \value -> log ("callback=" <> value))
  completions <- liftEffect (Ref.new 0)
  fiber <- forkAff do
    _ <- fromEffectFnAff pending
    liftEffect (Ref.modify_ (_ + 1) completions)
  waitFor isStarted
  killFiber (error "fixture cancellation") fiber
  waitFor isStopped
  liftEffect do
    log "context canceled"
    Ref.read completions >>= logShow
