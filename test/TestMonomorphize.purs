module TestMonomorphize where

import Prelude
import Effect.Console (log)

myIdentity :: Int -> Int -> Int
myIdentity x y = y

main = do
  let a = myIdentity 42 42
  log (show a)
