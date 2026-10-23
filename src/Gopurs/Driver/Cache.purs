module Gopurs.Driver.Cache
  ( Session, SourceModule, FrozenFFI, ReusedMemo, PartText, begin, plan, materialize, materializeSync, readParts, packages, publish, close ) where

import Prelude

import Data.Argonaut.Core (stringify)
import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Either (either)
import Data.Maybe (fromMaybe, isJust)
import Effect (Effect)
import Effect.Aff (Aff)
import Effect.Class (liftEffect)
import Effect.Exception (error, throwException)
import Gopurs.BuildCache (sessionExchange)
import Gopurs.Driver.Config (CompilerConfig)

type SourceModule = { name :: String, path :: String, hasForeign :: Boolean }
type FrozenFFI = { name :: String, path :: String, file :: String, found :: Boolean, digest :: String }
type Session =
  { status :: String, reason :: String, stage :: String, input :: String
  , output :: String, root :: String, count :: Int, ffi :: Array FrozenFFI, context :: String, sameTast :: Boolean
  , packageRequest :: String
  }
type Setting = { name :: String, value :: String }
type PartText = { name :: String, part :: String, text :: String }
type ReusedMemo =
  { name :: String, prepared :: String, optimizerEnvironment :: String
  , callerDemands :: String, emitterEnvironment :: String, metadata :: String, ffiDigest :: String
  }
type Request =
  { operation :: String, options :: Array Setting, force :: Boolean, stage :: String
  , modules :: Array SourceModule, entries :: Array String, diagnostics :: Boolean
  , moduleCache :: Boolean
  , owners :: Array String, parts :: Array String, reused :: Array ReusedMemo
  }

request :: String -> Request
request operation = { operation, options: [], force: false, stage: "", modules: [], entries: [], diagnostics: false, moduleCache: false, owners: [], parts: [], reused: [] }

exchange :: Request -> Aff Session
exchange = liftEffect <<< exchangeSync

exchangeSync :: Request -> Effect Session
exchangeSync message = do
  text <- sessionExchange (stringify (encodeJson message))
  json <- either (throwException <<< error) pure (jsonParser text)
  response <- either (throwException <<< error <<< printJsonDecodeError) pure (decodeJson json)
  when (response.status == "error") $ throwException (error
    ((if message.operation == "materialize" || message.operation == "packages" || message.operation == "publish" && not (Array.null message.reused) then "incremental cache: " else "") <> response.reason))
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
      , { name: "go-layout", value: if config.goPackages then "gopurs/go-packages-v1" else "gopurs/monolithic-v1" }
      , { name: "module-lazy", value: if config.optimizerJobs > 1 then "1" else "0" }
      ]
  }

plan :: Session -> Array SourceModule -> Aff (Array FrozenFFI)
plan session modules = _.ffi <$> exchange ((request "plan") { stage = session.stage, modules = modules })

materialize :: Session -> Array String -> Array String -> Aff Unit
materialize session owners parts = liftEffect (materializeSync session owners parts)

materializeSync :: Session -> Array String -> Array String -> Effect Unit
materializeSync session owners parts = void $ exchangeSync ((request "materialize") { stage = session.stage, owners = owners, parts = parts })

readParts :: Session -> Array String -> Array String -> Aff (Array PartText)
readParts session owners parts = liftEffect do
  text <- sessionExchange (stringify (encodeJson ((request "read-parts") { stage = session.stage, owners = owners, parts = parts })))
  json <- either (throwException <<< error <<< ("incremental cache: " <> _)) pure (jsonParser text)
  response <- either (throwException <<< error <<< ("incremental cache: " <> _) <<< printJsonDecodeError) pure
    (decodeJson json :: _ { status :: String, reason :: String })
  when (response.status /= "ready") $ throwException (error ("incremental cache: " <> response.reason))
  result <- either (throwException <<< error <<< ("incremental cache: " <> _) <<< printJsonDecodeError) pure
    (decodeJson json :: _ { contributions :: Array PartText })
  pure result.contributions

packages :: Session -> Array String -> Aff String
packages session entries = _.packageRequest <$> exchange ((request "packages") { stage = session.stage, entries = entries })

publish :: Session -> Array String -> Boolean -> Array ReusedMemo -> Aff Unit
publish session entries diagnostics reused = void $ exchange ((request "publish") { stage = session.stage, entries = entries, diagnostics = diagnostics, moduleCache = true, reused = reused })

close :: Session -> Aff Unit
close session = when (session.stage /= "") $ void $ exchange ((request "abort") { stage = session.stage })
