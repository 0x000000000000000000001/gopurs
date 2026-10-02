module Gopurs.NativeRecordArgs
  ( projectedArgument
  , candidateToShare
  , workerArguments
  ) where

import Prelude

import Data.Array as Array
import Data.Foldable (all, any)
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (fst)
import Gopurs.GoAst (GoType)
import Gopurs.NativeRecordArgs.Projection as Projection
import Gopurs.NativeRecordArgs.Source as Source
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ann, Binding, ExprType)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendSyntax(Abs, Accessor, EffectDefer, Local, Typed, UncurriedAbs, UncurriedEffectAbs))

-- This is a projection for a proven read-only parameter, not a representation
-- for arbitrary open records. Its tail must never be returned or reconstructed.
projectedArgument :: (ExprType -> GoType) -> ExprType -> Maybe GoType
projectedArgument toGoType ty = map
  (Projection.argumentType toGoType)
  (Projection.projectableFields ty)

-- The source proof controls sharing before PBO, not the final worker ABI.
candidateToShare :: Binding Ann -> Boolean
candidateToShare = Source.candidateToShare

-- Select each parameter independently from the final IR. A stale source proof
-- cannot enable projection, and the result must have a supported native ABI.
-- ModuleBindings publishes these types for both workers and direct callers;
-- partial and dynamic calls retain the ordinary boxed wrapper.
workerArguments :: (ExprType -> GoType) -> Array String -> TcoExpr -> Array ExprType -> ExprType -> Array GoType
workerArguments toGoType names body types result =
  if Projection.shareableResult result then Array.mapWithIndex choose types else map toGoType types
  where
  choose index ty = fromMaybe (toGoType ty) do
    name <- Array.index names index
    fields <- Projection.projectableFields ty
    if readOnlyUses name (map fst fields) body then Just (Projection.argumentType toGoType fields) else Nothing

isTarget :: String -> TcoExpr -> Boolean
isTarget target (TcoExpr _ syntax) = case syntax of
  Typed _ inner -> isTarget target inner
  Local ident level -> localId ident level == target
  _ -> false

containsTarget :: String -> TcoExpr -> Boolean
containsTarget target expr@(TcoExpr _ syntax) =
  isTarget target expr || any (containsTarget target) syntax

-- TCO identities use the emitter's localId (sanitized name and level), rather
-- than source lexical shadowing. Deferred reads must be entirely independent.
readOnlyUses :: String -> Array String -> TcoExpr -> Boolean
readOnlyUses target fields (TcoExpr _ syntax) = case syntax of
  Local ident level -> localId ident level /= target
  Accessor object (GetProp field) | isTarget target object -> Array.elem field fields
  Abs _ body -> not (containsTarget target body)
  UncurriedAbs _ body -> not (containsTarget target body)
  UncurriedEffectAbs _ body -> not (containsTarget target body)
  EffectDefer body -> not (containsTarget target body)
  _ -> all (readOnlyUses target fields) syntax
