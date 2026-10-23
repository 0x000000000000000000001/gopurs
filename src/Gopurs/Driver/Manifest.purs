module Gopurs.Driver.Manifest
  ( Manifest, ManifestEntry, buildManifest, buildManifestWithCache, buildManifestWithInterfaces, moduleInterface, readManifest, writeManifest, diffManifest
  , readFfiDigests, writeFfiDigests, readGlobalFunctions, writeGlobalFunctions
  ) where

import Prelude

import Data.Array as Array
import Data.Either (Either(..))
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Aff (Aff, attempt)
import Effect.Class (liftEffect)
import Gopurs.Cache.Wire as Wire
import Gopurs.ExprContext (ModuleFunctions)
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binding(..), ClassDecl, DataDecl, ExprType, Ident, Import(..), Module(..), ReExport)
import PureScript.Backend.Optimizer.Semantics (InlineDirectiveMap)

type ManifestEntry =
  { source :: String
  , interface :: String
  }

type Manifest =
  { schema :: Int
  , context :: String
  , directives :: String
  , modules :: Map String ManifestEntry
  }

type Interface =
  { imports :: Array String
  , exports :: Array Ident
  , reExports :: Array ReExport
  , dataDecls :: Array DataDecl
  , classDecls :: Array ClassDecl
  , foreign :: Map Ident (Maybe ExprType)
  , declarations :: Map Ident (Maybe ExprType)
  }

manifestPath :: String -> String
manifestPath output = output <> "/.gopurs-cache/build-manifest-v1.json"

ffiDigestsPath :: String -> String
ffiDigestsPath output = output <> "/.gopurs-cache/ffi-digests-v1.json"

globalFunctionsPath :: String -> String
globalFunctionsPath output = output <> "/.gopurs-cache/global-functions-v1.json"

readGlobalFunctions :: String -> Aff (Maybe ModuleFunctions)
readGlobalFunctions output = do
  text <- attempt (FS.readTextFile UTF8 (globalFunctionsPath output))
  pure case text of
    Left _ -> Nothing
    Right value -> Wire.read value

writeGlobalFunctions :: String -> ModuleFunctions -> Aff Unit
writeGlobalFunctions output functions = do
  text <- liftEffect (Wire.write functions)
  let path = globalFunctionsPath output
  FS.writeTextFile UTF8 (path <> ".next") text
  FS.rename (path <> ".next") path

readFfiDigests :: String -> Aff (Maybe (Map String String))
readFfiDigests output = do
  text <- attempt (FS.readTextFile UTF8 (ffiDigestsPath output))
  pure case text of
    Left _ -> Nothing
    Right value -> Wire.read value

writeFfiDigests :: String -> Map String String -> Aff Unit
writeFfiDigests output digests = do
  text <- liftEffect (Wire.write digests)
  let path = ffiDigestsPath output
  FS.writeTextFile UTF8 (path <> ".next") text
  FS.rename (path <> ".next") path

moduleInterface :: Module Ann -> Effect String
moduleInterface (Module value) = do
  let
    declarations = Map.fromFoldable (declarationTypes value.decls)
    interface =
      { imports: map (\(Import _ name) -> unwrap name) value.imports
      , exports: value.exports
      , reExports: value.reExports
      , dataDecls: value.dataDecls
      , classDecls: value.classDecls
      , foreign: value.foreign
      , declarations
      } :: Interface
  Wire.hash <$> Wire.write interface

declarationTypes :: Array (Bind Ann) -> Array (Tuple Ident (Maybe ExprType))
declarationTypes = Array.concatMap \bind -> case bind of
  NonRec (Binding (Ann annotation) ident _) -> [ Tuple ident annotation.type ]
  Rec bindings -> map (\(Binding (Ann annotation) ident _) -> Tuple ident annotation.type) bindings

buildManifest :: String -> String -> InlineDirectiveMap -> Array (Module Ann) -> Map String String -> Effect Manifest
buildManifest = buildManifestWithCache Nothing

buildManifestWithCache :: Maybe Manifest -> String -> String -> InlineDirectiveMap -> Array (Module Ann) -> Map String String -> Effect Manifest
buildManifestWithCache = buildManifestWithInterfaces Map.empty

-- Compact preparation supplies interfaces from its hash-verified projection,
-- anchored to the committed generation. Its inert headers cannot reconstruct
-- source interfaces, even when the optional build-manifest sidecar is absent.
buildManifestWithInterfaces :: Map String String -> Maybe Manifest -> String -> String -> InlineDirectiveMap -> Array (Module Ann) -> Map String String -> Effect Manifest
buildManifestWithInterfaces known previous _output context directives modules sourceHashes = do
  directivesKey <- Wire.hash <$> Wire.write directives
  let cached = case previous of
        Just saved | saved.schema == 1 && saved.context == context && saved.directives == directivesKey -> saved.modules
        _ -> Map.empty
  entries <- traverse (entry cached) modules
  pure
    { schema: 1
    , context
    , directives: directivesKey
    , modules: Map.fromFoldable entries
    }
  where
  entry cached (Module value) = do
    let
      name = unwrap value.name
      source = fromMaybe "" (Map.lookup name sourceHashes)
    interface <- case Map.lookup name known of
      Just exact -> pure exact
      Nothing -> case Map.lookup name cached of
        Just old | source /= "" && old.source == source -> pure old.interface
        _ -> moduleInterface (Module value)
    pure (Tuple name { source, interface })

writeManifest :: String -> Manifest -> Aff Unit
writeManifest output manifest = do
  let directory = output <> "/.gopurs-cache"
  void $ attempt (FS.mkdir directory)
  text <- liftEffect (Wire.write manifest)
  let path = manifestPath output
  FS.writeTextFile UTF8 (path <> ".next") text
  FS.rename (path <> ".next") path

readManifest :: String -> Aff (Maybe Manifest)
readManifest output = do
  text <- attempt (FS.readTextFile UTF8 (manifestPath output))
  pure case text of
    Left _ -> Nothing
    Right value -> Wire.read value

diffManifest :: Manifest -> Manifest -> Map String String
diffManifest previous current =
  Map.fromFoldable
    ( Array.mapMaybe
        (\(Tuple name entry) -> case Map.lookup name previous.modules of
          Nothing -> Just (Tuple name "added")
          Just old
            | old.source /= entry.source && old.interface /= entry.interface -> Just (Tuple name "source+interface")
            | old.source /= entry.source -> Just (Tuple name "source")
            | old.interface /= entry.interface -> Just (Tuple name "interface")
            | otherwise -> Nothing)
        (Map.toUnfoldable current.modules :: Array (Tuple String ManifestEntry))
    )
    <> Map.fromFoldable
      ( Array.mapMaybe
          (\(Tuple name _) -> if Map.member name current.modules then Nothing else Just (Tuple name "removed"))
          (Map.toUnfoldable previous.modules :: Array (Tuple String ManifestEntry))
      )
