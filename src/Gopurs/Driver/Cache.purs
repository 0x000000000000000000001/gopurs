module Gopurs.Driver.Cache
  ( Session, SourceModule, FrozenFFI, ReusedMemo, PartText, begin, withSession, withSessionUsing, plan, materialize, materializeSync, materializeInputs, materializeAllInputs, readParts, readCapsule, packages, publish, close ) where

import Prelude

import Data.Argonaut.Core (stringify)
import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Either (either, hush)
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Effect (Effect)
import Effect.Aff (Aff, attempt, bracket, joinFiber, launchAff)
import Effect.Class (liftEffect)
import Effect.Exception (error, throwException)
import Effect.Exception as Exception
import Effect.Ref as Ref
import Gopurs.Metrics as Metrics
import Gopurs.BuildCache (sessionExchange, sessionExchangeOnMiss)
import Gopurs.Driver.Config (CompilerConfig)

type SourceModule = { name :: String, path :: String, hasForeign :: Boolean }
type FrozenFFI = { name :: String, path :: String, file :: String, found :: Boolean, digest :: String }
type Session =
  { status :: String, reason :: String, stage :: String, input :: String, inputLayout :: String
  , output :: String, root :: String, cacheRoot :: String, count :: Int, ffi :: Array FrozenFFI, context :: String, sameTast :: Boolean
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
  , owners :: Array String, all :: Boolean, parts :: Array String, reused :: Array ReusedMemo
  }

request :: String -> Request
request operation = { operation, options: [], force: false, stage: "", modules: [], entries: [], diagnostics: false, moduleCache: false, owners: [], all: false, parts: [], reused: [] }

exchange :: Request -> Aff Session
exchange = liftEffect <<< exchangeSync

exchangeSync :: Request -> Effect Session
exchangeSync message = do
  text <- sessionExchange (stringify (encodeJson message))
  decodeResponse message text

decodeResponse :: Request -> String -> Effect Session
decodeResponse message text = do
  json <- either (throwException <<< error) pure (jsonParser text)
  response <- either (throwException <<< error <<< printJsonDecodeError) pure (decodeJson json)
  when (response.status == "error") $ throwException (error
    ((if message.operation == "materialize" || message.operation == "materialize-input" || message.operation == "packages" || message.operation == "publish" && not (Array.null message.reused) then "incremental cache: " else "") <> response.reason))
  pure response

begin :: CompilerConfig -> Aff Session
begin = exchange <<< beginRequest

beginRequest :: CompilerConfig -> Request
beginRequest config = (request "begin")
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
      , { name: "input-layout", value: "packed-v1" }
      ] <> if config.devMode then [ { name: "optimization-profile", value: "dev-v5" } ] else []
  }

-- Own the speculative fiber independently of session acquisition. Even a failed
-- or cancelled begin joins it before closing the session or allowing a retry.
withSession :: forall a b. CompilerConfig -> Maybe (String -> Aff (Maybe a)) -> (Session -> Maybe a -> Aff b) -> Aff b
withSession config preload = withSessionUsing acquire closeMeasured preload
  where
  message = beginRequest config
  acquire onMiss = do
    text <- case preload of
      Nothing -> sessionExchange (stringify (encodeJson message))
      Just _ -> sessionExchangeOnMiss (stringify (encodeJson message)) onMiss
    decodeResponse message text
  closeMeasured session = Metrics.measure "cache close" \_ -> close session

-- The request bridge is explicit so lifecycle tests can inject a failed begin
-- after notification, without adding test hooks to the native store.
withSessionUsing :: forall a b. ((String -> Effect Unit) -> Effect Session) -> (Session -> Aff Unit) -> Maybe (String -> Aff (Maybe a)) -> (Session -> Maybe a -> Aff b) -> Aff b
withSessionUsing acquire release preload use = bracket
  (liftEffect do
    session <- Ref.new Nothing
    fiber <- Ref.new Nothing
    pure { session, fiber })
  (\owned -> do
    pending <- liftEffect (Ref.read owned.fiber)
    case pending of
      Nothing -> pure unit
      Just fiber -> void (attempt (joinFiber fiber))
    acquired <- liftEffect (Ref.read owned.session)
    case acquired of
      Nothing -> pure unit
      Just session -> release session)
  (\owned -> do
    let onMiss seed = void $ Exception.try do
          existing <- Ref.read owned.fiber
          case existing, preload of
            Nothing, Just load -> do
              -- No suspension between launch and retaining the handle. Errors
              -- inside the optional task remain data, including fast failures.
              fiber <- launchAff (fromMaybe Nothing <<< hush <$> attempt (load seed))
              Ref.write (Just fiber) owned.fiber
            _, _ -> pure unit
    session <- Metrics.measure "cache lookup + freeze" \_ -> liftEffect do
      result <- acquire onMiss
      -- Register in the same Effect, before Aff can observe cancellation.
      Ref.write (Just result) owned.session
      pure result
    pending <- liftEffect (Ref.read owned.fiber)
    projected <- case pending of
      Nothing -> pure Nothing
      Just fiber -> Metrics.measure "preparation preload join" \_ ->
        fromMaybe Nothing <<< hush <$> attempt (joinFiber fiber)
    use session projected)

plan :: Session -> Array SourceModule -> Aff (Array FrozenFFI)
plan session modules = _.ffi <$> exchange ((request "plan") { stage = session.stage, modules = modules })

materialize :: Session -> Array String -> Array String -> Aff Unit
materialize session owners parts = liftEffect (materializeSync session owners parts)

materializeSync :: Session -> Array String -> Array String -> Effect Unit
materializeSync session owners parts = void $ exchangeSync ((request "materialize") { stage = session.stage, owners = owners, parts = parts })

materializeInputs :: Session -> Array String -> Aff Unit
materializeInputs session owners = when (session.inputLayout == "packed-v1") $
  void $ exchange ((request "materialize-input") { stage = session.stage, owners = owners })

materializeAllInputs :: Session -> Aff Unit
materializeAllInputs session = when (session.inputLayout == "packed-v1") $
  void $ exchange ((request "materialize-input") { stage = session.stage, all = true })

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

readCapsule :: Session -> String -> Aff (Maybe String)
readCapsule session reader = liftEffect do
  text <- sessionExchange (stringify (encodeJson ((request "read-capsule") { stage = session.stage, owners = [ reader ] })))
  json <- either (throwException <<< error <<< ("incremental cache: " <> _)) pure (jsonParser text)
  response <- either (throwException <<< error <<< ("incremental cache: " <> _) <<< printJsonDecodeError) pure
    (decodeJson json :: _ { status :: String, reason :: String, capsule :: Maybe String })
  when (response.status /= "ready") $ throwException (error ("incremental cache: " <> response.reason))
  pure response.capsule

packages :: Session -> Array String -> Aff String
packages session entries = _.packageRequest <$> exchange ((request "packages") { stage = session.stage, entries = entries })

publish :: Session -> Array String -> Boolean -> Array ReusedMemo -> Aff Unit
publish session entries diagnostics reused = void $ exchange ((request "publish") { stage = session.stage, entries = entries, diagnostics = diagnostics, moduleCache = true, reused = reused })

close :: Session -> Aff Unit
close session = when (session.stage /= "") $ void $ exchange ((request "abort") { stage = session.stage })
