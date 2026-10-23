module Gopurs.Driver.Prepare.Persistent (preparePersistent, publishPersistent) where

import Prelude
import Data.Argonaut.Core (stringify)
import Data.Argonaut.Decode (class DecodeJson, decodeJson)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Either (Either(..), hush)
import Data.Foldable (foldl)
import Data.List (List)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Effect.Aff (Aff, attempt, throwError)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Exception (error)
import Gopurs.Cache.Wire as Wire
import Gopurs.ClassMetadata (addClassDataDeclarations)
import Gopurs.Driver.Cache (Session)
import Gopurs.Driver.Prepare (PreparedCache, prepareLoadedModules, selectMainModules, sourceMetadata)
import Gopurs.Driver.Prepare.Proof as Proof
import Gopurs.Metrics as Metrics
import Gopurs.Monomorphization (buildGlobalAstMap)
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import PureScript.Backend.Optimizer.App (coreFnModulesFromOutput)
import PureScript.Backend.Optimizer.CoreFn (Ann, Module(..))
import PureScript.Backend.Optimizer.CoreFn.Usage (invalidateSourceUsageModule)
import PureScript.Backend.Optimizer.Monomorphize (InstantiationMap, monomorphizeWith, normalizeInstantiation, SpecializationLookup)
import PureScript.Backend.Optimizer.Semantics (InlineDirectiveMap)

directory :: String -> String
directory root = root <> "/.gopurs-cache/preparation-literals-v1"

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

readProof :: String -> Aff (Maybe Proof.State)
readProof root = do
  envelope <- readJSON (proofPath root) :: Aff (Maybe { hash :: String, payload :: String })
  pure do
    saved <- envelope
    if Wire.hash saved.payload /= saved.hash then Nothing
    else hush (jsonParser saved.payload) >>= hush <<< decodeJson

writeAtomic :: String -> String -> String -> Aff Unit
writeAtomic stage path text = do
  let temporary = stage <> "/preparation-" <> Wire.hash path <> ".next"
  FS.writeTextFile UTF8 temporary text
  FS.rename temporary path

sourceHash :: String -> Module Ann -> Aff (Tuple String String)
sourceHash input (Module mod) = do
  text <- FS.readTextFile UTF8 (input <> "/" <> unwrap mod.name <> "/corefn.json")
  pure (Tuple (unwrap mod.name) (Wire.hash text))

preparePersistent :: Session -> Int -> Maybe String -> InlineDirectiveMap -> Aff PreparedCache
preparePersistent session jobs mainModule directives = do
  loaded <- Metrics.measure "load TAST + sort" \_ -> coreFnModulesFromOutput session.input
  let modules = Array.fromFoldable loaded
  when (Array.length modules /= session.count) $
    throwError (error "TAST decoding did not produce every inventoried module")
  hashes <- Map.fromFoldable <$> traverse (sourceHash session.input) modules
  candidate <- Metrics.measure "compact preparation probe" \_ -> probe session mainModule directives loaded hashes
  case candidate of
    Just prepared -> do
      liftEffect $ Console.error ("[gopurs] compact preparation: " <> show (Set.size prepared.reusable) <> " modules proven reusable")
      pure prepared
    Nothing -> prepareLoadedModules session.input jobs mainModule directives loaded

