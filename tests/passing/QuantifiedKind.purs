module Main where

-- @dependencies: console

import Effect.Console (log)

data Proxy a = Proxy

test :: forall k (a :: k). Proxy a
test = Proxy

main = log "Done"
