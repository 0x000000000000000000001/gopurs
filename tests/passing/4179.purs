module Main where

import Prelude

import Data.Maybe (Maybe(..))
import Effect (Effect)
import Effect.Console (log)
import Test.Assert (assertEqual, assert)
import CustomAssert (assertThrows)
import Data.String (contains)
import Data.String.Pattern (Pattern(..))

force :: forall a b. (Unit -> b) -> b
force f = f unit

alpha = { backref: \_ -> bravo, x: 1 }
bravo = force \_ -> alpha.x


complicatedIdentity :: forall a. a -> a
complicatedIdentity = h
  where
  f :: forall a. Int -> { tick :: a -> a, tock :: a -> a }
  f n = { tick: if n <= 0 then identity else (f (n - 1)).tock identity, tock: \a -> g n a }

  g :: forall a. Int -> a -> a
  g = (\bit -> if bit then \n -> (f n).tick else const h) true

  h :: forall a. a -> a
  h = (\n -> (f n).tick) 10


foreign import runtimeImportImpl :: forall a. Maybe String -> (String -> Maybe String) -> String -> (Maybe String -> Effect a) -> Effect a

runtimeImport :: forall a. String -> (Maybe String -> Effect a) -> Effect a
runtimeImport = runtimeImportImpl Nothing Just

type ID = forall a. a -> a

main = do
  err <- assertThrows \_ ->
    let
      selfOwn = { a: 1, b: force \_ -> selfOwn.a }
    in selfOwn
  assert $ contains (Pattern "interface conversion") err || contains (Pattern "Attempt to read property") err

  err2 <- assertThrows \_ ->
    let
      f = (\_ -> { left: g identity, right: h identity }) unit

      g :: ID -> ID
      g x = (j x x x).right

      h :: ID -> ID -> { left :: ID, right :: ID }
      h x = j x x

      j x y z = { left: x y z, right: f.left }
    in f
  assert $ contains (Pattern "interface conversion") err2 || contains (Pattern "Attempt to read property") err2

  assertEqual { actual: bravo, expected: 1 }
  runtimeImport "InitializationError" \err3 -> do
    assertEqual { actual: (err3 /= Nothing), expected: true }
    log "Done"
