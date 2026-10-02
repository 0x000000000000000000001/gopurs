module Gopurs.DecoderSchemas (specializeDecoderSchemas) where

import Prelude

import Control.Monad.State (get, put, runState)
import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..), isJust)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.DecoderSchemas.Admission (admitSchema)
import Gopurs.DecoderSchemas.Source (DictionaryContext, replace)
import Gopurs.DecoderSchemas.Workers as Workers
import Gopurs.GoAst (GoDecl)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Qualified(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

type RewriteState =
  { next :: Int
  , names :: Set.Set String
  , sources :: Array (Tuple Ident NeutralExpr)
  , declarations :: Array GoDecl
  }

-- Admission proves the schema; Workers consumes that proof. This facade owns
-- the ABI gate, traversal, family-name reservation and final source publication.
specializeDecoderSchemas :: CodegenMetadata -> BackendModule -> { module :: BackendModule, declarations :: Array GoDecl }
specializeDecoderSchemas metadata mod
  | not (Map.member "Data.Argonaut.Decode.Internal.Record.schemaDecoderABI1" metadata.globalTypes) = { module: mod, declarations: [] }
  | otherwise =
      let
        context :: DictionaryContext
        context =
          { owner: mod.name
          , definitions: Map.fromFoldable (Array.concatMap (\g -> if g.recursive then [] else g.bindings) mod.bindings)
          }
        names = Set.fromFoldable (map (\(Tuple (Ident name) _) -> name) (Array.concatMap _.bindings mod.bindings))
        initial :: RewriteState
        initial = { next: 0, names, sources: [], declarations: [] }
        Tuple bindings result = runState (traverse (rewriteGroup context) mod.bindings) initial
        sources = if Array.null result.sources then [] else [ { recursive: false, bindings: result.sources } ]
      in { module: mod { bindings = bindings <> sources }, declarations: result.declarations }
  where
  prefix = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
  textEnabled = Map.member "Data.Argonaut.Decode.Parser.textDecoderABI1" metadata.globalTypes

  rewriteGroup context group
    | group.recursive = pure group
    | otherwise = do
        bindings <- traverse (\(Tuple ident expr) -> Tuple ident <$> rewrite context expr) group.bindings
        pure group { bindings = bindings }

  -- Prefer the outermost admitted construction and save it untouched. A rejected
  -- site permits descent, but recursive groups and local LetRec remain barriers.
  rewrite context original@(NeutralExpr syn) = case admitSchema context original of
    Just schema -> allocate schema original
    Nothing -> case syn of
      LetRec _ _ _ -> pure original
      _ -> NeutralExpr <$> traverse (rewrite context) syn

  allocate schema original = do
    state <- get
    let
      name = "__json_schema_" <> show state.next
      source = name <> "_source"
      next = state { next = state.next + 1 }
    -- A pre-existing descendant reserves the whole family, including its text
    -- workers and constructor getters. Failed names still advance the counter.
    if Array.any (\existing -> existing == name || isJust (String.stripPrefix (Pattern (name <> "_")) existing)) (Set.toUnfoldable state.names) then put next *> allocate schema original
    else do
      let generated = Workers.build { prefix, name, source, textEnabled } schema
      put next
        { names = Set.insert source (Set.insert name state.names)
        , sources = state.sources <> [ Tuple (Ident source) (NeutralExpr (Typed CoreFn.Any original)) ] <> generated.sources
        , declarations = state.declarations <> generated.declarations
        }
      pure (replace original (NeutralExpr (Var (Qualified (Just mod.name) (Ident name)))))
