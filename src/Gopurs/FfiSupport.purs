module Gopurs.FfiSupport
  ( extractFfiDecls
  , decodeFfiDecls
  , prepareFfi
  , bindGoSourceImpl
  , packageGoSourceImpl
  , packageGoFilesImpl
  ) where

import Prelude

import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Argonaut.Parser (jsonParser)
import Data.Maybe (Maybe(..))
import Data.Bifunctor (lmap)
import Data.Either (Either(..), either)
import Data.Array as Array
import Effect (Effect)
import Effect.Exception (message, throw, try)
import Gopurs.FfiTypes (FfiDecl)

foreign import extractFfiAstImpl :: String -> Effect String
foreign import prepareFfiAstImpl :: String -> String -> Effect String
foreign import bindGoSourceImpl :: String -> Effect String
foreign import packageGoSourceImpl :: String -> Effect String
foreign import packageGoFilesImpl :: String -> Effect String

decodeFfiDecls :: String -> Either String (Array FfiDecl)
decodeFfiDecls content = do
  json <- lmap ("Invalid JSON response: " <> _) (jsonParser content)
  declarations <- lmap (("Invalid FFI declarations: " <> _) <<< printJsonDecodeError) (decodeJson json)
  validateDeclarations declarations

extractFfiDecls :: { moduleName :: String, path :: String } -> String -> Effect (Array FfiDecl)
extractFfiDecls source content = withFfiContext source do
  json <- extractFfiAstImpl content
  either throw pure (decodeFfiDecls json)

prepareFfi :: { moduleName :: String, path :: String } -> String -> String -> Effect { decls :: Array FfiDecl, content :: String }
prepareFfi source prefix content = withFfiContext source do
  result <- prepareFfiAstImpl prefix content
  json <- either (throw <<< ("Invalid JSON response: " <> _)) pure (jsonParser result)
  prepared <- either (throw <<< ("Invalid FFI module: " <> _) <<< printJsonDecodeError) pure (decodeJson json)
  declarations <- either throw pure (validateDeclarations prepared.decls)
  pure (prepared { decls = declarations })

withFfiContext :: forall a. { moduleName :: String, path :: String } -> Effect a -> Effect a
withFfiContext source action = do
  result <- try action
  either
    (\error -> throw ("FFI error in module " <> source.moduleName <> " (" <> source.path <> "): " <> message error))
    pure
    result

validateDeclarations :: Array FfiDecl -> Either String (Array FfiDecl)
validateDeclarations declarations = case Array.find (\d -> d.schemaVersion /= 2) declarations of
  Just declaration -> Left ("Unsupported Go FFI signature schema " <> show declaration.schemaVersion <> "; expected 2. Rebuild the matching native/WASM parser and clear stale compiler artifacts.")
  Nothing -> Right declarations
