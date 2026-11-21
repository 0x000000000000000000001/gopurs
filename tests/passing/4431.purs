module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console (log)

data TypedCache :: (Type -> Type) -> Type -> Type
data TypedCache key a = Get (key a)

derive instance Functor k => Functor (TypedCache k)

main = log "Done"
