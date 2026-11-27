module Gopurs.Ownership.Continuations (share) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (StateT, get, put, runStateT)
import Control.Monad.Trans.Class (lift)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldMap)
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (GoExpr(..), rawGo)
import Gopurs.Ownership.Analysis (strip)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Pair(..))

type Shared = { blocks :: Array (Array GoExpr), leaves :: Int, saved :: Int }
type Sharing = StateT Shared Maybe

-- Only split a root decision region: its internal nodes contain conditions,
-- never bindings, writes or calls which can donate cells. The supplied terminal
-- emitter proves every leaf independently and starts its temporary counter at
-- zero, so structurally equal bodies already have identical local names.
-- A label starts a separate block; its variables cannot enter another leaf's
-- scope. No test, read, write, allocation or consuming call moves in time.
share
  :: (NeutralExpr -> Maybe GoExpr)
  -> (NeutralExpr -> Maybe (Array GoExpr))
  -> NeutralExpr
  -> Maybe (Array GoExpr)
share condition terminal original = do
  guard (case strip original of
    Branch _ _ -> true
    _ -> false)
  Tuple tree state <- runStateT (visit original) { blocks: [], leaves: 0, saved: 0 }
  -- Avoid extra jumps for tiny regions. This is a code-size threshold, not an
  -- ownership shortcut: every candidate leaf was proved before this decision.
  guard (state.leaves >= 4 && state.saved >= 32)
  pure (tree <> foldMap block (Array.mapWithIndex Tuple state.blocks))
  where
  label index = "__owned_continuation_" <> show index
  block (Tuple index body) = [ rawGo (label index <> ": {") ] <> body <> [ rawGo "}" ]

  visit :: NeutralExpr -> Sharing (Array GoExpr)
  visit expression = case strip expression of
    Branch branches fallback -> do
      cases <- traverse (\(Pair test body) -> do
        testExpr <- lift (condition test)
        statements <- visit body
        pure { testExpr, statements }) (NEA.toArray branches)
      otherwise <- visit fallback
      pure (Array.foldr (\branch rest -> [ GoIfElse branch.testExpr branch.statements rest ]) otherwise cases)
    _ -> do
      body <- lift (terminal expression)
      state <- get
      index <- case Array.elemIndex body state.blocks of
        Just existing -> do
          put (state { leaves = state.leaves + 1, saved = state.saved + Array.length body })
          pure existing
        Nothing -> do
          put (state { leaves = state.leaves + 1, blocks = Array.snoc state.blocks body })
          pure (Array.length state.blocks)
      pure [ rawGo ("goto " <> label index) ]
