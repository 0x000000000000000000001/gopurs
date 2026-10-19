module Preparation.Apply where

applyMethod :: forall a. { apply :: a -> a } -> a -> a
applyMethod method value = method.apply value
