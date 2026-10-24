module Gopurs.Driver.Prepare.Persistent (Projected, PreloadedProjection, preloadProjection, preparePersistent, preparePersistentForProfile, preparePersistentWithProjection, preparePersistentWithProjectionForProfile, publishPersistent) where

import Prelude
import Data.Argonaut.Core (stringify)
import Data.Argonaut.Decode (class DecodeJson, decodeJson)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Either (Either(..), hush)
import Data.Foldable (foldl)
import Data.List as List
import Data.Lazy (defer)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Effect.Aff (Aff, attempt)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Ref as Ref
import Gopurs.Cache.Wire as Wire
import Gopurs.Cache.LookupTable as LookupTable
import Gopurs.ClassMetadata (addClassDataDeclarations)
import Gopurs.Driver.Cache (Session)
import Gopurs.Driver.Cache as Cache
import Gopurs.Driver.Prepare (PreparedCache, prepareModulesWithCacheForProfile, selectMainModules)
import Gopurs.Driver.Manifest (moduleInterface)
import Gopurs.Driver.Prepare.Proof as Proof
import Gopurs.Driver.Prepare.Projection (Projection)
import Gopurs.Driver.Prepare.Projection as Projection
import Gopurs.Driver.Prepare.ProjectionCapsule as ProjectionCapsule
import Gopurs.Preparation (mapPreparationJobs, runPreparationJobs)
import Gopurs.Metrics as Metrics
import Gopurs.Monomorphization (buildGlobalAstMap)
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import PureScript.Backend.Optimizer.CoreFn.Json.Text (parseModule)
import PureScript.Backend.Optimizer.CoreFn (Ann, Module(..))
import PureScript.Backend.Optimizer.CoreFn.Usage (invalidateSourceUsageModule)
import PureScript.Backend.Optimizer.Monomorphize (InstantiationMap, monomorphizeWithReads, normalizeInstantiation, SpecializationLookup)
import PureScript.Backend.Optimizer.Semantics (InlineDirectiveMap)

directory :: String -> String
directory root = root <> "/preparation-literals-v1"

proofPath :: String -> String
proofPath root = directory root <> "/proof.json"

payloadPath :: String -> String -> String
payloadPath root digest = directory root <> "/" <> digest <> ".json"

owner :: String -> String
owner name = String.joinWith "." (fromMaybe [] (Array.init (String.split (Pattern ".") name)))

readJSON :: forall a. DecodeJson a => String -> Aff (Maybe a)
readJSON path = do
  text <- hush <$> attempt (FS.readTextFile UTF8 path)
  pure (text >>= hush <<< jsonParser >>= hush <<< decodeJson)

-- Bind the preparation base to the exact committed TAST generation, rather
-- than an independently updated sidecar manifest. FFI has its separate proof.
snapshotStamp :: String -> Aff (Maybe String)
snapshotStamp path = do
  parsed <- readJSON path :: Aff (Maybe { snapshot :: { compiler :: String, inputs :: Array { name :: String, kind :: String, digest :: String } } })
  pure do
    document <- parsed
    case Array.filter (\input -> input.name == "tast" && input.kind == "tast") document.snapshot.inputs of
      [ tast ] -> Just (Wire.hash (document.snapshot.compiler <> "/" <> tast.digest))
      _ -> Nothing

type CapturedProof = { hash :: String, value :: Proof.State }

readCapturedProof :: String -> Aff (Maybe CapturedProof)
readCapturedProof root = do
  envelope <- readJSON (proofPath root) :: Aff (Maybe { hash :: String, payload :: String })
  pure do
    saved <- envelope
    if Wire.hash saved.payload /= saved.hash then Nothing
    else do
      value <- hush (jsonParser saved.payload) >>= hush <<< decodeJson
      Just { hash: saved.hash, value }

readProof :: String -> Aff (Maybe Proof.State)
readProof root = map _.value <$> readCapturedProof root

writeAtomic :: String -> String -> String -> Aff Unit
writeAtomic stage path text = do
  let temporary = stage <> "/preparation-" <> Wire.hash path <> ".next"
  FS.writeTextFile UTF8 temporary text
  FS.rename temporary path

