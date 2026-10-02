module Gopurs.ClosedDictionaries (cacheClosedDictionaries) where

import Prelude

import Control.Monad.State (State, get, put, runState)
import Data.Array as Array
import Data.Foldable (foldl)
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.ClosedDictionaries.Admission as Admission
import Gopurs.CodegenState (CodegenMetadata)
import PureScript.Backend.Optimizer.Convert (BackendBindingGroup, BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType, Ident(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

type RewriteContext =
  { moduleName :: ModuleName
  , metadata :: CodegenMetadata
  , shared :: Admission.SharedDictionaries
  }

type LiftState =
  { next :: Int
  , names :: Set.Set String
  , lifted :: Array (Tuple Ident NeutralExpr)
  }

-- Closed dictionary constructions belong in lazy module-level cached getters.
-- Reuse proven original bindings first, then lift the outermost annotated sites.
-- Traversal order fixes fresh names and the order of the final binding group.
cacheClosedDictionaries :: CodegenMetadata -> BackendModule -> BackendModule
cacheClosedDictionaries metadata mod =
  let
    existing = foldl
      (\acc (Tuple (Ident name) _) -> Set.insert name acc)
      Set.empty
      (Array.concatMap _.bindings mod.bindings)

    context = { moduleName: mod.name, metadata, shared: Admission.sharedDictionaries metadata mod }
    initial = { next: 0, names: existing, lifted: [] }

    Tuple rewritten liftedState = runState (traverse (rewriteGroup context) mod.bindings) initial
  in
    if Array.null liftedState.lifted then mod { bindings = rewritten }
    else mod
      { bindings = rewritten <> [ { recursive: false, bindings: liftedState.lifted } ]
      }

rewriteGroup :: RewriteContext -> BackendBindingGroup Ident NeutralExpr -> State LiftState (BackendBindingGroup Ident NeutralExpr)
rewriteGroup context group
  | group.recursive = pure group
  | otherwise = do
      bindings <- traverse
        (\(Tuple ident expr) -> Tuple ident <$> rewriteBelowRoot context expr)
        group.bindings
      pure group { bindings = bindings }

-- A top-level body already has a getter. A rejected annotation also protects
-- its root from being reconsidered without that annotation. In both cases,
-- preserve all wrappers and visit only descendants; LetRec is a hard barrier.
rewriteBelowRoot :: RewriteContext -> NeutralExpr -> State LiftState NeutralExpr
rewriteBelowRoot context original@(NeutralExpr syn) =
  case syn of
    LetRec _ _ _ -> pure original
    Typed ty body -> NeutralExpr <<< Typed ty <$> rewriteBelowRoot context body
    TypeApp body ty -> NeutralExpr <<< flip TypeApp ty <$> rewriteBelowRoot context body
    _ -> NeutralExpr <$> traverse (rewriteExpr context) syn

rewriteExpr :: RewriteContext -> NeutralExpr -> State LiftState NeutralExpr
rewriteExpr context expr = case Admission.reusable context.shared expr of
  Just known -> pure (reference context.moduleName known.ident known.type)
  Nothing -> case Admission.liftable context.metadata expr of
    Just ty -> liftDictionary context.moduleName ty expr
    Nothing -> rewriteBelowRoot context expr

reference :: ModuleName -> Ident -> ExprType -> NeutralExpr
reference moduleName ident ty = NeutralExpr (Typed ty (NeutralExpr (Var (Qualified (Just moduleName) ident))))

-- Store the accepted expression verbatim: its children are not rewritten and
-- this new site is not deduplicated with other lifts or added to the reuse index.
liftDictionary :: ModuleName -> ExprType -> NeutralExpr -> State LiftState NeutralExpr
liftDictionary moduleName ty expr = do
  st <- get
  let
    allocated = allocate st.names st.next
    ident = Ident allocated.name
  put st
    { next = allocated.next
    , names = Set.insert allocated.name st.names
    , lifted = Array.snoc st.lifted (Tuple ident expr)
    }
  pure (reference moduleName ident ty)

allocate :: Set.Set String -> Int -> { name :: String, next :: Int }
allocate names index =
  let
    candidate = "__cached_dict_" <> show index
  in
    if Set.member candidate names then allocate names (index + 1)
    else { name: candidate, next: index + 1 }
