-- @inline export selectInts never
module FilterLibrary (Trail(..), selectInts) where

import Prelude

-- Names deliberately differ from the benchmark and from Data.List.
data Trail a = Stop | Cell a (Trail a)

turnAround :: Trail Int -> Trail Int
turnAround input = collect input Stop
  where
  collect Stop result = result
  collect (Cell value rest) result = collect rest (Cell value result)

-- Keep an actual imported worker so its callback ABI has to travel through
-- global metadata, including in the native parallel compiler.
selectInts :: (Int -> Boolean) -> Trail Int -> Trail Int
selectInts predicate input = collect input Stop
  where
  collect Stop result = turnAround result
  collect (Cell value rest) result =
    if predicate value then collect rest (Cell value result)
    else collect rest result
