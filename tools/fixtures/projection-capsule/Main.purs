module Test.ProjectionCapsuleNative where

import Prelude
import Data.Array as Array
import Data.Either (Either(..))
import Data.Lazy (defer)
import Data.List as List
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.Traversable (traverse, traverse_)
import Data.Time.Duration (Milliseconds(..))
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Aff (Aff, attempt, bracket, delay, launchAff_)
import Effect.Class (liftEffect)
import Effect.Ref as Ref
import Effect.Console (log)
import Effect.Exception (throw)
import Gopurs.Cache.LookupTable as Table
import Gopurs.Cache.Wire as Wire
import Gopurs.Preparation (mapPreparationJobs)
import Gopurs.Driver.Prepare.Projection (Projection)
import Gopurs.Driver.Prepare.ProjectionCapsule as Capsule
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

check :: Boolean -> String -> Effect Unit
check passed reason = unless passed (throw reason)

main :: Effect Unit
main = launchAff_ do
  fragments <- liftEffect capsuleCase
  traverse_ (dispatcherCase fragments) [ 1, 8 ]
  liftEffect $ log "PROJECTION-CAPSULE-NATIVE-PASS"

capsuleCase :: Effect (Array Capsule.Fragment)
capsuleCase = do
  let
    globals = Map.fromFoldable [ Tuple "A.present" Int, Tuple "B.new" String ]
    -- Keep all eleven other metadata fields
    -- concrete so the native representation is the same twelve-field record.
    projection :: Projection
    projection =
      { modules: List.Nil, interfaces: Map.singleton "A" "interface", mainModules: [ "A" ]
      , metadata:
          { classDeclsFields: Map.empty, ctorTypes: Map.empty, elidedCtors: Set.singleton "Elided"
          , enumAdts: Set.singleton "Enum", enumCtors: Set.singleton "Enum.Value"
          , ffiFunctions: Map.empty, globalFunctions: Map.empty
          , globalTypes: Table.fromMap globals
          , pointerAdtLeaves: Map.empty, pointerAdtNodes: Set.singleton "Pointer", pointerAdtPaths: Map.empty
          , reboxFields: Map.empty
          }
      }
    coverage = Set.fromFoldable [ "A.present", "A.absent" ]
    fragments = Capsule.create coverage projection
  bytes <- traverse Wire.write fragments
  case traverse Wire.read bytes >>= flip Capsule.open (defer \_ -> Right globals) of
    Nothing -> throw "capsule wire/open failed"
    Just opened -> do
      check (Table.lookup "A.present" opened.projection.metadata.globalTypes == Just Int) "covered present"
      check (Table.lookup "A.absent" opened.projection.metadata.globalTypes == Nothing) "covered absence"
      check (Table.lookup "B.new" opened.projection.metadata.globalTypes == Just String) "new present"
      check (Table.lookup "B.absent" opened.projection.metadata.globalTypes == Nothing) "new absence"
      before <- Wire.write projection.metadata
      after <- Wire.write opened.projection.metadata
      check (before == after) "full metadata bytes changed"
      repeated <- traverse Wire.write (Capsule.create coverage opened.projection)
      check (bytes == repeated) "capsule bytes changed after open/create"
      check (opened.projection.interfaces == projection.interfaces && opened.projection.mainModules == projection.mainModules) "projection headers changed"
      pure fragments

-- Exercise the production Aff dispatcher through the native compiler. Delayed
-- failure lets sibling jobs start; bracket releases must finish before retry.
dispatcherCase :: Array Capsule.Fragment -> Int -> Aff Unit
dispatcherCase original jobs = do
  let fragments = Array.concatMap (\fragment -> map (\index -> fragment { ordinal = index, count = 8 }) (Array.range 0 7)) original
  expected <- liftEffect $ traverse Wire.write fragments
  action <- pure $ mapPreparationJobs jobs (liftEffect <<< Wire.write) fragments
  first <- action
  second <- action
  liftEffect $ check (first == expected && second == expected) "dispatcher bytes/order/rerun changed"
  state <- liftEffect $ Ref.new { started: 0, finished: 0, active: 0, maximum: 0 }
  failed <- attempt $ mapPreparationJobs jobs (encodeWithFailure state) fragments
  observed <- liftEffect $ Ref.read state
  liftEffect do
    check (case failed of
      Left _ -> true
      Right _ -> false) "encoding failure escaped attempt"
    check (observed.active == 0 && observed.started == observed.finished) "encoding failure left unjoined jobs"
    check (observed.maximum == jobs) "configured encoder concurrency was not exercised"
  -- The caller publishes only the successful result after the failed join.
  retried <- mapPreparationJobs jobs (liftEffect <<< Wire.write) fragments
  liftEffect do
    check (retried == expected) "encoding retry bytes differ"
    log ("CAPSULE-DISPATCH-NATIVE-PASS jobs=" <> show jobs <> " started=" <> show observed.started <> " finished=" <> show observed.finished <> " active=" <> show observed.active)
  where
  encodeWithFailure state fragment = bracket
    (liftEffect $ Ref.modify_ (\s -> s { started = s.started + 1, active = s.active + 1, maximum = max s.maximum (s.active + 1) }) state)
    (\_ -> liftEffect $ Ref.modify_ (\s -> s { finished = s.finished + 1, active = s.active - 1 }) state)
    (\_ -> do
      delay (Milliseconds (if fragment.ordinal == 0 then 10.0 else 50.0))
      liftEffect do
        when (fragment.ordinal == 0) (throw "injected capsule encoder failure")
        Wire.write fragment)
