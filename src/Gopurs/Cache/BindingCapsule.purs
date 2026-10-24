module Gopurs.Cache.BindingCapsule (Capsule, codec, decode, encode, lookup, select) where

import Prelude

import Control.Alternative (guard)
import Data.Argonaut.Core (stringify)
import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Either (hush)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), isJust)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.Set (Set)
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Gopurs.Cache.Wire as Wire
import PureScript.Backend.Optimizer.Analysis (BackendAnalysis)
import PureScript.Backend.Optimizer.Convert (BackendImplementations)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (ExternImpl)
import PureScript.Backend.Optimizer.Observation as Observation

type Implementation = Tuple BackendAnalysis ExternImpl
type Anchor = { owner :: String, sha256 :: String }
type Coverage = { owner :: String, ident :: String, present :: Boolean }
type Envelope =
  { codec :: String, context :: String, reader :: String
  , anchors :: Array Anchor, coverage :: Array Coverage
  , payload :: String, payloadSha256 :: String
  }
type Capsule =
  { anchors :: Map String String
  , coverage :: Map (Qualified Ident) Boolean
  , values :: BackendImplementations
  }

codec :: String
codec = "pbo/binding-capsule-v1"

qualified :: Coverage -> Qualified Ident
qualified entry = Qualified (Just (ModuleName entry.owner)) (Ident entry.ident)

ordered :: forall a. (a -> a -> Ordering) -> Array a -> Boolean
ordered comparison values = Array.all (\(Tuple a b) -> comparison a b == LT) (Array.zip values (Array.drop 1 values))

-- The caller has already checked the text against its committed object SHA.
-- A capsule is only private lookup coverage, never a complete owner map.
decode :: String -> String -> Map String String -> String -> Maybe Capsule
decode context reader current text = do
  envelope <- hush (jsonParser text) >>= hush <<< decodeJson :: Maybe Envelope
  guard (envelope.codec == codec && envelope.context == context && envelope.reader == reader)
  guard (Wire.hash envelope.payload == envelope.payloadSha256)
  guard (ordered (\a b -> Wire.compareKey a.owner b.owner) envelope.anchors)
  guard (ordered (\a b -> Wire.compareKey (qualified a) (qualified b)) envelope.coverage)
  guard (Array.all (\anchor -> Map.lookup anchor.owner current == Just anchor.sha256) envelope.anchors)
  let anchors = Map.fromFoldable (map (\anchor -> Tuple anchor.owner anchor.sha256) envelope.anchors)
  let coverage = Map.fromFoldable (map (\entry -> Tuple (qualified entry) entry.present) envelope.coverage)
  guard (Map.keys anchors == Set.fromFoldable (map _.owner envelope.coverage))
  values <- Wire.read envelope.payload
  guard (Array.all (\entry -> Map.member (qualified entry) values == entry.present) envelope.coverage)
  guard (Array.all (\key -> Map.lookup key coverage == Just true) (Array.fromFoldable (Map.keys values)))
  pure { anchors, coverage, values }

-- Every entry comes from a complete verified owner or from previously verified
-- capsule coverage. Unknown reads, including false future reads, are omitted.
encode :: String -> String -> Map String String -> Map (Qualified Ident) (Maybe Implementation) -> Effect (Maybe String)
encode context reader current entries
  | Map.isEmpty entries = pure Nothing
  | otherwise = case rows of
      Nothing -> pure Nothing
      Just coverage -> do
        payload <- Wire.write (Map.mapMaybe identity entries)
        let envelope =
              { codec, context, reader
              , anchors: Array.sortBy (\a b -> Wire.compareKey a.owner b.owner) $ Array.mapMaybe
                  (\owner -> map (\sha256 -> { owner, sha256 }) (Map.lookup owner current))
                  (Set.toUnfoldable (Set.fromFoldable (map _.owner coverage)) :: Array String)
              , coverage: Array.sortBy (\a b -> Wire.compareKey (qualified a) (qualified b)) coverage
              , payload, payloadSha256: Wire.hash payload
              }
        pure (Just (stringify (encodeJson envelope)))
  where
  rows = traverse row (Map.toUnfoldable entries :: Array (Tuple (Qualified Ident) (Maybe Implementation)))
  row (Tuple (Qualified (Just owner) ident) value) = do
    guard (Map.member (unwrap owner) current)
    pure { owner: unwrap owner, ident: unwrap ident, present: isJust value }
  row _ = Nothing

-- Nothing means not loaded. Just Nothing is a verified negative query.
lookup :: Qualified Ident -> Capsule -> Maybe (Maybe Implementation)
lookup key capsule = map (const (Map.lookup key capsule.values)) (Map.lookup key capsule.coverage)

select
  :: Set String -> Map String String -> Map String (Maybe BackendImplementations)
  -> Maybe Capsule -> Observation.Reads -> Map (Qualified Ident) (Maybe Implementation)
select safe anchors complete capsule reads = Map.mapMaybeWithKey known reads.implementations
  where
  known key@(Qualified owner _) observed = do
    name <- map unwrap owner
    guard (Set.member name safe && Map.member name anchors)
    value <- case Map.lookup name complete of
      Just (Just implementations) -> Just (Map.lookup key implementations)
      _ -> capsule >>= lookup key
    -- A false observation is not an absence unless a complete verified map or
    -- already verified coverage actually lacks the binding.
    guard (isJust value == observed)
    pure value
