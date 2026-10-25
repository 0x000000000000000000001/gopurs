module Gopurs.Cache.Pool (Pool, create, intern, internRef, finish) where

import Effect (Effect)

-- One pool per portable artifact. Only canonical JSON fragments cross this
-- boundary; runtime constructor tags, pointers and values are never persisted.
foreign import data Pool :: Type
foreign import create :: Effect Pool
foreign import intern :: Pool -> String -> Effect Int
foreign import internRef :: Pool -> String -> Effect String
foreign import finish :: Pool -> Effect String
