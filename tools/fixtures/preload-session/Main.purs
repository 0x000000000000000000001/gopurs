module Test.PreloadSessionNative where

import Prelude
import Data.Array as Array
import Data.Either (Either(..))
import Data.Map as Map
import Data.Set as Set
import Data.Maybe (Maybe(..), isNothing)
import Data.Traversable (traverse_)
import Data.Time.Duration (Milliseconds(..))
import Effect (Effect)
import Effect.Aff (Aff, attempt, bracket, delay, forkAff, killFiber, launchAff_)
import Effect.Class (liftEffect)
import Effect.Console (log)
import Effect.Exception (error, throw)
import Effect.Ref as Ref
import Gopurs.Cache.HeaderPreload as Header
import Gopurs.Cache.Wire as Wire
import Gopurs.Preparation (mapPreparationJobs)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (EvalRef(..), InlineDirectiveMap)
import Gopurs.Driver.Cache as Cache
import Gopurs.Driver.Config (CompilerConfig)

foreign import setupFixture :: Effect Unit
foreign import editFixture :: Effect Unit
foreign import clearRecipe :: Effect Unit
foreign import remainingStages :: Effect Int
foreign import cleanupFixture :: Effect Unit

check :: Boolean -> String -> Effect Unit
check ok reason = unless ok (throw reason)

config :: CompilerConfig
config =
  { mainModule: Nothing, ffiDirectory: Nothing, rewriteLimit: 10000
  , prepareJobs: 8, optimizerJobs: 8, emission: { jobs: 8, pipelined: true }
  , allocationProfile: Nothing, forceBuild: false, goPackages: false, buildManifest: true
  }

type Counts = { started :: Int, finished :: Int, active :: Int, used :: Int, closed :: Int }
newCounts :: Effect (Ref.Ref Counts)
newCounts = Ref.new { started: 0, finished: 0, active: 0, used: 0, closed: 0 }

load :: Ref.Ref Counts -> Boolean -> String -> Aff (Maybe Int)
load counts fail _ = bracket
  (liftEffect $ Ref.modify_ (\s -> s { started = s.started + 1, active = s.active + 1 }) counts)
  (\_ -> liftEffect $ Ref.modify_ (\s -> s { finished = s.finished + 1, active = s.active - 1 }) counts)
  (\_ -> do
    delay (Milliseconds 25.0)
    liftEffect $ when fail (throw "optional preload error")
    pure (Just 42))

main :: Effect Unit
main = launchAff_ do
  traverse_ nativeCase [ "hit", "miss", "failure", "forced", "allocation", "cold" ]
  lifecycleCases
  headerCases
  liftEffect $ log "PRELOAD-SESSION-NATIVE-PASS"

nativeCase :: String -> Aff Unit
nativeCase kind = bracket (liftEffect setupFixture) (\_ -> liftEffect cleanupFixture) \_ -> do
  when (kind /= "hit") (liftEffect editFixture)
  when (kind == "cold") (liftEffect clearRecipe)
  counts <- liftEffect newCounts
  let current = config { forceBuild = kind == "forced", allocationProfile = if kind == "allocation" then Just "unused" else Nothing }
  Cache.withSession current (Just (load counts (kind == "failure"))) \session preloaded -> liftEffect do
    value <- Ref.read counts
    let notified = kind == "miss" || kind == "failure"
    check (value.started == (if notified then 1 else 0)) ("wrong native callback count: " <> kind)
    check (value.active == 0 && value.finished == value.started) "native preload not joined before use"
    check (preloaded == (if kind == "miss" then Just 42 else Nothing)) "wrong optional preload result"
    check ((session.status == "hit") == (kind == "hit")) "native hit/miss changed"
  stages <- liftEffect remainingStages
  liftEffect do
    check (stages == 0) "native session was not closed"
    log ("PRELOAD-NATIVE-CASE " <> kind)

