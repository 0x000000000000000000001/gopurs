module Gopurs.ReboxMetadata
  ( ReboxFields
  , ReboxFieldIndex
  , buildReboxFieldIndex
  ) where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (fromMaybe)
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..))
import Gopurs.ClassMetadata (ClassFields)
import Gopurs.ConstructorMetadata (ConstructorTypes)
import Gopurs.GoAst (sanitizeName)
import PureScript.Backend.Optimizer.CoreFn (ExprType)

type ReboxFields =
  { vars :: Array String
  , fields :: Array ExprType
  }

type ReboxFieldIndex = Map String ReboxFields

-- Match the former ascending scans: the first key wins, and constructors
-- take precedence over classes even when their generated names collide.
buildReboxFieldIndex :: ConstructorTypes -> ClassFields -> ReboxFieldIndex
buildReboxFieldIndex ctorTypes classFields =
  let
    constructors = Array.foldl
      (\index (Tuple key fields) -> addFields key fields index)
      Map.empty
      (Map.toUnfoldable ctorTypes :: Array (Tuple String ReboxFields))
  in
    Array.foldl
      (\index (Tuple key info) -> addFields key { vars: info.vars, fields: map (\field -> field."type") info.fields } index)
      constructors
      (Map.toUnfoldable classFields)

addFields :: String -> ReboxFields -> ReboxFieldIndex -> ReboxFieldIndex
addFields key fields index =
  let
    parts = String.split (Pattern ".") key
  in
    if Array.length parts < 2 then index
    else
      let
        ctorName = fromMaybe "" (Array.last parts)
        pkgName = String.joinWith "_" (Array.slice 0 (Array.length parts - 1) parts)
        suffix = pkgName <> "_" <> sanitizeName ctorName
        insertFirst name entries =
          if Map.member name entries then entries
          else Map.insert name fields entries
      in
        insertFirst ("Data_" <> suffix) (insertFirst ("Constructor_" <> suffix) index)
