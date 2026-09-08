-- @inline export shapeZero never
-- @inline export shapeOne never
-- @inline export shapeTwo never
-- @inline export shapeThree never
-- @inline export shapeFour never
-- @inline export shapeFive never
-- @inline export shapeSix never
-- @inline export makeFour never
-- @inline export produceFour never
module Main where

-- @dependencies: assert arrays prelude effect console refs
-- @snapshot-ffi

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

foreign import data Opaque :: Type
foreign import opaqueValue :: Effect Opaque

-- Native closed-record arguments must be boxed at these FFI call sites.
-- Unlike CompactRecordConsumers, no FFI function constructs the records.
foreign import nativeShape :: forall a. a -> String
foreign import nativeMap :: forall a. a -> String

type Zero = {}
type One = { count :: Int }
type Two = { count :: Int, label :: String }
type Three = { count :: Int, flag :: Boolean, number :: Number }
type Four = { child :: Two, count :: Int, label :: String, values :: Array Int }
type Five = { child :: Two, count :: Int, flag :: Boolean, label :: String, payload :: Opaque }
type Six = { child :: Two, count :: Int, flag :: Boolean, label :: String, number :: Number, values :: Array Int }

-- Each snapshot must contain a native struct parameter, then one orig := row
-- and CodeGen's RecordDictN call. Printer's boxed literal path is insufficient.
shapeZero :: Zero -> String
shapeZero row = nativeShape row

shapeOne :: One -> String
shapeOne row = nativeShape row

shapeTwo :: Two -> String
shapeTwo row = nativeShape row

shapeThree :: Three -> String
shapeThree row = nativeShape row

shapeFour :: Four -> String
shapeFour row = nativeShape row

shapeFive :: Five -> String
shapeFive row = nativeShape row

shapeSix :: Six -> String
shapeSix row = nativeShape row

makeFour :: Int -> Four
makeFour count = { child: { count: 2, label: "nested" }, count, label: "four", values: [-3, 0, 7] }

check :: String -> String -> String -> Effect Unit
check label expected actual = do
  assertEqual { expected, actual }
  log (label <> ": " <> actual)

produceFour :: Ref.Ref (Array String) -> Int -> Effect Four
produceFour trace count = do
  Ref.modify_ (_ <> ["produce"]) trace
  -- The snapshot, separately from the effect trace, checks this call is once.
  pure (makeFour count)

