-- @inline export track never
-- @inline export run never
-- @inline export worker never
-- @inline export produced never
-- @inline export input never
-- @inline export generic never
module Main where

-- @dependencies: assert prelude effect console refs either maybe arrays foldable-traversable
import Prelude
import Data.Either (Either(..))
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Test.Assert (assertEqual)

type Item = { value :: Int }

track :: Ref.Ref (Array Int) -> Int -> Int
track calls value = unsafePerformEffect do
  Ref.modify_ (\seen -> seen <> [ value ]) calls
  pure value

run :: Ref.Ref (Array Int) -> Array Int -> Either String (Array Item)
run calls values = traverse (\value ->
  let observed = track calls value
  in if observed < 0 then Left (show observed)
     else Right { value: observed * 2 }) values

-- Keep both the callback and array absent when specializing the dictionaries.
worker :: (Int -> Either String Item) -> Array Int -> Either String (Array Item)
worker = traverse

produced :: Ref.Ref (Array Int) -> Int -> Either String Item
produced calls = unsafePerformEffect do
  Ref.modify_ (\seen -> seen <> [ 200 ]) calls
  pure \value -> Right { value: track calls value }

input :: Ref.Ref (Array Int) -> Array Int
input calls = unsafePerformEffect do
  Ref.modify_ (\seen -> seen <> [ 300 ]) calls
  pure [ 4, 5 ]

-- The unknown applicative and the Maybe instance must retain their semantics.
generic :: forall f. Applicative f => (Int -> f Int) -> Array Int -> f (Array Int)
generic = traverse

main :: Effect Unit
main = do
  calls <- Ref.new []
  assertEqual
    { expected: Right [ { value: 6 }, { value: 14 }, { value: 4 }, { value: 18 } ]
    , actual: run calls [ 3, 7, 2, 9 ]
    }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 3, 7, 2, 9 ], actual }
  Ref.write [] calls
  assertEqual { expected: Left "-1", actual: run calls [ 3, -1, 8, -2, 9 ] }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 3, -1, 8, -2, 9 ], actual }
  Ref.write [] calls
  assertEqual { expected: Right [], actual: run calls [] }
  Ref.read calls >>= \actual -> assertEqual { expected: [], actual }
  held <- Ref.new (traverse (produced calls))
  Ref.read calls >>= \actual -> assertEqual { expected: [ 200 ], actual }
  callback <- Ref.read held
  assertEqual { expected: Right [ { value: 6 } ], actual: callback [ 6 ] }
  assertEqual { expected: Right [ { value: 7 } ], actual: callback [ 7 ] }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 200, 6, 7 ], actual }
  Ref.write [] calls
  assertEqual
    { expected: Right [ { value: 4 }, { value: 5 } ]
    , actual: traverse (produced calls) (input calls)
    }
  Ref.read calls >>= \actual -> assertEqual { expected: [ 200, 300, 4, 5 ], actual }
  savedWorker <- Ref.new worker
  useWorker <- Ref.read savedWorker
  assertEqual { expected: Right [ { value: 10 } ], actual: useWorker (\value -> Right { value }) [ 10 ] }
  assertEqual { expected: Just [ 2, 3 ], actual: generic (\value -> Just (value + 1)) [ 1, 2 ] }
  assertEqual { expected: Nothing, actual: generic (\value -> if value < 0 then Nothing else Just value) [ 1, -1, 2 ] }
  log "Done"
