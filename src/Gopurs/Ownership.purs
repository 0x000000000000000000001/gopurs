-- | Consuming workers for a closed, first-order forest language. These proofs
-- | belong to the final backend IR and never use source usage annotations.
module Gopurs.Ownership (prepare) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, foldMap)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Newtype (unwrap)
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (CodegenMetadata, FunctionInfo)
import Gopurs.GoAst (GoDecl, sanitizeName)
import Gopurs.Ownership.Analysis (freshTree, qualify, spine, strip)
import Gopurs.Ownership.Candidates as Candidates
import Gopurs.Ownership.Types (Candidate, Candidates)
import Gopurs.Ownership.Workers as Workers
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

-- Admission and name reservation precede the consuming-contract fixed point.
-- Publish declarations/signatures only for survivors, then rewrite fresh entry
-- calls. Ordinary function bindings remain available with persistent semantics.
prepare :: CodegenMetadata -> BackendModule ->
  { module :: BackendModule, declarations :: Array GoDecl, functions :: Map.Map String FunctionInfo }
prepare metadata mod =
  let initial = Candidates.collect metadata mod
      accepted = validateCandidates metadata mod.name initial
      workers = Array.fromFoldable (Map.values accepted)
      declarations = foldMap (fromMaybe [] <<< Workers.declarations metadata mod.name accepted) workers
      functions = functionSignatures workers
      bindings = map (\group -> group { bindings = map (\(Tuple name expr) ->
        Tuple name (rewrite metadata mod.name accepted expr)) group.bindings }) mod.bindings
  in { module: mod { bindings = bindings }, declarations, functions }

-- Removing a rejected function invalidates every caller which depended on its
-- consuming contract. The fixed point includes mutually recursive families.
-- Workers.declarations is an all-or-nothing proof/emission attempt; rerunning it
-- on the accepted set preserves the same per-worker fresh-name sequence.
validateCandidates :: CodegenMetadata -> ModuleName -> Candidates -> Candidates
validateCandidates metadata moduleName candidates =
  let accepted = Map.filter (isJust <<< Workers.declarations metadata moduleName candidates) candidates
  in if Map.size accepted == Map.size candidates then accepted else validateCandidates metadata moduleName accepted

functionSignatures :: Array Candidate -> Map.Map String FunctionInfo
functionSignatures = Map.fromFoldable <<< map (\fn -> Tuple (sanitizeName $ unwrap fn.worker)
  { fullName: fn.native, fArgs: fn.argTypes, fRet: fn.spec.goType, arity: Array.length fn.args })

-- Every tree argument must be fresh on entry. Scalar arguments retain their
-- ordinary evaluation; only the closed constructor trees need this proof.
freshCall :: CodegenMetadata -> ModuleName -> Candidates -> NeutralExpr -> Maybe { fn :: Candidate, args :: Array NeutralExpr }
freshCall metadata moduleName candidates expr = do
  let call = spine expr
  fn <- case strip call.head of
    Var qualified -> case qualify moduleName qualified of
      Qualified (Just mod) name | mod == moduleName -> Map.lookup name candidates
      _ -> Nothing
    _ -> Nothing
  guard (Array.length call.args == Array.length fn.args)
  let context = { metadata, moduleName, candidates, candidate: fn }
  guard (all (\(Tuple ty value) -> ty /= fn.spec.goType || freshTree context value)
    (Array.zip fn.argTypes call.args))
  pure { fn, args: call.args }

rewrite :: CodegenMetadata -> ModuleName -> Candidates -> NeutralExpr -> NeutralExpr
rewrite metadata moduleName candidates original@(NeutralExpr syn) =
  let choose = do
        { fn, args: supplied } <- freshCall metadata moduleName candidates original
        args <- NEA.fromArray (map (rewrite metadata moduleName candidates) supplied)
        let signature = Func (map (\ty -> if ty == fn.spec.goType then fn.spec.ty else Any) fn.argTypes) fn.spec.ty
        pure $ NeutralExpr $ Typed fn.spec.ty $ NeutralExpr $ App
          (NeutralExpr $ Typed signature $ NeutralExpr $ Var $ Qualified (Just moduleName) fn.worker) args
  in fromMaybe (NeutralExpr $ map (rewrite metadata moduleName candidates) syn) choose
