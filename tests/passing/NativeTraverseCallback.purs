-- @inline export track never
-- @inline export run never
-- @inline export partial never
-- @inline export staged never
module Main where

-- @dependencies: assert prelude effect console refs either arrays foldable-traversable
import Prelude
import Data.Either (Either(..))
import Data.TraversableWithIndex (traverseWithIndex)
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Test.Assert (assertEqual)

type Item = { index :: Int, value :: Int }

-- Observe evaluation order, including callbacks after the first error.
track :: Ref.Ref (Array Int) -> Int -> Int -> Int
track calls marker value = unsafePerformEffect do
  Ref.modify_ (\seen -> seen <> [ marker ]) calls
  pure value

run :: Ref.Ref (Array Int) -> Array Int -> Either String (Array Item)
run calls values = traverseWithIndex (\i value ->
  let observed = track calls i value
  in if observed < 0 then Left (show i)
     else Right { index: i, value: observed }) values

-- The capture is evaluated once when constructing this reusable traversal.
partial :: Ref.Ref (Array Int) -> Int -> Array Int -> Either String (Array Item)
partial calls offset =
  let captured = track calls 100 offset
  in traverseWithIndex (\i value -> Right { index: i, value: value + captured })

-- A computation between lambdas must keep the ordinary curried callback path.
staged :: Ref.Ref (Array Int) -> Array Int -> Either String (Array Item)
staged calls values = traverseWithIndex (\i ->
  let captured = track calls (100 + i) i
  in \value -> if value < 0 then Left (show captured)
     else Right { index: captured, value }) values

main :: Effect Unit
main = do
  calls <- Ref.new []
  assertEqual
    { expected: Right [ { index: 0, value: 3 }, { index: 1, value: 7 } ]
    , actual: run calls [ 3, 7 ]
    }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 0, 1 ], actual }
  Ref.write [] calls
  assertEqual { expected: Left "0", actual: run calls [ -1, 8, -2 ] }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 0, 1, 2 ], actual }
  Ref.write [] calls
  assertEqual { expected: Right [], actual: run calls [] }
  Ref.read calls >>= \actual -> assertEqual { expected: [], actual }
  held <- Ref.new (partial calls 10)
  Ref.read calls >>= \actual -> assertEqual { expected: [ 100 ], actual }
  callback <- Ref.read held
  assertEqual { expected: Right [ { index: 0, value: 13 } ], actual: callback [ 3 ] }
  assertEqual { expected: Right [ { index: 0, value: 17 } ], actual: callback [ 7 ] }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 100 ], actual }
  Ref.write [] calls
  assertEqual { expected: Left "1", actual: staged calls [ 3, -1, -2 ] }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 100, 101, 102 ], actual }
  log "Done"
