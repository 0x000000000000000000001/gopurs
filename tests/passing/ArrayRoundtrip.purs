-- Keep the conversion paths visible in the generated Go snapshot.
-- @inline export sumEvens never
-- @inline export sumArrayEvens never
-- @inline export sumRangeEvens never
module Main where

-- @dependencies: assert arrays prelude effect console refs

import Prelude

import Data.Array as Array
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

range :: Int -> Int -> Array Int
range start end = Array.range start end

filterEvens :: Array Int -> Array Int
filterEvens arr = Array.filter (\x -> mod x 2 == 0) arr

-- Same range -> filter -> fold expression as altbak's ArrayOps benchmark.
-- Before optimization, the filtered []Value is copied to []int64, then
-- immediately copied back to []Value for the fold.
sumEvens :: Int -> Int
sumEvens n = Array.foldl (+) 0 (filterEvens (range 1 n))

-- Array.range is inclusive and descends when start > end. Use an explicit
-- array input to cover genuinely empty inputs as well as sparse values.
sumArrayEvens :: Array Int -> Int
sumArrayEvens values = Array.foldl (+) 0 (filterEvens values)

sumRangeEvens :: Int -> Int -> Int
sumRangeEvens start end = Array.foldl (+) 0 (filterEvens (range start end))

check :: String -> Int -> Int -> Effect Unit
check label expected actual = do
  assertEqual { expected, actual }
  log (label <> ": " <> show actual)

checkArray :: String -> Int -> Array Int -> Effect Unit
checkArray label expected values = do
  inputRef <- Ref.new values
  input <- Ref.read inputRef
  check label expected (sumArrayEvens input)

checkRange :: String -> Int -> Int -> Int -> Effect Unit
checkRange label expected start end = do
  inputRef <- Ref.new { start, end }
  input <- Ref.read inputRef
  check label expected (sumRangeEvens input.start input.end)

checkBenchmark :: String -> Int -> Int -> Effect Unit
checkBenchmark label expected n = do
  inputRef <- Ref.new n
  input <- Ref.read inputRef
  check label expected (sumEvens input)

main :: Effect Unit
main = do
  checkArray "empty" 0 []
  checkArray "singleton odd" 0 [7]
  checkArray "singleton even" 8 [8]
  checkArray "singleton zero" 0 [0]
  checkArray "even only" 12 [2, 4, 6]
  checkArray "odd only" 0 [1, 3, 5]
  checkArray "mixed parity" 12 [1, 2, 3, 4, 5, 6]
  checkArray "negative values" (-6) [-5, -4, -3, -2, -1]
  checkArray "mixed signs" (-2) [-5, -4, -1, 0, 2, 7]
  checkArray "duplicates" (-4) [2, -4, 2, -4]

  -- Keep every partial sum within PureScript's Int32 bounds: this fixture
  -- tests representation conversions, not the backends' overflow behavior.
  checkArray "minimum int" (-2147483648) [-2147483648]
  checkArray "maximum int is odd" 0 [2147483647]
  checkArray "maximum even int" 2147483646 [2147483646]
  checkArray "bounds with cancellation" 0 [2147483646, -2147483648, 2]

  checkRange "ascending range" 12 1 6
  checkRange "descending range" 12 6 1
  checkRange "negative ascending range" (-6) (-5) (-1)
  checkRange "negative descending range" (-6) (-1) (-5)
  checkRange "range crossing zero" 0 (-3) 3
  -- Each boundary range has only two elements and one even value.
  checkRange "minimum boundary range" (-2147483648) (-2147483648) (-2147483647)
  checkRange "minimum boundary reversed" (-2147483648) (-2147483647) (-2147483648)
  checkRange "maximum boundary range" 2147483646 2147483646 2147483647
  checkRange "maximum boundary reversed" 2147483646 2147483647 2147483646

  checkBenchmark "benchmark n=0 descends" 0 0
  checkBenchmark "benchmark n=1 filters to empty" 0 1
  checkBenchmark "benchmark n=2" 2 2
  checkBenchmark "benchmark negative n" (-6) (-4)
  checkBenchmark "benchmark n=900" 202950 900

  log "Done"
