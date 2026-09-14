module Gopurs.GoFunctions (curriedFunction) where

import Prelude
import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (GoExpr(..), GoType(..))

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
