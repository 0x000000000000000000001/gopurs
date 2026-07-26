module DumpTypes where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import PureScript.Backend.Optimizer.Syntax (ExprType(..))
import Data.Set as Set
import Data.Array as Array
import PureScript.Backend.Optimizer.Monomorphize (InstantiationMap)

dumpTypes :: InstantiationMap -> Effect Unit
dumpTypes insts = do
  case Map.lookup "Control.Bind.bind" insts of
    Just s -> log ("TYPES: " <> show (Array.fromFoldable s))
    Nothing -> log "NO TYPES"
