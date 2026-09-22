-- @inline export observe never
-- @inline export tailResult never
-- @inline export decodeTracked never
-- @inline export decodeCustom never
module Main where

-- @dependencies: assert prelude effect console refs either maybe argonaut-core argonaut-codecs foreign-object tuples record
import Prelude
import Data.Argonaut.Core (Json, fromNumber, fromObject, jsonNull)
import Data.Argonaut.Decode.Class (class DecodeJson, class GDecodeJson, decodeJson, gDecodeJson)
import Data.Argonaut.Decode.Error (JsonDecodeError(..))
import Data.Either (Either(..))
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Foreign.Object as Object
import Prim.RowList as RL
import Record.Unsafe (unsafeHas)
import Test.Assert (assertEqual)
import Type.Proxy (Proxy(..))

calls :: Ref.Ref (Array Int)
calls = unsafePerformEffect (Ref.new [])

observe :: Int -> Int
observe value = unsafePerformEffect do
  Ref.modify_ (\seen -> seen <> [ value ]) calls
  pure value

newtype Probe = Probe Int

derive instance eqProbe :: Eq Probe
derive newtype instance showProbe :: Show Probe

instance decodeProbe :: DecodeJson Probe where
  decodeJson json = do
    value <- decodeJson json
    let observed = observe value
    if observed < 0 then Left (TypeMismatch "negative probe")
    else Right (Probe observed)

type Tracked = { a :: Probe, b :: Probe, c :: Probe }

decodeTracked :: Json -> Either JsonDecodeError Tracked
decodeTracked = decodeJson

object :: Array (Tuple String Json) -> Json
object = fromObject <<< Object.fromFoldable

-- The tail is intentionally not the standard row-list Nil/Cons decoder. Its
-- shared record must remain immutable when the ordinary Cons instance extends
-- it, and its method must execute only after the head field succeeds.
foreign import data CustomTail :: RL.RowList Type

sharedTail :: { tail :: Int }
sharedTail = { tail: 41 }

tailResult :: Object.Object Json -> Either JsonDecodeError { tail :: Int }
tailResult _ =
  let marker = observe 100
  in if marker == 100 then Right sharedTail else Left MissingValue

instance customTailDecoder :: GDecodeJson (tail :: Int) CustomTail where
  gDecodeJson value _ = tailResult value

decodeCustom :: Object.Object Json -> Either JsonDecodeError { head :: Int, tail :: Int }
decodeCustom value =
  gDecodeJson value (Proxy :: Proxy (RL.Cons "head" Int CustomTail))

main :: Effect Unit
main = do
  Ref.write [] calls
  assertEqual
    { expected: Right { a: Probe 1, b: Probe 2, c: Probe 3 }
    , actual: decodeTracked (object [ Tuple "c" (fromNumber 3.0), Tuple "a" (fromNumber 1.0), Tuple "b" (fromNumber 2.0) ])
    }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 1, 2, 3 ], actual }

  Ref.write [] calls
  assertEqual
    { expected: Left (AtKey "b" (TypeMismatch "negative probe"))
    , actual: decodeTracked (object [ Tuple "a" (fromNumber 1.0), Tuple "b" (fromNumber (-2.0)), Tuple "c" (fromNumber (-3.0)) ])
    }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 1, -2 ], actual }

  Ref.write [] calls
  assertEqual
    { expected: Left (AtKey "b" MissingValue)
    , actual: decodeTracked (object [ Tuple "a" (fromNumber 1.0), Tuple "c" (fromNumber 3.0) ])
    }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 1 ], actual }

  assertEqual
    { expected: Right { a: 7, b: Nothing }
    , actual: decodeJson (object [ Tuple "a" (fromNumber 7.0) ]) :: Either JsonDecodeError { a :: Int, b :: Maybe Int }
    }
  assertEqual
    { expected: Right { a: 7, b: Nothing }
    , actual: decodeJson (object [ Tuple "a" (fromNumber 7.0), Tuple "b" jsonNull ]) :: Either JsonDecodeError { a :: Int, b :: Maybe Int }
    }
  assertEqual
    { expected: Left (AtKey "outer" (AtKey "required" MissingValue))
    , actual: decodeJson (object [ Tuple "outer" (object []) ]) :: Either JsonDecodeError { outer :: { required :: Int } }
    }

  Ref.write [] calls
  first <- Ref.new (decodeCustom (Object.singleton "head" (fromNumber 1.0)))
  second <- Ref.new (decodeCustom (Object.singleton "head" (fromNumber 2.0)))
  Ref.read first >>= \actual -> assertEqual { expected: Right { head: 1, tail: 41 }, actual }
  Ref.read second >>= \actual -> assertEqual { expected: Right { head: 2, tail: 41 }, actual }
  assertEqual { expected: { tail: 41 }, actual: sharedTail }
  assertEqual { expected: false, actual: unsafeHas "head" sharedTail }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 100, 100 ], actual }

  Ref.write [] calls
  assertEqual { expected: Left (AtKey "head" MissingValue), actual: decodeCustom Object.empty }
  Ref.read calls >>= \actual -> assertEqual { expected: [], actual }
  log "Done"
