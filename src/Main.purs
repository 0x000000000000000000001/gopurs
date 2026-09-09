module Main where

import Prelude

import Effect (Effect)
import Effect.Console as Console
import Effect.Class (liftEffect)
import Effect.Aff (Aff, launchAff_, attempt)
import Node.FS.Aff as FS
import Node.Encoding (Encoding(..))
import Node.Process as Process
import Gopurs.CodeGen as CodeGen
import Effect.Unsafe (unsafePerformEffect)
import Data.Argonaut.Parser (jsonParser)
import Data.Either (Either(..))
import Data.Bifunctor (lmap)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Array as Array
import Data.Tuple (Tuple(..))
import Data.List as List
import Data.List (List)
import Data.Traversable (traverse)
import Data.Maybe (Maybe(..), isJust, fromMaybe)
import Data.Map as Map
import Data.Foldable (foldl)

import Data.Set as Set
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.String as String
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.Builder (buildModules)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.Monomorphize (collectInstantiations, InstantiationMap, collectAllTypes, monomorphize, transitiveCollect)
import PureScript.Backend.Optimizer.Semantics.Foreign (coreForeignSemantics)
import PureScript.Backend.Optimizer.CoreFn (Module(..), Ann, importName, Bind(..), Binding(..), ExprType(..), Ident(..))
import Gopurs.CodeGen (translate)
import Gopurs.AdtMetadata (buildPointerAdtMetadata, buildEnumAdtMetadata)
import Gopurs.ClassMetadata (ClassFields, buildClassFields, addClassDataDeclarations)
import Gopurs.ConstructorMetadata (ConstructorTypes, buildConstructorTypes, collectElidedConstructors)
import Gopurs.Runtime (runtimeGoCode)
import PureScript.Backend.Optimizer.FfiSupport (findFfiFile)
import Gopurs.FfiSupport (extractFfiAst)
import Gopurs.FfiTypes (FfiDecl)
import Gopurs.GlobalTypes (buildGlobalTypes)
import PureScript.Backend.Optimizer.App (coreFnModulesFromOutput, parseCLIArgs, writeCache, loadDirectives)
import Data.Argonaut.Decode (decodeJson)
import PureScript.Backend.Optimizer.Semantics (InlineDirectiveMap)

type PreparedData =
  { directives :: InlineDirectiveMap
  , elidedCtors :: Set.Set String
  , ctorTypes :: ConstructorTypes
  , globalTypes :: Map.Map String ExprType
  , classDeclsFields :: ClassFields
  , instantiations :: InstantiationMap
  , monomorphizedModules :: List (Module Ann)
  , adtTypes :: Set.Set ExprType
  , pointerAdtPaths :: Map.Map String { ctorName :: String, arity :: Int }
  , pointerAdtNodes :: Set.Set String
  , pointerAdtLeaves :: Map.Map String { nodeBaseStruct :: String, nodeCtor :: String }
  , enumAdts :: Set.Set String
  , enumCtors :: Set.Set String
  , targetMainModules :: Array String
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
  

  let globalAstMap = foldl (\acc (Module m) ->
        foldl (\acc' b -> case b of
          NonRec (Binding ann id e) -> Map.insert (unwrap m.name <> "." <> unwrap id) (Binding ann id e) acc'
          Rec binds -> foldl (\a (Binding ann id e) -> Map.insert (unwrap m.name <> "." <> unwrap id) (Binding ann id e) a) acc' binds
        ) acc m.decls
      ) Map.empty finalModulesWithClassDecls
      

  let rawInstantiations = foldl (collectInstantiations globalAstMap) Map.empty finalModulesWithClassDecls

  let transitiveInstantiations = transitiveCollect globalAstMap rawInstantiations

  let ffiGlobals = foldl (\acc (Module m) ->
         let modName = unwrap m.name
         in foldl (\acc' (Tuple (Ident name) _) -> Set.insert (modName <> "." <> name) acc') acc (Map.toUnfoldable m.foreign :: Array (Tuple Ident (Maybe ExprType)))
      ) Set.empty finalModulesWithClassDecls

  let instantiations = Map.filterKeys (\k -> not (Set.member k ffiGlobals) && case Map.lookup k globalTypes of
                                            Just t -> 
                                              let hasTV = hasTypeVariables t
                                              in hasTV
                                            Nothing -> false) transitiveInstantiations


  let monomorphizedModules =
        if Map.isEmpty instantiations then
          finalModulesWithClassDecls
        else
          map (monomorphize globalAstMap instantiations) finalModulesWithClassDecls


  let allTypes = foldl (\acc mod -> Set.union acc (collectAllTypes mod)) Set.empty finalModulesWithClassDecls
  let adtTypes = Set.filter (\t -> case t of
        ADT _ _ _ -> true
        _ -> false) allTypes
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
       , instantiations
       , monomorphizedModules
       , adtTypes
       , pointerAdtPaths
       , pointerAdtNodes
       , pointerAdtLeaves
       , enumAdts
       , enumCtors
       , targetMainModules
       }

