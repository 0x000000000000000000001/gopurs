module Main where

import Prelude

import Control.Lazy (defer)
import Control.Parallel (parTraverse)
import Effect (Effect)
import Effect.Ref as Ref
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Aff (Aff, launchAff_, attempt, bracket, forkAff, throwError)
import Effect.Aff.AVar as Avar
import Node.FS.Aff as FS
import Node.Encoding (Encoding(..))
import Node.Process as Process
import Gopurs.FfiBridge as FfiBridge
import Gopurs.Metrics as Metrics
import Gopurs.Emission (createEmitter, createPipelinedEmitter)
import Data.Array as Array
import Data.Foldable (foldl)
import Data.Either (Either(..), either)
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
import PureScript.Backend.Optimizer.Builder (buildModules, buildModulesParallel)
import PureScript.Backend.Optimizer.Cache as Cache
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
import Gopurs.Monomorphization (monomorphizeModulesWith)
import Gopurs.Preparation (runPreparationJobs)
import PureScript.Backend.Optimizer.Monomorphize (transitiveCollectWith)
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

    configuredPrepareJobs <- liftEffect (Process.lookupEnv "GOPURS_PREPARE_JOBS")
    let prepareJobs = fromMaybe 2 (configuredPrepareJobs >>= Int.fromString)
    monomorphizedModules <- monomorphizeModulesWith
      (\ast instantiations -> Metrics.measure "transitive specializations" \_ ->
        transitiveCollectWith (runPreparationJobs prepareJobs) ast instantiations) globalTypes finalModulesWithClassDecls
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
  configuredPboJobs <- liftEffect (Process.lookupEnv "GOPURS_PBO_JOBS")
  mbAllocProfile <- liftEffect (Process.lookupEnv "GOPURS_ALLOC_PROFILE")
  let
    emitJobs = max 1 (min 64 (fromMaybe 8 (configuredEmitJobs >>= Int.fromString)))
    pboJobs = max 1 (min 64 (fromMaybe 1 (configuredPboJobs >>= Int.fromString)))

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

  pboAttemptsRef <- liftEffect (Ref.new 0)
  pboCodegenRef <- liftEffect (Ref.new 0)
  Metrics.measure "optimize + emit" \_ ->
    bracket (liftEffect makeEmitter) _.cancel \emitter -> do
      let
        buildOpts =
          { directives: directives
          , analyzeCustom: \_ _ -> Nothing
          , foreignSemantics: coreForeignSemantics
          , traceIdents: Set.empty
          , rewriteLimit: fromMaybe 10_000 args.mbRewriteLimit
          -- Les tentatives peuvent être rejouées : les compteurs sont
          -- affichés à la fin, pas ici.
          , onPrepareModule: \_ m -> do
              liftEffect (Ref.modify_ (_ + 1) pboAttemptsRef)
              pure m
          -- Regenerate every module and its FFI output on each invocation.
          , onSkipModule: \_ _ -> pure Nothing
          -- Un appel par module finalisé, dans l'ordre canonique : la seule
          -- progression monotone, quel que soit l'ordonnancement.
          , onCodegenModule: \env coreFnModule backendMod _ -> do
              liftEffect (Ref.modify_ (_ + 1) pboCodegenRef)
              when (env.moduleIndex `mod` 100 == 0) $ liftEffect $ Console.error $
                "[gopurs] optimize + emit: module " <> show (env.moduleIndex + 1)
                  <> "/" <> show env.moduleCount <> " (" <> unwrap backendMod.name <> ")"
              emitter.enqueue
                { name: backendMod.name
                , imports: backendMod.imports
                , value: { coreFnModule, backendMod }
                }
          }
        sortedModules = List.fromFoldable monomorphizedModules
        onStats stats = liftEffect $ Console.error $
          "[gopurs] pbo stats: dispatched=" <> show stats.dispatched
            <> ", fallbackDispatched=" <> show stats.fallbackDispatched
            <> ", maxReady=" <> show stats.maxReady
            <> ", deferredAttempts=" <> show stats.deferredAttempts
            <> ", wakeups=" <> show stats.wakeups
            <> ", waitingPeak=" <> show stats.waitingPeak
            <> ", attemptMillis=" <> show (Int.round stats.attemptMillis)
            <> ", attemptMaxMillis=" <> show (Int.round stats.attemptMaxMillis)
            <> ", coordinatorMillis=" <> show (Int.round stats.coordinatorMillis)
            <> ", awaitMillis=" <> show (Int.round stats.awaitMillis)
            <> ", emitMillis=" <> show (Int.round stats.emitMillis)
      if pboJobs <= 1 then
        buildModules buildOpts sortedModules
      else do
        resultsVar <- Avar.empty
        let
          forkJob job = void $ forkAff do
            outcome <- attempt (job unit)
            Avar.put outcome resultsVar
          awaitJob = do
            outcome <- Avar.take resultsVar
            either throwError pure outcome
        buildModulesParallel
          { jobs: pboJobs, scheduler: { fork: forkJob, await: awaitJob }, onStats: Just onStats }
          buildOpts
          sortedModules
      emitter.finish
      attempts <- liftEffect (Ref.read pboAttemptsRef)
      codegen <- liftEffect (Ref.read pboCodegenRef)
      liftEffect $ Console.error $
        "[gopurs] pbo module attempts: " <> show attempts <> ", codegen: " <> show codegen

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

  case mbAllocProfile of
    Just path -> liftEffect (Cache.writeAllocProfile path)
    Nothing -> pure unit

  pure unit
