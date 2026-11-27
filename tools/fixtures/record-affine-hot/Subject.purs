module Subject where

import Prelude

type Snapshot = { total :: Int, borrowed :: Array Int, stamp :: String }

step :: Int -> Snapshot -> Snapshot
step 0 snapshot = snapshot
step remaining snapshot = step (remaining - 1) (snapshot { total = snapshot.total + 1 })

normalize :: Int -> Snapshot -> Snapshot
normalize 0 snapshot = snapshot
normalize remaining snapshot = normalize (remaining - 1) (snapshot { total = snapshot.total + 0 })

down :: Int -> Snapshot -> Snapshot
down 0 snapshot = snapshot
down remaining snapshot = down (remaining - 1) (snapshot { total = snapshot.total + (-1) })