readObject :: String -> String -> Aff (Maybe String)
readObject root digest = do
  text <- hush <$> attempt (FS.readTextFile UTF8 (payloadPath root digest))
  pure (text >>= \value -> if Wire.hash value == digest then Just value else Nothing)

-- Only the freeze's complete, distinct inventory can skip global TAST loading.
-- Old sessions and malformed inventories retain the complete preparation path.
readFrozenHashes :: Session -> Aff (Maybe (Map String String))
readFrozenHashes session = do
  frozen <- readJSON (session.stage <> "/source-hashes.json") :: Aff (Maybe (Array Proof.NamedHash))
  pure do
    entries <- frozen
    let hashes = Map.fromFoldable (map (\entry -> Tuple entry.name entry.hash) entries)
    if Array.length entries == session.count && Map.size hashes == session.count
      && Array.all (\entry -> entry.name /= "" && entry.hash /= "") entries
      then Just hashes else Nothing

readSource :: Session -> Map String String -> String -> Aff (Maybe (Module Ann))
readSource session hashes name = case Map.lookup name hashes of
  Nothing -> pure Nothing
  Just digest -> do
    Cache.materializeInputs session [ name ]
    text <- hush <$> attempt (FS.readTextFile UTF8 (session.input <> "/" <> name <> "/corefn.json"))
    pure do
      encoded <- text
      if Wire.hash encoded /= digest then Nothing
      else do
        source@(Module value) <- hush (parseModule encoded)
        if unwrap value.name == name then Just source else Nothing

-- A preload is private data, never a reusable proof on its own. Keep the exact
-- envelope identity so another publication cannot silently re-anchor it.
type PreloadedProjection =
  { root :: String, cacheRoot :: String, context :: String, count :: Int, proofHash :: String
  , projected :: Projected
  }

preloadProjection :: Int -> InlineDirectiveMap -> String -> Aff (Maybe PreloadedProjection)
preloadProjection jobs directives encodedSeed = case hush (jsonParser encodedSeed) >>= hush <<< decodeJson of
  Nothing -> pure Nothing
  Just (seed :: { root :: String, cacheRoot :: String, context :: String }) -> do
    captured <- readCapturedProof seed.cacheRoot
    directivesKey <- liftEffect (Wire.hash <$> Wire.write directives)
    case captured of
      Just saved | saved.value.schema == 5 && saved.value.context == seed.context
        && saved.value.directives == directivesKey -> do
          let count = Array.length saved.value.sources
          result <- readProjection { cacheRoot: seed.cacheRoot, count } jobs saved.value.projection saved.value.capsule
          pure (map (\projected -> { root: seed.root, cacheRoot: seed.cacheRoot, context: seed.context, count, proofHash: saved.hash, projected }) result)
      _ -> pure Nothing

preparePersistent :: Session -> Int -> Maybe String -> InlineDirectiveMap -> Aff PreparedCache
preparePersistent = preparePersistentWithProjection Nothing

preparePersistentForProfile :: Boolean -> Session -> Int -> Maybe String -> InlineDirectiveMap -> Aff PreparedCache
preparePersistentForProfile devMode = preparePersistentWithProjectionForProfile devMode Nothing

preparePersistentWithProjection :: Maybe PreloadedProjection -> Session -> Int -> Maybe String -> InlineDirectiveMap -> Aff PreparedCache
preparePersistentWithProjection = preparePersistentWithProjectionForProfile false

preparePersistentWithProjectionForProfile :: Boolean -> Maybe PreloadedProjection -> Session -> Int -> Maybe String -> InlineDirectiveMap -> Aff PreparedCache
preparePersistentWithProjectionForProfile devMode preloaded session jobs mainModule directives = do
  candidate <- Metrics.measure "compact preparation probe" \_ -> case preloaded of
    Nothing -> probe devMode session jobs mainModule directives
    Just _ -> probeWithProjection devMode preloaded session jobs mainModule directives
  case candidate of
    Just prepared -> do
      liftEffect $ Console.error ("[gopurs] compact preparation: " <> show (Set.size prepared.reusable) <> " modules proven reusable; TAST loaded on demand")
      pure prepared
    Nothing -> do
      Cache.materializeAllInputs session
      prepareModulesWithCacheForProfile devMode session.input session.count jobs mainModule directives

