module Gopurs.FilteredSumLoops (wrap) where

import Prelude

import Control.Alternative (guard)
import Data.Array.NonEmpty as NEA
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Gopurs.CongruenceSum as CongruenceSum
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoFunctions (Parameters)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Literal(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type Recurrence =
  { counter :: String, accumulator :: String
  , lower :: Int, modulus :: Int, residue :: Int
  }

-- Recognize the complete pure recurrence, including both selected and skipped
-- steps, before changing its execution count. Positive signed bounds prevent
-- decrement wraparound and make Go's remainder Euclidean. Wide entry values,
-- nonpositive bounds and every unrecognized expression retain the original loop.
wrap :: (TcoExpr -> Boolean) -> Parameters -> GoType -> TcoExpr -> GoExpr -> GoExpr
wrap self params result body original = case recognize self params body of
  Just recurrence | result == TypeInt64 ->
    let
      counter = GoVar (recurrence.counter <> "_loop")
      accumulator = GoVar (recurrence.accumulator <> "_loop")
      positive = GoBinOp ">=" counter (integer recurrence.lower)
      bounded = GoBinOp "<=" counter (integer top)
      signed = GoBinOp "==" accumulator (cast "int64" (cast "int32" accumulator))
      admitted = GoBinOp "&&" (GoBinOp "&&" positive bounded) signed
      value = GoCall (formula recurrence) [ counter, accumulator ]
    in GoBlock [ GoIfElse admitted [ GoReturn value ] [], original ]
  _ -> original

formula :: Recurrence -> GoExpr
formula { lower, modulus, residue } =
  GoFuncBlock [ Tuple "upper" TypeInt64, Tuple "seed" TypeInt64 ]
    (CongruenceSum.body { modulus, residue }
      { lower: integer lower, upper: GoVar "upper", seed: GoVar "seed" }) TypeInt64

recognize :: (TcoExpr -> Boolean) -> Parameters -> TcoExpr -> Maybe Recurrence
recognize self params body = do
  Tuple first second <- case params of
    [ Tuple first TypeInt64, Tuple second TypeInt64 ] | first /= second -> Just (Tuple first second)
    _ -> Nothing
  branch <- singleBranch body
  boundary <- case syntax branch.condition of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntOrd Syn.OpLt) left right) -> do
      counter <- localName left
      lower <- intLiteral right
      pure { counter, lower }
    _ -> Nothing
  let counter = boundary.counter
  guard (counter == first || counter == second)
  guard (positive32 boundary.lower)
  let accumulator = if counter == first then second else first
  guard (local accumulator branch.yes)
  args <- case syntax branch.no of
    Syn.App target args | self target -> Just (NEA.toArray args)
    Syn.UncurriedApp target args | self target -> Just args
    _ -> Nothing
  Tuple nextCounter nextAccumulator <- case args of
    [ a, b ] -> Just (if counter == first then Tuple a b else Tuple b a)
    _ -> Nothing
  guard case syntax nextCounter of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpSubtract) left right) -> local counter left && intLiteral right == Just 1
    _ -> false
  selection <- singleBranch nextAccumulator
  guard (local accumulator selection.no)
  guard case syntax selection.yes of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpAdd) left right) ->
      (local accumulator left && local counter right) || (local counter left && local accumulator right)
    _ -> false
  congruence <- CongruenceSum.recognize counter selection.condition
  pure { counter, accumulator, lower: boundary.lower, modulus: congruence.modulus, residue: congruence.residue }

singleBranch :: TcoExpr -> Maybe { condition :: TcoExpr, yes :: TcoExpr, no :: TcoExpr }
singleBranch value = case syntax value of
  Syn.Branch branches no -> case NEA.toArray branches of
    [ Pair condition yes ] -> Just { condition, yes, no }
    _ -> Nothing
  _ -> Nothing

syntax :: TcoExpr -> BackendSyntax TcoExpr
syntax (TcoExpr _ value) = case value of
  Syn.Typed _ inner -> syntax inner
  _ -> value

localName :: TcoExpr -> Maybe String
localName value = case syntax value of
  Syn.Local name level -> Just (localId name level)
  _ -> Nothing

local :: String -> TcoExpr -> Boolean
local name value = localName value == Just name

intLiteral :: TcoExpr -> Maybe Int
intLiteral value = case syntax value of
  Syn.Lit (LitInt number) -> Just number
  _ -> Nothing

positive32 :: Int -> Boolean
positive32 value = value >= 1 && value <= top

cast :: String -> GoExpr -> GoExpr
cast name value = GoCall (GoVar name) [ value ]

integer :: Int -> GoExpr
integer = cast "int64" <<< GoInt
