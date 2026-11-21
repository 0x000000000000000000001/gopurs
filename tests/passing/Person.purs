module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

data Person = Person { name :: String, age :: Number }

showPerson :: Person -> String
showPerson = \p -> case p of
  Person o -> o.name <> ", aged " <> show o.age

main = log "Done"