probe :: Boolean -> Session -> Int -> Maybe String -> InlineDirectiveMap -> Aff (Maybe PreparedCache)
probe devMode = probeWithProjection devMode Nothing

probeWithProjection :: Boolean -> Maybe PreloadedProjection -> Session -> Int -> Maybe String -> InlineDirectiveMap -> Aff (Maybe PreparedCache)
probeWithProjection devMode preloaded session jobs mainModule directives = do
  saved <- readCapturedProof session.cacheRoot
  frozen <- readFrozenHashes session
  committed <- snapshotStamp (session.stage <> "/prior/manifest.json")
  currentStamp <- snapshotStamp (session.stage <> "/session.json")
  directivesKey <- liftEffect (Wire.hash <$> Wire.write directives)
  case saved, frozen, committed, currentStamp of
    Just captured@{ value: old }, Just sourceHashes, Just stamp, Just current
      | old.schema == 5 && old.context == session.context && old.directives == directivesKey && old.snapshot == stamp -> do
          let
            previousHashes = Map.fromFoldable (map (\entry -> Tuple entry.name entry.hash) old.sources)
            changed = Array.filter (\(Tuple name hash) -> Map.lookup name previousHashes /= Just hash) (Map.toUnfoldable sourceHashes :: Array _)
            sameInventory = Map.keys previousHashes == Map.keys sourceHashes && Array.length old.sources == session.count
          case changed of
            [ Tuple changedName _ ] | sameInventory -> do
              changedSource <- readSource session sourceHashes changedName
              case changedSource of
                Just source -> do
                  currentProof <- liftEffect (Proof.moduleProofWithReads (Set.fromFoldable old.bodyReads) source)
                  let previousProof = Array.find (\entry -> entry.name == changedName) old.modules
                  case previousProof of
                    Just previous | Proof.permitsLiteralEdit (Set.fromFoldable old.bodyReads) previous currentProof -> do
                      projected <- case preloaded of
                        Just pending | pending.proofHash == captured.hash && pending.root == session.root && pending.cacheRoot == session.cacheRoot
                          && pending.context == session.context && pending.count == session.count ->
                            pure (Just pending.projected)
                        _ -> readProjection session jobs old.projection old.capsule
                      case projected of
                        Just result | validProjection sourceHashes result.projection ->
                          prepareProjected devMode session mainModule directives sourceHashes (old { capsule = result.capsule }) current source currentProof result.projection result.coverage
                        _ -> pure Nothing
                    _ -> pure Nothing
                _ -> pure Nothing
            _ -> pure Nothing
    _, _, _, _ -> pure Nothing

-- Hash every private fragment before scheduling its pure decoder. A missing or
-- corrupt fragment takes the existing full preparation path.
validFragmentDigests :: Int -> Array String -> Boolean
validFragmentDigests count digests = Array.length digests == Projection.fragmentCount count
  && Set.size (Set.fromFoldable digests) == Array.length digests

type Projected =
  { projection :: Projection
  , coverage :: Maybe (Set.Set String)
  , capsule :: Maybe (Array String)
  }

