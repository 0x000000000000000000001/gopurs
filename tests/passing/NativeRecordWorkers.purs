-- @inline export runA never
-- @inline export runB never
-- @inline export runC never
-- @inline export makeA never
-- @inline export makeB never
-- @inline export makeC never
module Main where

-- @dependencies: assert prelude effect console refs
import Prelude
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)
import Worker (score)

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
runA n = score (makeA n)
runB :: Int -> Int
runB n = score (makeB n)
runC :: Int -> Int
runC n = score (makeC n)
main :: Effect Unit
main = do
  seedRef <- Ref.new 5
  seed <- Ref.read seedRef
  aRef <- Ref.new (makeA seed)
  bRef <- Ref.new (makeB seed)
  cRef <- Ref.new (makeC seed)
  a <- Ref.read aRef
  b <- Ref.read bRef
  c <- Ref.read cRef
  assertEqual { expected: 2851, actual: score a }
  assertEqual { expected: -195, actual: score b }
  assertEqual { expected: 2851, actual: score c }
  let a2 = a { id = 9 }
  let b2 = b { name = "alpha" }
  let c2 = c { active = false }
  assertEqual { expected: 4435, actual: score a2 }
  assertEqual { expected: 1057, actual: score b2 }
  assertEqual { expected: 1057, actual: score c2 }
  savedA <- Ref.read aRef
  savedB <- Ref.read bRef
  savedC <- Ref.read cRef
  assertEqual { expected: a, actual: savedA }
  assertEqual { expected: b, actual: savedB }
  assertEqual { expected: c, actual: savedC }
  assertEqual { expected: 42, actual: a2.extra }
  assertEqual { expected: { retained: "keep" }, actual: b2.nested }
  assertEqual { expected: [1,2,3], actual: c2.values }
  assertEqual { expected: false, actual: c2.tail }
  assertEqual { expected: "last", actual: c2.z }
  log "Done"
