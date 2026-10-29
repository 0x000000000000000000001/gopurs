module Gopurs.Driver.Output (emitModule, emitPreparedModule, writeEntryPoints, writeRuntime, writePackageOutput, writePackageFiles) where

import Prelude

import Control.Lazy (defer)
import Data.Argonaut.Core (stringify)
import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Either (Either(..), either)
import Data.Foldable (for_)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Effect.Aff (Aff, attempt, throwError)
import Effect.Class (liftEffect)
import Effect.Exception (error, try)
import Gopurs.CodeGen (PreparedModule, prepareModule, translatePreparedWithFunctions)
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.ExprContext (ModuleFunctions)
import Gopurs.Driver.Cache (FrozenFFI)
import Gopurs.FfiBridge as FfiBridge
import Gopurs.FfiSupport (prepareFfi)
import Gopurs.FfiTypes (FfiDecl)
import Gopurs.GoSource (PackageRequest, PackageResult, PackageFilesResult, packageSources, packageFiles)
import Gopurs.GoSymbol (GoSymbol, generated)
import Gopurs.Runtime.Source (runtimeGoCode)
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import Node.FS.Sync as FSSync
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (Ann, ExprType, Ident, Module)

type ModuleFfi =
  { content :: String
  , decls :: Array FfiDecl
  , importsRuntime :: Boolean
  }

emitModule :: String -> Map.Map String FrozenFFI -> CodegenMetadata -> Module Ann -> BackendModule -> Aff ModuleFunctions
emitModule output frozen metadata coreFnModule backendMod = defer \_ ->
  emitPreparedModule output frozen metadata coreFnModule backendMod (prepareModule metadata backendMod)

emitPreparedModule :: String -> Map.Map String FrozenFFI -> CodegenMetadata -> Module Ann -> BackendModule -> PreparedModule -> Aff ModuleFunctions
emitPreparedModule output frozen metadata _ backendMod preparedModule = defer \_ -> do
  let
    moduleName = unwrap backendMod.name
    goName = moduleGoName moduleName
    foreigns = Map.toUnfoldable backendMod.foreign
    hasForeigns = not (Array.null foreigns)

  -- Read native signatures before translation: saturated FFI calls can then
  -- use direct workers. Missing FFI and malformed FFI are distinct paths.
  ffi <- if hasForeigns then do
    case Map.lookup moduleName frozen of
      Nothing -> throwError (error ("FFI inputs were not frozen for " <> moduleName))
      Just source | not source.found -> pure Nothing
      Just source -> do
        original <- FS.readTextFile UTF8 source.file
        -- Native liftEffect does not catch synchronous Effect exceptions.
        -- Move parser failures into Aff so supervision and CLI reporting run.
        result <- liftEffect $ try $ prepareFfi { moduleName, path: source.path } (goName <> "_") original
        prepared <- either throwError pure result
        pure $ Just
          { content: prepared.content
          , decls: prepared.decls
          , importsRuntime: String.contains (Pattern "\"gopurs/output/gopurs_runtime\"") original
          }
    else pure Nothing

  let
    ffiFunctions = case ffi of
      Just source -> FfiBridge.ffiFunctionInfos goName foreigns source.decls
      Nothing -> Map.empty
    translated = translatePreparedWithFunctions (metadata { ffiFunctions = ffiFunctions }) preparedModule

  FS.writeTextFile UTF8 (output <> "/purescript/" <> goName <> ".go") translated.code
  when hasForeigns $
    FS.writeTextFile UTF8 (output <> "/purescript/" <> goName <> "_ffi.go")
      (renderFfiModule goName backendMod foreigns ffi)
  pure translated.functions

renderFfiModule
  :: String
  -> BackendModule
  -> Array (Tuple Ident (Maybe ExprType))
  -> Maybe ModuleFfi
  -> String
