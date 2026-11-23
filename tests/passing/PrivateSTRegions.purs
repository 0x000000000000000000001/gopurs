module Main where

-- @dependencies: arrays assert console effect prelude refs st

import Prelude

import Control.Monad.ST as ST
import Control.Monad.ST.Ref as STRef
import Data.Array.ST as STArray
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

data Step a = Emit a Int | Stop Int | Failed Int

type Result a = { values :: Array a, next :: Int, failed :: Boolean }

-- The callback never receives the private references or mutable array. Both
-- terminal branches stop the loop; only the successful one publishes the array.
collect :: forall a. (Int -> Step a) -> Int -> Result a
collect step start = ST.run do
  values <- STArray.new
  cursor <- STRef.new start
  continue <- STRef.new true
  result <- STRef.new { values: [], next: start, failed: false }
  ST.while (STRef.read continue) do
    current <- STRef.read cursor
    case step current of
      Emit value next -> do
        _ <- STArray.push value values
        void (STRef.write next cursor)
      Stop next -> do
        frozen <- STArray.unsafeFreeze values
        _ <- STRef.write { values: frozen, next, failed: false } result
        void (STRef.write false continue)
      Failed next -> do
        _ <- STRef.write { values: [], next, failed: true } result
        void (STRef.write false continue)
  STRef.read result

-- write returns the assigned value and push returns the new array length.
operationResults :: Int -> Array Int
operationResults initial = ST.run do
  cell <- STRef.new initial
  values <- STArray.new
  assigned <- STRef.write (initial + 1) cell
  count <- STArray.push assigned values
  current <- STRef.read cell
  _ <- STArray.push (current + count) values
  STArray.unsafeFreeze values

skipLoop :: Int -> Int
skipLoop initial = ST.run do
  cell <- STRef.new initial
  ST.while (pure false) do
    current <- STRef.read cell
    void (STRef.write (current + 1) cell)
  STRef.read cell

-- Safe freeze takes a snapshot even when the mutable source is changed later.
freezeSnapshot :: Int -> { before :: Array Int, after :: Array Int }
freezeSnapshot initial = ST.run do
  values <- STArray.new
  _ <- STArray.push initial values
  before <- STArray.freeze values
  _ <- STArray.push (initial + 1) values
  after <- STArray.unsafeFreeze values
  pure { before, after }

-- The optimizer can remove these aliases, but runtime behavior must agree
-- whether the private-region optimization admits the remaining shape or not.
aliasRead :: Int -> Int
aliasRead initial = ST.run do
  cell <- STRef.new initial
  let alias = { cell }
  _ <- STRef.write (initial + 1) alias.cell
  let readLater _ = STRef.read alias.cell
  readLater unit

snapshotRead :: Int -> Int -> Int
snapshotRead initial = ST.run do
  cell <- STRef.new initial
  value <- STRef.read cell
  _ <- STRef.write (initial + 1) cell
  pure (\offset -> value + offset)

main :: Effect Unit
main = do
  stepRef <- Ref.new (\n -> if n < 4 then Emit (n * 3) (n + 1) else Stop (n + 7))
  step <- Ref.read stepRef
  keptRef <- Ref.new (collect step 1)
  first <- Ref.read keptRef
  assertEqual
    { expected: { values: [ 3, 6, 9 ], next: 11, failed: false }
    , actual: first
    }
  assertEqual
    { expected: { values: [ 9 ], next: 11, failed: false }
    , actual: collect step 3
    }
  assertEqual
    { expected: { values: [], next: 12, failed: false }
    , actual: collect step 5
    }
  preserved <- Ref.read keptRef
  assertEqual
    { expected: { values: [ 3, 6, 9 ], next: 11, failed: false }
    , actual: preserved
    }
  assertEqual
    { expected: { values: [], next: 91, failed: true }
    , actual: collect (\n -> if n < 2 then Emit n (n + 1) else Failed 91) 0
    }
  assertEqual
    { expected: { values: [ "λ😀", "λ😀" ], next: 2, failed: false }
    , actual: collect (\n -> if n < 2 then Emit "λ😀" (n + 1) else Stop n) 0
    }
  assertEqual
    { expected: { values: [ [ 0, 1 ], [ 1, 2 ] ], next: 2, failed: false }
    , actual: collect (\n -> if n < 2 then Emit [ n, n + 1 ] (n + 1) else Stop n) 0
    }
  input <- Ref.new 10 >>= Ref.read
  assertEqual { expected: [ 11, 12 ], actual: operationResults input }
  assertEqual { expected: 10, actual: skipLoop input }
  assertEqual
    { expected: { before: [ 10 ], after: [ 10, 11 ] }
    , actual: freezeSnapshot input
    }
  assertEqual { expected: 11, actual: aliasRead input }
  assertEqual { expected: 15, actual: snapshotRead input 5 }
  log "Done"
