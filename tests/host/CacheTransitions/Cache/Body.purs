-- @inline export value never
module Cache.Body (value) where

import Prelude

privateStep :: Int -> Int
privateStep n = n + 2

value :: Int -> Int
value n = privateStep n * 2
