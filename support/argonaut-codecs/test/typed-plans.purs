module Main where

-- @dependencies: assert prelude effect console refs either maybe argonaut-core argonaut-codecs foreign-object tuples arrays numbers foldable-traversable partial
import Prelude
import Data.Argonaut.Core (Json, fromArray, fromNumber, fromObject, fromString, jsonNull)
import Data.Argonaut.Decode.Class (class DecodeJson, decodeJson)
import Data.Argonaut.Decode ((.:), (.:?))
import Data.Argonaut.Decode.Decoders as D
import Data.Argonaut.Decode.Error (JsonDecodeError(..))
import Data.Argonaut.Decode.Parser (decodeJsonString, decodeJsonStringWith, parseJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Either (Either(..))
import Data.Foldable (for_)
import Data.Maybe (Maybe(..))
import Data.Number as Number
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Foreign.Object as Object
import Partial.Unsafe (unsafeCrashWith)
import Test.Assert (assert, assertEqual)

type Leaf = { count :: Int, note :: Maybe String }
type Nested = { groups :: Array (Array Leaf), optional :: Maybe Leaf }

parse :: String -> Json
parse text = case jsonParser text of
  Left err -> unsafeCrashWith err
  Right json -> json

object :: Array (Tuple String Json) -> Json
object = fromObject <<< Object.fromFoldable

decodeNested :: Json -> Either JsonDecodeError Nested
decodeNested = decodeJson

-- Custom element decoding must retain its transformation and exact error.
custom :: Json -> Either JsonDecodeError Int
custom json = do
  n <- D.decodeInt json
  if n < 0 then Left (TypeMismatch "positive") else Right (n + 10)

data Wrapped = Wrapped Int
derive instance eqWrapped :: Eq Wrapped
instance showWrapped :: Show Wrapped where
  show (Wrapped n) = "Wrapped " <> show n
instance decodeWrapped :: DecodeJson Wrapped where
  decodeJson = map Wrapped <<< custom

type Mixed = { rows :: Array { p :: Wrapped, q :: Maybe (Array (Maybe Int)) } }

decodeMixed :: Json -> Either JsonDecodeError Mixed
decodeMixed = decodeJson

type Owned = { json :: Json, obj :: Object.Object Json }

decodeOwned :: Json -> Either JsonDecodeError Owned
decodeOwned = decodeJson

decodeOwnedText :: String -> Either JsonDecodeError Owned
decodeOwnedText = decodeJsonStringWith decodeOwned

decodeJSONFields :: Json -> Either JsonDecodeError { json :: Json, name :: String }
decodeJSONFields = decodeJson

decodeJSONFieldsText :: String -> Either JsonDecodeError { json :: Json, name :: String }
decodeJSONFieldsText = decodeJsonStringWith decodeJSONFields

data Action = Label String (Maybe Int) | Batch (Array Leaf)
derive instance eqAction :: Eq Action
instance showAction :: Show Action where
  show (Label text count) = "Label " <> show text <> " " <> show count
  show (Batch rows) = "Batch " <> show rows
instance decodeAction :: DecodeJson Action where
  decodeJson json = do
    obj <- decodeJson json
    discriminator <- obj .: "kind"
    case discriminator of
      "label" -> Label <$> obj .: "text" <*> obj .:? "count"
      "batch" -> Batch <$> obj .: "rows"
      _ -> Left (TypeMismatch "Action kind")

decodeActions :: Json -> Either JsonDecodeError { actions :: Array Action }
decodeActions = decodeJson

data Checked = Checked
derive instance eqChecked :: Eq Checked
instance showChecked :: Show Checked where
  show _ = "Checked"
instance decodeChecked :: DecodeJson Checked where
  decodeJson json = do
    obj <- decodeJson json
    _ <- obj .: "value" :: Either JsonDecodeError String
    pure Checked

data EmptyObject = EmptyObject
derive instance eqEmptyObject :: Eq EmptyObject
instance showEmptyObject :: Show EmptyObject where
  show _ = "EmptyObject"
instance decodeEmptyObject :: DecodeJson EmptyObject where
  decodeJson json = do
    _ <- decodeJson json :: Either JsonDecodeError (Object.Object Json)
    pure EmptyObject

main :: Effect Unit
main = do
  let ownedInput = "{\"json\":{\"inner\":[\"owned-é🙂\"]},\"name\":\"kept\",\"obj\":{\"keep\":\"owned-object\"}}"
  assert $ decodeOwnedText ownedInput == (parseJson ownedInput >>= decodeOwned)
  assert $ decodeJSONFieldsText ownedInput == (parseJson ownedInput >>= decodeJSONFields)
  for_ [ "{\"groups\":[[{\"count\":7}]],\"optional\":null}"
       , "{\"groups\":[],\"groups\":[[{\"count\":2,\"co\\u0075nt\":3,\"note\":\"é🙂\"}]]}"
       , "{\"groups\":[[{\"note\":false}]],\"optional\":false}"
       , "{\"groups\":[],\"ignored\":{\"x\":1e400}}"
       , "{\"groups\":false,\"ignored\":\"\\q\"}"
       , "{\"groups\":false} trailing", "[", "null", "true", "{}"
       ] \text -> assertEqual { expected: parseJson text >>= decodeNested, actual: decodeJsonStringWith decodeNested text }
  for_ [ "{\"actions\":[{\"kind\":\"label\",\"text\":\"saved\",\"count\":null},{\"kind\":\"batch\",\"rows\":[{\"count\":9}]}]}"
       , "{\"actions\":[{\"kind\":\"unknown\",\"text\":false}]}"
       , "{\"actions\":[{\"count\":false,\"kind\":\"label\"}]}"
       , "{\"actions\":[{\"kind\":\"label\",\"text\":false,\"kind\":\"batch\",\"rows\":[]}]}"
       , "{\"actions\":[{\"kind\":\"batch\",\"rows\":[{}]}]}"
       ] \text -> assertEqual { expected: parseJson text >>= decodeActions, actual: decodeJsonStringWith decodeActions text }
  for_ [ "{\"rows\":[{\"p\":2,\"q\":[null,3]}]}", "{\"rows\":[{\"p\":-1,\"q\":false}]}" ] \text ->
    assertEqual { expected: parseJson text >>= decodeMixed, actual: decodeJsonStringWith decodeMixed text }
  assertEqual { expected: Right { count: 3, note: Nothing }, actual: decodeJsonString "{\"count\":3}" :: Either JsonDecodeError Leaf }
  assertEqual { expected: Right 13, actual: decodeJsonStringWith custom "3" }
  assertEqual { expected: Left (TypeMismatch "JSON"), actual: decodeJsonStringWith custom "3 trailing" }
  for_ [ "{\"value\":{\"a\":1,\"b\":[\"owned\",null],\"a\":2}}", "{\"value\":null}" ] \text -> do
    let original = parseJson text >>= decodeJson :: Either JsonDecodeError { value :: Json }
    assert $ (decodeJsonString text :: Either JsonDecodeError { value :: Json }) == original
    let originalObject = parseJson text >>= decodeJson :: Either JsonDecodeError { value :: Object.Object Json }
    assert $ (decodeJsonString text :: Either JsonDecodeError { value :: Object.Object Json }) == originalObject
  assertEqual
    { expected: Right { action: Label "single" Nothing }
    , actual: decodeJson (parse "{\"action\":{\"kind\":\"label\",\"text\":\"single\"}}") :: Either JsonDecodeError { action :: Action }
    }
  assertEqual
    { expected: Right { checked: Checked, empty: EmptyObject }
    , actual: decodeJson (parse "{\"checked\":{\"value\":\"discarded\"},\"empty\":{}}") :: Either JsonDecodeError { checked :: Checked, empty :: EmptyObject }
    }
  assertEqual
    { expected: Left (AtKey "checked" (AtKey "value" (TypeMismatch "String")))
    , actual: decodeJson (parse "{\"checked\":{\"value\":false}}") :: Either JsonDecodeError { checked :: Checked }
    }
  for_ [ parse "{\"actions\":[{\"kind\":\"label\",\"text\":\"saved\",\"count\":null},{\"kind\":\"batch\",\"rows\":[{\"count\":9}]}]}", object
    [ Tuple "actions" (fromArray [ object [ Tuple "kind" (fromString "label"), Tuple "text" (fromString "saved"), Tuple "count" jsonNull ], object [ Tuple "kind" (fromString "batch"), Tuple "rows" (fromArray [ object [ Tuple "count" (fromNumber 9.0) ] ]) ] ]) ] ] \input ->
    assertEqual { expected: Right { actions: [ Label "saved" Nothing, Batch [ { count: 9, note: Nothing } ] ] }, actual: decodeActions input }
  for_ [ Tuple "{\"kind\":\"label\",\"count\":false}" (AtKey "text" MissingValue)
       , Tuple "{\"kind\":\"label\",\"text\":\"x\",\"count\":false}" (AtKey "count" (TypeMismatch "Number"))
       , Tuple "{\"kind\":\"unknown\"}" (TypeMismatch "Action kind")
       , Tuple "{\"kind\":\"batch\",\"rows\":[{}]}" (AtKey "rows" (Named "Array" (AtIndex 0 (AtKey "count" MissingValue))))
       , Tuple "null" (TypeMismatch "Object")
       ] \(Tuple text err) -> assertEqual
        { expected: Left (AtKey "actions" (Named "Array" (AtIndex 0 err)))
        , actual: decodeActions (parse ("{\"actions\":[" <> text <> "]}"))
        }
  let expected = Right { groups: [ [ { count: 7, note: Nothing } ], [] ], optional: Just { count: 9, note: Just "kept" } }
  let raw = parse "{\"groups\":[[{\"count\":7}],[]],\"optional\":{\"count\":9,\"note\":\"kept\"}}"
  let boxed = object
        [ Tuple "groups" (fromArray [ fromArray [ object [ Tuple "count" (fromNumber 7.0) ] ], fromArray [] ])
        , Tuple "optional" (object [ Tuple "count" (fromNumber 9.0), Tuple "note" (fromString "kept") ])
        ]
  for_ [ raw, boxed ] \input -> assertEqual { expected, actual: decodeNested input }
  retained <- Ref.new (decodeNested raw)
  assertEqual
    { expected: Right { groups: [], optional: Nothing }
    , actual: decodeNested (parse "{\"groups\":[]}")
    }
  Ref.read retained >>= \actual -> assertEqual { expected, actual }
  assertEqual
    { expected: Left (AtKey "groups" (Named "Array" (AtIndex 0 (Named "Array" (AtIndex 0 (AtKey "count" (TypeMismatch "Number")))))))
    , actual: decodeNested (parse "{\"groups\":[[{\"count\":\"bad\",\"note\":false}]],\"optional\":42}")
    }
  assertEqual
    { expected: Left (AtKey "optional" (AtKey "count" MissingValue))
    , actual: decodeNested (parse "{\"groups\":[],\"optional\":{}}")
    }
  assertEqual { expected: Right {}, actual: decodeJson (object []) :: Either JsonDecodeError {} }
  assertEqual { expected: Left (TypeMismatch "Object"), actual: decodeJson jsonNull :: Either JsonDecodeError {} }
  assertEqual
    { expected: Right { value: Just {} }
    , actual: decodeJson (parse "{\"value\":{}}") :: Either JsonDecodeError { value :: Maybe {} }
    }

  let fields = Object.fromFoldable
        [ Tuple "leaf" (object [ Tuple "count" (fromNumber 3.0) ])
        , Tuple "null" jsonNull
        , Tuple "ints" (fromArray [ fromNumber 1.0, fromNumber 2.0 ])
        ]
  assertEqual { expected: Right { count: 3, note: Nothing }, actual: D.getField decodeJson fields "leaf" :: Either JsonDecodeError Leaf }
  assertEqual { expected: Right Nothing, actual: D.getFieldOptional decodeJson fields "absent" :: Either JsonDecodeError (Maybe Leaf) }
  assertEqual { expected: Right (Just { count: 3, note: Nothing }), actual: D.getFieldOptional decodeJson fields "leaf" :: Either JsonDecodeError (Maybe Leaf) }
  assertEqual { expected: Left (AtKey "null" (TypeMismatch "Object")), actual: D.getFieldOptional decodeJson fields "null" :: Either JsonDecodeError (Maybe Leaf) }
  assertEqual { expected: Right Nothing, actual: D.getFieldOptional' decodeJson fields "null" :: Either JsonDecodeError (Maybe Leaf) }
  assertEqual { expected: Right [ 1, 2 ], actual: D.getField decodeJson fields "ints" :: Either JsonDecodeError (Array Int) }
  assertEqual { expected: Right [ 11, 12 ], actual: D.getField (D.decodeArray custom) fields "ints" }
  assertEqual { expected: Left (Named "Array" (AtIndex 1 (TypeMismatch "positive"))), actual: D.decodeArray custom (parse "[1,-2,-3]") }
  for_ [ parse "{\"rows\":[{\"p\":2,\"q\":[null,3]}]}", object [ Tuple "rows" (fromArray [ object [ Tuple "p" (fromNumber 2.0), Tuple "q" (fromArray [ jsonNull, fromNumber 3.0 ]) ] ]) ] ] \input ->
    assertEqual
      { expected: Right { rows: [ { p: Wrapped 12, q: Just [ Nothing, Just 3 ] } ] }
      , actual: decodeJson input :: Either JsonDecodeError Mixed
      }
  assertEqual
    { expected: Left (AtKey "rows" (Named "Array" (AtIndex 0 (AtKey "p" (TypeMismatch "positive")))))
    , actual: decodeJson (parse "{\"rows\":[{\"p\":-1,\"q\":false},{\"p\":-2}]}") :: Either JsonDecodeError Mixed
    }

  let jsonFields = Object.fromFoldable
        [ Tuple "n" (fromNumber 7.0)
        , Tuple "nested" (fromArray [ jsonNull, object [ Tuple "text" (fromString "kept") ] ])
        ]
  for_ [ fromObject jsonFields, parse "{\"n\":7,\"nested\":[null,{\"text\":\"kept\"}]}" ] \input -> do
    assert $ (decodeJson input :: Either JsonDecodeError (Object.Object Json)) == Right jsonFields
    assert $ (decodeJson (object [ Tuple "value" input ]) :: Either JsonDecodeError { value :: Json }) == Right { value: input }
    assert $ (decodeJson (fromArray [ input, jsonNull ]) :: Either JsonDecodeError (Array (Maybe Json))) == Right [ Just input, Nothing ]
  assertEqual
    { expected: Right (Object.singleton "p" (Wrapped 12))
    , actual: decodeJson (parse "{\"p\":2}") :: Either JsonDecodeError (Object.Object Wrapped)
    }
  assertEqual
    { expected: Left (Named "ForeignObject" (TypeMismatch "positive"))
    , actual: decodeJson (parse "{\"p\":-2}") :: Either JsonDecodeError (Object.Object Wrapped)
    }
  assert $ (decodeJson jsonNull :: Either JsonDecodeError (Object.Object Json)) == Left (TypeMismatch "Object")

  -- Fused public accessors must retain the difference between missing, null,
  -- and a present Maybe value, on native/compact and constructed JSON alike.
  for_ [ parse "{\"n\":3,\"null\":null,\"bad\":false,\"array\":[1,-2,-3]}", object
    [ Tuple "n" (fromNumber 3.0), Tuple "null" jsonNull
    , Tuple "bad" (parse "false"), Tuple "array" (fromArray [ fromNumber 1.0, fromNumber (-2.0), fromNumber (-3.0) ])
    ] ] \json -> case D.decodeJObject json of
      Left err -> unsafeCrashWith (show err)
      Right obj -> do
        assertEqual { expected: Right 3, actual: D.getField decodeJson obj "n" :: Either JsonDecodeError Int }
        assertEqual { expected: Left (AtKey "missing" MissingValue), actual: D.getField decodeJson obj "missing" :: Either JsonDecodeError Int }
        assertEqual { expected: Left (AtKey "null" (TypeMismatch "Number")), actual: D.getField decodeJson obj "null" :: Either JsonDecodeError Int }
        assertEqual { expected: Right Nothing, actual: D.getFieldOptional decodeJson obj "missing" :: Either JsonDecodeError (Maybe Int) }
        assertEqual { expected: Left (AtKey "null" (TypeMismatch "Number")), actual: D.getFieldOptional decodeJson obj "null" :: Either JsonDecodeError (Maybe Int) }
        assertEqual { expected: Right (Just Nothing), actual: D.getFieldOptional decodeJson obj "null" :: Either JsonDecodeError (Maybe (Maybe Int)) }
        assertEqual { expected: Right Nothing, actual: D.getFieldOptional' decodeJson obj "null" :: Either JsonDecodeError (Maybe (Maybe Int)) }
        assertEqual { expected: Right (Just 3), actual: D.getFieldOptional' decodeJson obj "n" :: Either JsonDecodeError (Maybe Int) }
        assertEqual { expected: Left (AtKey "bad" (TypeMismatch "Number")), actual: D.getFieldOptional' decodeJson obj "bad" :: Either JsonDecodeError (Maybe Int) }
        assertEqual { expected: Right Nothing, actual: D.getFieldOptional custom obj "missing" }
        assertEqual { expected: Right Nothing, actual: D.getFieldOptional' custom obj "null" }
        assertEqual { expected: Right (Just 13), actual: D.getFieldOptional custom obj "n" }
        assertEqual { expected: Left (AtKey "array" (Named "Array" (AtIndex 1 (TypeMismatch "positive")))), actual: D.getField (D.decodeArray custom) obj "array" }

  -- Compare tagged access with the ordinary backend decoder at numeric edges.
  -- This checks preservation of that backend's Int policy, including values
  -- outside the portable 32-bit range, without imposing a new numeric policy.
  for_ [ -2147483648.0, 2147483647.0, 2147483648.0, -2147483649.0, 9007199254740991.0, -0.0, 1.5, Number.nan, Number.infinity ] \n -> do
    let json = fromNumber n
    let expectedInt = case D.decodeInt json of
          Left err -> Left (AtKey "n" err)
          Right value -> Right value
    assertEqual { expected: expectedInt, actual: D.getField decodeJson (Object.singleton "n" json) "n" :: Either JsonDecodeError Int }
  log "typed plans: Done"
