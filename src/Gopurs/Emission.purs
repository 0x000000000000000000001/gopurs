module Gopurs.Emission (createEmitter, createPipelinedEmitter) where

import Prelude

import Control.Lazy (defer)
import Data.Array as Array
import Data.Either (Either(..), either)
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

-- The producer owns the chain; batches are appended as fibers that first wait
-- for their predecessor, so emission stays strictly ordered while the producer
-- keeps converting. When more than `cap` batches are in flight, the producer
-- waits for the oldest one only (fine backpressure) instead of draining the
-- whole chain. Failures are sticky: the first error is stored and rethrown by
-- later enqueues and by `finish`.
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
      tail <- Ref.new Nothing
      inFlight <- Ref.new []
      cancelled <- Ref.new false
      failure <- Ref.new Nothing
      let
        cap = max 2 (jobs * 2)

        joinTail :: Aff Unit
        joinTail = do
          current <- liftEffect (Ref.read tail)
          case current of
            Nothing -> pure unit
            Just fiber -> do
              result <- attempt (joinFiber fiber)
              liftEffect (Ref.write Nothing tail)
              liftEffect (Ref.write [] inFlight)
              either throwError pure result

        waitForSlot = do
          queued <- liftEffect (Ref.read inFlight)
          when (Array.length queued >= cap) do
            case Array.uncons queued of
              Nothing -> pure unit
              Just { head: oldest, tail: rest } -> do
                result <- attempt (joinFiber oldest)
                liftEffect (Ref.write rest inFlight)
                case result of
                  Left err -> do
                    liftEffect (Ref.write (Just err) failure)
                    throwError err
                  Right _ -> pure unit
                waitForSlot

        launch batch = do
          waitForSlot
          previous <- liftEffect (Ref.read tail)
          invincible do
            fiber <- forkAff do
              result <- attempt do
                case previous of
                  Nothing -> pure unit
                  Just previousFiber -> void (joinFiber previousFiber)
                emit batch
              case result of
                Left err -> do
                  liftEffect (Ref.write (Just err) failure)
                  throwError err
                Right _ -> pure unit
            liftEffect (Ref.write (Just fiber) tail)
            liftEffect (Ref.modify_ (flip Array.snoc fiber) inFlight)

      emitter <- createEmitter jobs launch
      let
        enqueue entry = do
          stopped <- liftEffect (Ref.read cancelled)
          when stopped (throwError (error "Go emission cancelled"))
          stored <- liftEffect (Ref.read failure)
          case stored of
            Just err -> throwError err
            Nothing -> emitter.enqueue entry
        finish = do
          stopped <- liftEffect (Ref.read cancelled)
          unless stopped do
            emitter.finish
            joinTail
        cancel = do
          liftEffect (Ref.write true cancelled)
          -- Best effort: drain whatever is already in flight, ignoring errors
          -- (the build is failing for another reason).
          void (attempt joinTail)
      pure { enqueue, finish, cancel }
