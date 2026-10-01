module Gopurs.WorkerNames
  ( Names
  , fromModule
  , fresh
  , reserve
  ) where

import Prelude
import Data.Array as Array
import Data.Map as Map
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (fst)
import Gopurs.GoAst (sanitizeName)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (Ident(..))

-- A distinct PureScript name can still collide after Go sanitization. Keep both
-- inventories, including foreign declarations, throughout the ordered pass.
type Names = { source :: Set Ident, emitted :: Set String }

fromModule :: BackendModule -> Names
fromModule mod =
  let source = Set.fromFoldable (Array.concatMap (map fst <<< _.bindings) mod.bindings) <> Map.keys mod.foreign
  in { source, emitted: Set.map emittedName source }

fresh :: String -> Ident -> Names -> Ident
fresh suffix (Ident original) names = choose 0
  where
  choose index =
    let candidate = Ident (original <> suffix <> show index)
    in if Set.member candidate names.source || Set.member (emittedName candidate) names.emitted
      then choose (index + 1)
      else candidate

reserve :: Ident -> Names -> Names
reserve name names =
  { source: Set.insert name names.source, emitted: Set.insert (emittedName name) names.emitted }

emittedName :: Ident -> String
emittedName (Ident name) = sanitizeName name