readProjection :: forall r. { cacheRoot :: String, count :: Int | r } -> Int -> Array String -> Maybe (Array String) -> Aff (Maybe Projected)
readProjection session jobs digests capsule
  | not (validFragmentDigests session.count digests) = pure Nothing
  | otherwise = do
      encoded <- traverse (readObject session.cacheRoot) digests
      case traverse identity encoded of
        Nothing -> pure Nothing
        Just texts -> do
          -- Immutable full bytes are verified even on a capsule hit. No I/O is
          -- hidden in the Lazy; malformed complete data stays an explicit Left.
          let delayed = defer \_ -> case traverse Wire.read texts >>= Projection.merge of
                Nothing -> Left "complete projection codec"
                Just projection -> LookupTable.complete projection.metadata.globalTypes
          partial <- case capsule of
            Just refs | validFragmentDigests session.count refs -> do
              saved <- traverse (readObject session.cacheRoot) refs
              case traverse identity saved of
                Nothing -> pure Nothing
                Just bytes -> do
                  decoded <- runPreparationJobs jobs (map (\text _ -> Wire.read text) bytes)
                  pure (traverse identity decoded >>= flip ProjectionCapsule.open delayed)
            _ -> pure Nothing
          case partial of
            Just value -> pure (Just { projection: value.projection, coverage: Just value.coverage, capsule })
            Nothing -> do
              decoded <- runPreparationJobs jobs (map (\text _ -> Wire.read text) texts)
              pure (map (\projection -> { projection, coverage: Nothing, capsule: Nothing })
                (traverse identity decoded >>= Projection.merge))

validProjection :: Map String String -> Projection -> Boolean
validProjection hashes projection =
  let names = map (\(Module value) -> unwrap value.name) (Array.fromFoldable projection.modules)
  in Array.length names == Map.size hashes && Set.fromFoldable names == Map.keys hashes
    && Map.keys projection.interfaces == Map.keys hashes

prepareProjected :: Boolean -> Session -> Maybe String -> InlineDirectiveMap -> Map String String -> Proof.State -> String -> Module Ann -> Proof.ModuleProof -> Projection -> Maybe (Set.Set String) -> Aff (Maybe PreparedCache)
prepareProjected devMode session mainModule directives sourceHashes old current source@(Module changed) currentProof projection coverage = do
  demands <- liftEffect (Ref.new Set.empty)
  let
    index = Map.fromFoldable (map (\entry -> Tuple entry.name (Set.fromFoldable entry.keys)) old.index)
    payloads = Map.fromFoldable (map (\entry -> Tuple entry.name entry.hash) old.payloads)
    bodyReads = Set.fromFoldable old.bodyReads
    enrich = invalidateSourceUsageModule <<< addClassDataDeclarations
  sources <- liftEffect (Ref.new (Map.singleton (unwrap changed.name) (enrich source)))
  let
    load name = do
      available <- liftEffect (Ref.read sources)
      case Map.lookup name available of
        Just value -> pure (Just value)
        Nothing -> do
          loaded <- map enrich <$> readSource session sourceHashes name
          case loaded of
            Just value -> liftEffect (Ref.modify_ (Map.insert name value) sources)
            Nothing -> pure unit
          pure loaded
    -- A missed lookup is itself recorded. Load its owner and replay the pure
    -- transformation until all lookups have the same answers as the full AST.
    -- A lookup not covered by the committed full-preparation proof fails shut.
    rewrite own actual = do
      available <- liftEffect (Ref.read sources)
      let
        ast = Map.delete "Data.Ring.negate" (buildGlobalAstMap (List.fromFoldable (Map.values available)))
        Tuple result reads = monomorphizeWithReads ast (hashedLookup index) own actual
        pending = Set.filter (\name -> Map.member name sourceHashes && not (Map.member name available)) (Set.map owner reads)
      if not (Set.isEmpty (Set.difference reads bodyReads)) then pure Nothing
      else if Set.isEmpty pending then pure (Just result)
      else do
        loaded <- traverse load (Set.toUnfoldable pending :: Array String)
        if Array.all (case _ of
          Just _ -> true
          _ -> false) loaded then rewrite own actual
        else pure Nothing
    materialize name = do
      actual <- load name
      own <- readOwn session.cacheRoot index payloads name
      case actual, own of
        Just value, Just instantiations ->
          if Map.isEmpty index then pure (Just value) else rewrite instantiations value
        _, _ -> pure Nothing
  rewritten <- materialize (unwrap changed.name)
  pure do
    changedPrepared <- rewritten
    let
      preparedModules = map (\header@(Module value) -> if value.name == changed.name then changedPrepared else header) projection.modules
      prepared =
        { directives, modules: preparedModules, metadata: projection.metadata, metadataUnchanged: true
        , metadataDemands: Just (\keys -> Ref.modify_ (Set.union keys) demands)
        , mainModules: case mainModule of
            Just name -> [ name ]
            Nothing -> projection.mainModules
        , sources: map (\(Module value) -> { name: unwrap value.name, path: value.path, hasForeign: not (Map.isEmpty value.foreign) }) (Array.fromFoldable projection.modules)
        }
      proof = old
        { snapshot = current
        , sources = map (\(Tuple name hash) -> { name, hash }) (Map.toUnfoldable sourceHashes :: Array _)
        , modules = map (\entry -> if entry.name == currentProof.name then currentProof else entry) old.modules
        }
    Just
      { prepared, devMode, sourceHashes, sourceModules: projection.modules, sourceInterfaces: projection.interfaces
      , compactSourceDelta: Just (unwrap changed.name)
      , mono: { globalAstMap: Map.empty, instantiations: Map.empty, bodyReads }
      , reusable: Set.delete (unwrap changed.name) (Map.keys sourceHashes)
      , proof: Just proof, metadataDemandKeys: Just demands, metadataCoverage: coverage
      , metadataProjection: Just projection, materialize: Just materialize
      }

