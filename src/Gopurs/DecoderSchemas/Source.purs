module Gopurs.DecoderSchemas.Source
  ( DictionaryContext
  , strip
  , replace
  , standard
  , resolve
  , erased
  , symbolName
  ) where

import Prelude

import Data.Array.NonEmpty as NEA
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Literal(..), ModuleName(..), Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

-- The caller indexes original, nonrecursive module bindings once. Generated
-- source getters never enter this environment during rewriting.
type DictionaryContext =
  { owner :: ModuleName
  , definitions :: Map.Map Ident NeutralExpr
  }

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syn) = case syn of
  Typed _ inner -> strip inner
  TypeApp inner _ -> strip inner
  _ -> syn

replace :: NeutralExpr -> NeutralExpr -> NeutralExpr
replace (NeutralExpr syn) replacement = case syn of
  Typed ty inner -> NeutralExpr (Typed ty (replace inner replacement))
  TypeApp inner ty -> NeutralExpr (TypeApp (replace inner replacement) ty)
  _ -> replacement

standard :: String -> NeutralExpr -> Boolean
standard name expr = case strip expr of
  Var (Qualified (Just (ModuleName "Data.Argonaut.Decode.Class")) (Ident found)) -> name == found
  _ -> false

-- Flatten applications and qualified aliases with a bounded proof. Keep the
-- actual construction in its source getter, including application/argument order.
resolve :: DictionaryContext -> Int -> NeutralExpr -> Maybe (Tuple NeutralExpr (Array NeutralExpr))
resolve context fuel expr
  | fuel <= 0 = Nothing
  | otherwise = case strip expr of
      Var (Qualified (Just moduleName) name) | moduleName == context.owner -> case Map.lookup name context.definitions of
        Just value -> resolve context (fuel - 1) value
        Nothing -> Just (Tuple expr [])
      App fn args -> do
        Tuple head prior <- resolve context (fuel - 1) fn
        pure (Tuple head (prior <> NEA.toArray args))
      _ -> Just (Tuple expr [])

erased :: NeutralExpr -> Boolean
erased expr = case strip expr of
  PrimUndefined -> true
  _ -> false

symbolName :: DictionaryContext -> Int -> NeutralExpr -> Maybe String
symbolName context fuel expr = do
  Tuple head args <- resolve context fuel expr
  case args, strip head of
    [], Lit (LitRecord [ Prop "reflectSymbol" method ]) -> case strip method of
      Abs refs body | NEA.length refs == 1 -> case strip body of
        Lit (LitString value) -> Just value
        _ -> Nothing
      _ -> Nothing
    _, _ -> Nothing
