-- @inline Main.mark never
-- @inline Main.recordFactory never
-- @inline Main.failRead never
module Main where

-- @dependencies: prelude effect console refs arrays foldable-traversable either exceptions

import Prelude
import Data.Array as Array
import Data.Either (Either(..))
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Exception (throw, try)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

type Eight = { a :: Int, b :: Int, c :: Int, d :: Int, e :: Int, f :: Int, g :: Int, h :: Int }
type Trace = Ref.Ref (Array Int)

makeEight :: Int -> Eight
makeEight seed = { a: seed, b: seed + 1, c: seed + 2, d: seed + 3, e: seed + 4, f: seed + 5, g: seed + 6, h: seed + 7 }

fields :: Eight -> Array Int
fields r = [ r.a, r.b, r.c, r.d, r.e, r.f, r.g, r.h ]

readSum :: Ref.Ref Eight -> Effect Int
readSum reference = do
  r <- Ref.read reference
  pure (r.a + r.b + r.c + r.d + r.e + r.f + r.g + r.h)

readRepeated :: Ref.Ref Eight -> Effect Int
readRepeated reference = do
  r <- Ref.read reference
  pure (r.h + r.a + r.h - r.a + r.b - r.b)

readPartial :: Ref.Ref Eight -> Effect Int
readPartial reference = do
  r <- Ref.read reference
  pure (r.a - r.h)

readTwo :: Ref.Ref Eight -> Ref.Ref Eight -> Effect Int
readTwo first second = do
  l <- Ref.read first
  r <- Ref.read second
  pure (l.a + l.b + l.c + l.d + l.e + l.f + l.g + l.h + r.a + r.b + r.c + r.d + r.e + r.f + r.g + r.h)

readArithmetic :: Int -> Ref.Ref Eight -> Effect Int
readArithmetic delta reference = do
  r <- Ref.read reference
  pure ((r.a + r.b + r.c + r.d + r.e + r.f + r.g + r.h) * 3 + delta)

readUpdated :: Int -> Ref.Ref Eight -> Effect Int
readUpdated delta reference = do
  r <- Ref.read reference
  let updated = r { h = r.h + delta }
  pure (updated.a + updated.b + updated.c + updated.d + updated.e + updated.f + updated.g + updated.h + r.a + r.b + r.c + r.d + r.e + r.f + r.g + r.h)

readBoolean :: Ref.Ref Eight -> Effect Boolean
readBoolean reference = do
  r <- Ref.read reference
  pure (r.a < r.h && r.b == r.b)

mark :: Trace -> Int -> Int -> Int
mark trace tag value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events tag) trace
  pure value

recordFactory :: Trace -> Int -> Eight -> Eight
recordFactory trace tag value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events tag) trace
  pure value

failRead :: Trace -> Int -> Int
failRead trace value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events value) trace
  throw "selected failure"

withUnknown :: Trace -> Ref.Ref Eight -> Effect Int
withUnknown trace reference = do
  r <- Ref.read reference
  pure (r.a + r.b + r.c + mark trace 1 r.d + r.e + r.f + r.g + mark trace 2 r.h + r.a)

withFactories :: Trace -> Ref.Ref Eight -> Effect Int
withFactories trace reference = do
  r <- Ref.read reference
  pure ((recordFactory trace 3 r).a + (recordFactory trace 4 r).b
    + (recordFactory trace 5 r).c + (recordFactory trace 6 r).d
    + (recordFactory trace 7 r).e + (recordFactory trace 8 r).f
    + (recordFactory trace 9 r).g + (recordFactory trace 10 r).h)

withBranch :: Boolean -> Trace -> Ref.Ref Eight -> Effect Int
withBranch enabled trace reference = do
  r <- Ref.read reference
  pure (if enabled then r.a + r.b + r.c + r.d + r.e + r.f + r.g + r.h else failRead trace r.a)

main :: Effect Unit
main = do
  inputs <- Ref.new
    [ { left: 0, right: 19, delta: 1 }
    , { left: 2147483647, right: -2147483648, delta: 1 }
    , { left: -7, right: 9, delta: -3 }
    , { left: -2147483648, right: 2147483647, delta: -1 }
    , { left: 2147483639, right: -11, delta: 31 }
    , { left: -19, right: 3, delta: -2147483648 }
    ]
  cases <- Ref.read inputs
  for_ cases \input -> do
    first <- Ref.new (makeEight input.left)
    second <- Ref.new (makeEight input.right)
    total <- readSum first
    repeated <- readRepeated first
    partial <- readPartial first
    both <- readTwo first second
    arithmetic <- readArithmetic input.delta first
    updated <- readUpdated input.delta first
    boolean <- readBoolean first
    replay <- readSum first
    logShow [ total, repeated, partial, both, arithmetic, updated, replay ]
    logShow boolean
    originalFirst <- Ref.read first
    originalSecond <- Ref.read second
    logShow [ fields originalFirst, fields originalSecond ]
  reference <- Ref.new (makeEight 5)
  trace <- Ref.new []
  withUnknown trace reference >>= logShow
  withFactories trace reference >>= logShow
  Ref.read trace >>= logShow
  branchTrace <- Ref.new []
  withBranch true branchTrace reference >>= logShow
  Ref.read branchTrace >>= logShow
  outcome <- try (withBranch false branchTrace reference)
  logShow (case outcome of
    Left _ -> true
    Right _ -> false)
  Ref.read branchTrace >>= logShow
