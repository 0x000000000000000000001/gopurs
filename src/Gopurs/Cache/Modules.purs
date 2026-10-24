module Gopurs.Cache.Modules (ModuleCache, Memo, Keys, Part, initialize, loadHeaderTexts, loadHeaderTextsWithPreload, lookupAll, lookupHeaders, lookupBackend, lookupBackendWithReads, captureReads, emit, emitWithWorkerSnapshot, resolveWorkerSnapshot, warmWorkerSnapshot, readWorkerContribution, report, loadReads, headerSafeFromReads, publishReads, implementationOwners, warmImplementations, warmImplementationInputs, readCachedBinding, readImplementationsSync) where

import Prelude

import Control.Alternative (guard)
import Data.Argonaut.Core (Json, stringify)
import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Either (Either(..), hush)
import Data.Foldable (foldl, foldM, for_)
import Data.List (List(..), (:))
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust, maybe)
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Traversable (sequence, traverse)
import Data.Tuple (Tuple(..))
import Effect.Aff (Aff, attempt, throwError)
import Effect (Effect)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Exception (error, throwException)
import Effect.Ref as Ref
import Gopurs.Cache.HeaderPreload as HeaderPreload
import Gopurs.Cache.Wire as Wire
import Gopurs.Cache.BindingCapsule as BindingCapsule
import Gopurs.CodeGen as CodeGen
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.Driver.Cache (FrozenFFI, Session)
import Gopurs.Driver.Cache as Cache
import Gopurs.Driver.Output as Output
import Gopurs.Driver.Prepare (PreparedModules)
import Gopurs.ExprContext (ModuleFunctions)
import Gopurs.GoSymbol (GoReference(..))
import Gopurs.Metrics as Metrics
import Gopurs.Preparation (mapPreparationJobs, runPreparationJobs)
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import Node.FS.Sync as Sync
import PureScript.Backend.Optimizer.Analysis (BackendAnalysis)
import PureScript.Backend.Optimizer.Convert (BackendImplementations, BackendModule, ExternLookup(..))
import PureScript.Backend.Optimizer.CoreFn (Ann, Ident(..), Module(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.Observation as Observation
import PureScript.Backend.Optimizer.Semantics (EvalRef(..), ExternImpl, InlineDirectiveMap)

type Part = { codec :: String, object :: { sha256 :: String, bytes :: Number } }
type Keys = { prepared :: String, optimizerEnvironment :: String, callerDemands :: String }
type Memo =
  { name :: String, prepared :: String, optimizerEnvironment :: String
  , callerDemands :: String, emitterEnvironment :: String, parts :: Array Part
  , metadata :: Maybe String, ffiDigest :: Maybe String
  , capsule :: Maybe Part
  }
type ModuleCache =
  { session :: Session
  , candidates :: Map String Memo
  , keys :: Map String Keys
  , metadata :: String
  , metadataDemands :: Maybe (Set String -> Effect Unit)
  , invalidations :: Ref.Ref (Map String String)
  , restored :: Ref.Ref Int
  , sameContext :: Boolean
  , reads :: Ref.Ref (Map String Observation.Reads)
  , restoredPBO :: Ref.Ref (Set String)
  , priorImplementations :: Ref.Ref (Map String (Maybe BackendImplementations))
  , implementationAnchors :: Map String String
  , priorCapsules :: Ref.Ref (Map String (Maybe BindingCapsule.Capsule))
  , capsuleHits :: Ref.Ref Int
  , capsuleCreated :: Ref.Ref Int
  , priorWorkers :: Ref.Ref (Map String (Maybe ModuleFunctions))
  , currentDigests :: Ref.Ref (Map (Qualified Ident) String)
  , priorDigests :: Ref.Ref (Map (Qualified Ident) String)
  , backends :: Ref.Ref (Map String BackendModule)
  , ffiDigests :: Map String String
  , priorFfiDigests :: Map String String
  , headerSafe :: Set String
  , schedulingReads :: Map String (Array String)
  , lazyContributions :: Boolean
  , priorHeaderDirectives :: Map String InlineDirectiveMap
  , priorPartTexts :: Map String (Map String String)
  , reused :: Ref.Ref (Map String Cache.ReusedMemo)
  }

-- Versioned backend-v2 envelope: keep one durable object for the body and the
-- reads which justify reuse when the conservative prepared prefix has changed.
type CachedBackend = { body :: BackendModule, reads :: Maybe Observation.Reads }

initialize :: Int -> Session -> PreparedModules -> Set String -> Ref.Ref (Map String BackendModule) -> Map String String -> Map String String -> Set String -> Aff ModuleCache
initialize jobs session prepared reusableModules backends ffiDigests priorFfiDigests headerCandidates = do
  candidates <- Metrics.measure "module index" \_ -> do
    text <- FS.readTextFile UTF8 (session.stage <> "/prior/index.json")
    case hush (jsonParser text) >>= hush <<< decodeJson of
      Just values -> pure (Map.fromFoldable (map (\memo -> Tuple memo.name memo) (values :: Array Memo)))
      Nothing -> throwError (error "Invalid private module cache index")
  seed <- Metrics.measure "optimizer seed" \_ -> liftEffect (Wire.hash <$> Wire.write (Tuple session.context prepared.directives))
  let
    modules = Array.fromFoldable prepared.modules
    sameContext = case Array.head modules of
      Just (Module m) -> map _.optimizerEnvironment (Map.lookup (unwrap m.name) candidates) == Just seed
      Nothing -> false
  -- Source/interface equality does not prove that specializations are equal.
  -- Only identical TAST or an explicit preparation proof can reuse these keys.
  deltaReusable <- Metrics.measure "module reusable candidates" \_ -> pure $
    if Set.size reusableModules == 0 then Map.empty else Map.filterWithKey (\name _ -> Set.member name reusableModules) candidates
  let reusable = if session.sameTast then candidates else deltaReusable
  fingerprints <- Metrics.measure "prepared module keys" \_ -> mapJobs jobs (preparedFingerprint reusable) modules
  let reused = Array.length (Array.filter (\(Module m) -> Map.member (unwrap m.name) reusable) modules)
  liftEffect $ Console.error $ "[gopurs] prepared fingerprints: computed=" <> show (Array.length modules - reused) <> ", reused=" <> show reused
  metadata <- Metrics.measure "global metadata key" \_ -> do
    let previousMetadata = do
          guard (prepared.metadataUnchanged && sameContext)
          Module first <- Array.head modules
          memo <- Map.lookup (unwrap first.name) candidates
          value <- memo.metadata
          guard (Array.all (\(Module m) -> (Map.lookup (unwrap m.name) candidates >>= _.metadata) == Just value) modules)
          pure value
    case previousMetadata of
      Just value -> do
        liftEffect $ Console.error "[gopurs] global metadata fingerprint: reused"
        pure value
      Nothing -> liftEffect (Wire.hash <$> Wire.write prepared.metadata)
  keyed <- Metrics.measure "module prefix keys" \_ -> pure $ foldl add { prefix: seed, keys: Map.empty } fingerprints
  changedPrepared <- Metrics.measure "module changed keys" \_ -> pure $ Set.fromFoldable (Array.mapMaybe
      (\(Tuple name keys) -> case Map.lookup name candidates of
        Just previous | previous.prepared == keys.prepared && previous.callerDemands == keys.callerDemands -> Nothing
        _ -> Just name)
      (Map.toUnfoldable keyed.keys :: Array (Tuple String Keys)))
  committedInputsMatch <- Metrics.measure "module committed inputs" \_ -> pure $ Array.all (\(Module m) -> case Map.lookup (unwrap m.name) candidates of
      Just previous -> previous.metadata == Just metadata && previous.ffiDigest == Just (ffiDigestFor ffiDigests (unwrap m.name))
      Nothing -> false) modules
    -- Complete optimizer keys prove the module and its canonical predecessor
    -- prefix independently of the read-owner closure. In particular, a read
    -- of a future module remains absent when only that future body changes.
    -- Global metadata and all FFI still gate direct emitted-source restoration.
  exactKeySafe <- Metrics.measure "module exact headers" \_ -> pure $
    if not sameContext || not committedInputsMatch then Set.empty
      else Set.fromFoldable (Array.mapMaybe
        (\(Tuple name keys) -> case Map.lookup name candidates of
          Just previous
            | previous.prepared == keys.prepared
                && previous.callerDemands == keys.callerDemands
                && previous.optimizerEnvironment == keys.optimizerEnvironment -> Just name
          _ -> Nothing)
        (Map.toUnfoldable keyed.keys :: Array (Tuple String Keys)))
  readProof <- if not sameContext || not committedInputsMatch || Set.isEmpty headerCandidates then pure { safe: Set.empty, owners: Map.empty } else do
    -- Exact contributions are unchanged even when their old read-owner proof
    -- mentions a changed future module. They cannot propagate invalidation.
    -- Build the remaining read graph only for owners still requiring proof.
    let names = Array.filter (\name -> not (Set.member name exactKeySafe)) (map (\(Module m) -> unwrap m.name) modules)
    previousReads <- Metrics.measure "reads load" \_ -> loadReads (session.stage <> "/prior") names
    let decoded = Array.length (Array.filter (\(Tuple _ value) -> isJust value) (Map.toUnfoldable previousReads :: Array _))
    liftEffect $ Console.error ("[gopurs] reads decoded=" <> show decoded <> "/" <> show (Array.length names))
    safe <- Metrics.measure "module read closure" \_ -> pure $ headerSafeFromReads changedPrepared previousReads (Set.fromFoldable names)
    selected <- Metrics.measure "module read intersection" \_ -> pure $ Set.intersection headerCandidates safe
    pure { safe: selected, owners: previousReads }
  headerSafe <- Metrics.measure "module header union" \_ -> pure $ Set.union exactKeySafe readProof.safe
  -- These committed read owners are scheduling hints, never reuse evidence.
  -- Header hits need no lookup; only uncertain modules benefit from waiting
  -- for their prior dependencies before validating the current canonical snapshot.
  let schedulingReads = Map.mapMaybeWithKey (\name owners ->
        if Set.member name headerSafe then Nothing else owners) readProof.owners
  invalidations <- liftEffect (Ref.new Map.empty)
  restored <- liftEffect (Ref.new 0)
  reads <- liftEffect (Ref.new Map.empty)
  restoredPBO <- liftEffect (Ref.new Set.empty)
  priorImplementations <- liftEffect (Ref.new Map.empty)
  priorCapsules <- liftEffect (Ref.new Map.empty)
  capsuleHits <- liftEffect (Ref.new 0)
  capsuleCreated <- liftEffect (Ref.new 0)
  let implementationAnchors = Map.mapMaybe (\memo -> do
        part <- Array.index memo.parts 1
        guard (part.codec == "pbo/implementations-v1")
        pure part.object.sha256) candidates
  priorWorkers <- liftEffect (Ref.new Map.empty)
  currentDigests <- liftEffect (Ref.new Map.empty)
  priorDigests <- liftEffect (Ref.new Map.empty)
  reusedMemos <- liftEffect (Ref.new Map.empty)
  pure { session, candidates, keys: keyed.keys, metadata, metadataDemands: prepared.metadataDemands, invalidations, restored, sameContext, reads, restoredPBO, priorImplementations, implementationAnchors, priorCapsules, capsuleHits, capsuleCreated, priorWorkers, currentDigests, priorDigests, backends, ffiDigests, priorFfiDigests, headerSafe, schedulingReads, lazyContributions: false, priorHeaderDirectives: Map.empty, priorPartTexts: Map.empty, reused: reusedMemos }

  where
  add acc (Tuple name preparedKey) =
    let
      keys =
        { prepared: preparedKey
        , optimizerEnvironment: acc.prefix
        , callerDemands: Wire.hash ("specialized-module-v1/" <> preparedKey)
        }
    in
      { prefix: Wire.hash (acc.prefix <> "/" <> preparedKey)
      , keys: Map.insert name keys acc.keys
      }

mapJobs :: forall a b. Int -> (a -> Aff b) -> Array a -> Aff (Array b)
mapJobs = mapPreparationJobs

preparedFingerprint :: Map String Memo -> Module Ann -> Aff (Tuple String String)
-- Hash the actual specialized module, including embedded static arguments.
preparedFingerprint reusable mod@(Module m) = case Map.lookup (unwrap m.name) reusable of
  Just previous -> pure (Tuple (unwrap m.name) previous.prepared)
  Nothing -> Tuple (unwrap m.name) <$> liftEffect (Wire.hash <$> Wire.write mod)

keysFor :: ModuleCache -> String -> Aff Keys
keysFor cache name = case Map.lookup name cache.keys of
  Just keys -> pure keys
  Nothing -> throwError (error ("Missing prepared cache key: " <> name))

prior :: ModuleCache -> String -> String
prior cache name = cache.session.stage <> "/prior/" <> name <> "/"

target :: ModuleCache -> String -> String
target cache name = cache.session.stage <> "/modules/" <> name <> "/"

-- Private staging files can share the cached inode. Publication only reads
-- them; a missing hard-link support (cross-device) falls back to a copy.
linkOrCopy :: String -> String -> Aff Unit
linkOrCopy source destination = do
  linked <- attempt (FS.link source destination)
  case linked of
    Right _ -> pure unit
    Left _ -> FS.copyFile source destination

readPart :: forall a. Wire.Codec a => String -> Aff (Maybe a)
readPart path = do
  text <- hush <$> attempt (FS.readTextFile UTF8 path)
  pure (text >>= Wire.read)

-- When every complete optimizer key matches, no conversion reads the saved
-- implementations/directives. Validate only the business data consumed by the
-- emitter; the verified contribution bytes are copied unchanged at publication.
-- Any missing key or invalid body/worker falls back to the ordinary PBO builder,
-- which decodes contributions before making them visible to another module.
lookupAll :: Int -> ModuleCache -> List (Module Ann) -> Aff (Maybe (Array BackendModule))
lookupAll jobs cache modules = do
  let ordered = Array.fromFoldable modules
  let matches (Module m) = case Map.lookup (unwrap m.name) cache.keys of
        Just keys -> invalidationReason cache (unwrap m.name) keys == ""
        Nothing -> false
  if Array.all matches ordered then sequence <$> mapJobs jobs (lookupHeader cache) ordered
  else pure Nothing

-- Decode only independently proven headers. The builder restores these at its
-- canonical publication front; this private map does not expose contributions.
-- An invalid business codec remains an ordinary miss, including the existing
-- placeholder check and full fallback when that module is attempted.
lookupHeaders :: Int -> ModuleCache -> List (Module Ann) -> Aff (Map String BackendModule)
lookupHeaders jobs cache modules
  | not cache.lazyContributions = pure Map.empty
  | otherwise = do
      let selected = Array.filter (\(Module m) -> Set.member (unwrap m.name) cache.headerSafe) (Array.fromFoldable modules)
      decoded <- mapJobs jobs (lookupHeader cache) selected
      pure $ Map.fromFoldable $ Array.mapMaybe identity $ Array.zipWith
        (\(Module m) backend -> Tuple (unwrap m.name) <$> backend) selected decoded

invalidationReason :: ModuleCache -> String -> Keys -> String
invalidationReason cache name keys = case Map.lookup name cache.candidates of
  Nothing -> "no-entry"
  Just old | old.prepared /= keys.prepared || old.callerDemands /= keys.callerDemands -> "prepared-or-callers"
  Just old | old.optimizerEnvironment /= keys.optimizerEnvironment -> "optimizer-prefix"
  _ -> ""

preparedMatches :: ModuleCache -> String -> Keys -> Boolean
preparedMatches cache name keys = case Map.lookup name cache.candidates of
  Just old -> old.prepared == keys.prepared && old.callerDemands == keys.callerDemands
  Nothing -> false

ffiDigestFor :: Map String String -> String -> String
ffiDigestFor digests name = fromMaybe "no-foreign-v1" (Map.lookup name digests)

ffiMatches :: ModuleCache -> String -> Boolean
ffiMatches cache name =
  Map.lookup name cache.ffiDigests == Map.lookup name cache.priorFfiDigests
    && case Map.lookup name cache.candidates of
      Just previous -> previous.ffiDigest == Just (ffiDigestFor cache.ffiDigests name)
      Nothing -> false

-- A cached module only needs its implementations, directives and a header.
-- The whole back-end body is not consumed by the emitter: files and workers
-- are restored by hash from the private cache. This avoids decoding megabytes
-- per unchanged module in a one-shot process.
lookupHeader :: ModuleCache -> Module Ann -> Aff (Maybe BackendModule)
lookupHeader cache (Module m) = do
  let name = unwrap m.name
  keys <- keysFor cache name
  let sameFfi = ffiMatches cache name
  if not sameFfi || not (preparedMatches cache name keys) || not (Set.member name cache.headerSafe) then pure Nothing
  else do
    implementations <- if cache.lazyContributions then pure (Just Map.empty) else readImplementations cache name
    directives <- readHeaderDirectives cache name
    workers <- if cache.lazyContributions then pure (Just Map.empty) else readWorkers cache name
    case implementations, directives, workers of
      Just impls, Just dirs, Just _ -> do
        let owned =
              Array.all (\(Qualified mn _) -> mn == Just m.name) (Map.keys impls # Array.fromFoldable)
                && Array.all (case _ of
                  EvalExtern (Qualified mn _) -> mn == Just m.name
                  _ -> false) (Map.keys dirs # Array.fromFoldable)
        if not owned then pure Nothing
        else do
          let
            stub =
              { name: m.name
              , comments: []
              , imports: Set.empty
              , dataTypes: Map.empty
              , bindings: []
              , exports: Set.empty
              , reExports: Set.empty
              , dataDecls: []
              , classDecls: []
              , foreign: m.foreign
              , implementations: impls
              , directives: dirs
              }
          liftEffect do
            Ref.modify_ (Set.insert name) cache.restoredPBO
            Ref.modify_ (Map.delete name) cache.invalidations
          pure (Just stub)
      _, _, _ -> pure Nothing

readBodyEntry :: ModuleCache -> Module Ann -> Aff (Maybe CachedBackend)
readBodyEntry cache (Module m) = do
  let name = unwrap m.name
  ensureParts cache name [ "0" ]
  body <- case Map.lookup name cache.candidates >>= \memo -> Array.head memo.parts of
    Just part | part.codec == "pbo/backend-v2" -> readPart (prior cache name <> "0.json")
    _ -> pure Nothing
  workers <- readWorkers cache (unwrap m.name)
  pure do
    entry <- body
    _ <- workers
    let mod = entry.body
    guard (mod.name == m.name && Map.isEmpty mod.implementations && Map.isEmpty mod.directives)
    pure entry

lookupBackend :: ModuleCache -> Module Ann -> Aff (Maybe BackendModule)
lookupBackend cache mod = lookupBackendIn cache mod Nothing

lookupBackendWithReads :: ModuleCache -> Module Ann -> Observation.View -> Aff (Maybe BackendModule)
lookupBackendWithReads cache mod view = lookupBackendIn cache mod (Just view)

captureReads :: ModuleCache -> Module Ann -> Observation.Reads -> Aff Unit
captureReads cache (Module m) reads = liftEffect (Ref.modify_ (Map.insert (unwrap m.name) reads) cache.reads)

lookupBackendIn :: ModuleCache -> Module Ann -> Maybe Observation.View -> Aff (Maybe BackendModule)
lookupBackendIn cache coreFnModule@(Module m) view = do
  let name = unwrap m.name
  keys <- keysFor cache name
  let reason = invalidationReason cache name keys
  let sameFfi = ffiMatches cache name
  -- The resident map may contain a header stub from an earlier fast restore.
  -- Without today's header proof, decode the complete cached body instead.
  if sameFfi && Set.member name cache.headerSafe && (reason == "" || reason == "optimizer-prefix") then do
    previous <- liftEffect (Map.lookup name <$> Ref.read cache.backends)
    case previous of
      Just backend -> do
        liftEffect do
          Ref.modify_ (Set.insert name) cache.restoredPBO
          Ref.modify_ (Map.delete name) cache.invalidations
        pure (Just backend)
      Nothing -> lookupHeader cache coreFnModule
  else lookupBackendFromDisk cache coreFnModule view

lookupBackendFromDisk :: ModuleCache -> Module Ann -> Maybe Observation.View -> Aff (Maybe BackendModule)
lookupBackendFromDisk cache coreFnModule@(Module m) view = do
  let name = unwrap m.name
  keys <- keysFor cache name
  let reason = invalidationReason cache name keys
  if reason /= "" && not (reason == "optimizer-prefix" && cache.sameContext && isJust view) then do
    liftEffect (Ref.modify_ (Map.insert name reason) cache.invalidations)
    pure Nothing
  else do
    entry <- readBodyEntry cache coreFnModule
    accepted <- case entry of
      Nothing -> pure false
      Just _ | reason == "" -> pure true
      Just saved -> case view, saved.reads of
        Just inputs, Just reads -> validateReads cache inputs reads
        _, _ -> pure false
    implementations <- if accepted then readImplementations cache name else pure Nothing
    directives <- readHeaderDirectives cache name
    let decoded = do
          guard accepted
          mod <- map _.body entry
          impls <- implementations
          dirs <- directives
          guard (Array.all (\(Qualified mn _) -> mn == Just m.name) (Map.keys impls # Array.fromFoldable))
          guard (Array.all (case _ of
            EvalExtern (Qualified mn _) -> mn == Just m.name
            _ -> false) (Map.keys dirs # Array.fromFoldable))
          pure (mod { implementations = impls, directives = dirs })
    case decoded of
      Nothing -> liftEffect (Ref.modify_ (Map.insert name (if isJust entry && not accepted then "optimizer-reads" else "business-codec")) cache.invalidations)
      Just _ -> liftEffect do
        Ref.modify_ (Set.insert name) cache.restoredPBO
        Ref.modify_ (Map.delete name) cache.invalidations
    pure decoded

readImplementations :: ModuleCache -> String -> Aff (Maybe BackendImplementations)
readImplementations cache name | cache.lazyContributions && Map.member name cache.candidates =
  Just <$> liftEffect (readImplementationsSync cache name)
readImplementations cache name = do
  previous <- liftEffect (Ref.read cache.priorImplementations)
  case Map.lookup name previous of
    Just value -> pure value
    Nothing -> do
      value <- if Map.member name cache.candidates then readPart (prior cache name <> "1.json") else pure Nothing
      liftEffect (Ref.modify_ (Map.insert name value) cache.priorImplementations)
      pure value

-- Read verified small objects directly from the pinned private manifest.
-- Their decoding and canonical visibility still happen at the usual points.
loadHeaderTexts :: ModuleCache -> Aff ModuleCache
loadHeaderTexts = loadHeaderTextsWithPreload Nothing

loadHeaderTextsWithPreload :: Maybe HeaderPreload.Reservoir -> ModuleCache -> Aff ModuleCache
loadHeaderTextsWithPreload pending cache
  | not cache.lazyContributions || Set.isEmpty cache.headerSafe = pure cache
  | Just headers <- pending >>= HeaderPreload.select cache.session cache.headerSafe =
      pure (cache { priorHeaderDirectives = headers })
  | otherwise = do
      contributions <- Cache.readParts cache.session (Set.toUnfoldable cache.headerSafe) [ "2" ]
      texts <- foldM add Map.empty contributions
      unless (Array.length contributions == Set.size cache.headerSafe) $
        throwError (error "incremental cache: incomplete header contributions")
      pure (cache { priorPartTexts = texts })
  where
  add texts entry = do
    let parts = fromMaybe Map.empty (Map.lookup entry.name texts)
    unless (Set.member entry.name cache.headerSafe && entry.part == "2" && not (Map.member entry.part parts)) $
      throwError (error "incremental cache: unexpected header contribution")
    pure (Map.insert entry.name (Map.insert entry.part entry.text parts) texts)

readHeaderDirectives :: ModuleCache -> String -> Aff (Maybe InlineDirectiveMap)
readHeaderDirectives cache name = case Map.lookup name cache.priorHeaderDirectives of
  Just value -> pure (Just value)
  Nothing -> readCachedPart cache name "2"

readCachedPart :: forall a. Wire.Codec a => ModuleCache -> String -> String -> Aff (Maybe a)
readCachedPart cache name part = case Map.lookup name cache.priorPartTexts >>= Map.lookup part of
  Just text -> pure (Wire.read text)
  Nothing -> do
    ensureParts cache name [ part ]
    readPart (prior cache name <> part <> ".json")

ensureParts :: ModuleCache -> String -> Array String -> Aff Unit
ensureParts cache name parts = when (cache.lazyContributions && Map.member name cache.candidates) $
  Cache.materialize cache.session [ name ] parts

-- Observations are hints for private decoding, never evidence of visibility.
-- Only prior, unchanged owners can help this confirmed conversion miss.
implementationOwners :: ModuleCache -> Map String Int -> String -> Array String
implementationOwners cache ranks reader = case Map.lookup reader ranks of
  Nothing -> []
  Just rank -> Set.toUnfoldable $ Set.fromFoldable $ Array.filter
    (\owner -> Set.member owner cache.headerSafe && case Map.lookup owner ranks of
      Just ownerRank -> ownerRank < rank
      Nothing -> false)
    (fromMaybe [] (Map.lookup reader cache.schedulingReads))

decodeImplementations :: String -> String -> Maybe BackendImplementations
decodeImplementations name text = do
  value <- Wire.read text
  guard (Array.all (\(Qualified owner _) -> map unwrap owner == Just name) (Map.keys value # Array.fromFoldable))
  pure value

-- Run only after a conversion miss. Validate and read every immutable object
-- before the pure parallel tasks, then join and validate the entire batch
-- before adding anything to the private memo. The builder still controls rank
-- and finality; this never publishes an optimizer contribution.
warmImplementations :: Int -> ModuleCache -> Map String Int -> String -> Aff Unit
warmImplementations = warmImplementationOwners Set.empty

-- A capsule is a private hint for this confirmed miss. Uncovered new reads
-- still reach the complete owner loader after the builder's visibility guards.
warmImplementationInputs :: Int -> ModuleCache -> Map String Int -> String -> Aff Unit
warmImplementationInputs jobs cache ranks reader = do
  capsule <- loadBindingCapsule cache reader
  warmImplementationOwners (maybe Set.empty (Map.keys <<< _.anchors) capsule) jobs cache ranks reader

loadBindingCapsule :: ModuleCache -> String -> Aff (Maybe BindingCapsule.Capsule)
loadBindingCapsule cache reader
  | not cache.lazyContributions || not cache.sameContext = pure Nothing
  | otherwise = do
      previous <- liftEffect (Map.lookup reader <$> Ref.read cache.priorCapsules)
      case previous of
        Just value -> pure value
        Nothing -> do
          value <- case Map.lookup reader cache.candidates >>= _.capsule of
            Just part | part.codec == BindingCapsule.codec -> Metrics.measure "PBO binding capsule" \_ -> do
              text <- Cache.readCapsule cache.session reader
              pure (text >>= BindingCapsule.decode cache.session.context reader cache.implementationAnchors)
            _ -> pure Nothing
          liftEffect (Ref.modify_ (Map.insert reader value) cache.priorCapsules)
          pure value

warmImplementationOwners :: Set String -> Int -> ModuleCache -> Map String Int -> String -> Aff Unit
warmImplementationOwners covered jobs cache ranks reader
  | not cache.lazyContributions = pure unit
  | otherwise = do
      previous <- liftEffect (Ref.read cache.priorImplementations)
      let owners = Array.filter (\owner -> not (Set.member owner covered) && case Map.lookup owner previous of
            Just (Just _) -> false
            _ -> true) (implementationOwners cache ranks reader)
      unless (Array.null owners) $ Metrics.measure "PBO implementation warm" \_ -> do
        Cache.materialize cache.session owners [ "1" ]
        encoded <- attempt $ traverse (\name -> Tuple name <$> FS.readTextFile UTF8 (prior cache name <> "1.json")) owners
        texts <- case encoded of
          Left _ -> throwError (error "incremental cache: unreadable warm implementation contribution")
          Right texts -> pure texts
        decoded <- runPreparationJobs jobs (map (\(Tuple name text) _ -> Tuple name (decodeImplementations name text)) texts)
        complete <- traverse (\(Tuple name value) -> case value of
          Nothing -> throwError (error ("incremental cache: invalid implementation contribution: " <> name))
          Just implementations -> pure (Tuple name (Just implementations))) decoded
        let additions = Map.fromFoldable complete
        liftEffect (Ref.modify_ (\current -> Map.union current additions) cache.priorImplementations)
        liftEffect $ Console.error ("[gopurs] warmed implementation owners: " <> show (Array.length owners))

-- The builder calls this only for a finalized reused owner. Complete maps win;
-- a capsule miss never stands for absence. Fresh contributions bypass it.
readCachedBinding :: ModuleCache -> String -> String -> String -> Effect (Maybe (Maybe (Tuple BackendAnalysis ExternImpl)))
readCachedBinding cache reader owner ident = do
  let key = Qualified (Just (ModuleName owner)) (Ident ident)
  previous <- Ref.read cache.priorImplementations
  case Map.lookup owner previous of
    Just (Just implementations) -> pure (Just (Map.lookup key implementations))
    _ | Set.member owner cache.headerSafe -> do
      capsule <- map (\entries -> Map.lookup reader entries >>= identity) (Ref.read cache.priorCapsules)
      let covered = capsule >>= BindingCapsule.lookup key
      when (isJust covered) (Ref.modify_ (_ + 1) cache.capsuleHits)
      pure covered
    _ -> pure Nothing

publishBindingCapsule :: ModuleCache -> String -> Observation.Reads -> Aff Unit
publishBindingCapsule cache reader reads
  | not cache.lazyContributions || not cache.sameContext || not (Map.member reader cache.candidates) = pure unit
  | otherwise = do
      complete <- liftEffect (Ref.read cache.priorImplementations)
      capsule <- liftEffect (map (\entries -> Map.lookup reader entries >>= identity) (Ref.read cache.priorCapsules))
      let entries = BindingCapsule.select cache.headerSafe cache.implementationAnchors complete capsule reads
      encoded <- liftEffect (BindingCapsule.encode cache.session.context reader cache.implementationAnchors entries)
      for_ encoded \text -> do
        FS.writeTextFile UTF8 (target cache reader <> "capsule.json") text
        liftEffect (Ref.modify_ (_ + 1) cache.capsuleCreated)

-- Called only after the builder's rank/finality checks. These immutable bytes
-- belong to this session's committed manifest; never consult a public mirror.
readImplementationsSync :: ModuleCache -> String -> Effect BackendImplementations
readImplementationsSync cache name = do
  previous <- Ref.read cache.priorImplementations
  case Map.lookup name previous of
    Just (Just value) -> pure value
    _ -> do
      Cache.materializeSync cache.session [ name ] [ "1" ]
      text <- Sync.readTextFile UTF8 (prior cache name <> "1.json")
      case decodeImplementations name text of
        Just value -> do
          Ref.modify_ (Map.insert name (Just value)) cache.priorImplementations
          pure value
        _ -> throwException (error ("incremental cache: invalid implementation contribution: " <> name))

implementationDigest :: Ref.Ref (Map (Qualified Ident) String) -> Qualified Ident -> Tuple BackendAnalysis ExternImpl -> Aff String
implementationDigest memo name value = liftEffect do
  entries <- Ref.read memo
  case Map.lookup name entries of
    Just digest -> pure digest
    Nothing -> do
      digest <- Wire.hash <$> Wire.write value
      Ref.modify_ (Map.insert name digest) memo
      pure digest

validateReads :: ModuleCache -> Observation.View -> Observation.Reads -> Aff Boolean
validateReads cache view reads = do
  implementations <- foldM (\ok entry -> if ok then implementation entry else pure false) true
    (Map.toUnfoldable reads.implementations :: Array _)
  let directives = Array.all (\(Tuple name expected) -> view.directive name == expected)
        (Map.toUnfoldable reads.directives :: Array _)
  -- Force the pure directive lookups before reading their pending dependency
  -- state. A pending predecessor is never equivalent to an absent declaration.
  if not implementations || not directives then pure false
  else not <$> liftEffect view.pending
  where
  implementation (Tuple name@(Qualified (Just owner) (Ident ident)) present) = do
    restored <- liftEffect (Ref.read cache.restoredPBO)
    -- A saved positive read belongs to the same immutable contribution when
    -- its owner was restored and is visible in this attempt's snapshot. Avoid
    -- decoding that contribution just to prove its value unchanged. Absences
    -- still need a lookup: the old read may have hidden a future owner.
    if present && Set.member (unwrap owner) restored && view.availableOwner (unwrap owner) then pure true
    else case view.implementation (unwrap owner) ident of
      ExternPending -> pure false
      ExternMissing -> pure (not present)
      ExternFound _ | not present -> pure false
      ExternFound current -> do
        if Set.member (unwrap owner) restored then pure true
        else do
          previous <- readImplementations cache (unwrap owner)
          case previous >>= Map.lookup name of
            Nothing -> pure false
            Just expected -> do
              before <- implementationDigest cache.priorDigests name expected
              after <- implementationDigest cache.currentDigests name current
              pure (before == after)
  implementation _ = pure false

validateWorkers :: String -> ModuleFunctions -> Maybe ModuleFunctions
validateWorkers name functions = do
  guard (Array.all (isJust <<< String.stripPrefix (Pattern (name <> "."))) (Array.fromFoldable (Map.keys functions)))
  guard (Array.all (\info -> info.arity >= 0 && case info.reference of
    SymbolReference symbol -> symbol.owner /= "" && symbol.name /= ""
    LocalReference _ -> false) (Array.fromFoldable (Map.values functions)))
  pure functions

-- Decode only cached owners already visible in the captured emission snapshot.
-- Populate the private memo after the entire response is validated; visibility
-- still belongs to resolveWorkerSnapshot, including current and future absences.
warmWorkerSnapshot :: Int -> ModuleCache -> Map String (Maybe ModuleFunctions) -> Array BackendModule -> Aff Unit
warmWorkerSnapshot jobs cache snapshot backends
  | not cache.lazyContributions = pure unit
  | otherwise = do
      previous <- liftEffect (Ref.read cache.priorWorkers)
      let requested = Set.filter (\owner -> case Map.lookup owner snapshot of
            Just Nothing -> not (Map.member owner previous)
            _ -> false) (foldl (\owners backend -> Set.union owners (CodeGen.workerOwners backend)) Set.empty backends)
      unless (Set.isEmpty requested) $ do
        let owners = Set.toUnfoldable requested :: Array String
        for_ owners \name -> case Map.lookup name cache.candidates >>= \memo -> Array.index memo.parts 3 of
          Just part | part.codec == "gopurs/workers-v2" -> pure unit
          _ -> throwError (error ("incremental cache: unreadable workers for " <> name))
        contributions <- Cache.readParts cache.session owners [ "3" ]
        texts <- foldM (\values entry -> do
          unless (entry.part == "3" && Set.member entry.name requested && not (Map.member entry.name values)) $
            throwError (error "incremental cache: unexpected worker contribution")
          pure (Map.insert entry.name entry.text values)) Map.empty contributions
        unless (Map.size texts == Set.size requested) $
          throwError (error "incremental cache: incomplete worker contributions")
        decoded <- runPreparationJobs jobs (map (\(Tuple name text) _ -> Tuple name (Wire.read text >>= validateWorkers name)) (Map.toUnfoldable texts :: Array (Tuple String String)))
        complete <- traverse (\(Tuple name value) -> case value of
          Nothing -> throwError (error ("incremental cache: unreadable workers for " <> name))
          Just functions -> pure (Tuple name (Just functions))) decoded
        let additions = Map.fromFoldable complete
        liftEffect (Ref.modify_ (\current -> Map.union current additions) cache.priorWorkers)

readWorkers :: ModuleCache -> String -> Aff (Maybe ModuleFunctions)
readWorkers cache name = do
  previous <- liftEffect (Ref.read cache.priorWorkers)
  case Map.lookup name previous of
    Just value -> pure value
    Nothing -> do
      workers <- case Map.lookup name cache.candidates >>= \memo -> Array.index memo.parts 3 of
        Just part | part.codec == "gopurs/workers-v2" -> readCachedPart cache name "3"
        _ -> pure Nothing
      let value = workers >>= validateWorkers name
      liftEffect (Ref.modify_ (Map.insert name value) cache.priorWorkers)
      pure value

readWorkerContribution :: ModuleCache -> String -> Aff ModuleFunctions
readWorkerContribution cache name = do
  functions <- readWorkers cache name
  case functions of
    Nothing -> throwError (error ("incremental cache: unreadable workers for " <> name))
    Just value -> pure value

-- Each owner enters this table only after its emission batch completes. A
-- missing owner therefore remains absent even when its cache object exists.
-- Nothing is an immutable, proved cached contribution; Just is today's result.
resolveWorkerSnapshot
  :: (String -> Aff ModuleFunctions)
  -> Map String (Maybe ModuleFunctions)
  -> BackendModule
  -> Aff ModuleFunctions
resolveWorkerSnapshot load snapshot backend =
  foldM resolve Map.empty (Set.toUnfoldable (CodeGen.workerOwners backend) :: Array String)
  where
  resolve functions owner = case Map.lookup owner snapshot of
    Nothing -> pure functions
    Just (Just current) -> pure (Map.union current functions)
    Just Nothing -> (\previous -> Map.union previous functions) <$> load owner

emit
  :: ModuleCache -> Map String FrozenFFI -> CodegenMetadata -> String -> Boolean
  -> Module Ann -> BackendModule -> Aff { functions :: ModuleFunctions, workers :: String }
emit cache ffi metadata environment optimizedCached coreFnModule backend = do
  result <- emitWithWorkers Nothing cache ffi metadata environment optimizedCached coreFnModule backend
  pure { functions: result.functions, workers: result.workers }

-- The resolver captures the worker table before this emission batch. It must
-- never expose signatures from the current batch or a later canonical rank.
emitWithWorkerSnapshot
  :: (BackendModule -> Aff ModuleFunctions)
  -> ModuleCache -> Map String FrozenFFI -> CodegenMetadata -> String -> Boolean
  -> Module Ann -> BackendModule -> Aff { functions :: ModuleFunctions, workers :: String, deferred :: Boolean }
emitWithWorkerSnapshot resolve = emitWithWorkers (Just resolve)

emitWithWorkers
  :: Maybe (BackendModule -> Aff ModuleFunctions)
  -> ModuleCache -> Map String FrozenFFI -> CodegenMetadata -> String -> Boolean
  -> Module Ann -> BackendModule -> Aff { functions :: ModuleFunctions, workers :: String, deferred :: Boolean }
emitWithWorkers resolver cache ffi metadata _ optimizedCached coreFnModule backend = do
  let name = unwrap backend.name
  let directory = target cache name
  keys <- keysFor cache name
  let ffiDigest = ffiDigestFor cache.ffiDigests name
  if optimizedCached && preparedMatches cache name keys && Set.member name cache.headerSafe then do
    old <- case Map.lookup name cache.candidates of
      Nothing -> throwError (error "Lost module cache candidate")
      Just candidate -> pure candidate
    part <- case Array.index old.parts 3 of
      Nothing -> throwError (error ("Missing cached workers for " <> name))
      Just value | value.codec == "gopurs/workers-v2" -> pure value
      _ -> throwError (error ("incremental cache: invalid worker codec for " <> name))
    functions <- case resolver of
      Just _ -> pure (Just Map.empty)
      Nothing -> readWorkers cache name
    case functions of
      Nothing -> throwError (error ("incremental cache: unreadable workers for " <> name))
      Just value -> do
        let memo = { name, prepared: keys.prepared, optimizerEnvironment: keys.optimizerEnvironment, callerDemands: keys.callerDemands, emitterEnvironment: old.emitterEnvironment, metadata: cache.metadata, ffiDigest }
        if cache.lazyContributions then liftEffect (Ref.modify_ (Map.insert name memo) cache.reused)
        else do
          FS.mkdir directory
          for_ [ 0, 1, 2, 3 ] \i -> linkOrCopy (prior cache name <> show i <> ".json") (directory <> show i <> ".json")
          void $ attempt (linkOrCopy (prior cache name <> "reads.json") (directory <> "reads.json"))
          let goName = String.replaceAll (Pattern ".") (Replacement "_") name
          linkOrCopy (prior cache name <> "module.go") (cache.session.output <> "/purescript/" <> goName <> ".go")
          unless (Map.isEmpty backend.foreign) $
            linkOrCopy (prior cache name <> "ffi.go") (cache.session.output <> "/purescript/" <> goName <> "_ffi.go")
          FS.writeTextFile UTF8 (directory <> "memo.json") (stringify (encodeJson memo))
        liftEffect do
          Ref.modify_ (_ + 1) cache.restored
          Ref.modify_ (Map.insert name backend) cache.backends
        pure { functions: value, workers: part.object.sha256, deferred: isJust resolver }
  else do
    FS.mkdir directory
    -- An unsafe cached module keeps the previous emitter-environment proof:
    -- copy its PBO parts, then compare the full environment below.
    digests <- if optimizedCached then do
      ensureParts cache name [ "0", "1", "2", "reads" ]
      old <- case Map.lookup name cache.candidates of
        Nothing -> throwError (error "Lost module cache candidate")
        Just candidate -> pure candidate
      for_ [ 0, 1, 2 ] \i -> linkOrCopy (prior cache name <> show i <> ".json") (directory <> show i <> ".json")
      void $ attempt (linkOrCopy (prior cache name <> "reads.json") (directory <> "reads.json"))
      pure (map (_.object.sha256) (Array.take 3 old.parts))
    else do
      -- Keep the three PBO parts distinct; the backend part has empty contribution
      -- maps. A hit copies their verified private bytes instead of re-encoding them.
      observations <- liftEffect (Ref.read cache.reads)
      let
        emptyReads = { implementations: Map.empty, directives: Map.empty } :: Observation.Reads
        owned = fromMaybe emptyReads (Map.lookup name observations)
        ownerName (Qualified owner _) = case owner of
          Just moduleName -> unwrap moduleName
          Nothing -> ""
        owners =
          map ownerName (Array.fromFoldable (Map.keys owned.implementations) <> Array.fromFoldable (Map.keys owned.directives))
        readsText = if Array.null owners then "" else String.joinWith "\n" owners <> "\n"
      FS.writeTextFile UTF8 (directory <> "reads.json") readsText
      publishBindingCapsule cache name owned
      parts <- liftEffect $ sequence
            [ Wire.write ({ body: backend { implementations = Map.empty, directives = Map.empty }, reads: Map.lookup name observations } :: CachedBackend)
            , Wire.write backend.implementations
            , Wire.write backend.directives
            ]
      liftEffect (Ref.modify_ (Map.delete name) cache.reads)
      traverse (\(Tuple i text) -> FS.writeTextFile UTF8 (directory <> show i <> ".json") text $> Wire.hash text)
        (Array.mapWithIndex Tuple parts)
    available <- case resolver of
      Just resolve -> resolve backend
      Nothing -> pure metadata.globalFunctions
    let
      environment proof workers = liftEffect do
        workerKey <- Wire.hash <$> Wire.write workers
        Wire.hash <$> Wire.write
          { context: cache.session.context, metadata: cache.metadata, workers: workerKey, backend: digests
          , ffi: fromMaybe "no-foreign" (map _.digest (Map.lookup name ffi)), workerProof: proof
          }
      cached environmentKey = case Map.lookup name cache.candidates of
        Just old | old.emitterEnvironment == environmentKey -> case Array.index old.parts 3 of
          Just part -> map (\functions -> { functions, workers: part.object.sha256 }) <$> readWorkers cache name
          Nothing -> pure Nothing
        _ -> pure Nothing
      restore result = do
        ensureParts cache name [ "module", "ffi", "3" ]
        let goName = String.replaceAll (Pattern ".") (Replacement "_") name
        linkOrCopy (prior cache name <> "module.go") (cache.session.output <> "/purescript/" <> goName <> ".go")
        unless (Map.isEmpty backend.foreign) $
          linkOrCopy (prior cache name <> "ffi.go") (cache.session.output <> "/purescript/" <> goName <> "_ffi.go")
        linkOrCopy (prior cache name <> "3.json") (directory <> "3.json")
        liftEffect (Ref.modify_ (_ + 1) cache.restored)
        pure result
      inputWorkers = CodeGen.workerInputs available backend
    -- A key in this domain is committed only after checking containment below.
    -- Equal backend/metadata/compiler inputs reproduce the same transformations;
    -- all their imported signatures are then covered by this conservative view.
    -- PrepareModule's transformations do not consume imported worker signatures.
    inputEnvironment <- environment "input-workers-contained-v1" inputWorkers
    early <- cached inputEnvironment
    Tuple emitterEnvironment result <- case early of
      Just previous -> Tuple inputEnvironment <$> restore previous
      Nothing -> do
        let transformed = CodeGen.prepareModule (metadata { globalFunctions = available }) backend
        liftEffect $ for_ cache.metadataDemands \record ->
          record (Set.union (CodeGen.globalTypeDemands backend) (CodeGen.globalTypeDemands transformed.module))
        transformedWorkers <- case resolver of
          Just resolve -> CodeGen.workerInputs <$> resolve transformed.module <*> pure transformed.module
          Nothing -> pure transformed.workers
        let preparedModule = transformed { workers = transformedWorkers }
        let contained = Array.all (\key -> Map.member key inputWorkers) (Array.fromFoldable (Map.keys preparedModule.workers))
        environmentKey <- if contained then pure inputEnvironment
          else environment "prepared-workers-v1" preparedModule.workers
        -- A transformation may introduce an external reference. In that case the
        -- existing post-transformation proof remains available, after preparation.
        late <- if contained then pure Nothing else cached environmentKey
        emitted <- case late of
          Just previous -> restore previous
          Nothing -> do
            functions <- Output.emitPreparedModule cache.session.output ffi metadata coreFnModule backend preparedModule
            text <- liftEffect (Wire.write functions)
            FS.writeTextFile UTF8 (directory <> "3.json") text
            pure { functions, workers: Wire.hash text }
        pure (Tuple environmentKey emitted)
    let memo = { name, prepared: keys.prepared, optimizerEnvironment: keys.optimizerEnvironment, callerDemands: keys.callerDemands, emitterEnvironment, metadata: cache.metadata, ffiDigest, parts: [] :: Array Part }
    FS.writeTextFile UTF8 (directory <> "memo.json") (stringify (encodeJson memo))
    liftEffect (Ref.modify_ (Map.insert name backend) cache.backends)
    pure { functions: result.functions, workers: result.workers, deferred: false }

loadReads :: String -> Array String -> Aff (Map String (Maybe (Array String)))
-- Read only the hash-verified objects frozen from the committed module
-- manifest. The public reads-v1 mirror may belong to a different publication.
loadReads directory names = do
  let grouped = directory <> "/owners.json"
  content <- attempt (FS.readTextFile UTF8 grouped)
  case content of
    Left _ -> do
      exists <- liftEffect (Sync.exists grouped)
      if exists then throwError (error "incremental cache: unreadable read-owner index")
      else Map.fromFoldable <$> traverse read names
    Right text -> case hush (jsonParser text) >>= hush <<< decodeJson of
      Nothing -> throwError (error "incremental cache: invalid read-owner index")
      Just entries -> do
        decoded <- foldM add Map.empty (entries :: Array { name :: String, text :: Json })
        pure (Map.fromFoldable (map (\name -> Tuple name (fromMaybe Nothing (Map.lookup name decoded))) names))
  where
  owners = Array.filter (_ /= "") <<< String.split (Pattern "\n")
  add acc entry = do
    when (entry.name == "" || Map.member entry.name acc) $
      throwError (error "incremental cache: duplicate or empty read-owner name")
    -- Decode the required JSON field separately: null is unknown, whereas an
    -- empty string is a known empty proof. An omitted field is invalid above.
    case hush (decodeJson entry.text) :: Maybe (Maybe String) of
      Nothing -> throwError (error "incremental cache: invalid read-owner text")
      Just text -> pure (Map.insert entry.name (map owners text) acc)
  read name = do
    text <- attempt (FS.readTextFile UTF8 (directory <> "/" <> name <> "/reads.json"))
    pure (Tuple name (case text of
      Left _ -> Nothing
      Right value -> Just (owners value)))

headerSafeFromReads :: Set String -> Map String (Maybe (Array String)) -> Set String -> Set String
headerSafeFromReads changed reads names =
  let
    providers (Tuple name value) = case value of
      Nothing -> []
      Just owners -> map (\owner -> Tuple owner (Set.singleton name)) owners
    reverse = Map.fromFoldableWith Set.union (Array.concatMap providers (Map.toUnfoldable reads :: Array (Tuple String (Maybe (Array String)))))
    missing = Set.filter (\name -> case Map.lookup name reads of
      Just (Just _) -> false
      _ -> true) names
    initial = Set.union changed missing
    dependents = go (Set.toUnfoldable initial :: List String) initial
    -- Mark on enqueue: each owner enters the list once and each reverse edge
    -- is examined once. Rebuilding the pending Set on every pop was quadratic.
    go Nil seen = seen
    go (name : pending) seen =
      let
        next = fromMaybe Set.empty (Map.lookup name reverse)
        queued = foldl
          (\acc reader -> if Set.member reader acc.seen then acc
            else { pending: reader : acc.pending, seen: Set.insert reader acc.seen })
          { pending, seen } next
      in go queued.pending queued.seen
  in
    Set.difference names dependents

-- Keep the flat mirror for diagnostics and older tooling. Reuse reads only
-- the committed, hash-verified objects frozen under the private prior tree.
publishReads :: String -> String -> Array String -> Aff Unit
publishReads stage output names = do
  let directory = output <> "/.gopurs-cache/reads-v1"
  void $ attempt (FS.mkdir directory)
  for_ names \name -> void $ attempt
    (linkOrCopy (stage <> "/modules/" <> name <> "/reads.json") (directory <> "/" <> name <> ".json"))

report :: ModuleCache -> Int -> Int -> Aff Unit
report cache total optimized = do
  invalidations <- liftEffect (Ref.read cache.invalidations)
  restored <- liftEffect (Ref.read cache.restored)
  let reasons = Map.fromFoldableWith (+) (map (\reason -> Tuple reason 1) (Array.fromFoldable (Map.values invalidations)))
  liftEffect $ Console.error $
    "[gopurs] module cache: optimized=" <> show optimized <> ", reused=" <> show (total - optimized)
      <> ", emitted=" <> show (total - restored) <> ", restored=" <> show restored
      <> "; invalidations=" <> stringify (encodeJson (Map.toUnfoldable reasons :: Array (Tuple String Int)))
