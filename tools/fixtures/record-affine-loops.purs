module Main where

import Prelude

import Data.Array as Array
import Data.Foldable (for_)
import Data.String.CodeUnits as String
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref

type Snapshot =
  { mark :: Int
  , payload :: { total :: Int, borrowed :: Array Int }
  , tag :: String
  }

advance :: Int -> Snapshot -> Snapshot
advance 0 state = state
advance remaining state = advance (remaining - 1)
  (state { mark = state.mark + (-7)
         , payload = state.payload { total = state.payload.total + (remaining `mod` 7) }
         })

crossFields :: Int -> Snapshot -> Snapshot
crossFields 0 state = state
crossFields remaining state = crossFields (remaining - 1)
  (state { mark = state.mark + state.payload.total
         , payload = state.payload { total = state.payload.total + 1 }
         })

effectSteps :: Ref.Ref Int -> Int -> Snapshot -> Effect Snapshot
effectSteps _ 0 state = pure state
effectSteps trace remaining state = do
  Ref.modify_ (_ + 1) trace
  effectSteps trace (remaining - 1) (advance 1 state)

summary :: Snapshot -> Array Int
summary state = [ state.mark, state.payload.total, Array.length state.payload.borrowed, String.length state.tag ]

main :: Effect Unit
main = do
  seed <- Ref.new { mark: -2147483648, payload: { total: 2147483647, borrowed: [ 4, 7, 11 ] }, tag: "retain" }
  counts <- Ref.new [ 0, 1, 2, 6, 7, 8, 32, 1000 ]
  ns <- Ref.read counts
  for_ ns \count -> do
    source <- Ref.read seed
    logShow (summary (advance count source))
  source <- Ref.read seed
  logShow (summary source)
  let small = source { mark = 2, payload = source.payload { total = 3 } }
  logShow (summary (crossFields 4 small))
  trace <- Ref.new 0
  final <- effectSteps trace 3 small
  logShow (summary final)
  Ref.read trace >>= logShow
  logShow (summary small)
