module Main where

import Prelude

import Control.Monad.Rec.Class (Step(..), tailRecM)
import Control.Monad.ST (ST)
import Control.Monad.ST as ST
import Control.Monad.ST.Ref as STRef
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref

writeEffect :: Int -> Int -> Ref.Ref Int -> Effect Unit
writeEffect start limit ref = tailRecM step start
  where
  step index
    | index > limit = pure (Done unit)
    | otherwise = do
        Ref.write index ref
        pure (Loop (index + 1))

modifyEffect :: Int -> Int -> Ref.Ref Int -> Effect Unit
modifyEffect count delta ref = tailRecM step count
  where
  step remaining
    | remaining <= 0 = pure (Done unit)
    | otherwise = do
        _ <- Ref.modify (_ + delta) ref
        pure (Loop (remaining - 1))

writeST :: forall r. Int -> Int -> STRef.STRef r Int -> ST r Unit
writeST start limit ref = tailRecM step start
  where
  step index
    | index > limit = pure (Done unit)
    | otherwise = do
        _ <- STRef.write index ref
        pure (Loop (index + 1))

modifyST :: forall r. Int -> Int -> STRef.STRef r Int -> ST r Unit
modifyST count delta ref = tailRecM step count
  where
  step remaining
    | remaining <= 0 = pure (Done unit)
    | otherwise = do
        _ <- STRef.modify (_ + delta) ref
        pure (Loop (remaining - 1))

bumpEffect :: Int -> Ref.Ref Int -> Effect Unit
bumpEffect count ref = tailRecM step count
  where
  step remaining
    | remaining <= 0 = pure (Done unit)
    | otherwise = do
        _ <- Ref.modify (_ + 1) ref
        pure (Loop (remaining - 1))

bumpST :: forall r. Int -> STRef.STRef r Int -> ST r Unit
bumpST count ref = tailRecM step count
  where
  step remaining
    | remaining <= 0 = pure (Done unit)
    | otherwise = do
        _ <- STRef.modify (_ + 1) ref
        pure (Loop (remaining - 1))

-- Helper and caller levels deliberately overlap with different meanings.
-- Substitution must be simultaneous, so a count cannot become a reference.
effectPermuted :: Int -> Int -> Effect Int
effectPermuted seed count = do
  ref <- Ref.new seed
  bumpEffect count ref
  Ref.read ref

stPermuted :: Int -> Int -> Int
stPermuted seed count = ST.run do
  ref <- STRef.new seed
  bumpST count ref
  STRef.read ref

effectWritePermuted :: Int -> Int -> Int -> Effect Int
effectWritePermuted limit seed start = do
  ref <- Ref.new seed
  writeEffect start limit ref
  Ref.read ref

stWritePermuted :: Int -> Int -> Int -> Int
stWritePermuted limit seed start = ST.run do
  ref <- STRef.new seed
  writeST start limit ref
  STRef.read ref

effectWrite :: Int -> Int -> Int -> Effect Int
effectWrite start limit seed = do
  ref <- Ref.new seed
  writeEffect start limit ref
  Ref.read ref

effectModify :: Int -> Int -> Int -> Effect Int
effectModify count seed delta = do
  ref <- Ref.new seed
  modifyEffect count delta ref
  Ref.read ref

stWrite :: Int -> Int -> Int -> Int
stWrite start limit seed = ST.run do
  ref <- STRef.new seed
  writeST start limit ref
  STRef.read ref

stModify :: Int -> Int -> Int -> Int
stModify count seed delta = ST.run do
  ref <- STRef.new seed
  modifyST count delta ref
  STRef.read ref

-- An escaping reference and a read through a retained closure must continue
-- to refer to the actual cell after its mutating helper returns.
escaping :: Int -> Effect (Ref.Ref Int)
escaping count = do
  ref <- Ref.new 10
  modifyEffect count 3 ref
  pure ref

retained :: Int -> Effect (Effect Int)
retained count = do
  ref <- Ref.new 10
  modifyEffect count 2 ref
  pure do
    Ref.modify_ (_ + 1) ref
    Ref.read ref

-- Every iteration has an observer; replacing the whole body by the final
-- write would lose the history. target and history may even be aliases.
observed :: Int -> Ref.Ref Int -> Ref.Ref Int -> Effect Unit
observed count target history = tailRecM step count
  where
  step remaining
    | remaining <= 0 = pure (Done unit)
    | otherwise = do
        value <- Ref.read target
        Ref.modify_ (\total -> total * 10 + value) history
        Ref.modify_ (_ + 1) target
        pure (Loop (remaining - 1))

callbackLoop :: Int -> Effect Unit -> Ref.Ref Int -> Effect Unit
callbackLoop count callback target = tailRecM step count
  where
  step remaining
    | remaining <= 0 = pure (Done unit)
    | otherwise = do
        callback
        Ref.modify_ (_ + 1) target
        pure (Loop (remaining - 1))

main :: Effect Unit
main = do
  -- Reading the cases through a reference also exercises runtime parameters.
  inputs <- Ref.new
    [ { count: -7, seed: 15, delta: 4, start: 5, limit: 3 }
    , { count: 0, seed: -12, delta: 7, start: 0, limit: -1 }
    , { count: 1, seed: 2147483647, delta: 1, start: -2, limit: 1 }
    , { count: 2, seed: -2147483648, delta: -1, start: -3, limit: -1 }
    , { count: 5, seed: 23, delta: -7, start: 2, limit: 2 }
    , { count: 33, seed: 2147483640, delta: 11, start: 2147483644, limit: 2147483646 }
    , { count: 10000, seed: 19, delta: 3, start: 1, limit: 10000 }
    ]
  cases <- Ref.read inputs
  for_ cases \c -> do
    changed <- effectModify c.count c.seed c.delta
    written <- effectWrite c.start c.limit c.seed
    bumped <- effectPermuted c.seed c.count
    permutedWrite <- effectWritePermuted c.limit c.seed c.start
    logShow [ changed, written, stModify c.count c.seed c.delta, stWrite c.start c.limit c.seed
            , bumped, stPermuted c.seed c.count, permutedWrite, stWritePermuted c.limit c.seed c.start ]

  let action = effectModify 4 7 3
  first <- action
  second <- action
  logShow [ first, second ]

  mutableSeed <- Ref.new 8
  let deferred = do
        seed <- Ref.read mutableSeed
        effectModify 3 seed 2
  Ref.write 20 mutableSeed
  before <- deferred
  Ref.write (-5) mutableSeed
  after <- deferred
  logShow [ before, after ]

  external <- escaping 4
  Ref.modify_ (_ + 1) external
  Ref.read external >>= logShow
  captured <- retained 3
  a <- captured
  b <- captured
  logShow [ a, b ]

  shared <- Ref.new 2
  let alias = shared
  modifyEffect 3 2 shared
  original <- Ref.read alias
  writeEffect (-1) 3 alias
  final <- Ref.read shared
  logShow [ original, final ]

  target <- Ref.new 2
  history <- Ref.new 0
  observed 3 target history
  targetValue <- Ref.read target
  historyValue <- Ref.read history
  logShow [ targetValue, historyValue ]
  same <- Ref.new 1
  observed 2 same same
  Ref.read same >>= logShow

  ticks <- Ref.new 0
  cell <- Ref.new 1
  callbackLoop 3 (Ref.modify_ (_ + 1) ticks) cell
  tickValue <- Ref.read ticks
  cellValue <- Ref.read cell
  logShow [ tickValue, cellValue ]
