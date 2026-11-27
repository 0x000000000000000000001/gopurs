-- @inline Main.consume never
-- @inline Main.consumeDown never
-- @inline Main.consumeZero never
-- @inline Main.consumeMultiple never
-- @inline Main.withUnknown never
-- @inline Main.withExtraUse never
-- @inline Main.withInterveningEffect never
-- @inline Main.increase never
-- @inline Main.decrease never
-- @inline Main.normalize never
-- @inline Main.multiple never
-- @inline Main.observe never
module Main where

import Prelude
import Data.Foldable (for_)
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref

type Snapshot = { alpha :: Int, beta :: Int, delta :: Int, gamma :: Int, iota :: Int, omega :: Int, theta :: Int, zeta :: Int }

make :: Int -> Snapshot
make seed = { alpha: seed, beta: seed + 1, delta: seed + 2, gamma: seed + 3, iota: seed + 4, omega: seed + 5, theta: seed + 6, zeta: seed + 7 }

total :: Snapshot -> Int
total row = row.alpha + row.beta + row.delta + row.gamma + row.iota + row.omega + row.theta + row.zeta

increase :: Int -> Snapshot -> Snapshot
increase 0 row = row
increase remaining row = increase (remaining - 1) (row { theta = row.theta + 1 })

decrease :: Int -> Snapshot -> Snapshot
decrease 0 row = row
decrease remaining row = decrease (remaining - 1) (row { beta = row.beta - 1 })

normalize :: Int -> Snapshot -> Snapshot
normalize 0 row = row
normalize remaining row = normalize (remaining - 1) (row { gamma = row.gamma + 0 })

multiple :: Int -> Snapshot -> Snapshot
multiple 0 row = row
multiple remaining row = multiple (remaining - 1) (row { delta = row.delta + 3, zeta = row.zeta - 2 })

consume :: Int -> Ref.Ref Snapshot -> Effect Int
consume count reference = do
  original <- Ref.read reference
  let changed = increase count original
  pure (total changed + total original)

consumeDown :: Int -> Ref.Ref Snapshot -> Effect Int
consumeDown count reference = do
  original <- Ref.read reference
  let changed = decrease count original
  pure (total changed + total original)

consumeZero :: Int -> Ref.Ref Snapshot -> Effect Int
consumeZero count reference = do
  original <- Ref.read reference
  let changed = normalize count original
  pure (total changed + total original)

consumeMultiple :: Int -> Ref.Ref Snapshot -> Effect Int
consumeMultiple count reference = do
  original <- Ref.read reference
  let changed = multiple count original
  pure (total changed + total original)

observe :: Snapshot -> Int
observe row = row.theta + 11

withUnknown :: Int -> Ref.Ref Snapshot -> Effect Int
withUnknown count reference = do
  original <- Ref.read reference
  let changed = increase count original
  pure (total changed + observe changed + total original)

withExtraUse :: Int -> Ref.Ref Snapshot -> Effect (Array Int)
withExtraUse count reference = do
  original <- Ref.read reference
  let changed = increase count original
  pure [ total changed + total original, changed.theta ]

withInterveningEffect :: Int -> Ref.Ref Snapshot -> Effect Int
withInterveningEffect count reference = do
  original <- Ref.read reference
  let changed = increase count original
  Ref.write (make 700) reference
  fresh <- Ref.read reference
  pure (total changed + total original + total fresh)

main :: Effect Unit
main = do
  seeds <- Ref.new [ 0, -2147483648, 2147483647, -9 ]
  counts <- Ref.new [ 0, 1, 2, 7, 31 ]
  inputs <- Ref.read seeds
  amounts <- Ref.read counts
  for_ inputs \seed -> for_ amounts \count -> do
    reference <- Ref.new (make seed)
    up <- consume count reference
    down <- consumeDown count reference
    zero <- consumeZero count reference
    both <- consumeMultiple count reference
    unknown <- withUnknown count reference
    extra <- withExtraUse count reference
    original <- Ref.read reference
    logShow [ up, down, zero, both, unknown, total original ]
    logShow extra
  reference <- Ref.new (make 5)
  let saved = consume 7 reference
  saved >>= logShow
  Ref.write (make 19) reference
  saved >>= logShow
  withInterveningEffect 2 reference >>= logShow
  Ref.read reference >>= (logShow <<< total)