hasTypeVariables :: ExprType -> Boolean
hasTypeVariables (TypeVar v) = String.take 1 v == String.toLower (String.take 1 v) && v /= "gopurs_runtime.Value"

hasTypeVariables (Func args ret) = Array.any hasTypeVariables args || hasTypeVariables ret
hasTypeVariables (Array t) = hasTypeVariables t
hasTypeVariables (Record row) = hasTypeVariables row
hasTypeVariables (Row props tail) = 
  let tailHas = case tail of
        Nothing -> false
        Just t -> hasTypeVariables t
  in Array.any (\(Tuple _ v) -> hasTypeVariables v) props || tailHas
hasTypeVariables (TypeApp c args) = hasTypeVariables c || Array.any hasTypeVariables args
hasTypeVariables (ForAll _ body) = hasTypeVariables body
hasTypeVariables (ConstrainedType constraints body) = Array.any (\(Tuple _ a) -> Array.any hasTypeVariables a) constraints || hasTypeVariables body
hasTypeVariables Int = false
hasTypeVariables String = false
hasTypeVariables Char = false
hasTypeVariables Number = false
hasTypeVariables Boolean = false
hasTypeVariables Unit = false
hasTypeVariables (TypeLevelString _) = false
hasTypeVariables (ADT _ _ args) = Array.any hasTypeVariables args
hasTypeVariables Any = false

cacheVersion :: String
cacheVersion = "1.0.0"


emitModule :: PreparedData -> Maybe String -> Module Ann -> BackendModule -> Aff Unit
emitModule prepared mbFfiDir (Module coreFnMod) backendMod = do
  let modNameStr = unwrap backendMod.name
  let safeModName = String.replaceAll (Pattern ".") (Replacement "_") modNameStr
  let importsArray = map (\i -> String.split (Pattern ".") (unwrap (importName i))) coreFnMod.imports

  let goFile = translate prepared.enumAdts prepared.enumCtors prepared.pointerAdtPaths prepared.pointerAdtNodes prepared.pointerAdtLeaves prepared.adtTypes prepared.elidedCtors prepared.ctorTypes prepared.globalTypes prepared.instantiations prepared.classDeclsFields importsArray backendMod
  FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> ".go") goFile

  when (Array.length (Array.fromFoldable backendMod.foreign) > 0) do
    ffiPathMb <- liftEffect $ findFfiFile ".go" [] mbFfiDir modNameStr (Just coreFnMod.path)
    case ffiPathMb of
      Just ffiPath -> do
        content <- FS.readTextFile UTF8 ffiPath
        jsonStr <- liftEffect $ extractFfiAst modNameStr content
        let parsed = (jsonParser jsonStr >>= (decodeJson >>> lmap printJsonDecodeError)) :: Either String (Array FfiDecl)
        let ffiDecls = case parsed of
                         Right d -> d
                         Left err -> unsafePerformEffect (Console.log ("JSON Parse error for " <> modNameStr <> ": " <> err) *> pure [])

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

        let newContent = finalPkgLine <> "\n\n" <> importLine <> "\n" <> String.joinWith "\n" renamedContentLines <> "\n\n// --- Auto-generated FFI wrappers ---\n" <> CodeGen.generateFfiBridge safeModName backendMod.dataDecls prefixedFfiDecls (Map.toUnfoldable backendMod.foreign)
        FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> "_ffi.go") newContent
      Nothing -> do

        let dummyContent = "package purescript\n\nimport \"gopurs/output/gopurs_runtime\"\n\n" <> CodeGen.generateFfiBridge safeModName backendMod.dataDecls [] (Map.toUnfoldable backendMod.foreign)
        FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> "_ffi.go") dummyContent
  writeCache cacheVersion ("output/purescript/" <> safeModName <> ".gopurs-cache.json") backendMod

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