renderFfiModule goName backendMod foreigns ffi = case ffi of
  Nothing ->
    "package purescript\n\nimport \"gopurs/output/gopurs_runtime\"\n\n"
      <> FfiBridge.generateFfiBridge goName backendMod.dataDecls [] foreigns
  Just source ->
    let
      runtimeImport = if source.importsRuntime then "" else "import \"gopurs/output/gopurs_runtime\"\n"
      wrappers = FfiBridge.generateFfiBridge goName backendMod.dataDecls source.decls foreigns
      workers = FfiBridge.ffiValueWorkers goName backendMod.dataDecls foreigns source.decls
      workersBlock = if workers == "" then "" else "\n\n// --- Boxed-result FFI workers ---\n" <> workers
    in
      "package purescript\n\n" <> runtimeImport <> "\n" <> source.content
        <> "\n\n// --- Auto-generated FFI wrappers ---\n" <> wrappers <> workersBlock

writeRuntime :: String -> Aff Unit
writeRuntime output = do
  FS.mkdir (output <> "/gopurs_runtime")
  FS.writeTextFile UTF8 (output <> "/gopurs_runtime/runtime.go") runtimeGoCode
  FS.mkdir (output <> "/purescript")

writeEntryPoints :: String -> Array String -> Aff Unit
writeEntryPoints output modules = for_ modules \moduleName ->
  writeEntryPoint output moduleName (renderEntryPoint moduleName Nothing)

writeEntryPoint :: String -> String -> String -> Aff Unit
writeEntryPoint output moduleName source = do
  -- The shared entry follows traversal order; each module also keeps its own
  -- executable entry. An explicit --main selects exactly one module upstream.
  void $ attempt (FS.mkdir (output <> "/main"))
  FS.writeTextFile UTF8 (output <> "/main/main.go") source
  void $ attempt (FS.mkdir (output <> "/" <> moduleName))
  void $ attempt (FS.mkdir (output <> "/" <> moduleName <> "/main"))
  FS.writeTextFile UTF8 (output <> "/" <> moduleName <> "/main/main.go") source

-- Bind/type-check the complete private generation before writing fragments.
-- Publication and recovery own the destination afterwards; this writer never
-- promotes a staging tree to the public cache on its own.
writePackageOutput :: String -> Array String -> PackageRequest -> Aff PackageResult
writePackageOutput output modules request = do
  bound <- liftEffect $ try $ packageSources request
  result <- either throwError pure bound
  entries <- packageEntries modules result.entries
  void $ attempt (FS.mkdir (output <> "/purescript"))
  for_ result.files \file -> do
    case String.split (Pattern "/") file.path of
      [ directory, name ] | directory /= "" && name /= "" && directory /= "." && directory /= ".." -> do
        void $ attempt (FS.mkdir (output <> "/purescript/" <> directory))
        FS.writeTextFile UTF8 (output <> "/purescript/" <> file.path) file.content
      _ -> throwError (error ("Invalid generated package path: " <> file.path))
  for_ modules \moduleName -> case Map.lookup moduleName entries of
    Just target -> writeEntryPoint output moduleName (renderEntryPoint moduleName (Just { target, initialization: request.packagePath <> "/init_all" }))
    Nothing -> throwError (error ("Missing generated entry: " <> moduleName))
  pure result

packageEntries :: Array String -> Array { symbol :: GoSymbol, "package" :: String } -> Aff (Map.Map String String)
packageEntries modules placements = Map.fromFoldable <$> traverse entry modules
  where
  entry moduleName = do
    let goName = moduleGoName moduleName
        symbol = generated goName ("Get_" <> goName <> "_main")
    -- Validate entries against the exact response, before creating any output.
    case Array.find (\value -> value.symbol == symbol) placements of
      Just value -> pure (Tuple moduleName value."package")
      Nothing -> throwError (error ("Missing generated entry: " <> moduleName))

