module DumpInst where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.Monomorphize (InstantiationMap)
import PureScript.Backend.Optimizer.Syntax (ExprType(..))

dumpInst :: InstantiationMap -> Effect Unit
dumpInst insts = do
  case Map.lookup "Control.Bind.bind" insts of
    Just types -> log ("Control.Bind.bind has: " <> show types)
    Nothing -> log "Control.Bind.bind not found!"
