-- @inline export score never
module Worker where
import Prelude
score :: forall r. { id :: Int, name :: String, active :: Boolean | r } -> Int
score row =
  let
    a = row.id * 3 + (if row.active then 7 else 0)
    b = if row.active then a * 7 + 17 else a * 5 - 13
    c = if row.name == "alpha" then b + row.id * 11 else b - row.id * 13
    d = if c < 100 then c * 2 + 7 else c * 2 - 7
    e = if row.active then d + a * 3 else d - b
    f = if row.name == "beta" then e + 19 else e - 23
    g = if row.id < 0 then f - 29 else f + 31
    h = g + a + b + c + d + e + f
  in if h > 200 then h + b * 3 - a * 2 else h - b * 3 + a * 2
