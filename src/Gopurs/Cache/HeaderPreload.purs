module Gopurs.Cache.HeaderPreload
  ( Raw, Reservoir, withPreload, withDecodedPreload, withDecodedUsing, withOptional
  , readUsing, decode, preloadUsing, select, decodeDirectives ) where

import Prelude

import Control.Alternative (guard)
import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Either (Either(..), hush)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.Traversable (sequence)
import Data.Tuple (Tuple(..))
import Effect.Aff (Aff, attempt, bracket, joinFiber, launchAff)
import Effect.Class (liftEffect)
import Effect.Exception as Exception
import Gopurs.Cache.Wire as Wire
import Gopurs.Driver.Cache (PartText, Session)
import Gopurs.Driver.Cache as Cache
import Gopurs.Preparation (mapPreparationJobs)
import Gopurs.Metrics as Metrics
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import PureScript.Backend.Optimizer.CoreFn (Qualified(..))
import PureScript.Backend.Optimizer.Semantics (EvalRef(..), InlineDirectiveMap)

newtype Raw = Raw
  { stage :: String, root :: String, context :: String, count :: Int
  , entries :: Array PartText
  }

-- These are private decoded values, not reuse evidence. Only select can expose
-- them for the current session's independently validated headerSafe owners.
newtype Reservoir = Reservoir
  { stage :: String, root :: String, context :: String, count :: Int
  , headers :: Map String InlineDirectiveMap
  }

-- Only byte reading overlaps preparation. Decoding starts explicitly at the
-- FFI boundary, avoiding contention with preparation's own codec workers.
withPreload :: forall a. Session -> Boolean -> (Maybe (Aff (Maybe Raw)) -> Aff a) -> Aff a
withPreload session enabled = withOptional
  (if enabled then Just (readUsing session (readOwners session) (\owners -> Cache.readParts session owners [ "2" ])) else Nothing)

withDecodedPreload :: forall a. Int -> Maybe (Aff (Maybe Raw)) -> (Maybe (Aff (Maybe Reservoir)) -> Aff a) -> Aff a
withDecodedPreload jobs pending = withDecodedUsing (decode jobs)
  (map (\wait -> Metrics.measure "module header read join" (\_ -> wait)) pending)

-- Join the raw reader before launching codec workers. The caller nests this
-- bracket inside the reader's bracket, so both children finish before close.
withDecodedUsing :: forall a b c. (a -> Aff (Maybe b)) -> Maybe (Aff (Maybe a)) -> (Maybe (Aff (Maybe b)) -> Aff c) -> Aff c
withDecodedUsing _ Nothing use = use Nothing
withDecodedUsing decodeValue (Just pending) use = do
  raw <- pending
  withOptional (map decodeValue raw) use

-- The child is independent of caller cancellation. The bracket's protected
-- release always joins it before the outer session can close or retry.
withOptional :: forall a b. Maybe (Aff (Maybe a)) -> (Maybe (Aff (Maybe a)) -> Aff b) -> Aff b
withOptional Nothing use = use Nothing
withOptional (Just task) use = bracket
  (liftEffect $ Exception.try $ launchAff (fromMaybe Nothing <<< hush <$> attempt task))
  (case _ of
    Left _ -> pure unit
    Right fiber -> void (attempt (joinFiber fiber)))
  (case _ of
    Left _ -> use Nothing
    Right fiber -> use (Just (fromMaybe Nothing <<< hush <$> attempt (joinFiber fiber))))

-- Frozen source names are work hints only. Unknown prior owners or malformed
-- inventories abandon this optional work; normal late reads remain authoritative.
readOwners :: Session -> Aff (Maybe (Array String))
readOwners session = do
  text <- FS.readTextFile UTF8 (session.stage <> "/source-hashes.json")
  pure do
    json <- hush (jsonParser text)
    entries <- hush (decodeJson json) :: Maybe (Array { name :: String, hash :: String })
    let names = map _.name entries
    guard (Array.length names == session.count && Set.size (Set.fromFoldable names) == session.count)
    guard (Array.all (\entry -> entry.name /= "" && entry.hash /= "") entries)
    pure names

-- The production reader is Cache.readParts, which verifies bytes against the
-- pinned private manifest. An explicit reader also permits native lifecycle
-- tests without adding fault hooks to storage or bypassing its production checks.
preloadUsing :: Int -> Session -> Aff (Maybe (Array String)) -> (Array String -> Aff (Array PartText)) -> Aff (Maybe Reservoir)
preloadUsing jobs session inventory readParts = do
  raw <- readUsing session inventory readParts
  case raw of
    Nothing -> pure Nothing
    Just value -> decode jobs value

readUsing :: Session -> Aff (Maybe (Array String)) -> (Array String -> Aff (Array PartText)) -> Aff (Maybe Raw)
readUsing session inventory readParts = do
  owners <- inventory
  case owners of
    Nothing -> pure Nothing
    Just names -> do
      let expected = Set.fromFoldable names
      if Array.length names /= session.count || Set.size expected /= session.count then pure Nothing
      else do
        entries <- readParts names
        let actual = Set.fromFoldable (map _.name entries)
        if Array.length entries /= session.count || actual /= expected || not (Array.all (\entry -> entry.part == "2") entries) then pure Nothing
        else pure (Just (Raw { stage: session.stage, root: session.root, context: session.context, count: session.count, entries }))

decode :: Int -> Raw -> Aff (Maybe Reservoir)
decode jobs (Raw saved) = do
  decoded <- mapPreparationJobs jobs (\entry -> liftEffect $ pure (Tuple entry.name <$> decodeDirectives entry.name entry.text)) saved.entries
  pure do
    headers <- Map.fromFoldable <$> sequence decoded
    pure (Reservoir { stage: saved.stage, root: saved.root, context: saved.context, count: saved.count, headers })

decodeDirectives :: String -> String -> Maybe InlineDirectiveMap
decodeDirectives name text = do
  value <- Wire.read text
  guard (Array.all (case _ of
    EvalExtern (Qualified owner _) -> map unwrap owner == Just name
    _ -> false) (Map.keys value # Array.fromFoldable))
  pure value

select :: Session -> Set String -> Reservoir -> Maybe (Map String InlineDirectiveMap)
select session headerSafe (Reservoir saved) = do
  guard (saved.stage == session.stage && saved.root == session.root && saved.context == session.context && saved.count == session.count)
  guard (Array.all (flip Map.member saved.headers) (Set.toUnfoldable headerSafe :: Array String))
  pure (Map.filterWithKey (\name _ -> Set.member name headerSafe) saved.headers)
