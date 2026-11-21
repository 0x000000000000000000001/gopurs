module Main where

-- @dependencies: console prelude

import Prelude
import Effect.Console

infixl 4 what as ?!

what :: forall a b. a -> b -> a
what a _ = a

main = log $ "Done" ?! true
