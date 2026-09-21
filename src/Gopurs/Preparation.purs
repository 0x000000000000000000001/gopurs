module Gopurs.Preparation (runPreparationJobs) where

import Prelude

import Control.Lazy (defer)
import Control.Parallel (parTraverse)
import Data.Array as Array
import Effect.Aff (Aff)

-- Each round reads a fixed snapshot. Only the caller merges returned results,
-- in input order, before allowing the next round to observe them.
runPreparationJobs :: forall a. Int -> Array (Unit -> a) -> Aff (Array a)
runPreparationJobs configured tasks
  | configured <= 1 || Array.null tasks = defer \_ -> pure (map (_ $ unit) tasks)
  | otherwise = do
      let
        workers = min 8 configured
        count = Array.length tasks
        chunkSize = (count + workers - 1) / workers
        chunks = Array.filter (not <<< Array.null) $ map
          (\index -> Array.slice (index * chunkSize) ((index + 1) * chunkSize) tasks)
          (Array.range 0 (workers - 1))
      -- Partition once: repeatedly slicing the remaining jobs copies a growing
      -- amount of data. Each worker evaluates its chunk only when Aff runs it.
      results <- parTraverse (\chunk -> defer \_ -> pure (map (_ $ unit) chunk)) chunks
      pure (Array.concat results)
