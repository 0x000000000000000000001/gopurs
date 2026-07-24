module Gopurs.ShowTco where

import Prelude
import Data.Maybe (Maybe(..))
import Data.Array as Array
import Data.String as String
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))

showExprType :: ExprType -> String
showExprType = case _ of
  Int -> "Int"
  Number -> "Number"
  String -> "String"
  Char -> "Char"
  Boolean -> "Boolean"
  Any -> "Any"
  Func args ret -> "Func([" <> String.joinWith "," (map showExprType args) <> "]," <> showExprType ret <> ")"
  _ -> "OtherType"

showTco :: TcoExpr -> String
showTco (TcoExpr _ syn) = case syn of
  Var _ -> "Var"
  Local _ _ -> "Local"
  Lit _ -> "Lit"
  App f args -> "App(" <> showTco f <> "," <> show (Array.length (Array.fromFoldable args)) <> ")"
  Typed t inner -> "Typed(" <> showExprType t <> "," <> showTco inner <> ")"
  _ -> "Other"
