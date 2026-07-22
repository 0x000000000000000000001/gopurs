module Main where

import Prelude

import Effect (Effect)
import Effect.Class (liftEffect)
import Effect.Aff (Aff, launchAff_, attempt)
import Effect.Console as Console
import Node.FS.Aff as FS
import Node.FS.Stats as Stats
import Node.Encoding (Encoding(..))
import Node.Process as Process
import Data.Argonaut.Parser (jsonParser)
import Data.Either (Either(..), isRight)
import Data.Bifunctor (lmap)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Array as Array
import Data.List as List
import Data.Maybe (Maybe(..))
import Data.Map as Map
import Data.Set as Set
import Data.Traversable (traverse)
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.String as String
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.CoreFn.Json (decodeModule)
import PureScript.Backend.Optimizer.CoreFn.Sort (sortModules)
import PureScript.Backend.Optimizer.Builder (buildModules)
import PureScript.Backend.Optimizer.CoreFn (Module(..), Ann, importName)
import Gopurs.CodeGen (translate)
import Gopurs.Runtime (runtimeGoCode)
import Gopurs.FfiSupport (findFfiFile)

readCoreFnModule :: String -> Aff (Maybe (Module Ann))
readCoreFnModule filePath = do
  statRes <- attempt (FS.stat filePath)
  if isRight statRes then do
    contents <- FS.readTextFile UTF8 filePath
    case jsonParser contents >>= (lmap printJsonDecodeError <<< decodeModule) of
      Left err -> do
        liftEffect $ Console.error $ "Failed to decode " <> filePath <> ": " <> err
        pure Nothing
      Right mod -> pure (Just mod)
  else
    pure Nothing

main :: Effect Unit
main = launchAff_ do
  files <- FS.readdir "output"
  validDirs <- Array.filterA
    ( \f -> do
        stat <- FS.stat ("output/" <> f)
        pure (Stats.isDirectory stat)
    )
    files

  mbModules <- traverse (\dir -> readCoreFnModule ("output/" <> dir <> "/corefn.json")) validDirs
  let modulesArray = Array.catMaybes mbModules
  let modulesList = List.fromFoldable modulesArray
  let finalModules = sortModules modulesList

  _ <- attempt (FS.mkdir "output/gopurs_runtime")
  FS.writeTextFile UTF8 "output/gopurs_runtime/runtime.go" runtimeGoCode

  FS.writeTextFile UTF8 "output/go.mod" "module gopurs/output\n\ngo 1.22\n"

  buildModules
    { directives: Map.empty
    , analyzeCustom: \_ _ -> Nothing
    , foreignSemantics: Map.empty
    , traceIdents: Set.empty
    , onPrepareModule: \_ m -> pure m
    , onCodegenModule: \_ (Module coreFnMod) backendMod _ -> do
        let modNameStr = unwrap backendMod.name
        let importsArray = map (\i -> String.split (Pattern ".") (unwrap (importName i))) coreFnMod.imports
        let goFile = translate importsArray backendMod
        FS.writeTextFile UTF8 ("output/" <> modNameStr <> "/" <> String.replaceAll (Pattern ".") (Replacement "_") modNameStr <> ".go") goFile
        
        when (Array.length (Array.fromFoldable backendMod.foreign) > 0) do
          ffiPathMb <- liftEffect $ findFfiFile Nothing modNameStr (Just coreFnMod.path)
          case ffiPathMb of
            Just ffiPath -> do
              content <- FS.readTextFile UTF8 ffiPath
              FS.writeTextFile UTF8 ("output/" <> modNameStr <> "/" <> String.replaceAll (Pattern ".") (Replacement "_") modNameStr <> "_ffi.go") content
            Nothing -> pure unit
    }
    finalModules

  argv <- liftEffect Process.argv
  let mainModuleName = case Array.findIndex (_ == "--main") argv of
        Just i -> case Array.index argv (i + 1) of
          Just m -> m
          Nothing -> "Main"
        Nothing -> "Main"

  let pkgName = String.replaceAll (Pattern ".") (Replacement "_") mainModuleName
  let mainEntryPoint = "package main\n\nimport (\n\t\"gopurs/output/" <> mainModuleName <> "\"\n\t\"gopurs/output/gopurs_runtime\"\n)\n\nfunc main() {\n\tgopurs_runtime.Apply(" <> pkgName <> ".Get_Main(), gopurs_runtime.Value{})\n}\n"
  FS.writeTextFile UTF8 "output/main.go" mainEntryPoint
