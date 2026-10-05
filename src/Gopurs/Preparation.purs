module Gopurs.Preparation (runPreparationJobs) where

import Prelude

import Control.Lazy (defer)
import Control.Monad.Rec.Class (Step(..), tailRecM)
import Control.Parallel (parTraverse)
import Data.Array as Array
import Data.List (List(..), (:))
import Data.Tuple (Tuple(..), fst, snd)
import Effect.Aff (Aff)
import Effect.Class (liftEffect)
import Effect.Ref as Ref

-- Each round reads a fixed snapshot. Only the caller merges returned results,
-- in input order, before allowing the next round to observe them.
-- PBO's transitiveCollectWith owns that merge, its cache and its stopping rule.
-- This dispatcher only schedules a round: constructing Aff must not run a job,
-- and running the same Aff again must recompute every result.
runPreparationJobs :: forall a. Int -> Array (Unit -> a) -> Aff (Array a)
runPreparationJobs configured tasks
  | configured <= 1 || Array.null tasks = defer \_ -> pure (map (_ $ unit) tasks)
  | otherwise = do
      let
        count = Array.length tasks
        workers = min count (min 8 configured)
        chunkSize = max 1 (min 16 ((count + workers * 8 - 1) / (workers * 8)))
      -- Late rounds contain many cheap cache hits and a few expensive rewrites.
      -- Claim small slices atomically instead of assigning one large contiguous
      -- slice to each worker. The reference is fresh on every execution of Aff.
      next <- liftEffect (Ref.new 0)
      let
        worker = tailRecM
          (\results -> do
            start <- liftEffect $ Ref.modify' (\index -> { state: index + chunkSize, value: index }) next
            if start >= count then pure (Done results)
            else defer \_ -> pure $ Loop
              (Tuple start (map (_ $ unit) (Array.slice start (start + chunkSize) tasks)) : results))
          Nil
      results <- parTraverse (\_ -> worker) (Array.range 0 (workers - 1))
      -- Completion order is deliberately unconstrained; publication order is not.
      pure (Array.concatMap snd (Array.sortWith fst (Array.concatMap Array.fromFoldable results)))
