module Gopurs.Driver.Output (emitModule, emitPreparedModule, writeEntryPoints, writeRuntime) where

import Prelude

import Control.Lazy (defer)
import Data.Array as Array
import Data.Either (either)
import Data.Foldable (for_)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple)
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
import Gopurs.Runtime (runtimeGoCode)
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
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
writeEntryPoints output modules = for_ modules \moduleName -> do
  let source = renderEntryPoint moduleName
  -- The shared entry follows traversal order; each module also keeps its own
  -- executable entry. An explicit --main selects exactly one module upstream.
  void $ attempt (FS.mkdir (output <> "/main"))
  FS.writeTextFile UTF8 (output <> "/main/main.go") source
  void $ attempt (FS.mkdir (output <> "/" <> moduleName))
  void $ attempt (FS.mkdir (output <> "/" <> moduleName <> "/main"))
  FS.writeTextFile UTF8 (output <> "/" <> moduleName <> "/main/main.go") source

moduleGoName :: String -> String
moduleGoName = String.replaceAll (Pattern ".") (Replacement "_")

-- Keep the executable template readable, including its final newline. Module
-- bodies go through the Go AST/printer; this fixed wrapper only varies by name.
renderEntryPoint :: String -> String
renderEntryPoint moduleName = String.joinWith "\n"
  [ "package main"
  , ""
  , "import ("
  , "\t\"os\""
  , "\t\"runtime/pprof\""
  , "\t\"gopurs/output/purescript\""
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
