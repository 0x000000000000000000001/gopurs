module Gopurs.FfiSupport
  ( findFfiFile
  ) where

import Prelude

import Effect (Effect)
import Data.Maybe (Maybe)
import Data.Nullable (Nullable, toNullable, toMaybe)

foreign import findFfiFileImpl :: Nullable String -> String -> Nullable String -> Effect (Nullable String)

-- | Finds the .go FFI file corresponding to a PureScript module
-- | mbFfiDir: optional directory to search in (if not provided, searches .spago and local dirs)
-- | modName: the PureScript module name (e.g. "Data.Show")
-- | mbModulePath: the path to the original .purs file if known
findFfiFile :: Maybe String -> String -> Maybe String -> Effect (Maybe String)
findFfiFile mbFfiDir modName mbModulePath = do
  path <- findFfiFileImpl (toNullable mbFfiDir) modName (toNullable mbModulePath)
  pure (toMaybe path)
