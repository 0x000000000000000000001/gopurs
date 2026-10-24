module Gopurs.Cache.LookupTable
  ( LookupTable, fromMap, covered, complete, completeOrFail, lookupChecked, lookup, member
  ) where

import Prelude
import Data.Either (Either(..), either)
import Data.Foldable (all)
import Data.Lazy (Lazy, force)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), isJust)
import Data.Set (Set)
import Data.Set as Set
import Effect.Exception (throw)
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.Cache.Wire as Wire

data LookupTable a
  = Full (Map String a)
  | Covered (Set String) (Map String a) (Lazy (Either String (Map String a)))

fromMap :: forall a. Map String a -> LookupTable a
fromMap = Full

-- The caller must verify the capsule against its accepted preparation proof.
-- Coverage is a set of answers built from a complete table, never a claim that
-- these are all possible requests. Every uncovered key forces the full table.
covered :: forall a. Set String -> Map String a -> Lazy (Either String (Map String a)) -> Maybe (LookupTable a)
covered coverage values delayed
  | all (flip Set.member coverage) (Map.keys values) = Just (Covered coverage values delayed)
  | otherwise = Nothing

complete :: forall a. LookupTable a -> Either String (Map String a)
complete = case _ of
  Full values -> Right values
  Covered _ _ delayed -> force delayed

cacheFailure :: forall a. String -> a
cacheFailure reason = unsafePerformEffect (throw ("incremental cache: " <> reason))

completeOrFail :: forall a. LookupTable a -> Map String a
completeOrFail = either cacheFailure identity <<< complete

lookupChecked :: forall a. String -> LookupTable a -> Either String (Maybe a)
lookupChecked key = case _ of
  Full values -> Right (Map.lookup key values)
  table@(Covered coverage values _)
    | Set.member key coverage -> Right (Map.lookup key values)
    | otherwise -> Map.lookup key <$> complete table

-- The Lazy memoizes Either, so its native sync.Once never catches a thrown
-- decoder failure. This adapter raises only an explicit invalid-cache error;
-- it performs no I/O, tracing or mutation. New valid requests return normally.
lookup :: forall a. String -> LookupTable a -> Maybe a
lookup key = either cacheFailure identity <<< lookupChecked key

member :: forall a. String -> LookupTable a -> Boolean
member key = isJust <<< lookup key

-- Keep the exact complete Map wire format. No partial table can be serialized
-- as complete, including metadata fingerprints and full preparation output.
instance codecLookupTable :: Wire.Codec a => Wire.Codec (LookupTable a) where
  encode context table = do
    values <- either (throw <<< ("incremental cache: " <> _)) pure (complete table)
    Wire.encode context values
  decode context json = Full <$> Wire.decode context json
