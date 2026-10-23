module Gopurs.Cache.Modules (ModuleCache, Memo, Keys, Part, initialize, lookupAll, lookupBackend, lookupBackendWithReads, captureReads, emit, report, loadReads, headerSafeFromReads, publishReads, readImplementationsSync) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.Rec.Class (Step(..), tailRecM)
import Control.Parallel (parTraverse)
import Data.Argonaut.Core (stringify)
import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Either (Either(..), hush)
import Data.Foldable (foldl, foldM, for_)
import Data.List (List(..), (:))
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Traversable (sequence, traverse)
import Data.Tuple (Tuple(..), fst, snd)
import Effect.Aff (Aff, attempt, throwError)
import Effect (Effect)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Exception (error, throwException)
import Effect.Ref as Ref
import Gopurs.Cache.Wire as Wire
import Gopurs.CodeGen as CodeGen
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.Driver.Cache (FrozenFFI, Session)
import Gopurs.Driver.Cache as Cache
import Gopurs.Driver.Output as Output
import Gopurs.Driver.Prepare (PreparedModules)
import Gopurs.ExprContext (ModuleFunctions)
import Gopurs.GoSymbol (GoReference(..))
import Gopurs.Metrics as Metrics
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import Node.FS.Sync as Sync
import PureScript.Backend.Optimizer.Analysis (BackendAnalysis)
import PureScript.Backend.Optimizer.Convert (BackendImplementations, BackendModule, ExternLookup(..))
import PureScript.Backend.Optimizer.CoreFn (Ann, Ident(..), Module(..), Qualified(..))
import PureScript.Backend.Optimizer.Observation as Observation
import PureScript.Backend.Optimizer.Semantics (EvalRef(..), ExternImpl, InlineDirectiveMap)

type Part = { codec :: String, object :: { sha256 :: String, bytes :: Number } }
type Keys = { prepared :: String, optimizerEnvironment :: String, callerDemands :: String }
type Memo =
  { name :: String, prepared :: String, optimizerEnvironment :: String
  , callerDemands :: String, emitterEnvironment :: String, parts :: Array Part
  , metadata :: Maybe String, ffiDigest :: Maybe String
  }
type ModuleCache =
  { session :: Session
  , candidates :: Map String Memo
  , keys :: Map String Keys
  , metadata :: String
  , invalidations :: Ref.Ref (Map String String)
  , restored :: Ref.Ref Int
  , sameContext :: Boolean
  , reads :: Ref.Ref (Map String Observation.Reads)
  , restoredPBO :: Ref.Ref (Set String)
  , priorImplementations :: Ref.Ref (Map String (Maybe BackendImplementations))
  , priorWorkers :: Ref.Ref (Map String (Maybe ModuleFunctions))
  , currentDigests :: Ref.Ref (Map (Qualified Ident) String)
  , priorDigests :: Ref.Ref (Map (Qualified Ident) String)
  , backends :: Ref.Ref (Map String BackendModule)
  , ffiDigests :: Map String String
  , priorFfiDigests :: Map String String
  , headerSafe :: Set String
  , lazyContributions :: Boolean
  , reused :: Ref.Ref (Map String Cache.ReusedMemo)
  }

-- Versioned backend-v2 envelope: keep one durable object for the body and the
-- reads which justify reuse when the conservative prepared prefix has changed.
type CachedBackend = { body :: BackendModule, reads :: Maybe Observation.Reads }

