module Main where

-- @dependencies: console

import Effect.Console (log)

class C (t :: Type)
instance C (f a)

f :: C (Array String) => Int
f = 0

v :: Int
v = f

main = log "Done"
