-- @inline Main.literalInt never
-- @inline Main.literalText never
-- @inline Main.changing never
-- @inline Main.changingWithContinuation never
-- @inline Main.nested never
module Main where

-- @dependencies: prelude effect console refs

import Prelude
import Effect (Effect)
import Effect.Console (log, logShow)
import Effect.Ref as Ref

foreign import identity :: forall a. a -> Effect a
foreign import snapshot :: Effect String
foreign import reset :: Effect Unit

literalInt :: Unit -> Effect Int
literalInt _ = do
  value <- identity 41
  pure (value + 1)

literalText :: Unit -> Effect String
literalText _ = do
  value <- identity "retained"
  pure (value <> "!")

changing :: Ref.Ref Int -> Effect Int
changing ref = do
  input <- Ref.read ref
  value <- identity input
  pure value

changingWithContinuation :: Ref.Ref Int -> Effect Int
changingWithContinuation ref = do
  input <- Ref.read ref
  value <- identity input
  pure (value + 1)

nested :: Int -> Effect Int
nested input = do
  left <- identity input
  right <- identity (left + 1)
  pure (left + right)

main :: Effect Unit
main = do
  reset
  logShow =<< literalInt unit
  logShow =<< literalInt unit
  log =<< literalText unit
  ref <- Ref.new 7
  logShow =<< changing ref
  logShow =<< changingWithContinuation ref
  Ref.write 11 ref
  logShow =<< changing ref
  logShow =<< changingWithContinuation ref
  logShow =<< nested 10
  array <- identity [ 2, 4 ]
  logShow array
  record <- identity { count: 17 }
  logShow record.count
  callback <- identity (\value -> value + 3)
  logShow (callback 4)
  saved <- identity (identity 23)
  logShow =<< saved
  logShow =<< saved
  log =<< snapshot
