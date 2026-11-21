module Gopurs.AffineLoops (wrap, globalTarget, localTarget) where

import Prelude

import Control.Alternative (guard)
import Data.Array.NonEmpty as NEA
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.Tuple (Tuple(..))
import Gopurs.ExprContext (LocalEnv)
import Gopurs.GoAst (GoExpr(..), GoType(..), sanitizeName)
import Gopurs.GoFunctions (Parameters)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Literal(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type Recurrence = { counter :: String, accumulator :: String, increment :: Int }

-- Match the complete pure recurrence before changing its execution count. The
-- original loop remains the fallback for zero, negative and wide entry values.
-- No calls other than the exact tail self-call, callbacks or effects are admitted.
wrap :: (TcoExpr -> Boolean) -> Parameters -> GoType -> TcoExpr -> GoExpr -> GoExpr
wrap self params result body original = case recognize self params body of
  Just recurrence | result == TypeInt64 ->
    let
      counter = GoVar (recurrence.counter <> "_loop")
      accumulator = GoVar (recurrence.accumulator <> "_loop")
      positive = GoBinOp ">" counter (integer 0)
      bounded = GoBinOp "<=" counter (integer top)
      signed = GoBinOp "==" accumulator (cast "int64" (cast "int32" accumulator))
      admitted = GoBinOp "&&" (GoBinOp "&&" positive bounded) signed
      -- Both factors are signed 32-bit, with a positive counter. Their product
      -- plus the signed accumulator fits int64. IntMul intentionally has JS
      -- Number rounding semantics, so it cannot implement this exact sum.
      total = GoBinOp "+" accumulator (GoBinOp "*" counter (integer recurrence.increment))
      value = cast "int64" (cast "int32" total)
    in GoBlock [ GoIfElse admitted [ GoReturn value ] [], original ]
  _ -> original

recognize :: (TcoExpr -> Boolean) -> Parameters -> TcoExpr -> Maybe Recurrence
recognize self params body = do
  Tuple first second <- case params of
    [ Tuple first TypeInt64, Tuple second TypeInt64 ] | first /= second -> Just (Tuple first second)
    _ -> Nothing
  branch <- case syntax body of
    Syn.Branch branches step -> case NEA.toArray branches of
      [ Pair condition seed ] -> Just { condition, seed, step }
      _ -> Nothing
    _ -> Nothing
  counter <- case syntax branch.condition of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntOrd Syn.OpEq) left right)
      | literal 0 right -> localName left
      | literal 0 left -> localName right
    _ -> Nothing
  guard (counter == first || counter == second)
  let accumulator = if counter == first then second else first
  guard (local accumulator branch.seed)
  args <- case syntax branch.step of
    Syn.App target args | self target -> Just (NEA.toArray args)
    Syn.UncurriedApp target args | self target -> Just args
    _ -> Nothing
  Tuple nextCounter nextAccumulator <- case args of
    [ a, b ] -> Just (if counter == first then Tuple a b else Tuple b a)
    _ -> Nothing
  guard (case syntax nextCounter of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpSubtract) left right) -> local counter left && literal 1 right
    _ -> false)
  increment <- case syntax nextAccumulator of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpAdd) left right)
      | local accumulator left -> intLiteral right
      | local accumulator right -> intLiteral left
    _ -> Nothing
  guard (increment >= bottom && increment <= top)
  pure { counter, accumulator, increment }

syntax :: TcoExpr -> BackendSyntax TcoExpr
syntax (TcoExpr _ value) = case value of
  Syn.Typed _ inner -> syntax inner
  _ -> value

localName :: TcoExpr -> Maybe String
localName expr = case syntax expr of
  Syn.Local ident level -> Just (localId ident level)
  _ -> Nothing

local :: String -> TcoExpr -> Boolean
local name expr = localName expr == Just name

intLiteral :: TcoExpr -> Maybe Int
intLiteral expr = case syntax expr of
  Syn.Lit (LitInt value) -> Just value
  _ -> Nothing

literal :: Int -> TcoExpr -> Boolean
literal value expr = intLiteral expr == Just value

globalTarget :: String -> String -> TcoExpr -> Boolean
globalTarget moduleName name expr = case syntax expr of
  Syn.Var (Qualified (Just owner) (Ident ident)) ->
    String.replaceAll (String.Pattern ".") (String.Replacement "_") (unwrap owner) == moduleName
      && sanitizeName ident == name
  _ -> false

localTarget :: LocalEnv -> String -> TcoExpr -> Boolean
localTarget bound name expr = case localName expr >>= flip Map.lookup bound of
  Just binding -> binding.name == name
  Nothing -> false

cast :: String -> GoExpr -> GoExpr
cast name value = GoCall (GoVar name) [ value ]

integer :: Int -> GoExpr
integer = cast "int64" <<< GoInt
