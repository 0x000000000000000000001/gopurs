module PreparationProbe (main) where

import Prelude

import Control.Monad.Rec.Class (Step(..), tailRecM)
import Control.Parallel (parTraverse)
import Data.Argonaut.Core (stringify)
import Data.Argonaut.Decode (decodeJson)
import Data.Argonaut.Encode (encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Either (Either(..), hush)
import Data.Int as Int
import Data.List (List(..), (:))
import Data.List as List
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String.CodeUnits as CU
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst, snd)
import Effect (Effect)
import Effect.Aff (Aff, runAff_, throwError)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Exception (error, message)
import Effect.Ref as Ref
import Gopurs.Cache.Wire as Wire
import Gopurs.ClassMetadata (addClassDataDeclarations)
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.Driver.Prepare (prepareModules)
import Gopurs.GlobalTypes (buildGlobalTypes)
import Gopurs.Host.Metrics as Metrics
import Gopurs.Monomorphization (monomorphizeModulesWithDispatch)
import Gopurs.Preparation (runPreparationJobs)
import Node.Encoding (Encoding(..))
import Node.FS.Aff as FS
import Node.Process as Process
import PureScript.Backend.Optimizer.App (coreFnModulesFromOutput, loadDirectives)
import PureScript.Backend.Optimizer.CoreFn (Ann, Module(..))
import PureScript.Backend.Optimizer.CoreFn.Usage (invalidateSourceUsageModule)
import PureScript.Backend.Optimizer.Monomorphize (Instantiation, InstantiationMap, PreparedInstantiation, TransitiveResult, collectInstantiations, transitiveCollectWith)
import Type.Proxy (Proxy(..))

type Index = Array (Tuple String (Tuple String Int))

-- Match Cache.Modules' bounded dynamic batches. Each codec context remains
-- private, and result order is restored after all eight workers finish.
mapJobs :: forall a b. (a -> Aff b) -> Array a -> Aff (Array b)
mapJobs task values
  | Array.null values = pure []
  | otherwise = do
      let count = Array.length values
      let workers = min 8 count
      let chunkSize = max 1 (min 8 ((count + workers * 8 - 1) / (workers * 8)))
      next <- liftEffect (Ref.new 0)
      let worker = tailRecM
            (\results -> do
              start <- liftEffect $ Ref.modify' (\index -> { state: index + chunkSize, value: index }) next
              if start >= count then pure (Done results)
              else do
                completed <- traverse task (Array.slice start (start + chunkSize) values)
                pure (Loop (Tuple start completed : results)))
            Nil
      results <- parTraverse (\_ -> worker) (Array.range 0 (workers - 1))
      pure (Array.concatMap snd (Array.sortWith fst (Array.concatMap Array.fromFoldable results)))

writeSet :: forall a. Wire.Codec a => String -> String -> Array (Tuple String a) -> Aff Unit
writeSet root kind values = do
  let directory = root <> "/" <> kind
  FS.mkdir directory
  index <- Metrics.measure ("probe/encode-write/" <> kind) \_ -> mapJobs
    (\(Tuple name value) -> do
      text <- liftEffect (Wire.write value)
      let digest = Wire.hash text
      FS.writeTextFile UTF8 (directory <> "/" <> Wire.hash name <> ".json") text
      -- Wire's string payloads are hex UTF-16, so its representation is ASCII.
      pure (Tuple name (Tuple digest (CU.length text))))
    values
  FS.writeTextFile UTF8 (directory <> "/index.json") (stringify (encodeJson index))
  liftEffect $ Console.log ("[p00] " <> kind <> " objects=" <> show (Array.length index))

readIndex :: String -> Aff Index
readIndex directory = do
  text <- FS.readTextFile UTF8 (directory <> "/index.json")
  case hush (jsonParser text) >>= hush <<< decodeJson of
    Just index -> pure index
    Nothing -> throwError (error "Invalid probe index")

