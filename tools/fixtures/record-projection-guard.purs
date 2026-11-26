-- @inline Main.mark never
-- @inline Main.markAfter never
-- @inline Main.sourceFactory never
module Main where

-- @dependencies: prelude effect console refs arrays foldable-traversable

import Prelude
import Data.Array as Array
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

type Pair = { left :: Int, right :: Int }
type Eight = { a :: Int, b :: Int, c :: Int, d :: Int, e :: Int, f :: Int, g :: Int, h :: Int }
type Wide =
  { a :: Int, b :: Int, c :: Int, d :: Int, e :: Int, f :: Int, g :: Int, h :: Int
  , i :: Int, j :: Int, k :: Int, l :: Int, m :: Int, n :: Int, o :: Int, p :: Int }

pairFields :: Pair -> Array Int
pairFields r = [ r.left, r.right ]

eightFields :: Eight -> Array Int
eightFields r = [ r.a, r.b, r.c, r.d, r.e, r.f, r.g, r.h ]

wideFields :: Wide -> Array Int
wideFields r = [ r.a, r.b, r.c, r.d, r.e, r.f, r.g, r.h, r.i, r.j, r.k, r.l, r.m, r.n, r.o, r.p ]

updatePair :: Int -> Pair -> Pair
updatePair 0 r = r
updatePair n r = updatePair (n - 1) (r { right = r.right + 1 })

updateEight :: Int -> Eight -> Eight
updateEight 0 r = r
updateEight n r = updateEight (n - 1) (r { h = r.h + 1 })

updateWide :: Int -> Wide -> Wide
updateWide 0 r = r
updateWide n r = updateWide (n - 1) (r { p = r.p + 1 })

makeWide :: Int -> Wide
makeWide base =
  { a: base, b: base + 1, c: base + 2, d: base + 3
  , e: base + 4, f: base + 5, g: base + 6, h: base + 7
  , i: base + 8, j: base + 9, k: base + 10, l: base + 11
  , m: base + 12, n: base + 13, o: base + 14, p: base + 15 }

-- The reference boundary supplies a boxed record at runtime. Reading all
-- fields of both the old and new record makes destructive reuse observable.
inspectWide :: Int -> Ref.Ref Wide -> Effect (Array (Array Int))
inspectWide count reference = do
  source <- Ref.read reference
  let partial = updateWide count
  pure [ wideFields source, wideFields (partial source), wideFields (partial source), wideFields source ]

inspectEight :: Int -> Ref.Ref Eight -> Effect (Array (Array Int))
inspectEight count reference = do
  source <- Ref.read reference
  pure [ eightFields source, eightFields (updateEight count source), eightFields source ]

inspectPair :: Int -> Ref.Ref Pair -> Effect (Array (Array Int))
inspectPair count reference = do
  source <- Ref.read reference
  pure [ pairFields source, pairFields (updatePair count source), pairFields source ]

type Mixed = { a :: Int, borrowed :: Array Int, nested :: { flag :: Boolean, x :: Int }, shared :: Ref.Ref Int, tag :: String }

typeChanged :: Pair -> { left :: String, right :: Int }
typeChanged source = source { left = "changed" }

type Trace = Ref.Ref (Array Int)

mark :: Trace -> Int -> Int -> Int
mark trace tag value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events tag) trace
  pure value

markAfter :: Trace -> Int -> Int -> Int
markAfter trace previous value = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events (200 + previous)) trace
  pure value

sourceFactory :: Trace -> Wide -> Wide
sourceFactory trace source = unsafePerformEffect do
  Ref.modify_ (\events -> Array.snoc events 0) trace
  pure source

main :: Effect Unit
main = do
  inputs <- Ref.new
    [ { count: 0, seed: 19 }, { count: 1, seed: -7 }
    , { count: 2, seed: 2147483647 }, { count: 7, seed: -2147483648 }
    , { count: 31, seed: 2147483630 }, { count: 1000, seed: -19 }
    ]
  cases <- Ref.read inputs
  for_ cases \input -> do
    let wide = makeWide input.seed
    wideRef <- Ref.new wide
    eightRef <- Ref.new { a: wide.a, b: wide.b, c: wide.c, d: wide.d, e: wide.e, f: wide.f, g: wide.g, h: wide.h }
    pairRef <- Ref.new { left: wide.a, right: wide.b }
    inspectWide input.count wideRef >>= logShow
    inspectEight input.count eightRef >>= logShow
    inspectPair input.count pairRef >>= logShow
    original <- Ref.read pairRef
    let changed = typeChanged original
    logShow changed.left
    logShow [ original.left, original.right, changed.right ]
  cell <- Ref.new 10
  mixedRef <- Ref.new { a: 17, borrowed: [ 4, 7, 11 ], nested: { flag: true, x: 29 }, shared: cell, tag: "keep λ" }
  source <- Ref.read mixedRef
  let result = source { a = source.a + 3, nested = source.nested { x = source.nested.x + 1 } }
  Ref.write 77 result.shared
  sourceCell <- Ref.read source.shared
  resultCell <- Ref.read result.shared
  logShow [ source.a, result.a, source.nested.x, result.nested.x, sourceCell, resultCell ]
  logShow [ source.borrowed, result.borrowed ]
  logShow [ source.nested.flag, result.nested.flag ]
  logShow source.tag
  logShow result.tag
  trace <- Ref.new []
  traceInput <- Ref.new (makeWide 5)
  observed <- Ref.read traceInput
  let
    original = sourceFactory trace observed
    first = mark trace 1 (original.p + 1)
    second = markAfter trace first (original.a - 1)
    updated = original { p = first, a = second }
  logShow [ wideFields original, wideFields updated, wideFields original ]
  events <- Ref.read trace
  logShow events
