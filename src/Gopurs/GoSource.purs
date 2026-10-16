module Gopurs.GoSource
  ( SourceInput, BindRequest, bindSources
  , PackageInput, PackageRequest, PackageResult, packageSources
  , PackageFilesResult, packageFiles
  ) where

import Prelude
import Data.Argonaut.Core (Json, stringify)
import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Decode.Error (JsonDecodeError, printJsonDecodeError)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Either (Either, either)
import Effect (Effect)
import Effect.Exception (throw)
import Gopurs.FfiSupport (bindGoSourceImpl, packageGoSourceImpl, packageGoFilesImpl)
import Gopurs.GoCode (BoundGoCode)
import Gopurs.GoSymbol (GoSymbol)

type SourceInput =
  { owner :: String, path :: String, content :: String, initializers :: Array GoSymbol }

-- Exports is a JSON object mapping import paths to base64 gc export bytes.
-- Acquisition belongs to the build driver; binding is deterministic and has
-- no filesystem or subprocess dependency in any of the three transports.
type BindRequest =
  { packagePath :: String, goVersion :: String, arch :: String
  , files :: Array SourceInput, exports :: Json
  }

bindSources :: BindRequest -> Effect (Array BoundGoCode)
bindSources request = do
  response <- bindGoSourceImpl (stringify (encodeJson request))
  json <- either (throw <<< ("Invalid Go binding JSON: " <> _)) pure (jsonParser response)
  result <- either (throw <<< printJsonDecodeError) pure
    (decodeJson json :: Either JsonDecodeError { format :: String, files :: Array BoundGoCode })
  if result.format == "gopurs/go-source-v1" then pure result.files
  else throw ("Unsupported Go source binding format: " <> result.format)

type PackageInput =
  { owner :: String, moduleName :: String, path :: String, content :: String
  , initializers :: Array GoSymbol
  , records :: Array { "type" :: String, labels :: Array String }
  }

type PackageRequest =
  { packagePath :: String, goVersion :: String, arch :: String
  , files :: Array PackageInput, exports :: Json, entries :: Array GoSymbol
  }

type PackageResult =
  { format :: String, planHash :: String
  , packages :: Array { name :: String, dependencies :: Array String, declarations :: Int }
  , files :: Array { path :: String, content :: String }
  , entries :: Array { symbol :: GoSymbol, "package" :: String }
  }

packageSources :: PackageRequest -> Effect PackageResult
packageSources request = do
  response <- packageGoSourceImpl (stringify (encodeJson request))
  json <- either (throw <<< ("Invalid Go package JSON: " <> _)) pure (jsonParser response)
  result <- either (throw <<< printJsonDecodeError) pure
    (decodeJson json :: Either JsonDecodeError PackageResult)
  if result.format == "gopurs/go-packages-v1" then pure result
  else throw ("Unsupported Go package format: " <> result.format)

type PackageFilesResult =
  { format :: String, planHash :: String
  , packages :: Array { name :: String, dependencies :: Array String, declarations :: Int }
  , files :: Array { path :: String, kind :: String, moduleName :: String, sha256 :: String, bytes :: Number }
  , entries :: Array { symbol :: GoSymbol, "package" :: String }
  }

packageFiles :: String -> Effect PackageFilesResult
packageFiles request = do
  response <- packageGoFilesImpl request
  json <- either (throw <<< ("Invalid Go package file JSON: " <> _)) pure (jsonParser response)
  result <- either (throw <<< printJsonDecodeError) pure
    (decodeJson json :: Either JsonDecodeError PackageFilesResult)
  if result.format == "gopurs/go-package-files-v1" then pure result
  else throw ("Unsupported Go package file format: " <> result.format)