-- Fail/cancel the real PS lifecycle controller after it has launched the same
-- callback. The acquisition seam avoids inserting fault switches in Go storage.
lifecycleCases :: Aff Unit
lifecycleCases = do
  counts <- liftEffect newCounts
  let failedBegin notify = do
        notify "seed"
        throw "begin failed after callback"
  result <- attempt $ Cache.withSessionUsing failedBegin (release counts) (Just (load counts false)) (use counts)
  value <- liftEffect $ Ref.read counts
  liftEffect do
    check (case result of
      Left _ -> true
      Right _ -> false) "begin error not caught"
    check (value.started == 1 && value.finished == 1 && value.active == 0 && value.closed == 0 && value.used == 0) "begin error left an unjoined preload"
  cancelled <- liftEffect newCounts
  fiber <- forkAff $ Cache.withSessionUsing begin (release cancelled) (Just (load cancelled false)) (use cancelled)
  delay (Milliseconds 5.0)
  killFiber (error "caller cancelled") fiber
  stopped <- liftEffect $ Ref.read cancelled
  liftEffect do
    check (stopped.started == 1 && stopped.finished == 1 && stopped.active == 0 && stopped.closed == 1 && stopped.used == 0) "cancel did not join before close"
    log "PRELOAD-NATIVE-LIFECYCLE-PASS"
  fast <- liftEffect newCounts
  Cache.withSessionUsing begin (release fast) (Just (\_ -> pure (Just 7))) \_ immediate ->
    liftEffect $ check (immediate == Just 7) "immediate result lost before handle retention"
  where
  session = { status: "miss", reason: "test", stage: "fake", input: "", inputLayout: "", output: "", root: "", count: 0, ffi: [], context: "", sameTast: false, packageRequest: "" }
  begin notify = notify "seed" *> pure session
  release counts _ = liftEffect do
    value <- Ref.read counts
    check (value.active == 0 && value.started == value.finished) "release ran before preload joined"
    Ref.modify_ (\s -> s { closed = s.closed + 1 }) counts
  use counts _ _ = liftEffect $ Ref.modify_ (\s -> s { used = s.used + 1 }) counts

