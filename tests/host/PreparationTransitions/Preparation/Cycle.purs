module Preparation.Cycle where

import Prelude

forward :: forall a. Int -> a -> a
forward count value = if count <= 0 then value else backward (count - 1) value

backward :: forall a. Int -> a -> a
backward count value = if count <= 0 then value else forward (count - 1) value