-- Canonical specialization keys can contain very large nested type names.
-- Hash their complete strings with the same SHA-256 used by cache fingerprints.
hashedLookup :: Map String (Set.Set String) -> SpecializationLookup
hashedLookup index name = (\keys key -> Set.member (Wire.hash key) keys) <$> Map.lookup name index

readOwn :: String -> Map String (Set.Set String) -> Map String String -> String -> Aff (Maybe InstantiationMap)
readOwn root index payloads name = do
  let expected = Map.filterKeys (\global -> owner global == name) index
  case Map.lookup name payloads of
    Nothing -> pure (if Map.isEmpty expected then Just Map.empty else Nothing)
    Just digest -> do
      text <- hush <$> attempt (FS.readTextFile UTF8 (payloadPath root digest))
      pure do
        encoded <- text
        if Wire.hash encoded /= digest then Nothing
        else do
          decoded <- Wire.read encoded
          if map (Set.map Wire.hash <<< Map.keys) decoded == expected then Just decoded else Nothing

-- Call only after successful publication. Content-addressed own-payload files
-- are immutable; replacing the proof is one atomic rename. A failed later build
-- cannot pair this proof with different committed cache objects.
publishPersistent :: Session -> Int -> PreparedCache -> Aff Unit
publishPersistent session jobs cache = do
  result <- attempt (publishPersistentImpl session jobs cache)
  case result of
    Left _ -> liftEffect $ Console.error "[gopurs] compact preparation publication skipped: cache write unavailable"
    Right _ -> pure unit

