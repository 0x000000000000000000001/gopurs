module Gopurs.Emission (createEmitter, createPipelinedEmitter) where

import Prelude

import Control.Lazy (defer)
import Data.Array as Array
import Data.Either (either)
import Data.Maybe (Maybe(..))
import Data.Set (Set)
import Data.Set as Set
import Effect (Effect)
import Effect.Aff (Aff, attempt, forkAff, invincible, joinFiber, throwError)
import Effect.Class (liftEffect)
import Effect.Exception (error)
import Effect.Ref as Ref
import PureScript.Backend.Optimizer.CoreFn (ModuleName)

-- Only the sequential builder accesses the queue. Workers receive a completed
-- batch, and their results must be published before a dependent batch starts.
createEmitter
  :: forall a
   . Int
  -> (Array a -> Aff Unit)
  -> Effect
       { enqueue :: { name :: ModuleName, imports :: Set ModuleName, value :: a } -> Aff Unit
       , finish :: Aff Unit
       }
createEmitter jobs emit = do
  pending <- Ref.new []
  let
    finish = do
      batch <- liftEffect (Ref.read pending)
      unless (Array.null batch) do
        emit (map _.value batch)
        liftEffect (Ref.write [] pending)
    enqueue entry
      | jobs <= 1 = emit [ entry.value ]
      | otherwise = do
          batch <- liftEffect (Ref.read pending)
          -- Optimizer imports include references introduced by cross-module
          -- inlining, so checking only the original TAST imports is insufficient.
          when (Array.any (\queued -> Set.member queued.name entry.imports) batch) finish
          next <- liftEffect (Ref.modify (flip Array.snoc entry) pending)
          when (Array.length next >= jobs) finish
  pure { enqueue, finish }

-- The producer alone owns both queue references. It can prepare the next batch
-- while one worker emits, but waits before starting another worker. For jobs > 1,
-- cancel stops admission and drains the active worker without starting pending
-- work. It is non-preemptive: filesystem callbacks must finish before cleanup
-- returns. The producer must call it when exiting without finishing.
createPipelinedEmitter
  :: forall a
   . Int
  -> (Array a -> Aff Unit)
  -> Effect
       { enqueue :: { name :: ModuleName, imports :: Set ModuleName, value :: a } -> Aff Unit
       , finish :: Aff Unit
       , cancel :: Aff Unit
       }
createPipelinedEmitter jobs emit
  | jobs <= 1 = do
      emitter <- createEmitter jobs emit
      pure { enqueue: emitter.enqueue, finish: emitter.finish, cancel: pure unit }
  | otherwise = do
      active <- Ref.new Nothing
      cancelled <- Ref.new false
      let
        joinActive = do
          current <- liftEffect (Ref.read active)
          case current of
            Nothing -> pure unit
            Just fiber -> do
              result <- joinFiber fiber
              either throwError pure result
              liftEffect (Ref.write Nothing active)

        launch batch = do
          joinActive
          -- Register ownership before cancellation may resume in the producer.
          -- Construct emit inside the worker so its metadata snapshot is taken
          -- after the preceding worker has published its results.
          invincible do
            fiber <- forkAff (attempt (defer \_ -> emit batch))
            liftEffect (Ref.write (Just fiber) active)

        cancel = invincible do
          liftEffect (Ref.write true cancelled)
          current <- liftEffect (Ref.read active)
          case current of
            Nothing -> pure unit
            Just fiber -> do
              -- Killing an Aff using nonCanceler can leave its OS write running.
              -- Drain naturally, preserving the producer's primary failure.
              void (joinFiber fiber)
              liftEffect (Ref.write Nothing active)

      emitter <- createEmitter jobs launch
      let
        enqueue entry = do
          stopped <- liftEffect (Ref.read cancelled)
          when stopped (throwError (error "Go emission cancelled"))
          emitter.enqueue entry
        finish = do
          stopped <- liftEffect (Ref.read cancelled)
          unless stopped do
            emitter.finish
            joinActive
      pure { enqueue, finish, cancel }