main :: Effect Unit
main = do
  seedRef <- Ref.new 5
  seed <- Ref.read seedRef
  payload <- opaqueValue
  oneRef <- Ref.new ({ count: seed } :: One)
  twoRef <- Ref.new ({ label: "alpha", count: seed } :: Two)
  threeRef <- Ref.new ({ number: -1.25, flag: true, count: seed } :: Three)
  fourRef <- Ref.new (makeFour seed)
  fiveRef <- Ref.new
    ({ payload, label: "five", flag: true, count: seed, child: { count: 2, label: "nested" } } :: Five)
  sixRef <- Ref.new
    ({ values: [-3, 0, 7], number: -1.25, label: "six", flag: true, count: seed, child: { count: 2, label: "nested" } } :: Six)
  one <- Ref.read oneRef
  two <- Ref.read twoRef
  three <- Ref.read threeRef
  four <- Ref.read fourRef
  five <- Ref.read fiveRef
  six <- Ref.read sixRef

  -- Seven representation assertions: helpers 0..5 and the generic fallback.
  check "shape 0" "compact0:" (shapeZero {})
  check "shape 1" "compact1:count" (shapeOne one)
  check "shape 2" "compact2:count|label" (shapeTwo two)
  check "shape 3" "compact3:count|flag|number" (shapeThree three)
  check "shape 4" "compact4:child|count|label|values" (shapeFour four)
  check "shape 5" "compact5:child|count|flag|label|payload" (shapeFive five)
  check "shape 6" "generic6:child|count|flag|label|number|values" (shapeSix six)

  -- The map bridge must preserve boxed field values and their scalar tags.
  check "map 0" "{}" (nativeMap {})
  check "map 1" "{count=i:5}" (nativeMap one)
  check "map 2" "{count=i:5,label=s:alpha}" (nativeMap two)
  check "map 3" "{count=i:5,flag=b:true,number=n:-1.25}" (nativeMap three)
  check "map 4" "{child={count=i:2,label=s:nested},count=i:5,label=s:four,values=[i:-3,i:0,i:7]}" (nativeMap four)
  check "map 5" "{child={count=i:2,label=s:nested},count=i:5,flag=b:true,label=s:five,payload=s:opaque}" (nativeMap five)
  check "map 6" "{child={count=i:2,label=s:nested},count=i:5,flag=b:true,label=s:six,number=n:-1.25,values=[i:-3,i:0,i:7]}" (nativeMap six)

  -- Each update is consumed as a whole record, preventing accessor folding.
  check "update 1" "{count=i:-7}" (nativeMap (one { count = -7 }))
  check "update 2" "{count=i:5,label=s:}" (nativeMap (two { label = "" }))
  check "update 3" "{count=i:5,flag=b:false,number=n:2.5}" (nativeMap (three { flag = false, number = 2.5 }))
  check "update 4" "{child={count=i:2,label=s:nested},count=i:5,label=s:four,values=[]}" (nativeMap (four { values = [] }))
  check "update 5" "{child={count=i:2,label=s:nested},count=i:5,flag=b:false,label=s:five,payload=s:opaque}" (nativeMap (five { flag = false }))
  check "update 6" "{child={count=i:2,label=s:nested},count=i:5,flag=b:true,label=s:,number=n:2.5,values=[i:-3,i:0,i:7]}" (nativeMap (six { label = "", number = 2.5 }))

  -- Retain and reread every original after the updates above.
  savedOne <- Ref.read oneRef
  savedTwo <- Ref.read twoRef
  savedThree <- Ref.read threeRef
  savedFour <- Ref.read fourRef
  savedFive <- Ref.read fiveRef
  savedSix <- Ref.read sixRef
  check "retained 1" "{count=i:5}" (nativeMap savedOne)
  check "retained 2" "{count=i:5,label=s:alpha}" (nativeMap savedTwo)
  check "retained 3" "{count=i:5,flag=b:true,number=n:-1.25}" (nativeMap savedThree)
  check "retained 4" "{child={count=i:2,label=s:nested},count=i:5,label=s:four,values=[i:-3,i:0,i:7]}" (nativeMap savedFour)
  check "retained 5" "{child={count=i:2,label=s:nested},count=i:5,flag=b:true,label=s:five,payload=s:opaque}" (nativeMap savedFive)
  check "retained 6" "{child={count=i:2,label=s:nested},count=i:5,flag=b:true,label=s:six,number=n:-1.25,values=[i:-3,i:0,i:7]}" (nativeMap savedSix)

  check "native field access" "5:true:-1.25:nested:[-3,0,7]"
    (show three.count <> ":" <> show three.flag <> ":" <> show three.number <> ":" <> four.child.label <> ":" <> show four.values)
  check "nested update" "{child={count=i:2,label=s:changed},count=i:5,label=s:four,values=[i:-3,i:0,i:7]}"
    (nativeMap (four { child { label = "changed" } }))
  check "nested original" "{child={count=i:2,label=s:nested},count=i:5,label=s:four,values=[i:-3,i:0,i:7]}" (nativeMap four)

  trace <- Ref.new []
  Ref.modify_ (_ <> ["before"]) trace
  produced <- produceFour trace seed
  Ref.modify_ (_ <> ["consume:" <> shapeFour produced]) trace
  Ref.modify_ (_ <> ["after"]) trace
  actualTrace <- Ref.read trace
  assertEqual
    { expected: ["before", "produce", "consume:compact4:child|count|label|values", "after"]
    , actual: actualTrace
    }
  log ("effect order: " <> show actualTrace)
  log "Done"