restoreSet :: forall a. Wire.Codec a => Proxy a -> String -> String -> Aff Unit
restoreSet _ root kind = do
  let directory = root <> "/" <> kind
  reencode <- liftEffect (Process.lookupEnv "P00_REENCODE")
  case reencode of
    Just destination -> FS.mkdir (destination <> "/" <> kind)
    Nothing -> pure unit
  values <- Metrics.measure ("probe/read-validate-decode/" <> kind) \_ -> do
    index <- readIndex directory
    mapJobs
      (\entry@(Tuple name (Tuple digest bytes)) -> do
        text <- FS.readTextFile UTF8 (directory <> "/" <> Wire.hash name <> ".json")
        when (CU.length text /= bytes || Wire.hash text /= digest) $ throwError (error "Probe object integrity")
        case Wire.read text :: Maybe a of
          Just value -> pure (Tuple entry value)
          Nothing -> throwError (error ("Probe codec rejected " <> name)))
      index
  let label = case reencode of
        Just _ -> "probe/reencode-write/"
        Nothing -> "probe/verify-roundtrip/"
  Metrics.measure (label <> kind) \_ -> void $ mapJobs
    (\(Tuple (Tuple name (Tuple digest _)) value) -> do
      text <- liftEffect (Wire.write value)
      unless (Wire.hash text == digest) $ throwError (error ("Roundtrip differs: " <> name))
      case reencode of
        Just destination -> FS.writeTextFile UTF8 (destination <> "/" <> kind <> "/" <> Wire.hash name <> ".json") text
        Nothing -> pure unit)
    values
  case reencode of
    Just destination -> do
      index <- FS.readTextFile UTF8 (directory <> "/index.json")
      FS.writeTextFile UTF8 (destination <> "/" <> kind <> "/index.json") index
    Nothing -> pure unit
  liftEffect $ Console.log ("[p00] " <> kind <> " verified=" <> show (Array.length values))

required :: String -> Aff String
required name = do
  value <- liftEffect (Process.lookupEnv name)
  case value of
    Just text -> pure text
    Nothing -> throwError (error ("Missing " <> name))

run :: Aff Unit
run = do
  liftEffect (Metrics.setMemProfileRate 0)
  action <- required "P00_ACTION"
  root <- required "P00_OBJECTS"
  if action == "restore-source" then restoreSet (Proxy :: Proxy (Module Ann)) root "source"
  else if action == "restore-prepared" then restoreSet (Proxy :: Proxy (Module Ann)) root "prepared"
  else if action == "restore-metadata" then restoreSet (Proxy :: Proxy CodegenMetadata) root "metadata"
  else if action == "restore-direct" then restoreSet (Proxy :: Proxy (Map String Instantiation)) root "direct"
  else if action == "restore-callers" then restoreSet (Proxy :: Proxy InstantiationMap) root "callers"
  else if action == "restore-transitive" then restoreSet (Proxy :: Proxy (Map String Instantiation)) root "transitive"
  else if action == "restore-contributions" then restoreSet (Proxy :: Proxy PreparedInstantiation) root "contributions"
  else do
    input <- required "P00_INPUT"
    countText <- required "P00_COUNT"
    count <- case Int.fromString countText of
      Just value -> pure value
      Nothing -> throwError (error "Invalid P00_COUNT")
    if action == "encode-prepared" || action == "encode-metadata" then do
      directives <- loadDirectives
      prepared <- prepareModules input count 8 Nothing directives
      if action == "encode-metadata" then writeSet root "metadata" [ Tuple "metadata" prepared.metadata ]
      else writeSet root "prepared" (map (\mod@(Module m) -> Tuple (unwrap m.name) mod) (Array.fromFoldable prepared.modules))
    else do
      modules <- Metrics.measure "probe/load source" \_ -> coreFnModulesFromOutput input
      unless (List.length modules == count) $ throwError (error "Incomplete probe corpus")
      if action == "encode-source" then writeSet root "source" (map (\mod@(Module m) -> Tuple (unwrap m.name) mod) (Array.fromFoldable modules))
      else if action == "encode-demands" then do
        finalRound <- liftEffect (Ref.new ([] :: Array TransitiveResult))
        let collect ast direct = do
              writeSet root "direct" (Map.toUnfoldable direct :: Array _)
              callers <- Metrics.measure "probe/caller contributions" \_ -> runPreparationJobs 8
                (map (\mod@(Module m) _ -> Tuple (unwrap m.name)
                  (collectInstantiations ast Map.empty (invalidateSourceUsageModule (addClassDataDeclarations mod))))
                  (Array.fromFoldable modules))
              writeSet root "callers" callers
              transitive <- transitiveCollectWith
                (\tasks -> do
                  results <- runPreparationJobs 8 tasks
                  liftEffect (Ref.write results finalRound)
                  pure results)
                ast direct
              writeSet root "transitive" (Map.toUnfoldable transitive :: Array _)
              results <- liftEffect (Ref.read finalRound)
              writeSet root "contributions" (Array.mapMaybe (case _ of
                Nothing -> Nothing
                Just { key: Tuple name specialization, entry } -> Just (Tuple (name <> "/" <> specialization) entry)) results)
              pure transitive
        prepared <- monomorphizeModulesWithDispatch collect (runPreparationJobs 8)
          (buildGlobalTypes (Array.fromFoldable modules)) (map addClassDataDeclarations modules)
        liftEffect $ Console.log ("[p00] demand oracle modules=" <> show (List.length prepared))
      else throwError (error ("Unknown probe action: " <> action))

main :: Effect Unit
main = runAff_ (case _ of
  Right _ -> pure unit
  Left failure -> do
    Console.error (message failure)
    Process.exit' 1) run
