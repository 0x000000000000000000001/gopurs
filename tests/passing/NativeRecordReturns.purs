-- @inline export runA never
-- @inline export runB never
-- @inline export runC never
-- @inline export callDynamic never
-- @inline export firstSummary never
module Main where

-- @dependencies: assert prelude effect console refs either
import Prelude
import Data.Either (Either(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)
import Worker (decode)

type A = { id :: Int, name :: String, active :: Boolean, extra :: Int }
type B = { before :: Number, id :: Int, name :: String, active :: Boolean, nested :: { retained :: String } }
type C = { active :: Boolean, id :: Int, name :: String, values :: Array Int, tail :: Boolean, z :: String }
makeA :: Int -> A
makeA n = { id: n, name: "alpha", active: true, extra: 42 }
makeB :: Int -> B
makeB n = { before: 1.25, id: n, name: "beta", active: false, nested: { retained: "keep" } }
makeC :: Int -> C
makeC n = { active: true, id: n, name: "alpha", values: [1,2,3], tail: false, z: "last" }

runA :: Int -> Int
runA n = case decode (makeA n) of
  Left _ -> -1
  Right s -> s.id * 3 + (if s.active then 7 else 0)

runB :: Int -> Int
runB n = case decode (makeB n) of
  Left _ -> -1
  Right s -> s.id * 3 + (if s.active then 7 else 0)

runC :: Int -> Int
runC n = case decode (makeC n) of
  Left _ -> -1
  Right s -> s.id * 3 + (if s.active then 7 else 0)

-- A dynamic call still goes through the boxed wrapper: the record result is
-- converted at that boundary and must remain correct there.
callDynamic :: forall a. (a -> Either String { id :: Int, name :: String, active :: Boolean }) -> a -> String
callDynamic f value = case f value of
  Left e -> e
  Right s -> s.name

-- Returning the decoded record forces the payload to cross a function
-- boundary: native by value here, boxed then unboxed with the ordinary ABI.
firstSummary :: Int -> { id :: Int, name :: String, active :: Boolean }
firstSummary n = case decode (makeA n) of
  Left _ -> { id: -1, name: "", active: false }
  Right s -> s

main :: Effect Unit
main = do
  assertEqual { expected: 25, actual: runA 5 }
  assertEqual { expected: 18, actual: runB 5 }
  assertEqual { expected: 25, actual: runC 5 }
  assertEqual { expected: -1, actual: runA (-1) }
  assertEqual { expected: -1, actual: runB (-1) }
  assertEqual { expected: -1, actual: runC (-1) }
  let a = makeA 5
  let b = makeB 5
  let c = makeC 5
  let a2 = a { id = 9 }
  let b2 = b { name = "alpha" }
  let c2 = c { active = false }
  assertEqual { expected: 37, actual: runA 9 }
  assertEqual { expected: 18, actual: runB 5 }
  assertEqual { expected: 40, actual: runC 10 }
  assertEqual { expected: 42, actual: a2.extra }
  assertEqual { expected: { retained: "keep" }, actual: b2.nested }
  assertEqual { expected: [1,2,3], actual: c2.values }
  assertEqual { expected: "alpha", actual: a2.name }
  assertEqual { expected: "negative", actual: callDynamic decode (makeA (-3)) }
  assertEqual { expected: "beta", actual: callDynamic decode (makeB 3) }
  saved <- Ref.new (decode (makeA 3))
  held <- Ref.read saved
  let heldName = case held of
        Right s -> s.name
        Left _ -> ""
  assertEqual { expected: "alpha", actual: heldName }
  let first = firstSummary 7
  assertEqual { expected: 8, actual: first.id }
  assertEqual { expected: "alpha", actual: first.name }
  assertEqual { expected: true, actual: first.active }
  let missing = firstSummary (-1)
  assertEqual { expected: -1, actual: missing.id }
  assertEqual { expected: false, actual: missing.active }
  log "Done"
