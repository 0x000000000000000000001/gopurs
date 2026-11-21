module Main where

-- @dependencies: console

import Effect.Console (log)

data Proxy a = Proxy

type Prozy = Proxy

test1 = Proxy :: Prozy Int
test2 = Proxy :: Prozy "foo"

main = log "Done"
