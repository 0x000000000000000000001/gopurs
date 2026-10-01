module Gopurs.GoFunctions
  ( Parameters
  , namedParameters
  , loopParameters
  , iterationBindings
  , curriedFunction
  ) where

import Prelude
import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, rawGo)

type Parameters = Array (Tuple String GoType)

-- Syntax determines arity; missing type slots are boxed, surplus slots ignored.
namedParameters :: Array String -> Array GoType -> Parameters
namedParameters names types =
  Array.zipWith Tuple names (types <> Array.replicate (Array.length names - Array.length types) TypeValue)

loopParameters :: Parameters -> Parameters
loopParameters = map (\(Tuple name ty) -> Tuple (name <> "_loop") ty)

-- The mutable slots are separate from the source parameters. Rebind inside each
-- iteration so argument permutations and escaping closures see that iteration's
-- values, even after a tail jump writes the next set of slots.
iterationBindings :: Parameters -> Array GoExpr
iterationBindings = Array.concatMap (\(Tuple name ty) ->
  [ rawGo ("var " <> name <> " " <> goTypeToStr ty <> " = " <> name <> "_loop")
  , rawGo ("_ = " <> name)
  ])

-- Preserve the former GoFunc grouping: with more than ten parameters, emit
-- one unary layer at a time until the remaining group fits a runtime FuncN.
curriedFunction :: Array (Tuple String GoType) -> GoType -> GoExpr -> GoExpr
curriedFunction params result body = case Array.uncons params of
  Nothing -> curriedFunction [ Tuple "_" TypeValue ] result body
  Just { head, tail } | Array.length params > 10 ->
    wrap [ head ] (curriedFunction tail result body)
  _ -> wrap params body
  where
  wrap args inner =
    let
      name = if Array.length args == 1 then "Func" else "Func" <> show (Array.length args)
      stmts = case inner of
        GoBlock statements -> statements
        _ -> [ GoReturn inner ]
    in
      GoCall (GoSelector (GoVar "gopurs_runtime") name) [ GoFuncBlock args stmts result ]
