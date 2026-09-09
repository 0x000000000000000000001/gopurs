module Gopurs.FfiSupport
  ( extractFfiDecls
  , decodeFfiDecls
  ) where

import Prelude

import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Argonaut.Parser (jsonParser)
import Data.Bifunctor (lmap)
import Data.Either (Either, either)
import Effect (Effect)
import Effect.Exception (message, throw, try)
import Gopurs.FfiTypes (FfiDecl)

foreign import extractFfiAstImpl :: String -> Effect String

decodeFfiDecls :: String -> Either String (Array FfiDecl)
decodeFfiDecls content = do
  json <- lmap ("Invalid JSON response: " <> _) (jsonParser content)
  lmap (("Invalid FFI declarations: " <> _) <<< printJsonDecodeError) (decodeJson json)

extractFfiDecls :: { moduleName :: String, path :: String } -> String -> Effect (Array FfiDecl)
extractFfiDecls source content = do
  result <- try do
    json <- extractFfiAstImpl content
    either throw pure (decodeFfiDecls json)
  either
    (\error -> throw ("FFI error in module " <> source.moduleName <> " (" <> source.path <> "): " <> message error))
    pure
    result
