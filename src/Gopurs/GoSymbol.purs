module Gopurs.GoSymbol
  ( GoSymbol
  , GoReference(..)
  , generated
  , foreignSymbol
  , referenceName
  , Placement
  , RenderContext
  , renderReference
  ) where

import Prelude

import Data.Either (Either(..))
import Data.Generic.Rep (class Generic)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))

-- A source unit owns a symbol before the package plan places it. This is not
-- a Go import path, a PureScript type identity, or a runtime constructor tag.
-- Unit names are supplied by producers, never split back out of a Go name.
type GoSymbol = { owner :: String, name :: String }

data GoReference = LocalReference String | SymbolReference GoSymbol

derive instance Eq GoReference
derive instance Ord GoReference
derive instance Generic GoReference _

generated :: String -> String -> GoSymbol
generated unit name = { owner: "generated/" <> unit, name }

foreignSymbol :: String -> String -> GoSymbol
foreignSymbol unit name = { owner: "ffi/" <> unit, name }

referenceName :: GoReference -> String
referenceName (LocalReference name) = name
referenceName (SymbolReference symbol) = symbol.name

type Placement = { packagePath :: String, name :: String }
type RenderContext =
  { consumer :: String
  , symbols :: Map GoSymbol Placement
  , aliases :: Map String String
  }

-- The monolithic printer uses referenceName. A partitioned render must supply
-- the complete plan: a missing symbol or alias is an error, not a local name.
renderReference :: RenderContext -> GoReference -> Either String String
renderReference _ (LocalReference name) = Right name
renderReference context (SymbolReference symbol) = case Map.lookup symbol context.symbols of
  Nothing -> Left ("Unplaced Go symbol: " <> symbol.owner <> ":" <> symbol.name)
  Just target | target.packagePath == context.consumer -> Right target.name
  Just target -> case Map.lookup target.packagePath context.aliases of
    Nothing -> Left ("Missing Go import alias: " <> target.packagePath)
    Just alias -> Right (alias <> "." <> target.name)
