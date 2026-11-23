-- @inline export makeStack never
-- @inline export sumStack never
-- @inline export run never
module Main where

-- @dependencies: assert lists tuples prelude effect console refs

import Prelude

import Data.List (List(..))
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

data Direction = Leftward | Rightward | Stationary

type Cell = { offset :: Int, row :: Int }
type Stack = List (Tuple Cell Direction)

-- This path has only direct, monomorphic calls. It is a control for the cost
-- of choosing a shared representation where no generic boundary needs it.
makeStack :: Int -> Int -> Stack
makeStack size seed = go size Nil
  where
  go n rest
    | n <= 0 = rest
    | otherwise =
        let direction = if n `mod` 3 == 0 then Stationary else if n `mod` 3 == 1 then Leftward else Rightward
        in go (n - 1) (Cons (Tuple { offset: seed + n, row: n } direction) rest)

sumStack :: Stack -> Int
sumStack stack = go 0 stack
  where
  go total = case _ of
    Nil -> total
    Cons (Tuple cell direction) rest ->
      let
        delta = case direction of
          Leftward -> -1
          Rightward -> 1
          Stationary -> 0
      in go (total + cell.offset + cell.row + delta) rest

run :: Int -> Int -> Int
run count size = go count 0
  where
  go n total
    | n <= 0 = total
    | otherwise = go (n - 1) (total + sumStack (makeStack size (n `mod` 11)))

main :: Effect Unit
main = do
  seedRef <- Ref.new 3
  seed <- Ref.read seedRef
  assertEqual { expected: 21, actual: sumStack (makeStack 3 seed) }
  assertEqual { expected: 0, actual: sumStack (makeStack 0 seed) }
  assertEqual { expected: 30, actual: run seed 2 }
  log "Done"
