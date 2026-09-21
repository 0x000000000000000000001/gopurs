module Main where

import Prelude

import Control.Lazy (defer)
import Control.Parallel (parTraverse)
import Effect (Effect)
import Effect.Ref as Ref
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Aff (Aff, launchAff_, attempt, bracket)
import Node.FS.Aff as FS
import Node.Encoding (Encoding(..))
import Node.Process as Process
import Gopurs.FfiBridge as FfiBridge
import Gopurs.Metrics as Metrics
import Gopurs.Emission (createEmitter, createPipelinedEmitter)
import Data.Array as Array
import Data.Foldable (foldl)
import Data.Int as Int
import Data.List as List
import Data.List (List)
import Data.Traversable (traverse)
import Data.Maybe (Maybe(..), isJust, fromMaybe)
import Data.Map as Map

import Data.Set as Set
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.String as String
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.Builder (buildModules)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.Semantics.Foreign (coreForeignSemantics)
import PureScript.Backend.Optimizer.CoreFn (Module(..), Ann, Ident(..))
import Gopurs.CodeGen (translateWithFunctions)
import Gopurs.CodegenState (CodegenMetadata, CodegenMetadataRow)
import Gopurs.ExprContext (ModuleFunctions)
import Gopurs.AdtMetadata (buildPointerAdtMetadata, buildEnumAdtMetadata)
import Gopurs.ClassMetadata (buildClassFields, addClassDataDeclarations)
import Gopurs.ConstructorMetadata (buildConstructorTypes, collectElidedConstructors)
import Gopurs.Runtime (runtimeGoCode)
import PureScript.Backend.Optimizer.FfiSupport (findFfiFile)
import Gopurs.FfiSupport (prepareFfi)
import Gopurs.GlobalTypes (buildGlobalTypes)
import Gopurs.Monomorphization (monomorphizeModules)
import Gopurs.ReboxMetadata (buildReboxFieldIndex)
import PureScript.Backend.Optimizer.App (coreFnModulesFromOutput, parseCLIArgs, loadDirectives)
import PureScript.Backend.Optimizer.Semantics (InlineDirectiveMap)

type PreparedData =
  { directives :: InlineDirectiveMap
  , monomorphizedModules :: List (Module Ann)
  , targetMainModules :: Array String
  | CodegenMetadataRow
  }

loadAndPrepareModules :: { mbMainModule :: Maybe String } -> Aff PreparedData
loadAndPrepareModules args = do
  finalModules <- Metrics.measure "load TAST + sort" \_ -> coreFnModulesFromOutput "output"

  Metrics.measure "prepare + monomorphize" \_ -> do
    let elidedCtors = collectElidedConstructors (Array.fromFoldable finalModules)

    directives <- loadDirectives

    let ctorTypes = buildConstructorTypes (Array.fromFoldable finalModules)

    let globalTypes = buildGlobalTypes (Array.fromFoldable finalModules)
    let
      classDeclsFields = buildClassFields (Array.fromFoldable finalModules)
      reboxFields = buildReboxFieldIndex ctorTypes classDeclsFields
      finalModulesWithClassDecls = map addClassDataDeclarations finalModules

    let monomorphizedModules = monomorphizeModules globalTypes finalModulesWithClassDecls
    let
      { pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves } = buildPointerAdtMetadata (Array.fromFoldable finalModulesWithClassDecls)
      { enumAdts, enumCtors } = buildEnumAdtMetadata (Array.fromFoldable finalModules)

      targetMainModules = case args.mbMainModule of
        Just mainMod -> [ mainMod ]
        Nothing -> Array.mapMaybe (\(Module m) -> if isJust (Array.elemIndex (Ident "main") m.exports) then Just (unwrap m.name) else Nothing) (Array.fromFoldable finalModules)

    pure { directives
         , elidedCtors
         , ctorTypes
         , globalTypes
         , globalFunctions: Map.empty
         , classDeclsFields
         , reboxFields
         , monomorphizedModules
         , pointerAdtPaths
         , pointerAdtNodes
         , pointerAdtLeaves
         , enumAdts
         , enumCtors
         , targetMainModules
         }

emitModule :: CodegenMetadata -> Maybe String -> Module Ann -> BackendModule -> Aff ModuleFunctions
emitModule metadata mbFfiDir (Module coreFnMod) backendMod = defer \_ -> do
  let modNameStr = unwrap backendMod.name
  let safeModName = String.replaceAll (Pattern ".") (Replacement "_") modNameStr

  let translated = translateWithFunctions metadata backendMod
  FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> ".go") translated.code

  when (Array.length (Array.fromFoldable backendMod.foreign) > 0) do
    ffiPathMb <- liftEffect $ findFfiFile ".go" [] mbFfiDir modNameStr (Just coreFnMod.path)
    case ffiPathMb of
      Just ffiPath -> do
        content <- FS.readTextFile UTF8 ffiPath
        ffi <- liftEffect $ prepareFfi { moduleName: modNameStr, path: ffiPath } (safeModName <> "_") content

        let finalPkgLine = "package purescript"
        let hasImport = String.contains (Pattern "\"gopurs/output/gopurs_runtime\"") content
        let importLine = if hasImport then "" else "import \"gopurs/output/gopurs_runtime\"\n"

        let newContent = finalPkgLine <> "\n\n" <> importLine <> "\n" <> ffi.content <> "\n\n// --- Auto-generated FFI wrappers ---\n" <> FfiBridge.generateFfiBridge safeModName backendMod.dataDecls ffi.decls (Map.toUnfoldable backendMod.foreign)
        FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> "_ffi.go") newContent
      Nothing -> do

        let dummyContent = "package purescript\n\nimport \"gopurs/output/gopurs_runtime\"\n\n" <> FfiBridge.generateFfiBridge safeModName backendMod.dataDecls [] (Map.toUnfoldable backendMod.foreign)
        FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> "_ffi.go") dummyContent

  pure translated.functions

