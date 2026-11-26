module Gopurs.AffineLoops (wrap, globalTarget, localTarget) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (any, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.Tuple (Tuple(..))
import Gopurs.ExprContext (ExprContext, LocalEnv, TranslateExpr, childContext, flattenStmts)
import Gopurs.FilteredSumLoops as FilteredSumLoops
import Gopurs.GoAst (GoExpr(..), GoType(..), sanitizeName)
import Gopurs.GoFunctions (Parameters)
import Gopurs.RecordAffineLoops as RecordAffineLoops
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Literal(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type Recurrence = { counter :: String, accumulator :: String, increment :: Int, reads :: Array TcoExpr }

-- Match the complete pure recurrence before changing its execution count. The
-- original loop remains the fallback for zero, negative and wide entry values.
-- Unused global reads retain their first demand in source order. Generated
-- PureScript getters memoize initialization; foreign getters return a binding
-- materialized at package initialization (Printer/FfiBridge). Reading a Var
-- never invokes the function or Effect it may contain. Calls, accessors,
-- callbacks and effects are not admitted as reads.
wrap :: TranslateExpr -> ExprContext -> (TcoExpr -> Boolean) -> Parameters -> GoType -> TcoExpr -> GoExpr -> GoExpr
wrap translate context self params result body original = case recognize self params body of
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
    in GoBlock [ GoIfElse admitted (emitReads translate context recurrence.reads <> [ GoReturn value ]) [], original ]
  _ -> RecordAffineLoops.wrap self params result body
    (FilteredSumLoops.wrap self params result body original)

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
  prefix <- strictReads 64 branch.step
  args <- case syntax prefix.body of
    Syn.App target args | self target -> Just (NEA.toArray args)
    Syn.UncurriedApp target args | self target -> Just args
    _ -> Nothing
  Tuple nextCounter nextAccumulator <- case args of
    [ a, b ] -> Just (if counter == first then Tuple a b else Tuple b a)
    _ -> Nothing
  guard (case syntax nextCounter of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpSubtract) left right) -> local counter left && literal 1 right
    _ -> false)
  constant <- case syntax nextAccumulator of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpAdd) left right)
      | local accumulator left -> Just right
      | local accumulator right -> Just left
    _ -> Nothing
  -- Closed scalar evaluation may retain strict reads inside the constant
  -- operand. The other operand and counter update are already proven pure;
  -- these demands therefore follow the step prefix in either argument order.
  suffix <- strictReads (64 - Array.length prefix.reads) constant
  increment <- intLiteral suffix.body
  guard (increment >= bottom && increment <= top)
  pure { counter, accumulator, increment, reads: prefix.reads <> suffix.reads }

-- Match the whole suffix, including every local level, before discarding a
-- binding. Locals, projections and arbitrary initializers keep the old loop.
strictReads :: Int -> TcoExpr -> Maybe { reads :: Array TcoExpr, body :: TcoExpr }
strictReads remaining expression = case syntax expression of
  Syn.Let ident level value body -> do
    guard (remaining > 0)
    case syntax value of
      Syn.Var _ -> pure unit
      _ -> Nothing
    guard (not (usesLocal ident level body))
    suffix <- strictReads (remaining - 1) body
    pure { reads: Array.cons value suffix.reads, body: suffix.body }
  _ -> Just { reads: [], body: expression }

usesLocal :: Maybe Ident -> Syn.Level -> TcoExpr -> Boolean
usesLocal ident level (TcoExpr _ value) = case value of
  Syn.Local name at -> ident == name && level == at
  _ -> any (usesLocal ident level) value

-- Translate each demand normally so its representation and any strict
-- conversion remain intact. Never inherit the recurrence's Int result type.
emitReads :: TranslateExpr -> ExprContext -> Array TcoExpr -> Array GoExpr
emitReads translate context values =
  (foldl emit { next: 0, statements: [] } values).statements
  where
  emit acc value =
    let result = translate (childContext context Nothing) acc.next value
    in { next: result.nextId
       , statements: acc.statements <> flattenStmts result.stmts <> [ GoMutate "_" result.expr ]
       }

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
