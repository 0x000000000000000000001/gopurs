-- @inline Main.create never
module Main where

-- @dependencies: prelude effect console refs foldable-traversable

import Prelude
import Data.Foldable (traverse_)
import Effect (Effect)
import Effect.Console (log, logShow)
import Effect.Ref as Ref

foreign import runInt :: Effect Int -> Effect Int
foreign import runInt64 :: Effect Int -> Effect Int
foreign import runValue :: forall a. Effect a -> Effect a
foreign import tick :: Int -> Effect Int
foreign import reset :: Effect Unit
foreign import snapshot :: Effect String

cached :: Effect Int
cached = do
  value <- tick 40
  pure (value + 2)

cachedAlias :: Effect Int
cachedAlias = cached

pureAction :: Effect Int
pureAction = pure 7

create :: Int -> Effect Int
create input = do
  value <- tick input
  pure (value + 1)

main :: Effect Unit
main = do
  reset
  logShow =<< runInt cached
  logShow =<< runInt cachedAlias
  logShow =<< runInt64 cached
  logShow =<< runInt pureAction
  logShow =<< runInt (create 8)
  let actions = map create [ 2, 4, 6 ]
  traverse_ (\action -> logShow =<< runInt action) actions
  traverse_ (\action -> logShow =<< runInt action) actions
  ref <- Ref.new 10
  let dynamic = do
        input <- Ref.read ref
        value <- tick input
        pure (value + 1)
  logShow =<< runInt dynamic
  Ref.write 20 ref
  logShow =<< runInt dynamic
  logShow =<< runValue cached
  log =<< runValue (pure "hello")
  logShow =<< cached
  logShow =<< cached
  log =<< snapshot