initialize :: Int -> Session -> PreparedModules -> Set String -> Ref.Ref (Map String BackendModule) -> Map String String -> Map String String -> Set String -> Aff ModuleCache
initialize jobs session prepared reusableModules backends ffiDigests priorFfiDigests headerCandidates = do
  text <- FS.readTextFile UTF8 (session.stage <> "/prior/index.json")
  candidates <- case hush (jsonParser text) >>= hush <<< decodeJson of
    Just values -> pure (Map.fromFoldable (map (\memo -> Tuple memo.name memo) (values :: Array Memo)))
    Nothing -> throwError (error "Invalid private module cache index")
  seed <- liftEffect (Wire.hash <$> Wire.write (Tuple session.context prepared.directives))
  let modules = Array.fromFoldable prepared.modules
  -- Source/interface equality does not prove that specializations are equal.
  -- Only identical TAST or an explicit preparation proof can reuse these keys.
  let deltaReusable = if Set.size reusableModules == 0 then Map.empty else Map.filterWithKey (\name _ -> Set.member name reusableModules) candidates
  let reusable = if session.sameTast then candidates else deltaReusable
  fingerprints <- Metrics.measure "prepared module keys" \_ -> mapJobs jobs (preparedFingerprint reusable) modules
  let reused = Array.length (Array.filter (\(Module m) -> Map.member (unwrap m.name) reusable) modules)
  liftEffect $ Console.error $ "[gopurs] prepared fingerprints: computed=" <> show (Array.length modules - reused) <> ", reused=" <> show reused
  metadata <- Metrics.measure "global metadata key" \_ -> liftEffect (Wire.hash <$> Wire.write prepared.metadata)
  let
    sameContext = case Array.head modules of
      Just (Module m) -> map _.optimizerEnvironment (Map.lookup (unwrap m.name) candidates) == Just seed
      Nothing -> false
    keyed = foldl add { prefix: seed, keys: Map.empty } fingerprints
    changedPrepared = Set.fromFoldable (Array.mapMaybe
      (\(Tuple name keys) -> case Map.lookup name candidates of
        Just previous | previous.prepared == keys.prepared && previous.callerDemands == keys.callerDemands -> Nothing
        _ -> Just name)
      (Map.toUnfoldable keyed.keys :: Array (Tuple String Keys)))
    committedInputsMatch = Array.all (\(Module m) -> case Map.lookup (unwrap m.name) candidates of
      Just previous -> previous.metadata == Just metadata && previous.ffiDigest == Just (ffiDigestFor ffiDigests (unwrap m.name))
      Nothing -> false) modules
    -- Complete optimizer keys prove the module and its canonical predecessor
    -- prefix independently of the read-owner closure. In particular, a read
    -- of a future module remains absent when only that future body changes.
    -- Global metadata and all FFI still gate direct emitted-source restoration.
    exactKeySafe = if not sameContext || not committedInputsMatch then Set.empty
      else Set.fromFoldable (Array.mapMaybe
        (\(Tuple name keys) -> case Map.lookup name candidates of
          Just previous
            | previous.prepared == keys.prepared
                && previous.callerDemands == keys.callerDemands
                && previous.optimizerEnvironment == keys.optimizerEnvironment -> Just name
          _ -> Nothing)
        (Map.toUnfoldable keyed.keys :: Array (Tuple String Keys)))
  readSafe <- if not sameContext || not committedInputsMatch || Set.isEmpty headerCandidates then pure Set.empty else do
    let names = map (\(Module m) -> unwrap m.name) modules
    previousReads <- Metrics.measure "reads load" \_ -> loadReads (session.stage <> "/prior") names
    let decoded = Array.length (Array.filter (\(Tuple _ value) -> isJust value) (Map.toUnfoldable previousReads :: Array _))
    liftEffect $ Console.error ("[gopurs] reads decoded=" <> show decoded <> "/" <> show (Array.length names))
    pure (Set.intersection headerCandidates (headerSafeFromReads changedPrepared previousReads (Set.fromFoldable names)))
  let headerSafe = Set.union exactKeySafe readSafe
  invalidations <- liftEffect (Ref.new Map.empty)
  restored <- liftEffect (Ref.new 0)
  reads <- liftEffect (Ref.new Map.empty)
  restoredPBO <- liftEffect (Ref.new Set.empty)
  priorImplementations <- liftEffect (Ref.new Map.empty)
  priorWorkers <- liftEffect (Ref.new Map.empty)
  currentDigests <- liftEffect (Ref.new Map.empty)
  priorDigests <- liftEffect (Ref.new Map.empty)
  reusedMemos <- liftEffect (Ref.new Map.empty)
  pure { session, candidates, keys: keyed.keys, metadata, invalidations, restored, sameContext, reads, restoredPBO, priorImplementations, priorWorkers, currentDigests, priorDigests, backends, ffiDigests, priorFfiDigests, headerSafe, lazyContributions: false, reused: reusedMemos }

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
mapJobs configured task values
  | configured <= 1 || Array.null values = traverse task values
  | otherwise = do
      let count = Array.length values
      let workers = min 8 (min configured count)
      let chunkSize = max 1 (min 8 ((count + workers * 8 - 1) / (workers * 8)))
      next <- liftEffect (Ref.new 0)
      let worker = tailRecM
            (\results -> do
              start <- liftEffect $ Ref.modify' (\index -> { state: index + chunkSize, value: index }) next
              if start >= count then pure (Done results)
              else do
                completed <- traverse task (Array.slice start (start + chunkSize) values)
                pure (Loop (Tuple start completed : results)))
            Nil
      results <- parTraverse (\_ -> worker) (Array.range 0 (workers - 1))
      pure (Array.concatMap snd (Array.sortWith fst (Array.concatMap Array.fromFoldable results)))

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
    ensureParts cache name [ "2" ]
    implementations <- if cache.lazyContributions then pure (Just Map.empty) else readImplementations cache name
    directives <- readPart (prior cache name <> "2.json") :: Aff (Maybe InlineDirectiveMap)
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
    ensureParts cache name [ "2" ]
    directives <- readPart (prior cache name <> "2.json") :: Aff (Maybe InlineDirectiveMap)
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

