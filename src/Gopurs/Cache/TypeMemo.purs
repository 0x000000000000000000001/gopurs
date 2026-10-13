module Gopurs.Cache.TypeMemo (TypeMemo, create, lookup, insert) where

import Data.Unit (Unit)
import Effect (Effect)
import PureScript.Backend.Optimizer.CoreFn (ExprType)

-- A private shortcut for immutable types already assigned a canonical index.
-- Each artifact owns its memo. Host identity is never written to the artifact;
-- a miss still uses the structural table, including for equal distinct objects.
foreign import data TypeMemo :: Type
foreign import create :: Effect TypeMemo
foreign import lookup :: TypeMemo -> ExprType -> Effect Int
foreign import insert :: TypeMemo -> ExprType -> Int -> Effect Unit
