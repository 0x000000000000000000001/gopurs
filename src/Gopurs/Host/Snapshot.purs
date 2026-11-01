module Gopurs.Host.Snapshot (withSnapshot, release) where

import Prelude
import Effect (Effect)

foreign import withSnapshot :: Effect Unit -> Effect Unit
foreign import release :: Effect Unit
