module Main where

-- @dependencies: assert prelude effect console avar
-- @snapshot-ffi

import Prelude
import Effect (Effect)
import Effect.AVar (AVar)
import Effect.AVar as AVar
import Effect.Console (log)
import Test.Assert (assertEqual)

foreign import readWhileDraining :: AVar Int -> Effect Boolean

main :: Effect Unit
main = do
  avar <- AVar.new 42
  delivered <- readWhileDraining avar
  assertEqual { expected: true, actual: delivered }
  log "Done"
