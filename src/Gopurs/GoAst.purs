module Gopurs.GoAst where

import Prelude
import Data.Tuple (Tuple(..))

data GoExpr
  = GoVar String
  | GoString String
  | GoInt Int
  | GoCall GoExpr (Array GoExpr)
  | GoSelector GoExpr String
  | GoFunc String GoExpr
  | GoBlock (Array GoExpr)
  | GoReturn GoExpr
  | GoAssign String GoExpr
  | GoMap (Array (Tuple String GoExpr))
  | GoIIFE String GoExpr GoExpr
  | GoRecordAccess GoExpr String
  | GoBranch (Array (Tuple GoExpr GoExpr)) GoExpr
  | GoBinOp String GoExpr GoExpr
  | GoTypeAssertion GoExpr String
  | GoRaw String

derive instance eqGoExpr :: Eq GoExpr

type GoDecl =
  { identifier :: String
  , expression :: GoExpr
  }

type GoFile =
  { packageName :: String
  , imports :: Array String
  , decls :: Array GoDecl
  }
