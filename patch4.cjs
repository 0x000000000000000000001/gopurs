const fs = require('fs');
const content = fs.readFileSync('src/Main.purs', 'utf8');

const replacement = `
foreign import takeMonomorphizedModules :: PreparedData -> Effect (List (Module Ann))

runBuild :: { mbMainModule :: Maybe String, mbRewriteLimit :: Maybe Int, mbFfiDir :: Maybe String } -> Aff (Array String)
runBuild args = do
  prepared <- loadAndPrepareModules { mbMainModule: args.mbMainModule }

  _ <- attempt (FS.mkdir "output/gopurs_runtime")
  FS.writeTextFile UTF8 "output/gopurs_runtime/runtime.go" runtimeGoCode

  _ <- attempt (FS.mkdir "output/purescript")

  FS.writeTextFile UTF8 "output/go.mod" "module gopurs/output\\n\\ngo 1.22\\n"

  let _ = unsafePerformEffect (logMemory "Before buildModules")
  
  let directives = prepared.directives
  let enumAdts = prepared.enumAdts
  let enumCtors = prepared.enumCtors
  let pointerAdtPaths = prepared.pointerAdtPaths
  let pointerAdtNodes = prepared.pointerAdtNodes
  let pointerAdtLeaves = prepared.pointerAdtLeaves
  let adtTypes = prepared.adtTypes
  let elidedCtors = prepared.elidedCtors
  let ctorTypes = prepared.ctorTypes
  let globalTypes = prepared.globalTypes
  let instantiations = prepared.instantiations
  let classDeclsMap = prepared.classDeclsMap
  let classDeclsFields = prepared.classDeclsFields
  let targetMainModules = prepared.targetMainModules
  
  modules <- liftEffect $ takeMonomorphizedModules prepared
  
  buildModules
    { directives: directives
    , analyzeCustom: \\_ _ -> Nothing
    , foreignSemantics: coreForeignSemantics
    , traceIdents: Set.empty
    , rewriteLimit: fromMaybe 10_000 args.mbRewriteLimit
    , onPrepareModule: \\_ (Module m) -> pure (Module m)
    , onSkipModule: \\_ (Module coreFnMod) -> do
        let modNameStr = unwrap coreFnMod.name
        let safeModName = String.replaceAll (Pattern ".") (Replacement "_") modNameStr
        let cachePath = "output/purescript/" <> safeModName <> ".gopurs-cache.json"
        res <- pure Nothing
        case res of
          Just _ -> do
            let ffiPath = "output/purescript/" <> safeModName <> "_ffi.go"
            ffiStatRes <- attempt (FS.stat ffiPath)
            cacheStatRes <- attempt (FS.stat cachePath)
            case ffiStatRes, cacheStatRes of
              Right ffiStat, Right cacheStat ->
                if Stats.modifiedTimeMs ffiStat > Stats.modifiedTimeMs cacheStat then do
                  pure Nothing
                else
                  pure res
              _, _ -> pure res
          Nothing -> pure Nothing
    , onCodegenModule: \\_ (Module coreFnMod) backendMod _ -> do
        let modNameStr = unwrap backendMod.name
        let safeModName = String.replaceAll (Pattern ".") (Replacement "_") modNameStr
        let importsArray = map (\\i -> String.split (Pattern ".") (unwrap (importName i))) coreFnMod.imports

        let goFile = translate enumAdts enumCtors pointerAdtPaths pointerAdtNodes pointerAdtLeaves adtTypes elidedCtors ctorTypes globalTypes instantiations classDeclsMap classDeclsFields importsArray backendMod
        FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> ".go") goFile

        when (Array.length (Array.fromFoldable backendMod.foreign) > 0) do
          ffiPathMb <- liftEffect $ findFfiFile ".go" [] args.mbFfiDir modNameStr (Just coreFnMod.path)
          case ffiPathMb of
            Just ffiPath -> do
              content' <- FS.readTextFile UTF8 ffiPath
              jsonStr <- liftEffect $ extractFfiAst modNameStr content'
              let parsed = (jsonParser jsonStr >>= (decodeJson >>> lmap printJsonDecodeError)) :: Either String (Array FfiDecl)
              let ffiDecls = case parsed of
                               Right d -> d
                               Left err -> unsafePerformEffect (Console.log ("JSON Parse error for " <> modNameStr <> ": " <> err) *> pure [])
              
              let goFfiFile = translateFfiAst modNameStr ffiDecls content'
              FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> "_ffi.go") goFfiFile
            Nothing ->
              liftEffect $ Console.log ("Warning: Missing FFI file for " <> modNameStr <> ". Expected .go file but it was not found in specified paths.")

        writeCache cacheVersion ("output/purescript/" <> safeModName <> ".gopurs-cache.json") backendMod
    }
    modules
    
  pure targetMainModules

main :: Effect Unit
main = launchAff_ do
`;

const startIndex = content.indexOf('runBuild ::');
const newContent = content.substring(0, startIndex) + replacement + content.substring(content.indexOf('  argsRaw <- liftEffect Process.argv', startIndex));
fs.writeFileSync('src/Main.purs', newContent);