main :: Effect Unit
main = launchAff_ $ Metrics.measure "backend total" \_ -> do
  argsRaw <- liftEffect Process.argv
  let args = parseCLIArgs argsRaw

  prepared <- loadAndPrepareModules { mbMainModule: args.mbMainModule }
  globalFunctionsRef <- liftEffect (Ref.new Map.empty)
  configuredEmitJobs <- liftEffect (Process.lookupEnv "GOPURS_EMIT_JOBS")
  configuredPipeline <- liftEffect (Process.lookupEnv "GOPURS_PIPELINE")
  let emitJobs = max 1 (min 64 (fromMaybe 2 (configuredEmitJobs >>= Int.fromString)))

  Metrics.measure "runtime" \_ -> do
    _ <- attempt (FS.mkdir "output/gopurs_runtime")
    FS.writeTextFile UTF8 "output/gopurs_runtime/runtime.go" runtimeGoCode

    _ <- attempt (FS.mkdir "output/purescript")

    FS.writeTextFile UTF8 "output/go.mod" "module gopurs/output\n\ngo 1.22\n"

  let
    directives = prepared.directives
    monomorphizedModules = prepared.monomorphizedModules
    targetMainModules = prepared.targetMainModules
    metadata :: CodegenMetadata
    metadata =
      { enumAdts: prepared.enumAdts
      , enumCtors: prepared.enumCtors
      , pointerAdtPaths: prepared.pointerAdtPaths
      , pointerAdtNodes: prepared.pointerAdtNodes
      , pointerAdtLeaves: prepared.pointerAdtLeaves
      , elidedCtors: prepared.elidedCtors
      , ctorTypes: prepared.ctorTypes
      , globalTypes: prepared.globalTypes
      , globalFunctions: prepared.globalFunctions
      , classDeclsFields: prepared.classDeclsFields
      , reboxFields: prepared.reboxFields
      }

  let
    emitBatch batch = do
      globalFunctions <- liftEffect (Ref.read globalFunctionsRef)
      let emit entry = emitModule (metadata { globalFunctions = globalFunctions }) args.mbFfiDir entry.coreFnModule entry.backendMod
      functions <- if emitJobs == 1 then traverse emit batch else parTraverse emit batch
      -- No worker mutates shared metadata. Publish in the original module order.
      liftEffect (Ref.write (foldl (flip Map.union) globalFunctions functions) globalFunctionsRef)
    makeEmitter =
      if configuredPipeline /= Just "0" then createPipelinedEmitter emitJobs emitBatch
      else do
        emitter <- createEmitter emitJobs emitBatch
        pure { enqueue: emitter.enqueue, finish: emitter.finish, cancel: pure unit }

  Metrics.measure "optimize + emit" \_ ->
    bracket (liftEffect makeEmitter) _.cancel \emitter -> do
      buildModules
        { directives: directives
        , analyzeCustom: \_ _ -> Nothing
        , foreignSemantics: coreForeignSemantics
        , traceIdents: Set.empty
        , rewriteLimit: fromMaybe 10_000 args.mbRewriteLimit
        , onPrepareModule: \env (Module m) -> do
            when (env.moduleIndex `mod` 100 == 0) $ liftEffect $ Console.error $
              "[gopurs] optimize + emit: module " <> show (env.moduleIndex + 1)
                <> "/" <> show env.moduleCount <> " (" <> unwrap m.name <> ")"
            pure (Module m)
        -- Regenerate every module and its FFI output on each invocation.
        , onSkipModule: \_ _ -> pure Nothing
        , onCodegenModule: \_ coreFnModule backendMod _ -> do
            emitter.enqueue
              { name: backendMod.name
              , imports: backendMod.imports
              , value: { coreFnModule, backendMod }
              }
        }
        (List.fromFoldable monomorphizedModules)
      emitter.finish

  _ <- Metrics.measure "entry points" \_ -> traverse
    ( \mainMod -> do
        let pkgName = String.replaceAll (Pattern ".") (Replacement "_") mainMod
        let mainEntryPoint = "package main\n\nimport (\n\t\"os\"\n\t\"runtime/pprof\"\n\t\"gopurs/output/purescript\"\n\t\"gopurs/output/gopurs_runtime\"\n)\n\nfunc main() {\n\tif os.Getenv(\"PPROF\") == \"1\" {\n\t\tf, err := os.Create(\"cpu.prof\")\n\t\tif err != nil { panic(err) }\n\t\tpprof.StartCPUProfile(f)\n\t\tdefer pprof.StopCPUProfile()\n\t}\n\n\tgopurs_runtime.Apply(purescript.Get_" <> pkgName <> "_main(), gopurs_runtime.Value{})\n\n\tgopurs_runtime.EventLoopWait()\n\n\tif os.Getenv(\"PPROF\") == \"1\" {\n\t\tmf, err := os.Create(\"mem.prof\")\n\t\tif err != nil { panic(err) }\n\t\tpprof.WriteHeapProfile(mf)\n\t\tmf.Close()\n\t}\n}\n"
        
        _ <- attempt (FS.mkdir "output/main")
        FS.writeTextFile UTF8 ("output/main/main.go") mainEntryPoint

        _ <- attempt (FS.mkdir ("output/" <> mainMod))
        _ <- attempt (FS.mkdir ("output/" <> mainMod <> "/main"))
        FS.writeTextFile UTF8 ("output/" <> mainMod <> "/main/main.go") mainEntryPoint
    )
    targetMainModules

  pure unit
