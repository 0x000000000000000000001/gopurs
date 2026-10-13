module Gopurs.Driver.Cache
  ( Session, SourceModule, FrozenFFI, begin, plan, publish, close ) where

import Prelude

import Data.Argonaut.Core (stringify)
import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Either (either)
import Data.Maybe (fromMaybe, isJust)
import Effect.Aff (Aff, throwError)
import Effect.Class (liftEffect)
import Effect.Exception (error, try)
import Gopurs.BuildCache (sessionExchange)
import Gopurs.Driver.Config (CompilerConfig)

type SourceModule = { name :: String, path :: String, hasForeign :: Boolean }
type FrozenFFI = { name :: String, path :: String, file :: String, found :: Boolean, digest :: String }
type Session =
  { status :: String, reason :: String, stage :: String, input :: String
  , output :: String, count :: Int, ffi :: Array FrozenFFI, context :: String, sameTast :: Boolean
  }
type Setting = { name :: String, value :: String }
type Request =
  { operation :: String, options :: Array Setting, force :: Boolean, stage :: String
  , modules :: Array SourceModule, entries :: Array String, diagnostics :: Boolean
  , moduleCache :: Boolean
  }

request :: String -> Request
request operation = { operation, options: [], force: false, stage: "", modules: [], entries: [], diagnostics: false, moduleCache: false }

exchange :: Request -> Aff Session
exchange message = do
  result <- liftEffect $ try $ sessionExchange (stringify (encodeJson message))
  text <- either throwError pure result
  json <- either (throwError <<< error) pure (jsonParser text)
  response <- either (throwError <<< error <<< printJsonDecodeError) pure (decodeJson json)
  when (response.status == "error") $ throwError (error response.reason)
  pure response

begin :: CompilerConfig -> Aff Session
begin config = exchange $ (request "begin")
  { force = config.forceBuild || isJust config.allocationProfile
  , options =
      [ { name: "main", value: show config.mainModule }
      , { name: "ffi", value: fromMaybe "" config.ffiDirectory }
      , { name: "rewriteLimit", value: show config.rewriteLimit }
      , { name: "prepareJobs", value: show config.prepareJobs }
      , { name: "optimizerJobs", value: show config.optimizerJobs }
      , { name: "emissionJobs", value: show config.emission.jobs }
      , { name: "pipeline", value: show config.emission.pipelined }
      ]
  }

plan :: Session -> Array SourceModule -> Aff (Array FrozenFFI)
plan session modules = _.ffi <$> exchange ((request "plan") { stage = session.stage, modules = modules })

publish :: Session -> Array String -> Boolean -> Aff Unit
publish session entries diagnostics = void $ exchange ((request "publish") { stage = session.stage, entries = entries, diagnostics = diagnostics, moduleCache = true })

close :: Session -> Aff Unit
close session = when (session.stage /= "") $ void $ exchange ((request "abort") { stage = session.stage })
