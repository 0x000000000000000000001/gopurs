-- @inline export decode never
module Worker where

import Prelude
import Data.Either (Either(..))

type Summary = { id :: Int, name :: String, active :: Boolean }

-- A shared open-row worker whose native result carries a record payload
-- through Either. Every use below reads declared fields only.
decode :: forall r. { id :: Int, name :: String, active :: Boolean | r } -> Either String Summary
decode row =
  if row.id < 0 then
    Left "negative"
  else
    Right { id: row.id + 1, name: row.name, active: row.active }
