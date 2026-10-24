-- @inline Main.lookup never
module Main where

-- @dependencies: prelude effect console lists partial

import Prelude
import Data.List (List(..))
import Effect (Effect)
import Effect.Console (logShow)
import Partial.Unsafe (unsafeCrashWith)

-- The local worker consumes two lists and returns their selected element.
-- When a is a function, its remaining arguments belong to the caller.
lookup :: forall a. String -> List String -> List a -> a
lookup tag = go
  where
  go (Cons label labels) (Cons value values)
    | label == tag = value
    | otherwise = go labels values
  go _ _ = unsafeCrashWith "missing function"

equal :: String -> Int -> Int -> Boolean
equal tag x y =
  lookup tag (Cons "eq" (Cons "ne" Nil)) (Cons (==) (Cons (/=) Nil)) x y

render :: String -> Int -> String
render tag value =
  lookup tag (Cons "plain" (Cons "marked" Nil))
    (Cons show (Cons (\n -> "value=" <> show n) Nil)) value

order :: String -> Int -> Int -> Ordering
order tag x y =
  lookup tag (Cons "ascending" (Cons "descending" Nil))
    (Cons compare (Cons (flip compare) Nil)) x y

main :: Effect Unit
main = do
  logShow (equal "eq" 42 42)
  logShow (equal "eq" 42 41)
  logShow (equal "ne" 42 41)
  logShow (render "plain" 42)
  logShow (render "marked" 42)
  logShow (order "ascending" 1 2)
  logShow (order "descending" 1 2)