probe :: Session -> Maybe String -> InlineDirectiveMap -> List (Module Ann) -> Map String String -> Aff (Maybe PreparedCache)
probe session mainModule directives loaded sourceHashes = do
  saved <- readProof session.root
  committed <- snapshotStamp (session.stage <> "/prior/manifest.json")
  currentStamp <- snapshotStamp (session.stage <> "/session.json")
  directivesKey <- liftEffect (Wire.hash <$> Wire.write directives)
  case saved, committed, currentStamp of
    Just old, Just stamp, Just current
      | old.schema == 2 && old.context == session.context && old.directives == directivesKey && old.snapshot == stamp -> do
          let
            previousHashes = Map.fromFoldable (map (\entry -> Tuple entry.name entry.hash) old.sources)
            modules = Array.fromFoldable loaded
            changed = Array.filter (\(Module mod) -> Map.lookup (unwrap mod.name) previousHashes /= Map.lookup (unwrap mod.name) sourceHashes) modules
            sameInventory = Map.size previousHashes == Map.size sourceHashes && Array.all (\(Tuple name _) -> Map.member name previousHashes) (Map.toUnfoldable sourceHashes :: Array _)
          case changed of
            [ changedModule@(Module changedSource) ] | sameInventory -> do
              currentProof <- liftEffect (Proof.moduleProofWithReads (Set.fromFoldable old.bodyReads) changedModule)
              let previousProof = Array.find (\entry -> entry.name == unwrap changedSource.name) old.modules
              case previousProof of
                Just previous | Proof.permitsLiteralEdit (Set.fromFoldable old.bodyReads) previous currentProof -> do
                  let
                    enriched = map (invalidateSourceUsageModule <<< addClassDataDeclarations) loaded
                    ast = Map.delete "Data.Ring.negate" (buildGlobalAstMap enriched)
                    index = Map.fromFoldable (map (\entry -> Tuple entry.name (Set.fromFoldable entry.keys)) old.index)
                    sourceByName = Map.fromFoldable (map (\mod@(Module value) -> Tuple (unwrap value.name) mod) (Array.fromFoldable enriched))
                    payloads = Map.fromFoldable (map (\entry -> Tuple entry.name entry.hash) old.payloads)
                    materialize name = case Map.lookup name sourceByName of
                      Nothing -> pure Nothing
                      Just source -> do
                        own <- readOwn session.root index payloads name
                        pure (map (\instantiations -> if Map.isEmpty index then source else monomorphizeWith ast (hashedLookup index) instantiations source) own)
                  rewritten <- materialize (unwrap changedSource.name)
                  case rewritten of
                    Nothing -> pure Nothing
                    Just changedPrepared -> do
                      let
                        preparedModules = map (\source@(Module value) -> if value.name == changedSource.name then changedPrepared else Proof.placeholder (not (Map.isEmpty index)) source) enriched
                        metadata = sourceMetadata modules (map addClassDataDeclarations loaded)
                        prepared =
                          { directives, modules: preparedModules, metadata
                          , mainModules: selectMainModules mainModule modules
                          , sources: map (\(Module value) -> { name: unwrap value.name, path: value.path, hasForeign: not (Map.isEmpty value.foreign) }) modules
                          }
                        proof = old
                          { snapshot = current
                          , sources = map (\(Tuple name hash) -> { name, hash }) (Map.toUnfoldable sourceHashes :: Array _)
                          , modules = map (\entry -> if entry.name == currentProof.name then currentProof else entry) old.modules
                          }
                      pure (Just
                        { prepared, sourceHashes, sourceModules: loaded
                        , mono: { globalAstMap: ast, instantiations: Map.empty, bodyReads: Set.fromFoldable old.bodyReads }
                        , reusable: Set.delete (unwrap changedSource.name) (Map.keys sourceHashes)
                        , proof: Just proof, materialize: Just materialize
                        })
                _ -> pure Nothing
            _ -> pure Nothing
    _, _, _ -> pure Nothing

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
publishPersistent :: Session -> PreparedCache -> Aff Unit
publishPersistent session cache = do
  result <- attempt (publishPersistentImpl session cache)
  case result of
    Left _ -> liftEffect $ Console.error "[gopurs] compact preparation publication skipped: cache write unavailable"
    Right _ -> pure unit

publishPersistentImpl :: Session -> PreparedCache -> Aff Unit
publishPersistentImpl session cache = do
  stamp <- snapshotStamp (session.stage <> "/session.json")
  case stamp of
    Nothing -> pure unit
    Just snapshot -> do
      void $ attempt (FS.mkdir (directory session.root))
      directives <- liftEffect (Wire.hash <$> Wire.write cache.prepared.directives)
      reusable <- case cache.proof of
        Just proof -> pure (Just proof)
        Nothing -> do
          previous <- readProof session.root
          case previous of
            Just proof | proof.schema == 2 && proof.snapshot == snapshot
              && proof.context == session.context && proof.directives == directives -> do
                valid <- Array.all identity <$> traverse (\entry -> do
                  text <- hush <$> attempt (FS.readTextFile UTF8 (payloadPath session.root entry.hash))
                  pure (map Wire.hash text == Just entry.hash)) proof.payloads
                pure (if valid then Just proof else Nothing)
            _ -> pure Nothing
      state <- case reusable of
        Just proof -> pure (proof { snapshot = snapshot })
        Nothing -> do
          modules <- liftEffect (traverse (Proof.moduleProofWithReads cache.mono.bodyReads) (Array.fromFoldable cache.sourceModules))
          let
            index = map (\(Tuple name values) -> { name, keys: map Wire.hash (Set.toUnfoldable (Map.keys values) :: Array String) }) (Map.toUnfoldable cache.mono.instantiations :: Array _)
            grouped = foldl
              (\groups (Tuple name values) -> Map.insertWith Map.union (owner name) (Map.singleton name values) groups)
              Map.empty (Map.toUnfoldable cache.mono.instantiations :: Array _)
          payloads <- traverse (\(Tuple name values) -> do
            encoded <- liftEffect (Wire.write (map (map normalizeInstantiation) values))
            let hash = Wire.hash encoded
            writeAtomic session.stage (payloadPath session.root hash) encoded
            pure { name, hash }) (Map.toUnfoldable grouped :: Array _)
          pure
            { schema: 2, context: session.context, directives, snapshot
            , sources: map (\(Tuple name hash) -> { name, hash }) (Map.toUnfoldable cache.sourceHashes :: Array _)
            , modules, index, payloads, bodyReads: Set.toUnfoldable cache.mono.bodyReads
            }
      let payload = stringify (encodeJson state)
      writeAtomic session.stage (proofPath session.root) (stringify (encodeJson { hash: Wire.hash payload, payload }))
