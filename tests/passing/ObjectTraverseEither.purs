-- @inline Data.TraversableWithIndex.traverseWithIndex never
-- @inline export track never
-- @inline export run never
-- @inline export partial never
module Main where

-- @dependencies: assert prelude effect console refs either foreign-object foldable-traversable tuples
import Prelude
import Data.Either (Either(..))
import Data.Maybe (Maybe(..))
import Data.TraversableWithIndex (traverseWithIndex)
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Foreign.Object as Object
import Test.Assert (assertEqual)

track :: Ref.Ref (Array String) -> String -> Int -> Int
track calls key value = unsafePerformEffect do
  Ref.modify_ (\seen -> seen <> [ key ]) calls
  pure value

run :: Ref.Ref (Array String) -> Object.Object Int -> Either String (Object.Object Int)
run calls = traverseWithIndex \key value ->
  let observed = track calls key value
  in if observed < 0 then Left key else Right (observed * 2)

partial :: Ref.Ref (Array String) -> Int -> Object.Object Int -> Either String (Object.Object Int)
partial calls offset =
  let captured = track calls "capture" offset
  in traverseWithIndex (\_ value -> Right (value + captured))

main :: Effect Unit
main = do
  calls <- Ref.new []
  let input = Object.fromFoldable [ Tuple "a" 3, Tuple "b" 7 ]
  assertEqual
    { expected: Right (Object.fromFoldable [ Tuple "a" 6, Tuple "b" 14 ])
    , actual: run calls input
    }
  Ref.read calls >>= \actual -> assertEqual { expected: [ "a", "b" ], actual }
  Ref.write [] calls
  assertEqual
    { expected: Left "a"
    , actual: run calls (Object.fromFoldable [ Tuple "a" (-1), Tuple "b" 8, Tuple "c" (-2) ])
    }
  Ref.read calls >>= \actual -> assertEqual { expected: [ "a", "b", "c" ], actual }
  Ref.write [] calls
  assertEqual { expected: Right Object.empty, actual: run calls Object.empty }
  Ref.read calls >>= \actual -> assertEqual { expected: [], actual }
  held <- Ref.new (partial calls 10)
  Ref.read calls >>= \actual -> assertEqual { expected: [ "capture" ], actual }
  callback <- Ref.read held
  assertEqual { expected: Right (Object.singleton "a" 13), actual: callback (Object.singleton "a" 3) }
  assertEqual { expected: Right (Object.singleton "a" 17), actual: callback (Object.singleton "a" 7) }
  Ref.read calls >>= \actual -> assertEqual { expected: [ "capture" ], actual }
  assertEqual { expected: Just 3, actual: Object.lookup "a" input }
  log "Done"