writePackageFiles :: String -> String -> Array String -> String -> String -> String -> Aff Unit
writePackageFiles output cacheRoot modules requestPath resultPath overlayPath = do
  request <- FS.readTextFile UTF8 requestPath
  prepared <- liftEffect (FSSync.exists resultPath)
  Tuple result resultText <- if prepared then do
    -- The package session already produced a proven overlay result. Never
    -- call the full planner again on this path.
    text <- FS.readTextFile UTF8 resultPath
    json <- either (throwError <<< error) pure (jsonParser text)
    decoded <- either (throwError <<< error <<< printJsonDecodeError) pure
      (decodeJson json :: Either _ PackageFilesResult)
    pure (Tuple decoded text)
  else do
    emission <- liftEffect $ try $ packageFiles request
    value <- either throwError pure emission
    pure (Tuple value (stringify (encodeJson value)))
  entries <- packageEntries modules result.entries
  for_ modules \moduleName -> case Map.lookup moduleName entries of
    Just target -> writeEntryPoint output moduleName (renderEntryPoint moduleName (Just { target, initialization: "gopurs/output/purescript/init_all" }))
    Nothing -> throwError (error ("Missing generated entry: " <> moduleName))
  unless prepared $ FS.writeTextFile UTF8 resultPath resultText
  let planDir = cacheRoot <> "/go-packages/plan-v1"
  void $ attempt (FS.mkdir cacheRoot)
  void $ attempt (FS.mkdir (cacheRoot <> "/go-packages"))
  void $ attempt (FS.mkdir planDir)
  -- Stage files live below the public cache. Copy/rename is atomic and keeps a
  -- concurrent writer from exposing a half-written plan. The package result
  -- itself stays in the stage because publication validates it again.
  FS.rename requestPath (planDir <> "/request.json")
  FS.copyFile resultPath (resultPath <> ".plan")
  FS.rename (resultPath <> ".plan") (planDir <> "/result.json")
  overlay <- attempt (FS.readTextFile UTF8 overlayPath)
  case overlay of
    Right _ -> FS.rename overlayPath (planDir <> "/overlay.json")
    Left _ -> void $ attempt (FS.unlink (planDir <> "/overlay.json"))

moduleGoName :: String -> String
moduleGoName = String.replaceAll (Pattern ".") (Replacement "_")

-- Keep the executable template readable, including its final newline. Module
-- bodies go through the Go AST/printer; this fixed wrapper only varies by name.
renderEntryPoint :: String -> Maybe { target :: String, initialization :: String } -> String
renderEntryPoint moduleName placement = String.joinWith "\n"
  [ "package main"
  , ""
  , "import ("
  , "\t\"os\""
  , "\t\"runtime/pprof\""
  , case placement of
      Nothing -> "\t\"gopurs/output/purescript\""
      Just plan -> "\tpurescript \"" <> plan.target <> "\"\n\t_ \"" <> plan.initialization <> "\""
  , "\t\"gopurs/output/gopurs_runtime\""
  , ")"
  , ""
  , "func main() {"
  , "\tif os.Getenv(\"PPROF\") == \"1\" {"
  , "\t\tf, err := os.Create(\"cpu.prof\")"
  , "\t\tif err != nil { panic(err) }"
  , "\t\tpprof.StartCPUProfile(f)"
  , "\t\tdefer pprof.StopCPUProfile()"
  , "\t}"
  , ""
  , "\tgopurs_runtime.Apply(purescript.Get_" <> moduleGoName moduleName <> "_main(), gopurs_runtime.Value{})"
  , ""
  , "\tgopurs_runtime.EventLoopWait()"
  , ""
  , "\tif os.Getenv(\"PPROF\") == \"1\" {"
  , "\t\tmf, err := os.Create(\"mem.prof\")"
  , "\t\tif err != nil { panic(err) }"
  , "\t\tpprof.WriteHeapProfile(mf)"
  , "\t\tmf.Close()"
  , "\t}"
  , "}"
  , ""
  ]