ensureParts :: ModuleCache -> String -> Array String -> Aff Unit
ensureParts cache name parts = when (cache.lazyContributions && Map.member name cache.candidates) $
  Cache.materialize cache.session [ name ] parts

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
      case Wire.read text of
        Just value | Array.all (\(Qualified owner _) -> map unwrap owner == Just name) (Map.keys value # Array.fromFoldable) -> do
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
  implementation (Tuple name@(Qualified (Just owner) (Ident ident)) present) =
    case view.implementation (unwrap owner) ident of
      ExternPending -> pure false
      ExternMissing -> pure (not present)
      ExternFound _ | not present -> pure false
      ExternFound current -> do
        restored <- liftEffect (Ref.read cache.restoredPBO)
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

readWorkers :: ModuleCache -> String -> Aff (Maybe ModuleFunctions)
readWorkers cache name = do
  previous <- liftEffect (Ref.read cache.priorWorkers)
  case Map.lookup name previous of
    Just value -> pure value
    Nothing -> do
      ensureParts cache name [ "3" ]
      workers <- case Map.lookup name cache.candidates >>= \memo -> Array.index memo.parts 3 of
        Just part | part.codec == "gopurs/workers-v2" -> readPart (prior cache name <> "3.json")
        _ -> pure Nothing
      let value = do
            functions <- workers
            guard (Array.all (\info -> info.arity >= 0 && case info.reference of
              SymbolReference symbol -> symbol.owner /= "" && symbol.name /= ""
              LocalReference _ -> false) (Array.fromFoldable (Map.values functions)))
            pure functions
      liftEffect (Ref.modify_ (Map.insert name value) cache.priorWorkers)
      pure value

emit
  :: ModuleCache -> Map String FrozenFFI -> CodegenMetadata -> String -> Boolean
  -> Module Ann -> BackendModule -> Aff { functions :: ModuleFunctions, workers :: String }
emit cache ffi metadata _ optimizedCached coreFnModule backend = do
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
      Just value -> pure value
    functions <- readWorkers cache name
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
        pure { functions: value, workers: part.object.sha256 }
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
      parts <- liftEffect $ sequence
            [ Wire.write ({ body: backend { implementations = Map.empty, directives = Map.empty }, reads: Map.lookup name observations } :: CachedBackend)
            , Wire.write backend.implementations
            , Wire.write backend.directives
            ]
      liftEffect (Ref.modify_ (Map.delete name) cache.reads)
      traverse (\(Tuple i text) -> FS.writeTextFile UTF8 (directory <> show i <> ".json") text $> Wire.hash text)
        (Array.mapWithIndex Tuple parts)
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
      inputWorkers = CodeGen.workerInputs metadata.globalFunctions backend
    -- A key in this domain is committed only after checking containment below.
    -- Equal backend/metadata/compiler inputs reproduce the same transformations;
    -- all their imported signatures are then covered by this conservative view.
    -- PrepareModule's transformations do not consume imported worker signatures.
    inputEnvironment <- environment "input-workers-contained-v1" inputWorkers
    early <- cached inputEnvironment
    Tuple emitterEnvironment result <- case early of
      Just previous -> Tuple inputEnvironment <$> restore previous
      Nothing -> do
        let preparedModule = CodeGen.prepareModule metadata backend
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
    pure result

loadReads :: String -> Array String -> Aff (Map String (Maybe (Array String)))
-- Read only the hash-verified objects frozen from the committed module
-- manifest. The public reads-v1 mirror may belong to a different publication.
loadReads directory names = Map.fromFoldable <$> traverse read names
  where
  read name = do
    text <- attempt (FS.readTextFile UTF8 (directory <> "/" <> name <> "/reads.json"))
    pure (Tuple name (case text of
      Left _ -> Nothing
      Right value -> Just (Array.filter (_ /= "") (String.split (Pattern "\n") value))))

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
