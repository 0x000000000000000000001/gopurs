-- @inline Main.snapshot never
-- @inline Main.effectSnapshot never
-- @inline Main.rounds never
-- @inline Main.effectRounds never
-- @inline Main.localRounds never
-- @inline Main.moduloRounds never
-- @inline Main.localModuloRounds never
module Main where

-- @dependencies: prelude effect console refs

import Prelude
import Effect (Effect)
import Effect.Console (log, logShow)
import Effect.Ref as Ref

-- Closed record evaluation retains a strict global binding demand even when
-- the caller projects only the integer field.
snapshot :: Int -> { value :: Int, witness :: Unit }
snapshot remaining
  | remaining == 0 = { value: 0, witness: unit }
  | otherwise =
      let previous = snapshot (remaining - 1)
      in { value: previous.value + 1, witness: previous.witness }

effectSnapshot :: Int -> { value :: Int, witness :: String -> Effect Unit }
effectSnapshot remaining
  | remaining == 0 = { value: 0, witness: log }
  | otherwise =
      let previous = effectSnapshot (remaining - 1)
      in { value: previous.value + 1, witness: previous.witness }

rounds :: Int -> Int -> Int
rounds remaining total
  | remaining == 0 = total
  | otherwise = rounds (remaining - 1) (total + (snapshot 7).value)

effectRounds :: Int -> Int -> Int
effectRounds remaining total
  | remaining == 0 = total
  | otherwise = effectRounds (remaining - 1) (total + (effectSnapshot 5).value)

localRounds :: Int -> Int -> Int -> Int
localRounds offset count seed =
  let
    repeat remaining total
      | remaining == 0 = total
      | otherwise = repeat (remaining - 1) (total + (snapshot 11).value)
  in offset + repeat count seed

moduloRounds :: Int -> Int -> Int
moduloRounds remaining total
  | remaining == 0 = total
  | otherwise = moduloRounds (remaining - 1) (total + (remaining `mod` 7))

localModuloRounds :: Int -> Int -> Int -> Int
localModuloRounds offset count seed =
  let
    repeat remaining total
      | remaining == 0 = total
      | otherwise = repeat (remaining - 1) (total + (remaining `mod` 3))
  in offset + repeat count seed

main :: Effect Unit
main = do
  repetitions <- Ref.new 20
  count <- Ref.read repetitions
  maximum <- Ref.new 2147483647
  maxInt <- Ref.read maximum
  logShow [ rounds 0 maxInt, rounds 1 0, rounds count 0, rounds 5 (-40), rounds 2 maxInt ]
  logShow [ localRounds 3 count 0, localRounds 3 0 (-17), localRounds 0 2 maxInt ]
  logShow [ effectRounds 0 maxInt, effectRounds count 2, effectRounds 2 maxInt ]
  logShow [ moduloRounds 0 maxInt, moduloRounds count 0, moduloRounds 1 (-17), moduloRounds count maxInt ]
  logShow [ localModuloRounds 3 count 0, localModuloRounds 0 0 (-17) ]
