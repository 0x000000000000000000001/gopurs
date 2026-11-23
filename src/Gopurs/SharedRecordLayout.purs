module Gopurs.SharedRecordLayout (collectSharedRecordShapes, isSharedRecord) where

import Prelude

import Data.Array as Array
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), isJust)
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.AdtMetadata (PointerAdtPaths)
import Gopurs.SharedContainerUsage (foldFunctionTypes)
import PureScript.Backend.Optimizer.CoreFn (Ann, ExprType(..), Module(..))

type Fields = Array (Tuple String ExprType)
type Declaration = { arity :: Int, boxed :: Boolean }

-- A source function-value boundary is evidence of repeated record conversion,
-- not a proof of escape, ownership or a gain for every program. Restrict the
-- policy to flat nested records surrounding already boxed ADT payloads. Scalar
-- aggregates, native pointers, arrays and unknown fields retain native layouts.
collectSharedRecordShapes :: PointerAdtPaths -> Set String -> Array (Module Ann) -> Set String
collectSharedRecordShapes pointers enums modules =
  let
    observed = foldFunctionTypes
      { value: \ty acc -> acc { selected = collect ty acc.selected }
      , recursive: \ty acc -> acc { nativeLoops = collectNativeLoop ty acc.nativeLoops }
      }
      { selected: Set.empty, nativeLoops: Set.empty }
      modules
  in Set.difference observed.selected observed.nativeLoops
  where
  declarations :: Map.Map String Declaration
  declarations = Map.fromFoldable (Array.concatMap moduleDeclarations modules)

  moduleDeclarations (Module mod) = Array.mapMaybe (\declaration ->
    let name = unwrap mod.name <> "." <> declaration.name
    in
      if Map.member name pointers then Nothing
      else Just (Tuple name
        { arity: Array.length declaration.vars
        , boxed: not (Set.member name enums)
        })) mod.dataDecls

  leaf ty | scalar ty = Just false
  leaf ty = case applied ty [] of
    Just adt -> case Map.lookup adt.name declarations of
      Just declaration | declaration.arity == Array.length adt.args -> Just declaration.boxed
      _ -> Nothing
    Nothing -> Nothing

  fieldAllowed (Tuple _ (Record (Row fields Nothing))) =
    let visible = visibleFields fields
    in not (Array.null visible) && Array.all (\(Tuple _ ty) -> isJust (leaf ty)) visible
  fieldAllowed (Tuple _ ty) = isJust (leaf ty)

  hasBoxedNestedField (Tuple _ (Record (Row fields Nothing))) =
    Array.any (\(Tuple _ ty) -> leaf ty == Just true) (visibleFields fields)
  hasBoxedNestedField _ = false

  select fields selected =
    let visible = visibleFields fields
    in
      if Array.all fieldAllowed visible && Array.any hasBoxedNestedField visible then
        case recordKey fields of
          Just key -> Set.insert key selected
          Nothing -> selected
      else selected

  collect ty selected = case ty of
    Func args result -> collect result (foldl (flip collect) selected args)
    ForAll _ body -> collect body selected
    ConstrainedType _ body -> collect body selected
    Array element -> collect element selected
    Record (Row fields Nothing) -> collectFields fields (select fields selected)
    Record row -> collect row selected
    Row fields _ -> collectFields fields selected
    _ -> case applied ty [] of
      Just adt -> foldl (flip collect) selected adt.args
      Nothing -> selected

  collectFields fields selected = foldl (\acc (Tuple _ ty) -> collect ty acc) selected (visibleFields fields)

  -- A recursive worker carrying both a scalar control/state parameter and a
  -- direct record argument may perform many native updates after one callback.
  -- Preserve that family's native ABI even if a function-value use selected it.
  -- This conservative guard is independent of the result type; it protects
  -- reductions and loops returning their accumulator. It is not a complete
  -- frequency analysis or a universal performance guarantee.
  collectNativeLoop ty selected =
    let args = directArguments ty
    in
      if Array.any scalar args then foldl collectDirectRecord selected args
      else selected

  collectDirectRecord selected (Record (Row fields Nothing)) = case recordKey fields of
    Just key -> Set.insert key selected
    Nothing -> selected
  collectDirectRecord selected _ = selected

  directArguments (ForAll _ body) = directArguments body
  directArguments (ConstrainedType _ body) = directArguments body
  directArguments (Func args result) = args <> directArguments result
  directArguments _ = []

isSharedRecord :: Set String -> Fields -> Boolean
isSharedRecord selected fields =
  if Set.isEmpty selected then false
  else case recordKey fields of
    Just key -> Set.member key selected
    Nothing -> false

-- The selection above proves the declaration and exact arity of every ADT
-- leaf. Its boxed/enum representation is independent of the arguments, so the
-- key keeps the ADT identity and arity rather than scoped type variables. Bare
-- unknown fields cannot form a key. Length prefixes keep arbitrary labels and
-- type names distinct without depending on a host-specific hash.
recordKey :: Fields -> Maybe String
recordKey = fieldsKey true

fieldsKey :: Boolean -> Fields -> Maybe String
fieldsKey nested fields =
  let visible = visibleFields fields
  in
    if Array.null visible then Nothing
    else map (\parts -> "r(" <> String.joinWith "" parts <> ")") (traverse fieldKey visible)
  where
  fieldKey (Tuple name ty) = map (\key -> atom name <> key) (valueKey ty)
  valueKey (Record (Row innerFields Nothing)) | nested = fieldsKey false innerFields
  valueKey ty = case scalarKey ty of
    Just key -> Just key
    Nothing -> case applied ty [] of
      Just adt -> Just ("a" <> atom adt.name <> show (Array.length adt.args) <> ";")
      Nothing -> Nothing

atom :: String -> String
atom value = show (String.length value) <> ":" <> value

visibleFields :: Fields -> Fields
visibleFields = Array.sortBy (comparing \(Tuple label _) -> label)
  <<< Array.nubBy (comparing \(Tuple label _) -> label)

scalar :: ExprType -> Boolean
scalar = isJust <<< scalarKey

scalarKey :: ExprType -> Maybe String
scalarKey Int = Just "i"
scalarKey Number = Just "n"
scalarKey Boolean = Just "b"
scalarKey Char = Just "c"
scalarKey String = Just "s"
scalarKey _ = Nothing

applied :: ExprType -> Array ExprType -> Maybe { name :: String, args :: Array ExprType }
applied (TypeApp fn extra) collected = applied fn (extra <> collected)
applied (ADT name _ original) collected = Just { name, args: original <> collected }
applied _ _ = Nothing
