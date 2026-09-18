module Gopurs.Emission (createEmitter) where

import Prelude

import Data.Array as Array
import Data.Set (Set)
import Data.Set as Set
import Effect (Effect)
import Effect.Aff (Aff)
import Effect.Class (liftEffect)
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
