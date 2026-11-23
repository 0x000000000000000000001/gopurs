module Gopurs.SharedContainerLayout (canonicalContainerArguments) where

import Prelude

import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.AdtMetadata (PointerAdtPaths)
import Gopurs.GoAst (GoType(..), constructorNames)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

-- Source preparation selects products observed at higher-order boundaries.
-- Sharing their generic payloads avoids recursive copies of containing lists.
-- Purely native products retain their native fields. This selects a calling
-- representation; it is not an ownership or escape proof.
canonicalContainerArguments :: PointerAdtPaths -> Set.Set String -> Set.Set String -> String -> Array ExprType -> Array GoType -> Array GoType
canonicalContainerArguments paths enums elided name args mapped =
  if Array.length args /= Array.length mapped then mapped
  else if sharedPair name args then map (const TypeValue) mapped
  else case args of
    [ element ] | name == "Data.List.Types.List" && pointer name 1 -> case applied element [] of
      Just pair | sharedPair pair.name pair.args -> map (const TypeValue) mapped
      _ -> mapped
    _ -> mapped
  where
  pointer fullName arity = case Map.lookup fullName paths of
    Just info | info.arity == arity ->
      let
        parts = String.split (Pattern ".") fullName
        moduleName = String.joinWith "." (Array.take (Array.length parts - 1) parts)
        prefix = String.replaceAll (Pattern ".") (Replacement "_") moduleName
        names = constructorNames prefix info.ctorName
        legacyNames = constructorNames moduleName info.ctorName
      in not (Set.member names.structName elided || Set.member legacyNames.structName elided)
    _ -> false

  sharedPair fullName elements = fullName == "Data.Tuple.Tuple"
    && pointer fullName 2
    && case elements of
      [ record, tag ] -> scalarRecord record && enum tag && selected elements
      _ -> false

  selected elements = case Map.lookup "Data.Tuple.Tuple" paths of
    Just info -> Array.any (samePair elements) info.sharedPairs
    _ -> false

  samePair left right = case pairShape left, pairShape right of
    Just a, Just b -> a == b
    _, _ -> false

  pairShape [ Record (Row fields Nothing), tag ] = case applied tag [] of
    Just adt | Array.null adt.args -> Just
      { fields: Array.sortBy (comparing \(Tuple label _) -> label)
          (Array.nubBy (comparing \(Tuple label _) -> label) fields)
      , tag: adt.name
      }
    _ -> Nothing
  pairShape _ = Nothing

  enum ty = case applied ty [] of
    Just adt -> Array.null adt.args && Set.member adt.name enums
    _ -> false

  scalarRecord (Record (Row fields Nothing)) = not (Array.null fields)
    && Array.all (\(Tuple _ ty) -> scalar ty) fields
  scalarRecord _ = false

  scalar Int = true
  scalar Number = true
  scalar Boolean = true
  scalar Char = true
  scalar String = true
  scalar _ = false

  applied (TypeApp fn extra) collected = applied fn (extra <> collected)
  applied (ADT fullName _ original) collected = Just { name: fullName, args: original <> collected }
  applied _ _ = Nothing