-- Exercise the actual current PS header reservoir and lifecycle in native Go.
-- The reader seam returns the verified-read boundary's data, without test-only
-- switches in storage. Storage SHA rejection is covered by the native helper suite.
headerCases :: Aff Unit
headerCases = do
  alpha <- liftEffect (headerText "Alpha")
  beta <- liftEffect (headerText "Beta")
  let parts = [ part "Alpha" alpha, part "Beta" beta ]
  traverse_ (\jobs -> do
    loaded <- Header.preloadUsing jobs headerSession (pure (Just [ "Alpha", "Beta" ])) (\_ -> pure parts)
    case loaded of
      Nothing -> liftEffect $ throw "valid header reservoir rejected"
      Just reservoir -> liftEffect do
        check ((Map.keys <$> Header.select headerSession (Set.singleton "Alpha") reservoir) == Just (Set.singleton "Alpha")) "late safe subset changed"
        check (isNothing (Header.select (headerSession { stage = "different" }) (Set.singleton "Alpha") reservoir)) "different stage accepted"
        check (isNothing (Header.select (headerSession { root = "different" }) (Set.singleton "Alpha") reservoir)) "different root accepted"
        check (isNothing (Header.select (headerSession { context = "different" }) (Set.singleton "Alpha") reservoir)) "different context accepted"
        check (isNothing (Header.select (headerSession { count = 3 }) (Set.singleton "Alpha") reservoir)) "different count accepted"
        check (isNothing (Header.select headerSession (Set.singleton "Missing") reservoir)) "missing owner accepted"
    traverse_ (\entries -> do
      invalid <- Header.preloadUsing jobs headerSession (pure (Just [ "Alpha", "Beta" ])) (\_ -> pure entries)
      liftEffect $ check (isNothing invalid) "invalid header batch accepted")
      [ Array.take 1 parts
      , [ part "Alpha" alpha, part "Alpha" alpha ]
      , [ part "Alpha" alpha, part "Beta" "{}" ]
      , [ part "Alpha" alpha, part "Beta" alpha ]
      , [ part "Alpha" alpha, { name: "Beta", part: "3", text: beta } ]
      ]
    Header.withOptional (Just (Header.readUsing headerSession (pure (Just [ "Alpha", "Beta" ])) (\_ -> pure parts))) \raw ->
      Header.withDecodedPreload jobs raw \pending -> case pending of
        Nothing -> liftEffect $ throw "two-phase header handle missing"
        Just wait -> do
          result <- wait
          case loaded, result of
            Just direct, Just phased -> liftEffect do
              let allOwners = Set.fromFoldable [ "Alpha", "Beta" ]
              directWire <- Wire.write (Header.select headerSession allOwners direct)
              phasedWire <- Wire.write (Header.select headerSession allOwners phased)
              check (directWire == phasedWire) "two-phase headers differ from direct decode"
            _, _ -> liftEffect $ throw "two-phase valid reservoir rejected"
    duplicate <- Header.preloadUsing jobs headerSession (pure (Just [ "Alpha", "Alpha" ])) (\_ -> pure parts)
    liftEffect $ check (isNothing duplicate) "duplicate inventory accepted"
    ) [ 1, 8 ]
  traverse_ headerLifecycle [ "unused", "error", "cancel", "child-error" ]
  traverse_ headerPhases [ "success", "read-error", "read-caller-error", "read-cancel", "decode-error", "decode-caller-error", "decode-cancel" ]
  Header.withOptional (Nothing :: Maybe (Aff (Maybe Int))) \pending -> liftEffect $ check (isNothing pending) "disabled header preload ran"
  Header.withOptional (Just (pure (Just 7))) \pending -> case pending of
    Nothing -> liftEffect $ throw "immediate header handle lost"
    Just action -> action >>= \value -> liftEffect $ check (value == Just 7) "immediate header result lost"
  liftEffect $ log "HEADER-PRELOAD-NATIVE-PASS"
  where
  headerSession = { status: "miss", reason: "test", stage: "headers", input: "", inputLayout: "", output: "", root: "root", count: 2, ffi: [], context: "context", sameTast: false, packageRequest: "" }
  part name text = { name, part: "2", text }

headerText :: String -> Effect String
headerText owner = Wire.write (Map.singleton (EvalExtern (Qualified (Just (ModuleName owner)) (Ident "value"))) Map.empty :: InlineDirectiveMap)

headerLifecycle :: String -> Aff Unit
headerLifecycle kind = do
  counts <- liftEffect newCounts
  let child index = bracket
        (liftEffect $ Ref.modify_ (\s -> s { started = s.started + 1, active = s.active + 1 }) counts)
        (\_ -> liftEffect $ Ref.modify_ (\s -> s { finished = s.finished + 1, active = s.active - 1 }) counts)
        (\_ -> do
          delay (Milliseconds (if index == 0 then 20.0 else 50.0))
          liftEffect $ when (kind == "child-error" && index == 0) (throw "header child failure")
          pure index)
  let task = Just <$> mapPreparationJobs 8 child [ 0, 1, 2, 3 ]
  let action = bracket (pure unit)
        (\_ -> liftEffect do
          value <- Ref.read counts
          check (value.active == 0 && value.started == value.finished) "outer close preceded header child join"
          Ref.modify_ (\s -> s { closed = s.closed + 1 }) counts)
        (\_ -> Header.withOptional (Just task) \pending -> do
          if kind == "child-error" then case pending of
            Nothing -> liftEffect $ throw "header handle missing"
            Just wait -> wait >>= \value -> liftEffect $ check (isNothing value) "optional header error was not ignored"
          else do
            delay (Milliseconds 5.0)
            liftEffect $ when (kind == "error") (throw "header caller failure")
            when (kind == "cancel") (delay (Milliseconds 200.0)))
  if kind == "cancel" then do
    fiber <- forkAff action
    delay (Milliseconds 10.0)
    killFiber (error "header caller cancelled") fiber
  else do
    result <- attempt action
    liftEffect $ check ((case result of
      Left _ -> true
      Right _ -> false) == (kind == "error")) "header caller result changed"
  value <- liftEffect $ Ref.read counts
  liftEffect $ check (value.started == 4 && value.finished == 4 && value.active == 0 && value.closed == 1) ("header lifecycle not joined: " <> kind)


