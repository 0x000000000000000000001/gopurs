module Preparation.Provider where

import Prelude

first :: Int -> Int
first value = value + 1

second :: Int -> Int
second value = value + 2

method :: { apply :: Int -> Int }
method = { apply: first }