publishPersistentImpl :: Session -> Int -> PreparedCache -> Aff Unit
publishPersistentImpl session jobs cache = do
  stamp <- snapshotStamp (session.stage <> "/session.json")
  case stamp of
    Nothing -> pure unit
    Just snapshot -> do
      void $ attempt (FS.mkdir (directory session.cacheRoot))
      directives <- liftEffect (Wire.hash <$> Wire.write cache.prepared.directives)
      reusable <- case cache.proof of
        Just proof -> pure (Just proof)
        Nothing -> do
          previous <- readProof session.cacheRoot
          case previous of
            Just proof | proof.schema == 5 && proof.snapshot == snapshot
              && validFragmentDigests session.count proof.projection
              && proof.context == session.context && proof.directives == directives -> do
                valid <- Array.all identity <$> traverse (\entry -> do
                  text <- hush <$> attempt (FS.readTextFile UTF8 (payloadPath session.cacheRoot entry.hash))
                  pure (map Wire.hash text == Just entry.hash)) (proof.payloads <> map (\hash -> { name: "projection", hash }) proof.projection)
                pure (if valid then Just proof else Nothing)
            _ -> pure Nothing
      state <- case reusable of
        Just proof -> pure (proof { snapshot = snapshot })
        Nothing -> do
          -- Each proof owns its Wire encoder state. Join in source order before
          -- publishing any object, using the same bounded preparation workers.
          modules <- Metrics.measure "prepare source proofs" \_ ->
            mapPreparationJobs jobs (liftEffect <<< Proof.moduleProofWithReads cache.mono.bodyReads)
              (Array.fromFoldable cache.sourceModules)
          let
            index = map (\(Tuple name values) -> { name, keys: map Wire.hash (Set.toUnfoldable (Map.keys values) :: Array String) }) (Map.toUnfoldable cache.mono.instantiations :: Array _)
            grouped = foldl
              (\groups (Tuple name values) -> Map.insertWith Map.union (owner name) (Map.singleton name values) groups)
              Map.empty (Map.toUnfoldable cache.mono.instantiations :: Array _)
          payloads <- Metrics.measure "prepare specialization payloads" \_ -> traverse (\(Tuple name values) -> do
            encoded <- liftEffect (Wire.write (map (map normalizeInstantiation) values))
            let hash = Wire.hash encoded
            writeAtomic session.stage (payloadPath session.cacheRoot hash) encoded
            pure { name, hash }) (Map.toUnfoldable grouped :: Array _)
          interfaces <- Metrics.measure "prepare source interfaces" \_ ->
            liftEffect (Map.fromFoldable <$> traverse (\mod@(Module value) -> Tuple (unwrap value.name) <$> moduleInterface mod) (Array.fromFoldable cache.sourceModules))
          projection <- Metrics.measure "prepare projection fragments" \_ -> traverse (\fragment -> do
            encoded <- liftEffect (Wire.write fragment)
            let hash = Wire.hash encoded
            writeAtomic session.stage (payloadPath session.cacheRoot hash) encoded
            pure hash) (Projection.split
              { modules: map Proof.compactHeader cache.prepared.modules
              , metadata: cache.prepared.metadata
              , interfaces
              , mainModules: selectMainModules Nothing (Array.fromFoldable cache.sourceModules)
              })
          pure
            { schema: 5, context: session.context, directives, snapshot
            , sources: map (\(Tuple name hash) -> { name, hash }) (Map.toUnfoldable cache.sourceHashes :: Array _)
            , modules, index, payloads, projection, capsule: Nothing, bodyReads: Set.toUnfoldable cache.mono.bodyReads
            }
      withCapsule <- publishCapsule session jobs cache state
      let payload = stringify (encodeJson withCapsule)
      writeAtomic session.stage (proofPath session.cacheRoot) (stringify (encodeJson { hash: Wire.hash payload, payload }))

-- Capsule references are committed in the same proof as their complete source
-- fragments. A full oracle preserves an existing capsule; only a compact run
-- with new explicit demands creates one from the verified complete table.
publishCapsule :: Session -> Int -> PreparedCache -> Proof.State -> Aff Proof.State
publishCapsule session jobs cache state = case cache.metadataDemandKeys, cache.metadataProjection of
  Just reference, Just projection -> do
    requested <- liftEffect (Ref.read reference)
    let coverage = Set.union requested (fromMaybe Set.empty cache.metadataCoverage)
    if Set.isEmpty requested || (case cache.metadataCoverage of
      Just known -> Set.isEmpty (Set.difference requested known)
      Nothing -> false) then pure state
    else do
      -- Join all encoders before writing any object or publishing its proof.
      -- Wire's mutable encoder state belongs to each independent invocation.
      encoded <- mapPreparationJobs jobs (liftEffect <<< Wire.write)
        (ProjectionCapsule.create coverage projection)
      fragments <- traverse (\text -> do
        let digest = Wire.hash text
        writeAtomic session.stage (payloadPath session.cacheRoot digest) text
        pure digest) encoded
      pure (state { capsule = Just fragments })
  _, _ -> case state.capsule of
    Nothing -> pure state
    Just refs -> do
      objects <- traverse (readObject session.cacheRoot) refs
      pure (if validFragmentDigests session.count refs && Array.all (case _ of
        Just _ -> true
        Nothing -> false) objects then state else state { capsule = Nothing })
