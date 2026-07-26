module TestUnify where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import PureScript.Backend.Optimizer.Substitute (unify)
import PureScript.Backend.Optimizer.Syntax (ExprType(..))

polyBindType :: ExprType
polyBindType = Func [Record [Map.Tuple "bind" (Func [TypeVar "m", Func [TypeVar "a", TypeVar "m"]] (TypeVar "m"))], TypeVar "m", Func [TypeVar "a", TypeVar "m"]] (TypeVar "m")

testUnify :: Effect Unit
testUnify = do
  let subst = unify polyBindType polyBindType Map.empty
  log ("Is empty? " <> show (Map.isEmpty subst))
