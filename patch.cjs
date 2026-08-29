const fs = require('fs');
const content = fs.readFileSync('src/Main.purs', 'utf8');

const replacement = `
runBuild :: { mbMainModule :: Maybe String, mbRewriteLimit :: Maybe Int, mbFfiDir :: Maybe String } -> Aff (Array String)
runBuild args = do
  prepared <- loadAndPrepareModules { mbMainModule: args.mbMainModule }

  _ <- attempt (FS.mkdir "output/gopurs_runtime")
  FS.writeTextFile UTF8 "output/gopurs_runtime/runtime.go" runtimeGoCode

  _ <- attempt (FS.mkdir "output/purescript")

  FS.writeTextFile UTF8 "output/go.mod" "module gopurs/output\\n\\ngo 1.22\\n"

  let _ = unsafePerformEffect (logMemory "Before buildModules")
  
  buildModules
    { directives: prepared.directives
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

        let goFile = translate prepared.enumAdts prepared.enumCtors prepared.pointerAdtPaths prepared.pointerAdtNodes prepared.pointerAdtLeaves prepared.adtTypes prepared.elidedCtors prepared.ctorTypes prepared.globalTypes prepared.instantiations prepared.classDeclsMap prepared.classDeclsFields importsArray backendMod
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
              
              let goFfiFile = translateFfi modNameStr ffiDecls content'
              FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> "_ffi.go") goFfiFile
            Nothing ->
              liftEffect $ Console.log ("Warning: Missing FFI file for " <> modNameStr <> ". Expected .go file but it was not found in specified paths.")

        writeCache cacheVersion ("output/purescript/" <> safeModName <> ".gopurs-cache.json") backendMod
    }
    (List.fromFoldable prepared.monomorphizedModules)
    
  pure prepared.targetMainModules

main :: Effect Unit
main = launchAff_ do
  argsRaw <- liftEffect Process.argv
  let args = parseCLIArgs argsRaw
  
  targetMainModules <- runBuild args
`;

const startIndex = content.indexOf('main :: Effect Unit');
const newContent = content.substring(0, startIndex) + replacement + content.substring(content.indexOf('  _ <- traverse', startIndex));
fs.writeFileSync('src/Main.purs', newContent);
