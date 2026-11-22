-- @inline export throughCallback never
-- @inline export nativeResult never
module Main where

-- @dependencies: assert arrays lists tuples prelude effect console refs

import Prelude

import Data.Array as Array
import Data.List (List(..))
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

type Position = { column :: Int, line :: Int }
type Issue = { error :: String, position :: Position }
type State = { consumed :: Boolean, errors :: Array Issue, stream :: List Int }

throughCallback :: (State -> State) -> State -> State
throughCallback transform original =
  let
    intermediate = transform original
    finish _ = transform intermediate
  in
    finish unit

nativeResult :: State -> Tuple (List Int) (Array Issue)
nativeResult state = Tuple state.stream state.errors

setConsumed :: forall row. { consumed :: Boolean | row } -> { consumed :: Boolean | row }
setConsumed row = row { consumed = true }

main :: Effect Unit
main = do
  transformRef <- Ref.new \state -> state
    { consumed = true
    , errors = Array.snoc state.errors { error: "λ😀", position: { column: -2147483648, line: 2147483647 } }
    }
  transform <- Ref.read transformRef
  originalRef <- Ref.new { consumed: false, errors: [], stream: Cons 7 (Cons 9 Nil) }
  original <- Ref.read originalRef
  let result = throughCallback transform original
  assertEqual { expected: 2, actual: Array.length result.errors }
  assertEqual { expected: true, actual: result.consumed }
  preserved <- Ref.read originalRef
  assertEqual { expected: original, actual: preserved }
  assertEqual { expected: false, actual: original.consumed }
  assertEqual { expected: [], actual: original.errors }
  assertEqual { expected: Tuple result.stream result.errors, actual: nativeResult result }
  assertEqual
    { expected: Tuple Nil []
    , actual: nativeResult { consumed: false, errors: [], stream: Nil }
    }
  let forward flag = original { consumed = flag }
  callbackRef <- Ref.new forward
  callback <- Ref.read callbackRef
  assertEqual { expected: true, actual: (callback true).consumed }
  assertEqual { expected: false, actual: (callback false).consumed }
  assertEqual { expected: original.stream, actual: (callback true).stream }
  assertEqual { expected: { consumed: true, extra: "preserved" }, actual: setConsumed { consumed: false, extra: "preserved" } }
  log "Done"
