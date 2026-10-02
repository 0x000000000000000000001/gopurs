module Gopurs.NativeRecordArgs.Projection
  ( projectableFields
  , argumentType
  , shareableResult
  ) where

import Prelude

import Data.Array as Array
import Data.Foldable (all)
import Data.Maybe (Maybe(..))
import Data.String.CodeUnits as CodeUnits
import Data.Tuple (Tuple(..), fst)
import Gopurs.GoAst (GoType(..), sanitizeName)
import Gopurs.GoConversions (getUnboxedADT)
import Gopurs.GoTypes (visibleRecordFields)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

-- Only known scalar fields of a variable-tailed row can form a projection.
-- Resolve duplicate source labels before checking types and Go-name collisions;
-- the first visible field wins, then source labels determine layout order.
projectableFields :: ExprType -> Maybe (Array (Tuple String ExprType))
projectableFields (Record (Row fields (Just (TypeVar _)))) =
  let
    visible = visibleRecordFields fields
    names = map (sanitizeName <<< fst) visible
  in if not (Array.null visible) && all (\(Tuple _ ty) -> isScalar ty) visible
      && all usableFieldName names && Array.length (Array.nub names) == Array.length names then
    Just (Array.sortBy (comparing fst) visible)
  else Nothing
projectableFields _ = Nothing

-- Keep source labels here: Go rendering owns their sanitization. The use proof
-- must establish that the omitted tail cannot escape before selecting this ABI.
argumentType :: (ExprType -> GoType) -> Array (Tuple String ExprType) -> GoType
argumentType toGoType fields = TypeRecord (map (\(Tuple label ty) -> Tuple label (toGoType ty)) fields)

-- Quoted labels may not be usable Go fields even after sanitization. Unicode
-- names conservatively retain the ordinary ABI, as do colliding sanitized names.
usableFieldName :: String -> Boolean
usableFieldName name = name /= "_" && all
  (\char -> char == '_' || (char >= 'a' && char <= 'z')
    || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9'))
  (CodeUnits.toCharArray name)

isScalar :: ExprType -> Boolean
isScalar = case _ of
  Int -> true
  Number -> true
  String -> true
  Char -> true
  Boolean -> true
  _ -> false

-- Native Maybe/Either/Tuple results already have a worker ABI. Their payload
-- representation belongs to GoConversions, not to the argument projection.
shareableResult :: ExprType -> Boolean
shareableResult result = isScalar result || case getUnboxedADT result of
  Just _ -> true
  Nothing -> false
