module Gopurs.DecoderSchemas.Admission (admitSchema) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.DecoderSchemas.Programs (deriveProgram)
import Gopurs.DecoderSchemas.Source (DictionaryContext, erased, resolve, standard, strip, symbolName)
import Gopurs.DecoderSchemas.Types (Schema(..), tagged, weight)
import PureScript.Backend.Optimizer.CoreFn (Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

-- Only actual standard dictionary applications justify a worker. Types alone
-- do not identify a decoder, and scalar/opaque roots do not warrant a worker.
admitSchema :: DictionaryContext -> NeutralExpr -> Maybe Schema
admitSchema context expr = case strip expr of
  App fn args | standard "decodeJson" fn && closed expr -> case NEA.toArray args of
    [ dictionary ] -> do
      schema <- decoder context 64 dictionary
      let size = weight schema
      guard (size > 1 && size <= 128 && tagged schema)
      pure schema
    _ -> Nothing
  _ -> Nothing

-- This proof excludes every local, even a bound one: it is a closed dictionary
-- construction, not general lexical closure analysis. Effects also stop admission.
closed :: NeutralExpr -> Boolean
closed (NeutralExpr syn) = case syn of
  Local _ _ -> false
  LetRec _ _ _ -> false
  PrimEffect _ -> false
  EffectBind _ _ _ _ -> false
  EffectPure _ -> false
  EffectDefer _ -> false
  UncurriedEffectAbs _ _ -> false
  UncurriedEffectApp _ _ -> false
  _ -> all closed syn

decoder :: DictionaryContext -> Int -> NeutralExpr -> Maybe Schema
decoder context fuel expr = do
  Tuple head args <- resolve context fuel expr
  let recur = decoder context (fuel - 1)
  case args of
    [] | standard "decodeJsonInt" head -> pure (Scalar "Int")
    [] | standard "decodeJsonNumber" head -> pure (Scalar "Number")
    [] | standard "decodeJsonString" head -> pure (Scalar "String")
    [] | standard "decodeJsonBoolean" head -> pure (Scalar "Boolean")
    [] | standard "decodeJsonJson" head -> pure (Scalar "Json")
    [ inner ] | standard "decodeArray" head -> Sequence <$> recur inner
    [ inner ] | standard "decodeJsonMaybe" head -> Optional <$> recur inner
    [ row, proxy ] | standard "decodeRecord" head && erased proxy -> ObjectSchema <$> record context (fuel - 1) row
    [] -> case strip expr of
      Var (Qualified (Just moduleName) ident) -> do
        guard (moduleName /= context.owner || Map.member ident context.definitions)
        guard (case strip head of
          Var (Qualified (Just resolvedOwner) resolvedName) | resolvedOwner == context.owner -> Map.member resolvedName context.definitions
          _ -> true)
        pure (fromMaybe Custom (Derived <$> customProgram (fuel - 1) head))
      _ -> Nothing
    _ -> Nothing
  where
  -- Supply recursive dictionary proofs to the body analysis without letting it
  -- own alias resolution or admission budgets. Each read consumes the same fuel
  -- as before; failure keeps the dictionary as an opaque native-plan leaf.
  customProgram remaining head = do
    guard (remaining > 0)
    deriveProgram
      { resolveDictionary: resolve context remaining
      , readSchema: decoder context (remaining - 1)
      }
      head

record :: DictionaryContext -> Int -> NeutralExpr -> Maybe (Array (Tuple String Schema))
record context fuel expr = do
  Tuple head args <- resolve context fuel expr
  case args of
    [] | standard "gDecodeJsonNil" head -> pure []
    [ field, tail, symbol, cons, lacks ] | standard "gDecodeJsonCons" head && erased cons && erased lacks -> do
      key <- symbolName context (fuel - 1) symbol
      Tuple fieldHead fieldArgs <- resolve context (fuel - 1) field
      schema <- case fieldArgs of
        [ inner ] | standard "decodeFieldId" fieldHead -> decoder context (fuel - 1) inner
        [ inner ] | standard "decodeFieldMaybe" fieldHead -> Optional <$> decoder context (fuel - 1) inner
        _ -> Nothing
      rest <- record context (fuel - 1) tail
      pure (Array.cons (Tuple key schema) rest)
    _ -> Nothing
