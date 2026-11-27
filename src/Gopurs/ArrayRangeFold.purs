module Gopurs.ArrayRangeFold (emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIntrinsics.Callback as Callback
import Gopurs.CallAnalysis (qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.CongruenceSum as CongruenceSum
import Gopurs.ExprAnalysis (getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.ScalarArrayFold as ScalarArrayFold
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (Level(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type Pipeline =
  { callback :: TcoExpr, seed :: TcoExpr, predicate :: TcoExpr
  , start :: TcoExpr, end :: TcoExpr, range :: TcoExpr, filtered :: TcoExpr
  }

-- The exact Int-add fold has no effects and cannot fail. Interleaving its
-- private accumulator updates with the predicate therefore preserves the
-- predicate's invocation order, even for an opaque/effectful callback. Only
-- this immediately consumed range/filter pipeline loses its arrays.
emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
emit translate context@{ codegenStateRef, modNameStr } nextId fn args = do
  pipeline <- recognize context fn args
  let
    values = [ pipeline.callback, pipeline.seed, pipeline.predicate, pipeline.start, pipeline.end ]
    captured = CallArguments.capture { prefix: "rangeFold_arg_", markUsed: true }
      (\index currentId value -> case index of
        2 -> case Callback.nativeLambda translate context currentId 1 value of
          Just native -> native
          Nothing -> translate (childContext context Nothing) currentId value
        _ -> translate (childContext context Nothing) currentId value)
      nextId values
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    scalar index = coerceGoExpr codegenStateRef modNameStr (arg index) (argType index) TypeInt64
    -- Reconstruct the ordinary pipeline from references to already evaluated
    -- captures. Wide Int values keep rangeImpl's existing overflow/allocation
    -- behavior, without repeating seed, predicate creation or either bound.
    references = Array.mapWithIndex (\index original ->
      let name = Just (Ident ("rangeFoldCaptured" <> show index))
          level = Level (-1)
          TcoExpr analysis _ = original
          local = TcoExpr analysis (Syn.Local name level)
          -- The capture's native callback ABI is already in context.bound.
          -- Reapplying its source Func annotation requests boxed Value and
          -- would erase that ABI before ArrayIntrinsics can call it directly.
          value = case index, argType index of
            2, TypeFunc _ _ -> local
            _, _ -> TcoExpr analysis (Syn.Typed (getExprType original) local)
      in { value, key: localId name level, binding: { name: variableName (arg index), goType: argType index } }) values
    ref index = fromMaybe pipeline.seed (map _.value (Array.index references index))
    fallbackContext = context { bound = Array.foldl (\env entry -> Map.insert entry.key entry.binding env) context.bound references }
    range = replaceArguments pipeline.range [ ref 3, ref 4 ]
    filtered = replaceArguments pipeline.filtered [ ref 2, range ]
    -- ScalarArrayFold recognizes the builtin by name; translate that known
    -- callback as its original capture too, rather than evaluating its getter.
    fallbackTranslate ctx currentId value =
      if named "Data.Semiring" "intAdd" value then
        { stmts: StmtEmpty, expr: arg 0, exprType: argType 0, nextId: currentId }
      else translate ctx currentId value
  fallback <- ScalarArrayFold.emit fallbackTranslate fallbackContext captured.nextId fn [ pipeline.callback, ref 1, filtered ]
  let
    predicate = case argType 2 of
      TypeFunc [ expected ] result ->
        coerceGoExpr codegenStateRef modNameStr (GoCall (arg 2)
          [ coerceGoExpr codegenStateRef modNameStr (GoVar "current") TypeInt64 expected ]) result TypeBool
      _ -> coerceGoExpr codegenStateRef modNameStr
        (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply")
          [ arg 2, GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoVar "current" ] ]) TypeValue TypeBool
    start = scalar 3
    end = scalar 4
    fast = constantSum pipeline.predicate (scalar 1) start end <>
      loop predicate (scalar 1) start end
    slow = [ GoReturn (GoCall (GoFuncBlock [] (flattenStmts fallback.stmts <>
      [ GoReturn (coerceGoExpr codegenStateRef modNameStr fallback.expr fallback.exprType TypeInt64) ]) TypeInt64) []) ]
    admitted = GoBinOp "&&" (signed32 start) (signed32 end)
    expr = GoCall (GoFuncBlock [] [ GoIfElse admitted fast slow ] TypeInt64) []
  pure { stmts: captured.stmts, expr, exprType: TypeInt64, nextId: fallback.nextId }

recognize :: ExprContext -> TcoExpr -> Array TcoExpr -> Maybe Pipeline
recognize context fn args = do
  guard (named "Data.Foldable" "foldlArray" fn)
  Tuple callback (Tuple seed filtered) <- case args of
    [ callback, seed, filtered ] -> Just (Tuple callback (Tuple seed filtered))
    _ -> Nothing
  guard (named "Data.Semiring" "intAdd" callback)
  guard (integerTerm context seed && getExprType filtered == Array Int)
  Tuple predicate range <- uncurried "filterImpl" filtered
  guard (getExprType range == Array Int)
  guard case getExprType predicate of
    Func [ Int ] Boolean -> true
    _ -> false
  Tuple start end <- uncurried "rangeImpl" range
  guard (integerTerm context start && integerTerm context end)
  pure { callback, seed, predicate, start, end, range, filtered }

-- PBO can omit Typed around literals and local references. Literal syntax
-- proves Int directly; a local may use the native parameter representation
-- already established from TAST. An unannotated call still remains unknown.
integerTerm :: ExprContext -> TcoExpr -> Boolean
integerTerm context value
  | getExprType value == Int = true
  | getExprType value /= Any = false
  | otherwise = case unwrapTcoExpr value of
      Syn.Lit (LitInt _) -> true
      Syn.Local name level -> case Map.lookup (localId name level) context.bound of
        Just binding -> binding.goType == TypeInt64
        Nothing -> false
      _ -> false

uncurried :: String -> TcoExpr -> Maybe (Tuple TcoExpr TcoExpr)
uncurried name value = case unwrapTcoExpr value of
  Syn.UncurriedApp fn [ a, b ] | named "Data.Array" name fn -> Just (Tuple a b)
  _ -> Nothing

named :: String -> String -> TcoExpr -> Boolean
named moduleName name value = case qualifiedTarget value of
  Just target -> target.mbMod == Just (ModuleName moduleName) && target.name == name
  Nothing -> false

replaceArguments :: TcoExpr -> Array TcoExpr -> TcoExpr
replaceArguments original@(TcoExpr analysis syntax) args = case syntax of
  Syn.Typed ty inner -> TcoExpr analysis (Syn.Typed ty (replaceArguments inner args))
  Syn.TypeApp inner ty -> TcoExpr analysis (Syn.TypeApp (replaceArguments inner args) ty)
  Syn.UncurriedApp fn _ -> TcoExpr analysis (Syn.UncurriedApp fn args)
  _ -> original

-- capture always returns variable references; keeping that invariant here
-- avoids raw rendering of expressions into the fallback's local environment.
variableName :: GoExpr -> String
variableName (GoVar name) = name
variableName _ = "rangeFold_invalid_capture"

signed32 :: GoExpr -> GoExpr
signed32 value = GoBinOp "==" value
  (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ value ] ])

-- The outer guard has already established signed Int32 endpoints. Only a
-- direct pure lambda permits eliminating predicate calls. Keep all captures
-- above this branch, and retain the fused loop for nonpositive bounds or wide
-- seeds. Its ordering remains observable for every other predicate.
constantSum :: TcoExpr -> GoExpr -> GoExpr -> GoExpr -> Array GoExpr
constantSum predicate seed start end = case recognizePredicate predicate of
  Just congruence ->
    let
      positive = GoBinOp "&&" (GoBinOp ">" start (GoInt 0)) (GoBinOp ">" end (GoInt 0))
      admitted = GoBinOp "&&" positive (signed32 seed)
      lower = GoVar "lower"
      upper = GoVar "upper"
      -- Array.range includes both endpoints in either direction. Int32
      -- addition is commutative and the recognized predicate has no effects,
      -- so normalizing these private bounds preserves the exact result.
      normalize = GoIfElse (GoBinOp ">" lower upper)
        [ GoAssign "swap" lower, GoMutate "lower" upper, GoMutate "upper" (GoVar "swap") ] []
      formula = GoFuncBlock [ Tuple "lower" TypeInt64, Tuple "upper" TypeInt64, Tuple "seed" TypeInt64 ]
        ([ normalize ] <> CongruenceSum.body congruence { lower, upper, seed: GoVar "seed" }) TypeInt64
    in [ GoIfElse admitted [ GoReturn (GoCall formula [ start, end, seed ]) ] [] ]
  Nothing -> []

recognizePredicate :: TcoExpr -> Maybe CongruenceSum.Congruence
recognizePredicate predicate = case unwrapTcoExpr predicate of
  Syn.Abs args body -> case NEA.toArray args of
    [ Tuple name level ] -> CongruenceSum.recognize (localId name level) body
    _ -> Nothing
  _ -> Nothing

loop :: GoExpr -> GoExpr -> GoExpr -> GoExpr -> Array GoExpr
loop predicate seed start end =
  [ GoAssign "current" start
  , GoAssign "last" end
  , GoAssign "step" (GoCall (GoVar "int64") [ GoInt 1 ])
  , GoAssign "total" seed
  , GoIfElse (GoBinOp ">" (GoVar "current") (GoVar "last"))
      [ GoMutate "step" (GoCall (GoVar "int64") [ GoInt (-1) ]) ] []
  , GoFor "range_filter_fold"
      [ GoIfElse predicate
          [ GoMutate "total" (GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd") [ GoVar "total", GoVar "current" ]) ] []
      -- Inclusive endpoints, including descending and singleton ranges. Stop
      -- before stepping beyond either endpoint; the traversal itself uses
      -- int64 arithmetic, while only the accumulated sum wraps as PureScript Int.
      , GoIfElse (GoBinOp "==" (GoVar "current") (GoVar "last")) [ GoReturn (GoVar "total") ] []
      , GoMutate "current" (GoBinOp "+" (GoVar "current") (GoVar "step"))
      ]
  ]
