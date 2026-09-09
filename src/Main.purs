module Main where

import Prelude

import Effect (Effect)
import Effect.Class (liftEffect)
import Effect.Aff (Aff, launchAff_, attempt)
import Node.FS.Aff as FS
import Node.Encoding (Encoding(..))
import Node.Process as Process
import Gopurs.FfiBridge as FfiBridge
import Data.Array as Array
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
import Gopurs.CodeGen (CodegenMetadata, CodegenMetadataRow, translate)
import Gopurs.AdtMetadata (buildPointerAdtMetadata, buildEnumAdtMetadata)
import Gopurs.ClassMetadata (buildClassFields, addClassDataDeclarations)
import Gopurs.ConstructorMetadata (buildConstructorTypes, collectElidedConstructors)
import Gopurs.Runtime (runtimeGoCode)
import PureScript.Backend.Optimizer.FfiSupport (findFfiFile)
import Gopurs.FfiSupport (extractFfiDecls)
import Gopurs.GlobalTypes (buildGlobalTypes)
import Gopurs.Monomorphization (monomorphizeModules)
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
  finalModules <- coreFnModulesFromOutput "output"

  let elidedCtors = collectElidedConstructors (Array.fromFoldable finalModules)

  directives <- loadDirectives

  let ctorTypes = buildConstructorTypes (Array.fromFoldable finalModules)

  let globalTypes = buildGlobalTypes (Array.fromFoldable finalModules)
  let
    classDeclsFields = buildClassFields (Array.fromFoldable finalModules)
    finalModulesWithClassDecls = map addClassDataDeclarations finalModules

  let monomorphizedModules = monomorphizeModules globalTypes finalModulesWithClassDecls
  let
    { pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves } = buildPointerAdtMetadata (Array.fromFoldable finalModulesWithClassDecls)
    { enumAdts, enumCtors } = buildEnumAdtMetadata (Array.fromFoldable finalModulesWithClassDecls)

    targetMainModules = case args.mbMainModule of
      Just mainMod -> [ mainMod ]
      Nothing -> Array.mapMaybe (\(Module m) -> if isJust (Array.elemIndex (Ident "main") m.exports) then Just (unwrap m.name) else Nothing) (Array.fromFoldable finalModules)

  pure { directives
       , elidedCtors
       , ctorTypes
       , globalTypes
       , classDeclsFields
       , monomorphizedModules
       , pointerAdtPaths
       , pointerAdtNodes
       , pointerAdtLeaves
       , enumAdts
       , enumCtors
       , targetMainModules
       }

emitModule :: PreparedData -> Maybe String -> Module Ann -> BackendModule -> Aff Unit
emitModule prepared mbFfiDir (Module coreFnMod) backendMod = do
  let modNameStr = unwrap backendMod.name
  let safeModName = String.replaceAll (Pattern ".") (Replacement "_") modNameStr
  let
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
      , classDeclsFields: prepared.classDeclsFields
      }

  let goFile = translate metadata backendMod
  FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> ".go") goFile

  when (Array.length (Array.fromFoldable backendMod.foreign) > 0) do
    ffiPathMb <- liftEffect $ findFfiFile ".go" [] mbFfiDir modNameStr (Just coreFnMod.path)
    case ffiPathMb of
      Just ffiPath -> do
        content <- FS.readTextFile UTF8 ffiPath
        ffiDecls <- liftEffect $ extractFfiDecls { moduleName: modNameStr, path: ffiPath } content

        let lines = String.split (Pattern "\n") (String.replaceAll (Pattern "\r") (Replacement "") content)
        let otherLines = Array.filter (\l -> not (String.contains (Pattern "package ") l)) lines
        let finalPkgLine = "package purescript"
        let hasImport = String.contains (Pattern "\"gopurs/output/gopurs_runtime\"") content
        let importLine = if hasImport then "" else "import \"gopurs/output/gopurs_runtime\"\n"

        let prefixedFfiDecls = map (\d -> d { name = safeModName <> "_" <> d.name }) ffiDecls
        let renamedContentLines = map (\l -> Array.foldl (\acc decl ->
                                          if decl.isVar
                                          then String.replaceAll (Pattern ("var " <> decl.name)) (Replacement ("var " <> safeModName <> "_" <> decl.name)) acc
                                          else String.replaceAll (Pattern ("func " <> decl.name)) (Replacement ("func " <> safeModName <> "_" <> decl.name)) acc
                                       ) l ffiDecls) otherLines

        let newContent = finalPkgLine <> "\n\n" <> importLine <> "\n" <> String.joinWith "\n" renamedContentLines <> "\n\n// --- Auto-generated FFI wrappers ---\n" <> FfiBridge.generateFfiBridge safeModName backendMod.dataDecls prefixedFfiDecls (Map.toUnfoldable backendMod.foreign)
        FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> "_ffi.go") newContent
      Nothing -> do

        let dummyContent = "package purescript\n\nimport \"gopurs/output/gopurs_runtime\"\n\n" <> FfiBridge.generateFfiBridge safeModName backendMod.dataDecls [] (Map.toUnfoldable backendMod.foreign)
        FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> "_ffi.go") dummyContent

main :: Effect Unit
main = launchAff_ do
  argsRaw <- liftEffect Process.argv
  let args = parseCLIArgs argsRaw

  prepared <- loadAndPrepareModules { mbMainModule: args.mbMainModule }

  _ <- attempt (FS.mkdir "output/gopurs_runtime")
  FS.writeTextFile UTF8 "output/gopurs_runtime/runtime.go" runtimeGoCode

  _ <- attempt (FS.mkdir "output/purescript")

  FS.writeTextFile UTF8 "output/go.mod" "module gopurs/output\n\ngo 1.22\n"

  let
    directives = prepared.directives
    monomorphizedModules = prepared.monomorphizedModules
    targetMainModules = prepared.targetMainModules

  buildModules
    { directives: directives
    , analyzeCustom: \_ _ -> Nothing
    , foreignSemantics: coreForeignSemantics
    , traceIdents: Set.empty
    , rewriteLimit: fromMaybe 10_000 args.mbRewriteLimit
    , onPrepareModule: \_ (Module m) -> pure (Module m)
    -- Regenerate every module and its FFI output on each invocation.
    , onSkipModule: \_ _ -> pure Nothing
    , onCodegenModule: \_ coreFnModule backendMod _ ->
        emitModule prepared args.mbFfiDir coreFnModule backendMod
    }
    (List.fromFoldable monomorphizedModules)

  _ <- traverse
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