-- The boundary is observable without storage switches. Both phases use the
-- actual bounded worker dispatcher and the same nested brackets as Driver.
headerPhases :: String -> Aff Unit
headerPhases kind = do
  readCounts <- liftEffect newCounts
  decodeCounts <- liftEffect newCounts
  boundary <- liftEffect (Ref.new false)
  let task phase counts = Just <$> mapPreparationJobs 8 (\index -> bracket
        (liftEffect $ Ref.modify_ (\s -> s { started = s.started + 1, active = s.active + 1 }) counts)
        (\_ -> liftEffect $ Ref.modify_ (\s -> s { finished = s.finished + 1, active = s.active - 1 }) counts)
        (\_ -> do
          delay (Milliseconds (if index == 0 then 20.0 else 45.0))
          liftEffect $ when (kind == phase <> "-error" && index == 0) (throw "phase child failure")
          pure index)) [ 0, 1, 2, 3 ]
  let decoder _ = do
        reached <- liftEffect (Ref.read boundary)
        liftEffect $ check reached "header decode preceded FFI boundary"
        liftEffect $ Ref.modify_ (\s -> s { used = s.used + 1 }) decodeCounts
        task "decode" decodeCounts
  let action = bracket (pure unit)
        (\_ -> liftEffect $ traverse_ (\counts -> do
          value <- Ref.read counts
          check (value.active == 0 && value.started == value.finished) "close preceded phase join"
          Ref.modify_ (\s -> s { closed = s.closed + 1 }) counts) [ readCounts, decodeCounts ])
        (\_ -> Header.withOptional (Just (task "read" readCounts)) \raw -> do
          delay (Milliseconds 5.0)
          before <- liftEffect (Ref.read decodeCounts)
          liftEffect $ check (before.used == 0) "decode started during preparation"
          liftEffect $ when (kind == "read-caller-error") (throw "phase caller failure")
          if kind == "read-cancel" then delay (Milliseconds 200.0)
          else do
            liftEffect (Ref.write true boundary)
            Header.withDecodedUsing decoder raw \pending -> do
              if kind == "read-error" then liftEffect $ check (isNothing pending) "failed read launched decode"
              else if kind == "decode-cancel" then delay (Milliseconds 200.0)
              else if kind == "decode-caller-error" then do
                delay (Milliseconds 5.0)
                liftEffect $ throw "phase caller failure"
              else case pending of
                Nothing -> liftEffect $ throw "decode handle missing"
                Just wait -> do
                  result <- wait
                  liftEffect $ check (isNothing result == (kind == "decode-error")) "decode error changed result")
  if kind == "read-cancel" || kind == "decode-cancel" then do
    fiber <- forkAff action
    delay (Milliseconds (if kind == "read-cancel" then 10.0 else 60.0))
    killFiber (error "phase caller cancelled") fiber
  else do
    result <- attempt action
    liftEffect $ check ((case result of
      Left _ -> true
      Right _ -> false) == (kind == "read-caller-error" || kind == "decode-caller-error")) "phase caller error changed"
  first <- liftEffect (Ref.read readCounts)
  second <- liftEffect (Ref.read decodeCounts)
  let decoded = kind == "success" || kind == "decode-error" || kind == "decode-caller-error" || kind == "decode-cancel"
  liftEffect do
    check (first.started == 4 && first.finished == 4 && first.active == 0 && first.closed == 1) "raw phase did not join"
    check (second.used == (if decoded then 1 else 0) && second.started == (if decoded then 4 else 0) && second.finished == second.started && second.active == 0 && second.closed == 1) ("decode phase did not join: " <> kind)
    log ("HEADER-PHASES-NATIVE-PASS " <> kind)
